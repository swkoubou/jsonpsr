package parser

import "github.com/swkoubou/jsonpsr/tokenizer"

type Node struct {
	Kind
	tokens   []tokenizer.Token
	children []*Node
}

func NewNode(kind Kind, tokens []tokenizer.Token, children []*Node) *Node {
	return &Node{
		kind,
		tokens,
		children,
	}
}
func NewChildrenNode(children ...*Node) []*Node {
	return children
}
