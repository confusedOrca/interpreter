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
	newLxr := &Lexer{input: input}
	newLxr.readChar()
	return newLxr
}

// ------------------------
// Lexer Public Method
// ------------------------

func (lxr *Lexer) NextToken() token.Token {
	lxr.skipWhitespace()
	var tkn token.Token

	switch lxr.char {
	case '=', '!':
		if lxr.peekChar() == '=' {
			if lxr.char == '=' {
				tkn = newToken(token.EQ, "==")
			} else {
				tkn = newToken(token.NOT_EQ, "!=")
			}
			lxr.readChar()
		} else if lxr.char == '=' {
			tkn = newToken(token.ASSIGN, "=")
		} else {
			tkn = newToken(token.BANG, "!")
		}

	case '+', '-', '/', '*', '<', '>', ';', '(', ')', ',', '{', '}':
		tkn = newToken(charToTokenType[lxr.char], lxr.char)

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
			tkn.Literal = lxr.readBlock(isDigit)
			tkn.Type = token.INT
			return tkn
		}

		tkn = newToken(token.ILLEGAL, lxr.char)
	}

	lxr.readChar()
	return tkn
}

// ------------------------
// Lexer Private Methods
// ------------------------

func (lxr *Lexer) readChar() {
	if lxr.readPosition >= len(lxr.input) {
		lxr.char = NULLCHAR
	} else {
		lxr.char = lxr.input[lxr.readPosition]
	}

	lxr.position = lxr.readPosition
	lxr.readPosition += 1
}

func (lxr *Lexer) peekChar() byte {
	if lxr.readPosition >= len(lxr.input) {
		return NULLCHAR
	}

	return lxr.input[lxr.readPosition]
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
