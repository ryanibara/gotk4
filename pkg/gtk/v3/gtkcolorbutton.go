// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_color_button_get_type()), F: marshalColorButton},
	})
}

// ColorButton: the ColorButton is a button which displays the currently
// selected color and allows to open a color selection dialog to change the
// color. It is suitable widget for selecting a color in a preference dialog.
//
//
// CSS nodes
//
// GtkColorButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .color style class.
type ColorButton interface {
	Button
	Actionable
	Activatable
	Buildable
	ColorChooser

	// Alpha returns the current alpha value.
	Alpha() uint16
	// Color sets @color to be the current color in the ColorButton widget.
	Color() gdk.Color
	// RGBA sets @rgba to be the current color in the ColorButton widget.
	RGBA() gdk.RGBA
	// Title gets the title of the color selection dialog.
	Title() string
	// UseAlpha does the color selection dialog use the alpha channel ?
	UseAlpha() bool
	// SetAlpha sets the current opacity to be @alpha.
	SetAlpha(alpha uint16)
	// SetColor sets the current color to be @color.
	SetColor(color *gdk.Color)
	// SetRGBA sets the current color to be @rgba.
	SetRGBA(rgba *gdk.RGBA)
	// SetTitle sets the title for the color selection dialog.
	SetTitle(title string)
	// SetUseAlpha sets whether or not the color button should use the alpha
	// channel.
	SetUseAlpha(useAlpha bool)
}

// colorButton implements the ColorButton interface.
type colorButton struct {
	Button
	Actionable
	Activatable
	Buildable
	ColorChooser
}

var _ ColorButton = (*colorButton)(nil)

// WrapColorButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapColorButton(obj *externglib.Object) ColorButton {
	return ColorButton{
		Button:       WrapButton(obj),
		Actionable:   WrapActionable(obj),
		Activatable:  WrapActivatable(obj),
		Buildable:    WrapBuildable(obj),
		ColorChooser: WrapColorChooser(obj),
	}
}

func marshalColorButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColorButton(obj), nil
}

// NewColorButton constructs a class ColorButton.
func NewColorButton() ColorButton {
	var cret C.GtkColorButton
	var ret1 ColorButton

	cret = C.gtk_color_button_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ColorButton)

	return ret1
}

// NewColorButtonWithColor constructs a class ColorButton.
func NewColorButtonWithColor(color *gdk.Color) ColorButton {
	var arg1 *C.GdkColor

	arg1 = (*C.GdkColor)(unsafe.Pointer(color.Native()))

	var cret C.GtkColorButton
	var ret1 ColorButton

	cret = C.gtk_color_button_new_with_color(color)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ColorButton)

	return ret1
}

// NewColorButtonWithRGBA constructs a class ColorButton.
func NewColorButtonWithRGBA(rgba *gdk.RGBA) ColorButton {
	var arg1 *C.GdkRGBA

	arg1 = (*C.GdkRGBA)(unsafe.Pointer(rgba.Native()))

	var cret C.GtkColorButton
	var ret1 ColorButton

	cret = C.gtk_color_button_new_with_rgba(rgba)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ColorButton)

	return ret1
}

// Alpha returns the current alpha value.
func (b colorButton) Alpha() uint16 {
	var arg0 *C.GtkColorButton

	arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))

	var cret C.guint16
	var ret1 uint16

	cret = C.gtk_color_button_get_alpha(arg0)

	ret1 = C.guint16(cret)

	return ret1
}

// Color sets @color to be the current color in the ColorButton widget.
func (b colorButton) Color() gdk.Color {
	var arg0 *C.GtkColorButton

	arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))

	var arg1 C.GdkColor
	var ret1 *gdk.Color

	C.gtk_color_button_get_color(arg0, &arg1)

	*ret1 = gdk.WrapColor(unsafe.Pointer(arg1))

	return ret1
}

// RGBA sets @rgba to be the current color in the ColorButton widget.
func (b colorButton) RGBA() gdk.RGBA {
	var arg0 *C.GtkColorButton

	arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))

	var arg1 C.GdkRGBA
	var ret1 *gdk.RGBA

	C.gtk_color_button_get_rgba(arg0, &arg1)

	*ret1 = gdk.WrapRGBA(unsafe.Pointer(arg1))

	return ret1
}

// Title gets the title of the color selection dialog.
func (b colorButton) Title() string {
	var arg0 *C.GtkColorButton

	arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_color_button_get_title(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// UseAlpha does the color selection dialog use the alpha channel ?
func (b colorButton) UseAlpha() bool {
	var arg0 *C.GtkColorButton

	arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_color_button_get_use_alpha(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// SetAlpha sets the current opacity to be @alpha.
func (b colorButton) SetAlpha(alpha uint16) {
	var arg0 *C.GtkColorButton
	var arg1 C.guint16

	arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))
	arg1 = C.guint16(alpha)

	C.gtk_color_button_set_alpha(arg0, alpha)
}

// SetColor sets the current color to be @color.
func (b colorButton) SetColor(color *gdk.Color) {
	var arg0 *C.GtkColorButton
	var arg1 *C.GdkColor

	arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GdkColor)(unsafe.Pointer(color.Native()))

	C.gtk_color_button_set_color(arg0, color)
}

// SetRGBA sets the current color to be @rgba.
func (b colorButton) SetRGBA(rgba *gdk.RGBA) {
	var arg0 *C.GtkColorButton
	var arg1 *C.GdkRGBA

	arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GdkRGBA)(unsafe.Pointer(rgba.Native()))

	C.gtk_color_button_set_rgba(arg0, rgba)
}

// SetTitle sets the title for the color selection dialog.
func (b colorButton) SetTitle(title string) {
	var arg0 *C.GtkColorButton
	var arg1 *C.gchar

	arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_color_button_set_title(arg0, title)
}

// SetUseAlpha sets whether or not the color button should use the alpha
// channel.
func (b colorButton) SetUseAlpha(useAlpha bool) {
	var arg0 *C.GtkColorButton
	var arg1 C.gboolean

	arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))
	if useAlpha {
		arg1 = C.gboolean(1)
	}

	C.gtk_color_button_set_use_alpha(arg0, useAlpha)
}
