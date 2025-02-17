package gositter

type seq struct {
	expressions []Expression
}

// A sequence expression. Matches one or more sub expressions in order.
func Seq(exs ...Expression) Expression {
	return &seq{exs}
}

func (e *seq) bindRules(rules map[string]*rule) {
	for _, exp := range e.expressions {
		exp.bindRules(rules)
	}
}

func (e *seq) parse(input string) (SyntaxTree, string, error) {
	sts := make([]SyntaxTree, 0, 5)
	remainder := input
	var sub SyntaxTree
	var err error
	for _, ex := range e.expressions {
		sub, remainder, err = ex.parse(remainder)
		if err != nil {
			return nil, input, err
		} else {
			sts = append(sts, sub)
		}
	}
	t := &syntaxTree{"", sts}
	return t, remainder, nil
}
