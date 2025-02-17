package gositter

import (
	"bytes"
	"fmt"
)

type choice struct {
	expressions []expression
}

// A choice expression, matches the first succesful sub expression
func Choice(exs ...expression) expression {
	return &choice{exs}
}

func (e *choice) bindRules(rules map[string]*rule) {
	for _, ex := range e.expressions {
		ex.bindRules(rules)
	}
}

func (e *choice) parse(input string) (SyntaxTree, string, error) {
	errs := new(bytes.Buffer)
	for _, ex := range e.expressions {
		t, remainder, err := ex.parse(input)
		if err != nil {
			fmt.Fprintf(errs, "%s\n", err)
			continue
		}
		return &syntaxTree{"", []SyntaxTree{t}}, remainder, err
	}
	return nil, input, fmt.Errorf("No option parsed succesfully:\n %s", errs.String())
}
