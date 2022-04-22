package api

import (
	"fmt"
	"github.com/swkoubou/jsonpsr/parser"
	"log"
)

type Api struct {
	original *parser.Node
	jev      *parser.Node // json-element-value
}

func NewApi() *Api {
	return &Api{}
}

func (a *Api) Load(n *parser.Node) error {
	a.original = n
	if n.Kind == parser.JSON && n.Children[0].Kind == parser.ELEMENT && n.Children[0].Children[0].Kind == parser.VALUE {
		a.jev = n.Children[0].Children[0].Children[0]
	}
	if a.jev == nil {
		return fmt.Errorf("expected json-elemnt-value, but found %v", a.original.Kind.String())
	}
	if a.jev.Kind != parser.OBJECT && a.jev.Kind != parser.ARRAY {
		log.Fatalf("expected object or array, but found %v", a.jev.Kind.String())
	}
	return nil
}

func (a *Api) AGet(index int) *parser.Node {
	return nil
}

//func (a *Api) Get(path string) any {
//	pathSplit := strings.Split(path, "/")
//
//
//	return nil
//}
