# Kong Plugin Generator (Kongit)

[Kong's k8s API Gateway](https://konghq.com/products/api-gateway-platform) allows plugins to run and intercept incoming/outgoing requests for special business use cases such as authorization and authentication.  Historically, these plugins were written in Lua and only Lua.

Now, Kong allows Golang plugins using the [Go Plugin Developer Kit (PDK)](https://pkg.go.dev/github.com/Kong/go-pdk) to run your plugin. However, you need to run the go program as a unix socket and use a common interface for Kong to communicate with it which can cause friction for developers. 

This plugin generator will scaffold the Kong plugin with Golang support for you and template the unix socket interface.  After generation, you will have a fully working plugin and only have to add your specific business logic.  

## Getting started

### Installing

Install the plugin using `go get` 

```
go install github.com/jphorec/kongit@latest
```

Building from source

```
git clone https://github.com/jphorec/kongit.git

cd kongit

go install
```

### How to use

Kongit has two required flags: `--module` and `--name`. Module is the organization your new plugin will belong to. 
Name is the name of the plugin you would like to generate 

Example: 

```
kongit -m github.com/jphorec/redirect -n redirect
```

#### Whats generated

This example command will generate the following directory

```
redirect
├── Dockerfile
├── Makefile
├── README.md
├── cmd
│   └── redirect
│       ├── main.go
│       └── main_test.go
├── config.yml
├── go.mod
├── go.sum
└── redirect.go
```

The generated `README.md` will include instructions on how to run your plugin 
along with a `Makefile` for building a test docker image and running the docker image. 

The Make command `make docker && make docker-run` will build the plugin, build the docker image and run an instance of Kong's API Gateway with the generated plugin fully working. 

#### Available flags

```
Usage of kongit:
  -d string
        Optional -- directory shorthand
  -directory string
        Optional -- The name of the directory for your plugin - defaults to plugin name
  -m string
        Required -- module shorthand
  -module string
        Required -- The name of the module your plugin will belong to
  -n string
        Required -- name shorthand
  -name string
        Required -- The name of the plugin you want to create

```

