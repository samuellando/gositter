package gositter

import (
	"errors"
)

type grammar struct {
	root  *rule
	rules map[string]*rule
}

func CreateGrammar(root string, rules map[string]Expression) *grammar {
	ruleSet := make(map[string]*rule)
	for k, v := range rules {
		v.bindRules(ruleSet)
		ruleSet[k] = CreateRule(k, v)
	}
	return &grammar{ruleSet[root], ruleSet}
}

func (g *grammar) Parse(input string) (SyntaxTree, error) {
	t, remainder, err := g.root.Parse(input)
	if err != nil {
		return nil, err
	}
	if remainder != "" {
		return nil, errors.New("Remaining: " + remainder + " " + t.Tree())
	}
	return t, nil
}
