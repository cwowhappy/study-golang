package studybase

import (
	"fmt"
	"reflect"
	"unsafe"
)

func processSlice(s []int) {
	fmt.Printf("s: %p %#v\n", s, (*reflect.SliceHeader)(unsafe.Pointer(&s)))
	s = append(s, 4)
	fmt.Printf("s: %p %#v\n", s, (*reflect.SliceHeader)(unsafe.Pointer(&s)))
}

func SliceSample01() {
	s := make([]int, 0, 1)
	fmt.Printf("func[1] -> s: %p %#v\n", s, (*reflect.SliceHeader)(unsafe.Pointer(&s)))
	s = append(s, 4)
	fmt.Printf("func[2] -> s: %p %#v\n", s, (*reflect.SliceHeader)(unsafe.Pointer(&s)))
	processSlice(s)
}
