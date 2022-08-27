// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_GestureStylus_ConnectUp(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk3_GestureStylus_ConnectProximity(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk3_GestureStylus_ConnectMotion(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk3_GestureStylus_ConnectDown(gpointer, gdouble, gdouble, guintptr);
import "C"

// GType values.
var (
	GTypeGestureStylus = coreglib.Type(C.gtk_gesture_stylus_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGestureStylus, F: marshalGestureStylus},
	})
}

// GestureStylus is a Gesture implementation specific to stylus input. The
// provided signals just provide the basic information.
type GestureStylus struct {
	_ [0]func() // equal guard
	GestureSingle
}

var (
	_ Gesturer = (*GestureStylus)(nil)
)

func wrapGestureStylus(obj *coreglib.Object) *GestureStylus {
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

func marshalGestureStylus(p uintptr) (interface{}, error) {
	return wrapGestureStylus(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (gesture *GestureStylus) ConnectDown(f func(object, p0 float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "down", false, unsafe.Pointer(C._gotk4_gtk3_GestureStylus_ConnectDown), f)
}

func (gesture *GestureStylus) ConnectMotion(f func(object, p0 float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "motion", false, unsafe.Pointer(C._gotk4_gtk3_GestureStylus_ConnectMotion), f)
}

func (gesture *GestureStylus) ConnectProximity(f func(object, p0 float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "proximity", false, unsafe.Pointer(C._gotk4_gtk3_GestureStylus_ConnectProximity), f)
}

func (gesture *GestureStylus) ConnectUp(f func(object, p0 float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "up", false, unsafe.Pointer(C._gotk4_gtk3_GestureStylus_ConnectUp), f)
}
