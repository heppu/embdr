package main

import (
	"flag"
	"os"
	"testing"
)

func Test_main(_ *testing.T) {
	original := make([]string, len(os.Args))
	copy(original, os.Args)

	defer func() { os.Args = original }()

	os.Args = []string{"emdbr", "-p", "mypkg"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	main()

	os.Args = []string{"emdbr", "-p", "mypkg", "main.go"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	main()

	os.Args = []string{"emdbr", "-p", "mypkg", "-o", "../../target/out.go", "main.go"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	main()
}
