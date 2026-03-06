package main

import "fmt"

type Scanner struct {
	source               string
	tokens               []Token
	start, current, line int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) scanTokens() ([]Token, error) {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
    s.tokens = append(s.tokens, *NewToken(EOF, "", nil, s.line))

	return s.tokens, nil
}

func (s *Scanner) scanToken() {
	b := s.advance()

	switch b {
	case '(':
        s.addSimpleToken(LEFT_PAREN)
	case ')':
        s.addSimpleToken(RIGHT_PAREN)
	case '{':
        s.addSimpleToken(LEFT_BRACE)
	case '}':
        s.addSimpleToken(RIGHT_BRACE)
	case ',':
        s.addSimpleToken(COMMA)
	case '.':
        s.addSimpleToken(DOT)
	case ';':
        s.addSimpleToken(SEMICOLON)
	case '-':
        s.addSimpleToken(MINUS)
	case '+':
        s.addSimpleToken(PLUS)
	case '*':
        s.addSimpleToken(STAR)
    case ' ', '\t', '\r':
    case '\n':
        s.line++
    default:
        loxError(s.line, fmt.Sprintf("Unexpected character %c.", b))
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() byte {
	b := s.source[s.current]
	s.current++

	return b
}

func (s *Scanner) addSimpleToken(tokenType TokenType) {
	s.addToken(tokenType, nil)
}

func (s *Scanner) addToken(tokenType TokenType, value any) {
	lexeme := s.source[s.start:s.current]
	newToken := NewToken(tokenType, lexeme, value, s.line)
	s.tokens = append(s.tokens, *newToken)
}
