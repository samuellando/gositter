package gositter

import (
    "fmt"
)

type repeat1 struct {
	expression Expression
}

func Repeat1(ex Expression) Expression {
	return &repeat1{ex}
}

func (e *repeat1) bindRules(rules map[string]*rule) {
    e.expression.bindRules(rules)
}

func (e *repeat1) Parse(input string) (SyntaxTree, string, error) {
	ze := Repeat(e.expression)
	t, remainder, err := ze.Parse(input)
	if err != nil {
		return t, remainder, err
	}
	if t == nil {
		return t, remainder, fmt.Errorf("One or more expected")
	}
	return t, remainder, err
}

