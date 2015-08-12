//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package ode

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_odeiv2.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "reflect"

type GslOdeiv2Step struct {
	gogsl.GslReference
}

type GslOdeiv2StepType struct {
	gogsl.GslReference
}

type GslOdeiv2Driver struct {
	gogsl.GslReference
}

type GslOdeiv2Control struct {
	gogsl.GslReference
}

type GslOdeiv2ControlType struct {
	gogsl.GslReference
}

type GslOdeiv2Evolve struct {
	gogsl.GslReference
}

func StepAlloc(t *GslOdeiv2StepType, dim int) *GslOdeiv2Step {
	_ref := C.gsl_odeiv2_step_alloc((*C.gsl_odeiv2_step_type)(unsafe.Pointer(t.Ptr())), C.size_t(dim))
	_result := &GslOdeiv2Step{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func StepReset(s *GslOdeiv2Step) int32 {
	return int32(C.gsl_odeiv2_step_reset((*C.gsl_odeiv2_step)(unsafe.Pointer(s.Ptr()))))
}

func (x *GslOdeiv2Step) Dispose() {
	C.gsl_odeiv2_step_free((*C.gsl_odeiv2_step)(unsafe.Pointer(x.Ptr())))
}

func StepName(s *GslOdeiv2Step) string {
	return C.GoString(C.gsl_odeiv2_step_name((*C.gsl_odeiv2_step)(unsafe.Pointer(s.Ptr()))))
}

func StepOrder(s *GslOdeiv2Step) uint32 {
	return uint32(C.gsl_odeiv2_step_order((*C.gsl_odeiv2_step)(unsafe.Pointer(s.Ptr()))))
}

func StepSetDriver(s *GslOdeiv2Step, d *GslOdeiv2Driver) int32 {
	return int32(C.gsl_odeiv2_step_set_driver((*C.gsl_odeiv2_step)(unsafe.Pointer(s.Ptr())), (*C.gsl_odeiv2_driver)(unsafe.Pointer(d.Ptr()))))
}

func StepApply(s *GslOdeiv2Step, t float64, h float64, y []float64, yerr []float64, dydtIn []float64, dydtOut []float64, sys *GslOdeiv2System) int32 {
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&y))
	_slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&yerr))
	_slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&dydtIn))
	_slice_header_6 := (*reflect.SliceHeader)(unsafe.Pointer(&dydtOut))
	return int32(C.gsl_odeiv2_step_apply((*C.gsl_odeiv2_step)(unsafe.Pointer(s.Ptr())), C.double(t), C.double(h), (*C.double)(unsafe.Pointer(_slice_header_3.Data)), (*C.double)(unsafe.Pointer(_slice_header_4.Data)), (*C.double)(unsafe.Pointer(_slice_header_5.Data)), (*C.double)(unsafe.Pointer(_slice_header_6.Data)), (*C.gsl_odeiv2_system)(unsafe.Pointer(sys.CPtr()))))
}

func ControlStandardNew(epsAbs float64, epsRel float64, aY float64, aDydt float64) *GslOdeiv2Control {
	_ref := C.gsl_odeiv2_control_standard_new(C.double(epsAbs), C.double(epsRel), C.double(aY), C.double(aDydt))
	_result := &GslOdeiv2Control{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func ControlYNew(epsAbs float64, epsRel float64) *GslOdeiv2Control {
	_ref := C.gsl_odeiv2_control_y_new(C.double(epsAbs), C.double(epsRel))
	_result := &GslOdeiv2Control{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func ControlYpNew(epsAbs float64, epsRel float64) *GslOdeiv2Control {
	_ref := C.gsl_odeiv2_control_yp_new(C.double(epsAbs), C.double(epsRel))
	_result := &GslOdeiv2Control{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func ControlScaledNew(epsAbs float64, epsRel float64, aY float64, aDydt float64, scaleAbs []float64, dim int) *GslOdeiv2Control {
	_slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&scaleAbs))
	_ref := C.gsl_odeiv2_control_scaled_new(C.double(epsAbs), C.double(epsRel), C.double(aY), C.double(aDydt), (*C.double)(unsafe.Pointer(_slice_header_4.Data)), C.size_t(dim))
	_result := &GslOdeiv2Control{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func ControlAlloc(t *GslOdeiv2ControlType) *GslOdeiv2Control {
	_ref := C.gsl_odeiv2_control_alloc((*C.gsl_odeiv2_control_type)(unsafe.Pointer(t.Ptr())))
	_result := &GslOdeiv2Control{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func ControlInit(c *GslOdeiv2Control, epsAbs float64, epsRel float64, aY float64, aDydt float64) int32 {
	return int32(C.gsl_odeiv2_control_init((*C.gsl_odeiv2_control)(unsafe.Pointer(c.Ptr())), C.double(epsAbs), C.double(epsRel), C.double(aY), C.double(aDydt)))
}

func (x *GslOdeiv2Control) Dispose() {
	C.gsl_odeiv2_control_free((*C.gsl_odeiv2_control)(unsafe.Pointer(x.Ptr())))
}

func ControlHadjust(c *GslOdeiv2Control, s *GslOdeiv2Step, y []float64, yerr []float64, dydt []float64, h []float64) int32 {
	_slice_header_2 := (*reflect.SliceHeader)(unsafe.Pointer(&y))
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&yerr))
	_slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&dydt))
	_slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&h))
	return int32(C.gsl_odeiv2_control_hadjust((*C.gsl_odeiv2_control)(unsafe.Pointer(c.Ptr())), (*C.gsl_odeiv2_step)(unsafe.Pointer(s.Ptr())), (*C.double)(unsafe.Pointer(_slice_header_2.Data)), (*C.double)(unsafe.Pointer(_slice_header_3.Data)), (*C.double)(unsafe.Pointer(_slice_header_4.Data)), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func ControlName(c *GslOdeiv2Control) string {
	return C.GoString(C.gsl_odeiv2_control_name((*C.gsl_odeiv2_control)(unsafe.Pointer(c.Ptr()))))
}

func ControlErrlevel(c *GslOdeiv2Control, y float64, dydt float64, h float64, ind int, errlev []float64) int32 {
	_slice_header_5 := (*reflect.SliceHeader)(unsafe.Pointer(&errlev))
	return int32(C.gsl_odeiv2_control_errlevel((*C.gsl_odeiv2_control)(unsafe.Pointer(c.Ptr())), C.double(y), C.double(dydt), C.double(h), C.size_t(ind), (*C.double)(unsafe.Pointer(_slice_header_5.Data))))
}

func ControlSetDriver(c *GslOdeiv2Control, d *GslOdeiv2Driver) int32 {
	return int32(C.gsl_odeiv2_control_set_driver((*C.gsl_odeiv2_control)(unsafe.Pointer(c.Ptr())), (*C.gsl_odeiv2_driver)(unsafe.Pointer(d.Ptr()))))
}

func EvolveAlloc(dim int) *GslOdeiv2Evolve {
	_ref := C.gsl_odeiv2_evolve_alloc(C.size_t(dim))
	_result := &GslOdeiv2Evolve{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func EvolveApply(e *GslOdeiv2Evolve, con *GslOdeiv2Control, step *GslOdeiv2Step, sys *GslOdeiv2System, t1 float64, y []float64) (int32, float64, float64) {
	var _outptr_4 C.double
	var _outptr_6 C.double
	_slice_header_7 := (*reflect.SliceHeader)(unsafe.Pointer(&y))
	_result := int32(C.gsl_odeiv2_evolve_apply((*C.gsl_odeiv2_evolve)(unsafe.Pointer(e.Ptr())), (*C.gsl_odeiv2_control)(unsafe.Pointer(con.Ptr())), (*C.gsl_odeiv2_step)(unsafe.Pointer(step.Ptr())), (*C.gsl_odeiv2_system)(unsafe.Pointer(sys.CPtr())), &_outptr_4, C.double(t1), &_outptr_6, (*C.double)(unsafe.Pointer(_slice_header_7.Data))))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_4)), *(*float64)(unsafe.Pointer(&_outptr_6))
}

func EvolveApplyFixedStep(e *GslOdeiv2Evolve, con *GslOdeiv2Control, step *GslOdeiv2Step, sys *GslOdeiv2System, h float64, y []float64) (int32, float64) {
	var _outptr_4 C.double
	_slice_header_6 := (*reflect.SliceHeader)(unsafe.Pointer(&y))
	_result := int32(C.gsl_odeiv2_evolve_apply_fixed_step((*C.gsl_odeiv2_evolve)(unsafe.Pointer(e.Ptr())), (*C.gsl_odeiv2_control)(unsafe.Pointer(con.Ptr())), (*C.gsl_odeiv2_step)(unsafe.Pointer(step.Ptr())), (*C.gsl_odeiv2_system)(unsafe.Pointer(sys.CPtr())), &_outptr_4, C.double(h), (*C.double)(unsafe.Pointer(_slice_header_6.Data))))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_4))
}

func EvolveReset(e *GslOdeiv2Evolve) int32 {
	return int32(C.gsl_odeiv2_evolve_reset((*C.gsl_odeiv2_evolve)(unsafe.Pointer(e.Ptr()))))
}

func (x *GslOdeiv2Evolve) Dispose() {
	C.gsl_odeiv2_evolve_free((*C.gsl_odeiv2_evolve)(unsafe.Pointer(x.Ptr())))
}

func EvolveSetDriver(e *GslOdeiv2Evolve, d *GslOdeiv2Driver) int32 {
	return int32(C.gsl_odeiv2_evolve_set_driver((*C.gsl_odeiv2_evolve)(unsafe.Pointer(e.Ptr())), (*C.gsl_odeiv2_driver)(unsafe.Pointer(d.Ptr()))))
}

func DriverAllocYNew(sys *GslOdeiv2System, t *GslOdeiv2StepType, hstart float64, epsabs float64, epsrel float64) *GslOdeiv2Driver {
	_ref := C.gsl_odeiv2_driver_alloc_y_new((*C.gsl_odeiv2_system)(unsafe.Pointer(sys.CPtr())), (*C.gsl_odeiv2_step_type)(unsafe.Pointer(t.Ptr())), C.double(hstart), C.double(epsabs), C.double(epsrel))
	_result := &GslOdeiv2Driver{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func DriverAllocYpNew(sys *GslOdeiv2System, t *GslOdeiv2StepType, hstart float64, epsabs float64, epsrel float64) *GslOdeiv2Driver {
	_ref := C.gsl_odeiv2_driver_alloc_yp_new((*C.gsl_odeiv2_system)(unsafe.Pointer(sys.CPtr())), (*C.gsl_odeiv2_step_type)(unsafe.Pointer(t.Ptr())), C.double(hstart), C.double(epsabs), C.double(epsrel))
	_result := &GslOdeiv2Driver{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func DriverAllocStandardNew(sys *GslOdeiv2System, t *GslOdeiv2StepType, hstart float64, epsabs float64, epsrel float64, aY float64, aDydt float64) *GslOdeiv2Driver {
	_ref := C.gsl_odeiv2_driver_alloc_standard_new((*C.gsl_odeiv2_system)(unsafe.Pointer(sys.CPtr())), (*C.gsl_odeiv2_step_type)(unsafe.Pointer(t.Ptr())), C.double(hstart), C.double(epsabs), C.double(epsrel), C.double(aY), C.double(aDydt))
	_result := &GslOdeiv2Driver{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func DriverAllocScaledNew(sys *GslOdeiv2System, t *GslOdeiv2StepType, hstart float64, epsabs float64, epsrel float64, aY float64, aDydt float64, scaleAbs []float64) *GslOdeiv2Driver {
	_slice_header_7 := (*reflect.SliceHeader)(unsafe.Pointer(&scaleAbs))
	_ref := C.gsl_odeiv2_driver_alloc_scaled_new((*C.gsl_odeiv2_system)(unsafe.Pointer(sys.CPtr())), (*C.gsl_odeiv2_step_type)(unsafe.Pointer(t.Ptr())), C.double(hstart), C.double(epsabs), C.double(epsrel), C.double(aY), C.double(aDydt), (*C.double)(unsafe.Pointer(_slice_header_7.Data)))
	_result := &GslOdeiv2Driver{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func DriverSetHmin(d *GslOdeiv2Driver, hmin float64) int32 {
	return int32(C.gsl_odeiv2_driver_set_hmin((*C.gsl_odeiv2_driver)(unsafe.Pointer(d.Ptr())), C.double(hmin)))
}

func DriverSetHmax(d *GslOdeiv2Driver, hmax float64) int32 {
	return int32(C.gsl_odeiv2_driver_set_hmax((*C.gsl_odeiv2_driver)(unsafe.Pointer(d.Ptr())), C.double(hmax)))
}

func DriverSetNmax(d *GslOdeiv2Driver, nmax int) int32 {
	return int32(C.gsl_odeiv2_driver_set_nmax((*C.gsl_odeiv2_driver)(unsafe.Pointer(d.Ptr())), C.ulong(nmax)))
}

func DriverApply(d *GslOdeiv2Driver, t1 float64, y []float64) (int32, float64) {
	var _outptr_1 C.double
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&y))
	_result := int32(C.gsl_odeiv2_driver_apply((*C.gsl_odeiv2_driver)(unsafe.Pointer(d.Ptr())), &_outptr_1, C.double(t1), (*C.double)(unsafe.Pointer(_slice_header_3.Data))))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_1))
}

func DriverApplyFixedStep(d *GslOdeiv2Driver, h float64, n int, y []float64) (int32, float64) {
	var _outptr_1 C.double
	_slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&y))
	_result := int32(C.gsl_odeiv2_driver_apply_fixed_step((*C.gsl_odeiv2_driver)(unsafe.Pointer(d.Ptr())), &_outptr_1, C.double(h), C.ulong(n), (*C.double)(unsafe.Pointer(_slice_header_4.Data))))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_1))
}

func DriverReset(d *GslOdeiv2Driver) int32 {
	return int32(C.gsl_odeiv2_driver_reset((*C.gsl_odeiv2_driver)(unsafe.Pointer(d.Ptr()))))
}

func DriverResetHstart(d *GslOdeiv2Driver, hstart float64) int32 {
	return int32(C.gsl_odeiv2_driver_reset_hstart((*C.gsl_odeiv2_driver)(unsafe.Pointer(d.Ptr())), C.double(hstart)))
}

func (x *GslOdeiv2Driver) Dispose() {
	C.gsl_odeiv2_driver_free((*C.gsl_odeiv2_driver)(unsafe.Pointer(x.Ptr())))
}
