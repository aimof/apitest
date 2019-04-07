package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(400)
			return
		}
		if r.Header.Get("Wrong") != "" {
			w.WriteHeader(400)
			return
		}
		if r.Header.Get("Long") == "required" {
			fmt.Fprint(w, strings.Repeat("abcdefghijklmnopqrstuvwxyz\n", 30))
			return
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		b = bytes.TrimRight(b, "\x0a\x0d")
		if string(b) == "Hello, world!" {
			fmt.Fprint(w, "こんにちは、世界")
			return
		}
		fmt.Fprint(w, "Hello, world!")
	})
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
