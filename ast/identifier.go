package ast

import "github.com/confusedOrca/interpreter/token"

type Identifier struct {
	Token token.Token
	Value string
}

func (identifier *Identifier) expressionNode() {}
func (identifier *Identifier) TokenLiteral() string {
	return identifier.Token.Literal
}
