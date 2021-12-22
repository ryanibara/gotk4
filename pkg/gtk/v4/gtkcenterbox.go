// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_center_box_get_type()), F: marshalCenterBoxer},
	})
}

// CenterBox: GtkCenterBox arranges three children in a row, keeping the middle
// child centered as well as possible.
//
// !An example GtkCenterBox (centerbox.png)
//
// To add children to GtkCenterBox, use gtk.CenterBox.SetStartWidget(),
// gtk.CenterBox.SetCenterWidget() and gtk.CenterBox.SetEndWidget().
//
// The sizing and positioning of children can be influenced with the align and
// expand properties of the children.
//
//
// GtkCenterBox as GtkBuildable
//
// The GtkCenterBox implementation of the GtkBuildable interface supports
// placing children in the 3 positions by specifying “start”, “center” or “end”
// as the “type” attribute of a <child> element.
//
//
// CSS nodes
//
// GtkCenterBox uses a single CSS node with the name “box”,
//
// The first child of the GtkCenterBox will be allocated depending on the text
// direction, i.e. in left-to-right layouts it will be allocated on the left and
// in right-to-left layouts on the right.
//
// In vertical orientation, the nodes of the children are arranged from top to
// bottom.
//
//
// Accessibility
//
// GtkCenterBox uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type CenterBox struct {
	_ [0]func() // equal guard
	Widget

	*externglib.Object
	Orientable
}

var (
	_ Widgetter           = (*CenterBox)(nil)
	_ externglib.Objector = (*CenterBox)(nil)
)

func wrapCenterBox(obj *externglib.Object) *CenterBox {
	return &CenterBox{
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
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalCenterBoxer(p uintptr) (interface{}, error) {
	return wrapCenterBox(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCenterBox creates a new GtkCenterBox.
//
// The function returns the following values:
//
//    - centerBox: new GtkCenterBox.
//
func NewCenterBox() *CenterBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_center_box_new()

	var _centerBox *CenterBox // out

	_centerBox = wrapCenterBox(externglib.Take(unsafe.Pointer(_cret)))

	return _centerBox
}

// BaselinePosition gets the value set by
// gtk_center_box_set_baseline_position().
//
// The function returns the following values:
//
//    - baselinePosition: baseline position.
//
func (self *CenterBox) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkCenterBox       // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_box_get_baseline_position(_arg0)
	runtime.KeepAlive(self)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

// CenterWidget gets the center widget, or NULL if there is none.
//
// The function returns the following values:
//
//    - widget (optional): center widget.
//
func (self *CenterBox) CenterWidget() Widgetter {
	var _arg0 *C.GtkCenterBox // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_box_get_center_widget(_arg0)
	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// EndWidget gets the end widget, or NULL if there is none.
//
// The function returns the following values:
//
//    - widget (optional): end widget.
//
func (self *CenterBox) EndWidget() Widgetter {
	var _arg0 *C.GtkCenterBox // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_box_get_end_widget(_arg0)
	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// StartWidget gets the start widget, or NULL if there is none.
//
// The function returns the following values:
//
//    - widget (optional): start widget.
//
func (self *CenterBox) StartWidget() Widgetter {
	var _arg0 *C.GtkCenterBox // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_box_get_start_widget(_arg0)
	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// SetBaselinePosition sets the baseline position of a center box.
//
// This affects only horizontal boxes with at least one baseline aligned child.
// If there is more vertical space available than requested, and the baseline is
// not allocated by the parent then position is used to allocate the baseline
// wrt. the extra space available.
//
// The function takes the following parameters:
//
//    - position: GtkBaselinePosition.
//
func (self *CenterBox) SetBaselinePosition(position BaselinePosition) {
	var _arg0 *C.GtkCenterBox       // out
	var _arg1 C.GtkBaselinePosition // out

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkBaselinePosition(position)

	C.gtk_center_box_set_baseline_position(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)
}

// SetCenterWidget sets the center widget.
//
// To remove the existing center widget, pas NULL.
//
// The function takes the following parameters:
//
//    - child (optional): new center widget, or NULL.
//
func (self *CenterBox) SetCenterWidget(child Widgetter) {
	var _arg0 *C.GtkCenterBox // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	}

	C.gtk_center_box_set_center_widget(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetEndWidget sets the end widget.
//
// To remove the existing end widget, pass NULL.
//
// The function takes the following parameters:
//
//    - child (optional): new end widget, or NULL.
//
func (self *CenterBox) SetEndWidget(child Widgetter) {
	var _arg0 *C.GtkCenterBox // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	}

	C.gtk_center_box_set_end_widget(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetStartWidget sets the start widget.
//
// To remove the existing start widget, pass NULL.
//
// The function takes the following parameters:
//
//    - child (optional): new start widget, or NULL.
//
func (self *CenterBox) SetStartWidget(child Widgetter) {
	var _arg0 *C.GtkCenterBox // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	}

	C.gtk_center_box_set_start_widget(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}
