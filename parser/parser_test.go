package parser

import (
	"github.com/stretchr/testify/assert"
	"github.com/swkoubou/jsonpsr/tokenizer"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name   string
		in     []tokenizer.Token
		expect *Node
	}{
		{
			"only true",
			[]tokenizer.Token{
				{
					tokenizer.KEYWORD,
					"true",
				},
			},
			&Node{
				TRUE,
				[]tokenizer.Token{
					{
						tokenizer.KEYWORD,
						"true",
					},
				},
				nil,
			},
		},
	}

	psr := NewParser()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := psr.Parse(tt.in); !assert.Equal(t, tt.expect, actual) {
				t.Errorf("expected=%v, but actual=%v", tt.expect, actual)
			}
		})
	}
}
