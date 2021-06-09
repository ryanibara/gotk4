// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

// KeyvalConvertCase obtains the upper- and lower-case versions of the keyval
// @symbol. Examples of keyvals are K_KEY_a, K_KEY_Enter, K_KEY_F1, etc.
func KeyvalConvertCase(symbol uint) (lower uint, upper uint) {
	var _arg1 C.guint

	_arg1 = C.guint(symbol)

	var _arg2 C.guint
	var _arg3 C.guint

	C.gdk_keyval_convert_case(_arg1, &_arg2, &_arg3)

	var _lower uint
	var _upper uint

	_lower = (uint)(_arg2)
	_upper = (uint)(_arg3)

	return _lower, _upper
}

// KeyvalFromName converts a key name to a key value.
//
// The names are the same as those in the `gdk/gdkkeysyms.h` header file but
// without the leading “GDK_KEY_”.
func KeyvalFromName(keyvalName string) uint {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(keyvalName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.guint

	cret = C.gdk_keyval_from_name(_arg1)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// KeyvalIsLower returns true if the given key value is in lower case.
func KeyvalIsLower(keyval uint) bool {
	var _arg1 C.guint

	_arg1 = C.guint(keyval)

	var _cret C.gboolean

	cret = C.gdk_keyval_is_lower(_arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// KeyvalIsUpper returns true if the given key value is in upper case.
func KeyvalIsUpper(keyval uint) bool {
	var _arg1 C.guint

	_arg1 = C.guint(keyval)

	var _cret C.gboolean

	cret = C.gdk_keyval_is_upper(_arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// KeyvalName converts a key value into a symbolic name.
//
// The names are the same as those in the `gdk/gdkkeysyms.h` header file but
// without the leading “GDK_KEY_”.
func KeyvalName(keyval uint) string {
	var _arg1 C.guint

	_arg1 = C.guint(keyval)

	var _cret *C.gchar

	cret = C.gdk_keyval_name(_arg1)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// KeyvalToLower converts a key value to lower case, if applicable.
func KeyvalToLower(keyval uint) uint {
	var _arg1 C.guint

	_arg1 = C.guint(keyval)

	var _cret C.guint

	cret = C.gdk_keyval_to_lower(_arg1)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// KeyvalToUnicode: convert from a GDK key symbol to the corresponding ISO10646
// (Unicode) character.
func KeyvalToUnicode(keyval uint) uint32 {
	var _arg1 C.guint

	_arg1 = C.guint(keyval)

	var _cret C.guint32

	cret = C.gdk_keyval_to_unicode(_arg1)

	var _guint32 uint32

	_guint32 = (uint32)(_cret)

	return _guint32
}

// KeyvalToUpper converts a key value to upper case, if applicable.
func KeyvalToUpper(keyval uint) uint {
	var _arg1 C.guint

	_arg1 = C.guint(keyval)

	var _cret C.guint

	cret = C.gdk_keyval_to_upper(_arg1)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// UnicodeToKeyval: convert from a ISO10646 character to a key symbol.
func UnicodeToKeyval(wc uint32) uint {
	var _arg1 C.guint32

	_arg1 = C.guint32(wc)

	var _cret C.guint

	cret = C.gdk_unicode_to_keyval(_arg1)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// KeymapKey: a KeymapKey is a hardware key that can be mapped to a keyval.
type KeymapKey struct {
	native C.GdkKeymapKey
}

// WrapKeymapKey wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapKeymapKey(ptr unsafe.Pointer) *KeymapKey {
	if ptr == nil {
		return nil
	}

	return (*KeymapKey)(ptr)
}

func marshalKeymapKey(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapKeymapKey(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (k *KeymapKey) Native() unsafe.Pointer {
	return unsafe.Pointer(&k.native)
}

// Keycode gets the field inside the struct.
func (k *KeymapKey) Keycode() uint {
	var v uint
	v = (uint)(k.native.keycode)
	return v
}

// Group gets the field inside the struct.
func (k *KeymapKey) Group() int {
	var v int
	v = (int)(k.native.group)
	return v
}

// Level gets the field inside the struct.
func (k *KeymapKey) Level() int {
	var v int
	v = (int)(k.native.level)
	return v
}