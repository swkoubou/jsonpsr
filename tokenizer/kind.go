package tokenizer

type Kind int

const (
	_          Kind = iota
	LCUB            // {
	RCUB            // }
	LSQB            // [
	RSQB            // ]
	COLON           // :
	COMMA           // ,
	STRING          // "abc"
	NUMBER          // 123, 12.3, -12, ...
	KEYWORD         // true, false, null, ...
	WHITESPACE      // \n, \t, " ", ...
)

var kinds = [...]string{
	LCUB:       "LCUB",
	RCUB:       "RCUB",
	LSQB:       "LSQB",
	RSQB:       "RSQB",
	COLON:      "COLON",
	COMMA:      "COMMA",
	STRING:     "STRING",
	NUMBER:     "NUMBER",
	KEYWORD:    "KEYWORD",
	WHITESPACE: "WHITESPACE",
}

func (k Kind) String() string {
	return kinds[k]
}
