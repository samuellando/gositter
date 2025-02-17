// A library generating parsers, inspired by tree-sitter
package gositter

import (
	"fmt"
)

type grammar struct {
	root  *rule
	rules map[string]*rule
}

// Create a grammar from a set of rules, identified by a name and expression.
// The name of the root rule must be provided so that the parsing knows where to
// Start.
func CreateGrammar(root string, rules map[string]Expression) *grammar {
	ruleSet := make(map[string]*rule)
	for k, v := range rules {
		v.bindRules(ruleSet)
		ruleSet[k] = &rule{name: k, expression: v}
	}
	return &grammar{ruleSet[root], ruleSet}
}

// Parse the input string and return the Syntax Tree.
// Returns an error if any error occurs during the parsing or if the input is
// not fully parsed.
func (g *grammar) Parse(input string) (SyntaxTree, error) {
	t, remainder, err := g.root.parse(input)
	if err != nil {
		return nil, err
	}
	if remainder != "" {
		return nil, fmt.Errorf("Remaining: %s %s", remainder, t.Tree())
	}
	return t, nil
}
