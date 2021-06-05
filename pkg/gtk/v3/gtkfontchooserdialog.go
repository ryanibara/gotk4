// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_font_chooser_dialog_get_type()), F: marshalFontChooserDialog},
	})
}

// FontChooserDialog: the FontChooserDialog widget is a dialog for selecting a
// font. It implements the FontChooser interface.
//
//
// GtkFontChooserDialog as GtkBuildable
//
// The GtkFontChooserDialog implementation of the Buildable interface exposes
// the buttons with the names “select_button” and “cancel_button”.
type FontChooserDialog interface {
	Dialog
	Buildable
	FontChooser
}

// fontChooserDialog implements the FontChooserDialog interface.
type fontChooserDialog struct {
	Dialog
	Buildable
	FontChooser
}

var _ FontChooserDialog = (*fontChooserDialog)(nil)

// WrapFontChooserDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontChooserDialog(obj *externglib.Object) FontChooserDialog {
	return FontChooserDialog{
		Dialog:      WrapDialog(obj),
		Buildable:   WrapBuildable(obj),
		FontChooser: WrapFontChooser(obj),
	}
}

func marshalFontChooserDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontChooserDialog(obj), nil
}

// NewFontChooserDialog constructs a class FontChooserDialog.
func NewFontChooserDialog(title string, parent Window) FontChooserDialog {
	var arg1 *C.gchar
	var arg2 *C.GtkWindow

	arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	var cret C.GtkFontChooserDialog
	var ret1 FontChooserDialog

	cret = C.gtk_font_chooser_dialog_new(title, parent)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(FontChooserDialog)

	return ret1
}

type FontChooserDialogPrivate struct {
	native C.GtkFontChooserDialogPrivate
}

// WrapFontChooserDialogPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFontChooserDialogPrivate(ptr unsafe.Pointer) *FontChooserDialogPrivate {
	if ptr == nil {
		return nil
	}

	return (*FontChooserDialogPrivate)(ptr)
}

func marshalFontChooserDialogPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFontChooserDialogPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *FontChooserDialogPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}
