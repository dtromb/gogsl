//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////
package numint

/*
   #cgo pkg-config: --define-variable=prefix=. gsl
   #include <gsl/gsl_integration.h>
*/
import "C"

import "unsafe"
import "github.com/dtromb/gogsl"
import "reflect"

type GslIntegrationWorkspace struct {
	gogsl.GslReference
}

type GslIntegrationQawsTable struct {
	gogsl.GslReference
}

type GslIntegrationQawoTable struct {
	gogsl.GslReference
}

type GslIntegrationCquadWorkspace struct {
	gogsl.GslReference
}

type GslIntegrationGlfixedTable struct {
	gogsl.GslReference
}

type GslIntegrationQawoEnum uint32

func Qng(f *gogsl.GslFunction, a float64, b float64, epsabs float64, epsrel float64) (int32, float64, float64, int) {
	var _outptr_5 C.double
	var _outptr_6 C.double
	var _outptr_7 C.size_t
	gogsl.InitializeGslFunction(f)
	_result := int32(C.gsl_integration_qng((*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(a), C.double(b), C.double(epsabs), C.double(epsrel), &_outptr_5, &_outptr_6, &_outptr_7))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_5)), *(*float64)(unsafe.Pointer(&_outptr_6)), *(*int)(unsafe.Pointer(&_outptr_7))
}

func IntegrationWorkspaceAlloc(n int) *GslIntegrationWorkspace {
	_ref := C.gsl_integration_workspace_alloc(C.size_t(n))
	_result := &GslIntegrationWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslIntegrationWorkspace) Dispose() {
	C.gsl_integration_workspace_free((*C.gsl_integration_workspace)(unsafe.Pointer(x.Ptr())))
}

func Qag(f *gogsl.GslFunction, a float64, b float64, epsabs float64, epsrel float64, limit int, key int32, workspace *GslIntegrationWorkspace) (int32, float64, float64) {
	var _outptr_8 C.double
	var _outptr_9 C.double
	gogsl.InitializeGslFunction(f)
	_result := int32(C.gsl_integration_qag((*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(a), C.double(b), C.double(epsabs), C.double(epsrel), C.size_t(limit), C.int(key), (*C.gsl_integration_workspace)(unsafe.Pointer(workspace.Ptr())), &_outptr_8, &_outptr_9))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_8)), *(*float64)(unsafe.Pointer(&_outptr_9))
}

func Qags(f *gogsl.GslFunction, a float64, b float64, epsabs float64, epsrel float64, limit int, workspace *GslIntegrationWorkspace) (int32, float64, float64) {
	var _outptr_7 C.double
	var _outptr_8 C.double
	gogsl.InitializeGslFunction(f)
	_result := int32(C.gsl_integration_qags((*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(a), C.double(b), C.double(epsabs), C.double(epsrel), C.size_t(limit), (*C.gsl_integration_workspace)(unsafe.Pointer(workspace.Ptr())), &_outptr_7, &_outptr_8))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_7)), *(*float64)(unsafe.Pointer(&_outptr_8))
}

func Qagp(f *gogsl.GslFunction, pts []float64, npts int, epsabs float64, epsrel float64, limit int, workspace *GslIntegrationWorkspace) (int32, float64, float64) {
	var _outptr_7 C.double
	var _outptr_8 C.double
	gogsl.InitializeGslFunction(f)
	_slice_header_1 := (*reflect.SliceHeader)(unsafe.Pointer(&pts))
	_result := int32(C.gsl_integration_qagp((*C.gsl_function)(unsafe.Pointer(f.CPtr())), (*C.double)(unsafe.Pointer(_slice_header_1.Data)), C.size_t(npts), C.double(epsabs), C.double(epsrel), C.size_t(limit), (*C.gsl_integration_workspace)(unsafe.Pointer(workspace.Ptr())), &_outptr_7, &_outptr_8))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_7)), *(*float64)(unsafe.Pointer(&_outptr_8))
}

func Qagiu(f *gogsl.GslFunction, a float64, epsabs float64, epsrel float64, limit int, workspace *GslIntegrationWorkspace) (int32, float64, float64) {
	var _outptr_6 C.double
	var _outptr_7 C.double
	gogsl.InitializeGslFunction(f)
	_result := int32(C.gsl_integration_qagiu((*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(a), C.double(epsabs), C.double(epsrel), C.size_t(limit), (*C.gsl_integration_workspace)(unsafe.Pointer(workspace.Ptr())), &_outptr_6, &_outptr_7))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_6)), *(*float64)(unsafe.Pointer(&_outptr_7))
}

func Qagil(f *gogsl.GslFunction, b float64, epsabs float64, epsrel float64, limit int, workspace *GslIntegrationWorkspace) (int32, float64, float64) {
	var _outptr_6 C.double
	var _outptr_7 C.double
	gogsl.InitializeGslFunction(f)
	_result := int32(C.gsl_integration_qagil((*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(b), C.double(epsabs), C.double(epsrel), C.size_t(limit), (*C.gsl_integration_workspace)(unsafe.Pointer(workspace.Ptr())), &_outptr_6, &_outptr_7))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_6)), *(*float64)(unsafe.Pointer(&_outptr_7))
}

func Qawc(f *gogsl.GslFunction, a float64, b float64, c float64, epsabs float64, epsrel float64, limit int, workspace *GslIntegrationWorkspace) (int32, float64, float64) {
	var _outptr_8 C.double
	var _outptr_9 C.double
	gogsl.InitializeGslFunction(f)
	_result := int32(C.gsl_integration_qawc((*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(a), C.double(b), C.double(c), C.double(epsabs), C.double(epsrel), C.size_t(limit), (*C.gsl_integration_workspace)(unsafe.Pointer(workspace.Ptr())), &_outptr_8, &_outptr_9))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_8)), *(*float64)(unsafe.Pointer(&_outptr_9))
}

func QawsTableAlloc(alpha float64, beta float64, mu int32, nu int32) *GslIntegrationQawsTable {
	_ref := C.gsl_integration_qaws_table_alloc(C.double(alpha), C.double(beta), C.int(mu), C.int(nu))
	_result := &GslIntegrationQawsTable{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func QawsTableSet(t *GslIntegrationQawsTable, alpha float64, beta float64, mu int32, nu int32) int32 {
	return int32(C.gsl_integration_qaws_table_set((*C.gsl_integration_qaws_table)(unsafe.Pointer(t.Ptr())), C.double(alpha), C.double(beta), C.int(mu), C.int(nu)))
}

func (x *GslIntegrationQawsTable) Dispose() {
	C.gsl_integration_qaws_table_free((*C.gsl_integration_qaws_table)(unsafe.Pointer(x.Ptr())))
}

func Qaws(f *gogsl.GslFunction, a float64, b float64, t *GslIntegrationQawsTable, epsabs float64, epsrel float64, limit int, workspace *GslIntegrationWorkspace) (int32, float64, float64) {
	var _outptr_8 C.double
	var _outptr_9 C.double
	gogsl.InitializeGslFunction(f)
	_result := int32(C.gsl_integration_qaws((*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(a), C.double(b), (*C.gsl_integration_qaws_table)(unsafe.Pointer(t.Ptr())), C.double(epsabs), C.double(epsrel), C.size_t(limit), (*C.gsl_integration_workspace)(unsafe.Pointer(workspace.Ptr())), &_outptr_8, &_outptr_9))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_8)), *(*float64)(unsafe.Pointer(&_outptr_9))
}

func QawoTableAlloc(omega float64, l float64, sine GslIntegrationQawoEnum, n int) *GslIntegrationQawoTable {
	_ref := C.gsl_integration_qawo_table_alloc(C.double(omega), C.double(l), C.enum_gsl_integration_qawo_enum(sine), C.size_t(n))
	_result := &GslIntegrationQawoTable{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func QawoTableSet(t *GslIntegrationQawoTable, omega float64, l float64, sine GslIntegrationQawoEnum) int32 {
	return int32(C.gsl_integration_qawo_table_set((*C.gsl_integration_qawo_table)(unsafe.Pointer(t.Ptr())), C.double(omega), C.double(l), C.enum_gsl_integration_qawo_enum(sine)))
}

func QawoTableSetLength(t *GslIntegrationQawoTable, l float64) int32 {
	return int32(C.gsl_integration_qawo_table_set_length((*C.gsl_integration_qawo_table)(unsafe.Pointer(t.Ptr())), C.double(l)))
}

func (x *GslIntegrationQawoTable) Dispose() {
	C.gsl_integration_qawo_table_free((*C.gsl_integration_qawo_table)(unsafe.Pointer(x.Ptr())))
}

func Qawo(f *gogsl.GslFunction, a float64, epsabs float64, epsrel float64, limit int, workspace *GslIntegrationWorkspace, wf *GslIntegrationQawoTable) (int32, float64, float64) {
	var _outptr_7 C.double
	var _outptr_8 C.double
	gogsl.InitializeGslFunction(f)
	_result := int32(C.gsl_integration_qawo((*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(a), C.double(epsabs), C.double(epsrel), C.size_t(limit), (*C.gsl_integration_workspace)(unsafe.Pointer(workspace.Ptr())), (*C.gsl_integration_qawo_table)(unsafe.Pointer(wf.Ptr())), &_outptr_7, &_outptr_8))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_7)), *(*float64)(unsafe.Pointer(&_outptr_8))
}

func Qawf(f *gogsl.GslFunction, a float64, epsabs float64, limit int, workspace *GslIntegrationWorkspace, cycleWorkspace *GslIntegrationWorkspace, wf *GslIntegrationQawoTable) (int32, float64, float64) {
	var _outptr_7 C.double
	var _outptr_8 C.double
	gogsl.InitializeGslFunction(f)
	_result := int32(C.gsl_integration_qawf((*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(a), C.double(epsabs), C.size_t(limit), (*C.gsl_integration_workspace)(unsafe.Pointer(workspace.Ptr())), (*C.gsl_integration_workspace)(unsafe.Pointer(cycleWorkspace.Ptr())), (*C.gsl_integration_qawo_table)(unsafe.Pointer(wf.Ptr())), &_outptr_7, &_outptr_8))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_7)), *(*float64)(unsafe.Pointer(&_outptr_8))
}

func CquadWorkspaceAlloc(n int) *GslIntegrationCquadWorkspace {
	_ref := C.gsl_integration_cquad_workspace_alloc(C.size_t(n))
	_result := &GslIntegrationCquadWorkspace{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func (x *GslIntegrationCquadWorkspace) Dispose() {
	C.gsl_integration_cquad_workspace_free((*C.gsl_integration_cquad_workspace)(unsafe.Pointer(x.Ptr())))
}

func Cquad(f *gogsl.GslFunction, a float64, b float64, epsabs float64, epsrel float64, workspace *GslIntegrationCquadWorkspace) (int32, float64, float64, int) {
	var _outptr_6 C.double
	var _outptr_7 C.double
	var _outptr_8 C.size_t
	gogsl.InitializeGslFunction(f)
	_result := int32(C.gsl_integration_cquad((*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(a), C.double(b), C.double(epsabs), C.double(epsrel), (*C.gsl_integration_cquad_workspace)(unsafe.Pointer(workspace.Ptr())), &_outptr_6, &_outptr_7, &_outptr_8))
	return _result, *(*float64)(unsafe.Pointer(&_outptr_6)), *(*float64)(unsafe.Pointer(&_outptr_7)), *(*int)(unsafe.Pointer(&_outptr_8))
}

func GlfixedTableAlloc(n int) *GslIntegrationGlfixedTable {
	_ref := C.gsl_integration_glfixed_table_alloc(C.size_t(n))
	_result := &GslIntegrationGlfixedTable{}
	gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))
	return _result
}

func Glfixed(f *gogsl.GslFunction, a float64, b float64, t *GslIntegrationGlfixedTable) float64 {
	gogsl.InitializeGslFunction(f)
	return float64(C.gsl_integration_glfixed((*C.gsl_function)(unsafe.Pointer(f.CPtr())), C.double(a), C.double(b), (*C.gsl_integration_glfixed_table)(unsafe.Pointer(t.Ptr()))))
}

func GlfixedPoint(a float64, b float64, i int, xi []float64, wi []float64, t *GslIntegrationGlfixedTable) int32 {
	_slice_header_3 := (*reflect.SliceHeader)(unsafe.Pointer(&xi))
	_slice_header_4 := (*reflect.SliceHeader)(unsafe.Pointer(&wi))
	return int32(C.gsl_integration_glfixed_point(C.double(a), C.double(b), C.size_t(i), (*C.double)(unsafe.Pointer(_slice_header_3.Data)), (*C.double)(unsafe.Pointer(_slice_header_4.Data)), (*C.gsl_integration_glfixed_table)(unsafe.Pointer(t.Ptr()))))
}

func (x *GslIntegrationGlfixedTable) Dispose() {
	C.gsl_integration_glfixed_table_free((*C.gsl_integration_glfixed_table)(unsafe.Pointer(x.Ptr())))
}
