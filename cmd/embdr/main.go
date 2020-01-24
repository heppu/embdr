package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/heppu/embdr"
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

	var err error
	r := os.Stdin
	w := os.Stdout

	if input != "" {
		if r, err = os.Open(input); err != nil {
			logAndExit("Couldn't open input file: %s", err)
		}
		defer r.Close()
	}

	if output != "" {
		if r, err = os.Create(output); err != nil {
			logAndExit("Couldn't create output file: %s", err)
		}
		defer w.Close()
	}

	enc, err := embdr.NewEncoder(w)
	if err != nil {
		logAndExit("Couldn't create new encoder: %s", err)
	}
	defer enc.Close()

	if _, err := bufio.NewReader(r).WriteTo(enc); err != nil {
		logAndExit("Couldn't encode data: %s", err)
	}
}

func logAndExit(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
