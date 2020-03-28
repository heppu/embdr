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

var files = map[string]string{
	"template.tmpl": `
H4sIAAAAAAAC/4ySQWvcMBCFz55fMTGB2uDal9LDhj202T0UyqY0ewuBaK2xV6wtmZE229TovxdJ7h4K
hRyMxePN0zejaRq8N5KwJ00sHEm8KHfE3lyVO9w84O5hj9vNt30NMIn2JHrCecb6x3L2HkCNk2GHBWT5
4c2RzSHLWzNOTNY2/W81BYF0a6TSfXMQlj5/ihKz4ehWplHm7NSQQwnwKhi3zHsap0E42hiyO+O2v5R1
uMZUVe/oUuRusaA0ZPUHhxRMNl9COjWQxTWOYnqyjpXun9Nvhnn+iCx0T3h7orcKb1/FcCZcrbH+e6/1
HrJ8nqMDvc9X+ALznJzev1QxhLT0HjxA0+B3IyQyuTNri63RjrRD0yGNB5KSJF55L0diQi1GQmXRHSl+
k3BHnIS1wWqQRnngGrqzbmN0Ef2pgxKLdKjSQEqcIZPUGkmyQnMKncT2n0LRM2SqwxtzCq4sEWKeV/8b
M2QeIOOYHZLCI4aJ/yQhiYv0hEHYxBuvyqOT2+WdK4y7EExfz11H/Bhxi4WxLMvIFC64WaNWw79oxJww
4uaErhaYtCl1YPkyDAW/O2hxrdbI9f1gLBXl3XvqFi3Nu1hwyirUgIc/AQAA//9KyDOSSAMAAA==`,
}

// Load returns content of embedded template where name is the the path passed to emdbr.
func Load(name string) (string, error) {
	decoded, ok := files[name]
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
