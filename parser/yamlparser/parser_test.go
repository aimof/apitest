package yamlparser

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	want := Tests{
		Feature: "Sample Yaml",
		Scenarios: []Scenario{
			Scenario{
				Name:        "simple GET",
				Description: "Normal",
				Given: Given{
					Host: "http://localhost:8080",
				},
				Tests: []Test{
					Test{
						When: When{
							Host:   "",
							Path:   "/",
							Method: "GET",
							Header: nil,
							Body:   "",
						},
						Then: Then{
							Status:  200,
							Header:  nil,
							Format:  "",
							Require: nil,
							Retry:   true,
						},
					},
					Test{
						When: When{
							Host:   "",
							Path:   "/api",
							Method: "GET",
							Header: nil,
							Body:   "",
						},
						Then: Then{
							Status:  200,
							Header:  nil,
							Format:  "empty",
							Require: nil,
						},
					},
					Test{
						When: When{
							Host:   "http://localhost:8881",
							Path:   "/api/users/foo",
							Method: "POST",
							Header: nil,
							Body:   `{"Token": "Foo"}`,
						},
						Then: Then{
							Status:  200,
							Header:  []string{"key value", "key value"},
							Format:  "application/json",
							Require: []string{"foo", "bar"},
						},
					},
				},
			},
		},
	}
	ym := NewYamlParser()
	err := ym.parse("./test/bdd.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(ym.Tests, want) {
		t.Errorf("\ngot:  %v\nwant: %v", ym.Tests, want)
	}
}
