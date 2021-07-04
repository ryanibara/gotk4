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
		{T: externglib.Type(C.gtk_scrollbar_get_type()), F: marshalScrollbar},
	})
}

// Scrollbar: the `GtkScrollbar` widget is a horizontal or vertical scrollbar.
//
// !An example GtkScrollbar (scrollbar.png)
//
// Its position and movement are controlled by the adjustment that is passed to
// or created by [ctor@Gtk.Scrollbar.new]. See [class.Gtk.Adjustment] for more
// details. The [property@Gtk.Adjustment:value] field sets the position of the
// thumb and must be between [property@Gtk.Adjustment:lower] and
// [property@Gtk.Adjustment:upper] - [property@Gtk.Adjustment:page-size]. The
// [property@Gtk.Adjustment:page-size] represents the size of the visible
// scrollable area.
//
// The fields [property@Gtk.Adjustment:step-increment] and
// [property@Gtk.Adjustment:page-increment] fields are added to or subtracted
// from the [property@Gtk.Adjustment:value] when the user asks to move by a step
// (using e.g. the cursor arrow keys) or by a page (using e.g. the Page Down/Up
// keys).
//
//
// CSS nodes
//
// “` scrollbar ╰── range[.fine-tune] ╰── trough ╰── slider “`
//
// `GtkScrollbar` has a main CSS node with name scrollbar and a subnode for its
// contents. The main node gets the .horizontal or .vertical style classes
// applied, depending on the scrollbar's orientation.
//
// The range node gets the style class .fine-tune added when the scrollbar is in
// 'fine-tuning' mode.
//
// Other style classes that may be added to scrollbars inside
// [class@Gtk.ScrolledWindow] include the positional classes (.left, .right,
// .top, .bottom) and style classes related to overlay scrolling
// (.overlay-indicator, .dragging, .hovering).
//
//
// Accessibility
//
// `GtkScrollbar` uses the GTK_ACCESSIBLE_ROLE_SCROLLBAR role.
type Scrollbar interface {
	Widget
	Orientable

	Adjustment() Adjustment

	SetAdjustmentScrollbar(adjustment Adjustment)
}

// scrollbar implements the Scrollbar class.
type scrollbar struct {
	Widget
}

// WrapScrollbar wraps a GObject to the right type. It is
// primarily used internally.
func WrapScrollbar(obj *externglib.Object) Scrollbar {
	return scrollbar{
		Widget: WrapWidget(obj),
	}
}

func marshalScrollbar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapScrollbar(obj), nil
}

func NewScrollbar(orientation Orientation, adjustment Adjustment) Scrollbar {
	var _arg1 C.GtkOrientation // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)
	_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_scrollbar_new(_arg1, _arg2)

	var _scrollbar Scrollbar // out

	_scrollbar = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Scrollbar)

	return _scrollbar
}

func (s scrollbar) Adjustment() Adjustment {
	var _arg0 *C.GtkScrollbar  // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkScrollbar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrollbar_get_adjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Adjustment)

	return _adjustment
}

func (s scrollbar) SetAdjustmentScrollbar(adjustment Adjustment) {
	var _arg0 *C.GtkScrollbar  // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkScrollbar)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_scrollbar_set_adjustment(_arg0, _arg1)
}

func (s scrollbar) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s scrollbar) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s scrollbar) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s scrollbar) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s scrollbar) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s scrollbar) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s scrollbar) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b scrollbar) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}

func (o scrollbar) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o scrollbar) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
