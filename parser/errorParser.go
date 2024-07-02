package parser

import (
	"fmt"

	"github.com/confusedOrca/interpreter/token"
)

func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) peekError(tknType token.TokenType) {
	msg := fmt.Sprintf(
		"expected next token to be %s, got %s instead",
		tknType, parser.peekToken.Type,
	)

	parser.errors = append(parser.errors, msg)
}
