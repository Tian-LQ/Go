package errors

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHTTPErrorStruct(t *testing.T) {
	a := &HTTPError{
		Errors: make(map[string][]string),
	}
	a.Errors["body"] = []string{"can't be empty"}
	a.Code = 500
	b, err := json.Marshal(a)
	assert.NoError(t, err)
	t.Log(string(b))
}
