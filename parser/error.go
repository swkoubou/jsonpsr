package parser

import (
	"errors"
	"fmt"
	"github.com/swkoubou/jsonpsr/tokenizer"
)

type ParseError struct {
	tokenizer.Token
	Err error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("%v: %v", e.Token.String(), e.Err)
}

func (e *ParseError) Unwrap() error {
	return e.Err
}

// Errorの種類

var (
	ErrValueMissing      = errors.New("missing value of `key:value` pair")
	ErrColonMissing      = errors.New("missing colon of `key:value` pair")
	ErrUnexpectedKeyword = errors.New("unexpected keyword")
	ErrUnexpectedString  = errors.New("unexpected string")
	ErrUnexpectedNumber  = errors.New("unexpected number")
	ErrUnexpectedSymbol  = errors.New("unexpected symbol")
)
