//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package multiset

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_multiset.h>
   #include <unistd.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "os"

type GslMultiset struct {
   gogsl.GslReference
}

func MultisetAlloc(n int, k int) *GslMultiset {
   _ref := C.gsl_multiset_alloc(C.size_t(n), C.size_t(k))
   _result := &GslMultiset{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func MultisetCalloc(n int, k int) *GslMultiset {
   _ref := C.gsl_multiset_calloc(C.size_t(n), C.size_t(k))
   _result := &GslMultiset{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func InitFirst(c *GslMultiset) {
   C.gsl_multiset_init_first((*C.gsl_multiset)(unsafe.Pointer(c.Ptr())))
}

func InitLast(c *GslMultiset) {
   C.gsl_multiset_init_last((*C.gsl_multiset)(unsafe.Pointer(c.Ptr())))
}

func (x *GslMultiset) Dispose() {
   C.gsl_multiset_free((*C.gsl_multiset)(unsafe.Pointer(x.Ptr())))
}

func Memcpy(dest *GslMultiset, src *GslMultiset) int32 {
   return int32(C.gsl_multiset_memcpy((*C.gsl_multiset)(unsafe.Pointer(dest.Ptr())), (*C.gsl_multiset)(unsafe.Pointer(src.Ptr()))))
}

func Get(c *GslMultiset, i int) int {
   return int(C.gsl_multiset_get((*C.gsl_multiset)(unsafe.Pointer(c.Ptr())), C.size_t(i)))
}

func N(c *GslMultiset) int {
   return int(C.gsl_multiset_n((*C.gsl_multiset)(unsafe.Pointer(c.Ptr()))))
}

func K(c *GslMultiset) int {
   return int(C.gsl_multiset_k((*C.gsl_multiset)(unsafe.Pointer(c.Ptr()))))
}

func Valid(c *GslMultiset) int32 {
   return int32(C.gsl_multiset_valid((*C.gsl_multiset)(unsafe.Pointer(c.Ptr()))))
}

func Next(c *GslMultiset) int32 {
   return int32(C.gsl_multiset_next((*C.gsl_multiset)(unsafe.Pointer(c.Ptr()))))
}

func Prev(c *GslMultiset) int32 {
   return int32(C.gsl_multiset_prev((*C.gsl_multiset)(unsafe.Pointer(c.Ptr()))))
}

func Fwrite(stream *os.File, c *GslMultiset) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_multiset_fwrite(_file_0, (*C.gsl_multiset)(unsafe.Pointer(c.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func Fread(stream *os.File, c *GslMultiset) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_multiset_fread(_file_0, (*C.gsl_multiset)(unsafe.Pointer(c.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func Fprintf(stream *os.File, c *GslMultiset, format string) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _string_2 := C.CString(format)
   _result := int32(C.gsl_multiset_fprintf(_file_0, (*C.gsl_multiset)(unsafe.Pointer(c.Ptr())), _string_2))
   C.fclose(_file_0)
   C.free(unsafe.Pointer(_string_2))
   return _result
}

func Fscanf(stream *os.File, c *GslMultiset) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_multiset_fscanf(_file_0, (*C.gsl_multiset)(unsafe.Pointer(c.Ptr()))))
   C.fclose(_file_0)
   return _result
}

