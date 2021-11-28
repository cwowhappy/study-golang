package studycgo

import "C"
import (
	"fmt"
	"unsafe"
)

func func1() {
	goStr1, goStr2 := "Hello", "world"
	cstr1, cstr2 := C.CString(goStr1), C.CString(goStr2)
	defer C.free(unsafe.Pointer(cstr1))
	defer C.free(unsafe.Pointer(cstr2))
	cstr3 := C.cat(cstr1, cstr2)
	goStr3 := C.GoString(cstr3)
	fmt.Println(goStr3)
}
