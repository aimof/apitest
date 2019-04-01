package yamlmapper

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aimof/apitest"
	"github.com/go-yaml/yaml"
)

// YamlMapper .
type YamlMapper struct {
	tasks Tasks
}

// NewYamlMapper empty.
func NewYamlMapper() *YamlMapper {
	return &YamlMapper{}
}

// Tasks return apitest Tasks
func (ym *YamlMapper) Tasks(path string) (apitest.Tasks, error) {
	err := ym.parse(path)
	if err != nil {
		return nil, err
	}
	return ym.toApitestTasks()
}

func (ym *YamlMapper) toApitestTasks() (apitest.Tasks, error) {
	if ym == nil {
		return nil, errors.New("toApitest: YamlMapper.tasks is nil")
	}
	tasks := make(apitest.Tasks, 0, len(ym.tasks.Testcases))
	for i := range ym.tasks.Testcases {
		t, err := ym.toTask(i)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// W is apitest.WIface
type W struct {
	statuscode int
	body       []byte
}

// Want ia apitets.WIface.Want().
func (w W) Want() (int, []byte) { return w.statuscode, w.body }

func (ym *YamlMapper) toTask(i int) (apitest.Task, error) {
	if ym == nil {
		return apitest.Task{}, errors.New("toTask: ym is nil")
	}
	if i > len(ym.tasks.Testcases) {
		return apitest.Task{}, errors.New("toTask: i > ym.tasks.Testcase")
	}
	tc := ym.tasks.Testcases[i]
	task := apitest.Task{
		Name:  tc.Name,
		Retry: tc.Retry,
	}

	b := make([]byte, 0)
	if tc.BodyPath != "" {
		f, err := os.Open(filepath.Join(ym.tasks.Config.workdir, tc.BodyPath))
		if err != nil {
			return apitest.Task{}, err
		}
		b, err = ioutil.ReadAll(f)
		if err != nil {
			return apitest.Task{}, err
		}
		f.Close()
	}

	var err error
	task.Request, err = http.NewRequest(tc.Method, tc.URL, bytes.NewReader(bytes.TrimRight(b, "\x0a\x0d")))
	if err != nil {
		return apitest.Task{}, err
	}
	task.Request.Header = sliceToHeader(tc.Header)

	f, err := os.Open(filepath.Join(ym.tasks.Config.workdir, tc.Want.BodyPath))
	if err != nil {
		return apitest.Task{}, err
	}
	b, err = ioutil.ReadAll(f)
	if err != nil {
		return apitest.Task{}, err
	}
	f.Close()

	task.Want = W{tc.Want.Statuscode, []byte(bytes.TrimRight(b, "\x0a\x0d"))}
	return task, nil
}

// Map yaml file on path as operation
func (ym *YamlMapper) parse(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	f.Close()

	var t = Tasks{}
	err = yaml.Unmarshal(b, &t)
	if err != nil {
		return err
	}
	ym.tasks = t

	abs, err := filepath.Abs(filepath.Dir(path))
	if err != nil {
		return err
	}
	ym.tasks.Config.workdir = abs

	return err
}

func sliceToHeader(headerString []string) http.Header {
	header := make(http.Header, len(headerString))
row:
	for _, s := range headerString {
		values := strings.Split(s, ":")
		if len(values) < 2 {
			continue row
		}
		// if the format is invalid, header doesn't include the row.
		for _, value := range values {
			if value == "" {
				continue row
			}
		}
		header[values[0]] = values[1:]
	}
	return header
}
