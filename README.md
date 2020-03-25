# Embdr

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
