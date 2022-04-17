// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// glib.Type values for gtkappchooserdialog.go.
var GTypeAppChooserDialog = externglib.Type(C.gtk_app_chooser_dialog_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeAppChooserDialog, F: marshalAppChooserDialog},
	})
}

// AppChooserDialog: GtkAppChooserDialog shows a GtkAppChooserWidget inside a
// GtkDialog.
//
// !An example GtkAppChooserDialog (appchooserdialog.png)
//
// Note that GtkAppChooserDialog does not have any interesting methods of its
// own. Instead, you should get the embedded GtkAppChooserWidget using
// gtk.AppChooserDialog.GetWidget() and call its methods if the generic
// gtk.AppChooser interface is not sufficient for your needs.
//
// To set the heading that is shown above the GtkAppChooserWidget, use
// gtk.AppChooserDialog.SetHeading().
type AppChooserDialog struct {
	_ [0]func() // equal guard
	Dialog

	*externglib.Object
	AppChooser
}

var (
	_ externglib.Objector = (*AppChooserDialog)(nil)
	_ Widgetter           = (*AppChooserDialog)(nil)
)

func wrapAppChooserDialog(obj *externglib.Object) *AppChooserDialog {
	return &AppChooserDialog{
		Dialog: Dialog{
			Window: Window{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					Accessible: Accessible{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					ConstraintTarget: ConstraintTarget{
						Object: obj,
					},
				},
				Object: obj,
				Root: Root{
					NativeSurface: NativeSurface{
						Widget: Widget{
							InitiallyUnowned: externglib.InitiallyUnowned{
								Object: obj,
							},
							Object: obj,
							Accessible: Accessible{
								Object: obj,
							},
							Buildable: Buildable{
								Object: obj,
							},
							ConstraintTarget: ConstraintTarget{
								Object: obj,
							},
						},
					},
				},
				ShortcutManager: ShortcutManager{
					Object: obj,
				},
			},
		},
		Object: obj,
		AppChooser: AppChooser{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
		},
	}
}

func marshalAppChooserDialog(p uintptr) (interface{}, error) {
	return wrapAppChooserDialog(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAppChooserDialog creates a new GtkAppChooserDialog for the provided GFile.
//
// The dialog will show applications that can open the file.
//
// The function takes the following parameters:
//
//    - parent (optional): GtkWindow, or NULL.
//    - flags for this dialog.
//    - file: GFile.
//
// The function returns the following values:
//
//    - appChooserDialog: newly created GtkAppChooserDialog.
//
func NewAppChooserDialog(parent *Window, flags DialogFlags, file gio.FileOverrider) *AppChooserDialog {
	var _arg1 *C.GtkWindow     // out
	var _arg2 C.GtkDialogFlags // out
	var _arg3 *C.GFile         // out
	var _cret *C.GtkWidget     // in

	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(externglib.InternObject(parent).Native()))
	}
	_arg2 = C.GtkDialogFlags(flags)
	_arg3 = (*C.GFile)(unsafe.Pointer(externglib.InternObject(file).Native()))

	_cret = C.gtk_app_chooser_dialog_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(file)

	var _appChooserDialog *AppChooserDialog // out

	_appChooserDialog = wrapAppChooserDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _appChooserDialog
}

// NewAppChooserDialogForContentType creates a new GtkAppChooserDialog for the
// provided content type.
//
// The dialog will show applications that can open the content type.
//
// The function takes the following parameters:
//
//    - parent (optional): GtkWindow, or NULL.
//    - flags for this dialog.
//    - contentType: content type string.
//
// The function returns the following values:
//
//    - appChooserDialog: newly created GtkAppChooserDialog.
//
func NewAppChooserDialogForContentType(parent *Window, flags DialogFlags, contentType string) *AppChooserDialog {
	var _arg1 *C.GtkWindow     // out
	var _arg2 C.GtkDialogFlags // out
	var _arg3 *C.char          // out
	var _cret *C.GtkWidget     // in

	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(externglib.InternObject(parent).Native()))
	}
	_arg2 = C.GtkDialogFlags(flags)
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(contentType)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.gtk_app_chooser_dialog_new_for_content_type(_arg1, _arg2, _arg3)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(contentType)

	var _appChooserDialog *AppChooserDialog // out

	_appChooserDialog = wrapAppChooserDialog(externglib.Take(unsafe.Pointer(_cret)))

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
	var _cret *C.char                // in

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_app_chooser_dialog_get_heading(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Widget returns the GtkAppChooserWidget of this dialog.
//
// The function returns the following values:
//
//    - widget: GtkAppChooserWidget of self.
//
func (self *AppChooserDialog) Widget() Widgetter {
	var _arg0 *C.GtkAppChooserDialog // out
	var _cret *C.GtkWidget           // in

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_app_chooser_dialog_get_widget(_arg0)
	runtime.KeepAlive(self)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
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

// SetHeading sets the text to display at the top of the dialog.
//
// If the heading is not set, the dialog displays a default text.
//
// The function takes the following parameters:
//
//    - heading: string containing Pango markup.
//
func (self *AppChooserDialog) SetHeading(heading string) {
	var _arg0 *C.GtkAppChooserDialog // out
	var _arg1 *C.char                // out

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(heading)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_dialog_set_heading(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(heading)
}
