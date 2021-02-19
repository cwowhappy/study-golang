package studybase

import (
	"encoding/json"
	"fmt"
	"time"
)

type JSONTime struct {
	time.Time
}

func (j JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(`"2006-01-02 15:04:05"`), nil
}

func (j *JSONTime) UnmarshalJSON(b []byte) (err error) {
	if len(b) == 0 {
		j.Time = time.Time{}
		fmt.Println("a")
	} else {
		j.Time, err = time.Parse("\"2006-01-02 15:04:05\"", string(b))
		fmt.Println("a")
	}
	return
}

type Person struct {
	Name string `json:"name"`
	Birthday JSONTime `json:"birthday"`
}

func (p Person) toString() string {
	fmt.Printf("[Func:toString]Person p point=%p\n", &p)
	return fmt.Sprintf("Person[Name=%s, Age=%v]\n", p.Name, p.Birthday)
}

func BuildPerson() (person Person) {
	jsonContent := `{"name": "cwowhappy", "birthday": "2021-01-01 00:00:00"}`
	err := json.Unmarshal([]byte(jsonContent), &person)
	if err != nil {
		fmt.Println(err.Error())
	}
	return person
}
