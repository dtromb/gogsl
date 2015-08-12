//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package elljac

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_elljac.h>
*/
import "C"

import "reflect"
import "unsafe"

func E(u float64, m float64, sn []float64, cn []float64, dn []float64) int32 {
	_slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&sn))
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&cn))
	_slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&dn))
	return int32(C.gsl_sf_elljac_e(C.double(u), C.double(m), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), (*C.double)(unsafe.Pointer(_slice_header_3.Data)), (*C.double)(unsafe.Pointer(_slice_header_4.Data))))
}
