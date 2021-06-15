// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_visual_type_get_type()), F: marshalVisualType},
		{T: externglib.Type(C.gdk_visual_get_type()), F: marshalVisual},
	})
}

// VisualType: a set of values that describe the manner in which the pixel
// values for a visual are converted into RGB values for display.
type VisualType int

const (
	// VisualTypeStaticGray: each pixel value indexes a grayscale value
	// directly.
	VisualTypeStaticGray VisualType = 0
	// VisualTypeGrayscale: each pixel is an index into a color map that maps
	// pixel values into grayscale values. The color map can be changed by an
	// application.
	VisualTypeGrayscale VisualType = 1
	// VisualTypeStaticColor: each pixel value is an index into a predefined,
	// unmodifiable color map that maps pixel values into RGB values.
	VisualTypeStaticColor VisualType = 2
	// VisualTypePseudoColor: each pixel is an index into a color map that maps
	// pixel values into rgb values. The color map can be changed by an
	// application.
	VisualTypePseudoColor VisualType = 3
	// VisualTypeTrueColor: each pixel value directly contains red, green, and
	// blue components. Use gdk_visual_get_red_pixel_details(), etc, to obtain
	// information about how the components are assembled into a pixel value.
	VisualTypeTrueColor VisualType = 4
	// VisualTypeDirectColor: each pixel value contains red, green, and blue
	// components as for GDK_VISUAL_TRUE_COLOR, but the components are mapped
	// via a color table into the final output table instead of being converted
	// directly.
	VisualTypeDirectColor VisualType = 5
)

func marshalVisualType(p uintptr) (interface{}, error) {
	return VisualType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// QueryDepths: this function returns the available bit depths for the default
// screen. It’s equivalent to listing the visuals (gdk_list_visuals()) and then
// looking at the depth field in each visual, removing duplicates.
//
// The array returned by this function should not be freed.
func QueryDepths() []int {
	var _arg1 *C.gint
	var _arg2 C.gint // in

	C.gdk_query_depths(&_arg1, &_arg2)

	var _depths []int

	{
		src := unsafe.Slice(_arg1, _arg2)
		_depths = make([]int, _arg2)
		for i := 0; i < int(_arg2); i++ {
			_depths[i] = (int)(src[i])
		}
	}

	return _depths
}

// QueryVisualTypes: this function returns the available visual types for the
// default screen. It’s equivalent to listing the visuals (gdk_list_visuals())
// and then looking at the type field in each visual, removing duplicates.
//
// The array returned by this function should not be freed.
func QueryVisualTypes() []VisualType {
	var _arg1 *C.GdkVisualType
	var _arg2 C.gint // in

	C.gdk_query_visual_types(&_arg1, &_arg2)

	var _visualTypes []VisualType

	{
		src := unsafe.Slice(_arg1, _arg2)
		_visualTypes = make([]VisualType, _arg2)
		for i := 0; i < int(_arg2); i++ {
			_visualTypes[i] = VisualType(src[i])
		}
	}

	return _visualTypes
}

// Visual: a Visual contains information about a particular visual.
type Visual interface {
	gextras.Objector

	// BitsPerRGB returns the number of significant bits per red, green and blue
	// value.
	//
	// Not all GDK backend provide a meaningful value for this function.
	BitsPerRGB() int
	// BluePixelDetails obtains values that are needed to calculate blue pixel
	// values in TrueColor and DirectColor. The “mask” is the significant bits
	// within the pixel. The “shift” is the number of bits left we must shift a
	// primary for it to be in position (according to the "mask"). Finally,
	// "precision" refers to how much precision the pixel value contains for a
	// particular primary.
	BluePixelDetails() (mask uint32, shift int, precision int)
	// ByteOrder returns the byte order of this visual.
	//
	// The information returned by this function is only relevant when working
	// with XImages, and not all backends return meaningful information for
	// this.
	ByteOrder() ByteOrder
	// ColormapSize returns the size of a colormap for this visual.
	//
	// You have to use platform-specific APIs to manipulate colormaps.
	ColormapSize() int
	// Depth returns the bit depth of this visual.
	Depth() int
	// GreenPixelDetails obtains values that are needed to calculate green pixel
	// values in TrueColor and DirectColor. The “mask” is the significant bits
	// within the pixel. The “shift” is the number of bits left we must shift a
	// primary for it to be in position (according to the "mask"). Finally,
	// "precision" refers to how much precision the pixel value contains for a
	// particular primary.
	GreenPixelDetails() (mask uint32, shift int, precision int)
	// RedPixelDetails obtains values that are needed to calculate red pixel
	// values in TrueColor and DirectColor. The “mask” is the significant bits
	// within the pixel. The “shift” is the number of bits left we must shift a
	// primary for it to be in position (according to the "mask"). Finally,
	// "precision" refers to how much precision the pixel value contains for a
	// particular primary.
	RedPixelDetails() (mask uint32, shift int, precision int)
	// Screen gets the screen to which this visual belongs
	Screen() Screen
	// VisualType returns the type of visual this is (PseudoColor, TrueColor,
	// etc).
	VisualType() VisualType
}

// visual implements the Visual class.
type visual struct {
	gextras.Objector
}

var _ Visual = (*visual)(nil)

// WrapVisual wraps a GObject to the right type. It is
// primarily used internally.
func WrapVisual(obj *externglib.Object) Visual {
	return visual{
		Objector: obj,
	}
}

func marshalVisual(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVisual(obj), nil
}

// BitsPerRGB returns the number of significant bits per red, green and blue
// value.
//
// Not all GDK backend provide a meaningful value for this function.
func (v visual) BitsPerRGB() int {
	var _arg0 *C.GdkVisual // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(v.Native()))

	_cret = C.gdk_visual_get_bits_per_rgb(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// BluePixelDetails obtains values that are needed to calculate blue pixel
// values in TrueColor and DirectColor. The “mask” is the significant bits
// within the pixel. The “shift” is the number of bits left we must shift a
// primary for it to be in position (according to the "mask"). Finally,
// "precision" refers to how much precision the pixel value contains for a
// particular primary.
func (v visual) BluePixelDetails() (mask uint32, shift int, precision int) {
	var _arg0 *C.GdkVisual // out
	var _arg1 C.guint32    // in
	var _arg2 C.gint       // in
	var _arg3 C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(v.Native()))

	C.gdk_visual_get_blue_pixel_details(_arg0, &_arg1, &_arg2, &_arg3)

	var _mask uint32   // out
	var _shift int     // out
	var _precision int // out

	_mask = (uint32)(_arg1)
	_shift = (int)(_arg2)
	_precision = (int)(_arg3)

	return _mask, _shift, _precision
}

// ByteOrder returns the byte order of this visual.
//
// The information returned by this function is only relevant when working
// with XImages, and not all backends return meaningful information for
// this.
func (v visual) ByteOrder() ByteOrder {
	var _arg0 *C.GdkVisual   // out
	var _cret C.GdkByteOrder // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(v.Native()))

	_cret = C.gdk_visual_get_byte_order(_arg0)

	var _byteOrder ByteOrder // out

	_byteOrder = ByteOrder(_cret)

	return _byteOrder
}

// ColormapSize returns the size of a colormap for this visual.
//
// You have to use platform-specific APIs to manipulate colormaps.
func (v visual) ColormapSize() int {
	var _arg0 *C.GdkVisual // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(v.Native()))

	_cret = C.gdk_visual_get_colormap_size(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Depth returns the bit depth of this visual.
func (v visual) Depth() int {
	var _arg0 *C.GdkVisual // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(v.Native()))

	_cret = C.gdk_visual_get_depth(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// GreenPixelDetails obtains values that are needed to calculate green pixel
// values in TrueColor and DirectColor. The “mask” is the significant bits
// within the pixel. The “shift” is the number of bits left we must shift a
// primary for it to be in position (according to the "mask"). Finally,
// "precision" refers to how much precision the pixel value contains for a
// particular primary.
func (v visual) GreenPixelDetails() (mask uint32, shift int, precision int) {
	var _arg0 *C.GdkVisual // out
	var _arg1 C.guint32    // in
	var _arg2 C.gint       // in
	var _arg3 C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(v.Native()))

	C.gdk_visual_get_green_pixel_details(_arg0, &_arg1, &_arg2, &_arg3)

	var _mask uint32   // out
	var _shift int     // out
	var _precision int // out

	_mask = (uint32)(_arg1)
	_shift = (int)(_arg2)
	_precision = (int)(_arg3)

	return _mask, _shift, _precision
}

// RedPixelDetails obtains values that are needed to calculate red pixel
// values in TrueColor and DirectColor. The “mask” is the significant bits
// within the pixel. The “shift” is the number of bits left we must shift a
// primary for it to be in position (according to the "mask"). Finally,
// "precision" refers to how much precision the pixel value contains for a
// particular primary.
func (v visual) RedPixelDetails() (mask uint32, shift int, precision int) {
	var _arg0 *C.GdkVisual // out
	var _arg1 C.guint32    // in
	var _arg2 C.gint       // in
	var _arg3 C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(v.Native()))

	C.gdk_visual_get_red_pixel_details(_arg0, &_arg1, &_arg2, &_arg3)

	var _mask uint32   // out
	var _shift int     // out
	var _precision int // out

	_mask = (uint32)(_arg1)
	_shift = (int)(_arg2)
	_precision = (int)(_arg3)

	return _mask, _shift, _precision
}

// Screen gets the screen to which this visual belongs
func (v visual) Screen() Screen {
	var _arg0 *C.GdkVisual // out
	var _cret *C.GdkScreen // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(v.Native()))

	_cret = C.gdk_visual_get_screen(_arg0)

	var _screen Screen // out

	_screen = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Screen)

	return _screen
}

// VisualType returns the type of visual this is (PseudoColor, TrueColor,
// etc).
func (v visual) VisualType() VisualType {
	var _arg0 *C.GdkVisual    // out
	var _cret C.GdkVisualType // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(v.Native()))

	_cret = C.gdk_visual_get_visual_type(_arg0)

	var _visualType VisualType // out

	_visualType = VisualType(_cret)

	return _visualType
}
