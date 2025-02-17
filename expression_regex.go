package gositter

import (
	"fmt"
	"regexp"
)

type regex struct {
	pattern string
}

func Regex(v string) Expression {
	return &regex{v}
}

func (e *regex) bindRules(rules map[string]*rule) {
}

func (e *regex) Parse(input string) (SyntaxTree, string, error) {
	c := regexp.MustCompile("^" + e.pattern)
	loc := c.FindStringIndex(input)
	if loc != nil {
		var remainder string
		matchLen := loc[1] - loc[0]
		if matchLen == len(input) {
			remainder = ""
		} else {
			remainder = input[matchLen:]
		}
		return &token{"", input[loc[0]:loc[1]]}, remainder, nil
	} else {
		return nil, input, fmt.Errorf("Regex did not match '%s' '%s'", e.pattern, input)
	}
}
