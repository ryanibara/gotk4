// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_dbus_object_get_type()), F: marshalDBusObjector},
	})
}

// DBusObjectOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DBusObjectOverrider interface {
	// Interface gets the D-Bus interface with name interface_name associated
	// with object, if any.
	Interface(interfaceName string) DBusInterfacer
	// Interfaces gets the D-Bus interfaces associated with object.
	Interfaces() []DBusInterfacer
	// ObjectPath gets the object path for object.
	ObjectPath() string
	InterfaceAdded(interface_ DBusInterfacer)
	InterfaceRemoved(interface_ DBusInterfacer)
}

// DBusObject type is the base type for D-Bus objects on both the service side
// (see BusObjectSkeleton) and the client side (see BusObjectProxy). It is
// essentially just a container of interfaces.
type DBusObject struct {
	*externglib.Object
}

// DBusObjector describes DBusObject's interface methods.
type DBusObjector interface {
	externglib.Objector

	// Interface gets the D-Bus interface with name interface_name associated
	// with object, if any.
	Interface(interfaceName string) DBusInterfacer
	// Interfaces gets the D-Bus interfaces associated with object.
	Interfaces() []DBusInterfacer
	// ObjectPath gets the object path for object.
	ObjectPath() string
}

var _ DBusObjector = (*DBusObject)(nil)

func wrapDBusObject(obj *externglib.Object) *DBusObject {
	return &DBusObject{
		Object: obj,
	}
}

func marshalDBusObjector(p uintptr) (interface{}, error) {
	return wrapDBusObject(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Interface gets the D-Bus interface with name interface_name associated with
// object, if any.
//
// The function takes the following parameters:
//
//    - interfaceName d-Bus interface name.
//
func (object *DBusObject) Interface(interfaceName string) DBusInterfacer {
	var _arg0 *C.GDBusObject    // out
	var _arg1 *C.gchar          // out
	var _cret *C.GDBusInterface // in

	_arg0 = (*C.GDBusObject)(unsafe.Pointer(object.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(interfaceName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_object_get_interface(_arg0, _arg1)
	runtime.KeepAlive(object)
	runtime.KeepAlive(interfaceName)

	var _dBusInterface DBusInterfacer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			rv, ok := (externglib.CastObject(object)).(DBusInterfacer)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.DBusInterfacer")
			}
			_dBusInterface = rv
		}
	}

	return _dBusInterface
}

// Interfaces gets the D-Bus interfaces associated with object.
func (object *DBusObject) Interfaces() []DBusInterfacer {
	var _arg0 *C.GDBusObject // out
	var _cret *C.GList       // in

	_arg0 = (*C.GDBusObject)(unsafe.Pointer(object.Native()))

	_cret = C.g_dbus_object_get_interfaces(_arg0)
	runtime.KeepAlive(object)

	var _list []DBusInterfacer // out

	_list = make([]DBusInterfacer, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GDBusInterface)(v)
		var dst DBusInterfacer // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gio.DBusInterfacer is nil")
			}

			object := externglib.AssumeOwnership(objptr)
			rv, ok := (externglib.CastObject(object)).(DBusInterfacer)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.DBusInterfacer")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}

// ObjectPath gets the object path for object.
func (object *DBusObject) ObjectPath() string {
	var _arg0 *C.GDBusObject // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GDBusObject)(unsafe.Pointer(object.Native()))

	_cret = C.g_dbus_object_get_object_path(_arg0)
	runtime.KeepAlive(object)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ConnectInterfaceAdded: emitted when interface is added to object.
func (object *DBusObject) ConnectInterfaceAdded(f func(iface DBusInterfacer)) externglib.SignalHandle {
	return object.Connect("interface-added", f)
}

// ConnectInterfaceRemoved: emitted when interface is removed from object.
func (object *DBusObject) ConnectInterfaceRemoved(f func(iface DBusInterfacer)) externglib.SignalHandle {
	return object.Connect("interface-removed", f)
}
