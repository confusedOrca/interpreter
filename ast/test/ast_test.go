package ast_test

import (
	"testing"

	"github.com/confusedOrca/interpreter/ast"
	"github.com/confusedOrca/interpreter/token"
)

func TestString(t *testing.T) {
	tkn := token.Token{Type: token.LET, Literal: "let"}

	ident := &ast.Identifier{
		Token: token.Token{Type: token.IDENT, Literal: "myVar"},
		Value: "myVar"}

	value := &ast.Identifier{
		Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
		Value: "anotherVar",
	}

	ls := &ast.LetStatement{Token: tkn, Ident: ident, Value: value}
	p := &ast.Program{Statements: []ast.Statement{ls}}

	if p.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", p.String())
	}
}
