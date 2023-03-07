// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// AttrTypeGetName fetches the attribute type name.
//
// The attribute type name is the string passed in when registering the type
// using attr_type_register.
//
// The returned value is an interned string (see g_intern_string() for what that
// means) that should not be modified or freed.
//
// The function takes the following parameters:
//
//    - typ: attribute type ID to fetch the name for.
//
// The function returns the following values:
//
//    - utf8 (optional): type ID name (which may be NULL), or NULL if type is a
//      built-in Pango attribute type or invalid.
//
func AttrTypeGetName(typ AttrType) string {
	var _arg1 C.PangoAttrType // out
	var _cret *C.char         // in

	_arg1 = C.PangoAttrType(typ)

	_cret = C.pango_attr_type_get_name(_arg1)
	runtime.KeepAlive(typ)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}