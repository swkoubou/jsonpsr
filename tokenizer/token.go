package tokenizer

import "fmt"

type Token struct {
	Kind
	Raw string
}

func (t Token) String() string {
	return fmt.Sprintf("Token { Kind: %v, Raw: `%v` }", t.Kind.String(), t.Raw)
}
