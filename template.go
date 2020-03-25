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
H4sIAAAAAAAC/4xSTWvbQBA9a37FRAQqgSpdSg8OPrSxD4XilMa3EMhaO5YXS7tidh03Ffvfy+7KhhYK
ORgvT++9efPRNHhvJGFHmlg4knhW7oCduSJ3uHrAzcMW16tv2xpgFO1RdITThPWP+e09gBpGww4LyPLd
myObQ5a3ZhiZrG2632oMAOnWSKW7Zicsff4UIWbDka1Mo8zJqT6HEuBVMK6ZtzSMvXC0MmQ3xq1/Ketw
iUlVb+hc5G6moDRk9QeHFEg2n00uny0ucRDjk3WsdPec/iaYpo/IQneEt0d6q/D2VfQnwsUS60tt6z1k
+TRFBnqfL/AFpikxvX+poglp6T14gKbB70bIixqZ3Im1xdZoR9qh2SMNO5KS5DUcng/EhFoMhMqiO1D8
jcIdcBTWBqpBGuSOa9ifdPtXiSLqUkclFulRpSGVOEEmqTWSZIXmGDq7juQpCJ8hU3u8McfAzFJazPPq
f+OHzANkHP2DW1hu2MRPEpK4SKsNwCpWvSKPTq7n/VcYbySQvp72e+LHGLmYc5ZlGTOFAjdL1Kr/Nxox
pxjxokJnc5h0QXXI8qXvC3630cxaLJHr+95YKsq79+hmLM28mOOUVdCAhz8BAAD//11Q6DZgAwAA`,
}

// LoadTemplate returns content of embedded template where name is the the path passed to emdbr.
func LoadTemplate(name string) (string, error) {
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
