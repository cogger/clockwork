package mainspring

import (
	"github.com/cogger/clockwork/spring"
	"github.com/cogger/cogger"
	"golang.org/x/net/context"
)

func Wind(ctx context.Context, desired spring.Spring, springs spring.Assessembly) (cogger.Cog, error) {
	_, err := resolve(desired, springs)
	return desired, err
}

func resolve(desired spring.Spring, assessembly spring.Assessembly) ([]*node, error) {
	nodes, err := assessemblyToNodes(assessembly)
	if err != nil {
		return []*node{}, err
	}

	return nodes[desired.Name()].resolve([]*node{}, map[string]bool{})
}

func assessemblyToNodes(assessembly spring.Assessembly) (map[string]*node, error) {
	nodes := map[string]*node{}
	for _, name := range assessembly.Names() {
		sprng, _ := assessembly.Get(name)
		nodes[sprng.Name()] = &node{
			sprng: sprng,
			edges: []*node{},
		}
	}

	for _, node := range nodes {
		for _, dep := range node.sprng.DependsOn() {
			n, ok := nodes[dep]
			if !ok {
				return nodes, spring.ErrDoesNotExist
			}
			node.addEdge(n)
		}
	}
	return nodes, nil
}

type node struct {
	sprng spring.Spring
	edges []*node
}

func (n *node) addEdge(c *node) {
	n.edges = append(n.edges, c)
}

func (n *node) resolve(resolved []*node, unresolved map[string]bool) ([]*node, error) {
	unresolved[n.sprng.Name()] = true
	for _, c := range n.edges {
		found := false
		for i := 0; i < len(resolved); i++ {
			if resolved[i].sprng.Name() == c.sprng.Name() {
				found = true
			}
		}
		if !found {
			if !unresolved[c.sprng.Name()] {
				var err error
				resolved, err = c.resolve(resolved, unresolved)
				if err != nil {
					return resolved, err
				}
			} else {
				return resolved, ErrCircularDependency
			}
		}
	}
	resolved = append(resolved, n)
	delete(unresolved, n.sprng.Name())
	return resolved, nil
}
