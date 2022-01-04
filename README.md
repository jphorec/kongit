# Kong Plugin Generator (Kongit)

This plugin generator helps scaffold a Kong plugin with Golang support. 

## Getting started

### Installing

Install the plugin using `go get` 

```
// set the private organization in your .bashrc/.zrc
export GOPRIVATE=github.com/FigureTechnologies

go install github.com/FigureTechnologies/kongit@latest
```

Building from source

```
git clone https://github.com/FigureTechnologies/kongit.git

cd kongit

go install
```

### How to use

Kongit has two required flags: `--module` and `--name`. Module is the organization your new plugin will belong to. 
Name is the name of the plugin you would like to generate 

Example: 

```aidl
kongit -m github.com/FigureTechnologies/redirect -n redirect
```

This example command will generate the following directory

```aidl
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
along with a `Makefile` for building a test docker image and running the docker image

Available flags: 

```aidl
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

