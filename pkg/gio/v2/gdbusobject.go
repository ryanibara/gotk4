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
// extern GDBusInterface* _gotk4_gio2_DBusObjectIface_get_interface(GDBusObject*, gchar*);
// extern GList* _gotk4_gio2_DBusObjectIface_get_interfaces(GDBusObject*);
// extern gchar* _gotk4_gio2_DBusObjectIface_get_object_path(GDBusObject*);
// extern void _gotk4_gio2_DBusObjectIface_interface_added(GDBusObject*, GDBusInterface*);
// extern void _gotk4_gio2_DBusObjectIface_interface_removed(GDBusObject*, GDBusInterface*);
// extern void _gotk4_gio2_DBusObject_ConnectInterfaceAdded(gpointer, GDBusInterface*, guintptr);
// extern void _gotk4_gio2_DBusObject_ConnectInterfaceRemoved(gpointer, GDBusInterface*, guintptr);
import "C"

// glib.Type values for gdbusobject.go.
var GTypeDBusObject = externglib.Type(C.g_dbus_object_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeDBusObject, F: marshalDBusObject},
	})
}

// DBusObjectOverrider contains methods that are overridable.
type DBusObjectOverrider interface {
	externglib.Objector
	// Interface gets the D-Bus interface with name interface_name associated
	// with object, if any.
	//
	// The function takes the following parameters:
	//
	//    - interfaceName d-Bus interface name.
	//
	// The function returns the following values:
	//
	//    - dBusInterface (optional): NULL if not found, otherwise a BusInterface
	//      that must be freed with g_object_unref().
	//
	Interface(interfaceName string) DBusInterfaceOverrider
	// Interfaces gets the D-Bus interfaces associated with object.
	//
	// The function returns the following values:
	//
	//    - list of BusInterface instances. The returned list must be freed by
	//      g_list_free() after each element has been freed with
	//      g_object_unref().
	//
	Interfaces() []DBusInterfaceOverrider
	// ObjectPath gets the object path for object.
	//
	// The function returns the following values:
	//
	//    - utf8: string owned by object. Do not free.
	//
	ObjectPath() string
	// The function takes the following parameters:
	//
	InterfaceAdded(interface_ DBusInterfaceOverrider)
	// The function takes the following parameters:
	//
	InterfaceRemoved(interface_ DBusInterfaceOverrider)
}

// WrapDBusObjectOverrider wraps the DBusObjectOverrider
// interface implementation to access the instance methods.
func WrapDBusObjectOverrider(obj DBusObjectOverrider) *DBusObject {
	return wrapDBusObject(externglib.BaseObject(obj))
}

// DBusObject type is the base type for D-Bus objects on both the service side
// (see BusObjectSkeleton) and the client side (see BusObjectProxy). It is
// essentially just a container of interfaces.
type DBusObject struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*DBusObject)(nil)
)

// DBusObjector describes DBusObject's interface methods.
type DBusObjector interface {
	externglib.Objector

	// Interface gets the D-Bus interface with name interface_name associated
	// with object, if any.
	Interface(interfaceName string) DBusInterfaceOverrider
	// Interfaces gets the D-Bus interfaces associated with object.
	Interfaces() []DBusInterfaceOverrider
	// ObjectPath gets the object path for object.
	ObjectPath() string

	// Interface-added is emitted when interface is added to object.
	ConnectInterfaceAdded(func(iface DBusInterfaceOverrider)) externglib.SignalHandle
	// Interface-removed is emitted when interface is removed from object.
	ConnectInterfaceRemoved(func(iface DBusInterfaceOverrider)) externglib.SignalHandle
}

var _ DBusObjector = (*DBusObject)(nil)

func ifaceInitDBusObjector(gifacePtr, data C.gpointer) {
	iface := (*C.GDBusObjectIface)(unsafe.Pointer(gifacePtr))
	iface.get_interface = (*[0]byte)(C._gotk4_gio2_DBusObjectIface_get_interface)
	iface.get_interfaces = (*[0]byte)(C._gotk4_gio2_DBusObjectIface_get_interfaces)
	iface.get_object_path = (*[0]byte)(C._gotk4_gio2_DBusObjectIface_get_object_path)
	iface.interface_added = (*[0]byte)(C._gotk4_gio2_DBusObjectIface_interface_added)
	iface.interface_removed = (*[0]byte)(C._gotk4_gio2_DBusObjectIface_interface_removed)
}

//export _gotk4_gio2_DBusObjectIface_get_interface
func _gotk4_gio2_DBusObjectIface_get_interface(arg0 *C.GDBusObject, arg1 *C.gchar) (cret *C.GDBusInterface) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectOverrider)

	var _interfaceName string // out

	_interfaceName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	dBusInterface := iface.Interface(_interfaceName)

	if dBusInterface != nil {
		cret = (*C.GDBusInterface)(unsafe.Pointer(externglib.InternObject(dBusInterface).Native()))
		C.g_object_ref(C.gpointer(externglib.InternObject(dBusInterface).Native()))
	}

	return cret
}

//export _gotk4_gio2_DBusObjectIface_get_interfaces
func _gotk4_gio2_DBusObjectIface_get_interfaces(arg0 *C.GDBusObject) (cret *C.GList) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectOverrider)

	list := iface.Interfaces()

	for i := len(list) - 1; i >= 0; i-- {
		src := list[i]
		var dst *C.GDBusInterface // out
		dst = (*C.GDBusInterface)(unsafe.Pointer(externglib.InternObject(src).Native()))
		C.g_object_ref(C.gpointer(externglib.InternObject(src).Native()))
		cret = C.g_list_prepend(cret, C.gpointer(unsafe.Pointer(dst)))
	}

	return cret
}

//export _gotk4_gio2_DBusObjectIface_get_object_path
func _gotk4_gio2_DBusObjectIface_get_object_path(arg0 *C.GDBusObject) (cret *C.gchar) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectOverrider)

	utf8 := iface.ObjectPath()

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_gio2_DBusObjectIface_interface_added
func _gotk4_gio2_DBusObjectIface_interface_added(arg0 *C.GDBusObject, arg1 *C.GDBusInterface) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectOverrider)

	var _interface_ DBusInterfaceOverrider // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusInterfacer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DBusInterfaceOverrider)
			return ok
		})
		rv, ok := casted.(DBusInterfaceOverrider)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfacer")
		}
		_interface_ = rv
	}

	iface.InterfaceAdded(_interface_)
}

//export _gotk4_gio2_DBusObjectIface_interface_removed
func _gotk4_gio2_DBusObjectIface_interface_removed(arg0 *C.GDBusObject, arg1 *C.GDBusInterface) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectOverrider)

	var _interface_ DBusInterfaceOverrider // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusInterfacer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DBusInterfaceOverrider)
			return ok
		})
		rv, ok := casted.(DBusInterfaceOverrider)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfacer")
		}
		_interface_ = rv
	}

	iface.InterfaceRemoved(_interface_)
}

func wrapDBusObject(obj *externglib.Object) *DBusObject {
	return &DBusObject{
		Object: obj,
	}
}

func marshalDBusObject(p uintptr) (interface{}, error) {
	return wrapDBusObject(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_DBusObject_ConnectInterfaceAdded
func _gotk4_gio2_DBusObject_ConnectInterfaceAdded(arg0 C.gpointer, arg1 *C.GDBusInterface, arg2 C.guintptr) {
	var f func(iface DBusInterfaceOverrider)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(iface DBusInterfaceOverrider))
	}

	var _iface DBusInterfaceOverrider // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusInterfacer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DBusInterfaceOverrider)
			return ok
		})
		rv, ok := casted.(DBusInterfaceOverrider)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfacer")
		}
		_iface = rv
	}

	f(_iface)
}

// ConnectInterfaceAdded is emitted when interface is added to object.
func (object *DBusObject) ConnectInterfaceAdded(f func(iface DBusInterfaceOverrider)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(object, "interface-added", false, unsafe.Pointer(C._gotk4_gio2_DBusObject_ConnectInterfaceAdded), f)
}

//export _gotk4_gio2_DBusObject_ConnectInterfaceRemoved
func _gotk4_gio2_DBusObject_ConnectInterfaceRemoved(arg0 C.gpointer, arg1 *C.GDBusInterface, arg2 C.guintptr) {
	var f func(iface DBusInterfaceOverrider)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(iface DBusInterfaceOverrider))
	}

	var _iface DBusInterfaceOverrider // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusInterfacer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DBusInterfaceOverrider)
			return ok
		})
		rv, ok := casted.(DBusInterfaceOverrider)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfacer")
		}
		_iface = rv
	}

	f(_iface)
}

// ConnectInterfaceRemoved is emitted when interface is removed from object.
func (object *DBusObject) ConnectInterfaceRemoved(f func(iface DBusInterfaceOverrider)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(object, "interface-removed", false, unsafe.Pointer(C._gotk4_gio2_DBusObject_ConnectInterfaceRemoved), f)
}

// Interface gets the D-Bus interface with name interface_name associated with
// object, if any.
//
// The function takes the following parameters:
//
//    - interfaceName d-Bus interface name.
//
// The function returns the following values:
//
//    - dBusInterface (optional): NULL if not found, otherwise a BusInterface
//      that must be freed with g_object_unref().
//
func (object *DBusObject) Interface(interfaceName string) DBusInterfaceOverrider {
	var _arg0 *C.GDBusObject    // out
	var _arg1 *C.gchar          // out
	var _cret *C.GDBusInterface // in

	_arg0 = (*C.GDBusObject)(unsafe.Pointer(externglib.InternObject(object).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(interfaceName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_object_get_interface(_arg0, _arg1)
	runtime.KeepAlive(object)
	runtime.KeepAlive(interfaceName)

	var _dBusInterface DBusInterfaceOverrider // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(DBusInterfaceOverrider)
				return ok
			})
			rv, ok := casted.(DBusInterfaceOverrider)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfacer")
			}
			_dBusInterface = rv
		}
	}

	return _dBusInterface
}

// Interfaces gets the D-Bus interfaces associated with object.
//
// The function returns the following values:
//
//    - list of BusInterface instances. The returned list must be freed by
//      g_list_free() after each element has been freed with g_object_unref().
//
func (object *DBusObject) Interfaces() []DBusInterfaceOverrider {
	var _arg0 *C.GDBusObject // out
	var _cret *C.GList       // in

	_arg0 = (*C.GDBusObject)(unsafe.Pointer(externglib.InternObject(object).Native()))

	_cret = C.g_dbus_object_get_interfaces(_arg0)
	runtime.KeepAlive(object)

	var _list []DBusInterfaceOverrider // out

	_list = make([]DBusInterfaceOverrider, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GDBusInterface)(v)
		var dst DBusInterfaceOverrider // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gio.DBusInterfacer is nil")
			}

			object := externglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(DBusInterfaceOverrider)
				return ok
			})
			rv, ok := casted.(DBusInterfaceOverrider)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfacer")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}

// ObjectPath gets the object path for object.
//
// The function returns the following values:
//
//    - utf8: string owned by object. Do not free.
//
func (object *DBusObject) ObjectPath() string {
	var _arg0 *C.GDBusObject // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GDBusObject)(unsafe.Pointer(externglib.InternObject(object).Native()))

	_cret = C.g_dbus_object_get_object_path(_arg0)
	runtime.KeepAlive(object)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
