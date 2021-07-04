// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
	Scrollable

	Child() Widget

	ScrollToFocus() bool

	SetChildViewport(child Widget)

	SetScrollToFocusViewport(scrollToFocus bool)
}

// viewport implements the Viewport class.
type viewport struct {
	Widget
}

// WrapViewport wraps a GObject to the right type. It is
// primarily used internally.
func WrapViewport(obj *externglib.Object) Viewport {
	return viewport{
		Widget: WrapWidget(obj),
	}
}

func marshalViewport(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapViewport(obj), nil
}

func NewViewport(hadjustment Adjustment, vadjustment Adjustment) Viewport {
	var _arg1 *C.GtkAdjustment // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))
	_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	_cret = C.gtk_viewport_new(_arg1, _arg2)

	var _viewport Viewport // out

	_viewport = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Viewport)

	return _viewport
}

func (v viewport) Child() Widget {
	var _arg0 *C.GtkViewport // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

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

func (v viewport) SetChildViewport(child Widget) {
	var _arg0 *C.GtkViewport // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_viewport_set_child(_arg0, _arg1)
}

func (v viewport) SetScrollToFocusViewport(scrollToFocus bool) {
	var _arg0 *C.GtkViewport // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))
	if scrollToFocus {
		_arg1 = C.TRUE
	}

	C.gtk_viewport_set_scroll_to_focus(_arg0, _arg1)
}

func (s viewport) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s viewport) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s viewport) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s viewport) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s viewport) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s viewport) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s viewport) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b viewport) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}

func (s viewport) Border() (Border, bool) {
	return WrapScrollable(gextras.InternObject(s)).Border()
}

func (s viewport) HAdjustment() Adjustment {
	return WrapScrollable(gextras.InternObject(s)).HAdjustment()
}

func (s viewport) HScrollPolicy() ScrollablePolicy {
	return WrapScrollable(gextras.InternObject(s)).HScrollPolicy()
}

func (s viewport) VAdjustment() Adjustment {
	return WrapScrollable(gextras.InternObject(s)).VAdjustment()
}

func (s viewport) VScrollPolicy() ScrollablePolicy {
	return WrapScrollable(gextras.InternObject(s)).VScrollPolicy()
}

func (s viewport) SetHAdjustment(hadjustment Adjustment) {
	WrapScrollable(gextras.InternObject(s)).SetHAdjustment(hadjustment)
}

func (s viewport) SetHScrollPolicy(policy ScrollablePolicy) {
	WrapScrollable(gextras.InternObject(s)).SetHScrollPolicy(policy)
}

func (s viewport) SetVAdjustment(vadjustment Adjustment) {
	WrapScrollable(gextras.InternObject(s)).SetVAdjustment(vadjustment)
}

func (s viewport) SetVScrollPolicy(policy ScrollablePolicy) {
	WrapScrollable(gextras.InternObject(s)).SetVScrollPolicy(policy)
}
