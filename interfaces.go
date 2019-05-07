package apitest

import (
	"io"
	"net/http"
)

type Kicker interface {
	Kick(When) (*http.Response, error)
}

type ResponseComparer interface {
	MatchResponse(*http.Response, Then) (bool, error)
}

type Comparer interface {
	MatchBody(io.Reader, string) (bool, error)
}

type MockBuilder interface {
	Run([]Mock, Comparer) error
}
