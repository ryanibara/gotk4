// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gsk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_fixed_get_type()), F: marshalFixeder},
	})
}

// Fixeder describes Fixed's methods.
type Fixeder interface {
	// ChildPosition retrieves the translation transformation of the given child
	// `GtkWidget` in the `GtkFixed`.
	ChildPosition(widget Widgeter) (x float64, y float64)
	// ChildTransform retrieves the transformation for @widget set using
	// gtk_fixed_set_child_transform().
	ChildTransform(widget Widgeter) *gsk.Transform
	// Move sets a translation transformation to the given @x and @y coordinates
	// to the child @widget of the `GtkFixed`.
	Move(widget Widgeter, x float64, y float64)
	// Put adds a widget to a `GtkFixed` at the given position.
	Put(widget Widgeter, x float64, y float64)
	// Remove removes a child from @fixed.
	Remove(widget Widgeter)
	// SetChildTransform sets the transformation for @widget.
	SetChildTransform(widget Widgeter, transform *gsk.Transform)
}

// Fixed: `GtkFixed` places its child widgets at fixed positions and with fixed
// sizes.
//
// `GtkFixed` performs no automatic layout management.
//
// For most applications, you should not use this container! It keeps you from
// having to learn about the other GTK containers, but it results in broken
// applications. With `GtkFixed`, the following things will result in truncated
// text, overlapping widgets, and other display bugs:
//
// - Themes, which may change widget sizes.
//
// - Fonts other than the one you used to write the app will of course change
// the size of widgets containing text; keep in mind that users may use a larger
// font because of difficulty reading the default, or they may be using a
// different OS that provides different fonts.
//
// - Translation of text into other languages changes its size. Also, display of
// non-English text will use a different font in many cases.
//
// In addition, `GtkFixed` does not pay attention to text direction and thus may
// produce unwanted results if your app is run under right-to-left languages
// such as Hebrew or Arabic. That is: normally GTK will order containers
// appropriately for the text direction, e.g. to put labels to the right of the
// thing they label when using an RTL language, but it can’t do that with
// `GtkFixed`. So if you need to reorder widgets depending on the text
// direction, you would need to manually detect it and adjust child positions
// accordingly.
//
// Finally, fixed positioning makes it kind of annoying to add/remove UI
// elements, since you have to reposition all the other elements. This is a
// long-term maintenance problem for your application.
//
// If you know none of these things are an issue for your application, and
// prefer the simplicity of `GtkFixed`, by all means use the widget. But you
// should be aware of the tradeoffs.
type Fixed struct {
	Widget
}

var (
	_ Fixeder         = (*Fixed)(nil)
	_ gextras.Nativer = (*Fixed)(nil)
)

func wrapFixed(obj *externglib.Object) *Fixed {
	return &Fixed{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
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
	}
}

func marshalFixeder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFixed(obj), nil
}

// NewFixed creates a new `GtkFixed`.
func NewFixed() *Fixed {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_fixed_new()

	var _fixed *Fixed // out

	_fixed = wrapFixed(externglib.Take(unsafe.Pointer(_cret)))

	return _fixed
}

// ChildPosition retrieves the translation transformation of the given child
// `GtkWidget` in the `GtkFixed`.
//
// See also: [method@Gtk.Fixed.get_child_transform].
func (fixed *Fixed) ChildPosition(widget Widgeter) (x float64, y float64) {
	var _arg0 *C.GtkFixed  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.double     // in
	var _arg3 C.double     // in

	_arg0 = (*C.GtkFixed)(unsafe.Pointer(fixed.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))

	C.gtk_fixed_get_child_position(_arg0, _arg1, &_arg2, &_arg3)

	var _x float64 // out
	var _y float64 // out

	_x = float64(_arg2)
	_y = float64(_arg3)

	return _x, _y
}

// ChildTransform retrieves the transformation for @widget set using
// gtk_fixed_set_child_transform().
func (fixed *Fixed) ChildTransform(widget Widgeter) *gsk.Transform {
	var _arg0 *C.GtkFixed     // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.GskTransform // in

	_arg0 = (*C.GtkFixed)(unsafe.Pointer(fixed.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))

	_cret = C.gtk_fixed_get_child_transform(_arg0, _arg1)

	var _transform *gsk.Transform // out

	_transform = (*gsk.Transform)(unsafe.Pointer(_cret))
	C.gsk_transform_ref(_cret)
	runtime.SetFinalizer(_transform, func(v *gsk.Transform) {
		C.gsk_transform_unref((*C.GskTransform)(unsafe.Pointer(v)))
	})

	return _transform
}

// Move sets a translation transformation to the given @x and @y coordinates to
// the child @widget of the `GtkFixed`.
func (fixed *Fixed) Move(widget Widgeter, x float64, y float64) {
	var _arg0 *C.GtkFixed  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.double     // out
	var _arg3 C.double     // out

	_arg0 = (*C.GtkFixed)(unsafe.Pointer(fixed.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)

	C.gtk_fixed_move(_arg0, _arg1, _arg2, _arg3)
}

// Put adds a widget to a `GtkFixed` at the given position.
func (fixed *Fixed) Put(widget Widgeter, x float64, y float64) {
	var _arg0 *C.GtkFixed  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.double     // out
	var _arg3 C.double     // out

	_arg0 = (*C.GtkFixed)(unsafe.Pointer(fixed.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)

	C.gtk_fixed_put(_arg0, _arg1, _arg2, _arg3)
}

// Remove removes a child from @fixed.
func (fixed *Fixed) Remove(widget Widgeter) {
	var _arg0 *C.GtkFixed  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkFixed)(unsafe.Pointer(fixed.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))

	C.gtk_fixed_remove(_arg0, _arg1)
}

// SetChildTransform sets the transformation for @widget.
//
// This is a convenience function that retrieves the
// [class@Gtk.FixedLayoutChild] instance associated to @widget and calls
// [method@Gtk.FixedLayoutChild.set_transform].
func (fixed *Fixed) SetChildTransform(widget Widgeter, transform *gsk.Transform) {
	var _arg0 *C.GtkFixed     // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.GskTransform // out

	_arg0 = (*C.GtkFixed)(unsafe.Pointer(fixed.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))
	_arg2 = (*C.GskTransform)(unsafe.Pointer(transform))

	C.gtk_fixed_set_child_transform(_arg0, _arg1, _arg2)
}
