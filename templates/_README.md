# {{.Name}} Kong plugin

This plugin was generated with [kongit](https://github.com/FigureTechnologies/kongit)

# Development

This generated project brings in the Kong Golang developer package docs can be found here https://pkg.go.dev/github.com/Kong/go-pdk

Each Kong plugin runs its own server and has a socket connection back to Kong for communication. A plugin can have configuration paramters that 
are defined as a struct in your main plugin file `{{.Name}}.go`. 

If your dockerfile yaml is missing the configuration param definition then the server will fail to start. 

## Adding and changing congifugrations

The default config object is set as 
```
type Config struct {
	Param string `json:"param"`
}
```

### Note

- New parameters need a `json` attribute definition so the kong package knows how to deserialize the values when injecting them into a new plugin object at run time. 
- Parameter naming also require capitalized variables to make the values visible to the plugin library

### Configuration parameters are defined on the Kong config yaml as follows: 

```
_format_version: "1.1"
services:
- name: plugin-example
  url: https://example.com/
  routes:
  - paths:
    - "/"
  plugins:
  - name: {{.Name}}
    config:
      param: this-is-my-param
```

In this example service, we are using our plugin to intercept routes with the path `/` and 
by default having the service redirect to `https://example.com/` 

# How to run

A Dockerfile is generated as part of the plugin along with `Make` commands to start the docker instance and pass in the config yaml. 

## Make commands

```
make docker && make docker-run
```