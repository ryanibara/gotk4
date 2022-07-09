// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeVisualType returns the GType for the type VisualType.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeVisualType() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "VisualType").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalVisualType)
	return gtype
}

// GTypeVisual returns the GType for the type Visual.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeVisual() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "Visual").RegisteredGType())
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
	_gret := girepository.MustFind("Gdk", "list_visuals").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _list []*Visual // out

	_list = make([]*Visual, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
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
func QueryDepths() []int32 {
	var _outs [2]girepository.Argument

	girepository.MustFind("Gdk", "query_depths").Invoke(nil, _outs[:])

	var _depths []int32 // out

	{
		src := unsafe.Slice((**C.void)(_outs[0]), _outs[1])
		_depths = make([]int32, _outs[1])
		for i := 0; i < int(_outs[1]); i++ {
			_depths[i] = *(*int32)(unsafe.Pointer(src[i]))
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
	var _outs [2]girepository.Argument

	girepository.MustFind("Gdk", "query_visual_types").Invoke(nil, _outs[:])

	var _visualTypes []VisualType // out

	_visualTypes = make([]VisualType, _outs[1])
	copy(_visualTypes, unsafe.Slice((*VisualType)(unsafe.Pointer(_outs[0])), _outs[1]))

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
func (visual *Visual) BitsPerRGB() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(visual).Native()))

	_gret := girepository.MustFind("Gdk", "Visual").InvokeMethod("get_bits_per_rgb", _args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(visual)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

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
func (visual *Visual) BluePixelDetails() (mask uint32, shift, precision int32) {
	var _args [1]girepository.Argument
	var _outs [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(visual).Native()))

	girepository.MustFind("Gdk", "Visual").InvokeMethod("get_blue_pixel_details", _args[:], _outs[:])

	runtime.KeepAlive(visual)

	var _mask uint32     // out
	var _shift int32     // out
	var _precision int32 // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_mask = *(*uint32)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_shift = *(*int32)(unsafe.Pointer(_outs[1]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[2])) != nil {
		_precision = *(*int32)(unsafe.Pointer(_outs[2]))
	}

	return _mask, _shift, _precision
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
func (visual *Visual) ColormapSize() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(visual).Native()))

	_gret := girepository.MustFind("Gdk", "Visual").InvokeMethod("get_colormap_size", _args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(visual)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// Depth returns the bit depth of this visual.
//
// The function returns the following values:
//
//    - gint: bit depth of this visual.
//
func (visual *Visual) Depth() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(visual).Native()))

	_gret := girepository.MustFind("Gdk", "Visual").InvokeMethod("get_depth", _args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(visual)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

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
func (visual *Visual) GreenPixelDetails() (mask uint32, shift, precision int32) {
	var _args [1]girepository.Argument
	var _outs [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(visual).Native()))

	girepository.MustFind("Gdk", "Visual").InvokeMethod("get_green_pixel_details", _args[:], _outs[:])

	runtime.KeepAlive(visual)

	var _mask uint32     // out
	var _shift int32     // out
	var _precision int32 // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_mask = *(*uint32)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_shift = *(*int32)(unsafe.Pointer(_outs[1]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[2])) != nil {
		_precision = *(*int32)(unsafe.Pointer(_outs[2]))
	}

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
func (visual *Visual) RedPixelDetails() (mask uint32, shift, precision int32) {
	var _args [1]girepository.Argument
	var _outs [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(visual).Native()))

	girepository.MustFind("Gdk", "Visual").InvokeMethod("get_red_pixel_details", _args[:], _outs[:])

	runtime.KeepAlive(visual)

	var _mask uint32     // out
	var _shift int32     // out
	var _precision int32 // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_mask = *(*uint32)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_shift = *(*int32)(unsafe.Pointer(_outs[1]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[2])) != nil {
		_precision = *(*int32)(unsafe.Pointer(_outs[2]))
	}

	return _mask, _shift, _precision
}

// Screen gets the screen to which this visual belongs.
//
// The function returns the following values:
//
//    - screen to which this visual belongs.
//
func (visual *Visual) Screen() *Screen {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(visual).Native()))

	_gret := girepository.MustFind("Gdk", "Visual").InvokeMethod("get_screen", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(visual)

	var _screen *Screen // out

	_screen = wrapScreen(coreglib.Take(unsafe.Pointer(_cret)))

	return _screen
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
	_gret := girepository.MustFind("Gdk", "get_best").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
func VisualGetBestDepth() int32 {
	_gret := girepository.MustFind("Gdk", "get_best_depth").Invoke(nil, nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
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
func VisualGetBestWithDepth(depth int32) *Visual {
	var _args [1]girepository.Argument

	*(*C.gint)(unsafe.Pointer(&_args[0])) = C.gint(depth)

	_gret := girepository.MustFind("Gdk", "get_best_with_depth").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(depth)

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
	_gret := girepository.MustFind("Gdk", "get_system").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _visual *Visual // out

	_visual = wrapVisual(coreglib.Take(unsafe.Pointer(_cret)))

	return _visual
}
