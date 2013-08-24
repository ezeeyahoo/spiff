package flow

import (
	"github.com/vito/spiff/dynaml"
	"github.com/vito/spiff/yaml"
)

type Scope []map[string]yaml.Node

type Environment struct {
	Scope Scope
	Path  []string
}

func (e Environment) FindReference(path []string) dynaml.Node {
	root, found := resolveSymbol(path[0], e.Scope)
	if !found {
		return nil
	}

	return findInPath(path[1:], root)
}

func (e Environment) FindInStubs(path []string) dynaml.Node {
	return nil
}

func (e *Environment) PushScope(step map[string]yaml.Node) {
	e.Scope = append(e.Scope, step)
}

func (e *Environment) PushPath(step string) {
	e.Path = append(e.Path, step)
}

func findInPath(path []string, root yaml.Node) yaml.Node {
	here := root

	for _, step := range path {
		if here == nil {
			return nil
		}

		var found bool

		here, found = nextStep(step, here)
		if !found {
			return nil
		}
	}

	return here
}

func nextStep(step string, here yaml.Node) (yaml.Node, bool) {
	found := false
	switch here.(type) {
	case map[string]yaml.Node:
		found = true
		here = here.(map[string]yaml.Node)[step]
	default:
	}

	if !found {
		return nil, false
	}

	return here, true
}

func resolveSymbol(name string, context Scope) (yaml.Node, bool) {
	for i := len(context); i > 0; i-- {
		ctx := context[i-1]
		val := ctx[name]
		if val != nil {
			return val, true
		}
	}

	return nil, false
}
