// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// KeyvalConvertCase obtains the upper- and lower-case versions of the keyval
// symbol.
//
// Examples of keyvals are K_KEY_a, K_KEY_Enter, K_KEY_F1, etc.
//
// The function takes the following parameters:
//
//    - symbol: keyval.
//
func KeyvalConvertCase(symbol uint) (lower uint, upper uint) {
	var _arg1 C.guint // out
	var _arg2 C.guint // in
	var _arg3 C.guint // in

	_arg1 = C.guint(symbol)

	C.gdk_keyval_convert_case(_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(symbol)

	var _lower uint // out
	var _upper uint // out

	_lower = uint(_arg2)
	_upper = uint(_arg3)

	return _lower, _upper
}

// KeyvalFromName converts a key name to a key value.
//
// The names are the same as those in the gdk/gdkkeysyms.h header file but
// without the leading “GDK_KEY_”.
//
// The function takes the following parameters:
//
//    - keyvalName: key name.
//
func KeyvalFromName(keyvalName string) uint {
	var _arg1 *C.char // out
	var _cret C.guint // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(keyvalName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_keyval_from_name(_arg1)
	runtime.KeepAlive(keyvalName)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// KeyvalIsLower returns TRUE if the given key value is in lower case.
//
// The function takes the following parameters:
//
//    - keyval: key value.
//
func KeyvalIsLower(keyval uint) bool {
	var _arg1 C.guint    // out
	var _cret C.gboolean // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_is_lower(_arg1)
	runtime.KeepAlive(keyval)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// KeyvalIsUpper returns TRUE if the given key value is in upper case.
//
// The function takes the following parameters:
//
//    - keyval: key value.
//
func KeyvalIsUpper(keyval uint) bool {
	var _arg1 C.guint    // out
	var _cret C.gboolean // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_is_upper(_arg1)
	runtime.KeepAlive(keyval)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// KeyvalName converts a key value into a symbolic name.
//
// The names are the same as those in the gdk/gdkkeysyms.h header file but
// without the leading “GDK_KEY_”.
//
// The function takes the following parameters:
//
//    - keyval: key value.
//
func KeyvalName(keyval uint) string {
	var _arg1 C.guint // out
	var _cret *C.char // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_name(_arg1)
	runtime.KeepAlive(keyval)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// KeyvalToLower converts a key value to lower case, if applicable.
//
// The function takes the following parameters:
//
//    - keyval: key value.
//
func KeyvalToLower(keyval uint) uint {
	var _arg1 C.guint // out
	var _cret C.guint // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_to_lower(_arg1)
	runtime.KeepAlive(keyval)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// KeyvalToUnicode: convert from a GDK key symbol to the corresponding Unicode
// character.
//
// Note that the conversion does not take the current locale into consideration,
// which might be expected for particular keyvals, such as GDK_KEY_KP_Decimal.
//
// The function takes the following parameters:
//
//    - keyval: GDK key symbol.
//
func KeyvalToUnicode(keyval uint) uint32 {
	var _arg1 C.guint   // out
	var _cret C.guint32 // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_to_unicode(_arg1)
	runtime.KeepAlive(keyval)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// KeyvalToUpper converts a key value to upper case, if applicable.
//
// The function takes the following parameters:
//
//    - keyval: key value.
//
func KeyvalToUpper(keyval uint) uint {
	var _arg1 C.guint // out
	var _cret C.guint // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_to_upper(_arg1)
	runtime.KeepAlive(keyval)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// UnicodeToKeyval: convert from a Unicode character to a key symbol.
//
// The function takes the following parameters:
//
//    - wc: unicode character.
//
func UnicodeToKeyval(wc uint32) uint {
	var _arg1 C.guint32 // out
	var _cret C.guint   // in

	_arg1 = C.guint32(wc)

	_cret = C.gdk_unicode_to_keyval(_arg1)
	runtime.KeepAlive(wc)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}
