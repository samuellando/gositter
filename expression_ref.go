package gositter

import (
	"fmt"
)

type ref struct {
	name  string
	rules map[string]*rule
}

func Ref(v string) Expression {
	return &ref{name: v}
}

func (e *ref) bindRules(rules map[string]*rule) {
	e.rules = rules
}

func (e *ref) Parse(input string) (SyntaxTree, string, error) {
	rule, ok := e.rules[e.name]
	if !ok {
		return nil, input, fmt.Errorf("Reference not found")
	}
	return rule.Parse(input)
}
