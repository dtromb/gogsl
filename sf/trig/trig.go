//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package trig

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_trig.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func Sin(x float64) float64 {
   return float64(C.gsl_sf_sin(C.double(x)))
}

func SinE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_sin_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Cos(x float64) float64 {
   return float64(C.gsl_sf_cos(C.double(x)))
}

func CosE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_cos_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Hypot(x float64, y float64) float64 {
   return float64(C.gsl_sf_hypot(C.double(x), C.double(y)))
}

func HypotE(x float64, y float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_hypot_e(C.double(x), C.double(y), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Sinc(x float64) float64 {
   return float64(C.gsl_sf_sinc(C.double(x)))
}

func SincE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_sinc_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ComplexSinE(zr float64, zi float64, szr *sf.GslSfResult, szi *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_complex_sin_e(C.double(zr), C.double(zi), (*C.gsl_sf_result)(unsafe.Pointer(szr.Ptr())), (*C.gsl_sf_result)(unsafe.Pointer(szi.Ptr()))))
}

func ComplexCosE(zr float64, zi float64, czr *sf.GslSfResult, czi *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_complex_cos_e(C.double(zr), C.double(zi), (*C.gsl_sf_result)(unsafe.Pointer(czr.Ptr())), (*C.gsl_sf_result)(unsafe.Pointer(czi.Ptr()))))
}

func ComplexLogsinE(zr float64, zi float64, lszr *sf.GslSfResult, lszi *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_complex_logsin_e(C.double(zr), C.double(zi), (*C.gsl_sf_result)(unsafe.Pointer(lszr.Ptr())), (*C.gsl_sf_result)(unsafe.Pointer(lszi.Ptr()))))
}

func Lnsinh(x float64) float64 {
   return float64(C.gsl_sf_lnsinh(C.double(x)))
}

func LnsinhE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_lnsinh_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Lncosh(x float64) float64 {
   return float64(C.gsl_sf_lncosh(C.double(x)))
}

func LncoshE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_lncosh_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func PolarToRect(r float64, theta float64, x *sf.GslSfResult, y *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_polar_to_rect(C.double(r), C.double(theta), (*C.gsl_sf_result)(unsafe.Pointer(x.Ptr())), (*C.gsl_sf_result)(unsafe.Pointer(y.Ptr()))))
}

func RectToPolar(x float64, y float64, r *sf.GslSfResult, theta *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_rect_to_polar(C.double(x), C.double(y), (*C.gsl_sf_result)(unsafe.Pointer(r.Ptr())), (*C.gsl_sf_result)(unsafe.Pointer(theta.Ptr()))))
}

func AngleRestrictSymm(theta float64) float64 {
   return float64(C.gsl_sf_angle_restrict_symm(C.double(theta)))
}

func AngleRestrictPos(theta float64) float64 {
   return float64(C.gsl_sf_angle_restrict_pos(C.double(theta)))
}

func SinErrE(x float64, dx float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_sin_err_e(C.double(x), C.double(dx), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func CosErrE(x float64, dx float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_cos_err_e(C.double(x), C.double(dx), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

