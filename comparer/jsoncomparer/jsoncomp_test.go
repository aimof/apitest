package jsoncomparer

import (
	"strings"
	"testing"

	"github.com/aimof/jason"
)

func TestMatchBody(t *testing.T) {
	testCases := []struct {
		require string
		body    string
		want    bool
	}{
		{
			require: `match {"Token": "Foo"}`,
			body:    `{"Token": "Foo"}`,
			want:    true,
		},
		{
			require: `match {"Token": "Foo"}`,
			body:    `{"Token": "Bar"}`,
			want:    false,
		},
		{
			require: `match {"Token": "Foo"}`,
			body:    `{"Token": "Foo", "Other": "Bar"}`,
			want:    true,
		},
	}

	for _, tc := range testCases {
		got, err := MatchBody(strings.NewReader(tc.body), tc.require)
		if err != nil {
			t.Error(err)
			continue
		}
		if got != tc.want {
			t.Error()
		}
	}
}

func TestMatchValue(t *testing.T) {
	testcases := []struct {
		got   string
		want  string
		match bool
	}{
		{
			got:   "true",
			want:  "true",
			match: true,
		},
		{
			got:   `"Foo"`,
			want:  `"Bar"`,
			match: false,
		},
	}

	for i, tc := range testcases {
		gotV, err := jason.NewValue(strings.NewReader(tc.got))
		if err != nil {
			t.Error(err)
			continue
		}
		wantV, err := jason.NewValue(strings.NewReader(tc.want))
		if matchValue(gotV, wantV) != tc.match {
			t.Error(i)
		}
	}
}

func TestMatchNestedValue(t *testing.T) {
	testcases := []struct {
		got   string
		want  string
		match bool
	}{
		{
			got:   `{"Foo": "Bar"}`,
			want:  `{"Foo": "Bar"}`,
			match: true,
		},
		{
			got: `{
				"Foo": "Bar",
				"Fizz": "Buzz"
			}`,
			want:  `{"Foo": "Bar"}`,
			match: true,
		},
		{
			got: `{
				"Foo": "Foo",
				"Bar": [
					"Fizz",
					"Buzz"
				]
			}`,
			want: `{
				"Bar": "` + ArrayType + `"
			}`,
			match: true,
		},
		{
			got: `{
				"Foo": "Foo",
				"Bar": [
					"Fizz",
					"Buzz"
				]
			}`,
			want: `{
				"Bar": [
					"Fizz",
					"Bar"
				]
			}`,
			match: false,
		},
		{
			got:   `{"Fizz": null}`,
			want:  `{"Fizz": "` + NullType + `"}`,
			match: true,
		},
		{
			got: `{
				"Foo": "Foo",
				"Bar": {
					"Fizz": null,
					"Buzz": 3
				}
			}`,
			want: `{
				"Bar": {
					"Fizz": "` + NullType + `"
				}
			}`,
			match: true,
		},
		{
			got: `{
				"Foo": "Foo",
				"Bar": {
					"Fizz": null,
					"Buzz": 3
				}
			}`,
			want: `{
				"Bar": {
					"Buzz": 3
				}
			}`,
			match: true,
		},
	}

	for i, tc := range testcases {
		gotV, err := jason.NewValue(strings.NewReader(tc.got))
		if err != nil {
			t.Errorf("case %d: error: %v", i, err)
			continue
		}
		wantV, err := jason.NewValue(strings.NewReader(tc.want))
		if err != nil {
			t.Errorf("case %d: error: %v", i, err)
			continue
		}
		if matchValue(gotV, wantV) != tc.match {
			t.Errorf("case %d\ngot:  %v\nwant: %v", i, gotV, wantV)
		}
	}
}

func TestMatchType(t *testing.T) {
	testcases := []struct {
		got   []byte
		want  string
		match bool
	}{
		{
			got:   []byte("true"),
			want:  BoolType,
			match: true,
		},
		{
			got:   []byte("true"),
			want:  BoolType,
			match: true,
		},
		{
			got:   []byte("2"),
			want:  NumberType,
			match: true,
		},
		{
			got: []byte(`[
				"Fizz",
				"Buzz"
			]`),
			want:  ArrayType,
			match: true,
		},
		{
			got:   []byte(``),
			want:  ArrayType,
			match: false,
		},
	}

	for i, tc := range testcases {
		got, err := jason.NewValueFromBytes(tc.got)
		if err != nil {
			t.Error(i, err)
			continue
		}

		if tc.match != matchType(got, tc.want) {
			t.Error(i)
		}
	}
}
