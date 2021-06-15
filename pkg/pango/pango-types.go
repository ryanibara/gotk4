// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
import "C"

// Glyph: a `PangoGlyph` represents a single glyph in the output form of a
// string.
type Glyph uint32

// ExtentsToPixels converts extents from Pango units to device units.
//
// The conversion is done by dividing by the PANGO_SCALE factor and performing
// rounding.
//
// The @inclusive rectangle is converted by flooring the x/y coordinates and
// extending width/height, such that the final rectangle completely includes the
// original rectangle.
//
// The @nearest rectangle is converted by rounding the coordinates of the
// rectangle to the nearest device unit (pixel).
//
// The rule to which argument to use is: if you want the resulting device-space
// rectangle to completely contain the original rectangle, pass it in as
// @inclusive. If you want two touching-but-not-overlapping rectangles stay
// touching-but-not-overlapping after rounding to device units, pass them in as
// @nearest.
func ExtentsToPixels(inclusive *Rectangle, nearest *Rectangle) {
	var _arg1 *C.PangoRectangle // out
	var _arg2 *C.PangoRectangle // out

	_arg1 = (*C.PangoRectangle)(unsafe.Pointer(inclusive.Native()))
	_arg2 = (*C.PangoRectangle)(unsafe.Pointer(nearest.Native()))

	C.pango_extents_to_pixels(_arg1, _arg2)
}

// UnitsFromDouble converts a floating-point number to Pango units.
//
// The conversion is done by multiplying @d by PANGO_SCALE and rounding the
// result to nearest integer.
func UnitsFromDouble(d float64) int {
	var _arg1 C.double // out
	var _cret C.int    // in

	_arg1 = (C.double)(d)

	_cret = C.pango_units_from_double(_arg1)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// UnitsToDouble converts a number in Pango units to floating-point.
//
// The conversion is done by dividing @i by PANGO_SCALE.
func UnitsToDouble(i int) float64 {
	var _arg1 C.int    // out
	var _cret C.double // in

	_arg1 = (C.int)(i)

	_cret = C.pango_units_to_double(_arg1)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// Rectangle: the `PangoRectangle` structure represents a rectangle.
//
// `PangoRectangle` is frequently used to represent the logical or ink extents
// of a single glyph or section of text. (See, for instance,
// [method@Pango.Font.get_glyph_extents].)
type Rectangle struct {
	native C.PangoRectangle
}

// WrapRectangle wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRectangle(ptr unsafe.Pointer) *Rectangle {
	if ptr == nil {
		return nil
	}

	return (*Rectangle)(ptr)
}

// Native returns the underlying C source pointer.
func (r *Rectangle) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}
