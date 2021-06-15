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
		{T: externglib.Type(C.gtk_viewport_get_type()), F: marshalViewport},
	})
}

// Viewport: `GtkViewport` implements scrollability for widgets that lack their
// own scrolling capabilities.
//
// Use `GtkViewport` to scroll child widgets such as `GtkGrid`, `GtkBox`, and so
// on.
//
// The `GtkViewport` will start scrolling content only if allocated less than
// the child widget’s minimum size in a given orientation.
//
//
// CSS nodes
//
// `GtkViewport` has a single CSS node with name `viewport`.
//
//
// Accessibility
//
// `GtkViewport` uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type Viewport interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Scrollable

	// Child gets the child widget of @viewport.
	Child() Widget
	// ScrollToFocus gets whether the viewport is scrolling to keep the focused
	// child in view.
	ScrollToFocus() bool
	// SetChild sets the child widget of @viewport.
	SetChild(child Widget)
	// SetScrollToFocus sets whether the viewport should automatically scroll to
	// keep the focused child in view.
	SetScrollToFocus(scrollToFocus bool)
}

// viewport implements the Viewport class.
type viewport struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Scrollable
}

var _ Viewport = (*viewport)(nil)

// WrapViewport wraps a GObject to the right type. It is
// primarily used internally.
func WrapViewport(obj *externglib.Object) Viewport {
	return viewport{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Scrollable:       WrapScrollable(obj),
	}
}

func marshalViewport(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapViewport(obj), nil
}

// NewViewport constructs a class Viewport.
func NewViewport(hadjustment Adjustment, vadjustment Adjustment) Viewport {
	var _arg1 *C.GtkAdjustment // out
	var _arg2 *C.GtkAdjustment // out
	var _cret C.GtkViewport    // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))
	_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	_cret = C.gtk_viewport_new(_arg1, _arg2)

	var _viewport Viewport // out

	_viewport = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Viewport)

	return _viewport
}

// Child gets the child widget of @viewport.
func (v viewport) Child() Widget {
	var _arg0 *C.GtkViewport // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// ScrollToFocus gets whether the viewport is scrolling to keep the focused
// child in view.
func (v viewport) ScrollToFocus() bool {
	var _arg0 *C.GtkViewport // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_scroll_to_focus(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetChild sets the child widget of @viewport.
func (v viewport) SetChild(child Widget) {
	var _arg0 *C.GtkViewport // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_viewport_set_child(_arg0, _arg1)
}

// SetScrollToFocus sets whether the viewport should automatically scroll to
// keep the focused child in view.
func (v viewport) SetScrollToFocus(scrollToFocus bool) {
	var _arg0 *C.GtkViewport // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))
	if scrollToFocus {
		_arg1 = C.TRUE
	}

	C.gtk_viewport_set_scroll_to_focus(_arg0, _arg1)
}
