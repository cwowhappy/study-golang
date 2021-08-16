package main

import (
	"fmt"
)

func main()  {
	nums := []int{1, 2, 3}
	fmt.Println(nums)
	swapNum(nums)
	fmt.Println(nums)
}

func swapNum(nums []int) {
	for i, j := 0, len(nums) - 1; i < j; i, j = i + 1, j - 1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func addNum(nums []int, num int) {
	nums = append(nums, num)
}

