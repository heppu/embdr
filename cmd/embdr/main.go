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

const usage = `Embdr is tool for embedding templates in Go source code.


Usage:

	embdr -p <package_name> [-o output_file] [file_1 file_2 ...]

Flags:

	-p  package name
	-o  output file (default STDOUT)

Examples:

	Read index.tmpl file and encode it into templates.go file using package name 'mypkg':

		embdr -p mypkg -o templates.go index.tmpl


	Read list of files from STDIN and write encoded data to STDOUT using package name 'mypkg':

		find ./templates/ -name '*.tmpl' | ./embdr -p mypkg
`

func main() {
	var (
		pkgName string
		output  string
	)

	flag.StringVar(&pkgName, "p", "", "package name")
	flag.StringVar(&output, "o", "", "output file (default STDOUT)")

	flag.Usage = func() {
		fmt.Print(usage)
	}

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
	tmpl, err := embdr.Load(name)
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
