package main

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
	s.addSimpleToken(EOF)
	return s.tokens, nil
}

func (s *Scanner) scanToken() {
	b := s.advance()

	switch b {
	case '(':
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
