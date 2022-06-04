// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkcenterbox.go.
var GTypeCenterBox = coreglib.Type(C.gtk_center_box_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeCenterBox, F: marshalCenterBox},
	})
}

// CenterBoxOverrider contains methods that are overridable.
type CenterBoxOverrider interface {
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

	*coreglib.Object
	Orientable
}

var (
	_ Widgetter         = (*CenterBox)(nil)
	_ coreglib.Objector = (*CenterBox)(nil)
)

func classInitCenterBoxer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapCenterBox(obj *coreglib.Object) *CenterBox {
	return &CenterBox{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
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

func marshalCenterBox(p uintptr) (interface{}, error) {
	return wrapCenterBox(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCenterBox creates a new GtkCenterBox.
//
// The function returns the following values:
//
//    - centerBox: new GtkCenterBox.
//
func NewCenterBox() *CenterBox {
	_gret := girepository.MustFind("Gtk", "CenterBox").InvokeMethod("new_CenterBox", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _centerBox *CenterBox // out

	_centerBox = wrapCenterBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _centerBox
}

// CenterWidget gets the center widget, or NULL if there is none.
//
// The function returns the following values:
//
//    - widget (optional): center widget.
//
func (self *CenterBox) CenterWidget() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "CenterBox").InvokeMethod("get_center_widget", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "CenterBox").InvokeMethod("get_end_widget", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "CenterBox").InvokeMethod("get_start_widget", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
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

// SetCenterWidget sets the center widget.
//
// To remove the existing center widget, pas NULL.
//
// The function takes the following parameters:
//
//    - child (optional): new center widget, or NULL.
//
func (self *CenterBox) SetCenterWidget(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	girepository.MustFind("Gtk", "CenterBox").InvokeMethod("set_center_widget", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	girepository.MustFind("Gtk", "CenterBox").InvokeMethod("set_end_widget", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	girepository.MustFind("Gtk", "CenterBox").InvokeMethod("set_start_widget", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}
