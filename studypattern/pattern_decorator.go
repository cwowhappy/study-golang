package studypattern

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type SumFunc func(int64, int64) int64

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func getTime() time.Time {
	fmt.Println("getTime")
	return time.Now()
}

func timedSumFunc(f SumFunc) SumFunc {
	return func(start, end int64) int64 {
		// annotation: 先执行getTime()，在return后再执行defer定义的func
		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed (%s): %v ---\n",
				getFunctionName(f), time.Since(t))
		}(getTime())
		fmt.Println("timedSumFunc")
		return f(start, end)
	}
}

func Sum1(start, end int64) int64 {
	var sum int64
	sum = 0
	if start > end {
		start, end = end, start
	}
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}

func Sum2(start, end int64) int64 {
	if start > end {
		start, end = end, start
	}
	return (end - start + 1) * (end + start) / 2
}


func UpdateReturnValueInDefer() int {
	var value = 0

	defer func() {
		value += 1
		fmt.Printf("value = %d\n", value)
	}()

	return value
}