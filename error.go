package gogsl

/*
	#include <gsl/gsl_errno.h>
	#include <malloc.h>

	static void *GOGSL_ERROR_HANDLER;

	void *get_go_gsl_error_handler_ptr() {
		return GOGSL_ERROR_HANDLER;
	}

	void set_go_gsl_error_handler_ptr(void *ptr) {
		GOGSL_ERROR_HANDLER = ptr;
	}

	void go_gsl_error_handler_proxy (const char * reason, const char * file, int line, int gsl_errno) {
		invokeGoGslErrorHandler(reason, file, line, gsl_errno);
	}

	void *get_go_gsl_error_handler_proxy_addr() {
		return (void*)go_gsl_error_handler_proxy;
	}

	void call_gsl_error_handler(gsl_error_handler_t *handler, const char * reason, const char * file, int line, int gsl_errno) {
		handler(reason, file, line, gsl_errno);
	}

*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

type GslError int

const (
	GSL_SUCCESS  GslError = 0
	GSL_FAILURE  GslError = -1
	GSL_CONTINUE GslError = -2 /* iteration has not converged */
	GSL_EDOM     GslError = 1  /* input domain error, e.g sqrt(-1) */
	GSL_ERANGE   GslError = 2  /* output range error, e.g. exp(1e100) */
	GSL_EFAULT   GslError = 3  /* invalid poGslErrorer */
	GSL_EINVAL   GslError = 4  /* invalid argument supplied by user */
	GSL_EFAILED  GslError = 5  /* generic failure */
	GSL_EFACTOR  GslError = 6  /* factorization failed */
	GSL_ESANITY  GslError = 7  /* sanity check failed - shouldn't happen */
	GSL_ENOMEM   GslError = 8  /* malloc failed */
	GSL_EBADFUNC GslError = 9  /* problem with user-supplied function */
	GSL_ERUNAWAY GslError = 10 /* iterative process is out of control */
	GSL_EMAXITER GslError = 11 /* exceeded max number of iterations */
	GSL_EZERODIV GslError = 12 /* tried to divide by zero */
	GSL_EBADTOL  GslError = 13 /* user specified an invalid tolerance */
	GSL_ETOL     GslError = 14 /* failed to reach the specified tolerance */
	GSL_EUNDRFLW GslError = 15 /* underflow */
	GSL_EOVRFLW  GslError = 16 /* overflow  */
	GSL_ELOSS    GslError = 17 /* loss of accuracy */
	GSL_EROUND   GslError = 18 /* failed because of roundoff error */
	GSL_EBADLEN  GslError = 19 /* matrix, vector lengths are not conformant */
	GSL_ENOTSQR  GslError = 20 /* matrix not square */
	GSL_ESING    GslError = 21 /* apparent singularity detected */
	GSL_EDIVERGE GslError = 22 /* GslErroregral or series is divergent */
	GSL_EUNSUP   GslError = 23 /* requested feature is not supported by the hardware */
	GSL_EUNIMPL  GslError = 24 /* requested feature not (yet) implemented */
	GSL_ECACHE   GslError = 25 /* cache limit exceeded */
	GSL_ETABLE   GslError = 26 /* table limit exceeded */
	GSL_ENOPROG  GslError = 27 /* iteration is not making progress towards solution */
	GSL_ENOPROGJ GslError = 28 /* jacobian evaluations are not improving the solution */
	GSL_ETOLF    GslError = 29 /* cannot reach the specified tolerance in F */
	GSL_ETOLX    GslError = 30 /* cannot reach the specified tolerance in X */
	GSL_ETOLG    GslError = 31 /* cannot reach the specified tolerance in gradient */
	GSL_EOF      GslError = 32 /* end of file */
)

type GslErrorHandlerType func(reason string, file string, line int, gslError GslError)

type GslErrorHandler struct {
	Handler      GslErrorHandlerType
	cSideHandler bool
	cptr         uintptr
}

// const char * gsl_strerror (const GslError gsl_errno)
func (ge GslError) String() string {
	return C.GoString(C.gsl_strerror(C.int(ge)))
}

var global_GSL_HANDLER_HELD_REFERENCE *GslErrorHandler

// gsl_error_handler_t * gsl_set_error_handler (gsl_error_handler_t * new_handler)
func SetErrorHandler(handler *GslErrorHandler) *GslErrorHandler {
	var prevPtr uintptr
	var oldHandler uintptr
	if handler.cptr == 0 {
		if handler.cSideHandler {
			panic("nil gsl_error_handler_t pointer in reused native handler record")
		}
		handler.cptr = uintptr(C.go_gsl_error_handler_proxy)
		prevPtr = uintptr(C.get_go_gsl_error_handler_ptr())
		C.set_go_gsl_error_handler_ptr(unsafe.Pointer(handler))
	}
	if handler.cSideHandler {
		// Install the original handler and clear the Go handler structure global.
		prevPtr = uintptr(C.get_go_gsl_error_handler_ptr())
		C.set_go_gsl_error_handler_ptr(unsafe.Pointer(nil))
	} else {
		// Install the original handler and set the Go handler global to the given
		// go side structure - we also hold a reference to keep the Go GC from deallocating it.
		global_GSL_HANDLER_HELD_REFERENCE = handler
		prevPtr = uintptr(C.get_go_gsl_error_handler_ptr())
		C.set_go_gsl_error_handler_ptr(unsafe.Pointer(handler))
	}
	// Set the new handler active.
	oldHandler = uintptr(unsafe.Pointer(C.gsl_set_error_handler((*C.gsl_error_handler_t)(C.get_go_gsl_error_handler_proxy_addr()))))
	if oldHandler == 0 {
		// No previous handler installed.
		return nil
	}
	if prevPtr != 0 {
		// The previous handler was a Go handler, reuse the returned structure.
		return (*GslErrorHandler)(unsafe.Pointer(prevPtr))
	} else {
		// The previous handler was a C handler, create a new Go structure for it via closure.
		return &GslErrorHandler{
			cptr:         oldHandler,
			cSideHandler: true,
			Handler: func(reason string, file string, line int, gslError GslError) {
				cReason := C.CString(reason)
				cFile := C.CString(file)
				C.call_gsl_error_handler((*C.gsl_error_handler_t)(unsafe.Pointer(oldHandler)),
					cReason, cFile, C.int(line), C.int(gslError))
				C.free(unsafe.Pointer(cReason))
				C.free(unsafe.Pointer(cFile))
			},
		}
	}
}

// gsl_error_handler_t * gsl_set_error_handler_off ()
func SetErrorHandlerOff() *GslErrorHandler {
	oldHandler := uintptr(unsafe.Pointer(C.gsl_set_error_handler_off()))
	prevPtr := uintptr(C.get_go_gsl_error_handler_ptr())
	C.set_go_gsl_error_handler_ptr(unsafe.Pointer(nil))
	if oldHandler == 0 {
		// No previous handler installed.
		return nil
	}
	if prevPtr != 0 {
		// The previous handler was a Go handler, reuse the returned structure.
		return (*GslErrorHandler)(unsafe.Pointer(prevPtr))
	} else {
		// The previous handler was a C handler, create a new Go structure for it via closure.
		return &GslErrorHandler{
			cptr:         oldHandler,
			cSideHandler: true,
			Handler: func(reason string, file string, line int, gslError GslError) {
				cReason := C.CString(reason)
				cFile := C.CString(file)
				C.call_gsl_error_handler((*C.gsl_error_handler_t)(unsafe.Pointer(oldHandler)),
					cReason, cFile, C.int(line), C.int(gslError))
				C.free(unsafe.Pointer(cReason))
				C.free(unsafe.Pointer(cFile))
			},
		}
	}
}

func Error(reason string, gslErrno GslError) {
	cReason := C.CString(reason)
	pc, file, line, _ := runtime.Caller(1)
	cFile := C.CString(fmt.Sprintf("(0x%8.8x) %s", pc, file))
	C.gsl_error(cReason, cFile, C.int(line), C.int(gslErrno))
	C.free(unsafe.Pointer(cReason))
	C.free(unsafe.Pointer(cFile))
}
