package tokenizer

import "fmt"

type Token struct {
	Kind
	Raw string
	S   int
	E   int
}

func (t Token) String() string {
	// position S:E, S<=...<E, like python
	var position string
	if t.E-t.S == 1 {
		position = fmt.Sprintf("%v", t.S)
	} else {
		position = fmt.Sprintf("%v:%v", t.S, t.E)
	}
	return fmt.Sprintf("Token[%v]\t{ Kind: %v,\tRaw: `%v` }", position, t.Kind.String(), t.Raw)
}
