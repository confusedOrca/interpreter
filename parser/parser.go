package parser

import (
	"fmt"

	"github.com/confusedOrca/interpreter/ast"
	"github.com/confusedOrca/interpreter/lexer"
	"github.com/confusedOrca/interpreter/token"
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
	p := &Parser{
		lexer:          lxr,
		errors:         []string{},
		prefixParseFns: make(map[token.TokenType]prefixParseFn),
		infixParseFns:  make(map[token.TokenType]infixParseFn),
	}

	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)

	for tokenType := range precedences {
		p.registerInfix(tokenType, p.parseInfixExpression)
	}

	p.nextToken()
	p.nextToken()

	return p
}

// ------------------------
// Parser Public Method
// ------------------------

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		program.Statements = append(program.Statements, stmt)
		p.nextToken()
	}
	return program
}

func (parser *Parser) Errors() []string { return parser.errors }

// -------------------------------
// Parser Private Utility Methods
// -------------------------------

func (p *Parser) peekPrecedence() int {
	if precedence, ok := precedences[p.peekToken.Type]; ok {
		return precedence
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if precedence, ok := precedences[p.curToken.Type]; ok {
		return precedence
	}
	return LOWEST
}

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
// Parser FN Registration Methods
// ---------------------------------------------------------

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
