package ast_test

import (
	"testing"

	"github.com/confusedOrca/interpreter/ast"
	"github.com/confusedOrca/interpreter/token"
)

func TestString(t *testing.T) {
	letStmt := &ast.LetStatement{
		Token: token.Token{
			Type:    token.LET,
			Literal: "let",
		},
		Ident: &ast.Identifier{
			Token: token.Token{Type: token.IDENT, Literal: "myVar"},
			Value: "myVar",
		},
		Value: &ast.Identifier{
			Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
			Value: "anotherVar",
		},
	}

	program := &ast.Program{Statements: []ast.Statement{letStmt}}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
