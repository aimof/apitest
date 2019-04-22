package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aimof/jason"
)

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		return
	})
	http.HandleFunc("/api/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(400)
			return
		}
		v, err := jason.NewValue(r.Body)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		token := v.Get("Token")
		if token.Err != nil {
			w.WriteHeader(400)
			return
		}
		s, err := token.String()
		if err != nil {
			w.WriteHeader(400)
			return
		}
		if s != "Foo" {
			w.WriteHeader(400)
			return
		}

		name := v.Get("Name")
		if name.Err != nil {
			fmt.Fprint(w, `{"users": [1, 2, null]}`)
		}
		fmt.Fprint(w, `{"name": "Foo", "Info": "Bar"}`)
		return
	})
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
