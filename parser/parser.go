package parser

import (
	"fmt"

	"github.com/ByteHunter/glox/expression"
	"github.com/ByteHunter/glox/reporting"
	"github.com/ByteHunter/glox/token"
)

type ParseError struct {
	token   token.Token
	message string
}

func NewParseError(token token.Token, message string) *ParseError {
	return &ParseError{
		token:   token,
		message: message,
	}
}

func (p ParseError) Error() string {
	return fmt.Sprintf("ParseError %s", p.message)
}

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

	expr, err := p.Expression()
	if err != nil {
		return nil
	}

	return expr
}

func (p *Parser) Expression() (expression.Expression, error) {
	return p.Equality()
}

func (p *Parser) Equality() (expression.Expression, error) {
	expr, err := p.Comparison()
	if err != nil {
		return expr, err
	}

	for p.match(token.BANQ_EQUAL, token.EQUAL_EQUAL) {
		var operator token.Token = p.previous()
		right, err := p.Comparison()
		expr = expression.NewBinary(expr, operator, right)
		if err != nil {
			return expr, err
		}
	}

	return expr, nil
}

func (p *Parser) Comparison() (expression.Expression, error) {
	expr, err := p.Term()
	if err != nil {
		return expr, err
	}

	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		var operator token.Token = p.previous()
		right, err := p.Term()
		expr = expression.NewBinary(expr, operator, right)
		if err != nil {
			return expr, err
		}
	}

	return expr, nil
}

func (p *Parser) Term() (expression.Expression, error) {
	expr, err := p.Factor()
	if err != nil {
		return expr, err
	}

	for p.match(token.MINUS, token.PLUS) {
		var operator token.Token = p.previous()
		right, err := p.Factor()
		expr = expression.NewBinary(expr, operator, right)
		if err != nil {
			return expr, err
		}
	}

	return expr, nil
}

func (p *Parser) Factor() (expression.Expression, error) {
	expr, err := p.Unary()
	if err != nil {
		return expr, err
	}

	for p.match(token.SLASH, token.STAR) {
		var operator token.Token = p.previous()
		right, err := p.Unary()
		expr = expression.NewBinary(expr, operator, right)
		if err != nil {
			return expr, err
		}
	}

	return expr, nil
}

func (p *Parser) Unary() (expression.Expression, error) {
	if p.match(token.BANG, token.MINUS) {
		var operator token.Token = p.previous()
		right, err := p.Unary()
		return expression.NewUnary(operator, right), err
	}

	return p.Primary()
}

func (p *Parser) Primary() (expression.Expression, error) {
	if p.match(token.FALSE) {
		return expression.NewLiteral(false), nil
	}
	if p.match(token.TRUE) {
		return expression.NewLiteral(true), nil
	}
	if p.match(token.NIL) {
		return expression.NewLiteral(nil), nil
	}

	if p.match(token.NUMBER, token.STRING) {
		return expression.NewLiteral(p.previous().Literal), nil
	}

	if p.match(token.LEFT_PAREN) {
		expr, _ := p.Expression()
		_, err := p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		if err != nil {
			return expression.NewGrouping(expr), err
		}
		return expression.NewGrouping(expr), nil
	}

	reporting.LoxTokenError(p.peek(), "Expected expression")
	return nil, NewParseError(p.peek(), "Expected expression.")
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

	reporting.LoxTokenError(p.previous(), message)
	return p.previous(), NewParseError(p.previous(), message)
}

func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type == token.SEMICOLON {
			return
		}

		switch p.peek().Type {
		case
			token.CLASS,
			token.FUN,
			token.VAR,
			token.FOR,
			token.IF,
			token.WHILE,
			token.PRINT,
			token.RETURN:
			return
		}

		p.advance()
	}
}
