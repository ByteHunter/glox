package environment

import (
	"github.com/ByteHunter/glox/reporting"
	"github.com/ByteHunter/glox/token"
)

type Environment struct {
	Values map[string]any
}

func NewEnvironment() *Environment {
	return &Environment{
		Values: map[string]any{},
	}
}

func (e *Environment) Get(name token.Token) (any, error) {
	value, ok := e.Values[string(name.Lexeme)]

	if ok {
		return value, nil
	}

	return nil, reporting.NewRuntimeError(name, "Undefined variable '"+name.Lexeme+"'.")
}

func (e *Environment) Define(name string, value any) {
	e.Values[name] = value
}

func (e *Environment) Assign(name token.Token, value any) error {
	value, ok := e.Values[string(name.Lexeme)]
	if ok {
		e.Values[string(name.Lexeme)] = value
		return nil
	}

	return reporting.NewRuntimeError(name, "Undefined variable '" + name.Lexeme + "'.")
}
