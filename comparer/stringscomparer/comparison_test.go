package stringscomparer

import "testing"

type testcase struct {
	wildcard    string
	gotStatus   int
	gotBody     string
	wantStatus  int
	wantPattern string
	match       bool
}

func (tc testcase) Got() (int, []byte) {
	return tc.gotStatus, []byte(tc.gotBody)
}

func (tc testcase) Want() (int, []byte) {
	return tc.wantStatus, []byte(tc.wantPattern)
}

func TestCompare(t *testing.T) {
	testcases := []testcase{
		{"*", 200, "foo", 200, "f*o", true},
		{"*", 200, "foo", 500, "f*o", false},
		{"*", 200, "fo", 200, "abcdefoooooooo", false},
		{"*", 200, "Fizz fizz buzz Buzz", 200, "Fizz*Buzz", true},
	}

	for i, tc := range testcases {
		comp := NewComparer(tc.wildcard)
		if comp.Compare(tc, tc) != tc.match {
			t.Errorf("case: %d, wildcard: %s\ngot:  %d %s\nwant: %d %s", i, tc.wildcard, tc.gotStatus, tc.gotBody, tc.wantStatus, tc.wantPattern)
		}
	}
}

func TestCompareString(t *testing.T) {
	testCases := []struct {
		wildcard string
		pattern  string
		target   string
		match    bool
	}{
		// Match
		{wildcard: "*", pattern: "Do*yourself.", target: "Do it yourself.", match: true},
		{wildcard: "*", pattern: "Do*yourself.", target: "Don't repeat yourself.", match: true},
		// Not match
		{wildcard: "*", pattern: "Do*yourself", target: "foo", match: false},
		// Match: two or more wildcard
		{wildcard: "\n", pattern: "D\nR\nY", target: "Don't Repeat Yourself.", match: true},
		// Match: prefix is wildcard
		{wildcard: "*", pattern: "* yourself.", target: "Don't Repeat yourself.", match: true},
		{wildcard: "*", pattern: "test*", target: "testcase", match: true},
		{wildcard: "*", pattern: "Fizz*Buzz", target: "Fizz fizz buzz Buzz", match: true},
	}
	for i, tc := range testCases {
		comp := NewComparer(tc.wildcard)
		if comp.compareString(tc.target, tc.pattern) != tc.match {
			t.Errorf("case: %d,wildcard: %s\ntarget: %s\npattern: %s", i, comp.wildcard, tc.target, tc.pattern)
		}
	}
}
