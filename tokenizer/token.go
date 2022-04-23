package tokenizer

import "fmt"

type Token struct {
	Kind
	Raw string
	s   int
	e   int
}

func (t Token) String() string {
	// position s:e, s<=...<e, like python
	var position string
	if t.e-t.s == 1 {
		position = fmt.Sprintf("%v", t.s)
	} else {
		position = fmt.Sprintf("%v:%v", t.s, t.e)
	}
	return fmt.Sprintf("Token[%v]\t{ Kind: %v,\tRaw: `%v` }", position, t.Kind.String(), t.Raw)
}
