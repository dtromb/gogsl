//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package min

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_min.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"

type GslMinFminimizer struct {
   gogsl.GslReference
}

type GslMinFminimizerType struct {
   gogsl.GslReference
}

func FminimizerAlloc(t *GslMinFminimizerType) *GslMinFminimizer {
   _ref := C.gsl_min_fminimizer_alloc((*C.gsl_min_fminimizer_type)(unsafe.Pointer(t.Ptr())))
   _result := &GslMinFminimizer{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func FminimizerSet(s *GslMinFminimizer, f *gogsl.GslFunction, xMinimum float64, xLower float64, xUpper float64) int32 {
   gogsl.InitializeGslFunction(f)
   return int32(C.gsl_min_fminimizer_set((*C.gsl_min_fminimizer)(unsafe.Pointer(s.Ptr())), (*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(xMinimum), C.double(xLower), C.double(xUpper)))
}

func FminimizerSetWithValues(s *GslMinFminimizer, f *gogsl.GslFunction, xMinimum float64, fMinimum float64, xLower float64, fLower float64, xUpper float64, fUpper float64) int32 {
   gogsl.InitializeGslFunction(f)
   return int32(C.gsl_min_fminimizer_set_with_values((*C.gsl_min_fminimizer)(unsafe.Pointer(s.Ptr())), (*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(xMinimum), C.double(fMinimum), C.double(xLower), C.double(fLower), C.double(xUpper), C.double(fUpper)))
}

func (x *GslMinFminimizer) Dispose() {
   C.gsl_min_fminimizer_free((*C.gsl_min_fminimizer)(unsafe.Pointer(x.Ptr())))
}

func FminimizerName(s *GslMinFminimizer) string {
   return C.GoString(C.gsl_min_fminimizer_name((*C.gsl_min_fminimizer)(unsafe.Pointer(s.Ptr()))))
}

func FminimizerIterate(s *GslMinFminimizer) int32 {
   return int32(C.gsl_min_fminimizer_iterate((*C.gsl_min_fminimizer)(unsafe.Pointer(s.Ptr()))))
}

func FminimizerXMinimum(s *GslMinFminimizer) float64 {
   return float64(C.gsl_min_fminimizer_x_minimum((*C.gsl_min_fminimizer)(unsafe.Pointer(s.Ptr()))))
}

func FminimizerXUpper(s *GslMinFminimizer) float64 {
   return float64(C.gsl_min_fminimizer_x_upper((*C.gsl_min_fminimizer)(unsafe.Pointer(s.Ptr()))))
}

func FminimizerXLower(s *GslMinFminimizer) float64 {
   return float64(C.gsl_min_fminimizer_x_lower((*C.gsl_min_fminimizer)(unsafe.Pointer(s.Ptr()))))
}

func FminimizerFMinimum(s *GslMinFminimizer) float64 {
   return float64(C.gsl_min_fminimizer_f_minimum((*C.gsl_min_fminimizer)(unsafe.Pointer(s.Ptr()))))
}

func FminimizerFUpper(s *GslMinFminimizer) float64 {
   return float64(C.gsl_min_fminimizer_f_upper((*C.gsl_min_fminimizer)(unsafe.Pointer(s.Ptr()))))
}

func FminimizerFLower(s *GslMinFminimizer) float64 {
   return float64(C.gsl_min_fminimizer_f_lower((*C.gsl_min_fminimizer)(unsafe.Pointer(s.Ptr()))))
}

func TestInterval(xLower float64, xUpper float64, epsabs float64, epsrel float64) int32 {
   return int32(C.gsl_min_test_interval(C.double(xLower), C.double(xUpper), C.double(epsabs), C.double(epsrel)))
}

