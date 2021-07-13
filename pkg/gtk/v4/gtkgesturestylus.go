// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_stylus_get_type()), F: marshalGestureStyluser},
	})
}

// GestureStyluser describes GestureStylus's methods.
type GestureStyluser interface {
	// Axis returns the current value for the requested axis.
	Axis(axis gdk.AxisUse) (float64, bool)
	// Backlog returns the accumulated backlog of tracking information.
	Backlog() ([]gdk.TimeCoord, bool)
	// DeviceTool returns the GdkDeviceTool currently driving input through this
	// gesture.
	DeviceTool() *gdk.DeviceTool
}

// GestureStylus: GtkGestureStylus is a GtkGesture specific to stylus input.
//
// The provided signals just relay the basic information of the stylus events.
type GestureStylus struct {
	GestureSingle
}

var (
	_ GestureStyluser = (*GestureStylus)(nil)
	_ gextras.Nativer = (*GestureStylus)(nil)
)

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

func marshalGestureStyluser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGestureStylus(obj), nil
}

// NewGestureStylus creates a new GtkGestureStylus.
func NewGestureStylus() *GestureStylus {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_stylus_new()

	var _gestureStylus *GestureStylus // out

	_gestureStylus = wrapGestureStylus(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureStylus
}

// Axis returns the current value for the requested axis.
//
// This function must be called from the handler of one of the
// gtk.GestureStylus::down, gtk.GestureStylus::motion, gtk.GestureStylus::up or
// gtk.GestureStylus::proximity signals.
func (gesture *GestureStylus) Axis(axis gdk.AxisUse) (float64, bool) {
	var _arg0 *C.GtkGestureStylus // out
	var _arg1 C.GdkAxisUse        // out
	var _arg2 C.double            // in
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

// Backlog returns the accumulated backlog of tracking information.
//
// By default, GTK will limit rate of input events. On stylus input where
// accuracy of strokes is paramount, this function returns the accumulated
// coordinate/timing state before the emission of the current
// [Gtk.GestureStylus::motion] signal.
//
// This function may only be called within a gtk.GestureStylus::motion signal
// handler, the state given in this signal and obtainable through
// gtk.GestureStylus.GetAxis() express the latest (most up-to-date) state in
// motion history.
//
// The backlog is provided in chronological order.
func (gesture *GestureStylus) Backlog() ([]gdk.TimeCoord, bool) {
	var _arg0 *C.GtkGestureStylus // out
	var _arg1 *C.GdkTimeCoord
	var _arg2 C.guint    // in
	var _cret C.gboolean // in

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_stylus_get_backlog(_arg0, &_arg1, &_arg2)

	var _backlog []gdk.TimeCoord
	var _ok bool // out

	defer C.free(unsafe.Pointer(_arg1))
	_backlog = make([]gdk.TimeCoord, _arg2)
	copy(_backlog, unsafe.Slice((*gdk.TimeCoord)(unsafe.Pointer(_arg1)), _arg2))
	if _cret != 0 {
		_ok = true
	}

	return _backlog, _ok
}

// DeviceTool returns the GdkDeviceTool currently driving input through this
// gesture.
//
// This function must be called from the handler of one of the
// gtk.GestureStylus::down, gtk.GestureStylus::motion, gtk.GestureStylus::up or
// gtk.GestureStylus::proximity signals.
func (gesture *GestureStylus) DeviceTool() *gdk.DeviceTool {
	var _arg0 *C.GtkGestureStylus // out
	var _cret *C.GdkDeviceTool    // in

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_stylus_get_device_tool(_arg0)

	var _deviceTool *gdk.DeviceTool // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_deviceTool = &gdk.DeviceTool{
			Object: obj,
		}
	}

	return _deviceTool
}
