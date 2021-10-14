// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
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
		{T: externglib.Type(C.gtk_font_chooser_dialog_get_type()), F: marshalFontChooserDialogger},
	})
}

// FontChooserDialog widget is a dialog for selecting a font. It implements the
// FontChooser interface.
//
//
// GtkFontChooserDialog as GtkBuildable
//
// The GtkFontChooserDialog implementation of the Buildable interface exposes
// the buttons with the names “select_button” and “cancel_button”.
type FontChooserDialog struct {
	Dialog

	FontChooser
	*externglib.Object
}

func wrapFontChooserDialog(obj *externglib.Object) *FontChooserDialog {
	return &FontChooserDialog{
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
							Object: obj,
						},
					},
				},
			},
		},
		FontChooser: FontChooser{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalFontChooserDialogger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFontChooserDialog(obj), nil
}

// NewFontChooserDialog creates a new FontChooserDialog.
//
// The function takes the following parameters:
//
//    - title: title of the dialog, or NULL.
//    - parent: transient parent of the dialog, or NULL.
//
func NewFontChooserDialog(title string, parent *Window) *FontChooserDialog {
	var _arg1 *C.gchar     // out
	var _arg2 *C.GtkWindow // out
	var _cret *C.GtkWidget // in

	if title != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	if parent != nil {
		_arg2 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	}

	_cret = C.gtk_font_chooser_dialog_new(_arg1, _arg2)
	runtime.KeepAlive(title)
	runtime.KeepAlive(parent)

	var _fontChooserDialog *FontChooserDialog // out

	_fontChooserDialog = wrapFontChooserDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _fontChooserDialog
}

func (*FontChooserDialog) privateFontChooserDialog() {}
