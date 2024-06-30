package lexer

import (
	"github.com/confusedOrca/interpreter/token"
)

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	l.readChar()
	return l
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
