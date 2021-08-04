// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_color_button_get_type()), F: marshalColorButtonner},
	})
}

// ColorButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ColorButtonOverrider interface {
	ColorSet()
}

// ColorButton is a button which displays the currently selected color and
// allows to open a color selection dialog to change the color. It is suitable
// widget for selecting a color in a preference dialog.
//
//
// CSS nodes
//
// GtkColorButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .color style class.
type ColorButton struct {
	Button

	ColorChooser
	*externglib.Object
}

func wrapColorButton(obj *externglib.Object) *ColorButton {
	return &ColorButton{
		Button: Button{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
						Object: obj,
					},
				},
			},
			Actionable: Actionable{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					Object: obj,
				},
			},
			Activatable: Activatable{
				Object: obj,
			},
			Object: obj,
		},
		ColorChooser: ColorChooser{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalColorButtonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapColorButton(obj), nil
}

// NewColorButton creates a new color button.
//
// This returns a widget in the form of a small button containing a swatch
// representing the current selected color. When the button is clicked, a
// color-selection dialog will open, allowing the user to select a color. The
// swatch will be updated to reflect the new color when the user finishes.
func NewColorButton() *ColorButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_color_button_new()

	var _colorButton *ColorButton // out

	_colorButton = wrapColorButton(externglib.Take(unsafe.Pointer(_cret)))

	return _colorButton
}

// NewColorButtonWithColor creates a new color button.
//
// Deprecated: Use gtk_color_button_new_with_rgba() instead.
func NewColorButtonWithColor(color *gdk.Color) *ColorButton {
	var _arg1 *C.GdkColor  // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkColor)(gextras.StructNative(unsafe.Pointer(color)))

	_cret = C.gtk_color_button_new_with_color(_arg1)

	var _colorButton *ColorButton // out

	_colorButton = wrapColorButton(externglib.Take(unsafe.Pointer(_cret)))

	return _colorButton
}

// NewColorButtonWithRGBA creates a new color button.
func NewColorButtonWithRGBA(rgba *gdk.RGBA) *ColorButton {
	var _arg1 *C.GdkRGBA   // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(rgba)))

	_cret = C.gtk_color_button_new_with_rgba(_arg1)

	var _colorButton *ColorButton // out

	_colorButton = wrapColorButton(externglib.Take(unsafe.Pointer(_cret)))

	return _colorButton
}

// Alpha returns the current alpha value.
//
// Deprecated: Use gtk_color_chooser_get_rgba() instead.
func (button *ColorButton) Alpha() uint16 {
	var _arg0 *C.GtkColorButton // out
	var _cret C.guint16         // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_color_button_get_alpha(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// Color sets color to be the current color in the ColorButton widget.
//
// Deprecated: Use gtk_color_chooser_get_rgba() instead.
func (button *ColorButton) Color() gdk.Color {
	var _arg0 *C.GtkColorButton // out
	var _arg1 C.GdkColor        // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(button.Native()))

	C.gtk_color_button_get_color(_arg0, &_arg1)

	var _color gdk.Color // out

	_color = *(*gdk.Color)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _color
}

// Title gets the title of the color selection dialog.
func (button *ColorButton) Title() string {
	var _arg0 *C.GtkColorButton // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_color_button_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UseAlpha does the color selection dialog use the alpha channel ?
//
// Deprecated: Use gtk_color_chooser_get_use_alpha() instead.
func (button *ColorButton) UseAlpha() bool {
	var _arg0 *C.GtkColorButton // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_color_button_get_use_alpha(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAlpha sets the current opacity to be alpha.
//
// Deprecated: Use gtk_color_chooser_set_rgba() instead.
func (button *ColorButton) SetAlpha(alpha uint16) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 C.guint16         // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(button.Native()))
	_arg1 = C.guint16(alpha)

	C.gtk_color_button_set_alpha(_arg0, _arg1)
}

// SetColor sets the current color to be color.
//
// Deprecated: Use gtk_color_chooser_set_rgba() instead.
func (button *ColorButton) SetColor(color *gdk.Color) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 *C.GdkColor       // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.GdkColor)(gextras.StructNative(unsafe.Pointer(color)))

	C.gtk_color_button_set_color(_arg0, _arg1)
}

// SetTitle sets the title for the color selection dialog.
func (button *ColorButton) SetTitle(title string) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_color_button_set_title(_arg0, _arg1)
}

// SetUseAlpha sets whether or not the color button should use the alpha
// channel.
//
// Deprecated: Use gtk_color_chooser_set_use_alpha() instead.
func (button *ColorButton) SetUseAlpha(useAlpha bool) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(button.Native()))
	if useAlpha {
		_arg1 = C.TRUE
	}

	C.gtk_color_button_set_use_alpha(_arg0, _arg1)
}
