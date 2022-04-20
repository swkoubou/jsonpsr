package tokenizer

import (
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
				{STRING, "This is String"},
			},
		},
		{
			"one left curly bracket",
			"{",
			[]Token{
				{LCUB, "{"},
			},
		},
		{
			"KV string",
			"{\"key\":\"value\"}",
			[]Token{
				{LCUB, "{"},
				{STRING, "key"},
				{COLON, ":"},
				{STRING, "value"},
				{RCUB, "}"},
			},
		},
		{
			"KV string include escaped double quotation",
			`{"key":"val\"ue"}`,
			[]Token{
				{LCUB, "{"},
				{STRING, "key"},
				{COLON, ":"},
				{STRING, `val\"ue`},
				{RCUB, "}"},
			},
		},
		{
			"KV number(int)",
			`{"key":12}`,
			[]Token{
				{LCUB, "{"},
				{STRING, "key"},
				{COLON, ":"},
				{NUMBER, `12`},
				{RCUB, "}"},
			},
		},
		{
			"KV number(string)",
			`{"key":"12"}`,
			[]Token{
				{LCUB, "{"},
				{STRING, "key"},
				{COLON, ":"},
				{STRING, `12`},
				{RCUB, "}"},
			},
		},
		{
			"KV number(float)",
			`{"key":12.3}`,
			[]Token{
				{LCUB, "{"},
				{STRING, "key"},
				{COLON, ":"},
				{NUMBER, `12.3`},
				{RCUB, "}"},
			},
		},
		{
			"KV number(-int)",
			`{"key":-12}`,
			[]Token{
				{LCUB, "{"},
				{STRING, "key"},
				{COLON, ":"},
				{NUMBER, `-12`},
				{RCUB, "}"},
			},
		},
		{
			"KV array",
			`{"key":[12,"hello"]}`,
			[]Token{
				{LCUB, "{"},
				{STRING, "key"},
				{COLON, ":"},
				{LSQB, "["},
				{NUMBER, "12"},
				{COMMA, ","},
				{STRING, "hello"},
				{RSQB, "]"},
				{RCUB, "}"},
			},
		},
		{
			"KV include space",
			`{ "key" : [12, "hello"] }`,
			[]Token{
				{LCUB, "{"},
				{STRING, "key"},
				{COLON, ":"},
				{LSQB, "["},
				{NUMBER, `12`},
				{COMMA, ","},
				{STRING, `hello`},
				{RSQB, "]"},
				{RCUB, "}"},
			},
		},
	}
	tk := NewTokenizer()
	for _, tt := range tests {
		actual := tk.Tokenize(tt.in)
		if !assert.Equal(t, tt.expect, actual) {
			t.Fatalf("Test `%v` failed.", tt.name)
		}
	}
}
