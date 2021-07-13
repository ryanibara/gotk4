// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gsk/gsk.h>
import "C"

// TransformParse parses the given string into a transform and puts it in
// out_transform.
//
// Strings printed via gsk.Transform.ToString() can be read in again
// successfully using this function.
//
// If string does not describe a valid transform, FALSE is returned and NULL is
// put in out_transform.
func TransformParse(_string string) (*Transform, bool) {
	var _arg1 *C.char // out
	var _outTransform *Transform
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(_string)))

	_cret = C.gsk_transform_parse(_arg1, (**C.GskTransform)(unsafe.Pointer(&_outTransform)))

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _outTransform, _ok
}
