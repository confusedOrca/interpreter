package ast

import (
	"bytes"

	"github.com/confusedOrca/interpreter/token"
)

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
