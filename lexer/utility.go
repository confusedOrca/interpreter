package lexer

import (
	"github.com/confusedOrca/interpreter/token"
)

const NULLCHAR = 0

var charToTokenType = map[byte]token.TokenType{
	'+': token.PLUS,
	'-': token.MINUS,
	'/': token.SLASH,
	'*': token.ASTERISK,
	'<': token.LT,
	'>': token.GT,
	';': token.SEMICOLON,
	'(': token.LPAREN,
	')': token.RPAREN,
	',': token.COMMA,
	'{': token.LBRACE,
	'}': token.RBRACE,
}

type TokenLiteral interface {
	~byte | ~string
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
