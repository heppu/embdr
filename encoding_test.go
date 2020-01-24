package embdr_test

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/heppu/embdr"
)

func TestEncodeAndDecodeFiles(t *testing.T) {
	files := []string{"testdata/other.tmpl", "testdata/some.tmpl"}

	data, err := embdr.EncodeFiles(files...)
	noError(t, err, "encoding failed")

	if len(data) != len(files) {
		t.Fatalf("map should have %d entries but has %d", len(files), len(data))
	}

	for _, fullName := range files {
		name := path.Base(fullName)
		b, ok := data[name]
		if !ok {
			t.Fatalf("map should contain %v", name)
		}

		decoded, err := embdr.DecodeString(b)
		noError(t, err, "decoding %s failed", name)

		fileContent, err := ioutil.ReadFile("testdata/" + name)
		noError(t, err, "reading file %s failed", name)

		if string(decoded) != string(fileContent) {
			t.Fatalf("file %s doesn't match after decoding\nfile:\n%s\ndecoded:\n%s", name, fileContent, decoded)
		}
	}
}

func TestEncodeDir(t *testing.T) {
	dir := "testdata"
	data, err := embdr.EncodeDir(dir)
	noError(t, err, "encoding failed")

	files, err := ioutil.ReadDir(dir)
	noError(t, err, "reading files failed")

	if len(data) != len(files) {
		t.Fatalf("map should have %d entries but has %d", len(files), len(data))
	}

	for _, file := range files {
		name := file.Name()
		b, ok := data[name]
		if !ok {
			t.Fatalf("map should contain %v", name)
		}

		decoded, err := embdr.DecodeString(b)
		noError(t, err, "decoding %s failed", name)

		fileContent, err := ioutil.ReadFile("testdata/" + name)
		noError(t, err, "reading file %s failed", name)

		if string(decoded) != string(fileContent) {
			t.Fatalf("file %s doesn't match after decoding\nfile:\n%s\ndecoded:\n%s", name, fileContent, decoded)
		}
	}
}

func TestEncodeNonExistingFile(t *testing.T) {
	_, err := embdr.EncodeFiles("testdata/exist/do/i/not")
	if err == nil {
		t.Fatal("Encoding non existing file should have failed")
	}
}

func TestDecodingMalformedData(t *testing.T) {
	_, err := embdr.DecodeString("non-base64-string")
	if err == nil {
		t.Fatal("Decoding invalid data should have failed")
	}
}

func noError(t testing.TB, err error, format string, args ...interface{}) {
	if err != nil {
		t.Logf(format, args...)
		t.Fatal("Error:", err)
	}
}
