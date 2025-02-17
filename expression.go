package gositter

type Expression interface {
	parse(string) (SyntaxTree, string, error)
	bindRules(map[string]*rule)
}
