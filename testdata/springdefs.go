package testdata

type TestCase struct {
	Name     string
	Defs     []SpringDef
	Expected []string
}

type SpringDef struct {
	Name string
	Deps []string
}

var Tests = []TestCase{
	TestCase{
		Name: "simple dependency order",
		Defs: []SpringDef{
			SpringDef{Name: "a", Deps: []string{"b", "c"}},
			SpringDef{Name: "b"},
			SpringDef{Name: "c", Deps: []string{"b"}},
		},
		Expected: []string{"b", "c", "a"},
	},
	TestCase{
		Name: "deep dependency order",
		Defs: []SpringDef{
			SpringDef{Name: "a", Deps: []string{"b"}},
			SpringDef{Name: "b", Deps: []string{"c"}},
			SpringDef{Name: "c", Deps: []string{"d"}},
			SpringDef{Name: "d", Deps: []string{"e"}},
			SpringDef{Name: "e"},
		},
		Expected: []string{"e", "d", "c", "b", "a"},
	},
	TestCase{
		Name: "multibranch",
		Defs: []SpringDef{
			SpringDef{Name: "a", Deps: []string{"b", "c", "i"}},
			SpringDef{Name: "b", Deps: []string{"c", "e"}},
			SpringDef{Name: "c", Deps: []string{"d"}},
			SpringDef{Name: "d", Deps: []string{"e", "j"}},
			SpringDef{Name: "e", Deps: []string{"f"}},
			SpringDef{Name: "f", Deps: []string{"g"}},
			SpringDef{Name: "g"},
			SpringDef{Name: "h", Deps: []string{"e"}},
			SpringDef{Name: "i", Deps: []string{"j"}},
			SpringDef{Name: "j", Deps: []string{"k"}},
			SpringDef{Name: "k"},
		},
		Expected: []string{"g", "f", "e", "k", "j", "d", "c", "b", "i", "a"},
	},
	TestCase{
		Name: "multibranch",
		Defs: []SpringDef{
			SpringDef{Name: "a", Deps: []string{"b", "c", "i"}},
			SpringDef{Name: "b", Deps: []string{"c", "e"}},
			SpringDef{Name: "c", Deps: []string{"d", "j", "e"}},
			SpringDef{Name: "d", Deps: []string{"e", "j"}},
			SpringDef{Name: "e", Deps: []string{"f", "g"}},
			SpringDef{Name: "f", Deps: []string{"g", "i"}},
			SpringDef{Name: "g"},
			SpringDef{Name: "h", Deps: []string{"e"}},
			SpringDef{Name: "i", Deps: []string{"j", "k"}},
			SpringDef{Name: "j", Deps: []string{"k"}},
			SpringDef{Name: "k"},
		},
		Expected: []string{"g", "k", "j", "i", "f", "e", "d", "c", "b", "a"},
	},
}
