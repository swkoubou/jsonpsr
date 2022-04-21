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

// トークナイザと同じ。走査(?)に必要な基本機能。
func (p *Parser) curt() tokenizer.Token {
	return p.input[p.pos]
}
func (p *Parser) prev() tokenizer.Token {
	return p.input[p.pos-1]
}
func (p *Parser) next() tokenizer.Token {
	return p.input[p.pos+1]
}
func (p *Parser) goNext() {
	p.pos++
}
func (p *Parser) isEof() bool {
	return p.pos >= len(p.input)
}

func (p *Parser) Parse() []*Node {
	return nil
}

func (p *Parser) json() *Node {
	// json = element
	return p.element()
}
func (p *Parser) element() *Node {
	// element = value
	return p.value()
}
func (p *Parser) value() *Node {
	// value = object | array | string | number | "true" | "false" | "null"
	return nil
}

func (p *Parser) object() *Node {
	// object = "{" members? "}"
	return nil
}
func (p *Parser) members() *Node {
	// members = member ("," member)*
	return nil
}
func (p *Parser) member() *Node {
	// member = string ":" element
	return nil
}

func (p *Parser) array() *Node {
	// array = "[" elements? "]"
	return nil
}
func (p *Parser) elements() *Node {
	// elements = element ("," element)*
	return nil
}

func (p *Parser) string() *Node {
	return nil
}
func (p *Parser) number() *Node {
	return nil
}
func (p *Parser) true() *Node {
	return nil
}
func (p *Parser) false() *Node {
	return nil
}
func (p *Parser) null() *Node {
	return nil
}
