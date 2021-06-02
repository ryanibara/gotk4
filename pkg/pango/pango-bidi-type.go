// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <pango/pango.h>
import "C"

// BidiTypeForUnichar determines the bidirectional type of a character.
//
// The bidirectional type is specified in the Unicode Character Database.
//
// A simplified version of this function is available as
// [func@unichar_direction].
func BidiTypeForUnichar(ch uint32) BidiType {
	var arg1 C.gunichar

	arg1 = C.gunichar(ch)

	ret := C.pango_bidi_type_for_unichar(arg1)

	var ret0 BidiType

	ret0 = BidiType(ret)

	return ret0
}

// FindBaseDir searches a string the first character that has a strong
// direction, according to the Unicode bidirectional algorithm.
func FindBaseDir(text string, length int) Direction {
	var arg1 *C.gchar
	var arg2 C.gint

	arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gint(length)

	ret := C.pango_find_base_dir(arg1, arg2)

	var ret0 Direction

	ret0 = Direction(ret)

	return ret0
}

// GetMirrorChar returns the mirrored character of a Unicode character.
//
// Mirror characters are determined by the Unicode mirrored property.
//
// Use g_unichar_get_mirror_char() instead; the docs for that function provide
// full details.
func GetMirrorChar(ch uint32, mirroredCh uint32) bool {
	var arg1 C.gunichar
	var arg2 *C.gunichar

	arg1 = C.gunichar(ch)
	arg2 = (*C.gunichar)(mirroredCh)

	ret := C.pango_get_mirror_char(arg1, arg2)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// UnicharDirection determines the inherent direction of a character.
//
// The inherent direction is either PANGO_DIRECTION_LTR, PANGO_DIRECTION_RTL, or
// PANGO_DIRECTION_NEUTRAL.
//
// This function is useful to categorize characters into left-to-right letters,
// right-to-left letters, and everything else. If full Unicode bidirectional
// type of a character is needed, [type_func@Pango.BidiType.for_unichar] can be
// used instead.
func UnicharDirection(ch uint32) Direction {
	var arg1 C.gunichar

	arg1 = C.gunichar(ch)

	ret := C.pango_unichar_direction(arg1)

	var ret0 Direction

	ret0 = Direction(ret)

	return ret0
}
