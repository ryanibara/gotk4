// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// IOErrorFromErrno converts errno.h error codes into GIO error codes. The
// fallback value G_IO_ERROR_FAILED is returned for error codes not currently
// handled (but note that future GLib releases may return a more specific value
// instead).
//
// As errno is global and may be modified by intermediate function calls, you
// should save its value as soon as the call which sets it.
//
// The function takes the following parameters:
//
//    - errNo: error number as defined in errno.h.
//
// The function returns the following values:
//
//    - ioErrorEnum value for the given errno.h error number.
//
func IOErrorFromErrno(errNo int) IOErrorEnum {
	var _arg1 C.gint         // out
	var _cret C.GIOErrorEnum // in

	_arg1 = C.gint(errNo)

	_cret = C.g_io_error_from_errno(_arg1)
	runtime.KeepAlive(errNo)

	var _ioErrorEnum IOErrorEnum // out

	_ioErrorEnum = IOErrorEnum(_cret)

	return _ioErrorEnum
}
