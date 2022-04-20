package parser

import "github.com/swkoubou/jsonpsr/tokenizer"

type Parser struct {
	input []tokenizer.Token
	pos   int
}

func NewParser() *Parser {
	// 結局この関数New~が必要なのかはわからない。
	return &Parser{
		input: nil,
		pos:   0,
	}
}

func (p *Parser) Parse() {}
