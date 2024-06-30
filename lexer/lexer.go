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
	var tkn token.Token

	l.skipWhitespace()

	switch l.ch {

	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tkn = newToken(token.EQ, string(ch)+string(l.ch))
		} else {
			tkn = newToken(token.ASSIGN, l.ch)
		}

	case '+':
		tkn = newToken(token.PLUS, l.ch)
	case '-':
		tkn = newToken(token.MINUS, l.ch)

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tkn = newToken(token.NOT_EQ, string(ch)+string(l.ch))
		} else {
			tkn = newToken(token.BANG, l.ch)
		}

	case '/':
		tkn = newToken(token.SLASH, l.ch)
	case '*':
		tkn = newToken(token.ASTERISK, l.ch)
	case '<':
		tkn = newToken(token.LT, l.ch)
	case '>':
		tkn = newToken(token.GT, l.ch)
	case ';':
		tkn = newToken(token.SEMICOLON, l.ch)
	case '(':
		tkn = newToken(token.LPAREN, l.ch)
	case ')':
		tkn = newToken(token.RPAREN, l.ch)
	case ',':
		tkn = newToken(token.COMMA, l.ch)
	case '{':
		tkn = newToken(token.LBRACE, l.ch)
	case '}':
		tkn = newToken(token.RBRACE, l.ch)
	case 0:
		tkn.Literal = ""
		tkn.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tkn.Literal = l.readBlock(isLetter)
			tkn.Type = token.LookupIdent(tkn.Literal)
			return tkn
		}

		if isDigit(l.ch) {
			tkn.Type = token.INT
			tkn.Literal = l.readBlock(isDigit)
			return tkn
		}

		tkn = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tkn
}
