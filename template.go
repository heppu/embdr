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
H4sIAAAAAAAC/4ySQWvcMBCFz55fMTGB2uDal9JDwh7a7B4KZVOavYVAtNbYK9aWzEjONjX670WSm0Oh
kIOxeLx5+jQzTYN3RhL2pImFI4kX5U7YmzflFrf3uL8/4G777VADTKI9i55wWbD+sZ69B1DjZNhhAVl+
fHVkc8jy1owTk7VN/1tNQSDdGql03xyFpc+fosRsOLqVaZSZnRpyKAFeBOOO+UDjNAhHW0N2b9zul7IO
N5iq6j1ditytFpSGrP7gkILJ5mtIpwayuMFRTI/WsdL9U/otsCwfkYXuCa/P9Frh9YsYZsKbDdZ/77Xe
Q5YvS3Sg9/kNPsOyJKf3z1UMIS29Bw/QNPjdCIlMbmZtsTXakXZoOqTxSFKSjEB4ORETajESKovuRPGb
hDvhJKwlic6EEsk1dLNuY2wR/Ym+xCIdqtSMEhfIJLVGkqzQnMMr4tMfQ9ETZKrDK3MOrizRYZ5X/2sx
ZB4g45gdksIAQ7d/kpDERRpfELbxxjflwcndOuMK4x4E09e564gfIm6xMpZlGZnCBVcb1Gr4F42YE0bc
mvCqFSZtSR1YvgxDwe8OWl03G+T6bjCWivL2PXWrlvpdrDhlFWrAw58AAAD//4P8eLVEAwAA`,
}

// Load returns content of embedded file where name is the the path passed to embdr.
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
