package gogsl

/*

*/
import "C"

import (
	"reflect"
	"unsafe"
)

type FileMode string

var READ_ONLY FileMode = "r"
var READ_WRITE FileMode = "r+"
var TRUNC_WRITE_ONLY FileMode = "w"
var TRUNC_READ_WRITE FileMode = "w+"
var APPEND_ONLY FileMode = "a"
var READ_APPEND FileMode = "a+"

func (fm *FileMode) Ptr() uintptr {
	hdr := (*reflect.StringHeader)(unsafe.Pointer(fm))
	return uintptr(unsafe.Pointer(hdr.Data))
}
