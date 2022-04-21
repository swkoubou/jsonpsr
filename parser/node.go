package parser

import "github.com/swkoubou/jsonpsr/tokenizer"

type Node struct {
	tokens []tokenizer.Token
	child  *Node
}
