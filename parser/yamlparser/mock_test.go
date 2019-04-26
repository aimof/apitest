package yamlparser

import (
	"reflect"
	"testing"

	"github.com/aimof/apitest"
)

func TestToMock(t *testing.T) {
	var input = Tests{
		Mocks: []Mock{
			Mock{Name: "mock0", Port: 8080},
			Mock{Name: "mock1", Port: 8081},
		},
		Scenarios: []Scenario{
			Scenario{Tests: []Test{Test{
				Then: Then{Mocks: []MockBehavior{MockBehavior{
					MockName: "mock0",
					Given:    MockGiven{Method: "GET", Path: "/api", Format: "empty", Body: ""},
					Then:     MockThen{Body: "Foo"},
				}}},
			}}},
			Scenario{Tests: []Test{
				Test{
					Then: Then{Mocks: []MockBehavior{MockBehavior{
						MockName: "mock0",
						Given:    MockGiven{Method: "POST", Path: "/api/users", Format: "application/json", Body: "{}"},
						Then:     MockThen{Body: "Foo"},
					}}},
				},
				Test{
					Then: Then{Mocks: []MockBehavior{MockBehavior{
						MockName: "mock1",
						Given:    MockGiven{Method: "POST", Path: "/api", Format: "empty", Body: ""},
						Then:     MockThen{Body: "Foo"},
					}}},
				},
			}},
		},
	}

	var want = map[string]apitest.Mock{
		"mock0": apitest.Mock{
			Name: "mock0",
			Port: 8080,
			Behaviors: []apitest.MockBehavior{
				apitest.MockBehavior{
					Given: apitest.MockGiven{Method: "GET", Path: "/api", Format: "empty", Body: ""},
					Then:  apitest.MockThen{Body: "Foo"},
				},
				apitest.MockBehavior{
					Given: apitest.MockGiven{Method: "POST", Path: "/api/users", Format: "application/json", Body: "{}"},
					Then:  apitest.MockThen{Body: "Foo"},
				},
			},
		},
		"mock1": apitest.Mock{
			Name: "mock1",
			Port: 8081,
			Behaviors: []apitest.MockBehavior{
				apitest.MockBehavior{
					Given: apitest.MockGiven{Method: "POST", Path: "/api", Format: "empty", Body: ""},
					Then:  apitest.MockThen{Body: "Foo"},
				},
			},
		},
	}

	yp := &YamlParser{Tests: input}
	got, err := yp.toMocks()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot:  %v\nwant: %v", got, want)
	}
}
