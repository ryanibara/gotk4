// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_color_chooser_get_type()), F: marshalColorChooserer},
	})
}

// ColorChooserOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ColorChooserOverrider interface {
	// AddPalette adds a palette to the color chooser.
	//
	// If orientation is horizontal, the colors are grouped in rows, with
	// colors_per_line colors in each row. If horizontal is FALSE, the colors
	// are grouped in columns instead.
	//
	// The default color palette of gtk.ColorChooserWidget has 45 colors,
	// organized in columns of 5 colors (this includes some grays).
	//
	// The layout of the color chooser widget works best when the palettes have
	// 9-10 columns.
	//
	// Calling this function for the first time has the side effect of removing
	// the default color palette from the color chooser.
	//
	// If colors is NULL, removes all previously added palettes.
	//
	// The function takes the following parameters:
	//
	//    - orientation: GTK_ORIENTATION_HORIZONTAL if the palette should be
	//      displayed in rows, GTK_ORIENTATION_VERTICAL for columns.
	//    - colorsPerLine: number of colors to show in each row/column.
	//    - colors (optional) of the palette, or NULL.
	//
	AddPalette(orientation Orientation, colorsPerLine int, colors []gdk.RGBA)
	// The function takes the following parameters:
	//
	ColorActivated(color *gdk.RGBA)
	// RGBA gets the currently-selected color.
	//
	// The function returns the following values:
	//
	//    - color: GdkRGBA to fill in with the current color.
	//
	RGBA() *gdk.RGBA
	// SetRGBA sets the color.
	//
	// The function takes the following parameters:
	//
	//    - color: new color.
	//
	SetRGBA(color *gdk.RGBA)
}

// ColorChooser: GtkColorChooser is an interface that is implemented by widgets
// for choosing colors.
//
// Depending on the situation, colors may be allowed to have alpha
// (translucency).
//
// In GTK, the main widgets that implement this interface are
// gtk.ColorChooserWidget, gtk.ColorChooserDialog and gtk.ColorButton.
type ColorChooser struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*ColorChooser)(nil)
)

// ColorChooserer describes ColorChooser's interface methods.
type ColorChooserer interface {
	externglib.Objector

	// AddPalette adds a palette to the color chooser.
	AddPalette(orientation Orientation, colorsPerLine int, colors []gdk.RGBA)
	// RGBA gets the currently-selected color.
	RGBA() *gdk.RGBA
	// UseAlpha returns whether the color chooser shows the alpha channel.
	UseAlpha() bool
	// SetRGBA sets the color.
	SetRGBA(color *gdk.RGBA)
	// SetUseAlpha sets whether or not the color chooser should use the alpha
	// channel.
	SetUseAlpha(useAlpha bool)
}

var _ ColorChooserer = (*ColorChooser)(nil)

func wrapColorChooser(obj *externglib.Object) *ColorChooser {
	return &ColorChooser{
		Object: obj,
	}
}

func marshalColorChooserer(p uintptr) (interface{}, error) {
	return wrapColorChooser(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectColorActivated: emitted when a color is activated from the color
// chooser.
//
// This usually happens when the user clicks a color swatch, or a color is
// selected and the user presses one of the keys Space, Shift+Space, Return or
// Enter.
func (chooser *ColorChooser) ConnectColorActivated(f func(color *gdk.RGBA)) externglib.SignalHandle {
	return chooser.Connect("color-activated", f)
}

// AddPalette adds a palette to the color chooser.
//
// If orientation is horizontal, the colors are grouped in rows, with
// colors_per_line colors in each row. If horizontal is FALSE, the colors are
// grouped in columns instead.
//
// The default color palette of gtk.ColorChooserWidget has 45 colors, organized
// in columns of 5 colors (this includes some grays).
//
// The layout of the color chooser widget works best when the palettes have 9-10
// columns.
//
// Calling this function for the first time has the side effect of removing the
// default color palette from the color chooser.
//
// If colors is NULL, removes all previously added palettes.
//
// The function takes the following parameters:
//
//    - orientation: GTK_ORIENTATION_HORIZONTAL if the palette should be
//      displayed in rows, GTK_ORIENTATION_VERTICAL for columns.
//    - colorsPerLine: number of colors to show in each row/column.
//    - colors (optional) of the palette, or NULL.
//
func (chooser *ColorChooser) AddPalette(orientation Orientation, colorsPerLine int, colors []gdk.RGBA) {
	var _arg0 *C.GtkColorChooser // out
	var _arg1 C.GtkOrientation   // out
	var _arg2 C.int              // out
	var _arg4 *C.GdkRGBA         // out
	var _arg3 C.int

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = C.GtkOrientation(orientation)
	_arg2 = C.int(colorsPerLine)
	_arg3 = (C.int)(len(colors))
	_arg4 = (*C.GdkRGBA)(C.calloc(C.size_t(len(colors)), C.size_t(C.sizeof_GdkRGBA)))
	defer C.free(unsafe.Pointer(_arg4))
	{
		out := unsafe.Slice((*C.GdkRGBA)(_arg4), len(colors))
		for i := range colors {
			out[i] = *(*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer((&colors[i]))))
		}
	}

	C.gtk_color_chooser_add_palette(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(orientation)
	runtime.KeepAlive(colorsPerLine)
	runtime.KeepAlive(colors)
}

// RGBA gets the currently-selected color.
//
// The function returns the following values:
//
//    - color: GdkRGBA to fill in with the current color.
//
func (chooser *ColorChooser) RGBA() *gdk.RGBA {
	var _arg0 *C.GtkColorChooser // out
	var _arg1 C.GdkRGBA          // in

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(chooser.Native()))

	C.gtk_color_chooser_get_rgba(_arg0, &_arg1)
	runtime.KeepAlive(chooser)

	var _color *gdk.RGBA // out

	_color = (*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _color
}

// UseAlpha returns whether the color chooser shows the alpha channel.
//
// The function returns the following values:
//
//    - ok: TRUE if the color chooser uses the alpha channel, FALSE if not.
//
func (chooser *ColorChooser) UseAlpha() bool {
	var _arg0 *C.GtkColorChooser // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_color_chooser_get_use_alpha(_arg0)
	runtime.KeepAlive(chooser)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetRGBA sets the color.
//
// The function takes the following parameters:
//
//    - color: new color.
//
func (chooser *ColorChooser) SetRGBA(color *gdk.RGBA) {
	var _arg0 *C.GtkColorChooser // out
	var _arg1 *C.GdkRGBA         // out

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(color)))

	C.gtk_color_chooser_set_rgba(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(color)
}

// SetUseAlpha sets whether or not the color chooser should use the alpha
// channel.
//
// The function takes the following parameters:
//
//    - useAlpha: TRUE if color chooser should use alpha channel, FALSE if not.
//
func (chooser *ColorChooser) SetUseAlpha(useAlpha bool) {
	var _arg0 *C.GtkColorChooser // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(chooser.Native()))
	if useAlpha {
		_arg1 = C.TRUE
	}

	C.gtk_color_chooser_set_use_alpha(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(useAlpha)
}
