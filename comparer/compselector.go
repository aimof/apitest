package compselector

import (
	"github.com/aimof/apitest"
	"github.com/aimof/apitest/comparer/stringscomparer"
)

// Selector select configed comparers
type Selector map[string]apitest.Comparer

// Select Comparer from mode string
func (selector Selector) Select(s string) apitest.Comparer {
	if comp, ok := selector[s]; !ok {
		return stringscomparer.NewComparer("*")
	} else {
		return comp
	}
}
