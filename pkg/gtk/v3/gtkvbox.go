// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.gtk_vbox_get_type()), F: marshalVBox},
	})
}

// VBox: a VBox is a container that organizes child widgets into a single
// column.
//
// Use the Box packing interface to determine the arrangement, spacing, height,
// and alignment of VBox children.
//
// All children are allocated the same width.
//
// GtkVBox has been deprecated. You can use Box with a Orientable:orientation
// set to GTK_ORIENTATION_VERTICAL instead when calling gtk_box_new(), which is
// a very quick and easy change.
//
// If you have derived your own classes from GtkVBox, you can change the
// inheritance to derive directly from Box, and set the Orientable:orientation
// property to GTK_ORIENTATION_VERTICAL in your instance init function, with a
// call like:
//
//    gtk_orientable_set_orientation (GTK_ORIENTABLE (object),
//                                    GTK_ORIENTATION_VERTICAL);
//
// If you have a grid-like layout composed of nested boxes, and you don’t need
// first-child or last-child styling, the recommendation is to switch to Grid.
// For more information about migrating to Grid, see [Migrating from other
// containers to GtkGrid][gtk-migrating-GtkGrid].
type VBox interface {
	Box
	Buildable
	Orientable
}

// vBox implements the VBox interface.
type vBox struct {
	Box
	Buildable
	Orientable
}

var _ VBox = (*vBox)(nil)

// WrapVBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapVBox(obj *externglib.Object) VBox {
	return VBox{
		Box:        WrapBox(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalVBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVBox(obj), nil
}

// NewVBox constructs a class VBox.
func NewVBox(homogeneous bool, spacing int) VBox {
	var arg1 C.gboolean
	var arg2 C.gint

	if homogeneous {
		arg1 = C.gboolean(1)
	}
	arg2 = C.gint(spacing)

	var cret C.GtkVBox
	var ret1 VBox

	cret = C.gtk_vbox_new(homogeneous, spacing)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(VBox)

	return ret1
}
