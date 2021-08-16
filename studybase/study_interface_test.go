package studybase

import "testing"

func TestAnimal_Say(t *testing.T) {
	operator := GetOperator()
	operator.Say()
}
