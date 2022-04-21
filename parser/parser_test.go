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
			NewNode(
				JSON,
				"",
				nil,
				NewChildren(
					NewNode(
						ELEMENT,
						"",
						nil,
						NewChildren(
							NewNode(
								VALUE,
								"",
								nil,
								NewChildren(
									NewNode(
										TRUE,
										"",
										[]tokenizer.Token{
											{
												tokenizer.KEYWORD,
												"true",
											},
										},
										nil,
									),
								),
							),
						),
					),
				),
			),
		},
		{
			"kv string",
			[]tokenizer.Token{
				{tokenizer.LCUB, "{"},
				{tokenizer.STRING, "key"},
				{tokenizer.COLON, ":"},
				{tokenizer.STRING, "value"},
				{tokenizer.RCUB, "}"},
			},
			NewNode(
				JSON,
				"",
				nil,
				NewChildren(
					NewElementValueNode(
						NewChildren(
							NewNode(
								OBJECT,
								"",
								nil,
								NewChildren(
									NewNode(
										MEMBERS,
										"",
										nil,
										NewChildren(
											NewNode(
												MEMBER,
												"key",
												nil,
												NewChildren(
													NewElementValueNode(
														NewChildren(
															NewNode(
																STRING,
																"",
																[]tokenizer.Token{
																	{
																		tokenizer.STRING,
																		"value",
																	},
																},
																nil,
															)))),
											)),
									)),
							)))),
			),
		},
		{
			name: "kv string 2 pairs",
			in: []tokenizer.Token{
				{tokenizer.LCUB, "{"},
				{tokenizer.STRING, "key1"},
				{tokenizer.COLON, ":"},
				{tokenizer.STRING, "value1"},
				{tokenizer.COMMA, ","},
				{tokenizer.STRING, "key2"},
				{tokenizer.COLON, ":"},
				{tokenizer.STRING, "value2"},
				{tokenizer.RCUB, "}"},
			},
			expect: NewJsonElemValNode(NewChildren(NewNode(OBJECT, "", nil, NewChildren(
				NewNode(MEMBERS, "", nil, NewChildren(
					NewNode(MEMBER, "key1", nil, NewChildren(
						NewElemValString(tokenizer.Token{Kind: tokenizer.STRING, Raw: "value1"}))),
					NewNode(MEMBER, "key2", nil, NewChildren(
						NewElemValString(tokenizer.Token{Kind: tokenizer.STRING, Raw: "value2"}))),
				)),
			)))),
		},
		{
			name: "array: string, number, true, false, null",
			in: []tokenizer.Token{
				{tokenizer.LSQB, "["},
				{tokenizer.STRING, "string"},
				{tokenizer.COMMA, ","},
				{tokenizer.NUMBER, "123"},
				{tokenizer.COMMA, ","},
				{tokenizer.KEYWORD, "true"},
				{tokenizer.COMMA, ","},
				{tokenizer.KEYWORD, "false"},
				{tokenizer.COMMA, ","},
				{tokenizer.KEYWORD, "null"},
				{tokenizer.RSQB, "]"},
			},
			expect: NewJsonElemValNode(NewChildren(NewNode(ARRAY, "", nil, NewChildren(
				NewNode(ELEMENTS, "", nil, NewChildren(
					NewElemValString(tokenizer.Token{Kind: tokenizer.STRING, Raw: "string"}),
					NewElemValNumber(tokenizer.Token{Kind: tokenizer.NUMBER, Raw: "123"}),
					NewElemValTrue(tokenizer.Token{Kind: tokenizer.KEYWORD, Raw: "true"}),
					NewElemValFalse(tokenizer.Token{Kind: tokenizer.KEYWORD, Raw: "false"}),
					NewElemValNull(tokenizer.Token{Kind: tokenizer.KEYWORD, Raw: "null"}),
				)),
			)))),
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
