package gositter

import (
	"fmt"
)

type ref struct {
	name  string
	rules map[string]*rule
}

// A refernece expression, use to refer to another rule in the grammar
func Ref(v string) expression {
	return &ref{name: v}
}

func (e *ref) bindRules(rules map[string]*rule) {
	e.rules = rules
}

func (e *ref) parse(input string) (SyntaxTree, string, error) {
	rule, ok := e.rules[e.name]
	if !ok {
		return nil, input, fmt.Errorf("Reference not found")
	}
	return rule.parse(input)
}
