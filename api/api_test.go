package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/swkoubou/jsonpsr/parser"
	"github.com/swkoubou/jsonpsr/tokenizer"
	"testing"
)

func TestApi_Load(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect string
	}{
		{
			"kv string",
			`{"key": "value"}`,
			"value",
		},
	}

	tz := tokenizer.NewTokenizer()
	ps := parser.NewParser()
	api := NewApi()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := tz.Tokenize(tt.in)
			node := ps.Parse(tokens)
			if err := api.Load(node); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expect, "")
		})
	}
}
