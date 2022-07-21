// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GTypeOrientable returns the GType for the type Orientable.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeOrientable() coreglib.Type {
	gtype := coreglib.Type(C.gtk_orientable_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalOrientable)
	return gtype
}

// OrientableOverrider contains methods that are overridable.
type OrientableOverrider interface {
}

// Orientable: GtkOrientable interface is implemented by all widgets that can be
// oriented horizontally or vertically.
//
// GtkOrientable is more flexible in that it allows the orientation to be
// changed at runtime, allowing the widgets to “flip”.
//
// Orientable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Orientable struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Orientable)(nil)
)

// Orientabler describes Orientable's interface methods.
type Orientabler interface {
	coreglib.Objector

	// Orientation retrieves the orientation of the orientable.
	Orientation() Orientation
	// SetOrientation sets the orientation of the orientable.
	SetOrientation(orientation Orientation)
}

var _ Orientabler = (*Orientable)(nil)

func ifaceInitOrientabler(gifacePtr, data C.gpointer) {
}

func wrapOrientable(obj *coreglib.Object) *Orientable {
	return &Orientable{
		Object: obj,
	}
}

func marshalOrientable(p uintptr) (interface{}, error) {
	return wrapOrientable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Orientation retrieves the orientation of the orientable.
//
// The function returns the following values:
//
//    - orientation of the orientable.
//
func (orientable *Orientable) Orientation() Orientation {
	var _arg0 *C.GtkOrientable // out
	var _cret C.GtkOrientation // in

	_arg0 = (*C.GtkOrientable)(unsafe.Pointer(coreglib.InternObject(orientable).Native()))

	_cret = C.gtk_orientable_get_orientation(_arg0)
	runtime.KeepAlive(orientable)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

// SetOrientation sets the orientation of the orientable.
//
// The function takes the following parameters:
//
//    - orientation orientable’s new orientation.
//
func (orientable *Orientable) SetOrientation(orientation Orientation) {
	var _arg0 *C.GtkOrientable // out
	var _arg1 C.GtkOrientation // out

	_arg0 = (*C.GtkOrientable)(unsafe.Pointer(coreglib.InternObject(orientable).Native()))
	_arg1 = C.GtkOrientation(orientation)

	C.gtk_orientable_set_orientation(_arg0, _arg1)
	runtime.KeepAlive(orientable)
	runtime.KeepAlive(orientation)
}
