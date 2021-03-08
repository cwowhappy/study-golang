package studypattern

import "testing"

func TestPatternVisitor(t *testing.T)  {
	c := Circle{Radius: 10}
	r := Rectangle{Height: 100, Width: 200}
	shapes := []Shape{c, r}
	for _, s := range shapes {
		s.accept(JsonVisitor)
		s.accept(XmlVisitor)
	}

}
