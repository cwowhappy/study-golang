package studybase

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p *Person) toString() string {
	fmt.Printf("[Func:toString]Person p point=%p\n", p)
	return fmt.Sprintf("Person[Name=%s, Age=%d]\n", p.Name, p.Age)
}

func ProcessPerson() {
	p := new(Person)
	p.Name = "cwowhappy"
	p.Age = 10
	fmt.Printf("[Func:ProcessPerson]Person p point=%p\n", p)
	fmt.Println(p.toString())
}
