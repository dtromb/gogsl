//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package gamma

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_sf_result.h>
   #include <gsl/gsl_sf_gamma.h>
*/
import "C"

import "github.com/dtromb/gogsl/sf"
import "unsafe"

func Gamma(x float64) float64 {
   return float64(C.gsl_sf_gamma(C.double(x)))
}

func GammaE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_gamma_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Lngamma(x float64) float64 {
   return float64(C.gsl_sf_lngamma(C.double(x)))
}

func LngammaE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_lngamma_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func LngammaSgnE(x float64, resultLg *sf.GslSfResult) (int32, float64) {
   var _outptr_2 C.double
   _result := int32(C.gsl_sf_lngamma_sgn_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(resultLg.Ptr())), &_outptr_2))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_2))
}

func Gammastar(x float64) float64 {
   return float64(C.gsl_sf_gammastar(C.double(x)))
}

func GammastarE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_gammastar_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Gammainv(x float64) float64 {
   return float64(C.gsl_sf_gammainv(C.double(x)))
}

func GammainvE(x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_gammainv_e(C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func LngammaComplexE(zr float64, zi float64, lnr *sf.GslSfResult, arg *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_lngamma_complex_e(C.double(zr), C.double(zi), (*C.gsl_sf_result)(unsafe.Pointer(lnr.Ptr())), (*C.gsl_sf_result)(unsafe.Pointer(arg.Ptr()))))
}

func Fact(n uint32) float64 {
   return float64(C.gsl_sf_fact(C.uint(n)))
}

func FactE(n uint32, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_fact_e(C.uint(n), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Doublefact(n uint32) float64 {
   return float64(C.gsl_sf_doublefact(C.uint(n)))
}

func DoublefactE(n uint32, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_doublefact_e(C.uint(n), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Lnfact(n uint32) float64 {
   return float64(C.gsl_sf_lnfact(C.uint(n)))
}

func LnfactE(n uint32, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_lnfact_e(C.uint(n), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Lndoublefact(n uint32) float64 {
   return float64(C.gsl_sf_lndoublefact(C.uint(n)))
}

func LndoublefactE(n uint32, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_lndoublefact_e(C.uint(n), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Choose(n uint32, m uint32) float64 {
   return float64(C.gsl_sf_choose(C.uint(n), C.uint(m)))
}

func ChooseE(n uint32, m uint32, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_choose_e(C.uint(n), C.uint(m), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Lnchoose(n uint32, m uint32) float64 {
   return float64(C.gsl_sf_lnchoose(C.uint(n), C.uint(m)))
}

func LnchooseE(n uint32, m uint32, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_lnchoose_e(C.uint(n), C.uint(m), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Taylorcoeff(n int32, x float64) float64 {
   return float64(C.gsl_sf_taylorcoeff(C.int(n), C.double(x)))
}

func TaylorcoeffE(n int32, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_taylorcoeff_e(C.int(n), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Poch(a float64, x float64) float64 {
   return float64(C.gsl_sf_poch(C.double(a), C.double(x)))
}

func PochE(a float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_poch_e(C.double(a), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Lnpoch(a float64, x float64) float64 {
   return float64(C.gsl_sf_lnpoch(C.double(a), C.double(x)))
}

func LnpochE(a float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_lnpoch_e(C.double(a), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func LnpochSgnE(a float64, x float64, result *sf.GslSfResult) (int32, float64) {
   var _outptr_3 C.double
   _result := int32(C.gsl_sf_lnpoch_sgn_e(C.double(a), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr())), &_outptr_3))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_3))
}

func Pochrel(a float64, x float64) float64 {
   return float64(C.gsl_sf_pochrel(C.double(a), C.double(x)))
}

func PochrelE(a float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_pochrel_e(C.double(a), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GammaInc(a float64, x float64) float64 {
   return float64(C.gsl_sf_gamma_inc(C.double(a), C.double(x)))
}

func GammaIncE(a float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_gamma_inc_e(C.double(a), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GammaIncQ(a float64, x float64) float64 {
   return float64(C.gsl_sf_gamma_inc_Q(C.double(a), C.double(x)))
}

func GammaIncQE(a float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_gamma_inc_Q_e(C.double(a), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func GammaIncP(a float64, x float64) float64 {
   return float64(C.gsl_sf_gamma_inc_P(C.double(a), C.double(x)))
}

func GammaIncPE(a float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_gamma_inc_P_e(C.double(a), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Beta(a float64, b float64) float64 {
   return float64(C.gsl_sf_beta(C.double(a), C.double(b)))
}

func BetaE(a float64, b float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_beta_e(C.double(a), C.double(b), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func Lnbeta(a float64, b float64) float64 {
   return float64(C.gsl_sf_lnbeta(C.double(a), C.double(b)))
}

func LnbetaE(a float64, b float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_lnbeta_e(C.double(a), C.double(b), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

func BetaInc(a float64, b float64, x float64) float64 {
   return float64(C.gsl_sf_beta_inc(C.double(a), C.double(b), C.double(x)))
}

func BetaIncE(a float64, b float64, x float64, result *sf.GslSfResult) int32 {
   return int32(C.gsl_sf_beta_inc_e(C.double(a), C.double(b), C.double(x), (*C.gsl_sf_result)(unsafe.Pointer(result.Ptr()))))
}

