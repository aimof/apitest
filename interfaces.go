package apitest

import "net/http"

// Kicker kicks API.
type Kicker interface {
	// Kick params: http request, retry count.
	Kick(*http.Request, int) (GotIface, error)
}

// CompSelector selects Comparer and return it.
type CompSelector interface {
	Select(string) Comparer
}

// Comparer params are compmode, got, want'.
type Comparer interface {
	Compare(GotIface, WantIface) bool
}

// GotIface to compare.
type GotIface interface {
	Got() (int, []byte)
}

// WantIface to compare.
type WantIface interface {
	Want() (int, []byte)
}
