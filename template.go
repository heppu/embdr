// Code generated with go generate; DO NOT EDIT.

package embdr

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"errors"
	"io/ioutil"
)

var ErrTemplateDoesNotExist = errors.New("template doesn't exists")

var templates = map[string]string{
	"template.tmpl": `
H4sIAAAAAAAC/4xSTWvbQBA9a37FRAQqgSpdSg8JPrSxD704pfEtBLLWjuXF0q6YXcdNxf73sruyDoVC
DsbL03tv3nw0DT4YSdiRJhaOJF6UO2JnFuQe14+4fdzhZv1jVwOMoj2JjnCasP45v70HUMNo2GEBWb5/
d2RzyPLWDCOTtU33R40BIN0aqXTX7IWlr18ixGw4spVplDk71edQArwJxg3zjoaxF47WhuzWuM1vZR2u
MKnqLV2K3M0UlIas/uSQAsnms8n1s8UVDmJ8to6V7l7S3wTT9BlZ6I7w9kTvFd6+if5MeLfC+lrbeg9Z
Pk2Rgd7nd/gK05SY3r9W0YS09B48QNPgVYlM7szaYmu0I+3QHJCGPUlJcgmGlyMxoRYDobLojhR/o3BH
HIW1gWqQBrnnGg5n3S72RdSkTkos0qNKwylxgkxSayTJCs0pdLSM4jkIXyBTB7wxp8DMUlLM8+p/Y4fM
A2Qc/YNbWGrYwC8SkrhIKw3AOlZdkCcnN/PeK4y3EUjfz4cD8VOMXMw5y7KMmUKBmxVq1f8bjZhTjHhJ
obM5TLqcOmT51vcFf9hoZt2tkOuH3lgqyvuP6GYszbyY45RV0ICHvwEAAP//+0zduVgDAAA=`,
}

// Template returns content of embedded template where name is the the path passed to emdbr.
func Template(name string) (string, error) {
	decoded, ok := templates[name]
	if !ok {
		return "", ErrTemplateDoesNotExist
	}

	r, err := gzip.NewReader(base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(decoded)))
	if err != nil {
		return "", err
	}

	encoded, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	if err := r.Close(); err != nil {
		return "", err
	}

	return string(encoded), nil
}
