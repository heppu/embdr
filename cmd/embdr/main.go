package main

import (
	"flag"
	"io"
	"os"
)

func main() {
	var (
		packageName string
		output      string
		input       string
		public      bool
	)

	flag.StringVar(&input, "in", "", "input file or directory (default STDIN)")
	flag.StringVar(&output, "out", "", "output file (default STDOUT)")
	flag.StringVar(&packageName, "pkg", "", "name of generated package if not set only the data part generated")
	flag.BoolVar(&public, "public", false, "makes generared data public")
	flag.Parse()

	var (
		r io.Reader
		w io.Writer
	)

	if input == "" {
		r = os.Stdin
	}

	if output == "" {
		w = os.Stdout
	}

	r.Read([]byte{})
	w.Write([]byte{})
}
