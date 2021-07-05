// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_plug_get_type()), F: marshalPlug},
	})
}

// Plug: see Socket
type Plug interface {
	Object

	// AsComponent casts the class to the Component interface.
	AsComponent() Component

	// ID gets the unique ID of an Plug object, which can be used to embed
	// inside of an Socket using atk_socket_embed().
	//
	// Internally, this calls a class function that should be registered by the
	// IPC layer (usually at-spi2-atk). The implementor of an Plug object should
	// call this function (after atk-bridge is loaded) and pass the value to the
	// process implementing the Socket, so it could embed the plug.
	ID() string
	// SetChildPlug sets @child as accessible child of @plug and @plug as
	// accessible parent of @child. @child can be NULL.
	//
	// In some cases, one can not use the AtkPlug type directly as accessible
	// object for the toplevel widget of the application. For instance in the
	// gtk case, GtkPlugAccessible can not inherit both from GtkWindowAccessible
	// and from AtkPlug. In such a case, one can create, in addition to the
	// standard accessible object for the toplevel widget, an AtkPlug object,
	// and make the former the child of the latter by calling
	// atk_plug_set_child().
	SetChildPlug(child Object)
}

// plug implements the Plug class.
type plug struct {
	Object
}

// WrapPlug wraps a GObject to the right type. It is
// primarily used internally.
func WrapPlug(obj *externglib.Object) Plug {
	return plug{
		Object: WrapObject(obj),
	}
}

func marshalPlug(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPlug(obj), nil
}

// NewPlug creates a new Plug instance.
func NewPlug() Plug {
	var _cret *C.AtkObject // in

	_cret = C.atk_plug_new()

	var _plug Plug // out

	_plug = WrapPlug(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _plug
}

func (p plug) ID() string {
	var _arg0 *C.AtkPlug // out
	var _cret *C.gchar   // in

	_arg0 = (*C.AtkPlug)(unsafe.Pointer(p.Native()))

	_cret = C.atk_plug_get_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (p plug) SetChildPlug(child Object) {
	var _arg0 *C.AtkPlug   // out
	var _arg1 *C.AtkObject // out

	_arg0 = (*C.AtkPlug)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.AtkObject)(unsafe.Pointer(child.Native()))

	C.atk_plug_set_child(_arg0, _arg1)
}

func (p plug) AsComponent() Component {
	return WrapComponent(gextras.InternObject(p))
}