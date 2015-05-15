package spring

type Assessembly interface {
	Add(Spring) Spring
	Get(string) (Spring, error)
	Names() []string
	Clear() Assessembly
}

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
		err = ErrDoesNotExist
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
