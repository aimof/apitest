package yamlparser

import (
	"errors"

	"github.com/aimof/apitest"
)

func (yp YamlParser) toMocks() (map[string]apitest.Mock, error) {
	mocks := make(map[string]apitest.Mock, len(yp.Tests.Mocks))
	for _, m := range yp.Tests.Mocks {
		mocks[m.Name] = apitest.Mock{Name: m.Name, Port: m.Port, Behaviors: make([]apitest.MockBehavior, 0, 10)}
	}
	for _, s := range yp.Tests.Scenarios {
		for _, t := range s.Tests {
			for _, m := range t.Then.Mocks {
				mock, ok := mocks[m.MockName]
				if !ok {
					return nil, errors.New("undefined mock named: " + m.MockName)
				}
				b := apitest.MockBehavior{
					Given: apitest.MockGiven{
						Method: m.Given.Method,
						Path:   m.Given.Path,
						Format: m.Given.Format,
						Body:   m.Given.Body,
					},
					Then: apitest.MockThen{
						Body: m.Then.Body,
					},
				}
				mock.Behaviors = append(mock.Behaviors, b)
				mocks[m.MockName] = mock
			}
		}
	}
	return mocks, nil
}
