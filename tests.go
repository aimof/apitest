package apitest

import "net/http"

type Scenario struct {
	Name  string
	Given Given
	Tests []Test
}

type Test struct {
	When When
	Then Then
}

type Given struct {
	WorkDir string
}

type When struct {
	Request *http.Request
}

type Then struct {
	Status  int
	Header  []string
	Format  string
	Require []string
}
