package gositter

type SyntaxTree interface {
    Tree() string
    Value() string
    SetTag(string)
    Tag() string
    Find(string, ...bool) []SyntaxTree
}

type syntaxTree struct {
	tag   string
	value []SyntaxTree
}

func (t *syntaxTree) Find(tag string, recurse ...bool) []SyntaxTree {
	r := false
	if len(recurse) > 0 {
		r = recurse[0]
	}
	matches := make([]SyntaxTree, 0, 5)
	for _, st := range t.value {
		if st != nil {
			if st.Tag() == tag {
				matches = append(matches, st)
				if r {
					matches = append(matches, st.Find(tag)...)
				}
			} else {
				matches = append(matches, st.Find(tag)...)
			}
		}
	}
	return matches
}

func (t *syntaxTree) Tree() string {
	s := "(" + t.tag + " "
	for _, t := range t.value {
		if t != nil {
			s += t.Tree()
		}
	}
	return s + ")"
}

func (t *syntaxTree) Value() string {
	s := ""
	for _, t := range t.value {
		if t != nil {
			s += t.Value()
		}
	}
	return s
}

func (t *syntaxTree) Tag() string {
	return t.tag
}

func (t *syntaxTree) SetTag(tag string) {
	t.tag = tag
}
