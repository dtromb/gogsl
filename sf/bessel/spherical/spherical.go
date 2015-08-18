//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package spherical

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
   return float64(C.gsl_sf_bessel_j0(C.double(x)))
}

func J0E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_j0_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func J1(x float64) float64 {
   return float64(C.gsl_sf_bessel_j1(C.double(x)))
}

func J1E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_j1_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func J2(x float64) float64 {
   return float64(C.gsl_sf_bessel_j2(C.double(x)))
}

func J2E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_j2_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Jl(l int32, x float64) float64 {
   return float64(C.gsl_sf_bessel_jl(C.int(l), C.double(x)))
}

func JlE(l int32, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_jl_e(C.int(l), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func JlArray(lmax int32, x float64, resultArray []float64) int32 {
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
   return int32(C.gsl_sf_bessel_jl_array(C.int(lmax), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_2.Data))))
}

func JlSteedArray(lmax int32, x float64, resultArray []float64) int32 {
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
   return int32(C.gsl_sf_bessel_jl_steed_array(C.int(lmax), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_2.Data))))
}

func Y0(x float64) float64 {
   return float64(C.gsl_sf_bessel_y0(C.double(x)))
}

func Y0E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_y0_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Y1(x float64) float64 {
   return float64(C.gsl_sf_bessel_y1(C.double(x)))
}

func Y1E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_y1_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Y2(x float64) float64 {
   return float64(C.gsl_sf_bessel_y2(C.double(x)))
}

func Y2E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_y2_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Yl(l int32, x float64) float64 {
   return float64(C.gsl_sf_bessel_yl(C.int(l), C.double(x)))
}

func YlE(l int32, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_yl_e(C.int(l), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func YlArray(lmax int32, x float64, resultArray []float64) int32 {
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
   return int32(C.gsl_sf_bessel_yl_array(C.int(lmax), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_2.Data))))
}

func I0Scaled(x float64) float64 {
   return float64(C.gsl_sf_bessel_i0_scaled(C.double(x)))
}

func I0ScaledE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_i0_scaled_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func I1Scaled(x float64) float64 {
   return float64(C.gsl_sf_bessel_i1_scaled(C.double(x)))
}

func I1ScaledE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_i1_scaled_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func I2Scaled(x float64) float64 {
   return float64(C.gsl_sf_bessel_i2_scaled(C.double(x)))
}

func I2ScaledE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_i2_scaled_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func IlScaled(l int32, x float64) float64 {
   return float64(C.gsl_sf_bessel_il_scaled(C.int(l), C.double(x)))
}

func IlScaledE(l int32, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_il_scaled_e(C.int(l), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func IlScaledArray(lmax int32, x float64, resultArray []float64) int32 {
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
   return int32(C.gsl_sf_bessel_il_scaled_array(C.int(lmax), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_2.Data))))
}

func K0Scaled(x float64) float64 {
   return float64(C.gsl_sf_bessel_k0_scaled(C.double(x)))
}

func K0ScaledE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_k0_scaled_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func K1Scaled(x float64) float64 {
   return float64(C.gsl_sf_bessel_k1_scaled(C.double(x)))
}

func K1ScaledE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_k1_scaled_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func K2Scaled(x float64) float64 {
   return float64(C.gsl_sf_bessel_k2_scaled(C.double(x)))
}

func K2ScaledE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_k2_scaled_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func KlScaled(l int32, x float64) float64 {
   return float64(C.gsl_sf_bessel_kl_scaled(C.int(l), C.double(x)))
}

func KlScaledE(l int32, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_bessel_kl_scaled_e(C.int(l), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func KlScaledArray(lmax int32, x float64, resultArray []float64) int32 {
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
   return int32(C.gsl_sf_bessel_kl_scaled_array(C.int(lmax), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_2.Data))))
}

