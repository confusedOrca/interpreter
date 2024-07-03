package ast

import "github.com/confusedOrca/interpreter/token"

type ExpressionStatement struct {
	Token      token.Token // 1st tkn of expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
