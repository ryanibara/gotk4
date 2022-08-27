// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_GestureLongPress_ConnectPressed(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk4_GestureLongPress_ConnectCancelled(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeGestureLongPress = coreglib.Type(C.gtk_gesture_long_press_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGestureLongPress, F: marshalGestureLongPress},
	})
}

// GestureLongPress: GtkGestureLongPress is a GtkGesture for long presses.
//
// This gesture is also known as “Press and Hold”.
//
// When the timeout is exceeded, the gesture is triggering the
// gtk.GestureLongPress::pressed signal.
//
// If the touchpoint is lifted before the timeout passes, or if it drifts too
// far of the initial press point, the gtk.GestureLongPress::cancelled signal
// will be emitted.
//
// How long the timeout is before the ::pressed signal gets emitted is
// determined by the gtk.Settings:gtk-long-press-time setting. It can be
// modified by the gtk.GestureLongPress:delay-factor property.
type GestureLongPress struct {
	_ [0]func() // equal guard
	GestureSingle
}

var (
	_ Gesturer = (*GestureLongPress)(nil)
)

func wrapGestureLongPress(obj *coreglib.Object) *GestureLongPress {
	return &GestureLongPress{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureLongPress(p uintptr) (interface{}, error) {
	return wrapGestureLongPress(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectCancelled is emitted whenever a press moved too far, or was released
// before gtk.GestureLongPress::pressed happened.
func (gesture *GestureLongPress) ConnectCancelled(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "cancelled", false, unsafe.Pointer(C._gotk4_gtk4_GestureLongPress_ConnectCancelled), f)
}

// ConnectPressed is emitted whenever a press goes unmoved/unreleased longer
// than what the GTK defaults tell.
func (gesture *GestureLongPress) ConnectPressed(f func(x, y float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "pressed", false, unsafe.Pointer(C._gotk4_gtk4_GestureLongPress_ConnectPressed), f)
}

// NewGestureLongPress returns a newly created GtkGesture that recognizes long
// presses.
//
// The function returns the following values:
//
//    - gestureLongPress: newly created GtkGestureLongPress.
//
func NewGestureLongPress() *GestureLongPress {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_long_press_new()

	var _gestureLongPress *GestureLongPress // out

	_gestureLongPress = wrapGestureLongPress(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureLongPress
}

// DelayFactor returns the delay factor.
//
// The function returns the following values:
//
//    - gdouble: delay factor.
//
func (gesture *GestureLongPress) DelayFactor() float64 {
	var _arg0 *C.GtkGestureLongPress // out
	var _cret C.double               // in

	_arg0 = (*C.GtkGestureLongPress)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_cret = C.gtk_gesture_long_press_get_delay_factor(_arg0)
	runtime.KeepAlive(gesture)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetDelayFactor applies the given delay factor.
//
// The default long press time will be multiplied by this value. Valid values
// are in the range [0.5..2.0].
//
// The function takes the following parameters:
//
//    - delayFactor: delay factor to apply.
//
func (gesture *GestureLongPress) SetDelayFactor(delayFactor float64) {
	var _arg0 *C.GtkGestureLongPress // out
	var _arg1 C.double               // out

	_arg0 = (*C.GtkGestureLongPress)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))
	_arg1 = C.double(delayFactor)

	C.gtk_gesture_long_press_set_delay_factor(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(delayFactor)
}
