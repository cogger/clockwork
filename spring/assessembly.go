package spring

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
		springs: map[string]*node{},
	}
}

type defaultAssessembly struct {
	springs map[string]*node
}

func (da *defaultAssessembly) Add(sprng Spring) Spring {
	_, ok := da.springs[sprng.Name()]
	if ok {
		panic(ErrDuplicateName)
	}
	da.springs[sprng.Name()] = &node{
		sprng: sprng,
	}

	return sprng
}

func (da *defaultAssessembly) Get(name string) (Spring, error) {
	node, ok := da.springs[name]
	var err error
	if !ok {
		return nil, ErrDoesNotExist
	}
	return node.sprng, err
}

func (da *defaultAssessembly) Names() []string {
	names := make([]string, 0, len(da.springs))
	for name := range da.springs {
		names = append(names, name)
	}
	return names
}

func (da *defaultAssessembly) Clear() Assessembly {
	da.springs = map[string]*node{}
	return da
}

func (da *defaultAssessembly) resetNodes() {
	for _, n := range da.springs {
		n.edges = []*node{}
	}
}

func (da *defaultAssessembly) Order(sprng Spring) (Springs, error) {
	var sprngs Springs
	var err error

	da.resetNodes()

	for _, n := range da.springs {
		for _, dep := range n.sprng.DependsOn() {
			c, ok := da.springs[dep]
			if !ok {
				return sprngs, ErrDoesNotExist
			}
			n.addEdge(c)
		}
	}

	n, ok := da.springs[sprng.Name()]
	if !ok {
		return sprngs, ErrDoesNotExist
	}

	nodes, err := n.resolve([]*node{}, map[string]bool{})
	if err != nil {
		return sprngs, err
	}

	for _, n := range nodes {
		sprngs = append(sprngs, n.sprng)
	}
	return sprngs, err
}
