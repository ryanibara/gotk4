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

// glib.Type values for gtkaspectframe.go.
var GTypeAspectFrame = coreglib.Type(C.gtk_aspect_frame_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeAspectFrame, F: marshalAspectFrame},
	})
}

// AspectFrame: GtkAspectFrame preserves the aspect ratio of its child.
//
// The frame can respect the aspect ratio of the child widget, or use its own
// aspect ratio.
//
//
// CSS nodes
//
// GtkAspectFrame uses a CSS node with name frame.
type AspectFrame struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*AspectFrame)(nil)
)

func wrapAspectFrame(obj *coreglib.Object) *AspectFrame {
	return &AspectFrame{
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
	}
}

func marshalAspectFrame(p uintptr) (interface{}, error) {
	return wrapAspectFrame(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Child gets the child widget of self.
//
// The function returns the following values:
//
//    - widget (optional): child widget of self@.
//
func (self *AspectFrame) Child() Widgetter {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**AspectFrame)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "AspectFrame").InvokeMethod("get_child", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if _cret != nil {
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

// ObeyChild returns whether the child's size request should override the set
// aspect ratio of the GtkAspectFrame.
//
// The function returns the following values:
//
//    - ok: whether to obey the child's size request.
//
func (self *AspectFrame) ObeyChild() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**AspectFrame)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "AspectFrame").InvokeMethod("get_obey_child", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetChild sets the child widget of self.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (self *AspectFrame) SetChild(child Widgetter) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}
	*(**AspectFrame)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "AspectFrame").InvokeMethod("set_child", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetObeyChild sets whether the aspect ratio of the child's size request should
// override the set aspect ratio of the GtkAspectFrame.
//
// The function takes the following parameters:
//
//    - obeyChild: if TRUE, ratio is ignored, and the aspect ratio is taken from
//      the requistion of the child.
//
func (self *AspectFrame) SetObeyChild(obeyChild bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if obeyChild {
		_arg1 = C.TRUE
	}
	*(**AspectFrame)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "AspectFrame").InvokeMethod("set_obey_child", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(obeyChild)
}
