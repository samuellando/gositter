package gositter

type SyntaxTree interface {
	// Get a string reprisentaion of the parse tree
	Tree() string
	// Get the value stored in this tree, equivalent to the original text
	Value() string
	// Set the tag on this tree, used for finding
	SetTag(string)
	// Get the tag of this tree
	Tag() string
	// Find a tag in this tree, with optinal recursion (default should be false)
	Find(string, ...bool) []SyntaxTree
    // Returns all children of the root node
	Nodes() []SyntaxTree
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
	matches := make([]SyntaxTree, 0)
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

func (t *syntaxTree) Nodes() []SyntaxTree {
    nodes := make([]SyntaxTree, 0)
    for _, node := range t.value {
        if node != nil {
            nodes = append(nodes, node)
        }
    }
    return nodes
}

func (t *syntaxTree) Tag() string {
	return t.tag
}

func (t *syntaxTree) SetTag(tag string) {
	t.tag = tag
}
