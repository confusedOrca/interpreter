package lexer

import (
	"github.com/confusedOrca/interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

func New(input string) *Lexer {
	newLxr := &Lexer{
		input: input,
	}

	newLxr.readChar()
	return newLxr
}

func (lxr *Lexer) NextToken() token.Token {
	lxr.skipWhitespace()
	var tkn token.Token

	switch lxr.char {

	case '=':
		if lxr.peekChar() == '=' {
			lxr.readChar()
			tkn = newToken(token.EQ, "==")
		} else {
			tkn = newToken(token.ASSIGN, "=")
		}

	case '!':
		if lxr.peekChar() == '=' {
			lxr.readChar()
			tkn = newToken(token.NOT_EQ, "!=")
		} else {
			tkn = newToken(token.BANG, "!")
		}

	case '+':
		tkn = newToken(token.PLUS, lxr.char)
	case '-':
		tkn = newToken(token.MINUS, lxr.char)
	case '/':
		tkn = newToken(token.SLASH, lxr.char)
	case '*':
		tkn = newToken(token.ASTERISK, lxr.char)
	case '<':
		tkn = newToken(token.LT, lxr.char)
	case '>':
		tkn = newToken(token.GT, lxr.char)
	case ';':
		tkn = newToken(token.SEMICOLON, lxr.char)
	case '(':
		tkn = newToken(token.LPAREN, lxr.char)
	case ')':
		tkn = newToken(token.RPAREN, lxr.char)
	case ',':
		tkn = newToken(token.COMMA, lxr.char)
	case '{':
		tkn = newToken(token.LBRACE, lxr.char)
	case '}':
		tkn = newToken(token.RBRACE, lxr.char)

	case NULLCHAR:
		tkn.Literal = ""
		tkn.Type = token.EOF

	default:
		if isLetter(lxr.char) {
			tkn.Literal = lxr.readBlock(isLetter)
			tkn.Type = token.LookupIdent(tkn.Literal)
			return tkn
		}

		if isDigit(lxr.char) {
			tkn.Type = token.INT
			tkn.Literal = lxr.readBlock(isDigit)
			return tkn
		}

		tkn = newToken(token.ILLEGAL, lxr.char)
	}

	lxr.readChar()
	return tkn
}

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
