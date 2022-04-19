package tokenizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTokenizer(t *testing.T) {
	in := "hello"
	tokenizer := NewTokenizer(in)
	tokens := tokenizer.Tokenize()
	tokensStr := ""
	for _, token := range tokens {
		tokensStr += token.Raw
	}
	assert.Equal(t, in, tokensStr)
}
