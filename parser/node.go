package parser

import "github.com/swkoubou/jsonpsr/tokenizer"

type Node struct {
	Kind
	Key      string
	tokens   []tokenizer.Token
	children []*Node
}

func NewNode(kind Kind, key string, tokens []tokenizer.Token, children []*Node) *Node {
	return &Node{
		kind,
		key,
		tokens,
		children,
	}
}

func NewChildren(children ...*Node) []*Node {
	return children
}

func NewJsonElemValNode(children []*Node) *Node {
	return NewNode(
		JSON,
		"",
		nil,
		NewChildren(NewElementValueNode(children)),
	)
}

func NewElementValueNode(children []*Node) *Node {
	return NewNode(
		ELEMENT,
		"",
		nil,
		NewChildren(
			NewNode(
				VALUE,
				"",
				nil,
				children,
			)),
	)
}

func NewElemValString(tokens ...tokenizer.Token) *Node {
	return NewElementValueNode(
		NewChildren(
			NewNode(
				STRING,
				"",
				tokens,
				nil,
			)))
}

func NewElemValNumber(tokens ...tokenizer.Token) *Node {
	return NewElementValueNode(
		NewChildren(
			NewNode(
				NUMBER,
				"",
				tokens,
				nil,
			)))
}

func NewElemValTrue(tokens ...tokenizer.Token) *Node {
	return NewElementValueNode(
		NewChildren(
			NewNode(
				TRUE,
				"",
				tokens,
				nil,
			)))
}

func NewElemValFalse(tokens ...tokenizer.Token) *Node {
	return NewElementValueNode(
		NewChildren(
			NewNode(
				FALSE,
				"",
				tokens,
				nil,
			)))
}

func NewElemValNull(tokens ...tokenizer.Token) *Node {
	return NewElementValueNode(
		NewChildren(
			NewNode(
				NULL,
				"",
				tokens,
				nil,
			)))
}
