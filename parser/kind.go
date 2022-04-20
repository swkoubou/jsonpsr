package parser

// 参考
// - 仕様書 EN: https://www.json.org/json-en.html
// - 仕様書 JA: https://www.json.org/json-ja.html
// - 仕様書(?): https://datatracker.ietf.org/doc/html/rfc8259
// - EBNF: https://www.sigbus.info/compilerbook#単純な生成規則

type Kind int

const (
	ILLEGAL    Kind = iota
	OBJECT          // {}
	ARRAY           // []
	VALUE           // ...
	STRING          // ""
	NUMBER          // 123
	TRUE            // true
	FALSE           // false
	NULL            // null
	WHITESPACE      // " "
)

// ^\s+?([A-Z]*)\s+//(.*)\n?$
var kinds = [...]string{
	ILLEGAL:    "ILLEGAL",
	OBJECT:     "OBJECT",
	ARRAY:      "ARRAY",
	VALUE:      "VALUE",
	STRING:     "STRING",
	NUMBER:     "NUMBER",
	TRUE:       "TRUE",
	FALSE:      "FALSE",
	NULL:       "NULL",
	WHITESPACE: "WHITESPACE",
}

func (k Kind) String() string {
	return kinds[k]
}
