package ntuple

/*

	#cgo pkg-config: --define-variable=prefix=. gsl
	#include <gsl/gsl_ntuple.h>
	#include <gsl/gsl_histogram.h>
	#include <stdio.h>

	extern int gslNtupleSelectionFunctionCaller(void *x, void *cFunParamPtr);
	extern double gslNtupleValueFunctionCaller(void *x, void *cFunParamPtr);

	int _gsl_ntuple_selection_function_proxy(void *ntuple_data, void *params) {
		return gslNtupleSelectionFunctionCaller(ntuple_data, params);
	}

	double _gsl_ntuple_value_function_proxy(void *ntuple_data, void *params) {
		double val = gslNtupleValueFunctionCaller(ntuple_data, params);
		//printf("VALUE VALUE : %f\n", val);
		return val;
	}

	void _init_proxy_gsl_ntuple_select_function(void *goStruct, gsl_ntuple_select_fn *fn) {
		fn->function = _gsl_ntuple_selection_function_proxy;
		fn->params = goStruct;
	}

	void _init_proxy_gsl_ntuple_value_function(void *goStruct, gsl_ntuple_value_fn *fn) {
		fn->function = _gsl_ntuple_value_function_proxy;
		fn->params = goStruct;
	}

	size_t get_gsl_ntuple_selection_function_struct_size() {
		return sizeof(gsl_ntuple_select_fn);
	}

	size_t get_gsl_ntuple_value_function_struct_size() {
		return sizeof(gsl_ntuple_value_fn);
	}

*/
import "C"

import (
	"github.com/dtromb/gogsl"
	"github.com/dtromb/gogsl/histogram"
	"reflect"
	"unsafe"
)

var GSL_NTUPLE_SELECTION_FUNCTION_STRUCT_SIZE int32 = int32(C.get_gsl_ntuple_selection_function_struct_size())
var GSL_NTUPLE_VALUE_FUNCTION_STRUCT_SIZE int32 = int32(C.get_gsl_ntuple_value_function_struct_size())

type GslNtuple struct {
	gogsl.GslReference
	refLen      int
	elementSize int
	ref         interface{}
	baseType    reflect.Type
}

type OpenOperation int

const (
	OPEN OpenOperation = iota
	CREATE
)

//gsl_ntuple * gsl_ntuple_create (char * filename, void * ntuple_data, size_t size)
func CreateOrOpen(filename string, data interface{}, op OpenOperation) *GslNtuple {

	// Check the data type and make sure it is a slice of some concrete type.
	dataType := reflect.TypeOf(data)
	if dataType.Kind() != reflect.Slice {
		gogsl.Error("ntuple.Create() data argument must be a slice", gogsl.GSL_EINVAL)
		return nil
	}
	dataValue := reflect.ValueOf(data)
	baseType := dataType.Elem()
	dataSize := int(baseType.Size()) * dataValue.Len()
	if !dataValue.CanAddr() {
		x := reflect.New(dataValue.Type()).Elem()
		x.Set(dataValue)
		dataValue = x
	}
	srcHdr := (*reflect.SliceHeader)(unsafe.Pointer(dataValue.UnsafeAddr()))
	cFile := C.CString(filename)
	var tuplePtr *C.gsl_ntuple
	switch op {
	case OPEN:
		{
			tuplePtr = C.gsl_ntuple_open(cFile, unsafe.Pointer(srcHdr.Data), C.size_t(dataSize))
		}
	case CREATE:
		{
			tuplePtr = C.gsl_ntuple_create(cFile, unsafe.Pointer(srcHdr.Data), C.size_t(dataSize))
		}
	default:
		{
			gogsl.Error("unknown open operation", gogsl.GSL_EINVAL)
			return nil
		}
	}
	result := &GslNtuple{baseType: baseType, refLen: dataValue.Len(), elementSize: int(baseType.Size()), ref: data}
	gogsl.InitializeGslReference(result, uintptr(unsafe.Pointer(tuplePtr)))
	C.free(unsafe.Pointer(cFile))
	return result
}

func Open(filename string, data interface{}) *GslNtuple {
	return CreateOrOpen(filename, data, OPEN)
}

func Create(filename string, data interface{}) *GslNtuple {
	return CreateOrOpen(filename, data, CREATE)
}

func Close(nt *GslNtuple) int {
	result := int(C.gsl_ntuple_close((*C.gsl_ntuple)(unsafe.Pointer(nt.Ptr()))))
	gogsl.GslReferenceForceClear(&nt.GslReference)
	return result
}

func Read(nt *GslNtuple) int {
	return int(C.gsl_ntuple_read((*C.gsl_ntuple)(unsafe.Pointer(nt.Ptr()))))
}

func Write(nt *GslNtuple) int {
	return int(C.gsl_ntuple_write((*C.gsl_ntuple)(unsafe.Pointer(nt.Ptr()))))
}

func Bookdata(nt *GslNtuple) int {
	return int(C.gsl_ntuple_bookdata((*C.gsl_ntuple)(unsafe.Pointer(nt.Ptr()))))
}

func (nt *GslNtuple) Dispose() {
	if nt.Ptr() == 0 {
		return
	}
	C.gsl_ntuple_close((*C.gsl_ntuple)(unsafe.Pointer(nt.Ptr())))
}

func Project(h *histogram.GslHistogram, ntuple *GslNtuple,
	fValue *GslNtupleValueFunction,
	fSelection *GslNtupleSelectionFunction) int {
	fSelection.lock.Lock()
	defer fSelection.lock.Unlock()
	fValue.lock.Lock()
	defer fValue.lock.Unlock() // =(
	fSelection.ntuple = ntuple
	fValue.ntuple = ntuple
	result := int(C.gsl_ntuple_project((*C.gsl_histogram)(unsafe.Pointer(h.Ptr())),
		(*C.gsl_ntuple)(unsafe.Pointer(ntuple.Ptr())),
		(*C.gsl_ntuple_value_fn)(unsafe.Pointer(fValue.CPtr())),
		(*C.gsl_ntuple_select_fn)(unsafe.Pointer(fSelection.CPtr()))))
	fSelection.ntuple = nil
	fValue.ntuple = nil
	return result
}

func InitializeGslNtupleSelectionFunction(ptr *GslNtupleSelectionFunction) {
	ptr.cStructData = make([]byte, 0, GSL_NTUPLE_SELECTION_FUNCTION_STRUCT_SIZE)
	cPtr := ((*reflect.SliceHeader)(unsafe.Pointer(&ptr.cStructData))).Data
	C._init_proxy_gsl_ntuple_select_function(unsafe.Pointer(ptr), (*C.gsl_ntuple_select_fn)(unsafe.Pointer(cPtr)))
}

func InitializeGslNtupleValueFunction(ptr *GslNtupleValueFunction) {
	ptr.cStructData = make([]byte, 0, GSL_NTUPLE_VALUE_FUNCTION_STRUCT_SIZE)
	cPtr := ((*reflect.SliceHeader)(unsafe.Pointer(&ptr.cStructData))).Data
	C._init_proxy_gsl_ntuple_value_function(unsafe.Pointer(ptr), (*C.gsl_ntuple_value_fn)(unsafe.Pointer(cPtr)))
}

func (gf *GslNtupleSelectionFunction) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&gf.cStructData))).Data
}

func (gf *GslNtupleValueFunction) CPtr() uintptr {
	return ((*reflect.SliceHeader)(unsafe.Pointer(&gf.cStructData))).Data
}
