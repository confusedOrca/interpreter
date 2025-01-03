package ast

import "bytes"

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// --------------- Program Class ---------------

type Program struct{ Statements []Statement }

const emptyString = ""

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return emptyString
	}
	return p.Statements[0].TokenLiteral()
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}
