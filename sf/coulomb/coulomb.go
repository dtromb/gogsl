//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package coulomb

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_coulomb.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"
import "reflect"

func HydrogenicR1(z float64, r float64) float64 {
   return float64(C.gsl_sf_hydrogenicR_1(C.double(z), C.double(r)))
}

func HydrogenicR1E(z float64, r float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_hydrogenicR_1_e(C.double(z), C.double(r), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func HydrogenicR(n int32, l int32, z float64, r float64) float64 {
   return float64(C.gsl_sf_hydrogenicR(C.int(n), C.int(l), C.double(z), C.double(r)))
}

func HydrogenicRE(n int32, l int32, z float64, r float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_hydrogenicR_e(C.int(n), C.int(l), C.double(z), C.double(r), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func WaveFGE(eta float64, x float64, lF float64, k int32, f *sf.GslSfResult, fp *sf.GslSfResult, g *sf.GslSfResult, gp *sf.GslSfResult, expF []float64, expG []float64) int32 {
   _slice_header_8 := (*reflect.SliceHeader)(unsafe.Pointer(&expF))
   _slice_header_9 := (*reflect.SliceHeader)(unsafe.Pointer(&expG))
   return int32(C.gsl_sf_coulomb_wave_FG_e(C.double(eta), C.double(x), C.double(lF), C.int(k), (*C.gsl_sf_result)(unsafe.Pointer(f.Ptr())), (*C.gsl_sf_result)(unsafe.Pointer(fp.Ptr())), (*C.gsl_sf_result)(unsafe.Pointer(g.Ptr())), (*C.gsl_sf_result)(unsafe.Pointer(gp.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_8.Data)), (*C.double)(unsafe.Pointer(_slice_header_9.Data))))
}

func WaveFArray(lMin float64, kmax int32, eta float64, x float64, fcArray []float64, fExponent []float64) int32 {
   _slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&fcArray))
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&fExponent))
   return int32(C.gsl_sf_coulomb_wave_F_array(C.double(lMin), C.int(kmax), C.double(eta), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_4.Data)), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func WaveFGArray(lMin float64, kmax int32, eta float64, x float64, fcArray []float64, gcArray []float64, fExponent []float64, gExponent []float64) int32 {
   _slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&fcArray))
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&gcArray))
   _slice_header_6 := (*reflect.SliceHeader)(unsafe.Pointer(&fExponent))
   _slice_header_7 := (*reflect.SliceHeader)(unsafe.Pointer(&gExponent))
   return int32(C.gsl_sf_coulomb_wave_FG_array(C.double(lMin), C.int(kmax), C.double(eta), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_4.Data)), (*C.double)(unsafe.Pointer(_slice_header_5.Data)), (*C.double)(unsafe.Pointer(_slice_header_6.Data)), (*C.double)(unsafe.Pointer(_slice_header_7.Data))))
}

func WaveFGpArray(lMin float64, kmax int32, eta float64, x float64, fcArray []float64, fcpArray []float64, gcArray []float64, gcpArray []float64, fExponent []float64, gExponent []float64) int32 {
   _slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&fcArray))
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&fcpArray))
   _slice_header_6 := (*reflect.SliceHeader)(unsafe.Pointer(&gcArray))
   _slice_header_7 := (*reflect.SliceHeader)(unsafe.Pointer(&gcpArray))
   _slice_header_8 := (*reflect.SliceHeader)(unsafe.Pointer(&fExponent))
   _slice_header_9 := (*reflect.SliceHeader)(unsafe.Pointer(&gExponent))
   return int32(C.gsl_sf_coulomb_wave_FGp_array(C.double(lMin), C.int(kmax), C.double(eta), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_4.Data)), (*C.double)(unsafe.Pointer(_slice_header_5.Data)), (*C.double)(unsafe.Pointer(_slice_header_6.Data)), (*C.double)(unsafe.Pointer(_slice_header_7.Data)), (*C.double)(unsafe.Pointer(_slice_header_8.Data)), (*C.double)(unsafe.Pointer(_slice_header_9.Data))))
}

func WaveSphFArray(lMin float64, kmax int32, eta float64, x float64, fcArray []float64, fExponent []float64) int32 {
   _slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&fcArray))
   _slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&fExponent))
   return int32(C.gsl_sf_coulomb_wave_sphF_array(C.double(lMin), C.int(kmax), C.double(eta), C.double(x), (*C.double)(unsafe.Pointer(_slice_header_4.Data)), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func CLE(l float64, eta float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_coulomb_CL_e(C.double(l), C.double(eta), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func CLArray(lmin float64, kmax int32, eta float64, cl []float64) int32 {
   _slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&cl))
   return int32(C.gsl_sf_coulomb_CL_array(C.double(lmin), C.int(kmax), C.double(eta), (*C.double)(unsafe.Pointer(_slice_header_3.Data))))
}

