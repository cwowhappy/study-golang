package studybase

import "testing"

func TestBuildPerson(t *testing.T) {
	person := BuildPerson()
	t.Log(person.toString())
}

