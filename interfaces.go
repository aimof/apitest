package apitest

import "net/http"

type Kicker interface {
	Kick(When) (*http.Response, error)
}

type Comparer interface {
	Match(*http.Response, Then) (bool, error)
}
