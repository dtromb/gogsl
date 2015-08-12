//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package airy

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_airy.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func GslSfAiryAi(x float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_airy_Ai(C.double(x), C.gsl_mode_t(mode)))
}

func GslSfAiryAiE(x float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_airy_Ai_e(C.double(x), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GslSfAiryBi(x float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_airy_Bi(C.double(x), C.gsl_mode_t(mode)))
}

func GslSfAiryBiE(x float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_airy_Bi_e(C.double(x), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GslSfAiryAiScaled(x float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_airy_Ai_scaled(C.double(x), C.gsl_mode_t(mode)))
}

func GslSfAiryAiScaledE(x float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_airy_Ai_scaled_e(C.double(x), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GslSfAiryBiScaled(x float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_airy_Bi_scaled(C.double(x), C.gsl_mode_t(mode)))
}

func GslSfAiryBiScaledE(x float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_airy_Bi_scaled_e(C.double(x), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GslSfAiryAiDeriv(x float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_airy_Ai_deriv(C.double(x), C.gsl_mode_t(mode)))
}

func GslSfAiryAiDerivE(x float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_airy_Ai_deriv_e(C.double(x), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GslSfAiryBiDeriv(x float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_airy_Bi_deriv(C.double(x), C.gsl_mode_t(mode)))
}

func GslSfAiryBiDerivE(x float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_airy_Bi_deriv_e(C.double(x), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GslSfAiryAiDerivScaled(x float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_airy_Ai_deriv_scaled(C.double(x), C.gsl_mode_t(mode)))
}

func GslSfAiryAiDerivScaledE(x float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_airy_Ai_deriv_scaled_e(C.double(x), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GslSfAiryBiDerivScaled(x float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_airy_Bi_deriv_scaled(C.double(x), C.gsl_mode_t(mode)))
}

func GslSfAiryBiDerivScaledE(x float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_airy_Bi_deriv_scaled_e(C.double(x), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GslSfAiryZeroAiE(s uint32, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_airy_zero_Ai_e(C.uint(s), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GslSfAiryZeroBi(s uint32) float64 {
	return float64(C.gsl_sf_airy_zero_Bi(C.uint(s)))
}

func GslSfAiryZeroBiE(s uint32, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_airy_zero_Bi_e(C.uint(s), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GslSfAiryZeroAiDeriv(s uint32) float64 {
	return float64(C.gsl_sf_airy_zero_Ai_deriv(C.uint(s)))
}

func GslSfAiryZeroAiDerivE(s uint32, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_airy_zero_Ai_deriv_e(C.uint(s), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GslSfAiryZeroBiDeriv(s uint32) float64 {
	return float64(C.gsl_sf_airy_zero_Bi_deriv(C.uint(s)))
}

func GslSfAiryZeroBiDerivE(s uint32, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_airy_zero_Bi_deriv_e(C.uint(s), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}
