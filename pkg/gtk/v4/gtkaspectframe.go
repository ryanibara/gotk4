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
		{T: externglib.Type(C.gtk_aspect_frame_get_type()), F: marshalAspectFrame},
	})
}

// AspectFrame: `GtkAspectFrame` preserves the aspect ratio of its child.
//
// The frame can respect the aspect ratio of the child widget, or use its own
// aspect ratio.
//
//
// CSS nodes
//
// `GtkAspectFrame` uses a CSS node with name `frame`.
type AspectFrame interface {
	Widget

	Child() Widget

	ObeyChild() bool

	Ratio() float32

	Xalign() float32

	Yalign() float32

	SetChildAspectFrame(child Widget)

	SetObeyChildAspectFrame(obeyChild bool)

	SetRatioAspectFrame(ratio float32)

	SetXalignAspectFrame(xalign float32)

	SetYalignAspectFrame(yalign float32)
}

// aspectFrame implements the AspectFrame class.
type aspectFrame struct {
	Widget
}

// WrapAspectFrame wraps a GObject to the right type. It is
// primarily used internally.
func WrapAspectFrame(obj *externglib.Object) AspectFrame {
	return aspectFrame{
		Widget: WrapWidget(obj),
	}
}

func marshalAspectFrame(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAspectFrame(obj), nil
}

func NewAspectFrame(xalign float32, yalign float32, ratio float32, obeyChild bool) AspectFrame {
	var _arg1 C.float      // out
	var _arg2 C.float      // out
	var _arg3 C.float      // out
	var _arg4 C.gboolean   // out
	var _cret *C.GtkWidget // in

	_arg1 = C.float(xalign)
	_arg2 = C.float(yalign)
	_arg3 = C.float(ratio)
	if obeyChild {
		_arg4 = C.TRUE
	}

	_cret = C.gtk_aspect_frame_new(_arg1, _arg2, _arg3, _arg4)

	var _aspectFrame AspectFrame // out

	_aspectFrame = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(AspectFrame)

	return _aspectFrame
}

func (s aspectFrame) Child() Widget {
	var _arg0 *C.GtkAspectFrame // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_aspect_frame_get_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (s aspectFrame) ObeyChild() bool {
	var _arg0 *C.GtkAspectFrame // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_aspect_frame_get_obey_child(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s aspectFrame) Ratio() float32 {
	var _arg0 *C.GtkAspectFrame // out
	var _cret C.float           // in

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_aspect_frame_get_ratio(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

func (s aspectFrame) Xalign() float32 {
	var _arg0 *C.GtkAspectFrame // out
	var _cret C.float           // in

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_aspect_frame_get_xalign(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

func (s aspectFrame) Yalign() float32 {
	var _arg0 *C.GtkAspectFrame // out
	var _cret C.float           // in

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_aspect_frame_get_yalign(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

func (s aspectFrame) SetChildAspectFrame(child Widget) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 *C.GtkWidget      // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_aspect_frame_set_child(_arg0, _arg1)
}

func (s aspectFrame) SetObeyChildAspectFrame(obeyChild bool) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))
	if obeyChild {
		_arg1 = C.TRUE
	}

	C.gtk_aspect_frame_set_obey_child(_arg0, _arg1)
}

func (s aspectFrame) SetRatioAspectFrame(ratio float32) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 C.float           // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(ratio)

	C.gtk_aspect_frame_set_ratio(_arg0, _arg1)
}

func (s aspectFrame) SetXalignAspectFrame(xalign float32) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 C.float           // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(xalign)

	C.gtk_aspect_frame_set_xalign(_arg0, _arg1)
}

func (s aspectFrame) SetYalignAspectFrame(yalign float32) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 C.float           // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(yalign)

	C.gtk_aspect_frame_set_yalign(_arg0, _arg1)
}

func (s aspectFrame) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s aspectFrame) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s aspectFrame) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s aspectFrame) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s aspectFrame) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s aspectFrame) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s aspectFrame) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b aspectFrame) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
