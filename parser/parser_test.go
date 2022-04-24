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
			name: "only true",
			in: []tokenizer.Token{
				{
					tokenizer.KEYWORD,
					"true",
					0,
					5,
				},
			},
			expect: NewNode(
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
												0,
												5,
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
				{tokenizer.LCUB, "{", 0, 1},
				{tokenizer.STRING, "key", 1, 4},
				{tokenizer.COLON, ":", 4, 5},
				{tokenizer.STRING, "value", 5, 11},
				{tokenizer.RCUB, "}", 11, 12},
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
																		5, 11,
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
				{tokenizer.LCUB, "{", 0, 1},
				{tokenizer.STRING, "key1", 1, 5},
				{tokenizer.COLON, ":", 5, 6},
				{tokenizer.STRING, "value1", 6, 12},
				{tokenizer.COMMA, ",", 12, 13},
				{tokenizer.STRING, "key2", 13, 17},
				{tokenizer.COLON, ":", 17, 18},
				{tokenizer.STRING, "value2", 18, 24},
				{tokenizer.RCUB, "}", 24, 25},
			},
			expect: NewJsonElemValNode(NewChildren(NewNode(OBJECT, "", nil, NewChildren(
				NewNode(MEMBERS, "", nil, NewChildren(
					NewNode(MEMBER, "key1", nil, NewChildren(
						NewElemValString(tokenizer.Token{Kind: tokenizer.STRING, Raw: "value1", S: 6, E: 12}))),
					NewNode(MEMBER, "key2", nil, NewChildren(
						NewElemValString(tokenizer.Token{Kind: tokenizer.STRING, Raw: "value2", S: 18, E: 24}))),
				)),
			)))),
		},
		{
			name: "array: string, number, true, false, null",
			in: []tokenizer.Token{
				{tokenizer.LSQB, "[", 0, 1},
				{tokenizer.STRING, "string", 1, 7},
				{tokenizer.COMMA, ",", 7, 8},
				{tokenizer.NUMBER, "123", 8, 11},
				{tokenizer.COMMA, ",", 11, 12},
				{tokenizer.KEYWORD, "true", 12, 16},
				{tokenizer.COMMA, ",", 16, 17},
				{tokenizer.KEYWORD, "false", 17, 22},
				{tokenizer.COMMA, ",", 22, 23},
				{tokenizer.KEYWORD, "null", 23, 27},
				{tokenizer.RSQB, "]", 27, 28},
			},
			expect: NewJsonElemValNode(NewChildren(NewNode(ARRAY, "", nil, NewChildren(
				NewNode(ELEMENTS, "", nil, NewChildren(
					NewElemValString(tokenizer.Token{tokenizer.STRING, "string", 1, 7}),
					NewElemValNumber(tokenizer.Token{tokenizer.NUMBER, "123", 8, 11}),
					NewElemValTrue(tokenizer.Token{tokenizer.KEYWORD, "true", 12, 16}),
					NewElemValFalse(tokenizer.Token{tokenizer.KEYWORD, "false", 17, 22}),
					NewElemValNull(tokenizer.Token{tokenizer.KEYWORD, "null", 23, 27}),
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
