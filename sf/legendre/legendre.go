//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package legendre

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_legendre.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"
import "reflect"

func P1(x float64) float64 {
	return float64(C.gsl_sf_legendre_P1(C.double(x)))
}

func P2(x float64) float64 {
	return float64(C.gsl_sf_legendre_P2(C.double(x)))
}

func P3(x float64) float64 {
	return float64(C.gsl_sf_legendre_P3(C.double(x)))
}

func P1E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_legendre_P1_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func P2E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_legendre_P2_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func P3E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_legendre_P3_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Pl(l int32, x float64) float64 {
	return float64(C.gsl_sf_legendre_Pl(C.int(l), C.double(x)))
}

func PlE(l int32, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_legendre_Pl_e(C.int(l), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func PlArray(lmax int32, x float64, resultArray []float64) int32 {
	_slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
	return int32(C.gsl_sf_legendre_Pl_array(C.int(lmax), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_2.Data))))
}

func PlDerivArray(lmax int32, x float64, resultArray []float64, resultDerivArray []float64) int32 {
	_slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&resultDerivArray))
	return int32(C.gsl_sf_legendre_Pl_deriv_array(C.int(lmax), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), (*C.double)(unsafe.Pointer(_slice_header_3.Data))))
}

func Q0(x float64) float64 {
	return float64(C.gsl_sf_legendre_Q0(C.double(x)))
}

func Q0E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_legendre_Q0_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Q1(x float64) float64 {
	return float64(C.gsl_sf_legendre_Q1(C.double(x)))
}

func Q1E(x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_legendre_Q1_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Ql(l int32, x float64) float64 {
	return float64(C.gsl_sf_legendre_Ql(C.int(l), C.double(x)))
}

func QlE(l int32, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_legendre_Ql_e(C.int(l), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Plm(l int32, m int32, x float64) float64 {
	return float64(C.gsl_sf_legendre_Plm(C.int(l), C.int(m), C.double(x)))
}

func PlmE(l int32, m int32, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_legendre_Plm_e(C.int(l), C.int(m), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func PlmArray(lmax int32, m int32, x float64, resultArray []float64) int32 {
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
	return int32(C.gsl_sf_legendre_Plm_array(C.int(lmax), C.int(m), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_3.Data))))
}

func PlmDerivArray(lmax int32, m int32, x float64, resultArray []float64, resultDerivArray []float64) int32 {
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
	_slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&resultDerivArray))
	return int32(C.gsl_sf_legendre_Plm_deriv_array(C.int(lmax), C.int(m), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_3.Data)), (*C.double)(unsafe.Pointer(_slice_header_4.Data))))
}

func SphPlm(l int32, m int32, x float64) float64 {
	return float64(C.gsl_sf_legendre_sphPlm(C.int(l), C.int(m), C.double(x)))
}

func SphPlmE(l int32, m int32, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_legendre_sphPlm_e(C.int(l), C.int(m), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func SphPlmArray(lmax int32, m int32, x float64, resultArray []float64) int32 {
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
	return int32(C.gsl_sf_legendre_sphPlm_array(C.int(lmax), C.int(m), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_3.Data))))
}

func SphPlmDerivArray(lmax int32, m int32, x float64, resultArray []float64, resultDerivArray []float64) int32 {
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
	_slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&resultDerivArray))
	return int32(C.gsl_sf_legendre_sphPlm_deriv_array(C.int(lmax), C.int(m), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_3.Data)), (*C.double)(unsafe.Pointer(_slice_header_4.Data))))
}

func ArraySize(lmax int32, m int32) int32 {
	return int32(C.gsl_sf_legendre_array_size(C.int(lmax), C.int(m)))
}

func ConicalPHalf(lambda float64, x float64) float64 {
	return float64(C.gsl_sf_conicalP_half(C.double(lambda), C.double(x)))
}

func ConicalPHalfE(lambda float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_conicalP_half_e(C.double(lambda), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ConicalPMhalf(lambda float64, x float64) float64 {
	return float64(C.gsl_sf_conicalP_mhalf(C.double(lambda), C.double(x)))
}

func ConicalPMhalfE(lambda float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_conicalP_mhalf_e(C.double(lambda), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ConicalP0(lambda float64, x float64) float64 {
	return float64(C.gsl_sf_conicalP_0(C.double(lambda), C.double(x)))
}

func ConicalP0E(lambda float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_conicalP_0_e(C.double(lambda), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ConicalP1(lambda float64, x float64) float64 {
	return float64(C.gsl_sf_conicalP_1(C.double(lambda), C.double(x)))
}

func ConicalP1E(lambda float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_conicalP_1_e(C.double(lambda), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ConicalPSphReg(l int32, lambda float64, x float64) float64 {
	return float64(C.gsl_sf_conicalP_sph_reg(C.int(l), C.double(lambda), C.double(x)))
}

func ConicalPSphRegE(l int32, lambda float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_conicalP_sph_reg_e(C.int(l), C.double(lambda), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func ConicalPCylReg(m int32, lambda float64, x float64) float64 {
	return float64(C.gsl_sf_conicalP_cyl_reg(C.int(m), C.double(lambda), C.double(x)))
}

func ConicalPCylRegE(m int32, lambda float64, x float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_conicalP_cyl_reg_e(C.int(m), C.double(lambda), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func H3d0(lambda float64, eta float64) float64 {
	return float64(C.gsl_sf_legendre_H3d_0(C.double(lambda), C.double(eta)))
}

func H3d0E(lambda float64, eta float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_legendre_H3d_0_e(C.double(lambda), C.double(eta), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func H3d1(lambda float64, eta float64) float64 {
	return float64(C.gsl_sf_legendre_H3d_1(C.double(lambda), C.double(eta)))
}

func H3d1E(lambda float64, eta float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_legendre_H3d_1_e(C.double(lambda), C.double(eta), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func H3d(l int32, lambda float64, eta float64) float64 {
	return float64(C.gsl_sf_legendre_H3d(C.int(l), C.double(lambda), C.double(eta)))
}

func H3dE(l int32, lambda float64, eta float64, result *sf.GslSfResult) int32 {
	return int32(C.gsl_sf_legendre_H3d_e(C.int(l), C.double(lambda), C.double(eta), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func H3dArray(lmax int32, lambda float64, eta float64, resultArray []float64) int32 {
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&resultArray))
	return int32(C.gsl_sf_legendre_H3d_array(C.int(lmax), C.double(lambda), C.double(eta), (*C.double)(unsafe.Pointer(_slice_header_3.Data))))
}
