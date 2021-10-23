// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_color_get_type()), F: marshalColor},
	})
}

// Color is used to describe a color, similar to the XColor struct used in the
// X11 drawing API.
//
// Deprecated: Use RGBA.
//
// An instance of this type is always passed by reference.
type Color struct {
	*color
}

// color is the struct that's finalized.
type color struct {
	native *C.GdkColor
}

func marshalColor(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Color{&color{(*C.GdkColor)(b)}}, nil
}

// NewColor creates a new Color instance from the given
// fields.
func NewColor(pixel uint32, red, green, blue uint16) Color {
	var f0 C.guint32 // out
	f0 = C.guint32(pixel)
	var f1 C.guint16 // out
	f1 = C.guint16(red)
	var f2 C.guint16 // out
	f2 = C.guint16(green)
	var f3 C.guint16 // out
	f3 = C.guint16(blue)

	v := C.GdkColor{
		pixel: f0,
		red:   f1,
		green: f2,
		blue:  f3,
	}

	return *(*Color)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

// Pixel: for allocated colors, the pixel value used to draw this color on the
// screen. Not used anymore.
func (c *Color) Pixel() uint32 {
	var v uint32 // out
	v = uint32(c.native.pixel)
	return v
}

// Red: red component of the color. This is a value between 0 and 65535, with
// 65535 indicating full intensity.
func (c *Color) Red() uint16 {
	var v uint16 // out
	v = uint16(c.native.red)
	return v
}

// Green: green component of the color.
func (c *Color) Green() uint16 {
	var v uint16 // out
	v = uint16(c.native.green)
	return v
}

// Blue: blue component of the color.
func (c *Color) Blue() uint16 {
	var v uint16 // out
	v = uint16(c.native.blue)
	return v
}

// Copy makes a copy of a Color.
//
// The result must be freed using gdk_color_free().
//
// Deprecated: Use RGBA.
func (color *Color) Copy() *Color {
	var _arg0 *C.GdkColor // out
	var _cret *C.GdkColor // in

	_arg0 = (*C.GdkColor)(gextras.StructNative(unsafe.Pointer(color)))

	_cret = C.gdk_color_copy(_arg0)
	runtime.KeepAlive(color)

	var _ret *Color // out

	_ret = (*Color)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_ret)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_color_free((*C.GdkColor)(intern.C))
		},
	)

	return _ret
}

// Equal compares two colors.
//
// Deprecated: Use RGBA.
func (colora *Color) Equal(colorb *Color) bool {
	var _arg0 *C.GdkColor // out
	var _arg1 *C.GdkColor // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GdkColor)(gextras.StructNative(unsafe.Pointer(colora)))
	_arg1 = (*C.GdkColor)(gextras.StructNative(unsafe.Pointer(colorb)))

	_cret = C.gdk_color_equal(_arg0, _arg1)
	runtime.KeepAlive(colora)
	runtime.KeepAlive(colorb)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Hash: hash function suitable for using for a hash table that stores Colors.
//
// Deprecated: Use RGBA.
func (color *Color) Hash() uint {
	var _arg0 *C.GdkColor // out
	var _cret C.guint     // in

	_arg0 = (*C.GdkColor)(gextras.StructNative(unsafe.Pointer(color)))

	_cret = C.gdk_color_hash(_arg0)
	runtime.KeepAlive(color)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// String returns a textual specification of color in the hexadecimal form
// “\#rrrrggggbbbb” where “r”, “g” and “b” are hex digits representing the red,
// green and blue components respectively.
//
// The returned string can be parsed by gdk_color_parse().
//
// Deprecated: Use RGBA.
func (color *Color) String() string {
	var _arg0 *C.GdkColor // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GdkColor)(gextras.StructNative(unsafe.Pointer(color)))

	_cret = C.gdk_color_to_string(_arg0)
	runtime.KeepAlive(color)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ColorParse parses a textual specification of a color and fill in the red,
// green, and blue fields of a Color.
//
// The string can either one of a large set of standard names (taken from the
// X11 rgb.txt file), or it can be a hexadecimal value in the form “\#rgb”
// “\#rrggbb”, “\#rrrgggbbb” or “\#rrrrggggbbbb” where “r”, “g” and “b” are hex
// digits of the red, green, and blue components of the color, respectively.
// (White in the four forms is “\#fff”, “\#ffffff”, “\#fffffffff” and
// “\#ffffffffffff”).
//
// Deprecated: Use RGBA.
//
// The function takes the following parameters:
//
//    - spec: string specifying the color.
//
func ColorParse(spec string) (*Color, bool) {
	var _arg1 *C.gchar   // out
	var _arg2 C.GdkColor // in
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(spec)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_color_parse(_arg1, &_arg2)
	runtime.KeepAlive(spec)

	var _color *Color // out
	var _ok bool      // out

	_color = (*Color)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret != 0 {
		_ok = true
	}

	return _color, _ok
}
