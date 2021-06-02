// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <gdk/gdk.h>
import "C"

// KeyvalConvertCase obtains the upper- and lower-case versions of the keyval
// @symbol.
//
// Examples of keyvals are K_KEY_a, K_KEY_Enter, K_KEY_F1, etc.
func KeyvalConvertCase(symbol uint) (lower uint, upper uint) {
	var arg1 C.guint
	var arg2 *C.guint // out
	var arg3 *C.guint // out

	arg1 = C.guint(symbol)

	C.gdk_keyval_convert_case(arg1, &arg2, &arg3)

	var ret0 uint
	var ret1 uint

	ret0 = uint(arg2)

	ret1 = uint(arg3)

	return ret0, ret1
}

// KeyvalFromName converts a key name to a key value.
//
// The names are the same as those in the `gdk/gdkkeysyms.h` header file but
// without the leading “GDK_KEY_”.
func KeyvalFromName(keyvalName string) uint {
	var arg1 *C.char

	arg1 = (*C.gchar)(C.CString(keyvalName))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.gdk_keyval_from_name(arg1)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// KeyvalIsLower returns true if the given key value is in lower case.
func KeyvalIsLower(keyval uint) bool {
	var arg1 C.guint

	arg1 = C.guint(keyval)

	ret := C.gdk_keyval_is_lower(arg1)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// KeyvalIsUpper returns true if the given key value is in upper case.
func KeyvalIsUpper(keyval uint) bool {
	var arg1 C.guint

	arg1 = C.guint(keyval)

	ret := C.gdk_keyval_is_upper(arg1)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// KeyvalName converts a key value into a symbolic name.
//
// The names are the same as those in the `gdk/gdkkeysyms.h` header file but
// without the leading “GDK_KEY_”.
func KeyvalName(keyval uint) string {
	var arg1 C.guint

	arg1 = C.guint(keyval)

	ret := C.gdk_keyval_name(arg1)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// KeyvalToLower converts a key value to lower case, if applicable.
func KeyvalToLower(keyval uint) uint {
	var arg1 C.guint

	arg1 = C.guint(keyval)

	ret := C.gdk_keyval_to_lower(arg1)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// KeyvalToUnicode: convert from a GDK key symbol to the corresponding Unicode
// character.
//
// Note that the conversion does not take the current locale into consideration,
// which might be expected for particular keyvals, such as GDK_KEY_KP_Decimal.
func KeyvalToUnicode(keyval uint) uint32 {
	var arg1 C.guint

	arg1 = C.guint(keyval)

	ret := C.gdk_keyval_to_unicode(arg1)

	var ret0 uint32

	ret0 = uint32(ret)

	return ret0
}

// KeyvalToUpper converts a key value to upper case, if applicable.
func KeyvalToUpper(keyval uint) uint {
	var arg1 C.guint

	arg1 = C.guint(keyval)

	ret := C.gdk_keyval_to_upper(arg1)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// UnicodeToKeyval: convert from a Unicode character to a key symbol.
func UnicodeToKeyval(wc uint32) uint {
	var arg1 C.guint32

	arg1 = C.guint32(wc)

	ret := C.gdk_unicode_to_keyval(arg1)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}
