package parser

import (
	"errors"
	"fmt"

	"github.com/ByteHunter/glox/expression"
	"github.com/ByteHunter/glox/token"
)

type Parser struct {
	tokens  []token.Token
	current int
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) Parse() expression.Expression {
	if len(p.tokens) == 0 {
		return nil
	}
	return p.Expression()
}

func (p *Parser) Expression() expression.Expression {
	return p.Equality()
}

func (p *Parser) Equality() expression.Expression {
	var expr expression.Expression = p.Comparison()

	for p.match(token.BANQ_EQUAL, token.EQUAL_EQUAL) {
		var operator token.Token = p.previous()
		var right expression.Expression = p.Comparison()
		expr = expression.NewBinary(expr, operator, right)
	}

	return expr
}

func (p *Parser) Comparison() expression.Expression {
	var expr expression.Expression = p.Term()

	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		var operator token.Token = p.previous()
		var right expression.Expression = p.Term()
		expr = expression.NewBinary(expr, operator, right)
	}

	return expr
}

func (p *Parser) Term() expression.Expression {
	var expr expression.Expression = p.Factor()

	for p.match(token.MINUS, token.PLUS) {
		var operator token.Token = p.previous()
		var right expression.Expression = p.Factor()
		expr = expression.NewBinary(expr, operator, right)
	}

	return expr
}

func (p *Parser) Factor() expression.Expression {
	var expr expression.Expression = p.Unary()

	for p.match(token.MINUS, token.PLUS) {
		var operator token.Token = p.previous()
		var right expression.Expression = p.Unary()
		expr = expression.NewBinary(expr, operator, right)
	}

	return expr
}

func (p *Parser) Unary() expression.Expression {
	if p.match(token.BANG, token.MINUS) {
		var operator token.Token = p.previous()
		var right expression.Expression = p.Unary()
		return expression.NewUnary(operator, right)
	}

	return p.Primary()
}

func (p *Parser) Primary() expression.Expression {
	if p.match(token.FALSE) {
		return expression.NewLiteral(false)
	}
	if p.match(token.TRUE) {
		return expression.NewLiteral(true)
	}
	if p.match(token.NIL) {
		return expression.NewLiteral(nil)
	}

	if p.match(token.NUMBER, token.STRING) {
		return expression.NewLiteral(p.previous().Literal)
	}

	if p.match(token.LEFT_PAREN) {
		var expr expression.Expression = p.Expression()
		_, err := p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		if err != nil {
			fmt.Println(err.Error())
		}
		return expression.NewGrouping(expr)
	}

	return nil
}

func (p *Parser) match(tokenTypes ...token.TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(expected token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().Type == expected
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.current++
	}

	return p.previous()
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() token.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) consume(expected token.TokenType, message string) (token.Token, error) {
	if p.check(expected) {
		return p.advance(), nil
	}

	return token.Token{}, errors.New(message)
}
