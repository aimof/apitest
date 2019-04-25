package yamlparser

type Tests struct {
	Feature   string     `yaml:"Feature"`
	Scenarios []Scenario `yaml:"Scenarios"`
}

type Scenario struct {
	Name        string `yaml:"Scenario"`
	Description string `yaml:"description"`
	Given       Given  `yaml:"Given"`
	Tests       []Test `yaml:"Tests"`
}

type Given struct {
	Host    string `yaml:"host"`
	workdir string
}

type Test struct {
	When When `yaml:"When"`
	Then Then `yaml:"Then"`
}

type When struct {
	Host   string   `yaml:"host"`
	Path   string   `yaml:"path"`
	Method string   `yaml:"method"`
	Header []string `yaml:"header"`
	Body   string   `yaml:"body"`
}

type Then struct {
	Status  int      `yaml:"status"`
	Header  []string `yaml:"header"`
	Format  string   `yaml:"format"`
	Require []string `yaml:"require"`
}
