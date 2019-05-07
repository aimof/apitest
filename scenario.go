package apitest

import "net/http"

type Scenario struct {
	Name  string
	Who   string
	Given Given
	Tests []Test
}

type Given struct {
	WorkDir string
}

type Test struct {
	When When
	Then Then
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

type Mock struct {
	Name      string
	Port      int
	Behaviors []MockBehavior
}

type MockBehavior struct {
	Given MockGiven
	Then  MockThen
}

type MockGiven struct {
	Method string
	Path   string
	Format string
	Body   string
}

type MockThen struct {
	Body string
}
