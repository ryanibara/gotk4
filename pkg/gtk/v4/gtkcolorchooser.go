// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_ColorChooserInterface_add_palette(GtkColorChooser*, GtkOrientation, int, int, GdkRGBA*);
// extern void _gotk4_gtk4_ColorChooserInterface_color_activated(GtkColorChooser*, GdkRGBA*);
// extern void _gotk4_gtk4_ColorChooserInterface_get_rgba(GtkColorChooser*, GdkRGBA*);
// extern void _gotk4_gtk4_ColorChooserInterface_set_rgba(GtkColorChooser*, GdkRGBA*);
// extern void _gotk4_gtk4_ColorChooser_ConnectColorActivated(gpointer, GdkRGBA*, guintptr);
import "C"

// GTypeColorChooser returns the GType for the type ColorChooser.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeColorChooser() coreglib.Type {
	gtype := coreglib.Type(C.gtk_color_chooser_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalColorChooser)
	return gtype
}

// ColorChooserOverrider contains methods that are overridable.
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
	AddPalette(orientation Orientation, colorsPerLine int32, colors []gdk.RGBA)
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
//
// ColorChooser wraps an interface. This means the user can get the
// underlying type by calling Cast().
type ColorChooser struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ColorChooser)(nil)
)

// ColorChooserer describes ColorChooser's interface methods.
type ColorChooserer interface {
	coreglib.Objector

	// AddPalette adds a palette to the color chooser.
	AddPalette(orientation Orientation, colorsPerLine int32, colors []gdk.RGBA)
	// RGBA gets the currently-selected color.
	RGBA() *gdk.RGBA
	// UseAlpha returns whether the color chooser shows the alpha channel.
	UseAlpha() bool
	// SetRGBA sets the color.
	SetRGBA(color *gdk.RGBA)
	// SetUseAlpha sets whether or not the color chooser should use the alpha
	// channel.
	SetUseAlpha(useAlpha bool)

	// Color-activated is emitted when a color is activated from the color
	// chooser.
	ConnectColorActivated(func(color *gdk.RGBA)) coreglib.SignalHandle
}

var _ ColorChooserer = (*ColorChooser)(nil)

func ifaceInitColorChooserer(gifacePtr, data C.gpointer) {
	iface := (*C.GtkColorChooserInterface)(unsafe.Pointer(gifacePtr))
	iface.add_palette = (*[0]byte)(C._gotk4_gtk4_ColorChooserInterface_add_palette)
	iface.color_activated = (*[0]byte)(C._gotk4_gtk4_ColorChooserInterface_color_activated)
	iface.get_rgba = (*[0]byte)(C._gotk4_gtk4_ColorChooserInterface_get_rgba)
	iface.set_rgba = (*[0]byte)(C._gotk4_gtk4_ColorChooserInterface_set_rgba)
}

//export _gotk4_gtk4_ColorChooserInterface_add_palette
func _gotk4_gtk4_ColorChooserInterface_add_palette(arg0 *C.GtkColorChooser, arg1 C.GtkOrientation, arg2 C.int, arg3 C.int, arg4 *C.GdkRGBA) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ColorChooserOverrider)

	var _orientation Orientation // out
	var _colorsPerLine int32     // out
	var _colors []gdk.RGBA       // out

	_orientation = Orientation(arg1)
	_colorsPerLine = int32(arg2)
	if arg4 != nil {
		{
			src := unsafe.Slice((*C.GdkRGBA)(arg4), arg3)
			_colors = make([]gdk.RGBA, arg3)
			for i := 0; i < int(arg3); i++ {
				_colors[i] = *(*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
			}
		}
	}

	iface.AddPalette(_orientation, _colorsPerLine, _colors)
}

//export _gotk4_gtk4_ColorChooserInterface_color_activated
func _gotk4_gtk4_ColorChooserInterface_color_activated(arg0 *C.GtkColorChooser, arg1 *C.GdkRGBA) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ColorChooserOverrider)

	var _color *gdk.RGBA // out

	_color = (*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	iface.ColorActivated(_color)
}

//export _gotk4_gtk4_ColorChooserInterface_get_rgba
func _gotk4_gtk4_ColorChooserInterface_get_rgba(arg0 *C.GtkColorChooser, arg1 *C.GdkRGBA) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ColorChooserOverrider)

	color := iface.RGBA()

	*arg1 = *(*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(color)))
}

//export _gotk4_gtk4_ColorChooserInterface_set_rgba
func _gotk4_gtk4_ColorChooserInterface_set_rgba(arg0 *C.GtkColorChooser, arg1 *C.GdkRGBA) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ColorChooserOverrider)

	var _color *gdk.RGBA // out

	_color = (*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	iface.SetRGBA(_color)
}

func wrapColorChooser(obj *coreglib.Object) *ColorChooser {
	return &ColorChooser{
		Object: obj,
	}
}

func marshalColorChooser(p uintptr) (interface{}, error) {
	return wrapColorChooser(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_ColorChooser_ConnectColorActivated
func _gotk4_gtk4_ColorChooser_ConnectColorActivated(arg0 C.gpointer, arg1 *C.GdkRGBA, arg2 C.guintptr) {
	var f func(color *gdk.RGBA)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(color *gdk.RGBA))
	}

	var _color *gdk.RGBA // out

	_color = (*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	f(_color)
}

// ConnectColorActivated is emitted when a color is activated from the color
// chooser.
//
// This usually happens when the user clicks a color swatch, or a color is
// selected and the user presses one of the keys Space, Shift+Space, Return or
// Enter.
func (chooser *ColorChooser) ConnectColorActivated(f func(color *gdk.RGBA)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(chooser, "color-activated", false, unsafe.Pointer(C._gotk4_gtk4_ColorChooser_ConnectColorActivated), f)
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
func (chooser *ColorChooser) AddPalette(orientation Orientation, colorsPerLine int32, colors []gdk.RGBA) {
	var _arg0 *C.GtkColorChooser // out
	var _arg1 C.GtkOrientation   // out
	var _arg2 C.int              // out
	var _arg4 *C.GdkRGBA         // out
	var _arg3 C.int

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
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

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

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

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

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

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
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

	_arg0 = (*C.GtkColorChooser)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	if useAlpha {
		_arg1 = C.TRUE
	}

	C.gtk_color_chooser_set_use_alpha(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(useAlpha)
}
