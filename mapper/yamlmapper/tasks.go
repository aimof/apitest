package yamlmapper

// Tasks corresponds yaml file.
type Tasks struct {
	Config    Config     `yaml:"config"`
	Testcases []Testcase `yaml:"testcases"`
}

// Config of Operation
type Config struct {
	workdir  string `yaml:"workdir"`
	Timeout  int    `yaml:"timeout"`
	WildCard string `yaml:"wildcard"`
}

// Testcase has request and response
type Testcase struct {
	Name     string   `yaml:"name"`
	URL      string   `yaml:"url"`
	Method   string   `yaml:"method"`
	BodyPath string   `yaml:"bodypath"`
	Header   []string `yaml:"header"`
	Retry    int      `yaml:"retry"`
	Want     Want     `yaml:"want"`
}

// Want this response
type Want struct {
	Statuscode int    `yaml:"statuscode"`
	BodyPath   string `yaml:"bodypath"`
}
