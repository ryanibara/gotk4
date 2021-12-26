// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// UUIDStringIsValid parses the string str and verify if it is a UUID.
//
// The function accepts the following syntax:
//
// - simple forms (e.g. f81d4fae-7dec-11d0-a765-00a0c91e6bf6)
//
// Note that hyphens are required within the UUID string itself, as per the
// aforementioned RFC.
//
// The function takes the following parameters:
//
//    - str: string representing a UUID.
//
// The function returns the following values:
//
//    - ok: TRUE if str is a valid UUID, FALSE otherwise.
//
func UUIDStringIsValid(str string) bool {
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_uuid_string_is_valid(_arg1)
	runtime.KeepAlive(str)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UUIDStringRandom generates a random UUID (RFC 4122 version 4) as a string. It
// has the same randomness guarantees as #GRand, so must not be used for
// cryptographic purposes such as key generation, nonces, salts or one-time
// pads.
//
// The function returns the following values:
//
//    - utf8: string that should be freed with g_free().
//
func UUIDStringRandom() string {
	var _cret *C.gchar // in

	_cret = C.g_uuid_string_random()

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
