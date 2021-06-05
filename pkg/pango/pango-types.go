// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
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
	var arg1 *C.PangoRectangle
	var arg2 *C.PangoRectangle

	arg1 = (*C.PangoRectangle)(unsafe.Pointer(inclusive.Native()))
	arg2 = (*C.PangoRectangle)(unsafe.Pointer(nearest.Native()))

	C.pango_extents_to_pixels(inclusive, nearest)
}

// UnitsFromDouble converts a floating-point number to Pango units.
//
// The conversion is done by multiplying @d by PANGO_SCALE and rounding the
// result to nearest integer.
func UnitsFromDouble(d float64) int {
	var arg1 C.double

	arg1 = C.double(d)

	var cret C.int
	var ret1 int

	cret = C.pango_units_from_double(d)

	ret1 = C.int(cret)

	return ret1
}

// UnitsToDouble converts a number in Pango units to floating-point.
//
// The conversion is done by dividing @i by PANGO_SCALE.
func UnitsToDouble(i int) float64 {
	var arg1 C.int

	arg1 = C.int(i)

	var cret C.double
	var ret1 float64

	cret = C.pango_units_to_double(i)

	ret1 = C.double(cret)

	return ret1
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

func marshalRectangle(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRectangle(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *Rectangle) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// X gets the field inside the struct.
func (r *Rectangle) X() int {
	v = C.int(r.native.x)
}

// Y gets the field inside the struct.
func (r *Rectangle) Y() int {
	v = C.int(r.native.y)
}

// Width gets the field inside the struct.
func (r *Rectangle) Width() int {
	v = C.int(r.native.width)
}

// Height gets the field inside the struct.
func (r *Rectangle) Height() int {
	v = C.int(r.native.height)
}
