package embdr_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/heppu/embdr"
)

func TestEncodeFiles(t *testing.T) {
	filenames := []string{"testdata/valid/some_text.tmpl", "testdata/valid/some_html.tmpl"}

	files, err := embdr.EncodeFiles(filenames...)
	noErrorf(t, err, "encoding failed")

	for _, name := range filenames {
		content, err := ioutil.ReadFile(name)
		noErrorf(t, err, "reading file %s failed", name)

		parsed, ok := files[name]
		isTruef(t, ok, "map should contain %s", name)

		decoded, err := embdr.Decode(parsed)
		noErrorf(t, err, "decoding %s failed", name)

		isTruef(t, string(decoded) == string(content), "content of file %s doesn't match stored data", name)
	}
	isTruef(t, len(files) == len(filenames), "map should have %d entries but it has %d", len(filenames), len(files))
}

func TestEncodeNonExistingFile(t *testing.T) {
	_, err := embdr.EncodeFiles("testdata/exist/do/i/not")
	isTruef(t, err != nil, "Encoding non existing file should have failed")
}

func TestEncodeNonUTF8File(t *testing.T) {
	_, err := embdr.EncodeFiles("testdata/invalid")
	isTruef(t, err != nil, "Encoding non utf8 data should have failed")
}

func TestNLWriter_Write(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
		limit          int
	}{
		{
			name:           "underLimit",
			input:          "somedata",
			expectedOutput: "somedata",
			limit:          9,
		}, {
			name:           "overLimitOnce",
			input:          "somedata",
			expectedOutput: "somedat\na",
			limit:          7,
		}, {
			name:           "overLimitTwice",
			input:          "somedata",
			expectedOutput: "som\neda\nta",
			limit:          3,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			w := embdr.NewNLWriter(buf, tt.limit)
			_, err := w.Write([]byte(tt.input))
			noErrorf(t, err, "writing failed")
			isTruef(t, buf.String() == tt.expectedOutput, "input doesn't match expexted output\ninput:\n'%s'\noutput:\n'%s'", tt.input, buf.String())
		})
	}
}

func TestNLWriter_WriteMultipleTimes(t *testing.T) {
	expectedOutput := "he\nll\not\nhe\nre"

	buf := &bytes.Buffer{}
	w := embdr.NewNLWriter(buf, 2)

	_, err := w.Write([]byte("hello"))
	noErrorf(t, err, "writing failed")

	_, err = w.Write([]byte("there"))
	noErrorf(t, err, "writing failed")

	isTruef(t, buf.String() == expectedOutput, "input doesn't match expexted output\ninput:\n'%s'\noutput:\n'%s'", expectedOutput, buf.String())
}

func noErrorf(t testing.TB, err error, format string, args ...interface{}) {
	if err != nil {
		t.Logf(format, args...)
		t.Fatal("Error:", err)
	}
}

func isTruef(t testing.TB, b bool, format string, args ...interface{}) {
	if !b {
		t.Fatalf(format, args...)
	}
}
