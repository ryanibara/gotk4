// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_hbutton_box_get_type()), F: marshalHButtonBoxxer},
	})
}

// HButtonBoxxer describes HButtonBox's methods.
type HButtonBoxxer interface {
	privateHButtonBox()
}

type HButtonBox struct {
	ButtonBox
}

var (
	_ HButtonBoxxer   = (*HButtonBox)(nil)
	_ gextras.Nativer = (*HButtonBox)(nil)
)

func wrapHButtonBox(obj *externglib.Object) HButtonBoxxer {
	return &HButtonBox{
		ButtonBox: ButtonBox{
			Box: Box{
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
					},
				},
				Orientable: Orientable{
					Object: obj,
				},
			},
		},
	}
}

func marshalHButtonBoxxer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapHButtonBox(obj), nil
}

// NewHButtonBox creates a new horizontal button box.
//
// Deprecated: Use gtk_button_box_new() with GTK_ORIENTATION_HORIZONTAL instead.
func NewHButtonBox() *HButtonBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_hbutton_box_new()

	var _hButtonBox *HButtonBox // out

	_hButtonBox = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*HButtonBox)

	return _hButtonBox
}

func (*HButtonBox) privateHButtonBox() {}
