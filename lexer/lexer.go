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

const nullChar = 0

var charToTokenType = map[byte]token.TokenType{
	'+': token.PLUS, '-': token.MINUS, '/': token.SLASH, '*': token.ASTERISK,
	'<': token.LT, '>': token.GT, ';': token.SEMICOLON, ',': token.COMMA,
	'(': token.LPAREN, ')': token.RPAREN, '{': token.LBRACE, '}': token.RBRACE,
}

var stringToTokenType = map[string]token.TokenType{
	"=": token.ASSIGN, "==": token.EQ, "!": token.BANG, "!=": token.NOT_EQ,
}

func (lxr *Lexer) NextToken() token.Token {
	lxr.skipWhitespace()

	var tkn token.Token

	switch lxr.char {

	case '=', '!':
		caseString := string(lxr.char)
		if lxr.peekChar() == '=' {
			caseString += "="
			lxr.readChar()
		}
		tkn = newToken(stringToTokenType[caseString], caseString)

	case '+', '-', '/', '*', '<', '>', ';', '(', ')', ',', '{', '}':
		tkn = newToken(charToTokenType[lxr.char], lxr.char)

	case nullChar:
		tkn = newToken(token.EOF, "")

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
		lxr.char = nullChar
	} else {
		lxr.char = lxr.input[lxr.readPosition]
	}

	lxr.position = lxr.readPosition
	lxr.readPosition += 1
}

func (lxr *Lexer) peekChar() byte {
	if lxr.readPosition >= len(lxr.input) {
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

func (lxr *Lexer) skipWhitespace() {
	for isWhiteSpace(lxr.char) {
		lxr.readChar()
	}
}
