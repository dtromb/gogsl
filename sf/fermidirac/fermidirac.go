//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package fermidirac

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_fermi_dirac.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func M1(x float64) float64 {
   return float64(C.gsl_sf_fermi_dirac_m1(C.double(x)))
}

func M1E(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_fermi_dirac_m1_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Zero(x float64) float64 {
   return float64(C.gsl_sf_fermi_dirac_0(C.double(x)))
}

func ZeroE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_fermi_dirac_0_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func One(x float64) float64 {
   return float64(C.gsl_sf_fermi_dirac_1(C.double(x)))
}

func OneE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_fermi_dirac_1_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Two(x float64) float64 {
   return float64(C.gsl_sf_fermi_dirac_2(C.double(x)))
}

func TwoE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_fermi_dirac_2_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Int(j int32, x float64) float64 {
   return float64(C.gsl_sf_fermi_dirac_int(C.int(j), C.double(x)))
}

func IntE(j int32, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_fermi_dirac_int_e(C.int(j), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Mhalf(x float64) float64 {
   return float64(C.gsl_sf_fermi_dirac_mhalf(C.double(x)))
}

func MhalfE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_fermi_dirac_mhalf_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Half(x float64) float64 {
   return float64(C.gsl_sf_fermi_dirac_half(C.double(x)))
}

func HalfE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_fermi_dirac_half_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Threehalf(x float64) float64 {
   return float64(C.gsl_sf_fermi_dirac_3half(C.double(x)))
}

func ThreehalfE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_fermi_dirac_3half_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Inc0(x float64, b float64) float64 {
   return float64(C.gsl_sf_fermi_dirac_inc_0(C.double(x), C.double(b)))
}

func Inc0E(x float64, b float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_fermi_dirac_inc_0_e(C.double(x), C.double(b), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

