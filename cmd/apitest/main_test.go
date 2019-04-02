package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/aimof/apitest/mapper/yamlmapper"
)

func TestApitest(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
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
	ts := httptest.NewServer(mux)

	tasks, err := yamlmapper.NewYamlMapper().Tasks("./test/test.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	for i := range tasks {
		tasks[i].Request.URL, err = url.Parse(strings.Replace(tasks[i].Request.URL.String(), "$HOST", ts.URL, 1))
		if err != nil {
			t.Error(err)
		}
	}

	err = do(tasks)
	if err != nil {
		t.Error(err)
	}
}
