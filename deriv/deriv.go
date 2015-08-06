//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package deriv

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_deriv.h>
*/
import "C"

import "github.com/dtromb/gogsl"
import "unsafe"

func Central(f *gogsl.GslFunction, x float64, h float64) (int32, float64, float64) {
   var _outptr_3 C.double
   var _outptr_4 C.double
   gogsl.InitializeGslFunction(f)
   _result := int32(C.gsl_deriv_central((*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(x), C.double(h), &_outptr_3, &_outptr_4))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_3)), *(*float64)(unsafe.Pointer(&_outptr_4))
}

func Forward(f *gogsl.GslFunction, x float64, h float64) (int32, float64, float64) {
   var _outptr_3 C.double
   var _outptr_4 C.double
   gogsl.InitializeGslFunction(f)
   _result := int32(C.gsl_deriv_forward((*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(x), C.double(h), &_outptr_3, &_outptr_4))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_3)), *(*float64)(unsafe.Pointer(&_outptr_4))
}

func Backward(f *gogsl.GslFunction, x float64, h float64) (int32, float64, float64) {
   var _outptr_3 C.double
   var _outptr_4 C.double
   gogsl.InitializeGslFunction(f)
   _result := int32(C.gsl_deriv_backward((*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(x), C.double(h), &_outptr_3, &_outptr_4))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_3)), *(*float64)(unsafe.Pointer(&_outptr_4))
}

