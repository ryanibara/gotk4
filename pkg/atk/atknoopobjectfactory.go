// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

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
		{T: externglib.Type(C.atk_no_op_object_factory_get_type()), F: marshalNoOpObjectFactory},
	})
}

// NoOpObjectFactory: the AtkObjectFactory which creates an AtkNoOpObject. An
// instance of this is created by an AtkRegistry if no factory type has not been
// specified to create an accessible object of a particular type.
type NoOpObjectFactory interface {
	ObjectFactory
}

// noOpObjectFactory implements the NoOpObjectFactory class.
type noOpObjectFactory struct {
	ObjectFactory
}

// WrapNoOpObjectFactory wraps a GObject to the right type. It is
// primarily used internally.
func WrapNoOpObjectFactory(obj *externglib.Object) NoOpObjectFactory {
	return noOpObjectFactory{
		ObjectFactory: WrapObjectFactory(obj),
	}
}

func marshalNoOpObjectFactory(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNoOpObjectFactory(obj), nil
}

func NewNoOpObjectFactory() NoOpObjectFactory {
	var _cret *C.AtkObjectFactory // in

	_cret = C.atk_no_op_object_factory_new()

	var _noOpObjectFactory NoOpObjectFactory // out

	_noOpObjectFactory = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(NoOpObjectFactory)

	return _noOpObjectFactory
}
