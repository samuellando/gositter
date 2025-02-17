package gositter

type optional struct {
	expression Expression
}

func Optional(ex Expression) Expression {
	return &optional{ex}
}

func (e *optional) bindRules(rules map[string]*rule) {
    e.expression.bindRules(rules)
}

func (e *optional) Parse(input string) (SyntaxTree, string, error) {
	t, remainder, err := e.expression.Parse(input)
	if err != nil {
		return nil, remainder, nil
	}
	return t, remainder, err
}
