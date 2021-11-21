// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_color_chooser_dialog_get_type()), F: marshalColorChooserDialogger},
	})
}

// ColorChooserDialog: dialog for choosing a color.
//
// !An example GtkColorChooserDialog (colorchooser.png)
//
// GtkColorChooserDialog implements the gtk.ColorChooser interface and does not
// provide much API of its own.
//
// To create a GtkColorChooserDialog, use gtk.ColorChooserDialog.New.
//
// To change the initially selected color, use gtk.ColorChooser.SetRGBA(). To
// get the selected color use gtk.ColorChooser.GetRGBA().
type ColorChooserDialog struct {
	Dialog

	*externglib.Object
	ColorChooser
}

var (
	_ externglib.Objector = (*ColorChooserDialog)(nil)
	_ Widgetter           = (*ColorChooserDialog)(nil)
)

func wrapColorChooserDialog(obj *externglib.Object) *ColorChooserDialog {
	return &ColorChooserDialog{
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
		ColorChooser: ColorChooser{
			Object: obj,
		},
	}
}

func marshalColorChooserDialogger(p uintptr) (interface{}, error) {
	return wrapColorChooserDialog(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewColorChooserDialog creates a new GtkColorChooserDialog.
//
// The function takes the following parameters:
//
//    - title: title of the dialog, or NULL.
//    - parent: transient parent of the dialog, or NULL.
//
func NewColorChooserDialog(title string, parent *Window) *ColorChooserDialog {
	var _arg1 *C.char      // out
	var _arg2 *C.GtkWindow // out
	var _cret *C.GtkWidget // in

	if title != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	if parent != nil {
		_arg2 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	}

	_cret = C.gtk_color_chooser_dialog_new(_arg1, _arg2)
	runtime.KeepAlive(title)
	runtime.KeepAlive(parent)

	var _colorChooserDialog *ColorChooserDialog // out

	_colorChooserDialog = wrapColorChooserDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _colorChooserDialog
}
