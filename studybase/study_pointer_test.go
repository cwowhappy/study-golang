package studybase

import "testing"

func TestBytes2String(t *testing.T) {
	bytes := []byte{104, 101, 108, 108, 111}
	str := Bytes2String(bytes)
	if "hello" != str {
		t.Error()
	}
}