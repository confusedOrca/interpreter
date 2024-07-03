package lexer

import (
	"testing"

	"github.com/confusedOrca/interpreter/lexer"
)

func TestNextToken(t *testing.T) {
	lxr := lexer.New(test_input)

	for i, expTkn := range expectedTkns {
		tkn := lxr.NextToken()

		if tkn.Type != expTkn.Type {
			t.Fatalf("[%d] wrong tkntype. expected=%q, got=%q",
				i, expTkn.Type, tkn.Type)
		}

		if tkn.Literal != expTkn.Literal {
			t.Fatalf("[%d] wrong literal. expected=%q, got=%q",
				i, expTkn.Literal, tkn.Literal)
		}
	}
}
