// Code generated by girgen. DO NOT EDIT.

package gio

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
// extern gboolean _gotk4_gio2_DBusObjectSkeletonClass_authorize_method(void*, void*, void*);
// extern gboolean _gotk4_gio2_DBusObjectSkeleton_ConnectAuthorizeMethod(gpointer, void*, void*, guintptr);
import "C"

// glib.Type values for gdbusobjectskeleton.go.
var GTypeDBusObjectSkeleton = coreglib.Type(C.g_dbus_object_skeleton_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeDBusObjectSkeleton, F: marshalDBusObjectSkeleton},
	})
}

// DBusObjectSkeletonOverrider contains methods that are overridable.
type DBusObjectSkeletonOverrider interface {
	// The function takes the following parameters:
	//
	//    - interface_
	//    - invocation
	//
	// The function returns the following values:
	//
	AuthorizeMethod(interface_ DBusInterfaceSkeletonner, invocation *DBusMethodInvocation) bool
}

// DBusObjectSkeleton instance is essentially a group of D-Bus interfaces. The
// set of exported interfaces on the object may be dynamic and change at
// runtime.
//
// This type is intended to be used with BusObjectManager.
type DBusObjectSkeleton struct {
	_ [0]func() // equal guard
	*coreglib.Object

	DBusObject
}

var (
	_ coreglib.Objector = (*DBusObjectSkeleton)(nil)
)

func classInitDBusObjectSkeletonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GDBusObjectSkeletonClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GDBusObjectSkeletonClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface {
		AuthorizeMethod(interface_ DBusInterfaceSkeletonner, invocation *DBusMethodInvocation) bool
	}); ok {
		pclass.authorize_method = (*[0]byte)(C._gotk4_gio2_DBusObjectSkeletonClass_authorize_method)
	}
}

//export _gotk4_gio2_DBusObjectSkeletonClass_authorize_method
func _gotk4_gio2_DBusObjectSkeletonClass_authorize_method(arg0 *C.void, arg1 *C.void, arg2 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		AuthorizeMethod(interface_ DBusInterfaceSkeletonner, invocation *DBusMethodInvocation) bool
	})

	var _interface_ DBusInterfaceSkeletonner // out
	var _invocation *DBusMethodInvocation    // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusInterfaceSkeletonner is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusInterfaceSkeletonner)
			return ok
		})
		rv, ok := casted.(DBusInterfaceSkeletonner)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfaceSkeletonner")
		}
		_interface_ = rv
	}
	_invocation = wrapDBusMethodInvocation(coreglib.Take(unsafe.Pointer(arg2)))

	ok := iface.AuthorizeMethod(_interface_, _invocation)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapDBusObjectSkeleton(obj *coreglib.Object) *DBusObjectSkeleton {
	return &DBusObjectSkeleton{
		Object: obj,
		DBusObject: DBusObject{
			Object: obj,
		},
	}
}

func marshalDBusObjectSkeleton(p uintptr) (interface{}, error) {
	return wrapDBusObjectSkeleton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_DBusObjectSkeleton_ConnectAuthorizeMethod
func _gotk4_gio2_DBusObjectSkeleton_ConnectAuthorizeMethod(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 C.guintptr) (cret C.gboolean) {
	var f func(iface DBusInterfaceSkeletonner, invocation *DBusMethodInvocation) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(iface DBusInterfaceSkeletonner, invocation *DBusMethodInvocation) (ok bool))
	}

	var _iface DBusInterfaceSkeletonner   // out
	var _invocation *DBusMethodInvocation // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusInterfaceSkeletonner is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusInterfaceSkeletonner)
			return ok
		})
		rv, ok := casted.(DBusInterfaceSkeletonner)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfaceSkeletonner")
		}
		_iface = rv
	}
	_invocation = wrapDBusMethodInvocation(coreglib.Take(unsafe.Pointer(arg2)))

	ok := f(_iface, _invocation)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectAuthorizeMethod is emitted when a method is invoked by a remote caller
// and used to determine if the method call is authorized.
//
// This signal is like BusInterfaceSkeleton's
// BusInterfaceSkeleton::g-authorize-method signal, except that it is for the
// enclosing object.
//
// The default class handler just returns TRUE.
func (object *DBusObjectSkeleton) ConnectAuthorizeMethod(f func(iface DBusInterfaceSkeletonner, invocation *DBusMethodInvocation) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(object, "authorize-method", false, unsafe.Pointer(C._gotk4_gio2_DBusObjectSkeleton_ConnectAuthorizeMethod), f)
}

// NewDBusObjectSkeleton creates a new BusObjectSkeleton.
//
// The function takes the following parameters:
//
//    - objectPath: object path.
//
// The function returns the following values:
//
//    - dBusObjectSkeleton Free with g_object_unref().
//
func NewDBusObjectSkeleton(objectPath string) *DBusObjectSkeleton {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg0))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusObjectSkeleton").InvokeMethod("new_DBusObjectSkeleton", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(objectPath)

	var _dBusObjectSkeleton *DBusObjectSkeleton // out

	_dBusObjectSkeleton = wrapDBusObjectSkeleton(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusObjectSkeleton
}

// AddInterface adds interface_ to object.
//
// If object already contains a BusInterfaceSkeleton with the same interface
// name, it is removed before interface_ is added.
//
// Note that object takes its own reference on interface_ and holds it until
// removed.
//
// The function takes the following parameters:
//
//    - interface_: BusInterfaceSkeleton.
//
func (object *DBusObjectSkeleton) AddInterface(interface_ DBusInterfaceSkeletonner) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(object).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(interface_).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "DBusObjectSkeleton").InvokeMethod("add_interface", _args[:], nil)

	runtime.KeepAlive(object)
	runtime.KeepAlive(interface_)
}

// Flush: this method simply calls g_dbus_interface_skeleton_flush() on all
// interfaces belonging to object. See that method for when flushing is useful.
func (object *DBusObjectSkeleton) Flush() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(object).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gio", "DBusObjectSkeleton").InvokeMethod("flush", _args[:], nil)

	runtime.KeepAlive(object)
}

// RemoveInterface removes interface_ from object.
//
// The function takes the following parameters:
//
//    - interface_: BusInterfaceSkeleton.
//
func (object *DBusObjectSkeleton) RemoveInterface(interface_ DBusInterfaceSkeletonner) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(object).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(interface_).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "DBusObjectSkeleton").InvokeMethod("remove_interface", _args[:], nil)

	runtime.KeepAlive(object)
	runtime.KeepAlive(interface_)
}

// RemoveInterfaceByName removes the BusInterface with interface_name from
// object.
//
// If no D-Bus interface of the given interface exists, this function does
// nothing.
//
// The function takes the following parameters:
//
//    - interfaceName d-Bus interface name.
//
func (object *DBusObjectSkeleton) RemoveInterfaceByName(interfaceName string) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(object).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(interfaceName)))
	defer C.free(unsafe.Pointer(_arg1))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "DBusObjectSkeleton").InvokeMethod("remove_interface_by_name", _args[:], nil)

	runtime.KeepAlive(object)
	runtime.KeepAlive(interfaceName)
}

// SetObjectPath sets the object path for object.
//
// The function takes the following parameters:
//
//    - objectPath: valid D-Bus object path.
//
func (object *DBusObjectSkeleton) SetObjectPath(objectPath string) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(object).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg1))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "DBusObjectSkeleton").InvokeMethod("set_object_path", _args[:], nil)

	runtime.KeepAlive(object)
	runtime.KeepAlive(objectPath)
}
