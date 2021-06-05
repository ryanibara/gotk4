// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_rgba_get_type()), F: marshalRGBA},
	})
}

// RGBA: a RGBA is used to represent a (possibly translucent) color, in a way
// that is compatible with cairo’s notion of color.
type RGBA struct {
	native C.GdkRGBA
}

// WrapRGBA wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRGBA(ptr unsafe.Pointer) *RGBA {
	if ptr == nil {
		return nil
	}

	return (*RGBA)(ptr)
}

func marshalRGBA(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRGBA(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RGBA) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Red gets the field inside the struct.
func (r *RGBA) Red() float32 {
	v = C.float(r.native.red)
}

// Green gets the field inside the struct.
func (r *RGBA) Green() float32 {
	v = C.float(r.native.green)
}

// Blue gets the field inside the struct.
func (r *RGBA) Blue() float32 {
	v = C.float(r.native.blue)
}

// Alpha gets the field inside the struct.
func (r *RGBA) Alpha() float32 {
	v = C.float(r.native.alpha)
}

// Copy makes a copy of a RGBA.
//
// The result must be freed through gdk_rgba_free().
func (r *RGBA) Copy() *RGBA {
	var arg0 *C.GdkRGBA

	arg0 = (*C.GdkRGBA)(unsafe.Pointer(r.Native()))

	var cret *C.GdkRGBA
	var ret1 *RGBA

	cret = C.gdk_rgba_copy(arg0)

	ret1 = WrapRGBA(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *RGBA) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Equal compares two RGBA colors.
func (p *RGBA) Equal(p2 RGBA) bool {
	var arg0 C.gpointer
	var arg1 C.gpointer

	arg0 = (C.gpointer)(unsafe.Pointer(p.Native()))
	arg1 = (C.gpointer)(unsafe.Pointer(p2.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gdk_rgba_equal(arg0, p2)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Free frees a RGBA created with gdk_rgba_copy()
func (r *RGBA) Free() {
	var arg0 *C.GdkRGBA

	arg0 = (*C.GdkRGBA)(unsafe.Pointer(r.Native()))

	C.gdk_rgba_free(arg0)
}

// Hash: a hash function suitable for using for a hash table that stores RGBAs.
func (p *RGBA) Hash() uint {
	var arg0 C.gpointer

	arg0 = (C.gpointer)(unsafe.Pointer(p.Native()))

	var cret C.guint
	var ret1 uint

	cret = C.gdk_rgba_hash(arg0)

	ret1 = C.guint(cret)

	return ret1
}

// IsClear checks if an @rgba value is transparent. That is, drawing with the
// value would not produce any change.
func (r *RGBA) IsClear() bool {
	var arg0 *C.GdkRGBA

	arg0 = (*C.GdkRGBA)(unsafe.Pointer(r.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gdk_rgba_is_clear(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// IsOpaque checks if an @rgba value is opaque. That is, drawing with the value
// will not retain any results from previous contents.
func (r *RGBA) IsOpaque() bool {
	var arg0 *C.GdkRGBA

	arg0 = (*C.GdkRGBA)(unsafe.Pointer(r.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gdk_rgba_is_opaque(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Parse parses a textual representation of a color, filling in the @red,
// @green, @blue and @alpha fields of the @rgba RGBA.
//
// The string can be either one of: - A standard name (Taken from the X11
// rgb.txt file). - A hexadecimal value in the form “\#rgb”, “\#rrggbb”,
// “\#rrrgggbbb” or ”\#rrrrggggbbbb” - A hexadecimal value in the form “\#rgba”,
// “\#rrggbbaa”, or ”\#rrrrggggbbbbaaaa” - A RGB color in the form “rgb(r,g,b)”
// (In this case the color will have full opacity) - A RGBA color in the form
// “rgba(r,g,b,a)”
//
// Where “r”, “g”, “b” and “a” are respectively the red, green, blue and alpha
// color values. In the last two cases, “r”, “g”, and “b” are either integers in
// the range 0 to 255 or percentage values in the range 0% to 100%, and a is a
// floating point value in the range 0 to 1.
func (r *RGBA) Parse(spec string) bool {
	var arg0 *C.GdkRGBA
	var arg1 *C.char

	arg0 = (*C.GdkRGBA)(unsafe.Pointer(r.Native()))
	arg1 = (*C.char)(C.CString(spec))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ret1 bool

	cret = C.gdk_rgba_parse(arg0, spec)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// String returns a textual specification of @rgba in the form `rgb(r,g,b)` or
// `rgba(r,g,b,a)`, where “r”, “g”, “b” and “a” represent the red, green, blue
// and alpha values respectively. “r”, “g”, and “b” are represented as integers
// in the range 0 to 255, and “a” is represented as a floating point value in
// the range 0 to 1.
//
// These string forms are string forms that are supported by the CSS3 colors
// module, and can be parsed by gdk_rgba_parse().
//
// Note that this string representation may lose some precision, since “r”, “g”
// and “b” are represented as 8-bit integers. If this is a concern, you should
// use a different representation.
func (r *RGBA) String() string {
	var arg0 *C.GdkRGBA

	arg0 = (*C.GdkRGBA)(unsafe.Pointer(r.Native()))

	var cret *C.char
	var ret1 string

	cret = C.gdk_rgba_to_string(arg0)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}
