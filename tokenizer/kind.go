package tokenizer

type Kind int

const (
	ILLEGAL    Kind = iota
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
