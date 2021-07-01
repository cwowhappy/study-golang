package studybase

import (
	"fmt"
	"time"
)

const (
	JsonViewDateFormatLayout = "\"2006-01-02 -0700\""
	JsonViewDateTimeFormatLayout = "\"2006-01-02 15:04:05 -0700\""
)
type JsonViewDate time.Time
func (value JsonViewDate) MarshalJSON() ([]byte, error) {
	if time.Time(value).IsZero() {
		return []byte("\"\""), nil
	} else {
		// TODO 想想怎么把layout配置化
		return []byte(fmt.Sprintf("%s", time.Time(value).Format(JsonViewDateFormatLayout))), nil
	}
}
func (value *JsonViewDate) UnmarshalJSON(bytes []byte) (err error) {
	if len(bytes) == 0 {
		*value = JsonViewDate{}
	} else {
		v, err := time.Parse(JsonViewDateFormatLayout, string(bytes))
		if err == nil {
			*value = JsonViewDate(v)
		} else {
			*value = JsonViewDate{}
		}
	}
	return
}

type JsonViewDateTime time.Time
func (value JsonViewDateTime) MarshalJSON() ([]byte, error) {
	if time.Time(value).IsZero() {
		return []byte("\"\""), nil
	} else {
		// TODO 想想怎么把layout配置化
		return []byte(fmt.Sprintf("%s", time.Time(value).Format(JsonViewDateTimeFormatLayout))), nil
	}
}
func (value *JsonViewDateTime) UnmarshalJSON(bytes []byte) (err error) {
	if len(bytes) == 0 {
		*value = JsonViewDateTime{}
	} else {
		v, err := time.Parse(JsonViewDateTimeFormatLayout, string(bytes))
		if err == nil {
			*value = JsonViewDateTime(v)
		} else {
			*value = JsonViewDateTime{}
		}
	}
	return
}
