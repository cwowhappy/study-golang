package studybase

import (
	"encoding/json"
	"testing"
	"time"
)

func TestJsonViewDate_MarshalJSON(t *testing.T) {
	type A struct {
		Value JsonViewDate
	}
	a := A{Value: JsonViewDate(time.Now())}
	result, err := json.Marshal(a)
	if err == nil {
		t.Logf("result:%s", string(result))
	} else {
		t.Errorf(err.Error())
	}
}
func TestJsonViewDate_UnmarshalJSON(t *testing.T) {
	type A struct {
		Value JsonViewDate
	}
	jsonBytes := []byte(`{"Value":"2021-05-27 +0800"}`)
	var a A
	err := json.Unmarshal(jsonBytes, &a)
	if err == nil {
		t.Logf("result:%v", time.Time(a.Value))
	} else {
		t.Errorf(err.Error())
	}

}
