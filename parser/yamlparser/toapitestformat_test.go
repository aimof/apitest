package yamlparser

import (
	"reflect"
	"testing"
)

func TestParseHeader(t *testing.T) {
	headers := []struct {
		input string
		key   string
		value string
	}{
		{
			input: "Connection: keep-alive",
			key:   "Connection",
			value: "keep-alive",
		},
		{
			input: "foo",
			key:   "",
			value: "",
		},
		{
			input: "foo:foo:foo",
			key:   "",
			value: "",
		},
	}
	for i, h := range headers {
		k, v := parseHeader(h.input)
		if k != h.key || !reflect.DeepEqual(v, h.value) {
			t.Error(i)
		}
	}
}
