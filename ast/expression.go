package ast

import (
	"bytes"

	"github.com/confusedOrca/interpreter/token"
)

// ------------------------
// Identifier Class
// ------------------------

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// ------------------------
// IntegerLiteral Class
// ------------------------

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// ------------------------
// PrefixExpression Class
// ------------------------

type PrefixExpression struct {
	Token           token.Token
	Operator        string
	RightExpression Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.RightExpression.String())
	out.WriteString(")")
	return out.String()
}
