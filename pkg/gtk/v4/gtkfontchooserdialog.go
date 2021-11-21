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
		{T: externglib.Type(C.gtk_font_chooser_dialog_get_type()), F: marshalFontChooserDialogger},
	})
}

// FontChooserDialog: GtkFontChooserDialog widget is a dialog for selecting a
// font.
//
// !An example GtkFontChooserDialog (fontchooser.png)
//
// GtkFontChooserDialog implements the gtk.FontChooser interface and does not
// provide much API of its own.
//
// To create a GtkFontChooserDialog, use gtk.FontChooserDialog.New.
//
//
// GtkFontChooserDialog as GtkBuildable
//
// The GtkFontChooserDialog implementation of the GtkBuildable interface exposes
// the buttons with the names “select_button” and “cancel_button”.
type FontChooserDialog struct {
	Dialog

	*externglib.Object
	FontChooser
}

var (
	_ externglib.Objector = (*FontChooserDialog)(nil)
	_ Widgetter           = (*FontChooserDialog)(nil)
)

func wrapFontChooserDialog(obj *externglib.Object) *FontChooserDialog {
	return &FontChooserDialog{
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
		FontChooser: FontChooser{
			Object: obj,
		},
	}
}

func marshalFontChooserDialogger(p uintptr) (interface{}, error) {
	return wrapFontChooserDialog(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFontChooserDialog creates a new GtkFontChooserDialog.
//
// The function takes the following parameters:
//
//    - title: title of the dialog, or NULL.
//    - parent: transient parent of the dialog, or NULL.
//
func NewFontChooserDialog(title string, parent *Window) *FontChooserDialog {
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

	_cret = C.gtk_font_chooser_dialog_new(_arg1, _arg2)
	runtime.KeepAlive(title)
	runtime.KeepAlive(parent)

	var _fontChooserDialog *FontChooserDialog // out

	_fontChooserDialog = wrapFontChooserDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _fontChooserDialog
}
