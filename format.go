package gojsonfmt

import (
	"bytes"
	"encoding/json"
)

func Format(content []byte) ([]byte, error) {
	dst := new(bytes.Buffer)
	if err := json.Indent(dst, content, "", "  "); err != nil {
		return nil, err
	}

	return dst.Bytes(), nil
}
