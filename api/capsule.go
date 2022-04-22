package api

import "github.com/swkoubou/jsonpsr/parser"

type Capsule struct {
	node *parser.Node
}

func (c *Capsule) AGet(index int) (*Capsule, error) {
	if c.node.Kind != parser.ARRAY {
	}
	return nil, nil
}
func (c *Capsule) OGet(key string) {}
