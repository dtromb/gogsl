//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package dht

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_dht.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "reflect"

type GslDht struct {
	gogsl.GslReference
}

func Alloc(size int) *GslDht {
	_ref := C.gsl_dht_alloc(C.size_t(size))
	_result := &GslDht{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func Init(t *GslDht, nu float64, xmax float64) int32 {
	return int32(C.gsl_dht_init((*C.gsl_dht)(unsafe.Pointer(t.Ptr())), C.double(nu), C.double(xmax)))
}

func New(size int, nu float64, xmax float64) *GslDht {
	_ref := C.gsl_dht_new(C.size_t(size), C.double(nu), C.double(xmax))
	_result := &GslDht{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslDht) Dispose() {
	C.gsl_dht_free((*C.gsl_dht)(unsafe.Pointer(x.Ptr())))
}

func Apply(t *GslDht, fIn []float64, fOut []float64) int32 {
	_slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&fIn))
	_slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&fOut))
	return int32(C.gsl_dht_apply((*C.gsl_dht)(unsafe.Pointer(t.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data))))
}

func XSample(t *GslDht, n int32) float64 {
	return float64(C.gsl_dht_x_sample((*C.gsl_dht)(unsafe.Pointer(t.Ptr())), C.int(n)))
}

func KSample(t *GslDht, n int32) float64 {
	return float64(C.gsl_dht_k_sample((*C.gsl_dht)(unsafe.Pointer(t.Ptr())), C.int(n)))
}
