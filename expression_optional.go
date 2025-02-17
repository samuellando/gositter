package gositter

type optional struct {
	expression Expression
}

// Optional expression, matches zero or one occurences of the sub expression
func Optional(ex Expression) Expression {
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
