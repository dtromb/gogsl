//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package qrng

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_qrng.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "reflect"

type GslQrng struct {
	gogsl.GslReference
}

type GslQrngType struct {
	gogsl.GslReference
	cPtr uintptr
}

func QrngAlloc(t *GslQrngType, d uint32) *GslQrng {
	_ref := C.gsl_qrng_alloc((*C.gsl_qrng_type)(unsafe.Pointer(t.Ptr())), C.uint(d))
	_result := &GslQrng{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslQrng) Dispose() {
	C.gsl_qrng_free((*C.gsl_qrng)(unsafe.Pointer(x.Ptr())))
}

func Init(q *GslQrng) {
	C.gsl_qrng_init((*C.gsl_qrng)(unsafe.Pointer(q.Ptr())))
}

func Get(q *GslQrng, x []float64) int32 {
	_slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&x))
	return int32(C.gsl_qrng_get((*C.gsl_qrng)(unsafe.Pointer(q.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data))))
}

func Name(q *GslQrng) string {
	return C.GoString(C.gsl_qrng_name((*C.gsl_qrng)(unsafe.Pointer(q.Ptr()))))
}

func Size(q *GslQrng) int {
	return int(C.gsl_qrng_size((*C.gsl_qrng)(unsafe.Pointer(q.Ptr()))))
}

func State(q *GslQrng) *QrngState {
	return (*QrngState)(C.gsl_qrng_state((*C.gsl_qrng)(unsafe.Pointer(q.Ptr()))))
}

func Memcpy(dest *GslQrng, src *GslQrng) int32 {
	return int32(C.gsl_qrng_memcpy((*C.gsl_qrng)(unsafe.Pointer(dest.Ptr())), (*C.gsl_qrng)(unsafe.Pointer(src.Ptr()))))
}

func Clone(q *GslQrng) *GslQrng {
	_ref := C.gsl_qrng_clone((*C.gsl_qrng)(unsafe.Pointer(q.Ptr())))
	_result := &GslQrng{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}
