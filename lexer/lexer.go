package lexer

import (
	"github.com/confusedOrca/interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {

	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Type, tok.Literal = token.EQ, string(ch)+string(l.ch)
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Type, tok.Literal = token.NOT_EQ, string(ch)+string(l.ch)
		} else {
			tok = newToken(token.BANG, l.ch)
		}

	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readBlock(isLetter)
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}

		if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readBlock(isDigit)
			return tok
		}

		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readBlock(isValid func(byte) bool) string {
	position := l.position
	for isValid(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
