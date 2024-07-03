package parser

import (
	"fmt"

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

// ------------------------
// Parser Public Method
// ------------------------

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{Statements: []ast.Statement{}}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (parser *Parser) Errors() []string {
	return parser.errors
}

// -------------------------------
// Parser Private Utility Methods
// -------------------------------

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) curTokenIs(expType token.TokenType) bool {
	return p.curToken.Type == expType
}

func (p *Parser) peekTokenIs(expType token.TokenType) bool {
	return p.peekToken.Type == expType
}

func (p *Parser) expectPeek(expType token.TokenType) bool {
	if p.peekTokenIs(expType) {
		p.nextToken()
		return true
	}

	p.peekError(expType)
	return false
}

func (p *Parser) peekError(tknType token.TokenType) {
	msg := fmt.Sprintf(
		"expected next token to be %s, got %s instead",
		tknType, p.peekToken.Type,
	)

	p.errors = append(p.errors, msg)
}

// ---------------------------------------------------------
// Parser Private Parser FN & Parser FN Registration Methods
// ---------------------------------------------------------

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	expression.RightExpression = p.parseExpression(PREFIX)
	return expression
}
