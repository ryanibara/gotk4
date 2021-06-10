// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_aspect_frame_get_type()), F: marshalAspectFrame},
	})
}

// AspectFrame: the AspectFrame is useful when you want pack a widget so that it
// can resize but always retains the same aspect ratio. For instance, one might
// be drawing a small preview of a larger image. AspectFrame derives from Frame,
// so it can draw a label and a frame around the child. The frame will be
// “shrink-wrapped” to the size of the child.
//
//
// CSS nodes
//
// GtkAspectFrame uses a CSS node with name frame.
type AspectFrame interface {
	Frame
	Buildable

	// Set: set parameters for an existing AspectFrame.
	Set(xalign float32, yalign float32, ratio float32, obeyChild bool)
}

// aspectFrame implements the AspectFrame interface.
type aspectFrame struct {
	Frame
	Buildable
}

var _ AspectFrame = (*aspectFrame)(nil)

// WrapAspectFrame wraps a GObject to the right type. It is
// primarily used internally.
func WrapAspectFrame(obj *externglib.Object) AspectFrame {
	return AspectFrame{
		Frame:     WrapFrame(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalAspectFrame(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAspectFrame(obj), nil
}

// Set: set parameters for an existing AspectFrame.
func (a aspectFrame) Set(xalign float32, yalign float32, ratio float32, obeyChild bool) {
	var _arg0 *C.GtkAspectFrame
	var _arg1 C.gfloat
	var _arg2 C.gfloat
	var _arg3 C.gfloat
	var _arg4 C.gboolean

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(a.Native()))
	_arg1 = C.gfloat(xalign)
	_arg2 = C.gfloat(yalign)
	_arg3 = C.gfloat(ratio)
	if obeyChild {
		_arg4 = C.gboolean(1)
	}

	C.gtk_aspect_frame_set(_arg0, _arg1, _arg2, _arg3, _arg4)
}
