// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_app_chooser_dialog_get_type()), F: marshalAppChooserDialog},
	})
}

// AppChooserDialog shows a AppChooserWidget inside a Dialog.
//
// Note that AppChooserDialog does not have any interesting methods of its own.
// Instead, you should get the embedded AppChooserWidget using
// gtk_app_chooser_dialog_get_widget() and call its methods if the generic
// AppChooser interface is not sufficient for your needs.
//
// To set the heading that is shown above the AppChooserWidget, use
// gtk_app_chooser_dialog_set_heading().
type AppChooserDialog interface {
	Dialog
	AppChooser
	Buildable

	// Heading returns the text to display at the top of the dialog.
	Heading() string
	// Widget returns the AppChooserWidget of this dialog.
	Widget() Widget
	// SetHeading sets the text to display at the top of the dialog. If the
	// heading is not set, the dialog displays a default text.
	SetHeading(heading string)
}

// appChooserDialog implements the AppChooserDialog interface.
type appChooserDialog struct {
	Dialog
	AppChooser
	Buildable
}

var _ AppChooserDialog = (*appChooserDialog)(nil)

// WrapAppChooserDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapAppChooserDialog(obj *externglib.Object) AppChooserDialog {
	return AppChooserDialog{
		Dialog:     WrapDialog(obj),
		AppChooser: WrapAppChooser(obj),
		Buildable:  WrapBuildable(obj),
	}
}

func marshalAppChooserDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppChooserDialog(obj), nil
}

// NewAppChooserDialog constructs a class AppChooserDialog.
func NewAppChooserDialog(parent Window, flags DialogFlags, file gio.File) AppChooserDialog {
	var arg1 *C.GtkWindow
	var arg2 C.GtkDialogFlags
	var arg3 *C.GFile

	arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	arg2 = (C.GtkDialogFlags)(flags)
	arg3 = (*C.GFile)(unsafe.Pointer(file.Native()))

	var cret C.GtkAppChooserDialog
	var ret1 AppChooserDialog

	cret = C.gtk_app_chooser_dialog_new(parent, flags, file)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(AppChooserDialog)

	return ret1
}

// NewAppChooserDialogForContentType constructs a class AppChooserDialog.
func NewAppChooserDialogForContentType(parent Window, flags DialogFlags, contentType string) AppChooserDialog {
	var arg1 *C.GtkWindow
	var arg2 C.GtkDialogFlags
	var arg3 *C.gchar

	arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	arg2 = (C.GtkDialogFlags)(flags)
	arg3 = (*C.gchar)(C.CString(contentType))
	defer C.free(unsafe.Pointer(arg3))

	var cret C.GtkAppChooserDialog
	var ret1 AppChooserDialog

	cret = C.gtk_app_chooser_dialog_new_for_content_type(parent, flags, contentType)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(AppChooserDialog)

	return ret1
}

// Heading returns the text to display at the top of the dialog.
func (s appChooserDialog) Heading() string {
	var arg0 *C.GtkAppChooserDialog

	arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(s.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_app_chooser_dialog_get_heading(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Widget returns the AppChooserWidget of this dialog.
func (s appChooserDialog) Widget() Widget {
	var arg0 *C.GtkAppChooserDialog

	arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(s.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_app_chooser_dialog_get_widget(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// SetHeading sets the text to display at the top of the dialog. If the
// heading is not set, the dialog displays a default text.
func (s appChooserDialog) SetHeading(heading string) {
	var arg0 *C.GtkAppChooserDialog
	var arg1 *C.gchar

	arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(heading))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_app_chooser_dialog_set_heading(arg0, heading)
}
