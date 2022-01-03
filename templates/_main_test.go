package main

import (
	"testing"

	// currently dealing with an incompatibility issue with kong docker image
	// 2.4.1 and the latest go-pdk version
	// TODO: upgrade version once kong 2.6.0 is supported at Figure Tech
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
