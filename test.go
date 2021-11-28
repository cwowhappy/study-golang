package main

import (
	"fmt"
)

type Data interface {
	Value() [2]byte
}

type data struct {
	value [2]byte
}

func (d *data) Value() [2]byte {
	return d.value
}

func main()  {
	var num uint16
	nums := make([]int, 10)
	for i := 0; i < 10 ; i++ {
		nums[i] = i
	}
	num = num - 1
	fmt.Println(nums[num % 10])
	num = num + 1
	fmt.Println(nums[num % 10])
}
