package ast

import (
	"bytes"

	"github.com/confusedOrca/interpreter/token"
)

// ------------------------
// Identifier
// ------------------------

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// ------------------------
// IntegerLiteral
// ------------------------

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// ------------------------
// PrefixExpression
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
	out.WriteString("(" + pe.Operator + pe.RightExpression.String() + ")")
	return out.String()
}
