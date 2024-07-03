package parser

import (
	"github.com/confusedOrca/interpreter/ast"
	"github.com/confusedOrca/interpreter/token"
)

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.curToken.Type {

	case token.LET:
		return parser.parseLetStatement()

	case token.RETURN:
		return parser.parseReturnStatement()

	default:
		return parser.parseExpressionStatement()
	}
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: parser.curToken}

	if !parser.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Ident = &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}
	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	for !parser.curTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return stmt
}

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{
		Token: parser.curToken,
	}

	parser.nextToken()

	for !parser.curTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return stmt
}

func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{
		Token:      parser.curToken,
		Expression: parser.parseExpression(LOWEST),
	}

	if parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return stmt
}
