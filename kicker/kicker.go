package kicker

import (
	"net/http"

	"github.com/aimof/apitest"
)

// Kicker kicks API.
type Kicker struct {
	client *http.Client
}

// NewKicker with default http client.
func NewKicker() Kicker {
	return Kicker{
		client: new(http.Client),
	}
}

// Kick API
func (kicker Kicker) Kick(w apitest.When) (*http.Response, error) {
	return kicker.client.Do(w.Request)
}
