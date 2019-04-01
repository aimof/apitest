package apitest

import "net/http"

// Tasks is all tasks to do.
type Tasks []Task

// Task is a single step with single http request.
type Task struct {
	Name     string
	Request  *http.Request
	CompMode string
	Retry    int
	Want     WantIface
}

// Result is returned from this package.
type Result struct {
	Name  string
	Got   GotIface
	Match bool
}
