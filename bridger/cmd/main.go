package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"bridgelang.dev/bridgelib"
	"bridgelang.dev/bridgelib/data"
)

type FlagCollection struct {
	q  *bool
	fc *bool
	mC *int64
}

var Flags *FlagCollection

func main() {
	versioninfo := flag.Bool("version", false, "--version: Get version info")
	versioninfojson := flag.Bool("versionjson", false, "--versionjson: Get version info in JSON format")
	Flags = &FlagCollection{flag.Bool("q", false, "-q: Quiet, disable prints"),
		flag.Bool("fromcin", false, "--fromcin: Get file from std console input"),
		flag.Int64("maxmemory", 0, "--maxmemory: Max pages allocable by the VM")}
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
		fmt.Printf("Bridge Runner\nBridge Version: v%s (%s)\nChannel: %s\n", data.VERSIONSTR, data.VERSIONNAME, data.CHANNEL)
		os.Exit(0)
	}
	if !*Flags.q {
		fmt.Printf("Bridge Runner v%s\nIssue `bridger --version` for more information on this build\n", data.VERSIONSTR)
	}

	var rawBytes []byte
	if source := flag.Arg(0); source != "" && !*Flags.fc {
		var err error
		rawBytes, err = os.ReadFile(source)
		if err != nil {
			panic(err)
		}
	} else {
		if *Flags.fc {
			in := bufio.NewReader(os.Stdin)
			var err error
			rawBytes, err = io.ReadAll(in)
			if err != nil {
				panic(err)
			}
		} else {
			panic(errors.New("console input disabled while no source file was provided"))
		}
	}

	bbf, err := bridgelib.Deserialize(rawBytes)
	if err != nil {
		panic(err)
	}

	if bbf.FormatVersion != data.BBVERSION {
		fmt.Print("Warning: Different BridgeBytecode format version\n")
		if (bbf.FormatVersion & 0xFF000000) != (data.BBVERSION & 0xFF000000) {
			fmt.Printf("Error: BB format has a different major: %b != %b\n", (bbf.FormatVersion & 0xFF000000), (data.BBVERSION & 0xFF000000))
			os.Exit(-1)
		}
	}

	bc, err := bbf.InflateBytecode()
	if err != nil {
		panic(err)
	}

	deps, err := bridgelib.LoadDependencies(bbf.ImportsTable)
	if err != nil {
		panic(err)
	}

	inst, err := bridgelib.New(bc, bridgelib.BridgeContext{}, deps)
	if err != nil {
		panic(err)
	}
	defer inst.Close()

	nparams, err := inst.FuncParamCount("Main")
	if err != nil {
		panic(err)
	}

	var callErr error
	if nparams == 0 {
		_, callErr = inst.Call("Main")
	} else {
		arrPtr, count, merr := inst.WriteStrings(flag.Args()[1:])
		if merr != nil {
			panic(merr)
		}

		if nparams == 1 {
			_, callErr = inst.Call("Main", arrPtr)
		} else {
			_, callErr = inst.Call("Main", arrPtr, count)
		}
	}
	if callErr != nil {
		panic(callErr)
	}
}
