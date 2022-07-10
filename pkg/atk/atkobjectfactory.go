// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_atk1_ObjectFactoryClass_invalidate(void*);
import "C"

// GTypeObjectFactory returns the GType for the type ObjectFactory.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeObjectFactory() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Atk", "ObjectFactory").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalObjectFactory)
	return gtype
}

// ObjectFactoryOverrider contains methods that are overridable.
type ObjectFactoryOverrider interface {
	// Invalidate: inform factory that it is no longer being used to create
	// accessibles. When called, factory may need to inform Objects which it has
	// created that they need to be re-instantiated. Note: primarily used for
	// runtime replacement of ObjectFactorys in object registries.
	Invalidate()
}

// ObjectFactory: this class is the base object class for a factory used to
// create an accessible object for a specific GType. The function
// atk_registry_set_factory_type() is normally called to store in the registry
// the factory type to be used to create an accessible of a particular GType.
type ObjectFactory struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ObjectFactory)(nil)
)

func classInitObjectFactorier(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Atk", "ObjectFactoryClass")

	if _, ok := goval.(interface{ Invalidate() }); ok {
		o := pclass.StructFieldOffset("invalidate")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_atk1_ObjectFactoryClass_invalidate)
	}
}

//export _gotk4_atk1_ObjectFactoryClass_invalidate
func _gotk4_atk1_ObjectFactoryClass_invalidate(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Invalidate() })

	iface.Invalidate()
}

func wrapObjectFactory(obj *coreglib.Object) *ObjectFactory {
	return &ObjectFactory{
		Object: obj,
	}
}

func marshalObjectFactory(p uintptr) (interface{}, error) {
	return wrapObjectFactory(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CreateAccessible provides an Object that implements an accessibility
// interface on behalf of obj.
//
// The function takes the following parameters:
//
//    - obj: #GObject.
//
// The function returns the following values:
//
//    - object that implements an accessibility interface on behalf of obj.
//
func (factory *ObjectFactory) CreateAccessible(obj *coreglib.Object) *ObjectClass {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(obj.Native()))

	_info := girepository.MustFind("Atk", "ObjectFactory")
	_gret := _info.InvokeClassMethod("create_accessible", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(factory)
	runtime.KeepAlive(obj)

	var _object *ObjectClass // out

	_object = wrapObject(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _object
}

// Invalidate: inform factory that it is no longer being used to create
// accessibles. When called, factory may need to inform Objects which it has
// created that they need to be re-instantiated. Note: primarily used for
// runtime replacement of ObjectFactorys in object registries.
func (factory *ObjectFactory) Invalidate() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(factory).Native()))

	_info := girepository.MustFind("Atk", "ObjectFactory")
	_info.InvokeClassMethod("invalidate", _args[:], nil)

	runtime.KeepAlive(factory)
}
