package parser

import (
	"fmt"

	"github.com/confusedOrca/interpreter/ast"
	"github.com/confusedOrca/interpreter/token"
)

/*
the const assigned iota is 0 and all other const increments by 1
as we move down the list of constants.
*/
const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

func (parser *Parser) parseExpression(precedence int) ast.Expression {
	prefix := parser.prefixParseFns[parser.curToken.Type]
	if prefix == nil {
		parser.noPrefixParseFnError(parser.curToken.Type)
		return nil
	}

	leftExp := prefix()
	return leftExp
}

func (parser *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	parser.errors = append(parser.errors, msg)
}
