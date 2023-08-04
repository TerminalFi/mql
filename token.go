// Copyright (c) HashiCorp, Inc.

package mql

type token struct {
	Type  tokenType
	Value string
}

type tokenType int

const eof rune = -1

const (
	unknownToken tokenType = iota
	errToken
	eofToken
	whitespaceToken
	stringToken
	startLogicalExprToken
	endLogicalExprToken
	greaterThanToken
	greaterThanOrEqualToken
	lessThanToken
	lessThanOrEqualToken
	equalToken
	notEqualToken
	containsToken

	// keywords
	andToken
	orToken
)

var tokenTypeToString = map[tokenType]string{
	unknownToken:            "unknown",
	errToken:                "err",
	eofToken:                "eof",
	whitespaceToken:         "ws",
	stringToken:             "str",
	startLogicalExprToken:   "lparen",
	endLogicalExprToken:     "rparen",
	greaterThanToken:        "gt",
	greaterThanOrEqualToken: "gte",
	lessThanToken:           "lt",
	lessThanOrEqualToken:    "lte",
	equalToken:              "eq",
	notEqualToken:           "neq",
	containsToken:           "contains",
	andToken:                "and",
	orToken:                 "or",
}

// String returns a string of the tokenType and will return "Unknown" for
// invalid tokenTypes
func (t tokenType) String() string {
	s, ok := tokenTypeToString[t]
	switch ok {
	case true:
		return s
	default:
		return tokenTypeToString[unknownToken]
	}
}
