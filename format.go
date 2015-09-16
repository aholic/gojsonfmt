package gojsonfmt

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func JsonFmtHandleFunc(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.Write(body)
		w.Write([]byte{'\n'})
		w.Write([]byte(err.Error()))
		return
	}

	formatted, err := Format(body)
	if err != nil {
		w.Write(body)
		w.Write([]byte{'\n'})
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(formatted)
}

func Format(content []byte) ([]byte, error) {
	dst := new(bytes.Buffer)
	if err := json.Indent(dst, content, "", "  "); err != nil {
		return nil, err
	}

	return dst.Bytes(), nil
}
