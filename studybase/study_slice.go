package studybase

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func processSlice(s []int) {
	fmt.Printf("s: %p %#v\n", s, (*reflect.SliceHeader)(unsafe.Pointer(&s)))
	s = append(s, 4)
	fmt.Printf("s: %p %#v\n", s, (*reflect.SliceHeader)(unsafe.Pointer(&s)))
}

func SliceSample01() {
	s := make([]string, 0, 10)
	s = append(s, "1", "2", "3", "4")
	fmt.Printf("%s/n", strings.Join(s, ","))
	s = s[:0]
	s = append(s, "11")
	fmt.Printf("%s/n", strings.Join(s, ","))
}
