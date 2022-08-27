// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeColorSelectionDialog = coreglib.Type(C.gtk_color_selection_dialog_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeColorSelectionDialog, F: marshalColorSelectionDialog},
	})
}

// ColorSelectionDialogOverrides contains methods that are overridable.
type ColorSelectionDialogOverrides struct {
}

func defaultColorSelectionDialogOverrides(v *ColorSelectionDialog) ColorSelectionDialogOverrides {
	return ColorSelectionDialogOverrides{}
}

type ColorSelectionDialog struct {
	_ [0]func() // equal guard
	Dialog
}

var (
	_ Binner = (*ColorSelectionDialog)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ColorSelectionDialog, *ColorSelectionDialogClass, ColorSelectionDialogOverrides](
		GTypeColorSelectionDialog,
		initColorSelectionDialogClass,
		wrapColorSelectionDialog,
		defaultColorSelectionDialogOverrides,
	)
}

func initColorSelectionDialogClass(gclass unsafe.Pointer, overrides ColorSelectionDialogOverrides, classInitFunc func(*ColorSelectionDialogClass)) {
	if classInitFunc != nil {
		class := (*ColorSelectionDialogClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapColorSelectionDialog(obj *coreglib.Object) *ColorSelectionDialog {
	return &ColorSelectionDialog{
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
	}
}

func marshalColorSelectionDialog(p uintptr) (interface{}, error) {
	return wrapColorSelectionDialog(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
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

	_colorSelectionDialog = wrapColorSelectionDialog(coreglib.Take(unsafe.Pointer(_cret)))

	return _colorSelectionDialog
}

// ColorSelectionDialogClass: instance of this type is always passed by
// reference.
type ColorSelectionDialogClass struct {
	*colorSelectionDialogClass
}

// colorSelectionDialogClass is the struct that's finalized.
type colorSelectionDialogClass struct {
	native *C.GtkColorSelectionDialogClass
}

func (c *ColorSelectionDialogClass) ParentClass() *DialogClass {
	valptr := &c.native.parent_class
	var _v *DialogClass // out
	_v = (*DialogClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
