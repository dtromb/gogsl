//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package gogsl

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_math.h>
*/
import "C"

import "unsafe"

func GslIsnan(x float64) int32 {
	return int32(C.gsl_isnan(C.double(x)))
}

func GslIsinf(x float64) int32 {
	return int32(C.gsl_isinf(C.double(x)))
}

func GslFinite(x float64) int32 {
	return int32(C.gsl_finite(C.double(x)))
}

func GslLog1p(x float64) float64 {
	return float64(C.gsl_log1p(C.double(x)))
}

func GslExpm1(x float64) float64 {
	return float64(C.gsl_expm1(C.double(x)))
}

func GslHypot(x float64, y float64) float64 {
	return float64(C.gsl_hypot(C.double(x), C.double(y)))
}

func GslHypot3(x float64, y float64, z float64) float64 {
	return float64(C.gsl_hypot3(C.double(x), C.double(y), C.double(z)))
}

func GslAcosh(x float64) float64 {
	return float64(C.gsl_acosh(C.double(x)))
}

func GslAsinh(x float64) float64 {
	return float64(C.gsl_asinh(C.double(x)))
}

func GslAtanh(x float64) float64 {
	return float64(C.gsl_atanh(C.double(x)))
}

func GslLdexp(x float64, e int32) float64 {
	return float64(C.gsl_ldexp(C.double(x), C.int(e)))
}

func GslFrexp(x float64) (float64, int32) {
	var _outptr_1 C.int
	_result := float64(C.gsl_frexp(C.double(x), &_outptr_1))
	return _result, *(*int32)(unsafe.Pointer(&_outptr_1))
}

func GslPowInt(x float64, n int32) float64 {
	return float64(C.gsl_pow_int(C.double(x), C.int(n)))
}

func GslPowUint(x float64, n uint32) float64 {
	return float64(C.gsl_pow_uint(C.double(x), C.uint(n)))
}

func GslPow2(x float64) float64 {
	return float64(C.gsl_pow_2(C.double(x)))
}

func GslPow3(x float64) float64 {
	return float64(C.gsl_pow_3(C.double(x)))
}

func GslPow4(x float64) float64 {
	return float64(C.gsl_pow_4(C.double(x)))
}

func GslPow5(x float64) float64 {
	return float64(C.gsl_pow_5(C.double(x)))
}

func GslPow6(x float64) float64 {
	return float64(C.gsl_pow_6(C.double(x)))
}

func GslPow7(x float64) float64 {
	return float64(C.gsl_pow_7(C.double(x)))
}

func GslPow8(x float64) float64 {
	return float64(C.gsl_pow_8(C.double(x)))
}

func GslPow9(x float64) float64 {
	return float64(C.gsl_pow_9(C.double(x)))
}

func GslFcmp(x float64, y float64, epsilon float64) int32 {
	return int32(C.gsl_fcmp(C.double(x), C.double(y), C.double(epsilon)))
}
