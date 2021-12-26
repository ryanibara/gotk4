// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gsk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_fixed_get_type()), F: marshalFixedder},
	})
}

// Fixed: GtkFixed places its child widgets at fixed positions and with fixed
// sizes.
//
// GtkFixed performs no automatic layout management.
//
// For most applications, you should not use this container! It keeps you from
// having to learn about the other GTK containers, but it results in broken
// applications. With GtkFixed, the following things will result in truncated
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
// In addition, GtkFixed does not pay attention to text direction and thus may
// produce unwanted results if your app is run under right-to-left languages
// such as Hebrew or Arabic. That is: normally GTK will order containers
// appropriately for the text direction, e.g. to put labels to the right of the
// thing they label when using an RTL language, but it can’t do that with
// GtkFixed. So if you need to reorder widgets depending on the text direction,
// you would need to manually detect it and adjust child positions accordingly.
//
// Finally, fixed positioning makes it kind of annoying to add/remove UI
// elements, since you have to reposition all the other elements. This is a
// long-term maintenance problem for your application.
//
// If you know none of these things are an issue for your application, and
// prefer the simplicity of GtkFixed, by all means use the widget. But you
// should be aware of the tradeoffs.
type Fixed struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Fixed)(nil)
)

func wrapFixed(obj *externglib.Object) *Fixed {
	return &Fixed{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
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

func marshalFixedder(p uintptr) (interface{}, error) {
	return wrapFixed(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFixed creates a new GtkFixed.
//
// The function returns the following values:
//
//    - fixed: new GtkFixed.
//
func NewFixed() *Fixed {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_fixed_new()

	var _fixed *Fixed // out

	_fixed = wrapFixed(externglib.Take(unsafe.Pointer(_cret)))

	return _fixed
}

// ChildPosition retrieves the translation transformation of the given child
// GtkWidget in the GtkFixed.
//
// See also: gtk.Fixed.GetChildTransform().
//
// The function takes the following parameters:
//
//    - widget: child of fixed.
//
// The function returns the following values:
//
//    - x: horizontal position of the widget.
//    - y: vertical position of the widget.
//
func (fixed *Fixed) ChildPosition(widget Widgetter) (x float64, y float64) {
	var _arg0 *C.GtkFixed  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.double     // in
	var _arg3 C.double     // in

	_arg0 = (*C.GtkFixed)(unsafe.Pointer(fixed.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_fixed_get_child_position(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(fixed)
	runtime.KeepAlive(widget)

	var _x float64 // out
	var _y float64 // out

	_x = float64(_arg2)
	_y = float64(_arg3)

	return _x, _y
}

// ChildTransform retrieves the transformation for widget set using
// gtk_fixed_set_child_transform().
//
// The function takes the following parameters:
//
//    - widget: GtkWidget, child of fixed.
//
// The function returns the following values:
//
//    - transform (optional): GskTransform or NULL in case no transform has been
//      set on widget.
//
func (fixed *Fixed) ChildTransform(widget Widgetter) *gsk.Transform {
	var _arg0 *C.GtkFixed     // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.GskTransform // in

	_arg0 = (*C.GtkFixed)(unsafe.Pointer(fixed.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_fixed_get_child_transform(_arg0, _arg1)
	runtime.KeepAlive(fixed)
	runtime.KeepAlive(widget)

	var _transform *gsk.Transform // out

	if _cret != nil {
		_transform = (*gsk.Transform)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.gsk_transform_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_transform)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gsk_transform_unref((*C.GskTransform)(intern.C))
			},
		)
	}

	return _transform
}

// Move sets a translation transformation to the given x and y coordinates to
// the child widget of the GtkFixed.
//
// The function takes the following parameters:
//
//    - widget: child widget.
//    - x: horizontal position to move the widget to.
//    - y: vertical position to move the widget to.
//
func (fixed *Fixed) Move(widget Widgetter, x, y float64) {
	var _arg0 *C.GtkFixed  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.double     // out
	var _arg3 C.double     // out

	_arg0 = (*C.GtkFixed)(unsafe.Pointer(fixed.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)

	C.gtk_fixed_move(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(fixed)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// Put adds a widget to a GtkFixed at the given position.
//
// The function takes the following parameters:
//
//    - widget to add.
//    - x: horizontal position to place the widget at.
//    - y: vertical position to place the widget at.
//
func (fixed *Fixed) Put(widget Widgetter, x, y float64) {
	var _arg0 *C.GtkFixed  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.double     // out
	var _arg3 C.double     // out

	_arg0 = (*C.GtkFixed)(unsafe.Pointer(fixed.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)

	C.gtk_fixed_put(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(fixed)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// Remove removes a child from fixed.
//
// The function takes the following parameters:
//
//    - widget: child widget to remove.
//
func (fixed *Fixed) Remove(widget Widgetter) {
	var _arg0 *C.GtkFixed  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkFixed)(unsafe.Pointer(fixed.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_fixed_remove(_arg0, _arg1)
	runtime.KeepAlive(fixed)
	runtime.KeepAlive(widget)
}

// SetChildTransform sets the transformation for widget.
//
// This is a convenience function that retrieves the gtk.FixedLayoutChild
// instance associated to widget and calls gtk.FixedLayoutChild.SetTransform().
//
// The function takes the following parameters:
//
//    - widget: GtkWidget, child of fixed.
//    - transform (optional): transformation assigned to widget or NULL to reset
//      widget's transform.
//
func (fixed *Fixed) SetChildTransform(widget Widgetter, transform *gsk.Transform) {
	var _arg0 *C.GtkFixed     // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.GskTransform // out

	_arg0 = (*C.GtkFixed)(unsafe.Pointer(fixed.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	if transform != nil {
		_arg2 = (*C.GskTransform)(gextras.StructNative(unsafe.Pointer(transform)))
	}

	C.gtk_fixed_set_child_transform(_arg0, _arg1, _arg2)
	runtime.KeepAlive(fixed)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(transform)
}
