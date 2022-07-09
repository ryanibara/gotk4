// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeColorChooserDialog returns the GType for the type ColorChooserDialog.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeColorChooserDialog() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ColorChooserDialog").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalColorChooserDialog)
	return gtype
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
	_ [0]func() // equal guard
	Dialog

	*coreglib.Object
	ColorChooser
}

var (
	_ coreglib.Objector = (*ColorChooserDialog)(nil)
	_ Widgetter         = (*ColorChooserDialog)(nil)
)

func wrapColorChooserDialog(obj *coreglib.Object) *ColorChooserDialog {
	return &ColorChooserDialog{
		Dialog: Dialog{
			Window: Window{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
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
							InitiallyUnowned: coreglib.InitiallyUnowned{
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

func marshalColorChooserDialog(p uintptr) (interface{}, error) {
	return wrapColorChooserDialog(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewColorChooserDialog creates a new GtkColorChooserDialog.
//
// The function takes the following parameters:
//
//    - title (optional): title of the dialog, or NULL.
//    - parent (optional): transient parent of the dialog, or NULL.
//
// The function returns the following values:
//
//    - colorChooserDialog: new GtkColorChooserDialog.
//
func NewColorChooserDialog(title string, parent *Window) *ColorChooserDialog {
	var _args [2]girepository.Argument

	if title != "" {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_args[0]))
	}
	if parent != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	}

	_gret := girepository.MustFind("Gtk", "ColorChooserDialog").InvokeMethod("new_ColorChooserDialog", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(title)
	runtime.KeepAlive(parent)

	var _colorChooserDialog *ColorChooserDialog // out

	_colorChooserDialog = wrapColorChooserDialog(coreglib.Take(unsafe.Pointer(_cret)))

	return _colorChooserDialog
}
