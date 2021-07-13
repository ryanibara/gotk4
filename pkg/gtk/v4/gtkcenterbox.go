// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_center_box_get_type()), F: marshalCenterBoxer},
	})
}

// CenterBoxer describes CenterBox's methods.
type CenterBoxer interface {
	// BaselinePosition gets the value set by
	// gtk_center_box_set_baseline_position().
	BaselinePosition() BaselinePosition
	// CenterWidget gets the center widget, or NULL if there is none.
	CenterWidget() *Widget
	// EndWidget gets the end widget, or NULL if there is none.
	EndWidget() *Widget
	// StartWidget gets the start widget, or NULL if there is none.
	StartWidget() *Widget
	// SetBaselinePosition sets the baseline position of a center box.
	SetBaselinePosition(position BaselinePosition)
	// SetCenterWidget sets the center widget.
	SetCenterWidget(child Widgeter)
	// SetEndWidget sets the end widget.
	SetEndWidget(child Widgeter)
	// SetStartWidget sets the start widget.
	SetStartWidget(child Widgeter)
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
	Widget

	Orientable
}

var (
	_ CenterBoxer     = (*CenterBox)(nil)
	_ gextras.Nativer = (*CenterBox)(nil)
)

func wrapCenterBox(obj *externglib.Object) *CenterBox {
	return &CenterBox{
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
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalCenterBoxer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCenterBox(obj), nil
}

// NewCenterBox creates a new GtkCenterBox.
func NewCenterBox() *CenterBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_center_box_new()

	var _centerBox *CenterBox // out

	_centerBox = wrapCenterBox(externglib.Take(unsafe.Pointer(_cret)))

	return _centerBox
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *CenterBox) Native() uintptr {
	return v.Widget.InitiallyUnowned.Object.Native()
}

// BaselinePosition gets the value set by
// gtk_center_box_set_baseline_position().
func (self *CenterBox) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkCenterBox       // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_box_get_baseline_position(_arg0)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

// CenterWidget gets the center widget, or NULL if there is none.
func (self *CenterBox) CenterWidget() *Widget {
	var _arg0 *C.GtkCenterBox // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_box_get_center_widget(_arg0)

	var _widget *Widget // out

	_widget = wrapWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _widget
}

// EndWidget gets the end widget, or NULL if there is none.
func (self *CenterBox) EndWidget() *Widget {
	var _arg0 *C.GtkCenterBox // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_box_get_end_widget(_arg0)

	var _widget *Widget // out

	_widget = wrapWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _widget
}

// StartWidget gets the start widget, or NULL if there is none.
func (self *CenterBox) StartWidget() *Widget {
	var _arg0 *C.GtkCenterBox // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_box_get_start_widget(_arg0)

	var _widget *Widget // out

	_widget = wrapWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _widget
}

// SetBaselinePosition sets the baseline position of a center box.
//
// This affects only horizontal boxes with at least one baseline aligned child.
// If there is more vertical space available than requested, and the baseline is
// not allocated by the parent then position is used to allocate the baseline
// wrt. the extra space available.
func (self *CenterBox) SetBaselinePosition(position BaselinePosition) {
	var _arg0 *C.GtkCenterBox       // out
	var _arg1 C.GtkBaselinePosition // out

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkBaselinePosition(position)

	C.gtk_center_box_set_baseline_position(_arg0, _arg1)
}

// SetCenterWidget sets the center widget.
//
// To remove the existing center widget, pas NULL.
func (self *CenterBox) SetCenterWidget(child Widgeter) {
	var _arg0 *C.GtkCenterBox // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_center_box_set_center_widget(_arg0, _arg1)
}

// SetEndWidget sets the end widget.
//
// To remove the existing end widget, pass NULL.
func (self *CenterBox) SetEndWidget(child Widgeter) {
	var _arg0 *C.GtkCenterBox // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_center_box_set_end_widget(_arg0, _arg1)
}

// SetStartWidget sets the start widget.
//
// To remove the existing start widget, pass NULL.
func (self *CenterBox) SetStartWidget(child Widgeter) {
	var _arg0 *C.GtkCenterBox // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_center_box_set_start_widget(_arg0, _arg1)
}
