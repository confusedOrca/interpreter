package lexer

import (
	"github.com/confusedOrca/interpreter/token"
)

const NULLCHAR = 0

func (lxr *Lexer) readChar() {
	if lxr.readPosition >= len(lxr.input) {
		lxr.char = NULLCHAR
	} else {
		lxr.char = lxr.input[lxr.readPosition]
	}

	lxr.position = lxr.readPosition
	lxr.readPosition += 1
}

func (lxr *Lexer) readBlock(isValid func(byte) bool) string {
	startPosition := lxr.position

	for isValid(lxr.char) {
		lxr.readChar()
	}

	return lxr.input[startPosition:lxr.position]
}

func (lxr *Lexer) skipWhitespace() {
	for isWhiteSpace(lxr.char) {
		lxr.readChar()
	}
}

func (lxr *Lexer) peekChar() byte {
	if lxr.readPosition >= len(lxr.input) {
		return NULLCHAR
	} else {
		return lxr.input[lxr.readPosition]
	}
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
