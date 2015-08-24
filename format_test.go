package gojsonfmt

import "testing"

func TestFormat(t *testing.T) {
	content := [][]byte{
		[]byte(`{"a":"b", "b":"c"}`),
		[]byte(`{"a":"b", "b":}`),
	}

	expect := []byte("{\n  \"a\": \"b\",\n  \"b\": \"c\"\n}")

	if out, err := Format(content[0]); err != nil {
		t.Errorf("expect: %v, actual: %v", nil, err)
	} else {
		if string(out) != string(expect) {
			t.Errorf("expect: %v, actual: %v", string(expect), string(out))
		}
	}

	if out, err := Format(content[1]); err == nil {
		t.Errorf("expect: not nil, actual: %v", err)
	} else {
		if out != nil {
			t.Errorf("expect: %v, actual: %v", nil, out)
		}
	}
}
