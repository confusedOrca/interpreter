package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := input_case
	tests := test_case

	l := NewLexer(input)

	for i, tt := range tests {
		tkn := l.NextToken()

		if tkn.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tkn.Type)
		}
		if tkn.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tkn.Literal)
		}
	}
}
