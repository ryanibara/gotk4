// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_key_get_type()), F: marshalEventControllerKey},
	})
}

// EventControllerKey: `GtkEventControllerKey` is an event controller that
// provides access to key events.
type EventControllerKey interface {
	EventController

	ForwardEventControllerKey(widget Widget) bool

	Group() uint

	ImContext() IMContext

	SetImContextEventControllerKey(imContext IMContext)
}

// eventControllerKey implements the EventControllerKey class.
type eventControllerKey struct {
	EventController
}

// WrapEventControllerKey wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventControllerKey(obj *externglib.Object) EventControllerKey {
	return eventControllerKey{
		EventController: WrapEventController(obj),
	}
}

func marshalEventControllerKey(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventControllerKey(obj), nil
}

func NewEventControllerKey() EventControllerKey {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_event_controller_key_new()

	var _eventControllerKey EventControllerKey // out

	_eventControllerKey = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(EventControllerKey)

	return _eventControllerKey
}

func (c eventControllerKey) ForwardEventControllerKey(widget Widget) bool {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkWidget             // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_event_controller_key_forward(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c eventControllerKey) Group() uint {
	var _arg0 *C.GtkEventControllerKey // out
	var _cret C.guint                  // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_event_controller_key_get_group(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (c eventControllerKey) ImContext() IMContext {
	var _arg0 *C.GtkEventControllerKey // out
	var _cret *C.GtkIMContext          // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_event_controller_key_get_im_context(_arg0)

	var _imContext IMContext // out

	_imContext = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(IMContext)

	return _imContext
}

func (c eventControllerKey) SetImContextEventControllerKey(imContext IMContext) {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkIMContext          // out

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkIMContext)(unsafe.Pointer(imContext.Native()))

	C.gtk_event_controller_key_set_im_context(_arg0, _arg1)
}
