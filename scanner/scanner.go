package scanner

import (
	"fmt"
	"strconv"

	"github.com/ByteHunter/glox/reporting"
	"github.com/ByteHunter/glox/token"
)

var keywords = map[string]token.TokenType{
	"and":    token.AND,
	"class":  token.CLASS,
	"else":   token.ELSE,
	"false":  token.FALSE,
	"fun":    token.FUN,
	"for":    token.FOR,
	"if":     token.IF,
	"nil":    token.NIL,
	"or":     token.OR,
	"print":  token.PRINT,
	"return": token.RETURN,
	"super":  token.SUPER,
	"this":   token.THIS,
	"true":   token.TRUE,
	"var":    token.VAR,
	"while":  token.WHILE,
}

type Scanner struct {
	source               string
	tokens               []token.Token
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

func (s *Scanner) GetTokens() []token.Token {
	return s.tokens
}

func (s *Scanner) ScanTokens() ([]token.Token, error) {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, *token.NewToken(token.EOF, "", nil, s.line))

	return s.tokens, nil
}

func (s *Scanner) scanToken() {
	b := s.advance()

	switch b {
	case '(':
		s.addSimpleToken(token.LEFT_PAREN)
	case ')':
		s.addSimpleToken(token.RIGHT_PAREN)
	case '{':
		s.addSimpleToken(token.LEFT_BRACE)
	case '}':
		s.addSimpleToken(token.RIGHT_BRACE)
	case ',':
		s.addSimpleToken(token.COMMA)
	case '.':
		s.addSimpleToken(token.DOT)
	case ';':
		s.addSimpleToken(token.SEMICOLON)
	case '-':
		s.addSimpleToken(token.MINUS)
	case '+':
		s.addSimpleToken(token.PLUS)
	case '*':
		s.addSimpleToken(token.STAR)
	case '!':
		if s.match('=') {
			s.addSimpleToken(token.BANG_EQUAL)
		} else {
			s.addSimpleToken(token.BANG)
		}
	case '=':
		if s.match('=') {
			s.addSimpleToken(token.EQUAL_EQUAL)
		} else {
			s.addSimpleToken(token.EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addSimpleToken(token.LESS_EQUAL)
		} else {
			s.addSimpleToken(token.LESS)
		}
	case '>':
		if s.match('=') {
			s.addSimpleToken(token.GREATER_EQUAL)
		} else {
			s.addSimpleToken(token.GREATER)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addSimpleToken(token.SLASH)
		}
	case '"':
		s.string()
	case ' ', '\t', '\r':
	case '\n':
		s.line++
	default:
		if s.isDigit(b) {
			s.number()
			break
		}
		if s.isAlpha(b) {
			s.identifier()
			break
		}
		reporting.LoxError(s.line, fmt.Sprintf("Unexpected character %c.", b))
	}
}

func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}

	value := s.source[s.start:s.current]
	tokenType, isKeyword := keywords[value]

	if !isKeyword {
		tokenType = token.IDENTIFIER
	}

	s.addSimpleToken(tokenType)
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		reporting.LoxError(s.line, "Unterminated string.")
		return
	}

	s.advance()

	value := s.source[s.start+1 : s.current-1]
	s.addToken(token.STRING, value)
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}
	}

	value := s.source[s.start:s.current]
	float_value := s.parseFloat(value)

	s.addToken(token.NUMBER, float_value)
}

func (s *Scanner) parseFloat(value string) float64 {
	float_value, err := strconv.ParseFloat(value, 64)
	if err != nil {
		reporting.LoxError(s.line, "Could not parse the literal as float")
	}

	return float_value
}

func (s *Scanner) isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func (s *Scanner) isAlpha(b byte) bool {
	return b == '_' ||
		(b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z')
}

func (s *Scanner) isAlphaNumeric(b byte) bool {
	return s.isAlpha(b) || s.isDigit(b)
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++

	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return '\x00'
	}

	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return '\x00'
	}

	return s.source[s.current+1]
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() byte {
	b := s.source[s.current]
	s.current++

	return b
}

func (s *Scanner) addSimpleToken(tokenType token.TokenType) {
	s.addToken(tokenType, nil)
}

func (s *Scanner) addToken(tokenType token.TokenType, value any) {
	lexeme := s.source[s.start:s.current]
	newToken := token.NewToken(tokenType, lexeme, value, s.line)
	s.tokens = append(s.tokens, *newToken)
}
