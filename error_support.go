package gogsl

/*
	#include <gsl/gsl_errno.h>
	
	extern void *get_go_gsl_error_handler_ptr();
	
*/
import "C"

import "unsafe"

//export invokeGoGslErrorHandler
func invokeGoGslErrorHandler(cReason uintptr, cFile uintptr, line int, gsl_errno int) {
	reason := C.GoString((*C.char)(unsafe.Pointer(cReason)))
	file := C.GoString((*C.char)(unsafe.Pointer(cFile)))
	handler := (*GslErrorHandler)(unsafe.Pointer(C.get_go_gsl_error_handler_ptr()))
	if handler != nil {
		handler.Handler(reason, file, line, GslError(gsl_errno))
	}
	
}