package parser

import (
	"github.com/confusedOrca/interpreter/ast"
	"github.com/confusedOrca/interpreter/lexer"
	"github.com/confusedOrca/interpreter/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	lexer          *lexer.Lexer
	curToken       token.Token
	peekToken      token.Token
	errors         []string
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(lxr *lexer.Lexer) *Parser {
	newParser := &Parser{
		lexer:          lxr,
		errors:         []string{},
		prefixParseFns: make(map[token.TokenType]prefixParseFn),
	}

	newParser.registerPrefix(token.IDENT, newParser.parseIdentifier)
	newParser.registerPrefix(token.INT, newParser.parseIntegerLiteral)
	newParser.registerPrefix(token.BANG, newParser.parsePrefixExpression)
	newParser.registerPrefix(token.MINUS, newParser.parsePrefixExpression)

	newParser.nextToken()
	newParser.nextToken()

	return newParser
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

func (parser *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	parser.prefixParseFns[tokenType] = fn
}

func (parser *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	parser.infixParseFns[tokenType] = fn
}

func (parser *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: parser.curToken,
		Value: parser.curToken.Literal,
	}
}

func (parser *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    parser.curToken,
		Operator: parser.curToken.Literal,
	}
	parser.nextToken()
	expression.RightExpression = parser.parseExpression(PREFIX)
	return expression
}
