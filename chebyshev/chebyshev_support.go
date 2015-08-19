package chebyshev

/*
	#include <unistd.h>
	#include <gsl/gsl_chebyshev.h>
*/
import "C"

import (
	"reflect"
	"unsafe"
)

// We copy these out at the expense of a little performance to
// keep memory management safety.

//double * gsl_cheb_coeffs (const gsl_cheb_series * cs)
func Coeffs(cs *GslChebSeries) []float64 {
	ptr := unsafe.Pointer(C.gsl_cheb_coeffs((*C.gsl_cheb_series)(unsafe.Pointer(cs.Ptr()))))
	order := Order(cs)
	hdr := &reflect.SliceHeader{Len: order, Cap: order, Data: uintptr(ptr)}
	cf := make([]float64, order)
	copy(cf, *(*[]float64)(unsafe.Pointer(hdr)))
	return cf
}
