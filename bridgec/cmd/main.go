package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	bridgeparser "bridgelang.dev/bridgec/internal/bridge-parser"
	"bridgelang.dev/bridgelib"
	"bridgelang.dev/bridgelib/data"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"tinygo.org/x/go-llvm"
)

type FlagCollection struct {
	q     *bool
	fc    *bool
	dllvm *bool
	pa    *bool
}

var Flags *FlagCollection

func main() {
	versioninfo := flag.Bool("version", false, "--version: Get version info")
	versioninfojson := flag.Bool("versionjson", false, "--versionjson: Get version info in JSON format")
	Flags = &FlagCollection{flag.Bool("q", false, "-q: Quiet, disable prints"),
		flag.Bool("fromcin", false, "--fromcin: Get file from std console input"),
		flag.Bool("dumpllvm", false, "--dumpllvm: Dump LLVM Module for inspection"),
		flag.Bool("printast", false, "--printast: Print all AST nodes")}
	flag.Parse()
	if *versioninfojson {
		output := make(map[string]string)
		output["version"] = data.VERSIONSTR
		output["codename"] = data.VERSIONNAME
		output["channel"] = data.CHANNEL
		out, err := json.Marshal(output)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", string(out))
		os.Exit(0)
	}
	if *versioninfo {
		fmt.Printf("Bridge Compiler\nBridge Version: v%s (%s)\nChannel: %s\n", data.VERSIONSTR, data.VERSIONNAME, data.CHANNEL)
		os.Exit(0)
	}
	if !*Flags.q {
		fmt.Printf("Bridge Compiler v%s\nIssue `bridgec --version` for more information on this build\n", data.VERSIONSTR)
	}

	if source := flag.Arg(0); source != "" && !*Flags.fc {
		stream, err := antlr.NewFileStream(source)
		if err != nil {
			panic(err)
		}
		parseStep(stream, strings.ReplaceAll(path.Base(strings.ReplaceAll(source, path.Ext(source), "")), ".", ""))
	} else {
		if *Flags.fc {
			in := bufio.NewReader(os.Stdin)
			var err error
			rawBytes, err := io.ReadAll(in)
			if err != nil {
				panic(err)
			}
			parseStep(string(rawBytes), "program")
			return
		}
		panic(errors.New("console input disabled while no source file was provided"))
	}

	fmt.Printf("Done!\n")
}

func parseStep(stream any, modname string) {
	if stream == nil {
		panic(errors.New("input stream was nil"))
	}
	var cstream antlr.CharStream
	switch stream := stream.(type) {
	case *antlr.FileStream:
		cstream = stream
	default:
		cstream = antlr.NewInputStream(stream.(string))
	}
	llvm.InitializeAllTargetInfos()
	llvm.InitializeAllTargets()
	llvm.InitializeAllTargetMCs()
	llvm.InitializeAllAsmPrinters()
	bp := bridgeparser.Parse(cstream)
	bp.BuildParseTrees = true
	tree := bp.Start()
	llvmctx := llvm.NewContext()
	visitor := &bridgeparser.Visitor{PrintAst: *Flags.pa, Llvmctx: llvmctx, Llvmmod: llvmctx.NewModule(modname), Builder: llvmctx.NewBuilder()}
	defer visitor.Llvmmod.Dispose()
	visitor.MakeVisitor()
	visitor.Visit(tree)

	if err := llvm.VerifyModule(visitor.Llvmmod, llvm.VerifierFailureAction(llvm.ReturnStatusAction)); err != nil {
		panic(err)
	}

	if *Flags.dllvm {
		visitor.Llvmmod.Dump()
	}

	target, err := llvm.GetTargetFromTriple("wasm64-unknown-unknown")
	if err != nil {
		panic(err)
	}

	machine := target.CreateTargetMachine("wasm64-unknown-unknown", "", "", llvm.CodeGenOptLevel(llvm.CodeGenLevelAggressive), llvm.RelocMode(llvm.RelocDefault), llvm.CodeModel(llvm.CodeModelJITDefault))
	defer machine.Dispose()

	buffer, err := machine.EmitToMemoryBuffer(visitor.Llvmmod, llvm.CodeGenFileType(llvm.ObjectFile))
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	deflater, err := gzip.NewWriterLevel(&buf, 9)
	if err != nil {
		panic(err)
	}

	if _, err = deflater.Write(buffer.Bytes()); err != nil {
		panic(err)
	}
	if err = deflater.Close(); err != nil {
		panic(err)
	}

	importTable := make(map[string]string)
	for name, version := range visitor.LibDeps {
		importTable[name] = version
	}
	iTableBytes, err := json.Marshal(importTable)
	if err != nil {
		panic(err)
	}

	bbf := &bridgelib.BBFile{
		FormatVersion: data.BBVERSION,
		ImportsTable:  string(iTableBytes),
		ByteCode:      buf.Bytes(),
	}
	encoded, err := bridgelib.Serialize(bbf)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(fmt.Sprintf("%s.bb", modname), encoded, 0644)
	if err != nil {
		panic(err)
	}
}
