// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtkcolorseldialog.go.
var GTypeColorSelectionDialog = externglib.Type(C.gtk_color_selection_dialog_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeColorSelectionDialog, F: marshalColorSelectionDialog},
	})
}

// ColorSelectionDialogOverrider contains methods that are overridable.
type ColorSelectionDialogOverrider interface {
	externglib.Objector
}

// WrapColorSelectionDialogOverrider wraps the ColorSelectionDialogOverrider
// interface implementation to access the instance methods.
func WrapColorSelectionDialogOverrider(obj ColorSelectionDialogOverrider) *ColorSelectionDialog {
	return wrapColorSelectionDialog(externglib.BaseObject(obj))
}

type ColorSelectionDialog struct {
	_ [0]func() // equal guard
	Dialog
}

var (
	_ Binner = (*ColorSelectionDialog)(nil)
)

func classInitColorSelectionDialogger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapColorSelectionDialog(obj *externglib.Object) *ColorSelectionDialog {
	return &ColorSelectionDialog{
		Dialog: Dialog{
			Window: Window{
				Bin: Bin{
					Container: Container{
						Widget: Widget{
							InitiallyUnowned: externglib.InitiallyUnowned{
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
	}
}

func marshalColorSelectionDialog(p uintptr) (interface{}, error) {
	return wrapColorSelectionDialog(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewColorSelectionDialog creates a new ColorSelectionDialog.
//
// The function takes the following parameters:
//
//    - title: string containing the title text for the dialog.
//
// The function returns the following values:
//
//    - colorSelectionDialog: ColorSelectionDialog.
//
func NewColorSelectionDialog(title string) *ColorSelectionDialog {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_color_selection_dialog_new(_arg1)
	runtime.KeepAlive(title)

	var _colorSelectionDialog *ColorSelectionDialog // out

	_colorSelectionDialog = wrapColorSelectionDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _colorSelectionDialog
}

// ColorSelection retrieves the ColorSelection widget embedded in the dialog.
//
// The function returns the following values:
//
//    - widget: embedded ColorSelection.
//
func (colorsel *ColorSelectionDialog) ColorSelection() Widgetter {
	var _arg0 *C.GtkColorSelectionDialog // out
	var _cret *C.GtkWidget               // in

	_arg0 = (*C.GtkColorSelectionDialog)(unsafe.Pointer(externglib.InternObject(colorsel).Native()))

	_cret = C.gtk_color_selection_dialog_get_color_selection(_arg0)
	runtime.KeepAlive(colorsel)

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
