package parser

// 参考
// - 仕様書 EN: https://www.json.org/json-en.html
// - 仕様書 JA: https://www.json.org/json-ja.html
// - 仕様書(?): https://datatracker.ietf.org/doc/html/rfc8259
// - EBNF: https://www.sigbus.info/compilerbook#単純な生成規則

/*
json syntax ebnf?

json = element
element = value
value = object | array | string | number | "true" | "false" | "null"

object = "{" members? "}"
members = member ("," member)*
member = string ":" element

array = "[" elements? "]"
elements = element ("," element)*

*/

type Kind int

const (
	ILLEGAL Kind = iota
	JSON
	VALUE
	STRING
	NUMBER
	TRUE
	FALSE
	NULL
	OBJECT
	MEMBERS
	MEMBER
	ARRAY
	ELEMENTS
	ELEMENT
)

// ^\s+?([A-Z]*)\s+//(.*)\n?$
// ^\s*?([A-Z]*)\n?$
// $1: "$1",
var kinds = [...]string{
	ILLEGAL:  "ILLEGAL",
	JSON:     "JSON",
	VALUE:    "VALUE",
	STRING:   "STRING",
	NUMBER:   "NUMBER",
	TRUE:     "TRUE",
	FALSE:    "FALSE",
	NULL:     "NULL",
	OBJECT:   "OBJECT",
	MEMBERS:  "MEMBERS",
	MEMBER:   "MEMBER",
	ARRAY:    "ARRAY",
	ELEMENTS: "ELEMENTS",
	ELEMENT:  "ELEMENT",
}

func (k Kind) String() string {
	return kinds[k]
}
