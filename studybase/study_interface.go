package studybase

import (
	"fmt"
	"time"
)

type Operator interface {
	Say()
}

type Animal struct {
	Name string
}

func (animal Animal) Say() {
	fmt.Printf("Hi, my name is %s[%v] \n", animal.Name, time.Now())
}

func GetOperator() Operator {
	return Animal{Name: "kitty"}
}