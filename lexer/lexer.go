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

// ------------------------ Lexer Public Method ------------------------

func (lxr *Lexer) NextToken() token.Token {
	lxr.skipWhitespace()

	var tkn token.Token
	switch lxr.char {
	case '+', '-', '/', '*', '<', '>', ';', '(', ')', ',', '{', '}':
		tkn = newToken(mapToTokenType[string(lxr.char)], lxr.char)

	case '=', '!':
		caseString := string(lxr.char)
		if lxr.peekChar() == '=' {
			caseString += "="
			lxr.readChar()
		}
		tkn = newToken(mapToTokenType[caseString], caseString)

	case nullChar:
		tkn = newToken(token.EOF, emptyString)

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

// ------------------------ Lexer Private Methods ------------------------

func (lxr *Lexer) isOutOfBound() bool {
	return lxr.readPosition >= len(lxr.input)
}

func (lxr *Lexer) readChar() {
	if lxr.isOutOfBound() {
		lxr.char = nullChar
	} else {
		lxr.char = lxr.input[lxr.readPosition]
	}
	lxr.position = lxr.readPosition
	lxr.readPosition += 1
}

func (lxr *Lexer) peekChar() byte {
	if lxr.isOutOfBound() {
		return nullChar
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

func (l *Lexer) skipWhitespace() {
	for isWhiteSpace(l.char) {
		l.readChar()
	}
}
