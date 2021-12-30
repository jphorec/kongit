package main

import (
	"testing"

	"github.com/Kong/go-pdk/test"
	{{.NameImport}} {{.Module}}
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
