// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk3_ColorSelectionClass_color_changed(void*);
// extern void _gotk4_gtk3_ColorSelection_ConnectColorChanged(gpointer, guintptr);
import "C"

// GTypeColorSelection returns the GType for the type ColorSelection.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeColorSelection() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ColorSelection").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalColorSelection)
	return gtype
}

// ColorSelectionOverrider contains methods that are overridable.
type ColorSelectionOverrider interface {
	ColorChanged()
}

type ColorSelection struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer       = (*ColorSelection)(nil)
	_ coreglib.Objector = (*ColorSelection)(nil)
)

func classInitColorSelectioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "ColorSelectionClass")

	if _, ok := goval.(interface{ ColorChanged() }); ok {
		o := pclass.StructFieldOffset("color_changed")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ColorSelectionClass_color_changed)
	}
}

//export _gotk4_gtk3_ColorSelectionClass_color_changed
func _gotk4_gtk3_ColorSelectionClass_color_changed(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ColorChanged() })

	iface.ColorChanged()
}

func wrapColorSelection(obj *coreglib.Object) *ColorSelection {
	return &ColorSelection{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalColorSelection(p uintptr) (interface{}, error) {
	return wrapColorSelection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_ColorSelection_ConnectColorChanged
func _gotk4_gtk3_ColorSelection_ConnectColorChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectColorChanged: this signal is emitted when the color changes in the
// ColorSelection according to its update policy.
func (colorsel *ColorSelection) ConnectColorChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(colorsel, "color-changed", false, unsafe.Pointer(C._gotk4_gtk3_ColorSelection_ConnectColorChanged), f)
}

// NewColorSelection creates a new GtkColorSelection.
//
// The function returns the following values:
//
//    - colorSelection: new ColorSelection.
//
func NewColorSelection() *ColorSelection {
	_info := girepository.MustFind("Gtk", "ColorSelection")
	_gret := _info.InvokeClassMethod("new_ColorSelection", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _colorSelection *ColorSelection // out

	_colorSelection = wrapColorSelection(coreglib.Take(unsafe.Pointer(_cret)))

	return _colorSelection
}

// CurrentAlpha returns the current alpha value.
//
// The function returns the following values:
//
//    - guint16: integer between 0 and 65535.
//
func (colorsel *ColorSelection) CurrentAlpha() uint16 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_gret := _info.InvokeClassMethod("get_current_alpha", _args[:], nil)
	_cret := *(*C.guint16)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(colorsel)

	var _guint16 uint16 // out

	_guint16 = uint16(*(*C.guint16)(unsafe.Pointer(&_cret)))

	return _guint16
}

// CurrentColor sets color to be the current color in the GtkColorSelection
// widget.
//
// Deprecated: Use gtk_color_selection_get_current_rgba() instead.
//
// The function returns the following values:
//
//    - color to fill in with the current color.
//
func (colorsel *ColorSelection) CurrentColor() *gdk.Color {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_info.InvokeClassMethod("get_current_color", _args[:], _outs[:])

	runtime.KeepAlive(colorsel)

	var _color *gdk.Color // out

	_color = (*gdk.Color)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))

	return _color
}

// CurrentRGBA sets rgba to be the current color in the GtkColorSelection
// widget.
//
// The function returns the following values:
//
//    - rgba to fill in with the current color.
//
func (colorsel *ColorSelection) CurrentRGBA() *gdk.RGBA {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_info.InvokeClassMethod("get_current_rgba", _args[:], _outs[:])

	runtime.KeepAlive(colorsel)

	var _rgba *gdk.RGBA // out

	_rgba = (*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))

	return _rgba
}

// HasOpacityControl determines whether the colorsel has an opacity control.
//
// The function returns the following values:
//
//    - ok: TRUE if the colorsel has an opacity control, FALSE if it does't.
//
func (colorsel *ColorSelection) HasOpacityControl() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_gret := _info.InvokeClassMethod("get_has_opacity_control", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(colorsel)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// HasPalette determines whether the color selector has a color palette.
//
// The function returns the following values:
//
//    - ok: TRUE if the selector has a palette, FALSE if it hasn't.
//
func (colorsel *ColorSelection) HasPalette() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_gret := _info.InvokeClassMethod("get_has_palette", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(colorsel)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// PreviousAlpha returns the previous alpha value.
//
// The function returns the following values:
//
//    - guint16: integer between 0 and 65535.
//
func (colorsel *ColorSelection) PreviousAlpha() uint16 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_gret := _info.InvokeClassMethod("get_previous_alpha", _args[:], nil)
	_cret := *(*C.guint16)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(colorsel)

	var _guint16 uint16 // out

	_guint16 = uint16(*(*C.guint16)(unsafe.Pointer(&_cret)))

	return _guint16
}

// PreviousColor fills color in with the original color value.
//
// Deprecated: Use gtk_color_selection_get_previous_rgba() instead.
//
// The function returns the following values:
//
//    - color to fill in with the original color value.
//
func (colorsel *ColorSelection) PreviousColor() *gdk.Color {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_info.InvokeClassMethod("get_previous_color", _args[:], _outs[:])

	runtime.KeepAlive(colorsel)

	var _color *gdk.Color // out

	_color = (*gdk.Color)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))

	return _color
}

// PreviousRGBA fills rgba in with the original color value.
//
// The function returns the following values:
//
//    - rgba to fill in with the original color value.
//
func (colorsel *ColorSelection) PreviousRGBA() *gdk.RGBA {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_info.InvokeClassMethod("get_previous_rgba", _args[:], _outs[:])

	runtime.KeepAlive(colorsel)

	var _rgba *gdk.RGBA // out

	_rgba = (*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))

	return _rgba
}

// IsAdjusting gets the current state of the colorsel.
//
// The function returns the following values:
//
//    - ok: TRUE if the user is currently dragging a color around, and FALSE if
//      the selection has stopped.
//
func (colorsel *ColorSelection) IsAdjusting() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_gret := _info.InvokeClassMethod("is_adjusting", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(colorsel)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetCurrentAlpha sets the current opacity to be alpha.
//
// The first time this is called, it will also set the original opacity to be
// alpha too.
//
// The function takes the following parameters:
//
//    - alpha: integer between 0 and 65535.
//
func (colorsel *ColorSelection) SetCurrentAlpha(alpha uint16) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))
	*(*C.guint16)(unsafe.Pointer(&_args[1])) = C.guint16(alpha)

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_info.InvokeClassMethod("set_current_alpha", _args[:], nil)

	runtime.KeepAlive(colorsel)
	runtime.KeepAlive(alpha)
}

// SetCurrentColor sets the current color to be color.
//
// The first time this is called, it will also set the original color to be
// color too.
//
// Deprecated: Use gtk_color_selection_set_current_rgba() instead.
//
// The function takes the following parameters:
//
//    - color to set the current color with.
//
func (colorsel *ColorSelection) SetCurrentColor(color *gdk.Color) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(color)))

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_info.InvokeClassMethod("set_current_color", _args[:], nil)

	runtime.KeepAlive(colorsel)
	runtime.KeepAlive(color)
}

// SetCurrentRGBA sets the current color to be rgba.
//
// The first time this is called, it will also set the original color to be rgba
// too.
//
// The function takes the following parameters:
//
//    - rgba to set the current color with.
//
func (colorsel *ColorSelection) SetCurrentRGBA(rgba *gdk.RGBA) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(rgba)))

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_info.InvokeClassMethod("set_current_rgba", _args[:], nil)

	runtime.KeepAlive(colorsel)
	runtime.KeepAlive(rgba)
}

// SetHasOpacityControl sets the colorsel to use or not use opacity.
//
// The function takes the following parameters:
//
//    - hasOpacity: TRUE if colorsel can set the opacity, FALSE otherwise.
//
func (colorsel *ColorSelection) SetHasOpacityControl(hasOpacity bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))
	if hasOpacity {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_info.InvokeClassMethod("set_has_opacity_control", _args[:], nil)

	runtime.KeepAlive(colorsel)
	runtime.KeepAlive(hasOpacity)
}

// SetHasPalette shows and hides the palette based upon the value of
// has_palette.
//
// The function takes the following parameters:
//
//    - hasPalette: TRUE if palette is to be visible, FALSE otherwise.
//
func (colorsel *ColorSelection) SetHasPalette(hasPalette bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))
	if hasPalette {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_info.InvokeClassMethod("set_has_palette", _args[:], nil)

	runtime.KeepAlive(colorsel)
	runtime.KeepAlive(hasPalette)
}

// SetPreviousAlpha sets the “previous” alpha to be alpha.
//
// This function should be called with some hesitations, as it might seem
// confusing to have that alpha change.
//
// The function takes the following parameters:
//
//    - alpha: integer between 0 and 65535.
//
func (colorsel *ColorSelection) SetPreviousAlpha(alpha uint16) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))
	*(*C.guint16)(unsafe.Pointer(&_args[1])) = C.guint16(alpha)

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_info.InvokeClassMethod("set_previous_alpha", _args[:], nil)

	runtime.KeepAlive(colorsel)
	runtime.KeepAlive(alpha)
}

// SetPreviousColor sets the “previous” color to be color.
//
// This function should be called with some hesitations, as it might seem
// confusing to have that color change. Calling
// gtk_color_selection_set_current_color() will also set this color the first
// time it is called.
//
// Deprecated: Use gtk_color_selection_set_previous_rgba() instead.
//
// The function takes the following parameters:
//
//    - color to set the previous color with.
//
func (colorsel *ColorSelection) SetPreviousColor(color *gdk.Color) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(color)))

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_info.InvokeClassMethod("set_previous_color", _args[:], nil)

	runtime.KeepAlive(colorsel)
	runtime.KeepAlive(color)
}

// SetPreviousRGBA sets the “previous” color to be rgba.
//
// This function should be called with some hesitations, as it might seem
// confusing to have that color change. Calling
// gtk_color_selection_set_current_rgba() will also set this color the first
// time it is called.
//
// The function takes the following parameters:
//
//    - rgba to set the previous color with.
//
func (colorsel *ColorSelection) SetPreviousRGBA(rgba *gdk.RGBA) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(rgba)))

	_info := girepository.MustFind("Gtk", "ColorSelection")
	_info.InvokeClassMethod("set_previous_rgba", _args[:], nil)

	runtime.KeepAlive(colorsel)
	runtime.KeepAlive(rgba)
}

// ColorSelectionPaletteFromString parses a color palette string; the string is
// a colon-separated list of color names readable by gdk_color_parse().
//
// The function takes the following parameters:
//
//    - str: string encoding a color palette.
//
// The function returns the following values:
//
//    - colors: return location for allocated array of Color.
//    - ok: TRUE if a palette was successfully parsed.
//
func ColorSelectionPaletteFromString(str string) ([]gdk.Color, bool) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_args[0]))

	_info := girepository.MustFind("Gtk", "palette_from_string")
	_gret := _info.Invoke(_args[:], _outs[:])
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(str)

	var _colors []gdk.Color // out
	var _ok bool            // out

	defer C.free(unsafe.Pointer(_outs[0]))
	{
		src := unsafe.Slice((**C.void)(_outs[0]), _outs[1])
		_colors = make([]gdk.Color, _outs[1])
		for i := 0; i < int(_outs[1]); i++ {
			_colors[i] = *(*gdk.Color)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(&_colors[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					{
						args := [1]girepository.Argument{(*C.void)(intern.C)}
						girepository.MustFind("Gdk", "Color").InvokeRecordMethod("free", args[:], nil)
					}
				},
			)
		}
	}
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _colors, _ok
}
