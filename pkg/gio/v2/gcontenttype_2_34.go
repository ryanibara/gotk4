// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// ContentTypeGetGenericIconName gets the generic icon name for a content type.
//
// See the shared-mime-info
// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
// specification for more on the generic icon name.
//
// The function takes the following parameters:
//
//    - typ: content type string.
//
// The function returns the following values:
//
//    - utf8 (optional): registered generic icon name for the given type, or NULL
//      if unknown. Free with g_free().
//
func ContentTypeGetGenericIconName(typ string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_content_type_get_generic_icon_name(_arg1)
	runtime.KeepAlive(typ)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// ContentTypeGetSymbolicIcon gets the symbolic icon for a content type.
//
// The function takes the following parameters:
//
//    - typ: content type string.
//
// The function returns the following values:
//
//    - icon: symbolic #GIcon corresponding to the content type. Free the
//      returned object with g_object_unref().
//
func ContentTypeGetSymbolicIcon(typ string) *Icon {
	var _arg1 *C.gchar // out
	var _cret *C.GIcon // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_content_type_get_symbolic_icon(_arg1)
	runtime.KeepAlive(typ)

	var _icon *Icon // out

	_icon = wrapIcon(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _icon
}