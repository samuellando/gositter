package gositter

type token struct {
	tag   string
	value string
}

func (t *token) Tree() string {
	return "(" + t.tag + " " + t.value + ")"
}

func (t *token) Value() string {
	return t.value
}

func (t *token) Tag() string {
	return t.tag
}

func (t *token) SetTag(tag string) {
	t.tag = tag
}

func (t *token) Find(tag string, recurse ...bool) []SyntaxTree {
	return []SyntaxTree{}
}
