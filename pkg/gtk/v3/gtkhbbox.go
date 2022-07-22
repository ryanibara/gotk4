// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
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
	GTypeHButtonBox = coreglib.Type(C.gtk_hbutton_box_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeHButtonBox, F: marshalHButtonBox},
	})
}

// HButtonBoxOverrider contains methods that are overridable.
type HButtonBoxOverrider interface {
}

type HButtonBox struct {
	_ [0]func() // equal guard
	ButtonBox
}

var (
	_ Containerer       = (*HButtonBox)(nil)
	_ coreglib.Objector = (*HButtonBox)(nil)
)

func classInitHButtonBoxer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapHButtonBox(obj *coreglib.Object) *HButtonBox {
	return &HButtonBox{
		ButtonBox: ButtonBox{
			Box: Box{
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
				Object: obj,
				Orientable: Orientable{
					Object: obj,
				},
			},
		},
	}
}

func marshalHButtonBox(p uintptr) (interface{}, error) {
	return wrapHButtonBox(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewHButtonBox creates a new horizontal button box.
//
// Deprecated: Use gtk_button_box_new() with GTK_ORIENTATION_HORIZONTAL instead.
//
// The function returns the following values:
//
//    - hButtonBox: new button box Widget.
//
func NewHButtonBox() *HButtonBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_hbutton_box_new()

	var _hButtonBox *HButtonBox // out

	_hButtonBox = wrapHButtonBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _hButtonBox
}
