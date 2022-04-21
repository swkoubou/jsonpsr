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
