package mockbuilder

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/aimof/apitest"
)

type stringComparer struct{}

func (sc stringComparer) MatchBody(r io.Reader, s string) (bool, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return false, err
	}
	if string(b) != s {
		return false, errors.New("")
	}
	return true, nil
}

func TestBuild(t *testing.T) {
	mock := apitest.Mock{
		Name: "mock0",
		Port: 8080,
		Behaviors: []apitest.MockBehavior{
			apitest.MockBehavior{
				Given: apitest.MockGiven{
					Method: "GET",
					Path:   "/",
					Format: "empty",
					Body:   "",
				},
				Then: apitest.MockThen{
					Body: `{"Fizz": "Buzz"}`,
				},
			},
			apitest.MockBehavior{
				Given: apitest.MockGiven{
					Method: "POST",
					Path:   "/",
					Format: "application/json",
					Body:   "{}",
				},
				Then: apitest.MockThen{
					Body: `{"Foo": "Bar"}`,
				},
			},
		},
	}

	sc := stringComparer{}

	mb := NewMockBuilder()
	go mb.Run(mock, sc)

	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != `{"Fizz": "Buzz"}` {
		t.Error(string(b))
	}

	req, err := http.NewRequest("POST", "http://localhost:8080", strings.NewReader("{}"))
	if err != nil {
		t.Fatal(err)
	}

	client := http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != `{"Fizz": "Buzz"}` {
		t.Error(string(b))
	}
}
