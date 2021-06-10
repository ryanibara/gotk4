// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_scrollable_get_type()), F: marshalScrollable},
	})
}

// ScrollableOverrider contains methods that are overridable. This
// interface is a subset of the interface Scrollable.
type ScrollableOverrider interface {
	// Border returns the size of a non-scrolling border around the outside of
	// the scrollable. An example for this would be treeview headers. GTK+ can
	// use this information to display overlayed graphics, like the overshoot
	// indication, at the right position.
	Border() (Border, bool)
}

// Scrollable is an interface that is implemented by widgets with native
// scrolling ability.
//
// To implement this interface you should override the Scrollable:hadjustment
// and Scrollable:vadjustment properties.
//
//
// Creating a scrollable widget
//
// All scrollable widgets should do the following.
//
// - When a parent widget sets the scrollable child widget’s adjustments, the
// widget should populate the adjustments’ Adjustment:lower, Adjustment:upper,
// Adjustment:step-increment, Adjustment:page-increment and Adjustment:page-size
// properties and connect to the Adjustment::value-changed signal.
//
// - Because its preferred size is the size for a fully expanded widget, the
// scrollable widget must be able to cope with underallocations. This means that
// it must accept any value passed to its WidgetClass.size_allocate() function.
//
// - When the parent allocates space to the scrollable child widget, the widget
// should update the adjustments’ properties with new values.
//
// - When any of the adjustments emits the Adjustment::value-changed signal, the
// scrollable widget should scroll its contents.
type Scrollable interface {
	gextras.Objector
	ScrollableOverrider

	// SetHAdjustment sets the horizontal adjustment of the Scrollable.
	SetHAdjustment(hadjustment Adjustment)
	// SetHScrollPolicy sets the ScrollablePolicy to determine whether
	// horizontal scrolling should start below the minimum width or below the
	// natural width.
	SetHScrollPolicy(policy ScrollablePolicy)
	// SetVAdjustment sets the vertical adjustment of the Scrollable.
	SetVAdjustment(vadjustment Adjustment)
	// SetVScrollPolicy sets the ScrollablePolicy to determine whether vertical
	// scrolling should start below the minimum height or below the natural
	// height.
	SetVScrollPolicy(policy ScrollablePolicy)
}

// scrollable implements the Scrollable interface.
type scrollable struct {
	gextras.Objector
}

var _ Scrollable = (*scrollable)(nil)

// WrapScrollable wraps a GObject to a type that implements interface
// Scrollable. It is primarily used internally.
func WrapScrollable(obj *externglib.Object) Scrollable {
	return Scrollable{
		Objector: obj,
	}
}

func marshalScrollable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapScrollable(obj), nil
}

// Border returns the size of a non-scrolling border around the outside of
// the scrollable. An example for this would be treeview headers. GTK+ can
// use this information to display overlayed graphics, like the overshoot
// indication, at the right position.
func (s scrollable) Border() (Border, bool) {
	var _arg0 *C.GtkScrollable

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(s.Native()))

	var _border Border
	var _cret C.gboolean

	_cret = C.gtk_scrollable_get_border(_arg0, (*C.GtkBorder)(unsafe.Pointer(&_border)))

	var _ok bool

	if _cret {
		_ok = true
	}

	return _border, _ok
}

// SetHAdjustment sets the horizontal adjustment of the Scrollable.
func (s scrollable) SetHAdjustment(hadjustment Adjustment) {
	var _arg0 *C.GtkScrollable
	var _arg1 *C.GtkAdjustment

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))

	C.gtk_scrollable_set_hadjustment(_arg0, _arg1)
}

// SetHScrollPolicy sets the ScrollablePolicy to determine whether
// horizontal scrolling should start below the minimum width or below the
// natural width.
func (s scrollable) SetHScrollPolicy(policy ScrollablePolicy) {
	var _arg0 *C.GtkScrollable
	var _arg1 C.GtkScrollablePolicy

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkScrollablePolicy)(policy)

	C.gtk_scrollable_set_hscroll_policy(_arg0, _arg1)
}

// SetVAdjustment sets the vertical adjustment of the Scrollable.
func (s scrollable) SetVAdjustment(vadjustment Adjustment) {
	var _arg0 *C.GtkScrollable
	var _arg1 *C.GtkAdjustment

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	C.gtk_scrollable_set_vadjustment(_arg0, _arg1)
}

// SetVScrollPolicy sets the ScrollablePolicy to determine whether vertical
// scrolling should start below the minimum height or below the natural
// height.
func (s scrollable) SetVScrollPolicy(policy ScrollablePolicy) {
	var _arg0 *C.GtkScrollable
	var _arg1 C.GtkScrollablePolicy

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkScrollablePolicy)(policy)

	C.gtk_scrollable_set_vscroll_policy(_arg0, _arg1)
}
