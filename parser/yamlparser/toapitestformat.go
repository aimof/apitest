package yamlparser

import (
	"net/http"
	"strings"

	"github.com/aimof/apitest"
)

func toApitestScenario(in Scenario) (apitest.Scenario, error) {
	out := apitest.Scenario{
		Name: in.Name,
		Given: apitest.Given{
			WorkDir: in.Given.workdir,
		},
	}
	out.Tests = make([]apitest.Test, 0, len(in.Tests))
	for _, test := range in.Tests {
		w, err := toRequest(test.When)
		if err != nil {
			return apitest.Scenario{}, err
		}
		out.Tests = append(out.Tests, apitest.Test{
			When: w,
			Then: apitest.Then{
				Status:  test.Then.Status,
				Header:  test.Then.Header,
				Format:  test.Then.Format,
				Require: test.Then.Require,
				Retry:   test.Then.Retry,
			},
		})
	}
	return out, nil
}

func toRequest(w When) (apitest.When, error) {
	req, err := http.NewRequest(w.Method, w.Host+w.Path, strings.NewReader(w.Body))
	if err != nil {
		return apitest.When{}, err
	}
	return apitest.When{Request: req}, nil
}

func parseHeader(s string) (key, value string) {
	s = strings.ReplaceAll(s, " ", "")
	words := strings.Split(s, ":")
	if len(words) != 2 {
		return "", ""
	}

	if words[0] == "" || words[1] == "" {
		return "", ""
	}
	key = words[0]
	value = words[1]
	return key, value
}
