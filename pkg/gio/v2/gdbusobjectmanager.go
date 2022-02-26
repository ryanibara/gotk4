// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern GDBusInterface* _gotk4_gio2_DBusObjectManagerIface_get_interface(GDBusObjectManager*, gchar*, gchar*);
// extern GDBusObject* _gotk4_gio2_DBusObjectManagerIface_get_object(GDBusObjectManager*, gchar*);
// extern GList* _gotk4_gio2_DBusObjectManagerIface_get_objects(GDBusObjectManager*);
// extern gchar* _gotk4_gio2_DBusObjectManagerIface_get_object_path(GDBusObjectManager*);
// extern void _gotk4_gio2_DBusObjectManagerIface_interface_added(GDBusObjectManager*, GDBusObject*, GDBusInterface*);
// extern void _gotk4_gio2_DBusObjectManagerIface_interface_removed(GDBusObjectManager*, GDBusObject*, GDBusInterface*);
// extern void _gotk4_gio2_DBusObjectManagerIface_object_added(GDBusObjectManager*, GDBusObject*);
// extern void _gotk4_gio2_DBusObjectManagerIface_object_removed(GDBusObjectManager*, GDBusObject*);
// extern void _gotk4_gio2_DBusObjectManager_ConnectInterfaceAdded(gpointer, GDBusObject*, GDBusInterface*, guintptr);
// extern void _gotk4_gio2_DBusObjectManager_ConnectInterfaceRemoved(gpointer, GDBusObject*, GDBusInterface*, guintptr);
// extern void _gotk4_gio2_DBusObjectManager_ConnectObjectAdded(gpointer, GDBusObject*, guintptr);
// extern void _gotk4_gio2_DBusObjectManager_ConnectObjectRemoved(gpointer, GDBusObject*, guintptr);
import "C"

// glib.Type values for gdbusobjectmanager.go.
var GTypeDBusObjectManager = externglib.Type(C.g_dbus_object_manager_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeDBusObjectManager, F: marshalDBusObjectManager},
	})
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
	Interface(objectPath, interfaceName string) DBusInterfacer
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
	GetObject(objectPath string) DBusObjector
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
	Objects() []DBusObjector
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
type DBusObjectManager struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*DBusObjectManager)(nil)
)

// DBusObjectManagerer describes DBusObjectManager's interface methods.
type DBusObjectManagerer interface {
	externglib.Objector

	// Interface gets the interface proxy for interface_name at object_path, if
	// any.
	Interface(objectPath, interfaceName string) DBusInterfacer
	// GetObject gets the BusObjectProxy at object_path, if any.
	GetObject(objectPath string) DBusObjector
	// ObjectPath gets the object path that manager is for.
	ObjectPath() string
	// Objects gets all BusObject objects known to manager.
	Objects() []DBusObjector

	// Interface-added: emitted when interface is added to object.
	ConnectInterfaceAdded(func(object DBusObjector, iface DBusInterfacer)) externglib.SignalHandle
	// Interface-removed: emitted when interface has been removed from object.
	ConnectInterfaceRemoved(func(object DBusObjector, iface DBusInterfacer)) externglib.SignalHandle
	// Object-added: emitted when object is added to manager.
	ConnectObjectAdded(func(object DBusObjector)) externglib.SignalHandle
	// Object-removed: emitted when object is removed from manager.
	ConnectObjectRemoved(func(object DBusObjector)) externglib.SignalHandle
}

var _ DBusObjectManagerer = (*DBusObjectManager)(nil)

func ifaceInitDBusObjectManagerer(gifacePtr, data C.gpointer) {
	iface := (*C.GDBusObjectManagerIface)(unsafe.Pointer(gifacePtr))
	iface.get_interface = (*[0]byte)(C._gotk4_gio2_DBusObjectManagerIface_get_interface)
	iface.get_object = (*[0]byte)(C._gotk4_gio2_DBusObjectManagerIface_get_object)
	iface.get_object_path = (*[0]byte)(C._gotk4_gio2_DBusObjectManagerIface_get_object_path)
	iface.get_objects = (*[0]byte)(C._gotk4_gio2_DBusObjectManagerIface_get_objects)
	iface.interface_added = (*[0]byte)(C._gotk4_gio2_DBusObjectManagerIface_interface_added)
	iface.interface_removed = (*[0]byte)(C._gotk4_gio2_DBusObjectManagerIface_interface_removed)
	iface.object_added = (*[0]byte)(C._gotk4_gio2_DBusObjectManagerIface_object_added)
	iface.object_removed = (*[0]byte)(C._gotk4_gio2_DBusObjectManagerIface_object_removed)
}

//export _gotk4_gio2_DBusObjectManagerIface_get_interface
func _gotk4_gio2_DBusObjectManagerIface_get_interface(arg0 *C.GDBusObjectManager, arg1 *C.gchar, arg2 *C.gchar) (cret *C.GDBusInterface) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectManagerOverrider)

	var _objectPath string    // out
	var _interfaceName string // out

	_objectPath = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_interfaceName = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	dBusInterface := iface.Interface(_objectPath, _interfaceName)

	cret = (*C.GDBusInterface)(unsafe.Pointer(externglib.InternObject(dBusInterface).Native()))
	C.g_object_ref(C.gpointer(externglib.InternObject(dBusInterface).Native()))

	return cret
}

//export _gotk4_gio2_DBusObjectManagerIface_get_object
func _gotk4_gio2_DBusObjectManagerIface_get_object(arg0 *C.GDBusObjectManager, arg1 *C.gchar) (cret *C.GDBusObject) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectManagerOverrider)

	var _objectPath string // out

	_objectPath = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	dBusObject := iface.GetObject(_objectPath)

	cret = (*C.GDBusObject)(unsafe.Pointer(externglib.InternObject(dBusObject).Native()))
	C.g_object_ref(C.gpointer(externglib.InternObject(dBusObject).Native()))

	return cret
}

//export _gotk4_gio2_DBusObjectManagerIface_get_object_path
func _gotk4_gio2_DBusObjectManagerIface_get_object_path(arg0 *C.GDBusObjectManager) (cret *C.gchar) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectManagerOverrider)

	utf8 := iface.ObjectPath()

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_gio2_DBusObjectManagerIface_get_objects
func _gotk4_gio2_DBusObjectManagerIface_get_objects(arg0 *C.GDBusObjectManager) (cret *C.GList) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectManagerOverrider)

	list := iface.Objects()

	for i := len(list) - 1; i >= 0; i-- {
		src := list[i]
		var dst *C.GDBusObject // out
		dst = (*C.GDBusObject)(unsafe.Pointer(externglib.InternObject(src).Native()))
		C.g_object_ref(C.gpointer(externglib.InternObject(src).Native()))
		cret = C.g_list_prepend(cret, C.gpointer(unsafe.Pointer(dst)))
	}

	return cret
}

//export _gotk4_gio2_DBusObjectManagerIface_interface_added
func _gotk4_gio2_DBusObjectManagerIface_interface_added(arg0 *C.GDBusObjectManager, arg1 *C.GDBusObject, arg2 *C.GDBusInterface) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectManagerOverrider)

	var _object DBusObjector       // out
	var _interface_ DBusInterfacer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusObjector is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
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

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
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
func _gotk4_gio2_DBusObjectManagerIface_interface_removed(arg0 *C.GDBusObjectManager, arg1 *C.GDBusObject, arg2 *C.GDBusInterface) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectManagerOverrider)

	var _object DBusObjector       // out
	var _interface_ DBusInterfacer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusObjector is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
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

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
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
func _gotk4_gio2_DBusObjectManagerIface_object_added(arg0 *C.GDBusObjectManager, arg1 *C.GDBusObject) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectManagerOverrider)

	var _object DBusObjector // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusObjector is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
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
func _gotk4_gio2_DBusObjectManagerIface_object_removed(arg0 *C.GDBusObjectManager, arg1 *C.GDBusObject) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectManagerOverrider)

	var _object DBusObjector // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusObjector is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
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

func wrapDBusObjectManager(obj *externglib.Object) *DBusObjectManager {
	return &DBusObjectManager{
		Object: obj,
	}
}

func marshalDBusObjectManager(p uintptr) (interface{}, error) {
	return wrapDBusObjectManager(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_DBusObjectManager_ConnectInterfaceAdded
func _gotk4_gio2_DBusObjectManager_ConnectInterfaceAdded(arg0 C.gpointer, arg1 *C.GDBusObject, arg2 *C.GDBusInterface, arg3 C.guintptr) {
	var f func(object DBusObjector, iface DBusInterfacer)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
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

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
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

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
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

// ConnectInterfaceAdded: emitted when interface is added to object.
//
// This signal exists purely as a convenience to avoid having to connect signals
// to all objects managed by manager.
func (manager *DBusObjectManager) ConnectInterfaceAdded(f func(object DBusObjector, iface DBusInterfacer)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(manager, "interface-added", false, unsafe.Pointer(C._gotk4_gio2_DBusObjectManager_ConnectInterfaceAdded), f)
}

//export _gotk4_gio2_DBusObjectManager_ConnectInterfaceRemoved
func _gotk4_gio2_DBusObjectManager_ConnectInterfaceRemoved(arg0 C.gpointer, arg1 *C.GDBusObject, arg2 *C.GDBusInterface, arg3 C.guintptr) {
	var f func(object DBusObjector, iface DBusInterfacer)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
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

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
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

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
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

// ConnectInterfaceRemoved: emitted when interface has been removed from object.
//
// This signal exists purely as a convenience to avoid having to connect signals
// to all objects managed by manager.
func (manager *DBusObjectManager) ConnectInterfaceRemoved(f func(object DBusObjector, iface DBusInterfacer)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(manager, "interface-removed", false, unsafe.Pointer(C._gotk4_gio2_DBusObjectManager_ConnectInterfaceRemoved), f)
}

//export _gotk4_gio2_DBusObjectManager_ConnectObjectAdded
func _gotk4_gio2_DBusObjectManager_ConnectObjectAdded(arg0 C.gpointer, arg1 *C.GDBusObject, arg2 C.guintptr) {
	var f func(object DBusObjector)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
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

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
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

// ConnectObjectAdded: emitted when object is added to manager.
func (manager *DBusObjectManager) ConnectObjectAdded(f func(object DBusObjector)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(manager, "object-added", false, unsafe.Pointer(C._gotk4_gio2_DBusObjectManager_ConnectObjectAdded), f)
}

//export _gotk4_gio2_DBusObjectManager_ConnectObjectRemoved
func _gotk4_gio2_DBusObjectManager_ConnectObjectRemoved(arg0 C.gpointer, arg1 *C.GDBusObject, arg2 C.guintptr) {
	var f func(object DBusObjector)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
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

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
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

// ConnectObjectRemoved: emitted when object is removed from manager.
func (manager *DBusObjectManager) ConnectObjectRemoved(f func(object DBusObjector)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(manager, "object-removed", false, unsafe.Pointer(C._gotk4_gio2_DBusObjectManager_ConnectObjectRemoved), f)
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
func (manager *DBusObjectManager) Interface(objectPath, interfaceName string) DBusInterfacer {
	var _arg0 *C.GDBusObjectManager // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.gchar              // out
	var _cret *C.GDBusInterface     // in

	_arg0 = (*C.GDBusObjectManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(interfaceName)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_dbus_object_manager_get_interface(_arg0, _arg1, _arg2)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(objectPath)
	runtime.KeepAlive(interfaceName)

	var _dBusInterface DBusInterfacer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.DBusInterfacer is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DBusInterfacer)
			return ok
		})
		rv, ok := casted.(DBusInterfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfacer")
		}
		_dBusInterface = rv
	}

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
func (manager *DBusObjectManager) GetObject(objectPath string) DBusObjector {
	var _arg0 *C.GDBusObjectManager // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusObject        // in

	_arg0 = (*C.GDBusObjectManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_object_manager_get_object(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(objectPath)

	var _dBusObject DBusObjector // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.DBusObjector is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DBusObjector)
			return ok
		})
		rv, ok := casted.(DBusObjector)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusObjector")
		}
		_dBusObject = rv
	}

	return _dBusObject
}

// ObjectPath gets the object path that manager is for.
//
// The function returns the following values:
//
//    - utf8: string owned by manager. Do not free.
//
func (manager *DBusObjectManager) ObjectPath() string {
	var _arg0 *C.GDBusObjectManager // out
	var _cret *C.gchar              // in

	_arg0 = (*C.GDBusObjectManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.g_dbus_object_manager_get_object_path(_arg0)
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
func (manager *DBusObjectManager) Objects() []DBusObjector {
	var _arg0 *C.GDBusObjectManager // out
	var _cret *C.GList              // in

	_arg0 = (*C.GDBusObjectManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.g_dbus_object_manager_get_objects(_arg0)
	runtime.KeepAlive(manager)

	var _list []DBusObjector // out

	_list = make([]DBusObjector, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GDBusObject)(v)
		var dst DBusObjector // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gio.DBusObjector is nil")
			}

			object := externglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(DBusObjector)
				return ok
			})
			rv, ok := casted.(DBusObjector)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusObjector")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}
