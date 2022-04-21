package parser

import (
	"github.com/swkoubou/jsonpsr/tokenizer"
	"log"
)

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
func (p *Parser) consume() tokenizer.Token {
	t := p.curt()
	p.goNext()
	return t
}
func (p *Parser) goNext() {
	p.pos++
}
func (p *Parser) isEof() bool {
	return p.pos >= len(p.input)
}

func (p *Parser) Parse(tokens []tokenizer.Token) *Node {
	p.input = tokens
	p.pos = 0
	return p.json()
}

func (p *Parser) json() *Node {
	// json = element
	return NewNode(
		JSON,
		"",
		nil,
		NewChildren(p.element()),
	)
}

func (p *Parser) element() *Node {
	// element = value
	return NewNode(
		ELEMENT,
		"",
		nil,
		NewChildren(p.value()),
	)
}

func (p *Parser) value() *Node {
	// value = object | array | string | number | "true" | "false" | "null"
	switch p.curt().Kind {
	case tokenizer.LCUB:
		// {
		return NewNode(
			VALUE,
			"",
			nil,
			NewChildren(p.object()),
		)
	case tokenizer.LSQB:
		// [
		return NewNode(
			VALUE,
			"",
			nil,
			NewChildren(p.array()),
		)
	case tokenizer.STRING:
		return NewNode(
			VALUE,
			"",
			nil,
			NewChildren(p.string()),
		)
	case tokenizer.NUMBER:
		return NewNode(
			VALUE,
			"",
			nil,
			NewChildren(p.number()),
		)
	case tokenizer.KEYWORD:
		// "true", "false", "null"
		switch p.curt().Raw {
		case "true":
			return NewNode(
				VALUE,
				"",
				nil,
				NewChildren(p.true()),
			)
		case "false":
			return NewNode(
				VALUE,
				"",
				nil,
				NewChildren(p.false()),
			)
		case "null":
			return NewNode(
				VALUE,
				"",
				nil,
				NewChildren(p.null()),
			)
		}
	}
	// or panic
	return &Node{
		ILLEGAL,
		"",
		nil,
		nil,
	}
}

func (p *Parser) object() *Node {
	// object = "{" members? "}"
	obj := NewNode(
		OBJECT,
		"",
		nil,
		nil,
	)
	p.goNext() // "{"
	if p.curt().Kind != tokenizer.RCUB {
		// not empty
		obj.children = NewChildren(p.members())
	}
	p.goNext() // "}"
	return obj
}

func (p *Parser) members() *Node {
	// members = member ("," member)*
	var children []*Node
	// 一個は必ずある。
	children = append(children, p.member())
	// ("," member)*は、0個以上なので、あるかわからない。
	for p.curt().Kind == tokenizer.COMMA {
		p.goNext() // ","
		children = append(children, p.member())
	}
	return NewNode(
		MEMBERS,
		"",
		nil,
		children,
	)
}

func (p *Parser) member() *Node {
	// member = string ":" element
	key := p.consume()
	if key.Kind != tokenizer.STRING {
		// err
		log.Fatalf("<member> expected token.string but found %v", key.String())
	}
	if colon := p.consume(); colon.Kind != tokenizer.COLON {
		// err
		log.Fatalf("<member> expected token.colon but found %v", key.String())
	}
	return NewNode(
		MEMBER,
		key.Raw,
		nil,
		NewChildren(p.element()),
	)
}

func (p *Parser) array() *Node {
	// array = "[" elements? "]"
	p.goNext() // "["
	arr := NewNode(
		ARRAY,
		"",
		nil,
		nil,
	)
	if p.curt().Kind != tokenizer.RSQB {
		arr.children = NewChildren(p.elements())
	}
	p.goNext() // "]"
	return arr
}

func (p *Parser) elements() *Node {
	// elements = element ("," element)*
	var elems []*Node
	elems = append(elems, p.element())
	for p.curt().Kind == tokenizer.COMMA {
		p.goNext() // ","
		elems = append(elems, p.element())
	}

	return NewNode(
		ELEMENTS,
		"",
		nil,
		elems,
	)
}

func (p *Parser) string() *Node {
	t := p.consume()
	return NewNode(
		STRING,
		"",
		[]tokenizer.Token{t},
		nil,
	)
}

func (p *Parser) number() *Node {
	t := p.consume()
	return NewNode(
		NUMBER,
		"",
		[]tokenizer.Token{t},
		nil,
	)
}

func (p *Parser) true() *Node {
	t := p.consume()
	return &Node{
		TRUE,
		"",
		[]tokenizer.Token{t},
		nil,
	}
}

func (p *Parser) false() *Node {
	t := p.consume()
	return &Node{
		FALSE,
		"",
		[]tokenizer.Token{t},
		nil,
	}
}

func (p *Parser) null() *Node {
	t := p.consume()
	return &Node{
		NULL,
		"",
		[]tokenizer.Token{t},
		nil,
	}
}
