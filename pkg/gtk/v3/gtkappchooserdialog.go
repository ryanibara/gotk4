// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkappchooserdialog.go.
var GTypeAppChooserDialog = coreglib.Type(C.gtk_app_chooser_dialog_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeAppChooserDialog, F: marshalAppChooserDialog},
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

func classInitAppChooserDialogger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

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

// Heading returns the text to display at the top of the dialog.
//
// The function returns the following values:
//
//    - utf8 (optional): text to display at the top of the dialog, or NULL, in
//      which case a default text is displayed.
//
func (self *AppChooserDialog) Heading() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "AppChooserDialog").InvokeMethod("get_heading", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "AppChooserDialog").InvokeMethod("get_widget", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(heading)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "AppChooserDialog").InvokeMethod("set_heading", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(heading)
}
