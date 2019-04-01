package yamlmapper

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
)

func TestToTask(t *testing.T) {
	ym := YamlMapper{
		tasks: Tasks{
			Config: Config{
				workdir:  "./test",
				Timeout:  30,
				WildCard: "*",
			},
			Testcases: []Testcase{
				{
					Name:     "POSTSomething",
					URL:      "http://localhost/",
					Method:   "POST",
					Retry:    -1,
					Header:   []string{"Foo:Bar:Fizz"},
					BodyPath: "./postSomethingBody.txt",
					Want: Want{
						Statuscode: 200,
						BodyPath:   "./postSomethingPattern.txt",
					},
				},
			},
		},
	}
	got, err := ym.toApitestTasks()
	if err != nil {
		t.Error(err)
		return
	}
	if len(got) != 1 {
		t.Error()
		return
	}
	getTask := got[0]
	if getTask.Name != "POSTSomething" || getTask.Retry != -1 {
		t.Error()
	}
	n, b := getTask.Want.Want()
	if n != 200 || string(b) != "1\n2\n*Fizz" {
		t.Errorf("got: %d %v", n, b)
	}
	req := getTask.Request
	if req.Method != "POST" || req.URL.String() != "http://localhost/" {
		t.Error()
	}
	if h, ok := req.Header["Foo"]; !ok {
		t.Error()
	} else {
		if !reflect.DeepEqual(h, []string{"Bar", "Fizz"}) {
			t.Error()
		}
	}
	b, err = ioutil.ReadAll(req.Body)
	if err != nil {
		t.Error()
	}
	if string(b) != "Hello, world!" {
		t.Errorf("got: %v", b)
	}
}

func TestParse(t *testing.T) {
	workDir, err := filepath.Abs(filepath.Base("./test/"))
	if err != nil {
		t.Error(err)
		return
	}
	want := Tasks{
		Config: Config{
			workdir: workDir,
			Timeout: 30,
		},
		Testcases: []Testcase{
			{
				Name:   "GetSomething",
				URL:    "http://localhost/",
				Method: "GET",
				Retry:  -1,
				Want: Want{
					Statuscode: 200,
					BodyPath:   "foo*bar",
				},
			},
			{
				Name:   "PostSomething",
				URL:    "http://localhost/api",
				Method: "POST",
				Header: []string{
					"key:value",
				},
				BodyPath: "foo.txt",
				Want: Want{
					Statuscode: 200,
					BodyPath:   "*",
				},
			},
		},
	}
	ym := new(YamlMapper)
	err = ym.parse("./test/mapper_test.yaml")
	if err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(ym.tasks, want) {
		t.Errorf("\ngot:  %v\nwant: %v", ym.tasks, want)
	}
}

func TestSliceToHeader(t *testing.T) {
	s := []string{
		"Key:Value0:Value1",
		"",
		"Foo:Bar",
	}
	h := sliceToHeader(s)
	if len(h) != 2 {
		t.Error()
	}
	if v, ok := h["Key"]; !ok {
		t.Error()
	} else if !reflect.DeepEqual(v, []string{"Value0", "Value1"}) {
		t.Errorf("got:  %v", v)
	}
	if v, ok := h["Foo"]; !ok {
		t.Error()
	} else if !reflect.DeepEqual(v, []string{"Bar"}) {
		t.Errorf("got:  %v", v)
	}
}
