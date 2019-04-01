package stringscomparer

import (
	"strings"

	"github.com/aimof/apitest"
)

// Comparer compare's two string.
type Comparer struct {
	wildcard string
}

// NewComparer : want is compare target.
func NewComparer(wildcard string) Comparer {
	return Comparer{
		wildcard: wildcard,
	}
}

// Compare Iface.
func (comp Comparer) Compare(g apitest.GotIface, w apitest.WantIface) bool {
	gotStatus, gotBody := g.Got()
	wantStatus, wantPattern := w.Want()
	return gotStatus == wantStatus && comp.compareString(string(gotBody), string(wantPattern))
}

// compareString and return match or not.
func (comp Comparer) compareString(got, want string) bool {
	targetStrings := strings.Split(want, comp.wildcard)
	for _, s := range targetStrings {
		if s == "" {
			continue
		}
		index := strings.Index(got, s)
		if index == -1 {
			return false
		}
		if index+len(s)-1 > len(got) {
			return false
		} else if index+len(s)-1 < len(got) {
			got = got[index+len(s)-1:]
		} else {
			got = ""
		}
	}
	return true
}
