package expression

import "github.com/ByteHunter/glox/token"

type Expression any

type Binary struct {
	Expression
	left     Expression
	operator token.Token
	right    Expression
}

func NewBinary(left Expression, operator token.Token, right Expression) *Binary {
	return &Binary{
		left: left,
		operator: operator,
		right: right,
	}
}
