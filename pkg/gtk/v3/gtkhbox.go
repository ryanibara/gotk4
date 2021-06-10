// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_hbox_get_type()), F: marshalHBox},
	})
}

// HBox is a container that organizes child widgets into a single row.
//
// Use the Box packing interface to determine the arrangement, spacing, width,
// and alignment of HBox children.
//
// All children are allocated the same height.
//
// GtkHBox has been deprecated. You can use Box instead, which is a very quick
// and easy change. If you have derived your own classes from GtkHBox, you can
// simply change the inheritance to derive directly from Box. No further changes
// are needed, since the default value of the Orientable:orientation property is
// GTK_ORIENTATION_HORIZONTAL.
//
// If you have a grid-like layout composed of nested boxes, and you don’t need
// first-child or last-child styling, the recommendation is to switch to Grid.
// For more information about migrating to Grid, see [Migrating from other
// containers to GtkGrid][gtk-migrating-GtkGrid].
type HBox interface {
	Box
	Buildable
	Orientable
}

// hBox implements the HBox interface.
type hBox struct {
	Box
	Buildable
	Orientable
}

var _ HBox = (*hBox)(nil)

// WrapHBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapHBox(obj *externglib.Object) HBox {
	return HBox{
		Box:        WrapBox(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalHBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapHBox(obj), nil
}
