package main

import "fmt"

func main() {
	m := make(map[int]struct{}, 3)
	m[1] = struct{}{}
	m[2] = struct{}{}
	m[3] = struct{}{}
	m[4] = struct{}{}
	m[5] = struct{}{}
	fmt.Println(len(m))

	l := make([]int, 0, 3)
	l = append(l, 1)
	l = append(l, 2)
	l = append(l, 3)
	l = append(l, 4)
	l = append(l, 5)
	fmt.Println(len(l))
	fmt.Println(l)
}
