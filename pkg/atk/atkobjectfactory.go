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
		{T: externglib.Type(C.atk_object_factory_get_type()), F: marshalObjectFactory},
	})
}

// ObjectFactory: this class is the base object class for a factory used to
// create an accessible object for a specific GType. The function
// atk_registry_set_factory_type() is normally called to store in the registry
// the factory type to be used to create an accessible of a particular GType.
type ObjectFactory interface {
	gextras.Objector

	// CreateAccessibleObjectFactory provides an Object that implements an
	// accessibility interface on behalf of @obj
	CreateAccessibleObjectFactory(obj gextras.Objector) Object
	// AccessibleType gets the GType of the accessible which is created by the
	// factory.
	AccessibleType() externglib.Type
	// InvalidateObjectFactory: inform @factory that it is no longer being used
	// to create accessibles. When called, @factory may need to inform Objects
	// which it has created that they need to be re-instantiated. Note:
	// primarily used for runtime replacement of ObjectFactorys in object
	// registries.
	InvalidateObjectFactory()
}

// objectFactory implements the ObjectFactory class.
type objectFactory struct {
	gextras.Objector
}

// WrapObjectFactory wraps a GObject to the right type. It is
// primarily used internally.
func WrapObjectFactory(obj *externglib.Object) ObjectFactory {
	return objectFactory{
		Objector: obj,
	}
}

func marshalObjectFactory(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapObjectFactory(obj), nil
}

func (f objectFactory) CreateAccessibleObjectFactory(obj gextras.Objector) Object {
	var _arg0 *C.AtkObjectFactory // out
	var _arg1 *C.GObject          // out
	var _cret *C.AtkObject        // in

	_arg0 = (*C.AtkObjectFactory)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer(obj.Native()))

	_cret = C.atk_object_factory_create_accessible(_arg0, _arg1)

	var _object Object // out

	_object = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Object)

	return _object
}

func (f objectFactory) AccessibleType() externglib.Type {
	var _arg0 *C.AtkObjectFactory // out
	var _cret C.GType             // in

	_arg0 = (*C.AtkObjectFactory)(unsafe.Pointer(f.Native()))

	_cret = C.atk_object_factory_get_accessible_type(_arg0)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

func (f objectFactory) InvalidateObjectFactory() {
	var _arg0 *C.AtkObjectFactory // out

	_arg0 = (*C.AtkObjectFactory)(unsafe.Pointer(f.Native()))

	C.atk_object_factory_invalidate(_arg0)
}