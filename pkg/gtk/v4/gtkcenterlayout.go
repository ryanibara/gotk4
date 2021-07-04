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
		{T: externglib.Type(C.gtk_center_layout_get_type()), F: marshalCenterLayout},
	})
}

// CenterLayout: `GtkCenterLayout` is a layout manager that manages up to three
// children.
//
// The start widget is allocated at the start of the layout (left in
// left-to-right locales and right in right-to-left ones), and the end widget at
// the end.
//
// The center widget is centered regarding the full width of the layout's.
type CenterLayout interface {
	LayoutManager

	BaselinePosition() BaselinePosition

	CenterWidget() Widget

	EndWidget() Widget

	Orientation() Orientation

	StartWidget() Widget

	SetBaselinePositionCenterLayout(baselinePosition BaselinePosition)

	SetCenterWidgetCenterLayout(widget Widget)

	SetEndWidgetCenterLayout(widget Widget)

	SetOrientationCenterLayout(orientation Orientation)

	SetStartWidgetCenterLayout(widget Widget)
}

// centerLayout implements the CenterLayout class.
type centerLayout struct {
	LayoutManager
}

// WrapCenterLayout wraps a GObject to the right type. It is
// primarily used internally.
func WrapCenterLayout(obj *externglib.Object) CenterLayout {
	return centerLayout{
		LayoutManager: WrapLayoutManager(obj),
	}
}

func marshalCenterLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCenterLayout(obj), nil
}

func NewCenterLayout() CenterLayout {
	var _cret *C.GtkLayoutManager // in

	_cret = C.gtk_center_layout_new()

	var _centerLayout CenterLayout // out

	_centerLayout = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(CenterLayout)

	return _centerLayout
}

func (s centerLayout) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkCenterLayout    // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_center_layout_get_baseline_position(_arg0)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

func (s centerLayout) CenterWidget() Widget {
	var _arg0 *C.GtkCenterLayout // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_center_layout_get_center_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (s centerLayout) EndWidget() Widget {
	var _arg0 *C.GtkCenterLayout // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_center_layout_get_end_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (s centerLayout) Orientation() Orientation {
	var _arg0 *C.GtkCenterLayout // out
	var _cret C.GtkOrientation   // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_center_layout_get_orientation(_arg0)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

func (s centerLayout) StartWidget() Widget {
	var _arg0 *C.GtkCenterLayout // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_center_layout_get_start_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (s centerLayout) SetBaselinePositionCenterLayout(baselinePosition BaselinePosition) {
	var _arg0 *C.GtkCenterLayout    // out
	var _arg1 C.GtkBaselinePosition // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkBaselinePosition(baselinePosition)

	C.gtk_center_layout_set_baseline_position(_arg0, _arg1)
}

func (s centerLayout) SetCenterWidgetCenterLayout(widget Widget) {
	var _arg0 *C.GtkCenterLayout // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_center_layout_set_center_widget(_arg0, _arg1)
}

func (s centerLayout) SetEndWidgetCenterLayout(widget Widget) {
	var _arg0 *C.GtkCenterLayout // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_center_layout_set_end_widget(_arg0, _arg1)
}

func (s centerLayout) SetOrientationCenterLayout(orientation Orientation) {
	var _arg0 *C.GtkCenterLayout // out
	var _arg1 C.GtkOrientation   // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkOrientation(orientation)

	C.gtk_center_layout_set_orientation(_arg0, _arg1)
}

func (s centerLayout) SetStartWidgetCenterLayout(widget Widget) {
	var _arg0 *C.GtkCenterLayout // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_center_layout_set_start_widget(_arg0, _arg1)
}
