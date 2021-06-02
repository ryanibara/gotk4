// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
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

	// BaselinePosition returns the baseline position of the layout.
	BaselinePosition() BaselinePosition
	// CenterWidget returns the center widget of the layout.
	CenterWidget() Widget
	// EndWidget returns the end widget of the layout.
	EndWidget() Widget
	// Orientation gets the current orienration of the layout manager.
	Orientation() Orientation
	// StartWidget returns the start widget fo the layout.
	StartWidget() Widget
	// SetBaselinePosition sets the new baseline position of @self
	SetBaselinePosition(baselinePosition BaselinePosition)
	// SetCenterWidget sets the new center widget of @self.
	//
	// To remove the existing center widget, pass nil.
	SetCenterWidget(widget Widget)
	// SetEndWidget sets the new end widget of @self.
	//
	// To remove the existing center widget, pass nil.
	SetEndWidget(widget Widget)
	// SetOrientation sets the orientation of @self.
	SetOrientation(orientation Orientation)
	// SetStartWidget sets the new start widget of @self.
	//
	// To remove the existing start widget, pass nil.
	SetStartWidget(widget Widget)
}

// centerLayout implements the CenterLayout interface.
type centerLayout struct {
	LayoutManager
}

var _ CenterLayout = (*centerLayout)(nil)

// WrapCenterLayout wraps a GObject to the right type. It is
// primarily used internally.
func WrapCenterLayout(obj *externglib.Object) CenterLayout {
	return CenterLayout{
		LayoutManager: WrapLayoutManager(obj),
	}
}

func marshalCenterLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCenterLayout(obj), nil
}

// NewCenterLayout constructs a class CenterLayout.
func NewCenterLayout() CenterLayout {
	ret := C.gtk_center_layout_new()

	var ret0 CenterLayout

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(CenterLayout)

	return ret0
}

// BaselinePosition returns the baseline position of the layout.
func (s centerLayout) BaselinePosition() BaselinePosition {
	var arg0 *C.GtkCenterLayout

	arg0 = (*C.GtkCenterLayout)(s.Native())

	ret := C.gtk_center_layout_get_baseline_position(arg0)

	var ret0 BaselinePosition

	ret0 = BaselinePosition(ret)

	return ret0
}

// CenterWidget returns the center widget of the layout.
func (s centerLayout) CenterWidget() Widget {
	var arg0 *C.GtkCenterLayout

	arg0 = (*C.GtkCenterLayout)(s.Native())

	ret := C.gtk_center_layout_get_center_widget(arg0)

	var ret0 Widget

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Widget)

	return ret0
}

// EndWidget returns the end widget of the layout.
func (s centerLayout) EndWidget() Widget {
	var arg0 *C.GtkCenterLayout

	arg0 = (*C.GtkCenterLayout)(s.Native())

	ret := C.gtk_center_layout_get_end_widget(arg0)

	var ret0 Widget

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Widget)

	return ret0
}

// Orientation gets the current orienration of the layout manager.
func (s centerLayout) Orientation() Orientation {
	var arg0 *C.GtkCenterLayout

	arg0 = (*C.GtkCenterLayout)(s.Native())

	ret := C.gtk_center_layout_get_orientation(arg0)

	var ret0 Orientation

	ret0 = Orientation(ret)

	return ret0
}

// StartWidget returns the start widget fo the layout.
func (s centerLayout) StartWidget() Widget {
	var arg0 *C.GtkCenterLayout

	arg0 = (*C.GtkCenterLayout)(s.Native())

	ret := C.gtk_center_layout_get_start_widget(arg0)

	var ret0 Widget

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Widget)

	return ret0
}

// SetBaselinePosition sets the new baseline position of @self
func (s centerLayout) SetBaselinePosition(baselinePosition BaselinePosition) {
	var arg0 *C.GtkCenterLayout
	var arg1 C.GtkBaselinePosition

	arg0 = (*C.GtkCenterLayout)(s.Native())
	arg1 = (C.GtkBaselinePosition)(baselinePosition)

	C.gtk_center_layout_set_baseline_position(arg0, arg1)
}

// SetCenterWidget sets the new center widget of @self.
//
// To remove the existing center widget, pass nil.
func (s centerLayout) SetCenterWidget(widget Widget) {
	var arg0 *C.GtkCenterLayout
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkCenterLayout)(s.Native())
	arg1 = (*C.GtkWidget)(widget.Native())

	C.gtk_center_layout_set_center_widget(arg0, arg1)
}

// SetEndWidget sets the new end widget of @self.
//
// To remove the existing center widget, pass nil.
func (s centerLayout) SetEndWidget(widget Widget) {
	var arg0 *C.GtkCenterLayout
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkCenterLayout)(s.Native())
	arg1 = (*C.GtkWidget)(widget.Native())

	C.gtk_center_layout_set_end_widget(arg0, arg1)
}

// SetOrientation sets the orientation of @self.
func (s centerLayout) SetOrientation(orientation Orientation) {
	var arg0 *C.GtkCenterLayout
	var arg1 C.GtkOrientation

	arg0 = (*C.GtkCenterLayout)(s.Native())
	arg1 = (C.GtkOrientation)(orientation)

	C.gtk_center_layout_set_orientation(arg0, arg1)
}

// SetStartWidget sets the new start widget of @self.
//
// To remove the existing start widget, pass nil.
func (s centerLayout) SetStartWidget(widget Widget) {
	var arg0 *C.GtkCenterLayout
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkCenterLayout)(s.Native())
	arg1 = (*C.GtkWidget)(widget.Native())

	C.gtk_center_layout_set_start_widget(arg0, arg1)
}
