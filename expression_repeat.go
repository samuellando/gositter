package gositter

type repeat struct {
	expression expression
}

// Repetition expression, matches zero or more occurences of the sub expression
func Repeat(ex expression) expression {
	return &repeat{ex}
}

func (e *repeat) bindRules(rules map[string]*rule) {
	e.expression.bindRules(rules)
}

func (e *repeat) parse(input string) (SyntaxTree, string, error) {
	sts := make([]SyntaxTree, 0, 5)
	remainder := input
	var sub SyntaxTree
	var err error
	for {
		sub, remainder, err = e.expression.parse(remainder)
		if err != nil {
			break
		} else {
			sts = append(sts, sub)
		}
	}
	if len(sts) == 0 {
		return nil, remainder, err
	}
	t := &syntaxTree{"", sts}
	return t, remainder, nil
}
