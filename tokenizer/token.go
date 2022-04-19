package tokenizer

type Token struct {
	Kind
	Raw string
}

//func NewToken(k Kind, r string) *Token {
//	return &Token{
//		k,
//		r,
//	}
//}
