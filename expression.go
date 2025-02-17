package gositter

type expression interface {
	parse(string) (SyntaxTree, string, error)
	bindRules(map[string]*rule)
}
