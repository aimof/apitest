package mockbuilder

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aimof/apitest"
)

type MockBuilder struct {
	server *http.Server
}

func NewMockBuilder() *MockBuilder {
	return &MockBuilder{
		server: &http.Server{},
	}
}

func (mb *MockBuilder) Run(mock apitest.Mock, comp apitest.Comparer) {
	mb.server.Addr = ":" + strconv.Itoa(mock.Port)
	mb.server.Handler = mb.buildMux(mock.Behaviors, comp)
	go mb.server.ListenAndServe()
}

func (mb *MockBuilder) buildMux(behaviors []apitest.MockBehavior, comp apitest.Comparer) *http.ServeMux {
	mux := http.NewServeMux()
	for _, behavior := range behaviors {
		path, handler := handleFunc(behavior, comp)
		mux.HandleFunc(path, handler)
	}

	return mux
}

func handleFunc(behavior apitest.MockBehavior, comp apitest.Comparer) (string, func(http.ResponseWriter, *http.Request)) {
	return behavior.Given.Path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != behavior.Given.Method {
			w.WriteHeader(500)
			return
		}
		if behavior.Given.Format == "empty" {
			fmt.Fprint(w, behavior.Then.Body)
			return
		}

		match, err := comp.MatchBody(r.Body, behavior.Given.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		if match {
			fmt.Fprint(w, behavior.Then.Body)
			return
		}
		w.WriteHeader(500)
	}
}
