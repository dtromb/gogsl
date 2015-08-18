//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package combination

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_combination.h>
   #include <unistd.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "os"

type GslCombination struct {
   gogsl.GslReference
}

func CombinationAlloc(n int, k int) *GslCombination {
   _ref := C.gsl_combination_alloc(C.size_t(n), C.size_t(k))
   _result := &GslCombination{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func CombinationCalloc(n int, k int) *GslCombination {
   _ref := C.gsl_combination_calloc(C.size_t(n), C.size_t(k))
   _result := &GslCombination{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func InitFirst(c *GslCombination) {
   C.gsl_combination_init_first((*C.gsl_combination)(unsafe.Pointer(c.Ptr())))
}

func InitLast(c *GslCombination) {
   C.gsl_combination_init_last((*C.gsl_combination)(unsafe.Pointer(c.Ptr())))
}

func (x *GslCombination) Dispose() {
   C.gsl_combination_free((*C.gsl_combination)(unsafe.Pointer(x.Ptr())))
}

func Memcpy(dest *GslCombination, src *GslCombination) int32 {
   return int32(C.gsl_combination_memcpy((*C.gsl_combination)(unsafe.Pointer(dest.Ptr())), (*C.gsl_combination)(unsafe.Pointer(src.Ptr()))))
}

func Get(c *GslCombination, i int) int {
   return int(C.gsl_combination_get((*C.gsl_combination)(unsafe.Pointer(c.Ptr())), C.size_t(i)))
}

func N(c *GslCombination) int {
   return int(C.gsl_combination_n((*C.gsl_combination)(unsafe.Pointer(c.Ptr()))))
}

func K(c *GslCombination) int {
   return int(C.gsl_combination_k((*C.gsl_combination)(unsafe.Pointer(c.Ptr()))))
}

func Valid(c *GslCombination) int32 {
   return int32(C.gsl_combination_valid((*C.gsl_combination)(unsafe.Pointer(c.Ptr()))))
}

func Next(c *GslCombination) int32 {
   return int32(C.gsl_combination_next((*C.gsl_combination)(unsafe.Pointer(c.Ptr()))))
}

func Prev(c *GslCombination) int32 {
   return int32(C.gsl_combination_prev((*C.gsl_combination)(unsafe.Pointer(c.Ptr()))))
}

func Fwrite(stream *os.File, c *GslCombination) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_combination_fwrite(_file_0, (*C.gsl_combination)(unsafe.Pointer(c.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func Fread(stream *os.File, c *GslCombination) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_combination_fread(_file_0, (*C.gsl_combination)(unsafe.Pointer(c.Ptr()))))
   C.fclose(_file_0)
   return _result
}

func Fprintf(stream *os.File, c *GslCombination, format string) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _string_2 := C.CString(format)
   _result := int32(C.gsl_combination_fprintf(_file_0, (*C.gsl_combination)(unsafe.Pointer(c.Ptr())), _string_2))
   C.fclose(_file_0)
   C.free(unsafe.Pointer(_string_2))
   return _result
}

func Fscanf(stream *os.File, c *GslCombination) int32 {
   _file_0 := C.fdopen(C.dup(C.int(stream.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))
   _result := int32(C.gsl_combination_fscanf(_file_0, (*C.gsl_combination)(unsafe.Pointer(c.Ptr()))))
   C.fclose(_file_0)
   return _result
}

