package embdr_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"text/template"

	"github.com/heppu/embdr"
)

func TestTemplate(t *testing.T) {
	filename := "template.tmpl"

	expected, err := ioutil.ReadFile(filename)
	noErrorf(t, err, "failed to read %s", filename)

	got, err := embdr.Load(filename)
	noErrorf(t, err, "failed get template for %s", filename)

	isTruef(t, string(expected) == got, "content of file %s doesn't match stored data", filename)
}

func TestTemplateNotFound(t *testing.T) {
	_, err := embdr.Load("foo")
	isTruef(t, err == embdr.ErrTemplateDoesNotExist, "expected 'embdr.ErrTemplateDoesNotExist' error, got '%s'", err)
}

func ExampleTemplate() {
	const name = "template.tmpl"
	tmpl, err := embdr.Load(name)
	if err != nil {
		log.Fatalf("Couldn't get template with name: %s", err)
	}

	t, err := template.New(name).Parse(tmpl)
	if err != nil {
		log.Fatalf("Couldn't load internal templates: %s", err)
	}

	type GenParams struct {
		Package   string
		Templates map[string]string
	}

	if err := t.Execute(os.Stdout, GenParams{}); err != nil {
		log.Fatalf("Couldn't generate output: %s", err)
	}
}
