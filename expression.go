package gositter

type Expression interface {
	Parse(string) (SyntaxTree, string, error)
	bindRules(map[string]*rule)
}
