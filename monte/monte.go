//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package monte

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_monte_plain.h>
   #include <gsl/gsl_monte_miser.h>
   #include <gsl/gsl_monte_vegas.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "github.com/dtromb/gogsl/rng"
import "reflect"

type GslMontePlainState struct {
   gogsl.GslReference
}

type GslMonteMiserState struct {
   gogsl.GslReference
}

type GslMonteMiserParams struct {
   gogsl.GslReference
}

type GslMonteVegasState struct {
   gogsl.GslReference
}

type GslMonteVegasParams struct {
   gogsl.GslReference
}

func PlainAlloc(dim int) *GslMontePlainState {
   _ref := C.gsl_monte_plain_alloc(C.size_t(dim))
   _result := &GslMontePlainState{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func PlainInit(s *GslMontePlainState) int32 {
   return int32(C.gsl_monte_plain_init((*C.gsl_monte_plain_state)(unsafe.Pointer(s.Ptr()))))
}

func PlainIntegrate(f *GslMonteFunction, xl []float64, xu []float64, dim int, calls int, r *rng.GslRng, s *GslMontePlainState) (int32, float64, float64) {
   var _outptr_7 C.double
   var _outptr_8 C.double
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xl))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&xu))
   _result := int32(C.gsl_monte_plain_integrate((*C.gsl_monte_function)(unsafe.Pointer(f.CPtr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(dim), C.size_t(calls), (*C.gsl_rng)(unsafe.Pointer(r.Ptr())), (*C.gsl_monte_plain_state)(unsafe.Pointer(s.Ptr())), &_outptr_7, &_outptr_8))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_7)), *(*float64)(unsafe.Pointer(&_outptr_8))
}

func (x *GslMontePlainState) Dispose() {
   C.gsl_monte_plain_free((*C.gsl_monte_plain_state)(unsafe.Pointer(x.Ptr())))
}

func MiserAlloc(dim int) *GslMonteMiserState {
   _ref := C.gsl_monte_miser_alloc(C.size_t(dim))
   _result := &GslMonteMiserState{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func MiserInit(s *GslMonteMiserState) int32 {
   return int32(C.gsl_monte_miser_init((*C.gsl_monte_miser_state)(unsafe.Pointer(s.Ptr()))))
}

func MiserIntegrate(f *GslMonteFunction, xl []float64, xu []float64, dim int, calls int, r *rng.GslRng, s *GslMonteMiserState) (int32, float64, float64) {
   var _outptr_7 C.double
   var _outptr_8 C.double
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xl))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&xu))
   _result := int32(C.gsl_monte_miser_integrate((*C.gsl_monte_function)(unsafe.Pointer(f.CPtr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(dim), C.size_t(calls), (*C.gsl_rng)(unsafe.Pointer(r.Ptr())), (*C.gsl_monte_miser_state)(unsafe.Pointer(s.Ptr())), &_outptr_7, &_outptr_8))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_7)), *(*float64)(unsafe.Pointer(&_outptr_8))
}

func (x *GslMonteMiserState) Dispose() {
   C.gsl_monte_miser_free((*C.gsl_monte_miser_state)(unsafe.Pointer(x.Ptr())))
}

func MiserParamsGet(s *GslMonteMiserState, params *GslMonteMiserParams) {
   C.gsl_monte_miser_params_get((*C.gsl_monte_miser_state)(unsafe.Pointer(s.Ptr())), (*C.gsl_monte_miser_params)(unsafe.Pointer(params.Ptr())))
}

func MiserParamsSet(s *GslMonteMiserState, params *GslMonteMiserParams) {
   C.gsl_monte_miser_params_set((*C.gsl_monte_miser_state)(unsafe.Pointer(s.Ptr())), (*C.gsl_monte_miser_params)(unsafe.Pointer(params.Ptr())))
}

func VegasAlloc(dim int) *GslMonteVegasState {
   _ref := C.gsl_monte_vegas_alloc(C.size_t(dim))
   _result := &GslMonteVegasState{}
   gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
   return _result
}

func VegasInit(s *GslMonteVegasState) int32 {
   return int32(C.gsl_monte_vegas_init((*C.gsl_monte_vegas_state)(unsafe.Pointer(s.Ptr()))))
}

func VegasIntegrate(f *GslMonteFunction, xl []float64, xu []float64, dim int, calls int, r *rng.GslRng, s *GslMonteVegasState) (int32, float64, float64) {
   var _outptr_7 C.double
   var _outptr_8 C.double
   _slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&xl))
   _slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&xu))
   _result := int32(C.gsl_monte_vegas_integrate((*C.gsl_monte_function)(unsafe.Pointer(f.CPtr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), C.size_t(dim), C.size_t(calls), (*C.gsl_rng)(unsafe.Pointer(r.Ptr())), (*C.gsl_monte_vegas_state)(unsafe.Pointer(s.Ptr())), &_outptr_7, &_outptr_8))
   return _result, *(*float64)(unsafe.Pointer(&_outptr_7)), *(*float64)(unsafe.Pointer(&_outptr_8))
}

func (x *GslMonteVegasState) Dispose() {
   C.gsl_monte_vegas_free((*C.gsl_monte_vegas_state)(unsafe.Pointer(x.Ptr())))
}

func VegasChisq(s *GslMonteVegasState) float64 {
   return float64(C.gsl_monte_vegas_chisq((*C.gsl_monte_vegas_state)(unsafe.Pointer(s.Ptr()))))
}

func VegasRunval(s *GslMonteVegasState) (float64, float64) {
   var _outptr_1 C.double
   var _outptr_2 C.double
   C.gsl_monte_vegas_runval((*C.gsl_monte_vegas_state)(unsafe.Pointer(s.Ptr())), &_outptr_1, &_outptr_2)
   return *(*float64)(unsafe.Pointer(&_outptr_1)), *(*float64)(unsafe.Pointer(&_outptr_2))
}

func VegasParamsGet(s *GslMonteVegasState, params *GslMonteVegasParams) {
   C.gsl_monte_vegas_params_get((*C.gsl_monte_vegas_state)(unsafe.Pointer(s.Ptr())), (*C.gsl_monte_vegas_params)(unsafe.Pointer(params.Ptr())))
}

func VegasParamsSet(s *GslMonteVegasState, params *GslMonteVegasParams) {
   C.gsl_monte_vegas_params_set((*C.gsl_monte_vegas_state)(unsafe.Pointer(s.Ptr())), (*C.gsl_monte_vegas_params)(unsafe.Pointer(params.Ptr())))
}

