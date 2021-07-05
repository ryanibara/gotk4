// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_vbox_get_type()), F: marshalVBox},
	})
}

// VBox is a container that organizes child widgets into a single column.
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

	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsOrientable casts the class to the Orientable interface.
	AsOrientable() Orientable
}

// vBox implements the VBox class.
type vBox struct {
	Box
}

// WrapVBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapVBox(obj *externglib.Object) VBox {
	return vBox{
		Box: WrapBox(obj),
	}
}

func marshalVBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVBox(obj), nil
}

// NewVBox creates a new VBox.
//
// Deprecated: since version 3.2.
func NewVBox(homogeneous bool, spacing int) VBox {
	var _arg1 C.gboolean   // out
	var _arg2 C.gint       // out
	var _cret *C.GtkWidget // in

	if homogeneous {
		_arg1 = C.TRUE
	}
	_arg2 = C.gint(spacing)

	_cret = C.gtk_vbox_new(_arg1, _arg2)

	var _vBox VBox // out

	_vBox = WrapVBox(externglib.Take(unsafe.Pointer(_cret)))

	return _vBox
}

func (v vBox) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(v))
}

func (v vBox) AsOrientable() Orientable {
	return WrapOrientable(gextras.InternObject(v))
}
