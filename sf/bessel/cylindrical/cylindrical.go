//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package cylindrical

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_bessel.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"
import "reflect"

func J0(x float64) float64 {
	return float64(C.gsl_sf_bessel_J0(C.double(x)))
}

func J0E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_J0_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func J1(x float64) float64 {
	return float64(C.gsl_sf_bessel_J1(C.double(x)))
}

func J1E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_J1_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Jn(n int32, x float64) float64 {
	return float64(C.gsl_sf_bessel_Jn(C.int(n), C.double(x)))
}

func JnE(n int32, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_Jn_e(C.int(n), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func JnArray(nmin int32, nmax int32, x float64, resultArray []float64) int32 {
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
	return int32(C.gsl_sf_bessel_Jn_array(C.int(nmin), C.int(nmax), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_3.Data))))
}

func Y0(x float64) float64 {
	return float64(C.gsl_sf_bessel_Y0(C.double(x)))
}

func Y0E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_Y0_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Y1(x float64) float64 {
	return float64(C.gsl_sf_bessel_Y1(C.double(x)))
}

func Y1E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_Y1_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Yn(n int32, x float64) float64 {
	return float64(C.gsl_sf_bessel_Yn(C.int(n), C.double(x)))
}

func YnE(n int32, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_Yn_e(C.int(n), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func YnArray(nmin int32, nmax int32, x float64, resultArray []float64) int32 {
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
	return int32(C.gsl_sf_bessel_Yn_array(C.int(nmin), C.int(nmax), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_3.Data))))
}

func I0(x float64) float64 {
	return float64(C.gsl_sf_bessel_I0(C.double(x)))
}

func I0E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_I0_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func I1(x float64) float64 {
	return float64(C.gsl_sf_bessel_I1(C.double(x)))
}

func I1E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_I1_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func In(n int32, x float64) float64 {
	return float64(C.gsl_sf_bessel_In(C.int(n), C.double(x)))
}

func InE(n int32, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_In_e(C.int(n), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func InArray(nmin int32, nmax int32, x float64, resultArray []float64) int32 {
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
	return int32(C.gsl_sf_bessel_In_array(C.int(nmin), C.int(nmax), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_3.Data))))
}

func I0Scaled(x float64) float64 {
	return float64(C.gsl_sf_bessel_I0_scaled(C.double(x)))
}

func I0ScaledE(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_I0_scaled_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func I1Scaled(x float64) float64 {
	return float64(C.gsl_sf_bessel_I1_scaled(C.double(x)))
}

func I1ScaledE(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_I1_scaled_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func InScaled(n int32, x float64) float64 {
	return float64(C.gsl_sf_bessel_In_scaled(C.int(n), C.double(x)))
}

func InScaledE(n int32, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_In_scaled_e(C.int(n), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func InScaledArray(nmin int32, nmax int32, x float64, resultArray []float64) int32 {
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
	return int32(C.gsl_sf_bessel_In_scaled_array(C.int(nmin), C.int(nmax), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_3.Data))))
}

func K0(x float64) float64 {
	return float64(C.gsl_sf_bessel_K0(C.double(x)))
}

func K0E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_K0_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func K1(x float64) float64 {
	return float64(C.gsl_sf_bessel_K1(C.double(x)))
}

func K1E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_K1_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Kn(n int32, x float64) float64 {
	return float64(C.gsl_sf_bessel_Kn(C.int(n), C.double(x)))
}

func KnE(n int32, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_Kn_e(C.int(n), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func KnArray(nmin int32, nmax int32, x float64, resultArray []float64) int32 {
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
	return int32(C.gsl_sf_bessel_Kn_array(C.int(nmin), C.int(nmax), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_3.Data))))
}

func K0Scaled(x float64) float64 {
	return float64(C.gsl_sf_bessel_K0_scaled(C.double(x)))
}

func K0ScaledE(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_K0_scaled_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func K1Scaled(x float64) float64 {
	return float64(C.gsl_sf_bessel_K1_scaled(C.double(x)))
}

func K1ScaledE(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_K1_scaled_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func KnScaled(n int32, x float64) float64 {
	return float64(C.gsl_sf_bessel_Kn_scaled(C.int(n), C.double(x)))
}

func KnScaledE(n int32, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_bessel_Kn_scaled_e(C.int(n), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func KnScaledArray(nmin int32, nmax int32, x float64, resultArray []float64) int32 {
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
	return int32(C.gsl_sf_bessel_Kn_scaled_array(C.int(nmin), C.int(nmax), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_3.Data))))
}
