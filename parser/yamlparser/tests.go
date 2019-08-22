package yamlparser

type Tests struct {
	Feature   string     `yaml:"Feature"`
	Scenarios []Scenario `yaml:"Scenarios"`
	Mocks     []Mock     `yaml:"Mock"`
}

type Scenario struct {
	Name        string `yaml:"Scenario"`
	Description string `yaml:"description"`
	Number      int
	Who         string `yaml:"Who"`
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
	Mocks   []MockBehavior `yaml:"Mocks"`
	Status  int            `yaml:"status"`
	Header  []string       `yaml:"header"`
	Format  string         `yaml:"format"`
	Require []string       `yaml:"require"`
	Retry   bool           `yaml:"retry"`
}

type MockBehavior struct {
	MockName string    `yaml:"Mock"`
	Given    MockGiven `yaml:"Given"`
	Then     MockThen  `yaml:"Then"`
}

type MockGiven struct {
	Method string `yaml:"method"`
	Path   string `yaml:"path"`
	Format string `yaml:"format"`
	Body   string `yaml:"body"`
}

type MockThen struct {
	Body string `yaml:"body"`
}

type Mock struct {
	Name string `yaml:"Name"`
	Port int    `yaml:"port"`
}
