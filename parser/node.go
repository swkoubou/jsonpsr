package parser

import "github.com/swkoubou/jsonpsr/tokenizer"

type Node struct {
	Kind
	tokens   []tokenizer.Token
	children []*Node
}
