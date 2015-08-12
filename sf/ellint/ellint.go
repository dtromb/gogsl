//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package ellint

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_ellint.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func Kcomp(k float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_ellint_Kcomp(C.double(k), C.gsl_mode_t(mode)))
}

func KcompE(k float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_ellint_Kcomp_e(C.double(k), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Ecomp(k float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_ellint_Ecomp(C.double(k), C.gsl_mode_t(mode)))
}

func EcompE(k float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_ellint_Ecomp_e(C.double(k), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Pcomp(k float64, n float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_ellint_Pcomp(C.double(k), C.double(n), C.gsl_mode_t(mode)))
}

func PcompE(k float64, n float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_ellint_Pcomp_e(C.double(k), C.double(n), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func F(phi float64, k float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_ellint_F(C.double(phi), C.double(k), C.gsl_mode_t(mode)))
}

func FE(phi float64, k float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_ellint_F_e(C.double(phi), C.double(k), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func E(phi float64, k float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_ellint_E(C.double(phi), C.double(k), C.gsl_mode_t(mode)))
}

func EE(phi float64, k float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_ellint_E_e(C.double(phi), C.double(k), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func P(phi float64, k float64, n float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_ellint_P(C.double(phi), C.double(k), C.double(n), C.gsl_mode_t(mode)))
}

func PE(phi float64, k float64, n float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_ellint_P_e(C.double(phi), C.double(k), C.double(n), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func D(phi float64, k float64, n float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_ellint_D(C.double(phi), C.double(k), C.double(n), C.gsl_mode_t(mode)))
}

func DE(phi float64, k float64, n float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_ellint_D_e(C.double(phi), C.double(k), C.double(n), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func RC(x float64, y float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_ellint_RC(C.double(x), C.double(y), C.gsl_mode_t(mode)))
}

func RCE(x float64, y float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_ellint_RC_e(C.double(x), C.double(y), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func RD(x float64, y float64, z float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_ellint_RD(C.double(x), C.double(y), C.double(z), C.gsl_mode_t(mode)))
}

func RDE(x float64, y float64, z float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_ellint_RD_e(C.double(x), C.double(y), C.double(z), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func RF(x float64, y float64, z float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_ellint_RF(C.double(x), C.double(y), C.double(z), C.gsl_mode_t(mode)))
}

func RFE(x float64, y float64, z float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_ellint_RF_e(C.double(x), C.double(y), C.double(z), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func RJ(x float64, y float64, z float64, p float64, mode sf.GslMode) float64 {
	return float64(C.gsl_sf_ellint_RJ(C.double(x), C.double(y), C.double(z), C.double(p), C.gsl_mode_t(mode)))
}

func RJE(x float64, y float64, z float64, p float64, mode sf.GslMode, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_ellint_RJ_e(C.double(x), C.double(y), C.double(z), C.double(p), C.gsl_mode_t(mode), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}
