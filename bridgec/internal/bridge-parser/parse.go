package bridgeparser

import (
	parser "bridgelang.dev/bridgec-libs/antlr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func Parse(source antlr.CharStream) *parser.BridgeParser {
	blex := parser.NewBridgeLexer(source)
	bparse := parser.NewBridgeParser(antlr.NewCommonTokenStream(blex, antlr.TokenDefaultChannel))

	return bparse
}
