// Package embdr provides functions for encoding and decoding data in gz compressed base64 format.
package embdr

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"unicode/utf8"
)

// EncodeFiles encodes multiple files to map and uses filename as a key.
func EncodeFiles(filenames ...string) (map[string]string, error) {
	output := make(map[string]string, len(filenames))
	for _, filename := range filenames {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		encoded, err := Encode(data)
		if err != nil {
			return nil, fmt.Errorf("encoding %s failed: %w", filename, err)
		}

		output[filename] = string(encoded)
	}

	return output, nil
}

func Encode(data []byte) ([]byte, error) {
	if !utf8.Valid(data) {
		return nil, fmt.Errorf("invalid utf8 input")
	}

	buf := &bytes.Buffer{}
	enc := base64.NewEncoder(base64.StdEncoding, NewNLWriter(buf, 80))
	w, err := gzip.NewWriterLevel(enc, gzip.BestCompression)
	if err != nil {
		return nil, fmt.Errorf("couldn't create encoder: %w", err)
	}

	if _, err := w.Write(data); err != nil {
		return nil, err
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	if err := enc.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Decode(s string) ([]byte, error) {
	r, err := gzip.NewReader(base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(s)))
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if err := r.Close(); err != nil {
		return nil, err
	}

	return data, nil
}

type NLWriter struct {
	w          io.Writer
	buf        *bytes.Buffer
	lineLength int
	c          int
}

func NewNLWriter(w io.Writer, lineLength int) *NLWriter {
	return &NLWriter{
		w:          w,
		buf:        bytes.NewBuffer(make([]byte, 0, lineLength)),
		lineLength: lineLength,
	}
}

func (w *NLWriter) Write(p []byte) (int, error) {
	w.buf.Reset()
	last := len(p) - 1
	for i, b := range p {
		_ = w.buf.WriteByte(b)
		w.c++
		if w.c == w.lineLength && i != last {
			_ = w.buf.WriteByte('\n')
			w.c = 0
		}
	}

	return w.w.Write(w.buf.Bytes())
}
