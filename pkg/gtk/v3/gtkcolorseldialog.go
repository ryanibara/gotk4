// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_color_selection_dialog_get_type()), F: marshalColorSelectionDialogger},
	})
}

type ColorSelectionDialog struct {
	Dialog

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ Binner = (*ColorSelectionDialog)(nil)
)

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

func marshalColorSelectionDialogger(p uintptr) (interface{}, error) {
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

	_arg0 = (*C.GtkColorSelectionDialog)(unsafe.Pointer(colorsel.Native()))

	_cret = C.gtk_color_selection_dialog_get_color_selection(_arg0)
	runtime.KeepAlive(colorsel)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}
