// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern GDBusInterface* _gotk4_gio2_DBusObjectManagerIface_get_interface(void*, void*, void*);
// extern GDBusObject* _gotk4_gio2_DBusObjectManagerIface_get_object(void*, void*);
// extern GList* _gotk4_gio2_DBusObjectManagerIface_get_objects(void*);
// extern gchar* _gotk4_gio2_DBusObjectManagerIface_get_object_path(void*);
// extern void _gotk4_gio2_DBusObjectManagerIface_interface_added(void*, void*, void*);
// extern void _gotk4_gio2_DBusObjectManagerIface_interface_removed(void*, void*, void*);
// extern void _gotk4_gio2_DBusObjectManagerIface_object_added(void*, void*);
// extern void _gotk4_gio2_DBusObjectManagerIface_object_removed(void*, void*);
// extern void _gotk4_gio2_DBusObjectManager_ConnectInterfaceAdded(gpointer, void*, void*, guintptr);
// extern void _gotk4_gio2_DBusObjectManager_ConnectInterfaceRemoved(gpointer, void*, void*, guintptr);
// extern void _gotk4_gio2_DBusObjectManager_ConnectObjectAdded(gpointer, void*, guintptr);
// extern void _gotk4_gio2_DBusObjectManager_ConnectObjectRemoved(gpointer, void*, guintptr);
import "C"

// GTypeDBusObjectManager returns the GType for the type DBusObjectManager.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDBusObjectManager() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "DBusObjectManager").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDBusObjectManager)
	return gtype
}

// DBusObjectManagerOverrider contains methods that are overridable.
type DBusObjectManagerOverrider interface {
	// Interface gets the interface proxy for interface_name at object_path, if
	// any.
	//
	// The function takes the following parameters:
	//
	//    - objectPath: object path to look up.
	//    - interfaceName d-Bus interface name to look up.
	//
	// The function returns the following values:
	//
	//    - dBusInterface instance or NULL. Free with g_object_unref().
	//
	Interface(objectPath, interfaceName string) *DBusInterface
	// GetObject gets the BusObjectProxy at object_path, if any.
	//
	// The function takes the following parameters:
	//
	//    - objectPath: object path to look up.
	//
	// The function returns the following values:
	//
	//    - dBusObject or NULL. Free with g_object_unref().
	//
	GetObject(objectPath string) *DBusObject
	// ObjectPath gets the object path that manager is for.
	//
	// The function returns the following values:
	//
	//    - utf8: string owned by manager. Do not free.
	//
	ObjectPath() string
	// Objects gets all BusObject objects known to manager.
	//
	// The function returns the following values:
	//
	//    - list of BusObject objects. The returned list should be freed with
	//      g_list_free() after each element has been freed with
	//      g_object_unref().
	//
	Objects() []*DBusObject
	// The function takes the following parameters:
	//
	//    - object
	//    - interface_
	//
	InterfaceAdded(object DBusObjector, interface_ DBusInterfacer)
	// The function takes the following parameters:
	//
	//    - object
	//    - interface_
	//
	InterfaceRemoved(object DBusObjector, interface_ DBusInterfacer)
	// The function takes the following parameters:
	//
	ObjectAdded(object DBusObjector)
	// The function takes the following parameters:
	//
	ObjectRemoved(object DBusObjector)
}

// DBusObjectManager type is the base type for service- and client-side
// implementations of the standardized org.freedesktop.DBus.ObjectManager
// (http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-objectmanager)
// interface.
//
// See BusObjectManagerClient for the client-side implementation and
// BusObjectManagerServer for the service-side implementation.
//
// DBusObjectManager wraps an interface. This means the user can get the
// underlying type by calling Cast().
type DBusObjectManager struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DBusObjectManager)(nil)
)

// DBusObjectManagerer describes DBusObjectManager's interface methods.
type DBusObjectManagerer interface {
	coreglib.Objector

	// Interface gets the interface proxy for interface_name at object_path, if
	// any.
	Interface(objectPath, interfaceName string) *DBusInterface
	// GetObject gets the BusObjectProxy at object_path, if any.
	GetObject(objectPath string) *DBusObject
	// ObjectPath gets the object path that manager is for.
	ObjectPath() string
	// Objects gets all BusObject objects known to manager.
	Objects() []*DBusObject

	// Interface-added is emitted when interface is added to object.
	ConnectInterfaceAdded(func(object DBusObjector, iface DBusInterfacer)) coreglib.SignalHandle
	// Interface-removed is emitted when interface has been removed from object.
	ConnectInterfaceRemoved(func(object DBusObjector, iface DBusInterfacer)) coreglib.SignalHandle
	// Object-added is emitted when object is added to manager.
	ConnectObjectAdded(func(object DBusObjector)) coreglib.SignalHandle
	// Object-removed is emitted when object is removed from manager.
	ConnectObjectRemoved(func(object DBusObjector)) coreglib.SignalHandle
}

var _ DBusObjectManagerer = (*DBusObjectManager)(nil)

func ifaceInitDBusObjectManagerer(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Gio", "DBusObjectManagerIface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_interface"))) = unsafe.Pointer(C._gotk4_gio2_DBusObjectManagerIface_get_interface)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_object"))) = unsafe.Pointer(C._gotk4_gio2_DBusObjectManagerIface_get_object)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_object_path"))) = unsafe.Pointer(C._gotk4_gio2_DBusObjectManagerIface_get_object_path)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_objects"))) = unsafe.Pointer(C._gotk4_gio2_DBusObjectManagerIface_get_objects)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("interface_added"))) = unsafe.Pointer(C._gotk4_gio2_DBusObjectManagerIface_interface_added)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("interface_removed"))) = unsafe.Pointer(C._gotk4_gio2_DBusObjectManagerIface_interface_removed)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("object_added"))) = unsafe.Pointer(C._gotk4_gio2_DBusObjectManagerIface_object_added)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("object_removed"))) = unsafe.Pointer(C._gotk4_gio2_DBusObjectManagerIface_object_removed)
}

//export _gotk4_gio2_DBusObjectManagerIface_get_interface
func _gotk4_gio2_DBusObjectManagerIface_get_interface(arg0 *C.void, arg1 *C.void, arg2 *C.void) (cret *C.GDBusInterface) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectManagerOverrider)

	var _objectPath string    // out
	var _interfaceName string // out

	_objectPath = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_interfaceName = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	dBusInterface := iface.Interface(_objectPath, _interfaceName)

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(dBusInterface).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(dBusInterface).Native()))

	return cret
}

//export _gotk4_gio2_DBusObjectManagerIface_get_object
func _gotk4_gio2_DBusObjectManagerIface_get_object(arg0 *C.void, arg1 *C.void) (cret *C.GDBusObject) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectManagerOverrider)

	var _objectPath string // out

	_objectPath = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	dBusObject := iface.GetObject(_objectPath)

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(dBusObject).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(dBusObject).Native()))

	return cret
}

//export _gotk4_gio2_DBusObjectManagerIface_get_object_path
func _gotk4_gio2_DBusObjectManagerIface_get_object_path(arg0 *C.void) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectManagerOverrider)

	utf8 := iface.ObjectPath()

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_gio2_DBusObjectManagerIface_get_objects
func _gotk4_gio2_DBusObjectManagerIface_get_objects(arg0 *C.void) (cret *C.GList) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectManagerOverrider)

	list := iface.Objects()

	for i := len(list) - 1; i >= 0; i-- {
		src := list[i]
		var dst *C.void // out
		dst = (*C.void)(unsafe.Pointer(coreglib.InternObject(src).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(src).Native()))
		cret = C.g_list_prepend(cret, C.gpointer(unsafe.Pointer(dst)))
	}

	return cret
}

//export _gotk4_gio2_DBusObjectManagerIface_interface_added
func _gotk4_gio2_DBusObjectManagerIface_interface_added(arg0 *C.void, arg1 *C.void, arg2 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectManagerOverrider)

	var _object DBusObjector       // out
	var _interface_ DBusInterfacer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusObjector is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusObjector)
			return ok
		})
		rv, ok := casted.(DBusObjector)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusObjector")
		}
		_object = rv
	}
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gio.DBusInterfacer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusInterfacer)
			return ok
		})
		rv, ok := casted.(DBusInterfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfacer")
		}
		_interface_ = rv
	}

	iface.InterfaceAdded(_object, _interface_)
}

//export _gotk4_gio2_DBusObjectManagerIface_interface_removed
func _gotk4_gio2_DBusObjectManagerIface_interface_removed(arg0 *C.void, arg1 *C.void, arg2 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectManagerOverrider)

	var _object DBusObjector       // out
	var _interface_ DBusInterfacer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusObjector is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusObjector)
			return ok
		})
		rv, ok := casted.(DBusObjector)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusObjector")
		}
		_object = rv
	}
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gio.DBusInterfacer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusInterfacer)
			return ok
		})
		rv, ok := casted.(DBusInterfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfacer")
		}
		_interface_ = rv
	}

	iface.InterfaceRemoved(_object, _interface_)
}

//export _gotk4_gio2_DBusObjectManagerIface_object_added
func _gotk4_gio2_DBusObjectManagerIface_object_added(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectManagerOverrider)

	var _object DBusObjector // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusObjector is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusObjector)
			return ok
		})
		rv, ok := casted.(DBusObjector)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusObjector")
		}
		_object = rv
	}

	iface.ObjectAdded(_object)
}

//export _gotk4_gio2_DBusObjectManagerIface_object_removed
func _gotk4_gio2_DBusObjectManagerIface_object_removed(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectManagerOverrider)

	var _object DBusObjector // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusObjector is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusObjector)
			return ok
		})
		rv, ok := casted.(DBusObjector)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusObjector")
		}
		_object = rv
	}

	iface.ObjectRemoved(_object)
}

func wrapDBusObjectManager(obj *coreglib.Object) *DBusObjectManager {
	return &DBusObjectManager{
		Object: obj,
	}
}

func marshalDBusObjectManager(p uintptr) (interface{}, error) {
	return wrapDBusObjectManager(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_DBusObjectManager_ConnectInterfaceAdded
func _gotk4_gio2_DBusObjectManager_ConnectInterfaceAdded(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 C.guintptr) {
	var f func(object DBusObjector, iface DBusInterfacer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(object DBusObjector, iface DBusInterfacer))
	}

	var _object DBusObjector  // out
	var _iface DBusInterfacer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusObjector is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusObjector)
			return ok
		})
		rv, ok := casted.(DBusObjector)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusObjector")
		}
		_object = rv
	}
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gio.DBusInterfacer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusInterfacer)
			return ok
		})
		rv, ok := casted.(DBusInterfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfacer")
		}
		_iface = rv
	}

	f(_object, _iface)
}

// ConnectInterfaceAdded is emitted when interface is added to object.
//
// This signal exists purely as a convenience to avoid having to connect signals
// to all objects managed by manager.
func (manager *DBusObjectManager) ConnectInterfaceAdded(f func(object DBusObjector, iface DBusInterfacer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(manager, "interface-added", false, unsafe.Pointer(C._gotk4_gio2_DBusObjectManager_ConnectInterfaceAdded), f)
}

//export _gotk4_gio2_DBusObjectManager_ConnectInterfaceRemoved
func _gotk4_gio2_DBusObjectManager_ConnectInterfaceRemoved(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 C.guintptr) {
	var f func(object DBusObjector, iface DBusInterfacer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(object DBusObjector, iface DBusInterfacer))
	}

	var _object DBusObjector  // out
	var _iface DBusInterfacer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusObjector is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusObjector)
			return ok
		})
		rv, ok := casted.(DBusObjector)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusObjector")
		}
		_object = rv
	}
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gio.DBusInterfacer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusInterfacer)
			return ok
		})
		rv, ok := casted.(DBusInterfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfacer")
		}
		_iface = rv
	}

	f(_object, _iface)
}

// ConnectInterfaceRemoved is emitted when interface has been removed from
// object.
//
// This signal exists purely as a convenience to avoid having to connect signals
// to all objects managed by manager.
func (manager *DBusObjectManager) ConnectInterfaceRemoved(f func(object DBusObjector, iface DBusInterfacer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(manager, "interface-removed", false, unsafe.Pointer(C._gotk4_gio2_DBusObjectManager_ConnectInterfaceRemoved), f)
}

//export _gotk4_gio2_DBusObjectManager_ConnectObjectAdded
func _gotk4_gio2_DBusObjectManager_ConnectObjectAdded(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(object DBusObjector)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(object DBusObjector))
	}

	var _object DBusObjector // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusObjector is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusObjector)
			return ok
		})
		rv, ok := casted.(DBusObjector)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusObjector")
		}
		_object = rv
	}

	f(_object)
}

// ConnectObjectAdded is emitted when object is added to manager.
func (manager *DBusObjectManager) ConnectObjectAdded(f func(object DBusObjector)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(manager, "object-added", false, unsafe.Pointer(C._gotk4_gio2_DBusObjectManager_ConnectObjectAdded), f)
}

//export _gotk4_gio2_DBusObjectManager_ConnectObjectRemoved
func _gotk4_gio2_DBusObjectManager_ConnectObjectRemoved(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(object DBusObjector)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(object DBusObjector))
	}

	var _object DBusObjector // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusObjector is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusObjector)
			return ok
		})
		rv, ok := casted.(DBusObjector)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusObjector")
		}
		_object = rv
	}

	f(_object)
}

// ConnectObjectRemoved is emitted when object is removed from manager.
func (manager *DBusObjectManager) ConnectObjectRemoved(f func(object DBusObjector)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(manager, "object-removed", false, unsafe.Pointer(C._gotk4_gio2_DBusObjectManager_ConnectObjectRemoved), f)
}

// Interface gets the interface proxy for interface_name at object_path, if any.
//
// The function takes the following parameters:
//
//    - objectPath: object path to look up.
//    - interfaceName d-Bus interface name to look up.
//
// The function returns the following values:
//
//    - dBusInterface instance or NULL. Free with g_object_unref().
//
func (manager *DBusObjectManager) Interface(objectPath, interfaceName string) *DBusInterface {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(interfaceName)))
	defer C.free(unsafe.Pointer(_args[2]))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(manager)
	runtime.KeepAlive(objectPath)
	runtime.KeepAlive(interfaceName)

	var _dBusInterface *DBusInterface // out

	_dBusInterface = wrapDBusInterface(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusInterface
}

// GetObject gets the BusObjectProxy at object_path, if any.
//
// The function takes the following parameters:
//
//    - objectPath: object path to look up.
//
// The function returns the following values:
//
//    - dBusObject or NULL. Free with g_object_unref().
//
func (manager *DBusObjectManager) GetObject(objectPath string) *DBusObject {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_args[1]))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(manager)
	runtime.KeepAlive(objectPath)

	var _dBusObject *DBusObject // out

	_dBusObject = wrapDBusObject(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusObject
}

// ObjectPath gets the object path that manager is for.
//
// The function returns the following values:
//
//    - utf8: string owned by manager. Do not free.
//
func (manager *DBusObjectManager) ObjectPath() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(manager)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Objects gets all BusObject objects known to manager.
//
// The function returns the following values:
//
//    - list of BusObject objects. The returned list should be freed with
//      g_list_free() after each element has been freed with g_object_unref().
//
func (manager *DBusObjectManager) Objects() []*DBusObject {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(manager)

	var _list []*DBusObject // out

	_list = make([]*DBusObject, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst *DBusObject // out
		dst = wrapDBusObject(coreglib.AssumeOwnership(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}
