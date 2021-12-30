package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Project struct {
	Dir        string
	Name       string
	NameImport string
	Org        string
}

type TemplateMapping struct {
	template string
	path     string
}

var org string
var name string
var dir string

//go:embed templates/*
var templates embed.FS

func init() {
	flag.StringVar(&name, "name", "", "Required -- The name of the plugin you want to create")
	flag.StringVar(&name, "n", "", "Required -- name shorthand")

	flag.StringVar(&org, "org", "", "Required -- The name of the github organization your plugin will belong to")
	flag.StringVar(&org, "o", "", "Required -- org shorthand")

	flag.StringVar(&dir, "directory", "", "Optional -- The name of the directory for your plugin - defaults to plugin name")
	flag.StringVar(&dir, "d", name, "Optional -- directory shorthand")
}
func main() {
	flag.Parse()
	if len(org) == 0 {
		fmt.Println("--org is required")
		os.Exit(1)
	}

	if len(name) == 0 {
		fmt.Println("--name is required")
		os.Exit(1)
	}

	if len(dir) == 0 {
		dir = name
	}

	cmdPath := "cmd/{project}"

	nameImport := name

	pluginEntry := strings.ReplaceAll(cmdPath, "{project}", name)

	var templateMap = []TemplateMapping{
		{"templates/_main.go", fmt.Sprintf("%s/%s/main.go", dir, pluginEntry)},
		{"templates/_plugin.go", fmt.Sprintf("%s/%s.go", dir, name)},
		{"templates/_go.mod", fmt.Sprintf("%s/go.mod", dir)},
		{"templates/_Dockerfile", fmt.Sprintf("%s/Dockerfile", dir)},
		{"templates/_main_test.go", fmt.Sprintf("%s/%s/main_test.go", dir, pluginEntry)},
		{"templates/_Makefile", fmt.Sprintf("%s/Makefile", dir)},
		{"templates/_README.md", fmt.Sprintf("%s/README.md", dir)},
		{"templates/_config.yml", fmt.Sprintf("%s/config.yml", dir)},
	}

	if strings.ContainsAny(name, "-") {
		nameImport = strings.ReplaceAll(name, "-", "")
	}
	project := &Project{
		Dir:        dir,
		Name:       name,
		NameImport: nameImport,
		Org:        org,
	}
	os.Mkdir(dir, 0755)
	os.MkdirAll(dir+"/"+pluginEntry, 0755)
	for _, mapping := range templateMap {
		generateFile(*project, mapping.template, mapping.path)
	}
}

func generateFile(project Project, fileTemplate string, path string) bool {
	var buf bytes.Buffer
	t, err := template.ParseFS(templates, fileTemplate)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	err = t.Execute(&buf, project)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return WriteToFile(buf.Bytes(), path)
}

func WriteToFile(input []byte, file string) bool {
	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer f.Close()

	f.Write(input)
	return true
}
