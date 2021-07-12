// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
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
		{T: externglib.Type(C.g_dbus_object_proxy_get_type()), F: marshalDBusObjectProxier},
	})
}

// DBusObjectProxier describes DBusObjectProxy's methods.
type DBusObjectProxier interface {
	// Connection gets the connection that @proxy is for.
	Connection() *DBusConnection
}

// DBusObjectProxy is an object used to represent a remote object with one or
// more D-Bus interfaces. Normally, you don't instantiate a BusObjectProxy
// yourself - typically BusObjectManagerClient is used to obtain it.
type DBusObjectProxy struct {
	*externglib.Object

	DBusObject
}

var (
	_ DBusObjectProxier = (*DBusObjectProxy)(nil)
	_ gextras.Nativer   = (*DBusObjectProxy)(nil)
)

func wrapDBusObjectProxy(obj *externglib.Object) *DBusObjectProxy {
	return &DBusObjectProxy{
		Object: obj,
		DBusObject: DBusObject{
			Object: obj,
		},
	}
}

func marshalDBusObjectProxier(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDBusObjectProxy(obj), nil
}

// NewDBusObjectProxy creates a new BusObjectProxy for the given connection and
// object path.
func NewDBusObjectProxy(connection DBusConnectioner, objectPath string) *DBusObjectProxy {
	var _arg1 *C.GDBusConnection  // out
	var _arg2 *C.gchar            // out
	var _cret *C.GDBusObjectProxy // in

	_arg1 = (*C.GDBusConnection)(unsafe.Pointer((connection).(gextras.Nativer).Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))

	_cret = C.g_dbus_object_proxy_new(_arg1, _arg2)

	var _dBusObjectProxy *DBusObjectProxy // out

	_dBusObjectProxy = wrapDBusObjectProxy(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusObjectProxy
}

// Connection gets the connection that @proxy is for.
func (proxy *DBusObjectProxy) Connection() *DBusConnection {
	var _arg0 *C.GDBusObjectProxy // out
	var _cret *C.GDBusConnection  // in

	_arg0 = (*C.GDBusObjectProxy)(unsafe.Pointer(proxy.Native()))

	_cret = C.g_dbus_object_proxy_get_connection(_arg0)

	var _dBusConnection *DBusConnection // out

	_dBusConnection = wrapDBusConnection(externglib.Take(unsafe.Pointer(_cret)))

	return _dBusConnection
}
