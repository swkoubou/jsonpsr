package tokenizer

type Tokenizer struct {
	input []rune
	pos   int
}

func NewTokenizer(in string) *Tokenizer {
	return &Tokenizer{
		input: []rune(in),
		pos:   0,
	}
}

func (t *Tokenizer) curt() rune {
	return t.input[t.pos]
}
func (t *Tokenizer) next() rune {
	return t.input[t.pos+1]
}

func (t *Tokenizer) goNext() {
	t.pos++
}

func (t *Tokenizer) isEof() bool {
	return t.pos >= len(t.input)
}

func (t *Tokenizer) Tokenize() []Token {
	var result []Token
	for !t.isEof() {
		result = append(result, *NewToken(ILLEGAL, string(t.curt())))
		t.goNext()
	}
	return result
}
