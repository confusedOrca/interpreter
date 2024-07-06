package lexer

import (
	"github.com/confusedOrca/interpreter/token"
)

type TokenLiteral interface {
	~byte | ~string
}

const nullChar = 0
const emptyString = ""

var mapToTokenType = map[string]token.TokenType{
	"+": token.PLUS, "-": token.MINUS, "/": token.SLASH, "*": token.ASTERISK,
	"<": token.LT, ">": token.GT, ";": token.SEMICOLON, ",": token.COMMA,
	"(": token.LPAREN, ")": token.RPAREN, "{": token.LBRACE, "}": token.RBRACE,
	"=": token.ASSIGN, "==": token.EQ, "!": token.BANG, "!=": token.NOT_EQ,
}

func newToken[T TokenLiteral](tokenType token.TokenType, ch T) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
