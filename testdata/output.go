package mypkg

import (
	"errors"
	"fmt"
	"html/template"
)

const (
	a = "asd"
	b = "asd"
)

var ErrNoTemplates = errors.New("no templates")

func ParseTemplates(t *template.Template) (*template.Template, error) {
	if len(templataMap) == 0 {
		return nil, fmt.Errorf("no templates to parse")
	}

	for name, data := range templataMap {
		if t == nil {
			t = template.New(name)
		}

		var tmpl *template.Template
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}

		if _, err := tmpl.Parse(data); err != nil {
			return nil, err
		}
	}

	return t, nil
}
