// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_long_press_get_type()), F: marshalGestureLongPress},
	})
}

// GestureLongPress is a Gesture implementation able to recognize long presses,
// triggering the GestureLongPress::pressed after the timeout is exceeded.
//
// If the touchpoint is lifted before the timeout passes, or if it drifts too
// far of the initial press point, the GestureLongPress::cancelled signal will
// be emitted.
type GestureLongPress interface {
	GestureSingle
}

// gestureLongPress implements the GestureLongPress class.
type gestureLongPress struct {
	GestureSingle
}

var _ GestureLongPress = (*gestureLongPress)(nil)

// WrapGestureLongPress wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureLongPress(obj *externglib.Object) GestureLongPress {
	return gestureLongPress{
		GestureSingle: WrapGestureSingle(obj),
	}
}

func marshalGestureLongPress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureLongPress(obj), nil
}

// NewGestureLongPress constructs a class GestureLongPress.
func NewGestureLongPress(widget Widget) GestureLongPress {
	var _arg1 *C.GtkWidget          // out
	var _cret C.GtkGestureLongPress // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_gesture_long_press_new(_arg1)

	var _gestureLongPress GestureLongPress // out

	_gestureLongPress = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(GestureLongPress)

	return _gestureLongPress
}
