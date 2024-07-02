package parser

import (
	"testing"

	"github.com/confusedOrca/interpreter/ast"
	"github.com/confusedOrca/interpreter/lexer"
	"github.com/confusedOrca/interpreter/parser"
)

func TestLetStatements(t *testing.T) {
	lxr := lexer.New(test_input)
	parser := parser.New(lxr)
	program := parser.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	for i, expIdent := range expectedIdents {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, expIdent.Name) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("stmt.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("stmt not *ast.LetStatement. got=%T", stmt)
		return false
	}

	if letStmt.Ident.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Ident.Value)
		return false
	}

	if letStmt.Ident.TokenLiteral() != name {
		t.Errorf("stmt.Name not '%s'. got=%s", name, letStmt.Ident)
		return false
	}

	return true
}
