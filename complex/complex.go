//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package complex

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_complex.h>
   #include <gsl/gsl_complex_math.h>
*/
import "C"

import "unsafe"

func GslComplexRect(x float64, y float64) complex128 {
	_result := C.gsl_complex_rect(C.double(x), C.double(y))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexPolar(r float64, theta float64) complex128 {
	_result := C.gsl_complex_polar(C.double(r), C.double(theta))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArg(z complex128) float64 {
	_arg_0 := GoComplexToGsl(z)
	return float64(C.gsl_complex_arg(*(*C.gsl_complex)(unsafe.Pointer(_arg_0))))
}

func GslComplexAbs(z complex128) float64 {
	_arg_0 := GoComplexToGsl(z)
	return float64(C.gsl_complex_abs(*(*C.gsl_complex)(unsafe.Pointer(_arg_0))))
}

func GslComplexAbs2(z complex128) float64 {
	_arg_0 := GoComplexToGsl(z)
	return float64(C.gsl_complex_abs2(*(*C.gsl_complex)(unsafe.Pointer(_arg_0))))
}

func GslComplexLogabs(z complex128) float64 {
	_arg_0 := GoComplexToGsl(z)
	return float64(C.gsl_complex_logabs(*(*C.gsl_complex)(unsafe.Pointer(_arg_0))))
}

func GslComplexAdd(a complex128, b complex128) complex128 {
	_arg_0 := GoComplexToGsl(a)
	_arg_1 := GoComplexToGsl(b)
	_result := C.gsl_complex_add(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), *(*C.gsl_complex)(unsafe.Pointer(_arg_1)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexSub(a complex128, b complex128) complex128 {
	_arg_0 := GoComplexToGsl(a)
	_arg_1 := GoComplexToGsl(b)
	_result := C.gsl_complex_sub(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), *(*C.gsl_complex)(unsafe.Pointer(_arg_1)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexMul(a complex128, b complex128) complex128 {
	_arg_0 := GoComplexToGsl(a)
	_arg_1 := GoComplexToGsl(b)
	_result := C.gsl_complex_mul(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), *(*C.gsl_complex)(unsafe.Pointer(_arg_1)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexDiv(a complex128, b complex128) complex128 {
	_arg_0 := GoComplexToGsl(a)
	_arg_1 := GoComplexToGsl(b)
	_result := C.gsl_complex_div(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), *(*C.gsl_complex)(unsafe.Pointer(_arg_1)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexAddReal(a complex128, x float64) complex128 {
	_arg_0 := GoComplexToGsl(a)
	_result := C.gsl_complex_add_real(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), C.double(x))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexSubReal(a complex128, x float64) complex128 {
	_arg_0 := GoComplexToGsl(a)
	_result := C.gsl_complex_sub_real(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), C.double(x))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexMulReal(a complex128, x float64) complex128 {
	_arg_0 := GoComplexToGsl(a)
	_result := C.gsl_complex_mul_real(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), C.double(x))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexDivReal(a complex128, x float64) complex128 {
	_arg_0 := GoComplexToGsl(a)
	_result := C.gsl_complex_div_real(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), C.double(x))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexAddImag(a complex128, y float64) complex128 {
	_arg_0 := GoComplexToGsl(a)
	_result := C.gsl_complex_add_imag(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), C.double(y))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexSubImag(a complex128, y float64) complex128 {
	_arg_0 := GoComplexToGsl(a)
	_result := C.gsl_complex_sub_imag(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), C.double(y))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexMulImag(a complex128, y float64) complex128 {
	_arg_0 := GoComplexToGsl(a)
	_result := C.gsl_complex_mul_imag(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), C.double(y))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexDivImag(a complex128, y float64) complex128 {
	_arg_0 := GoComplexToGsl(a)
	_result := C.gsl_complex_div_imag(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), C.double(y))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexConjugate(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_conjugate(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexInverse(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_inverse(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexNegative(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_negative(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexSqrt(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_sqrt(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexSqrtReal(x float64) complex128 {
	_result := C.gsl_complex_sqrt_real(C.double(x))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexPow(z complex128, a complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_arg_1 := GoComplexToGsl(a)
	_result := C.gsl_complex_pow(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), *(*C.gsl_complex)(unsafe.Pointer(_arg_1)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexPowReal(z complex128, x float64) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_pow_real(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), C.double(x))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexExp(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_exp(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexLog(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_log(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexLog10(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_log10(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexLogB(z complex128, b complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_arg_1 := GoComplexToGsl(b)
	_result := C.gsl_complex_log_b(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)), *(*C.gsl_complex)(unsafe.Pointer(_arg_1)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexSin(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_sin(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexCos(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_cos(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexTan(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_tan(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexSec(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_sec(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexCsc(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_csc(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexCot(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_cot(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArcsin(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_arcsin(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArcsinReal(z float64) complex128 {
	_result := C.gsl_complex_arcsin_real(C.double(z))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArccos(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_arccos(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArccosReal(z float64) complex128 {
	_result := C.gsl_complex_arccos_real(C.double(z))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArctan(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_arctan(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArcsec(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_arcsec(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArcsecReal(z float64) complex128 {
	_result := C.gsl_complex_arcsec_real(C.double(z))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArccsc(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_arccsc(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArccscReal(z float64) complex128 {
	_result := C.gsl_complex_arccsc_real(C.double(z))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArccot(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_arccot(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexSinh(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_sinh(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexCosh(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_cosh(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexTanh(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_tanh(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexSech(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_sech(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexCsch(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_csch(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexCoth(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_coth(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArcsinh(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_arcsinh(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArccosh(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_arccosh(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArccoshReal(z float64) complex128 {
	_result := C.gsl_complex_arccosh_real(C.double(z))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArctanh(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_arctanh(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArctanhReal(z float64) complex128 {
	_result := C.gsl_complex_arctanh_real(C.double(z))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArcsech(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_arcsech(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArccsch(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_arccsch(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}

func GslComplexArccoth(z complex128) complex128 {
	_arg_0 := GoComplexToGsl(z)
	_result := C.gsl_complex_arccoth(*(*C.gsl_complex)(unsafe.Pointer(_arg_0)))
	return GslComplexToGo(uintptr(unsafe.Pointer(&_result)))
}
