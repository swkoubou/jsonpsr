package tokenizer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenizer_Tokenize(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect []Token
	}{
		{
			"only string",
			"\"This is String\"",
			[]Token{
				{STRING, "This is String", 0, 16},
			},
		},
		{
			"one left curly bracket",
			"{",
			[]Token{
				{LCUB, "{", 0, 1},
			},
		},
		{
			"KV string",
			"{\"key\":\"value\"}",
			[]Token{
				{LCUB, "{", 0, 1},
				{STRING, "key", 1, 6},
				{COLON, ":", 6, 7},
				{STRING, "value", 7, 14},
				{RCUB, "}", 14, 15},
			},
		},
		{
			"KV string include escaped double quotation",
			`{"key":"val\"ue"}`,
			[]Token{
				{LCUB, "{", 0, 1},
				{STRING, "key", 1, 6},
				{COLON, ":", 6, 7},
				{STRING, `val\"ue`, 7, 16},
				{RCUB, "}", 16, 17},
			},
		},
		{
			"KV number(int)",
			`{"key":12}`,
			[]Token{
				{LCUB, "{", 0, 1},
				{STRING, "key", 1, 6},
				{COLON, ":", 6, 7},
				{NUMBER, `12`, 7, 9},
				{RCUB, "}", 9, 10},
			},
		},
		{
			"KV number(string)",
			`{"key":"12"}`,
			[]Token{
				{LCUB, "{", 0, 1},
				{STRING, "key", 1, 6},
				{COLON, ":", 6, 7},
				{STRING, `12`, 7, 11},
				{RCUB, "}", 11, 12},
			},
		},
		{
			"KV number(float)",
			`{"key":12.3}`,
			[]Token{
				{LCUB, "{", 0, 1},
				{STRING, "key", 1, 6},
				{COLON, ":", 6, 7},
				{NUMBER, `12.3`, 7, 11},
				{RCUB, "}", 11, 12},
			},
		},
		{
			"KV number(-int)",
			`{"key":-12}`,
			[]Token{
				{LCUB, "{", 0, 1},
				{STRING, "key", 1, 6},
				{COLON, ":", 6, 7},
				{NUMBER, `-12`, 7, 10},
				{RCUB, "}", 10, 11},
			},
		},
		{
			"KV array",
			`{"key":[12,"hello"]}`,
			[]Token{
				{LCUB, "{", 0, 1},
				{STRING, "key", 1, 6},
				{COLON, ":", 6, 7},
				{LSQB, "[", 7, 8},
				{NUMBER, "12", 8, 10},
				{COMMA, ",", 10, 11},
				{STRING, "hello", 11, 18},
				{RSQB, "]", 18, 19},
				{RCUB, "}", 19, 20},
			},
		},
		{
			"KV include space",
			`{ "key" : [12, "hello"] }`,
			[]Token{
				{LCUB, "{", 0, 1},
				{STRING, "key", 1 + 1, 6 + 1},
				{COLON, ":", 6 + 2, 7 + 2},
				{LSQB, "[", 7 + 3, 8 + 3},
				{NUMBER, `12`, 8 + 3, 10 + 3},
				{COMMA, ",", 10 + 3, 11 + 3},
				{STRING, `hello`, 11 + 4, 18 + 4},
				{RSQB, "]", 18 + 4, 19 + 4},
				{RCUB, "}", 19 + 5, 20 + 5},
			},
		},
		{
			"Array string",
			`["hello", "bob"]`,
			[]Token{
				{LSQB, "[", 0, 1},
				{STRING, `hello`, 1, 8},
				{COMMA, ",", 8, 9},
				{STRING, `bob`, 9 + 1, 14 + 1},
				{RSQB, "]", 14 + 1, 15 + 1},
			},
		},
	}
	tk := NewTokenizer()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual, err := tk.Tokenize(tt.in); !assert.Equal(t, tt.expect, actual) || err != nil {
				t.Errorf("expected=%v, but actual=%v", tt.expect, actual)
			}
		})
	}
}

func TestToken_String(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect []string
	}{
		{
			"KV string",
			`{`,
			[]string{
				"Token[0]\t{ Kind: LCUB,\tRaw: `{` }",
			},
		},
	}
	tk := NewTokenizer()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tk.Tokenize(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, len(tt.expect), len(actual))
			for i, str := range tt.expect {
				if !assert.Equal(t, str, actual[i].String()) {
					t.Errorf("expected=%v, but actual=%v", tt.expect, actual[i].String())
				}
			}
		})
	}
}

func TestTokenizer_Tokenize_2(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect []string
	}{
		{
			"KV string",
			`{`,
			[]string{},
		},
		{
			"KV string",
			`{"key":"value"}`,
			[]string{},
		},
	}
	tk := NewTokenizer()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens, err := tk.Tokenize(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			for _, tok := range tokens {
				fmt.Println(tok.String())
			}
		})
	}
}
