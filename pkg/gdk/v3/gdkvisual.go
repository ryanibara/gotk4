// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

// GTypeVisualType returns the GType for the type VisualType.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeVisualType() coreglib.Type {
	gtype := coreglib.Type(C.gdk_visual_type_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalVisualType)
	return gtype
}

// GTypeVisual returns the GType for the type Visual.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeVisual() coreglib.Type {
	gtype := coreglib.Type(C.gdk_visual_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalVisual)
	return gtype
}

// VisualType: set of values that describe the manner in which the pixel values
// for a visual are converted into RGB values for display.
type VisualType C.gint

const (
	// VisualStaticGray: each pixel value indexes a grayscale value directly.
	VisualStaticGray VisualType = iota
	// VisualGrayscale: each pixel is an index into a color map that maps pixel
	// values into grayscale values. The color map can be changed by an
	// application.
	VisualGrayscale
	// VisualStaticColor: each pixel value is an index into a predefined,
	// unmodifiable color map that maps pixel values into RGB values.
	VisualStaticColor
	// VisualPseudoColor: each pixel is an index into a color map that maps
	// pixel values into rgb values. The color map can be changed by an
	// application.
	VisualPseudoColor
	// VisualTrueColor: each pixel value directly contains red, green, and blue
	// components. Use gdk_visual_get_red_pixel_details(), etc, to obtain
	// information about how the components are assembled into a pixel value.
	VisualTrueColor
	// VisualDirectColor: each pixel value contains red, green, and blue
	// components as for GDK_VISUAL_TRUE_COLOR, but the components are mapped
	// via a color table into the final output table instead of being converted
	// directly.
	VisualDirectColor
)

func marshalVisualType(p uintptr) (interface{}, error) {
	return VisualType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for VisualType.
func (v VisualType) String() string {
	switch v {
	case VisualStaticGray:
		return "StaticGray"
	case VisualGrayscale:
		return "Grayscale"
	case VisualStaticColor:
		return "StaticColor"
	case VisualPseudoColor:
		return "PseudoColor"
	case VisualTrueColor:
		return "TrueColor"
	case VisualDirectColor:
		return "DirectColor"
	default:
		return fmt.Sprintf("VisualType(%d)", v)
	}
}

// ListVisuals lists the available visuals for the default screen. (See
// gdk_screen_list_visuals()) A visual describes a hardware image data format.
// For example, a visual might support 24-bit color, or 8-bit color, and might
// expect pixels to be in a certain format.
//
// Call g_list_free() on the return value when you’re finished with it.
//
// Deprecated: Use gdk_screen_list_visuals (gdk_screen_get_default ()).
//
// The function returns the following values:
//
//    - list: a list of visuals; the list must be freed, but not its contents.
//
func ListVisuals() []*Visual {
	var _cret *C.GList // in

	_cret = C.gdk_list_visuals()

	var _list []*Visual // out

	_list = make([]*Visual, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GdkVisual)(v)
		var dst *Visual // out
		dst = wrapVisual(coreglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// QueryDepths: this function returns the available bit depths for the default
// screen. It’s equivalent to listing the visuals (gdk_list_visuals()) and then
// looking at the depth field in each visual, removing duplicates.
//
// The array returned by this function should not be freed.
//
// Deprecated: Visual selection should be done using
// gdk_screen_get_system_visual() and gdk_screen_get_rgba_visual().
//
// The function returns the following values:
//
//    - depths: return location for available depths.
//
func QueryDepths() []int {
	var _arg1 *C.gint // in
	var _arg2 C.gint  // in

	C.gdk_query_depths(&_arg1, &_arg2)

	var _depths []int // out

	{
		src := unsafe.Slice((*C.gint)(_arg1), _arg2)
		_depths = make([]int, _arg2)
		for i := 0; i < int(_arg2); i++ {
			_depths[i] = int(src[i])
		}
	}

	return _depths
}

// QueryVisualTypes: this function returns the available visual types for the
// default screen. It’s equivalent to listing the visuals (gdk_list_visuals())
// and then looking at the type field in each visual, removing duplicates.
//
// The array returned by this function should not be freed.
//
// Deprecated: Visual selection should be done using
// gdk_screen_get_system_visual() and gdk_screen_get_rgba_visual().
//
// The function returns the following values:
//
//    - visualTypes: return location for the available visual types.
//
func QueryVisualTypes() []VisualType {
	var _arg1 *C.GdkVisualType // in
	var _arg2 C.gint           // in

	C.gdk_query_visual_types(&_arg1, &_arg2)

	var _visualTypes []VisualType // out

	_visualTypes = make([]VisualType, _arg2)
	copy(_visualTypes, unsafe.Slice((*VisualType)(unsafe.Pointer(_arg1)), _arg2))

	return _visualTypes
}

// Visual contains information about a particular visual.
type Visual struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Visual)(nil)
)

func wrapVisual(obj *coreglib.Object) *Visual {
	return &Visual{
		Object: obj,
	}
}

func marshalVisual(p uintptr) (interface{}, error) {
	return wrapVisual(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// BitsPerRGB returns the number of significant bits per red, green and blue
// value.
//
// Not all GDK backend provide a meaningful value for this function.
//
// Deprecated: Use gdk_visual_get_red_pixel_details() and its variants to learn
// about the pixel layout of TrueColor and DirectColor visuals.
//
// The function returns the following values:
//
//    - gint: number of significant bits per color value for visual.
//
func (visual *Visual) BitsPerRGB() int {
	var _arg0 *C.GdkVisual // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(coreglib.InternObject(visual).Native()))

	_cret = C.gdk_visual_get_bits_per_rgb(_arg0)
	runtime.KeepAlive(visual)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// BluePixelDetails obtains values that are needed to calculate blue pixel
// values in TrueColor and DirectColor. The “mask” is the significant bits
// within the pixel. The “shift” is the number of bits left we must shift a
// primary for it to be in position (according to the "mask"). Finally,
// "precision" refers to how much precision the pixel value contains for a
// particular primary.
//
// The function returns the following values:
//
//    - mask (optional): pointer to a #guint32 to be filled in, or NULL.
//    - shift (optional): pointer to a #gint to be filled in, or NULL.
//    - precision (optional): pointer to a #gint to be filled in, or NULL.
//
func (visual *Visual) BluePixelDetails() (mask uint32, shift, precision int) {
	var _arg0 *C.GdkVisual // out
	var _arg1 C.guint32    // in
	var _arg2 C.gint       // in
	var _arg3 C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(coreglib.InternObject(visual).Native()))

	C.gdk_visual_get_blue_pixel_details(_arg0, &_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(visual)

	var _mask uint32   // out
	var _shift int     // out
	var _precision int // out

	_mask = uint32(_arg1)
	_shift = int(_arg2)
	_precision = int(_arg3)

	return _mask, _shift, _precision
}

// ByteOrder returns the byte order of this visual.
//
// The information returned by this function is only relevant when working with
// XImages, and not all backends return meaningful information for this.
//
// Deprecated: This information is not useful.
//
// The function returns the following values:
//
//    - byteOrder stating the byte order of visual.
//
func (visual *Visual) ByteOrder() ByteOrder {
	var _arg0 *C.GdkVisual   // out
	var _cret C.GdkByteOrder // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(coreglib.InternObject(visual).Native()))

	_cret = C.gdk_visual_get_byte_order(_arg0)
	runtime.KeepAlive(visual)

	var _byteOrder ByteOrder // out

	_byteOrder = ByteOrder(_cret)

	return _byteOrder
}

// ColormapSize returns the size of a colormap for this visual.
//
// You have to use platform-specific APIs to manipulate colormaps.
//
// Deprecated: This information is not useful, since GDK does not provide APIs
// to operate on colormaps.
//
// The function returns the following values:
//
//    - gint: size of a colormap that is suitable for visual.
//
func (visual *Visual) ColormapSize() int {
	var _arg0 *C.GdkVisual // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(coreglib.InternObject(visual).Native()))

	_cret = C.gdk_visual_get_colormap_size(_arg0)
	runtime.KeepAlive(visual)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Depth returns the bit depth of this visual.
//
// The function returns the following values:
//
//    - gint: bit depth of this visual.
//
func (visual *Visual) Depth() int {
	var _arg0 *C.GdkVisual // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(coreglib.InternObject(visual).Native()))

	_cret = C.gdk_visual_get_depth(_arg0)
	runtime.KeepAlive(visual)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// GreenPixelDetails obtains values that are needed to calculate green pixel
// values in TrueColor and DirectColor. The “mask” is the significant bits
// within the pixel. The “shift” is the number of bits left we must shift a
// primary for it to be in position (according to the "mask"). Finally,
// "precision" refers to how much precision the pixel value contains for a
// particular primary.
//
// The function returns the following values:
//
//    - mask (optional): pointer to a #guint32 to be filled in, or NULL.
//    - shift (optional): pointer to a #gint to be filled in, or NULL.
//    - precision (optional): pointer to a #gint to be filled in, or NULL.
//
func (visual *Visual) GreenPixelDetails() (mask uint32, shift, precision int) {
	var _arg0 *C.GdkVisual // out
	var _arg1 C.guint32    // in
	var _arg2 C.gint       // in
	var _arg3 C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(coreglib.InternObject(visual).Native()))

	C.gdk_visual_get_green_pixel_details(_arg0, &_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(visual)

	var _mask uint32   // out
	var _shift int     // out
	var _precision int // out

	_mask = uint32(_arg1)
	_shift = int(_arg2)
	_precision = int(_arg3)

	return _mask, _shift, _precision
}

// RedPixelDetails obtains values that are needed to calculate red pixel values
// in TrueColor and DirectColor. The “mask” is the significant bits within the
// pixel. The “shift” is the number of bits left we must shift a primary for it
// to be in position (according to the "mask"). Finally, "precision" refers to
// how much precision the pixel value contains for a particular primary.
//
// The function returns the following values:
//
//    - mask (optional): pointer to a #guint32 to be filled in, or NULL.
//    - shift (optional): pointer to a #gint to be filled in, or NULL.
//    - precision (optional): pointer to a #gint to be filled in, or NULL.
//
func (visual *Visual) RedPixelDetails() (mask uint32, shift, precision int) {
	var _arg0 *C.GdkVisual // out
	var _arg1 C.guint32    // in
	var _arg2 C.gint       // in
	var _arg3 C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(coreglib.InternObject(visual).Native()))

	C.gdk_visual_get_red_pixel_details(_arg0, &_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(visual)

	var _mask uint32   // out
	var _shift int     // out
	var _precision int // out

	_mask = uint32(_arg1)
	_shift = int(_arg2)
	_precision = int(_arg3)

	return _mask, _shift, _precision
}

// Screen gets the screen to which this visual belongs.
//
// The function returns the following values:
//
//    - screen to which this visual belongs.
//
func (visual *Visual) Screen() *Screen {
	var _arg0 *C.GdkVisual // out
	var _cret *C.GdkScreen // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(coreglib.InternObject(visual).Native()))

	_cret = C.gdk_visual_get_screen(_arg0)
	runtime.KeepAlive(visual)

	var _screen *Screen // out

	_screen = wrapScreen(coreglib.Take(unsafe.Pointer(_cret)))

	return _screen
}

// VisualType returns the type of visual this is (PseudoColor, TrueColor, etc).
//
// The function returns the following values:
//
//    - visualType stating the type of visual.
//
func (visual *Visual) VisualType() VisualType {
	var _arg0 *C.GdkVisual    // out
	var _cret C.GdkVisualType // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(coreglib.InternObject(visual).Native()))

	_cret = C.gdk_visual_get_visual_type(_arg0)
	runtime.KeepAlive(visual)

	var _visualType VisualType // out

	_visualType = VisualType(_cret)

	return _visualType
}

// VisualGetBest: get the visual with the most available colors for the default
// GDK screen. The return value should not be freed.
//
// Deprecated: Visual selection should be done using
// gdk_screen_get_system_visual() and gdk_screen_get_rgba_visual().
//
// The function returns the following values:
//
//    - visual: best visual.
//
func VisualGetBest() *Visual {
	var _cret *C.GdkVisual // in

	_cret = C.gdk_visual_get_best()

	var _visual *Visual // out

	_visual = wrapVisual(coreglib.Take(unsafe.Pointer(_cret)))

	return _visual
}

// VisualGetBestDepth: get the best available depth for the default GDK screen.
// “Best” means “largest,” i.e. 32 preferred over 24 preferred over 8 bits per
// pixel.
//
// Deprecated: Visual selection should be done using
// gdk_screen_get_system_visual() and gdk_screen_get_rgba_visual().
//
// The function returns the following values:
//
//    - gint: best available depth.
//
func VisualGetBestDepth() int {
	var _cret C.gint // in

	_cret = C.gdk_visual_get_best_depth()

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// VisualGetBestType: return the best available visual type for the default GDK
// screen.
//
// Deprecated: Visual selection should be done using
// gdk_screen_get_system_visual() and gdk_screen_get_rgba_visual().
//
// The function returns the following values:
//
//    - visualType: best visual type.
//
func VisualGetBestType() VisualType {
	var _cret C.GdkVisualType // in

	_cret = C.gdk_visual_get_best_type()

	var _visualType VisualType // out

	_visualType = VisualType(_cret)

	return _visualType
}

// VisualGetBestWithBoth combines gdk_visual_get_best_with_depth() and
// gdk_visual_get_best_with_type().
//
// Deprecated: Visual selection should be done using
// gdk_screen_get_system_visual() and gdk_screen_get_rgba_visual().
//
// The function takes the following parameters:
//
//    - depth: bit depth.
//    - visualType: visual type.
//
// The function returns the following values:
//
//    - visual (optional): best visual with both depth and visual_type, or NULL
//      if none.
//
func VisualGetBestWithBoth(depth int, visualType VisualType) *Visual {
	var _arg1 C.gint          // out
	var _arg2 C.GdkVisualType // out
	var _cret *C.GdkVisual    // in

	_arg1 = C.gint(depth)
	_arg2 = C.GdkVisualType(visualType)

	_cret = C.gdk_visual_get_best_with_both(_arg1, _arg2)
	runtime.KeepAlive(depth)
	runtime.KeepAlive(visualType)

	var _visual *Visual // out

	if _cret != nil {
		_visual = wrapVisual(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _visual
}

// VisualGetBestWithDepth: get the best visual with depth depth for the default
// GDK screen. Color visuals and visuals with mutable colormaps are preferred
// over grayscale or fixed-colormap visuals. The return value should not be
// freed. NULL may be returned if no visual supports depth.
//
// Deprecated: Visual selection should be done using
// gdk_screen_get_system_visual() and gdk_screen_get_rgba_visual().
//
// The function takes the following parameters:
//
//    - depth: bit depth.
//
// The function returns the following values:
//
//    - visual: best visual for the given depth.
//
func VisualGetBestWithDepth(depth int) *Visual {
	var _arg1 C.gint       // out
	var _cret *C.GdkVisual // in

	_arg1 = C.gint(depth)

	_cret = C.gdk_visual_get_best_with_depth(_arg1)
	runtime.KeepAlive(depth)

	var _visual *Visual // out

	_visual = wrapVisual(coreglib.Take(unsafe.Pointer(_cret)))

	return _visual
}

// VisualGetBestWithType: get the best visual of the given visual_type for the
// default GDK screen. Visuals with higher color depths are considered better.
// The return value should not be freed. NULL may be returned if no visual has
// type visual_type.
//
// Deprecated: Visual selection should be done using
// gdk_screen_get_system_visual() and gdk_screen_get_rgba_visual().
//
// The function takes the following parameters:
//
//    - visualType: visual type.
//
// The function returns the following values:
//
//    - visual: best visual of the given type.
//
func VisualGetBestWithType(visualType VisualType) *Visual {
	var _arg1 C.GdkVisualType // out
	var _cret *C.GdkVisual    // in

	_arg1 = C.GdkVisualType(visualType)

	_cret = C.gdk_visual_get_best_with_type(_arg1)
	runtime.KeepAlive(visualType)

	var _visual *Visual // out

	_visual = wrapVisual(coreglib.Take(unsafe.Pointer(_cret)))

	return _visual
}

// VisualGetSystem: get the system’s default visual for the default GDK screen.
// This is the visual for the root window of the display. The return value
// should not be freed.
//
// Deprecated: Use gdk_screen_get_system_visual (gdk_screen_get_default ()).
//
// The function returns the following values:
//
//    - visual: system visual.
//
func VisualGetSystem() *Visual {
	var _cret *C.GdkVisual // in

	_cret = C.gdk_visual_get_system()

	var _visual *Visual // out

	_visual = wrapVisual(coreglib.Take(unsafe.Pointer(_cret)))

	return _visual
}
