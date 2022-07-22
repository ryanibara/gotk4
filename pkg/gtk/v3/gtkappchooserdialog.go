// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeAppChooserDialog = coreglib.Type(C.gtk_app_chooser_dialog_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAppChooserDialog, F: marshalAppChooserDialog},
	})
}

// AppChooserDialogOverrider contains methods that are overridable.
type AppChooserDialogOverrider interface {
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
	_ [0]func() // equal guard
	Dialog

	*coreglib.Object
	AppChooser
}

var (
	_ coreglib.Objector = (*AppChooserDialog)(nil)
	_ Binner            = (*AppChooserDialog)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:     GTypeAppChooserDialog,
		GoType:    reflect.TypeOf((*AppChooserDialog)(nil)),
		InitClass: initClassAppChooserDialog,
	})
}

func initClassAppChooserDialog(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ InitAppChooserDialog(*AppChooserDialogClass) }); ok {
		klass := (*AppChooserDialogClass)(gextras.NewStructNative(gclass))
		goval.InitAppChooserDialog(klass)
	}
}

func wrapAppChooserDialog(obj *coreglib.Object) *AppChooserDialog {
	return &AppChooserDialog{
		Dialog: Dialog{
			Window: Window{
				Bin: Bin{
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
				},
			},
		},
		Object: obj,
		AppChooser: AppChooser{
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
	}
}

func marshalAppChooserDialog(p uintptr) (interface{}, error) {
	return wrapAppChooserDialog(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAppChooserDialog creates a new AppChooserDialog for the provided #GFile,
// to allow the user to select an application for it.
//
// The function takes the following parameters:
//
//    - parent (optional) or NULL.
//    - flags for this dialog.
//    - file: #GFile.
//
// The function returns the following values:
//
//    - appChooserDialog: newly created AppChooserDialog.
//
func NewAppChooserDialog(parent *Window, flags DialogFlags, file gio.Filer) *AppChooserDialog {
	var _arg1 *C.GtkWindow     // out
	var _arg2 C.GtkDialogFlags // out
	var _arg3 *C.GFile         // out
	var _cret *C.GtkWidget     // in

	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	}
	_arg2 = C.GtkDialogFlags(flags)
	_arg3 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	_cret = C.gtk_app_chooser_dialog_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(file)

	var _appChooserDialog *AppChooserDialog // out

	_appChooserDialog = wrapAppChooserDialog(coreglib.Take(unsafe.Pointer(_cret)))

	return _appChooserDialog
}

// NewAppChooserDialogForContentType creates a new AppChooserDialog for the
// provided content type, to allow the user to select an application for it.
//
// The function takes the following parameters:
//
//    - parent (optional) or NULL.
//    - flags for this dialog.
//    - contentType: content type string.
//
// The function returns the following values:
//
//    - appChooserDialog: newly created AppChooserDialog.
//
func NewAppChooserDialogForContentType(parent *Window, flags DialogFlags, contentType string) *AppChooserDialog {
	var _arg1 *C.GtkWindow     // out
	var _arg2 C.GtkDialogFlags // out
	var _arg3 *C.gchar         // out
	var _cret *C.GtkWidget     // in

	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	}
	_arg2 = C.GtkDialogFlags(flags)
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(contentType)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.gtk_app_chooser_dialog_new_for_content_type(_arg1, _arg2, _arg3)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(contentType)

	var _appChooserDialog *AppChooserDialog // out

	_appChooserDialog = wrapAppChooserDialog(coreglib.Take(unsafe.Pointer(_cret)))

	return _appChooserDialog
}

// Heading returns the text to display at the top of the dialog.
//
// The function returns the following values:
//
//    - utf8 (optional): text to display at the top of the dialog, or NULL, in
//      which case a default text is displayed.
//
func (self *AppChooserDialog) Heading() string {
	var _arg0 *C.GtkAppChooserDialog // out
	var _cret *C.gchar               // in

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_app_chooser_dialog_get_heading(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Widget returns the AppChooserWidget of this dialog.
//
// The function returns the following values:
//
//    - widget of self.
//
func (self *AppChooserDialog) Widget() Widgetter {
	var _arg0 *C.GtkAppChooserDialog // out
	var _cret *C.GtkWidget           // in

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_app_chooser_dialog_get_widget(_arg0)
	runtime.KeepAlive(self)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// SetHeading sets the text to display at the top of the dialog. If the heading
// is not set, the dialog displays a default text.
//
// The function takes the following parameters:
//
//    - heading: string containing Pango markup.
//
func (self *AppChooserDialog) SetHeading(heading string) {
	var _arg0 *C.GtkAppChooserDialog // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(heading)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_dialog_set_heading(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(heading)
}

// AppChooserDialogClass: instance of this type is always passed by reference.
type AppChooserDialogClass struct {
	*appChooserDialogClass
}

// appChooserDialogClass is the struct that's finalized.
type appChooserDialogClass struct {
	native *C.GtkAppChooserDialogClass
}

// ParentClass: parent class.
func (a *AppChooserDialogClass) ParentClass() *DialogClass {
	valptr := &a.native.parent_class
	var v *DialogClass // out
	v = (*DialogClass)(gextras.NewStructNative(unsafe.Pointer((&*valptr))))
	return v
}
