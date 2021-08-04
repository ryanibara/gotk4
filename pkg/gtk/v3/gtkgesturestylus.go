// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_stylus_get_type()), F: marshalGestureStylusser},
	})
}

// GestureStylus is a Gesture implementation specific to stylus input. The
// provided signals just provide the basic information
type GestureStylus struct {
	GestureSingle
}

func wrapGestureStylus(obj *externglib.Object) *GestureStylus {
	return &GestureStylus{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureStylusser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGestureStylus(obj), nil
}

// NewGestureStylus creates a new GestureStylus.
func NewGestureStylus(widget Widgetter) *GestureStylus {
	var _arg1 *C.GtkWidget  // out
	var _cret *C.GtkGesture // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_gesture_stylus_new(_arg1)

	var _gestureStylus *GestureStylus // out

	_gestureStylus = wrapGestureStylus(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureStylus
}

// Axis returns the current value for the requested axis. This function must be
// called from either the GestureStylus:down, GestureStylus:motion,
// GestureStylus:up or GestureStylus:proximity signals.
func (gesture *GestureStylus) Axis(axis gdk.AxisUse) (float64, bool) {
	var _arg0 *C.GtkGestureStylus // out
	var _arg1 C.GdkAxisUse        // out
	var _arg2 C.gdouble           // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(gesture.Native()))
	_arg1 = C.GdkAxisUse(axis)

	_cret = C.gtk_gesture_stylus_get_axis(_arg0, _arg1, &_arg2)

	var _value float64 // out
	var _ok bool       // out

	_value = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// DeviceTool returns the DeviceTool currently driving input through this
// gesture. This function must be called from either the GestureStylus::down,
// GestureStylus::motion, GestureStylus::up or GestureStylus::proximity signal
// handlers.
func (gesture *GestureStylus) DeviceTool() *gdk.DeviceTool {
	var _arg0 *C.GtkGestureStylus // out
	var _cret *C.GdkDeviceTool    // in

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_stylus_get_device_tool(_arg0)

	var _deviceTool *gdk.DeviceTool // out

	if _cret != nil {
		{
			obj := externglib.Take(unsafe.Pointer(_cret))
			_deviceTool = &gdk.DeviceTool{
				Object: obj,
			}
		}
	}

	return _deviceTool
}
