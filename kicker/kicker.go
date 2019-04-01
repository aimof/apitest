package kicker

import (
	"io/ioutil"
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
func (kicker Kicker) Kick(req *http.Request, retry int) (apitest.GotIface, error) {
	resp, err := kicker.client.Do(req)
	if err != nil {
		return G{}, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return G{}, err
	}
	return G{resp.StatusCode, b}, nil
}

// G is apitest.GotIface
type G struct {
	statuscode int
	body       []byte
}

// Got is apitest.GotIface.Got
func (g G) Got() (int, []byte) {
	return g.statuscode, g.body
}
