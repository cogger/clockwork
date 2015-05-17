package spring

type node struct {
	sprng Spring
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
