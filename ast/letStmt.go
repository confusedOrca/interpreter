package ast

import "github.com/confusedOrca/interpreter/token"

type LetStatement struct {
	Token token.Token
	Ident *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
