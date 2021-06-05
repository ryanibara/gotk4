// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_color_get_type()), F: marshalColor},
	})
}

// ColorParse parses a textual specification of a color and fill in the @red,
// @green, and @blue fields of a Color.
//
// The string can either one of a large set of standard names (taken from the
// X11 `rgb.txt` file), or it can be a hexadecimal value in the form “\#rgb”
// “\#rrggbb”, “\#rrrgggbbb” or “\#rrrrggggbbbb” where “r”, “g” and “b” are hex
// digits of the red, green, and blue components of the color, respectively.
// (White in the four forms is “\#fff”, “\#ffffff”, “\#fffffffff” and
// “\#ffffffffffff”).
func ColorParse(spec string) (color Color, ok bool) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(spec))
	defer C.free(unsafe.Pointer(arg1))

	var arg2 C.GdkColor
	var ret2 *Color
	var cret C.gboolean
	var ret2 bool

	cret = C.gdk_color_parse(spec, &arg2)

	*ret2 = WrapColor(unsafe.Pointer(arg2))
	ret2 = C.bool(cret) != C.false

	return ret2, ret2
}

// Color: a Color is used to describe a color, similar to the XColor struct used
// in the X11 drawing API.
type Color struct {
	native C.GdkColor
}

// WrapColor wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapColor(ptr unsafe.Pointer) *Color {
	if ptr == nil {
		return nil
	}

	return (*Color)(ptr)
}

func marshalColor(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapColor(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *Color) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// Pixel gets the field inside the struct.
func (c *Color) Pixel() uint32 {
	v = C.guint32(c.native.pixel)
}

// Red gets the field inside the struct.
func (c *Color) Red() uint16 {
	v = C.guint16(c.native.red)
}

// Green gets the field inside the struct.
func (c *Color) Green() uint16 {
	v = C.guint16(c.native.green)
}

// Blue gets the field inside the struct.
func (c *Color) Blue() uint16 {
	v = C.guint16(c.native.blue)
}

// Copy makes a copy of a Color.
//
// The result must be freed using gdk_color_free().
func (c *Color) Copy() *Color {
	var arg0 *C.GdkColor

	arg0 = (*C.GdkColor)(unsafe.Pointer(c.Native()))

	var cret *C.GdkColor
	var ret1 *Color

	cret = C.gdk_color_copy(arg0)

	ret1 = WrapColor(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Color) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Equal compares two colors.
func (c *Color) Equal(colorb *Color) bool {
	var arg0 *C.GdkColor
	var arg1 *C.GdkColor

	arg0 = (*C.GdkColor)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GdkColor)(unsafe.Pointer(colorb.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gdk_color_equal(arg0, colorb)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Free frees a Color created with gdk_color_copy().
func (c *Color) Free() {
	var arg0 *C.GdkColor

	arg0 = (*C.GdkColor)(unsafe.Pointer(c.Native()))

	C.gdk_color_free(arg0)
}

// Hash: a hash function suitable for using for a hash table that stores Colors.
func (c *Color) Hash() uint {
	var arg0 *C.GdkColor

	arg0 = (*C.GdkColor)(unsafe.Pointer(c.Native()))

	var cret C.guint
	var ret1 uint

	cret = C.gdk_color_hash(arg0)

	ret1 = C.guint(cret)

	return ret1
}

// String returns a textual specification of @color in the hexadecimal form
// “\#rrrrggggbbbb” where “r”, “g” and “b” are hex digits representing the red,
// green and blue components respectively.
//
// The returned string can be parsed by gdk_color_parse().
func (c *Color) String() string {
	var arg0 *C.GdkColor

	arg0 = (*C.GdkColor)(unsafe.Pointer(c.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gdk_color_to_string(arg0)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}
