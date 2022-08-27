// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// SCALE: scale between dimensions used for Pango distances and device units.
//
// The definition of device units is dependent on the output device; it will
// typically be pixels for a screen, and points for a printer. PANGO_SCALE is
// currently 1024, but this may be changed in the future.
//
// When setting font sizes, device units are always considered to be points (as
// in "12 point font"), rather than pixels.
const SCALE = 1024

// Glyph: PangoGlyph represents a single glyph in the output form of a string.
type Glyph = uint32

// Rectangle: PangoRectangle structure represents a rectangle.
//
// PangoRectangle is frequently used to represent the logical or ink extents of
// a single glyph or section of text. (See, for instance,
// pango.Font.GetGlyphExtents().)
//
// An instance of this type is always passed by reference.
type Rectangle struct {
	*rectangle
}

// rectangle is the struct that's finalized.
type rectangle struct {
	native *C.PangoRectangle
}

// NewRectangle creates a new Rectangle instance from the given
// fields. Beware that this function allocates on the Go heap; be careful
// when using it!
func NewRectangle(x, y, width, height int) Rectangle {
	var f0 C.int // out
	f0 = C.int(x)
	var f1 C.int // out
	f1 = C.int(y)
	var f2 C.int // out
	f2 = C.int(width)
	var f3 C.int // out
	f3 = C.int(height)

	v := C.PangoRectangle{
		x:      f0,
		y:      f1,
		width:  f2,
		height: f3,
	}

	return *(*Rectangle)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

// X coordinate of the left side of the rectangle.
func (r *Rectangle) X() int {
	valptr := &r.native.x
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Y coordinate of the the top side of the rectangle.
func (r *Rectangle) Y() int {
	valptr := &r.native.y
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Width: width of the rectangle.
func (r *Rectangle) Width() int {
	valptr := &r.native.width
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Height: height of the rectangle.
func (r *Rectangle) Height() int {
	valptr := &r.native.height
	var _v int // out
	_v = int(*valptr)
	return _v
}

// X coordinate of the left side of the rectangle.
func (r *Rectangle) SetX(x int) {
	valptr := &r.native.x
	*valptr = C.int(x)
}

// Y coordinate of the the top side of the rectangle.
func (r *Rectangle) SetY(y int) {
	valptr := &r.native.y
	*valptr = C.int(y)
}

// Width: width of the rectangle.
func (r *Rectangle) SetWidth(width int) {
	valptr := &r.native.width
	*valptr = C.int(width)
}

// Height: height of the rectangle.
func (r *Rectangle) SetHeight(height int) {
	valptr := &r.native.height
	*valptr = C.int(height)
}
