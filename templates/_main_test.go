package main

import (
	"testing"

	{{.NameImport}} "{{.Module}}"
	"github.com/Kong/go-pdk/test"
	"github.com/stretchr/testify/assert"
)


var config = &{{.NameImport}}.Config{
	Param: "this-is-my-param",
}

func TestExample(t *testing.T) {
	env, err := test.New(t, test.Request{
		Method: "GET",
		Url:    "http://example.com",
	})
	assert.NoError(t, err)

	env.DoHttp(config)

	assert.Equal(t, 200, env.ClientRes.Status)
}
