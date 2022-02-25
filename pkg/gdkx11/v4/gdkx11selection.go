// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/x11/gdkx.h>
import "C"

// X11FreeCompoundText frees the data returned from
// gdk_x11_display_string_to_compound_text().
//
// The function takes the following parameters:
//
//    - ctext: pointer stored in ctext from a call to
//      gdk_x11_display_string_to_compound_text().
//
func X11FreeCompoundText(ctext *byte) {
	var _arg1 *C.guchar // out

	_arg1 = (*C.guchar)(unsafe.Pointer(ctext))

	C.gdk_x11_free_compound_text(_arg1)
	runtime.KeepAlive(ctext)
}

// StringToCompoundText: convert a string from the encoding of the current
// locale into a form suitable for storing in a window property.
//
// The function takes the following parameters:
//
//    - str: nul-terminated string.
//
// The function returns the following values:
//
//    - encoding: location to store the encoding (to be used as the type for the
//      property).
//    - format: location to store the format of the property.
//    - ctext: location to store newly allocated data for the property.
//    - gint: 0 upon success, non-zero upon failure.
//
func (display *X11Display) StringToCompoundText(str string) (encoding string, format int, ctext []byte, gint int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 *C.char       // in
	var _arg3 C.int         // in
	var _arg4 *C.guchar     // in
	var _arg5 C.int         // in
	var _cret C.int         // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_x11_display_string_to_compound_text(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)
	runtime.KeepAlive(display)
	runtime.KeepAlive(str)

	var _encoding string // out
	var _format int      // out
	var _ctext []byte    // out
	var _gint int        // out

	_encoding = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	_format = int(_arg3)
	defer C.free(unsafe.Pointer(_arg4))
	_ctext = make([]byte, _arg5)
	copy(_ctext, unsafe.Slice((*byte)(unsafe.Pointer(_arg4)), _arg5))
	_gint = int(_cret)

	return _encoding, _format, _ctext, _gint
}

// UTF8ToCompoundText converts from UTF-8 to compound text.
//
// The function takes the following parameters:
//
//    - str: UTF-8 string.
//
// The function returns the following values:
//
//    - encoding: location to store resulting encoding.
//    - format: location to store format of the result.
//    - ctext: location to store the data of the result.
//    - ok: TRUE if the conversion succeeded, otherwise FALSE.
//
func (display *X11Display) UTF8ToCompoundText(str string) (string, int, []byte, bool) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 *C.char       // in
	var _arg3 C.int         // in
	var _arg4 *C.guchar     // in
	var _arg5 C.int         // in
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_x11_display_utf8_to_compound_text(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)
	runtime.KeepAlive(display)
	runtime.KeepAlive(str)

	var _encoding string // out
	var _format int      // out
	var _ctext []byte    // out
	var _ok bool         // out

	_encoding = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	_format = int(_arg3)
	defer C.free(unsafe.Pointer(_arg4))
	_ctext = make([]byte, _arg5)
	copy(_ctext, unsafe.Slice((*byte)(unsafe.Pointer(_arg4)), _arg5))
	if _cret != 0 {
		_ok = true
	}

	return _encoding, _format, _ctext, _ok
}
