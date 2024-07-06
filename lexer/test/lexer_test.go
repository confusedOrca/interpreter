package lexer_test

import (
	"testing"

	"github.com/confusedOrca/interpreter/lexer"
	"github.com/confusedOrca/interpreter/token"
)

func TestNextToken(t *testing.T) {
	lxr := lexer.New(test_input)

	for i, tt := range expectedTkns {
		tkn := lxr.NextToken()

		if tkn.Type != tt.Type {
			t.Fatalf("[%d] wrong tkntype. expected=%q, got=%q",
				i, tt.Type, tkn.Type)
		}

		if tkn.Literal != tt.Literal {
			t.Fatalf("[%d] wrong literal. expected=%q, got=%q",
				i, tt.Literal, tkn.Literal)
		}
	}
}

var test_input = `let five = 5;
				let ten = 10;
				let add = fn(x, y) {
				x + y;
				};
				let result = add(five, ten);
				!-/*5;
				5 < 10 > 5;
				if (5 < 10) {
				return true;
				} else {
				return false;
				}
				10 == 10;
				10 != 9;`

var expectedTkns = []struct {
	Type    token.TokenType
	Literal string
}{
	{token.LET, "let"},
	{token.IDENT, "five"},
	{token.ASSIGN, "="},
	{token.INT, "5"},
	{token.SEMICOLON, ";"},
	{token.LET, "let"},
	{token.IDENT, "ten"},
	{token.ASSIGN, "="},
	{token.INT, "10"},
	{token.SEMICOLON, ";"},
	{token.LET, "let"},
	{token.IDENT, "add"},
	{token.ASSIGN, "="},
	{token.FUNCTION, "fn"},
	{token.LPAREN, "("},
	{token.IDENT, "x"},
	{token.COMMA, ","},
	{token.IDENT, "y"},
	{token.RPAREN, ")"},
	{token.LBRACE, "{"},
	{token.IDENT, "x"},
	{token.PLUS, "+"},
	{token.IDENT, "y"},
	{token.SEMICOLON, ";"},
	{token.RBRACE, "}"},
	{token.SEMICOLON, ";"},
	{token.LET, "let"},
	{token.IDENT, "result"},
	{token.ASSIGN, "="},
	{token.IDENT, "add"},
	{token.LPAREN, "("},
	{token.IDENT, "five"},
	{token.COMMA, ","},
	{token.IDENT, "ten"},
	{token.RPAREN, ")"},
	{token.SEMICOLON, ";"},
	{token.BANG, "!"},
	{token.MINUS, "-"},
	{token.SLASH, "/"},
	{token.ASTERISK, "*"},
	{token.INT, "5"},
	{token.SEMICOLON, ";"},
	{token.INT, "5"},
	{token.LT, "<"},
	{token.INT, "10"},
	{token.GT, ">"},
	{token.INT, "5"},
	{token.SEMICOLON, ";"},
	{token.IF, "if"},
	{token.LPAREN, "("},
	{token.INT, "5"},
	{token.LT, "<"},
	{token.INT, "10"},
	{token.RPAREN, ")"},
	{token.LBRACE, "{"},
	{token.RETURN, "return"},
	{token.TRUE, "true"},
	{token.SEMICOLON, ";"},
	{token.RBRACE, "}"},
	{token.ELSE, "else"},
	{token.LBRACE, "{"},
	{token.RETURN, "return"},
	{token.FALSE, "false"},
	{token.SEMICOLON, ";"},
	{token.RBRACE, "}"},
	{token.INT, "10"},
	{token.EQ, "=="},
	{token.INT, "10"},
	{token.SEMICOLON, ";"},
	{token.INT, "10"},
	{token.NOT_EQ, "!="},
	{token.INT, "9"},
	{token.SEMICOLON, ";"},
	{token.EOF, ""},
}
