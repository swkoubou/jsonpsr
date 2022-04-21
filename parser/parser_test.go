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
