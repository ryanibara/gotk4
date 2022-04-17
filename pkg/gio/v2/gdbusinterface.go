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
// extern GDBusInterfaceInfo* _gotk4_gio2_DBusInterfaceIface_get_info(GDBusInterface*);
// extern GDBusObject* _gotk4_gio2_DBusInterfaceIface_dup_object(GDBusInterface*);
// extern void _gotk4_gio2_DBusInterfaceIface_set_object(GDBusInterface*, GDBusObject*);
import "C"

// glib.Type values for gdbusinterface.go.
var GTypeDBusInterface = externglib.Type(C.g_dbus_interface_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeDBusInterface, F: marshalDBusInterface},
	})
}

// DBusInterfaceOverrider contains methods that are overridable.
type DBusInterfaceOverrider interface {
	// DupObject gets the BusObject that interface_ belongs to, if any.
	//
	// The function returns the following values:
	//
	//    - dBusObject (optional) or NULL. The returned reference should be freed
	//      with g_object_unref().
	//
	DupObject() *DBusObject
	// Info gets D-Bus introspection information for the D-Bus interface
	// implemented by interface_.
	//
	// The function returns the following values:
	//
	//    - dBusInterfaceInfo Do not free.
	//
	Info() *DBusInterfaceInfo
	// SetObject sets the BusObject for interface_ to object.
	//
	// Note that interface_ will hold a weak reference to object.
	//
	// The function takes the following parameters:
	//
	//    - object (optional) or NULL.
	//
	SetObject(object DBusObjector)
}

// DBusInterface type is the base type for D-Bus interfaces both on the service
// side (see BusInterfaceSkeleton) and client side (see BusProxy).
//
// DBusInterface wraps an interface. This means the user can get the
// underlying type by calling Cast().
type DBusInterface struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*DBusInterface)(nil)
)

// DBusInterfacer describes DBusInterface's interface methods.
type DBusInterfacer interface {
	externglib.Objector

	// GetObject gets the BusObject that interface_ belongs to, if any.
	GetObject() *DBusObject
	// Info gets D-Bus introspection information for the D-Bus interface
	// implemented by interface_.
	Info() *DBusInterfaceInfo
	// SetObject sets the BusObject for interface_ to object.
	SetObject(object DBusObjector)
}

var _ DBusInterfacer = (*DBusInterface)(nil)

func ifaceInitDBusInterfacer(gifacePtr, data C.gpointer) {
	iface := (*C.GDBusInterfaceIface)(unsafe.Pointer(gifacePtr))
	iface.dup_object = (*[0]byte)(C._gotk4_gio2_DBusInterfaceIface_dup_object)
	iface.get_info = (*[0]byte)(C._gotk4_gio2_DBusInterfaceIface_get_info)
	iface.set_object = (*[0]byte)(C._gotk4_gio2_DBusInterfaceIface_set_object)
}

//export _gotk4_gio2_DBusInterfaceIface_dup_object
func _gotk4_gio2_DBusInterfaceIface_dup_object(arg0 *C.GDBusInterface) (cret *C.GDBusObject) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusInterfaceOverrider)

	dBusObject := iface.DupObject()

	if dBusObject != nil {
		cret = (*C.GDBusObject)(unsafe.Pointer(externglib.InternObject(dBusObject).Native()))
		C.g_object_ref(C.gpointer(externglib.InternObject(dBusObject).Native()))
	}

	return cret
}

//export _gotk4_gio2_DBusInterfaceIface_get_info
func _gotk4_gio2_DBusInterfaceIface_get_info(arg0 *C.GDBusInterface) (cret *C.GDBusInterfaceInfo) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusInterfaceOverrider)

	dBusInterfaceInfo := iface.Info()

	cret = (*C.GDBusInterfaceInfo)(gextras.StructNative(unsafe.Pointer(dBusInterfaceInfo)))

	return cret
}

//export _gotk4_gio2_DBusInterfaceIface_set_object
func _gotk4_gio2_DBusInterfaceIface_set_object(arg0 *C.GDBusInterface, arg1 *C.GDBusObject) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusInterfaceOverrider)

	var _object DBusObjector // out

	if arg1 != nil {
		{
			objptr := unsafe.Pointer(arg1)

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
	}

	iface.SetObject(_object)
}

func wrapDBusInterface(obj *externglib.Object) *DBusInterface {
	return &DBusInterface{
		Object: obj,
	}
}

func marshalDBusInterface(p uintptr) (interface{}, error) {
	return wrapDBusInterface(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// GetObject gets the BusObject that interface_ belongs to, if any.
//
// The function returns the following values:
//
//    - dBusObject (optional) or NULL. The returned reference should be freed
//      with g_object_unref().
//
func (interface_ *DBusInterface) GetObject() *DBusObject {
	var _arg0 *C.GDBusInterface // out
	var _cret *C.GDBusObject    // in

	_arg0 = (*C.GDBusInterface)(unsafe.Pointer(externglib.InternObject(interface_).Native()))

	_cret = C.g_dbus_interface_dup_object(_arg0)
	runtime.KeepAlive(interface_)

	var _dBusObject *DBusObject // out

	if _cret != nil {
		_dBusObject = wrapDBusObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _dBusObject
}

// Info gets D-Bus introspection information for the D-Bus interface implemented
// by interface_.
//
// The function returns the following values:
//
//    - dBusInterfaceInfo Do not free.
//
func (interface_ *DBusInterface) Info() *DBusInterfaceInfo {
	var _arg0 *C.GDBusInterface     // out
	var _cret *C.GDBusInterfaceInfo // in

	_arg0 = (*C.GDBusInterface)(unsafe.Pointer(externglib.InternObject(interface_).Native()))

	_cret = C.g_dbus_interface_get_info(_arg0)
	runtime.KeepAlive(interface_)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	_dBusInterfaceInfo = (*DBusInterfaceInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_dbus_interface_info_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_dBusInterfaceInfo)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_dbus_interface_info_unref((*C.GDBusInterfaceInfo)(intern.C))
		},
	)

	return _dBusInterfaceInfo
}

// SetObject sets the BusObject for interface_ to object.
//
// Note that interface_ will hold a weak reference to object.
//
// The function takes the following parameters:
//
//    - object (optional) or NULL.
//
func (interface_ *DBusInterface) SetObject(object DBusObjector) {
	var _arg0 *C.GDBusInterface // out
	var _arg1 *C.GDBusObject    // out

	_arg0 = (*C.GDBusInterface)(unsafe.Pointer(externglib.InternObject(interface_).Native()))
	if object != nil {
		_arg1 = (*C.GDBusObject)(unsafe.Pointer(externglib.InternObject(object).Native()))
	}

	C.g_dbus_interface_set_object(_arg0, _arg1)
	runtime.KeepAlive(interface_)
	runtime.KeepAlive(object)
}
