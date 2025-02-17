package gositter

type rule struct {
	name       string
	expression Expression
}

func CreateRule(name string, ex Expression) *rule {
	return &rule{name, ex}
}

func (r *rule) Name() string {
	return r.name
}

func (r *rule) Parse(input string) (SyntaxTree, string, error) {
	t, remainder, err := r.expression.Parse(input)
	if err != nil {
		return t, remainder, err
	}
	if t != nil {
		t.SetTag(r.name)
	}
	return t, remainder, err
}
