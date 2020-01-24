// Package embdr provides functions for encoding and decoding files
// in gz compressed base64 (RFC 4648) format and embdrding them in go binaries.
package embdr

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"io/ioutil"
	"path"
)

type Data map[string]string

type Encoder struct {
	gz   *gzip.Writer
	base io.WriteCloser
}

func (e *Encoder) Write(p []byte) (n int, err error) {
	return e.gz.Write(p)
}

func (e *Encoder) Close() error {
	if err := e.gz.Close(); err != nil {
		_ = e.base.Close()
		return err
	}
	return e.base.Close()
}

// EncodeDir encodes all files in the root of given directory into a map and uses file's basename as a key.
// If multiple files have a same basename only last one will be present on returned data.
func EncodeDir(dir string) (Data, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	output := Data{}
	for _, file := range files {
		data, err := EncodeFile(path.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}
		output[file.Name()] = data
	}

	return output, nil
}

// EncodeFiles encodes multiple files to map and uses file's basename as a key.
// If multiple files have a same basename only last one will be present on returned data.
func EncodeFiles(filenames ...string) (Data, error) {
	output := Data{}
	for _, filename := range filenames {
		data, err := EncodeFile(filename)
		if err != nil {
			return nil, err
		}
		output[path.Base(filename)] = data
	}

	return output, nil
}

// EncodeFile encodes single file as gzip compressed base64 string.
func EncodeFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return EncodeBytes(data)
}

// EncodeFile encodes single file as gzip compressed base64 string.
func EncodeBytes(data []byte) (string, error) {
	buf := &bytes.Buffer{}
	wc, err := NewEncoder(buf)
	if err != nil {
		return "", err
	}

	if _, err = wc.Write(data); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// NewEncoder returns encoder which encodes data as gzip compressed base64 string.
func NewEncoder(w io.Writer) (enc *Encoder, err error) {
	enc = &Encoder{}
	enc.base = base64.NewEncoder(base64.RawStdEncoding, w)
	if enc.gz, err = gzip.NewWriterLevel(enc.base, gzip.BestCompression); err != nil {
		return nil, err
	}

	return enc, nil
}

// DecodeString decodes gzip compressed base64 string.
func DecodeString(data string) ([]byte, error) {
	r, err := NewDecoder(bytes.NewBufferString(data))
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(r)
}

// NewDecoder creates gzip compressed base64 string decoder.
func NewDecoder(r io.Reader) (io.Reader, error) {
	return gzip.NewReader(base64.NewDecoder(base64.RawStdEncoding, r))
}
