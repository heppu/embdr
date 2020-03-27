# Embdr
[![Documentation](https://godoc.org/github.com/heppu/embdr?status.svg)](https://pkg.go.dev//github.com/heppu/embdr)
[![Go Report Card](https://goreportcard.com/badge/github.com/heppu/embdr)](https://goreportcard.com/report/github.com/heppu/embdr)
[![codecov](https://codecov.io/gh/heppu/embdr/branch/master/graph/badge.svg)](https://codecov.io/gh/heppu/embdr)
[![Release](https://img.shields.io/github/release/heppu/embdr.svg?label=Release)](https://github.com/heppu/embdr/releases)
[![GitHub issues](https://img.shields.io/github/issues/heppu/embdr.svg)](https://github.com/heppu/embdr/issues)
[![license](https://img.shields.io/github/license/heppu/embdr.svg?maxAge=2592000)](https://github.com/heppu/embdr/LICENSE)


Emdeb Go templates with zero 3rd party dependencies.

Just generate and load.

Generate:
```
embdr -p mypkg -o templates.go index.tmpl`
```

Load:
```go
tmpl, err := LoadTemplate("index.tmpl")
```

## Install

go get github.com/heppu/embdr/cmd/embdr

## Usage


```
embdr -p <package_name> [-o output_file] [file_1 file_2 ...]

Flags:

	-p  package name
	-o  output file (default STDOUT)
```

### Embed single file

```bash
embdr -p mypkg -o templates.go index.tmpl
```

### Embed all files from folder

```bash
find ./templates/ -name '*.tmpl' | ./embdr -p mypkg -o templates.go
```

### Using go generate directive

Define template file.

```
{{- /* user.tmpl */ -}}

Hi there {{.Name}}!
```

Then create file that has the //go:generate directive.

```go
// main.go

package main

//go:generate embdr -p main -o templates.go user.tmpl

import (
	"log"
	"os"
	"text/template"
)

type User struct {
	Name string
}

func main() {
	const name = "user.tmpl"

	data, err := LoadTemplate(name)
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New(name).Parse(data)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(os.Stdout, User{Name: "Gopher"})
	if err != nil {
		log.Fatal(err)
	}
}
```

Now just generate embedded templates with go generate and you are ready to go!

```bash
$ go generate ./...
$ go run main.go templates.go
Hi there Gopher!
```

## Why one more embedding tool?

There are almost as many static asset embedding tools for go as there are http routers. Why I decided to make one more was to have something extremely simple with no additional knobs and pulls. Just give the input, ouput and package name and don't care about anything else. I also didn't want force users to use any external 3rd party dependencies.
