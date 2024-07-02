package parser

import (
	"github.com/confusedOrca/interpreter/ast"
	"github.com/confusedOrca/interpreter/lexer"
	"github.com/confusedOrca/interpreter/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(lxr *lexer.Lexer) *Parser {
	newParser := &Parser{
		lexer: lxr,
	}

	newParser.nextToken()
	newParser.nextToken()

	return newParser
}

func (parser *Parser) nextToken() {
	parser.curToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for parser.curToken.Type != token.EOF {
		stmt := parser.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		parser.nextToken()
	}
	return program
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.curToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	default:
		return nil
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
	return false
}
