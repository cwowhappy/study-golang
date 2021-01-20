package studypattern

import (
	"fmt"
	"testing"
)

func TestPatternDecorator(t *testing.T) {
	sum1 := timedSumFunc(Sum1)
	sum2 := timedSumFunc(Sum2)

	fmt.Printf("%d, %d\n", sum1(-10000, 10000000), sum2(-10000, 10000000))
}

func TestUpdateReturnValueInDefer(t *testing.T) {
	result := 0
	r := UpdateReturnValueInDefer()
	if result != r {
		t.Error("error")
	}
}
