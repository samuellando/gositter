package gositter

type rule struct {
	name       string
	expression expression
}

func (r *rule) parse(input string) (SyntaxTree, string, error) {
	t, remainder, err := r.expression.parse(input)
	if err != nil {
		return t, remainder, err
	}
	if t != nil {
		t.SetTag(r.name)
	}
	return t, remainder, err
}
