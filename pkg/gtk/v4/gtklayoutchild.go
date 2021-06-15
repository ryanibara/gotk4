// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_layout_child_get_type()), F: marshalLayoutChild},
	})
}

// LayoutChild: `GtkLayoutChild` is the base class for objects that are meant to
// hold layout properties.
//
// If a `GtkLayoutManager` has per-child properties, like their packing type, or
// the horizontal and vertical span, or the icon name, then the layout manager
// should use a `GtkLayoutChild` implementation to store those properties.
//
// A `GtkLayoutChild` instance is only ever valid while a widget is part of a
// layout.
type LayoutChild interface {
	gextras.Objector

	// ChildWidget retrieves the `GtkWidget` associated to the given
	// @layout_child.
	ChildWidget() Widget
	// LayoutManager retrieves the `GtkLayoutManager` instance that created the
	// given @layout_child.
	LayoutManager() LayoutManager
}

// layoutChild implements the LayoutChild class.
type layoutChild struct {
	gextras.Objector
}

var _ LayoutChild = (*layoutChild)(nil)

// WrapLayoutChild wraps a GObject to the right type. It is
// primarily used internally.
func WrapLayoutChild(obj *externglib.Object) LayoutChild {
	return layoutChild{
		Objector: obj,
	}
}

func marshalLayoutChild(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLayoutChild(obj), nil
}

// ChildWidget retrieves the `GtkWidget` associated to the given
// @layout_child.
func (l layoutChild) ChildWidget() Widget {
	var _arg0 *C.GtkLayoutChild // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkLayoutChild)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_layout_child_get_child_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// LayoutManager retrieves the `GtkLayoutManager` instance that created the
// given @layout_child.
func (l layoutChild) LayoutManager() LayoutManager {
	var _arg0 *C.GtkLayoutChild   // out
	var _cret *C.GtkLayoutManager // in

	_arg0 = (*C.GtkLayoutChild)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_layout_child_get_layout_manager(_arg0)

	var _layoutManager LayoutManager // out

	_layoutManager = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(LayoutManager)

	return _layoutManager
}
