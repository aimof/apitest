package yamlparser

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aimof/apitest"
	yaml "gopkg.in/yaml.v2"
)

// YamlParser .
type YamlParser struct {
	Tests Tests
}

// NewYamlParser empty.
func NewYamlParser() *YamlParser {
	return &YamlParser{Tests{}}
}

func (yp *YamlParser) Parse(path string) ([]apitest.Scenario, error) {
	err := yp.parse(path)
	if err != nil {
		return nil, err
	}
	err = yp.afterParse(path)
	if err != nil {
		return nil, err
	}
	_, err = yp.toMocks()
	if err != nil {
		return nil, err
	}
	scenarios := make([]apitest.Scenario, 0, len(yp.Tests.Scenarios))
	for _, s := range yp.Tests.Scenarios {
		scenario, err := toApitestScenario(s)
		if err != nil {
			return nil, err
		}
		scenarios = append(scenarios, scenario)
	}
	return scenarios, nil
}

func (ym *YamlParser) parse(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	var t Tests
	err = yaml.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	ym.Tests = t
	return nil
}

func (ym *YamlParser) afterParse(path string) error {
	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	for i := range ym.Tests.Scenarios {
		ym.Tests.Scenarios[i].Given.workdir = filepath.Dir(abs)
		for j := range ym.Tests.Scenarios[i].Tests {
			ym.Tests.Scenarios[i].Tests[j].When.Host = ym.Tests.Scenarios[i].Given.Host
		}
	}
	return nil
}
