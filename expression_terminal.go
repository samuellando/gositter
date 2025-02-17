package gositter

import (
	"fmt"
	"strings"
)

type terminal struct {
	pattern string
}

// A terminal (or literal) expression.
func Terminal(v string) expression {
	return &terminal{v}
}

func (e *terminal) bindRules(rules map[string]*rule) {
}

func (e *terminal) parse(input string) (SyntaxTree, string, error) {
	if strings.HasPrefix(input, e.pattern) {
		var remainder string
		if len(input) == len(e.pattern) {
			remainder = ""
		} else {
			remainder = input[len(e.pattern):]
		}
		return &token{"", e.pattern}, remainder, nil
	} else {
		return nil, input, fmt.Errorf("Terminal did not match: '%s' '%s'", e.pattern, input)
	}
}
