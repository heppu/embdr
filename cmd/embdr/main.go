package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/heppu/embdr"
)

type GenParams struct {
	Package   string
	Templates map[string]string
}

func main() {
	var (
		pkgName string
		output  string
	)

	flag.StringVar(&output, "o", "", "name of generated output file (default STDOUT)")
	flag.StringVar(&pkgName, "p", "", "name of generated package")
	flag.Parse()

	if pkgName == "" {
		flag.Usage()
		os.Exit(1)
	}

	inputFiles := flag.Args()
	if len(inputFiles) == 0 {
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			logAndExitf("Could not read filenames from STDIN: %s", err)
		}
		inputFiles = strings.Fields(string(input))
	}

	w := os.Stdout
	if output != "" {
		var err error
		if w, err = os.Create(output); err != nil {
			logAndExitf("Couldn't create output file: %s", err)
		}
		defer w.Close()
	}

	data, err := embdr.EncodeFiles(inputFiles...)
	if err != nil {
		logAndExitf("Encoding failed: %s", err)
	}

	const name = "template.tmpl"
	tmpl, err := embdr.Template(name)
	if err != nil {
		logAndExitf("Couldn't get template with name: %s", err)
	}

	t, err := template.New(name).Parse(tmpl)
	if err != nil {
		logAndExitf("Couldn't load internal templates: %s", err)
	}

	if err := t.Execute(w, GenParams{Package: pkgName, Templates: data}); err != nil {
		logAndExitf("Couldn't generate output: %s", err)
	}
}

func logAndExitf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
