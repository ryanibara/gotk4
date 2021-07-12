// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_color_chooser_dialog_get_type()), F: marshalColorChooserDialoger},
	})
}

// ColorChooserDialoger describes ColorChooserDialog's methods.
type ColorChooserDialoger interface {
	privateColorChooserDialog()
}

// ColorChooserDialog: dialog for choosing a color.
//
// !An example GtkColorChooserDialog (colorchooser.png)
//
// `GtkColorChooserDialog` implements the [iface@Gtk.ColorChooser] interface and
// does not provide much API of its own.
//
// To create a `GtkColorChooserDialog`, use [ctor@Gtk.ColorChooserDialog.new].
//
// To change the initially selected color, use
// [method@Gtk.ColorChooser.set_rgba]. To get the selected color use
// [method@Gtk.ColorChooser.get_rgba].
type ColorChooserDialog struct {
	Dialog

	ColorChooser
}

var (
	_ ColorChooserDialoger = (*ColorChooserDialog)(nil)
	_ gextras.Nativer      = (*ColorChooserDialog)(nil)
)

func wrapColorChooserDialog(obj *externglib.Object) *ColorChooserDialog {
	return &ColorChooserDialog{
		Dialog: Dialog{
			Window: Window{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
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
				Root: Root{
					NativeSurface: NativeSurface{
						Widget: Widget{
							InitiallyUnowned: externglib.InitiallyUnowned{
								Object: obj,
							},
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
		ColorChooser: ColorChooser{
			Object: obj,
		},
	}
}

func marshalColorChooserDialoger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapColorChooserDialog(obj), nil
}

// NewColorChooserDialog creates a new `GtkColorChooserDialog`.
func NewColorChooserDialog(title string, parent Windower) *ColorChooserDialog {
	var _arg1 *C.char      // out
	var _arg2 *C.GtkWindow // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	_arg2 = (*C.GtkWindow)(unsafe.Pointer((parent).(gextras.Nativer).Native()))

	_cret = C.gtk_color_chooser_dialog_new(_arg1, _arg2)

	var _colorChooserDialog *ColorChooserDialog // out

	_colorChooserDialog = wrapColorChooserDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _colorChooserDialog
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *ColorChooserDialog) Native() uintptr {
	return v.Dialog.Window.Widget.InitiallyUnowned.Object.Native()
}

func (*ColorChooserDialog) privateColorChooserDialog() {}
