package parser

import "github.com/confusedOrca/interpreter/token"

func (parser *Parser) nextToken() {
	parser.curToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

func (parser *Parser) curTokenIs(expType token.TokenType) bool {
	return parser.curToken.Type == expType
}

func (parser *Parser) peekTokenIs(expType token.TokenType) bool {
	return parser.peekToken.Type == expType
}

func (parser *Parser) expectPeek(expType token.TokenType) bool {
	if parser.peekTokenIs(expType) {
		parser.nextToken()
		return true
	}

	parser.peekError(expType)
	return false
}
