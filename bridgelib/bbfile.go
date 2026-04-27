package bridgelib

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

var bbMagic = [4]byte{0x69, 0x42, 0x0B, 0xBF}

type BBFile struct {
	FormatVersion uint32
	ImportsTable  string
	ByteCode      []byte
}

// Serialize serialises f into the BB binary format
//
// Layout:
//
//	    [4]  magic number
//		[4]  version
//		[4]  imports len
//		[n]  imports bytes
//		[4]  bytecode len
//		[n]  bytecode
func Serialize(f *BBFile) ([]byte, error) {
	var buf bytes.Buffer

	buf.Write(bbMagic[:])

	if err := binary.Write(&buf, binary.LittleEndian, f.FormatVersion); err != nil {
		return nil, err
	}

	importBytes := []byte(f.ImportsTable)
	if len(importBytes) > 0xFFFFFFFF {
		return nil, errors.New("bbfile: module name too long")
	}
	if err := binary.Write(&buf, binary.LittleEndian, uint32(len(importBytes))); err != nil {
		return nil, err
	}
	buf.Write(importBytes)

	if err := binary.Write(&buf, binary.LittleEndian, uint32(len(f.ByteCode))); err != nil {
		return nil, err
	}
	buf.Write(f.ByteCode)

	return buf.Bytes(), nil
}

func (f *BBFile) InflateBytecode() ([]byte, error) {
	gr, err := gzip.NewReader(bytes.NewReader(f.ByteCode))
	if err != nil {
		return nil, fmt.Errorf("bbfile: inflate: %w", err)
	}
	return io.ReadAll(gr)
}

// Deserialize deserialises data from a BB binary
func Deserialize(data []byte) (*BBFile, error) {
	r := bytes.NewReader(data)

	var magic [4]byte
	if _, err := io.ReadFull(r, magic[:]); err != nil {
		return nil, fmt.Errorf("bbfile: read magic: %w", err)
	}
	if !bytes.Equal(magic[:], bbMagic[:]) {
		return nil, errors.New("bbfile: invalid magic")
	}

	var fmtVer uint32
	if err := binary.Read(r, binary.LittleEndian, &fmtVer); err != nil {
		return nil, fmt.Errorf("bbfile: read format version: %w", err)
	}

	var importLen uint32
	if err := binary.Read(r, binary.LittleEndian, &importLen); err != nil {
		return nil, fmt.Errorf("bbfile: read name length: %w", err)
	}
	importBytes := make([]byte, importLen)
	if _, err := io.ReadFull(r, importBytes); err != nil {
		return nil, fmt.Errorf("bbfile: read module name: %w", err)
	}

	var bcLen uint32
	if err := binary.Read(r, binary.LittleEndian, &bcLen); err != nil {
		return nil, fmt.Errorf("bbfile: read bytecode length: %w", err)
	}
	bCode := make([]byte, bcLen)
	if _, err := io.ReadFull(r, bCode); err != nil {
		return nil, fmt.Errorf("bbfile: read bytecode: %w", err)
	}

	return &BBFile{
		FormatVersion: fmtVer,
		ImportsTable:  string(importBytes),
		ByteCode:      bCode,
	}, nil
}
