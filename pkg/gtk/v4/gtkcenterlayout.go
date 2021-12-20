// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_center_layout_get_type()), F: marshalCenterLayouter},
	})
}

// CenterLayout: GtkCenterLayout is a layout manager that manages up to three
// children.
//
// The start widget is allocated at the start of the layout (left in
// left-to-right locales and right in right-to-left ones), and the end widget at
// the end.
//
// The center widget is centered regarding the full width of the layout's.
type CenterLayout struct {
	LayoutManager
}

var (
	_ LayoutManagerer = (*CenterLayout)(nil)
)

func wrapCenterLayout(obj *externglib.Object) *CenterLayout {
	return &CenterLayout{
		LayoutManager: LayoutManager{
			Object: obj,
		},
	}
}

func marshalCenterLayouter(p uintptr) (interface{}, error) {
	return wrapCenterLayout(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCenterLayout creates a new GtkCenterLayout.
//
// The function returns the following values:
//
//    - centerLayout: newly created GtkCenterLayout.
//
func NewCenterLayout() *CenterLayout {
	var _cret *C.GtkLayoutManager // in

	_cret = C.gtk_center_layout_new()

	var _centerLayout *CenterLayout // out

	_centerLayout = wrapCenterLayout(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _centerLayout
}

// BaselinePosition returns the baseline position of the layout.
//
// The function returns the following values:
//
//    - baselinePosition: current baseline position of self.
//
func (self *CenterLayout) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkCenterLayout    // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_layout_get_baseline_position(_arg0)
	runtime.KeepAlive(self)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

// CenterWidget returns the center widget of the layout.
//
// The function returns the following values:
//
//    - widget (optional): current center widget of self.
//
func (self *CenterLayout) CenterWidget() Widgetter {
	var _arg0 *C.GtkCenterLayout // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_layout_get_center_widget(_arg0)
	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// EndWidget returns the end widget of the layout.
//
// The function returns the following values:
//
//    - widget (optional): current end widget of self.
//
func (self *CenterLayout) EndWidget() Widgetter {
	var _arg0 *C.GtkCenterLayout // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_layout_get_end_widget(_arg0)
	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Orientation gets the current orienration of the layout manager.
//
// The function returns the following values:
//
//    - orientation: current orientation of self.
//
func (self *CenterLayout) Orientation() Orientation {
	var _arg0 *C.GtkCenterLayout // out
	var _cret C.GtkOrientation   // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_layout_get_orientation(_arg0)
	runtime.KeepAlive(self)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

// StartWidget returns the start widget fo the layout.
//
// The function returns the following values:
//
//    - widget (optional): current start widget of self.
//
func (self *CenterLayout) StartWidget() Widgetter {
	var _arg0 *C.GtkCenterLayout // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_layout_get_start_widget(_arg0)
	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// SetBaselinePosition sets the new baseline position of self.
//
// The function takes the following parameters:
//
//    - baselinePosition: new baseline position.
//
func (self *CenterLayout) SetBaselinePosition(baselinePosition BaselinePosition) {
	var _arg0 *C.GtkCenterLayout    // out
	var _arg1 C.GtkBaselinePosition // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkBaselinePosition(baselinePosition)

	C.gtk_center_layout_set_baseline_position(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(baselinePosition)
}

// SetCenterWidget sets the new center widget of self.
//
// To remove the existing center widget, pass NULL.
//
// The function takes the following parameters:
//
//    - widget (optional): new center widget.
//
func (self *CenterLayout) SetCenterWidget(widget Widgetter) {
	var _arg0 *C.GtkCenterLayout // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.gtk_center_layout_set_center_widget(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// SetEndWidget sets the new end widget of self.
//
// To remove the existing center widget, pass NULL.
//
// The function takes the following parameters:
//
//    - widget (optional): new end widget.
//
func (self *CenterLayout) SetEndWidget(widget Widgetter) {
	var _arg0 *C.GtkCenterLayout // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.gtk_center_layout_set_end_widget(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// SetOrientation sets the orientation of self.
//
// The function takes the following parameters:
//
//    - orientation: new orientation.
//
func (self *CenterLayout) SetOrientation(orientation Orientation) {
	var _arg0 *C.GtkCenterLayout // out
	var _arg1 C.GtkOrientation   // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkOrientation(orientation)

	C.gtk_center_layout_set_orientation(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(orientation)
}

// SetStartWidget sets the new start widget of self.
//
// To remove the existing start widget, pass NULL.
//
// The function takes the following parameters:
//
//    - widget (optional): new start widget.
//
func (self *CenterLayout) SetStartWidget(widget Widgetter) {
	var _arg0 *C.GtkCenterLayout // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.gtk_center_layout_set_start_widget(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}
