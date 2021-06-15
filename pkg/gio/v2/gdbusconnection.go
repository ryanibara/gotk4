// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
import "C"

// BusGetFinish finishes an operation started with g_bus_get().
//
// The returned object is a singleton, that is, shared with other callers of
// g_bus_get() and g_bus_get_sync() for @bus_type. In the event that you need a
// private message bus connection, use g_dbus_address_get_for_bus_sync() and
// g_dbus_connection_new_for_address().
//
// Note that the returned BusConnection object will (usually) have the
// BusConnection:exit-on-close property set to true.
func BusGetFinish(res AsyncResult) (DBusConnection, error) {
	var _arg1 *C.GAsyncResult    // out
	var _cret *C.GDBusConnection // in
	var _cerr *C.GError          // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	_cret = C.g_bus_get_finish(_arg1, &_cerr)

	var _dBusConnection DBusConnection // out
	var _goerr error                   // out

	_dBusConnection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(DBusConnection)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _dBusConnection, _goerr
}

// BusGetSync: synchronously connects to the message bus specified by @bus_type.
// Note that the returned object may shared with other callers, e.g. if two
// separate parts of a process calls this function with the same @bus_type, they
// will share the same object.
//
// This is a synchronous failable function. See g_bus_get() and
// g_bus_get_finish() for the asynchronous version.
//
// The returned object is a singleton, that is, shared with other callers of
// g_bus_get() and g_bus_get_sync() for @bus_type. In the event that you need a
// private message bus connection, use g_dbus_address_get_for_bus_sync() and
// g_dbus_connection_new_for_address().
//
// Note that the returned BusConnection object will (usually) have the
// BusConnection:exit-on-close property set to true.
func BusGetSync(busType BusType, cancellable Cancellable) (DBusConnection, error) {
	var _arg1 C.GBusType         // out
	var _arg2 *C.GCancellable    // out
	var _cret *C.GDBusConnection // in
	var _cerr *C.GError          // in

	_arg1 = (C.GBusType)(busType)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_bus_get_sync(_arg1, _arg2, &_cerr)

	var _dBusConnection DBusConnection // out
	var _goerr error                   // out

	_dBusConnection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(DBusConnection)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _dBusConnection, _goerr
}

// DBusInterfaceVTable: virtual table for handling properties and method calls
// for a D-Bus interface.
//
// Since 2.38, if you want to handle getting/setting D-Bus properties
// asynchronously, give nil as your get_property() or set_property() function.
// The D-Bus call will be directed to your @method_call function, with the
// provided @interface_name set to "org.freedesktop.DBus.Properties".
//
// Ownership of the BusMethodInvocation object passed to the method_call()
// function is transferred to your handler; you must call one of the methods of
// BusMethodInvocation to return a reply (possibly empty), or an error. These
// functions also take ownership of the passed-in invocation object, so unless
// the invocation object has otherwise been referenced, it will be then be
// freed. Calling one of these functions may be done within your method_call()
// implementation but it also can be done at a later point to handle the method
// asynchronously.
//
// The usual checks on the validity of the calls is performed. For `Get` calls,
// an error is automatically returned if the property does not exist or the
// permissions do not allow access. The same checks are performed for `Set`
// calls, and the provided value is also checked for being the correct type.
//
// For both `Get` and `Set` calls, the BusMethodInvocation passed to the
// @method_call handler can be queried with
// g_dbus_method_invocation_get_property_info() to get a pointer to the
// BusPropertyInfo of the property.
//
// If you have readable properties specified in your interface info, you must
// ensure that you either provide a non-nil @get_property() function or provide
// implementations of both the `Get` and `GetAll` methods on
// org.freedesktop.DBus.Properties interface in your @method_call function. Note
// that the required return type of the `Get` call is `(v)`, not the type of the
// property. `GetAll` expects a return value of type `a{sv}`.
//
// If you have writable properties specified in your interface info, you must
// ensure that you either provide a non-nil @set_property() function or provide
// an implementation of the `Set` call. If implementing the call, you must
// return the value of type G_VARIANT_TYPE_UNIT.
type DBusInterfaceVTable struct {
	native C.GDBusInterfaceVTable
}

// WrapDBusInterfaceVTable wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusInterfaceVTable(ptr unsafe.Pointer) *DBusInterfaceVTable {
	if ptr == nil {
		return nil
	}

	return (*DBusInterfaceVTable)(ptr)
}

// Native returns the underlying C source pointer.
func (d *DBusInterfaceVTable) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// DBusSubtreeVTable: virtual table for handling subtrees registered with
// g_dbus_connection_register_subtree().
type DBusSubtreeVTable struct {
	native C.GDBusSubtreeVTable
}

// WrapDBusSubtreeVTable wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusSubtreeVTable(ptr unsafe.Pointer) *DBusSubtreeVTable {
	if ptr == nil {
		return nil
	}

	return (*DBusSubtreeVTable)(ptr)
}

// Native returns the underlying C source pointer.
func (d *DBusSubtreeVTable) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}
