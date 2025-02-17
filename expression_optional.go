package gositter

type optional struct {
	expression expression
}

// Optional expression, matches zero or one occurences of the sub expression
func Optional(ex expression) expression {
	return &optional{ex}
}

func (e *optional) bindRules(rules map[string]*rule) {
	e.expression.bindRules(rules)
}

func (e *optional) parse(input string) (SyntaxTree, string, error) {
	t, remainder, err := e.expression.parse(input)
	if err != nil {
		return nil, remainder, nil
	}
	return t, remainder, err
}
