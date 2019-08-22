package jsoncomparer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aimof/apitest"
	"github.com/aimof/apitest/logger"
	"github.com/aimof/jason"
)

// JSONComparer has no field.
type JSONComparer struct{}

// NewJSONComparer initialize JsonComparer.
func NewJSONComparer() *JSONComparer {
	return &JSONComparer{}
}

// Match response and then section
func (jc *JSONComparer) MatchResp(resp *http.Response, then apitest.Then) (bool, error) {
	if resp.StatusCode != then.Status {
		return false, fmt.Errorf("jsoncomparer.Match: Statuscodes are not match. got: %d: want: %d", resp.StatusCode, then.Status)
	}
	if then.Format == "empty" {
		return true, nil
	}
	for _, s := range then.Require {
		if match, err := jc.MatchBody(resp.Body, s); err != nil {
			return false, err
		} else if !match {
			logger.Info("Match: not match. require: " + s)
			return false, err
		}
	}

	return true, nil
}

func (jc *JSONComparer) MatchBody(body io.Reader, require string) (bool, error) {
	got, err := jason.NewValue(body)
	if err != nil {
		return false, errors.New("matchJson: body is illigal json format")
	}

	if !strings.HasPrefix(require, Match+" ") {
		return false, errors.New("matchJson: invalid format")
	}
	r := strings.TrimPrefix(require, Match+" ")
	want, err := jason.NewValue(strings.NewReader(r))
	if err != nil {
		return false, errors.New("matchJson: require is illigal json format:" + require)
	}

	match := matchValue(got, want)
	if match {
		return true, nil
	}
	logger.Debug(fmt.Sprintf("matchJSON: not Match. got: %v, want %v", got.Interface(), want.Interface()))
	return false, nil
}

// return match
func matchValue(got, want *jason.Value) bool {
	switch want.Interface().(type) {
	case string:
		if want.Interface().(string) == "" {
			return got.Interface().(string) == ""
		}
		if want.Interface().(string)[0] == '#' {
			return matchType(got, want.Interface().(string))
		}
		gotString, ok := got.Interface().(string)
		if !ok {
			return false
		}
		return gotString == want.Interface().(string)
	case json.Number:
		gotNumber, ok := got.Interface().(json.Number)
		if !ok {
			return false
		}
		return gotNumber == want.Interface().(json.Number)
	case bool:
		return got.Interface().(bool) == want.Interface().(bool)
	case []interface{}:
		wantSlice, ok := want.Interface().([]interface{})
		if !ok {
			return false
		}
		gotSlice, ok := got.Interface().([]interface{})
		if !ok {
			return false
		}

		if len(gotSlice) > len(wantSlice) {
			return false
		}

		for i := range wantSlice {
			wantV := want.Get(i)
			if wantV.Err != nil {
				return false
			}
			gotV := got.Get(i)
			if gotV.Err != nil {
				return false
			}
			if !matchValue(gotV, wantV) {
				return false
			}
		}
	case map[string]interface{}:
		wantChildren, err := want.GetAll()
		if err != nil {
			return false
		}
		for k, wantV := range wantChildren {
			gotV := got.Get(k)
			if gotV.Err != nil {
				return false
			}
			if !matchValue(gotV, wantV) {
				return false
			}
		}
		return true

	}

	return false
}

func matchType(got *jason.Value, want string) bool {
	if got.Interface() == nil {
		return want == NullType
	}
	switch got.Interface().(type) {
	case string:
		return want == StringType
	case json.Number:
		return want == NumberType
	case bool:
		return want == BoolType
	case map[string]interface{}:
		return want == ObjectType
	case []interface{}:
		return want == ArrayType
	default:
		return false
	}
}
