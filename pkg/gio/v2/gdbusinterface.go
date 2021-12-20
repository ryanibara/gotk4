// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_dbus_interface_get_type()), F: marshalDBusInterfacer},
	})
}

// DBusInterfaceOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DBusInterfaceOverrider interface {
	// DupObject gets the BusObject that interface_ belongs to, if any.
	//
	// The function returns the following values:
	//
	//    - dBusObject (optional) or NULL. The returned reference should be freed
	//      with g_object_unref().
	//
	DupObject() DBusObjector
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
	GetObject() DBusObjector
	// Info gets D-Bus introspection information for the D-Bus interface
	// implemented by interface_.
	Info() *DBusInterfaceInfo
	// SetObject sets the BusObject for interface_ to object.
	SetObject(object DBusObjector)
}

var _ DBusInterfacer = (*DBusInterface)(nil)

func wrapDBusInterface(obj *externglib.Object) *DBusInterface {
	return &DBusInterface{
		Object: obj,
	}
}

func marshalDBusInterfacer(p uintptr) (interface{}, error) {
	return wrapDBusInterface(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// GetObject gets the BusObject that interface_ belongs to, if any.
//
// The function returns the following values:
//
//    - dBusObject (optional) or NULL. The returned reference should be freed
//      with g_object_unref().
//
func (interface_ *DBusInterface) GetObject() DBusObjector {
	var _arg0 *C.GDBusInterface // out
	var _cret *C.GDBusObject    // in

	_arg0 = (*C.GDBusInterface)(unsafe.Pointer(interface_.Native()))

	_cret = C.g_dbus_interface_dup_object(_arg0)
	runtime.KeepAlive(interface_)

	var _dBusObject DBusObjector // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			casted := object.Cast()
			rv, ok := casted.(DBusObjector)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.DBusObjector")
			}
			_dBusObject = rv
		}
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

	_arg0 = (*C.GDBusInterface)(unsafe.Pointer(interface_.Native()))

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

	_arg0 = (*C.GDBusInterface)(unsafe.Pointer(interface_.Native()))
	if object != nil {
		_arg1 = (*C.GDBusObject)(unsafe.Pointer(object.Native()))
	}

	C.g_dbus_interface_set_object(_arg0, _arg1)
	runtime.KeepAlive(interface_)
	runtime.KeepAlive(object)
}
