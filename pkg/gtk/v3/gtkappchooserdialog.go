// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
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
		{T: externglib.Type(C.gtk_app_chooser_dialog_get_type()), F: marshalAppChooserDialoger},
	})
}

// AppChooserDialoger describes AppChooserDialog's methods.
type AppChooserDialoger interface {
	// Heading returns the text to display at the top of the dialog.
	Heading() string
	// Widget returns the AppChooserWidget of this dialog.
	Widget() *Widget
	// SetHeading sets the text to display at the top of the dialog.
	SetHeading(heading string)
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
type AppChooserDialog struct {
	Dialog

	AppChooser
}

var (
	_ AppChooserDialoger = (*AppChooserDialog)(nil)
	_ gextras.Nativer    = (*AppChooserDialog)(nil)
)

func wrapAppChooserDialog(obj *externglib.Object) *AppChooserDialog {
	return &AppChooserDialog{
		Dialog: Dialog{
			Window: Window{
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
						},
					},
				},
			},
		},
		AppChooser: AppChooser{
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
			},
		},
	}
}

func marshalAppChooserDialoger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAppChooserDialog(obj), nil
}

// NewAppChooserDialog creates a new AppChooserDialog for the provided #GFile,
// to allow the user to select an application for it.
func NewAppChooserDialog(parent *Window, flags DialogFlags, file gio.Filer) *AppChooserDialog {
	var _arg1 *C.GtkWindow     // out
	var _arg2 C.GtkDialogFlags // out
	var _arg3 *C.GFile         // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	_arg2 = C.GtkDialogFlags(flags)
	_arg3 = (*C.GFile)(unsafe.Pointer((file).(gextras.Nativer).Native()))

	_cret = C.gtk_app_chooser_dialog_new(_arg1, _arg2, _arg3)

	var _appChooserDialog *AppChooserDialog // out

	_appChooserDialog = wrapAppChooserDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _appChooserDialog
}

// NewAppChooserDialogForContentType creates a new AppChooserDialog for the
// provided content type, to allow the user to select an application for it.
func NewAppChooserDialogForContentType(parent *Window, flags DialogFlags, contentType string) *AppChooserDialog {
	var _arg1 *C.GtkWindow     // out
	var _arg2 C.GtkDialogFlags // out
	var _arg3 *C.gchar         // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	_arg2 = C.GtkDialogFlags(flags)
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(contentType)))

	_cret = C.gtk_app_chooser_dialog_new_for_content_type(_arg1, _arg2, _arg3)

	var _appChooserDialog *AppChooserDialog // out

	_appChooserDialog = wrapAppChooserDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _appChooserDialog
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *AppChooserDialog) Native() uintptr {
	return v.Dialog.Window.Bin.Container.Widget.InitiallyUnowned.Object.Native()
}

// Heading returns the text to display at the top of the dialog.
func (self *AppChooserDialog) Heading() string {
	var _arg0 *C.GtkAppChooserDialog // out
	var _cret *C.gchar               // in

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_dialog_get_heading(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Widget returns the AppChooserWidget of this dialog.
func (self *AppChooserDialog) Widget() *Widget {
	var _arg0 *C.GtkAppChooserDialog // out
	var _cret *C.GtkWidget           // in

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_dialog_get_widget(_arg0)

	var _widget *Widget // out

	_widget = wrapWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _widget
}

// SetHeading sets the text to display at the top of the dialog. If the heading
// is not set, the dialog displays a default text.
func (self *AppChooserDialog) SetHeading(heading string) {
	var _arg0 *C.GtkAppChooserDialog // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(heading)))

	C.gtk_app_chooser_dialog_set_heading(_arg0, _arg1)
}
