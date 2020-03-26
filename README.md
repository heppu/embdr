# Embdr
[![Documentation](https://godoc.org/github.com/heppu/embdr?status.svg)](https://pkg.go.dev//github.com/heppu/embdr)
[![Go Report Card](https://goreportcard.com/badge/github.com/heppu/embdr)](https://goreportcard.com/report/github.com/heppu/embdr)
[![codecov](https://codecov.io/gh/heppu/embdr/branch/master/graph/badge.svg)](https://codecov.io/gh/heppu/embdr)
[![Release](https://img.shields.io/github/release/heppu/embdr.svg?label=Release)](https://github.com/heppu/embdr/releases)
[![GitHub issues](https://img.shields.io/github/issues/heppu/embdr.svg)](https://github.com/heppu/embdr/issues)
[![license](https://img.shields.io/github/license/heppu/embdr.svg?maxAge=2592000)](https://github.com/heppu/embdr/LICENSE)


Simplest static asset embedding tool for Go with zero 3rd party dependencies.

Just generate and load.

Generate:
```
embdr -p mypkg -o templates.go index.tmpl`
```

Load:
```
tmpl, err := LoadTemplate(name)
```

## Install

go get github.com/heppu/embdr/cmd/embdr

## Usage

### Embed single file

```bash
embdr -p mypkg -o templates.go index.tmpl
```

### Embed all files from folder

```bash
find ./templates/ -name '*.tmpl' | ./embdr -p mypkg -o templates.go
```

### Using go generate directive

```go
package mypkg

//go:generate embdr -p mypkg -o templates.go user.tmpl user.tmpl

import "text/template"

var userTmpl template.Template

func init() {
	const name = "user.tmpl"

	tmpl, err := LoadTemplate(name)
	if err != nil {
		log.Fatal(err)
	}

	userTmpl, err := template.New(name).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
}

func RenderUser(user User) error {
	return t.Execute(w, user)
}
```

```bash
go generate ./...
```
