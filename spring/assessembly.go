package spring

import "github.com/gonum/graph/concrete"

//Assessembly interface describes how springs should be assessed from storage
type Assessembly interface {
	Add(Spring) Spring
	Get(string) (Spring, error)
	Names() []string
	Clear() Assessembly
	Order(Spring) (Springs, error)
}

//NewAssessembly creates an Assessembly from the default impliementation
func NewAssessembly() Assessembly {
	return &defaultAssessembly{
		springs: map[string]Spring{},
	}
}

type defaultAssessembly struct {
	springs map[string]Spring
}

func (da *defaultAssessembly) Add(sprng Spring) Spring {
	_, ok := da.springs[sprng.Name()]
	if ok {
		panic(ErrDuplicateName)
	}
	da.springs[sprng.Name()] = sprng

	return sprng
}

func (da *defaultAssessembly) Get(name string) (Spring, error) {
	sprng, ok := da.springs[name]
	var err error
	if !ok {
		return nil, ErrDoesNotExist
	}
	return sprng, err
}

func (da *defaultAssessembly) Names() []string {
	names := make([]string, 0, len(da.springs))
	for name := range da.springs {
		names = append(names, name)
	}
	return names
}

func (da *defaultAssessembly) Clear() Assessembly {
	da.springs = map[string]Spring{}
	return da
}

func (da *defaultAssessembly) Order(sprng Spring) (Springs, error) {
	directed := concrete.NewDirectedGraph()
	if _, ok := da.springs[sprng.Name()]; !ok {
		return Springs{}, ErrDoesNotExist
	}

	for _, node := range da.springs {
		directed.AddNode(node)
		for _, name := range node.DependsOn() {
			dependency, ok := da.springs[name]
			if !ok {
				return Springs{}, ErrDoesNotExist
			}
			directed.AddDirectedEdge(concrete.Edge{F: node, T: dependency}, 1)
		}
	}

	var resolve func(Spring, []Spring, map[string]bool) ([]Spring, error)

	resolve = func(sprng Spring, resolved []Spring, unresolved map[string]bool) ([]Spring, error) {
		unresolved[sprng.Name()] = true
		for _, successor := range directed.Successors(sprng) {
			c := successor.(Spring)
			found := false
			for i := 0; i < len(resolved); i++ {
				if resolved[i].Name() == c.Name() {
					found = true
				}
			}

			if !found {
				if !unresolved[c.Name()] {
					var err error
					resolved, err = resolve(c, resolved, unresolved)
					if err != nil {
						return resolved, err
					}
				} else {
					return resolved, ErrCircularDependency
				}
			}
		}
		resolved = append(resolved, sprng)
		delete(unresolved, sprng.Name())
		return resolved, nil
	}

	nodes, err := resolve(sprng, []Spring{}, map[string]bool{})
	return nodes, err
}
