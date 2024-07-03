package parser_test

import (
	"testing"

	"github.com/confusedOrca/interpreter/ast"
	"github.com/confusedOrca/interpreter/lexer"
	"github.com/confusedOrca/interpreter/parser"
)

// --------------------------
// Let Statements Test
// --------------------------

var letStmt_input = `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`
var letStmt_expIdents = []struct {
	Name string
}{
	{"x"},
	{"y"},
	{"foobar"},
}

func TestLetStatements(t *testing.T) {
	lxr := lexer.New(letStmt_input)
	parser := parser.New(lxr)
	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	for i, expIdent := range letStmt_expIdents {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, expIdent.Name) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, identName string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("stmt.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("stmt not *ast.LetStatement. got=%T", stmt)
		return false
	}

	if letStmt.Ident.Value != identName {
		t.Errorf("letStmt.Ident.Value not '%s'. got=%s", identName, letStmt.Ident.Value)
		return false
	}

	if letStmt.Ident.TokenLiteral() != identName {
		t.Errorf("stmt.Name not '%s'. got=%s", identName, letStmt.Ident)
		return false
	}

	return true
}

// --------------------------
// Return Statement Test
// --------------------------

var retStmt_input = `
	return 5;
	return 10;
	return 993322;
	`

func TestReturnStatements(t *testing.T) {
	lxr := lexer.New(retStmt_input)
	parser := parser.New(lxr)
	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement, got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}