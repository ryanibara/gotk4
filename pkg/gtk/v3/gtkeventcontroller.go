// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_get_type()), F: marshalEventController},
	})
}

// EventController is a base, low-level implementation for event controllers.
// Those react to a series of Events, and possibly trigger actions as a
// consequence of those.
type EventController interface {
	gextras.Objector

	PropagationPhase() PropagationPhase

	Widget() Widget

	ResetEventController()

	SetPropagationPhaseEventController(phase PropagationPhase)
}

// eventController implements the EventController class.
type eventController struct {
	gextras.Objector
}

// WrapEventController wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventController(obj *externglib.Object) EventController {
	return eventController{
		Objector: obj,
	}
}

func marshalEventController(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventController(obj), nil
}

func (c eventController) PropagationPhase() PropagationPhase {
	var _arg0 *C.GtkEventController // out
	var _cret C.GtkPropagationPhase // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_event_controller_get_propagation_phase(_arg0)

	var _propagationPhase PropagationPhase // out

	_propagationPhase = PropagationPhase(_cret)

	return _propagationPhase
}

func (c eventController) Widget() Widget {
	var _arg0 *C.GtkEventController // out
	var _cret *C.GtkWidget          // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_event_controller_get_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (c eventController) ResetEventController() {
	var _arg0 *C.GtkEventController // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	C.gtk_event_controller_reset(_arg0)
}

func (c eventController) SetPropagationPhaseEventController(phase PropagationPhase) {
	var _arg0 *C.GtkEventController // out
	var _arg1 C.GtkPropagationPhase // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))
	_arg1 = C.GtkPropagationPhase(phase)

	C.gtk_event_controller_set_propagation_phase(_arg0, _arg1)
}
