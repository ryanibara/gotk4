// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_color_chooser_get_type()), F: marshalColorChooser},
	})
}

// ColorChooser is an interface that is implemented by widgets for choosing
// colors. Depending on the situation, colors may be allowed to have alpha
// (translucency).
//
// In GTK+, the main widgets that implement this interface are
// ColorChooserWidget, ColorChooserDialog and ColorButton.
type ColorChooser interface {
	gextras.Objector

	// AddPalette adds a palette to the color chooser. If @orientation is
	// horizontal, the colors are grouped in rows, with @colors_per_line colors
	// in each row. If @horizontal is false, the colors are grouped in columns
	// instead.
	//
	// The default color palette of ColorChooserWidget has 27 colors, organized
	// in columns of 3 colors. The default gray palette has 9 grays in a single
	// row.
	//
	// The layout of the color chooser widget works best when the palettes have
	// 9-10 columns.
	//
	// Calling this function for the first time has the side effect of removing
	// the default color and gray palettes from the color chooser.
	//
	// If @colors is nil, removes all previously added palettes.
	AddPalette(orientation Orientation, colorsPerLine int, colors []gdk.RGBA)
	// RGBA gets the currently-selected color.
	RGBA() gdk.RGBA
	// UseAlpha returns whether the color chooser shows the alpha channel.
	UseAlpha() bool
	// SetRGBA sets the color.
	SetRGBA(color gdk.RGBA)
	// SetUseAlpha sets whether or not the color chooser should use the alpha
	// channel.
	SetUseAlpha(useAlpha bool)
}

// colorChooser implements the ColorChooser interface.
type colorChooser struct {
	gextras.Objector
}

var _ ColorChooser = (*colorChooser)(nil)

// WrapColorChooser wraps a GObject to a type that implements
// interface ColorChooser. It is primarily used internally.
func WrapColorChooser(obj *externglib.Object) ColorChooser {
	return colorChooser{
		Objector: obj,
	}
}

func marshalColorChooser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColorChooser(obj), nil
}

func (c colorChooser) AddPalette(orientation Orientation, colorsPerLine int, colors []gdk.RGBA) {
	var _arg0 *C.GtkColorChooser // out
	var _arg1 C.GtkOrientation   // out
	var _arg2 C.gint             // out
	var _arg4 *C.GdkRGBA
	var _arg3 C.gint

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(c.Native()))
	_arg1 = C.GtkOrientation(orientation)
	_arg2 = C.gint(colorsPerLine)
	_arg3 = C.gint(len(colors))
	_arg4 = (*C.GdkRGBA)(unsafe.Pointer(&colors[0]))

	C.gtk_color_chooser_add_palette(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (c colorChooser) RGBA() gdk.RGBA {
	var _arg0 *C.GtkColorChooser // out
	var _arg1 *C.GdkRGBA         // in

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(c.Native()))

	C.gtk_color_chooser_get_rgba(_arg0, &_arg1)

	var _color gdk.RGBA // out

	_color = (gdk.RGBA)(unsafe.Pointer(_arg1))

	return _color
}

func (c colorChooser) UseAlpha() bool {
	var _arg0 *C.GtkColorChooser // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_color_chooser_get_use_alpha(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c colorChooser) SetRGBA(color gdk.RGBA) {
	var _arg0 *C.GtkColorChooser // out
	var _arg1 *C.GdkRGBA         // out

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkRGBA)(unsafe.Pointer(color))

	C.gtk_color_chooser_set_rgba(_arg0, _arg1)
}

func (c colorChooser) SetUseAlpha(useAlpha bool) {
	var _arg0 *C.GtkColorChooser // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(c.Native()))
	if useAlpha {
		_arg1 = C.TRUE
	}

	C.gtk_color_chooser_set_use_alpha(_arg0, _arg1)
}
