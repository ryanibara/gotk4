// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
// GType _gotk4_gio2_DBusProxyTypeFunc(GDBusObjectManagerClient*, gchar*, gchar*, gpointer);
// extern void callbackDelete(gpointer);
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_dbus_object_manager_client_get_type()), F: marshalDBusObjectManagerClienter},
	})
}

// DBusObjectManagerClientOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DBusObjectManagerClientOverrider interface {
	InterfaceProxySignal(objectProxy *DBusObjectProxy, interfaceProxy *DBusProxy, senderName, signalName string, parameters *glib.Variant)
}

// DBusObjectManagerClient is used to create, monitor and delete object proxies
// for remote objects exported by a BusObjectManagerServer (or any code
// implementing the org.freedesktop.DBus.ObjectManager
// (http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-objectmanager)
// interface).
//
// Once an instance of this type has been created, you can connect to the
// BusObjectManager::object-added and BusObjectManager::object-removed signals
// and inspect the BusObjectProxy objects returned by
// g_dbus_object_manager_get_objects().
//
// If the name for a BusObjectManagerClient is not owned by anyone at object
// construction time, the default behavior is to request the message bus to
// launch an owner for the name. This behavior can be disabled using the
// G_DBUS_OBJECT_MANAGER_CLIENT_FLAGS_DO_NOT_AUTO_START flag. It's also worth
// noting that this only works if the name of interest is activatable in the
// first place. E.g. in some cases it is not possible to launch an owner for the
// requested name. In this case, BusObjectManagerClient object construction
// still succeeds but there will be no object proxies (e.g.
// g_dbus_object_manager_get_objects() returns the empty list) and the
// BusObjectManagerClient:name-owner property is NULL.
//
// The owner of the requested name can come and go (for example consider a
// system service being restarted) – BusObjectManagerClient handles this case
// too; simply connect to the #GObject::notify signal to watch for changes on
// the BusObjectManagerClient:name-owner property. When the name owner vanishes,
// the behavior is that BusObjectManagerClient:name-owner is set to NULL (this
// includes emission of the #GObject::notify signal) and then
// BusObjectManager::object-removed signals are synthesized for all currently
// existing object proxies. Since BusObjectManagerClient:name-owner is NULL when
// this happens, you can use this information to disambiguate a synthesized
// signal from a genuine signal caused by object removal on the remote
// BusObjectManager. Similarly, when a new name owner appears,
// BusObjectManager::object-added signals are synthesized while
// BusObjectManagerClient:name-owner is still NULL. Only when all object proxies
// have been added, the BusObjectManagerClient:name-owner is set to the new name
// owner (this includes emission of the #GObject::notify signal). Furthermore,
// you are guaranteed that BusObjectManagerClient:name-owner will alternate
// between a name owner (e.g. :1.42) and NULL even in the case where the name of
// interest is atomically replaced
//
// Ultimately, BusObjectManagerClient is used to obtain BusProxy instances. All
// signals (including the org.freedesktop.DBus.Properties::PropertiesChanged
// signal) delivered to BusProxy instances are guaranteed to originate from the
// name owner. This guarantee along with the behavior described above, means
// that certain race conditions including the "half the proxy is from the old
// owner and the other half is from the new owner" problem cannot happen.
//
// To avoid having the application connect to signals on the returned
// BusObjectProxy and BusProxy objects, the BusObject::interface-added,
// BusObject::interface-removed, BusProxy::g-properties-changed and
// BusProxy::g-signal signals are also emitted on the BusObjectManagerClient
// instance managing these objects. The signals emitted are
// BusObjectManager::interface-added, BusObjectManager::interface-removed,
// BusObjectManagerClient::interface-proxy-properties-changed and
// BusObjectManagerClient::interface-proxy-signal.
//
// Note that all callbacks and signals are emitted in the [thread-default main
// context][g-main-context-push-thread-default] that the BusObjectManagerClient
// object was constructed in. Additionally, the BusObjectProxy and BusProxy
// objects originating from the BusObjectManagerClient object will be created in
// the same context and, consequently, will deliver signals in the same main
// loop.
type DBusObjectManagerClient struct {
	*externglib.Object

	AsyncInitable
	DBusObjectManager
	Initable
}

func wrapDBusObjectManagerClient(obj *externglib.Object) *DBusObjectManagerClient {
	return &DBusObjectManagerClient{
		Object: obj,
		AsyncInitable: AsyncInitable{
			Object: obj,
		},
		DBusObjectManager: DBusObjectManager{
			Object: obj,
		},
		Initable: Initable{
			Object: obj,
		},
	}
}

func marshalDBusObjectManagerClienter(p uintptr) (interface{}, error) {
	return wrapDBusObjectManagerClient(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewDBusObjectManagerClientFinish finishes an operation started with
// g_dbus_object_manager_client_new().
//
// The function takes the following parameters:
//
//    - res obtained from the ReadyCallback passed to
//    g_dbus_object_manager_client_new().
//
func NewDBusObjectManagerClientFinish(res AsyncResulter) (*DBusObjectManagerClient, error) {
	var _arg1 *C.GAsyncResult       // out
	var _cret *C.GDBusObjectManager // in
	var _cerr *C.GError             // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	_cret = C.g_dbus_object_manager_client_new_finish(_arg1, &_cerr)
	runtime.KeepAlive(res)

	var _dBusObjectManagerClient *DBusObjectManagerClient // out
	var _goerr error                                      // out

	_dBusObjectManagerClient = wrapDBusObjectManagerClient(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dBusObjectManagerClient, _goerr
}

// NewDBusObjectManagerClientForBusFinish finishes an operation started with
// g_dbus_object_manager_client_new_for_bus().
//
// The function takes the following parameters:
//
//    - res obtained from the ReadyCallback passed to
//    g_dbus_object_manager_client_new_for_bus().
//
func NewDBusObjectManagerClientForBusFinish(res AsyncResulter) (*DBusObjectManagerClient, error) {
	var _arg1 *C.GAsyncResult       // out
	var _cret *C.GDBusObjectManager // in
	var _cerr *C.GError             // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	_cret = C.g_dbus_object_manager_client_new_for_bus_finish(_arg1, &_cerr)
	runtime.KeepAlive(res)

	var _dBusObjectManagerClient *DBusObjectManagerClient // out
	var _goerr error                                      // out

	_dBusObjectManagerClient = wrapDBusObjectManagerClient(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dBusObjectManagerClient, _goerr
}

// NewDBusObjectManagerClientForBusSync: like
// g_dbus_object_manager_client_new_sync() but takes a Type instead of a
// BusConnection.
//
// This is a synchronous failable constructor - the calling thread is blocked
// until a reply is received. See g_dbus_object_manager_client_new_for_bus() for
// the asynchronous version.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - busType: Type.
//    - flags: zero or more flags from the BusObjectManagerClientFlags
//    enumeration.
//    - name: owner of the control object (unique or well-known name).
//    - objectPath: object path of the control object.
//    - getProxyTypeFunc function or NULL to always construct BusProxy proxies.
//
func NewDBusObjectManagerClientForBusSync(ctx context.Context, busType BusType, flags DBusObjectManagerClientFlags, name, objectPath string, getProxyTypeFunc DBusProxyTypeFunc) (*DBusObjectManagerClient, error) {
	var _arg8 *C.GCancellable                 // out
	var _arg1 C.GBusType                      // out
	var _arg2 C.GDBusObjectManagerClientFlags // out
	var _arg3 *C.gchar                        // out
	var _arg4 *C.gchar                        // out
	var _arg5 C.GDBusProxyTypeFunc            // out
	var _arg6 C.gpointer
	var _arg7 C.GDestroyNotify
	var _cret *C.GDBusObjectManager // in
	var _cerr *C.GError             // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg8 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GBusType(busType)
	_arg2 = C.GDBusObjectManagerClientFlags(flags)
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg4))
	if getProxyTypeFunc != nil {
		_arg5 = (*[0]byte)(C._gotk4_gio2_DBusProxyTypeFunc)
		_arg6 = C.gpointer(gbox.Assign(getProxyTypeFunc))
		_arg7 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_cret = C.g_dbus_object_manager_client_new_for_bus_sync(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, &_cerr)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(busType)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(name)
	runtime.KeepAlive(objectPath)
	runtime.KeepAlive(getProxyTypeFunc)

	var _dBusObjectManagerClient *DBusObjectManagerClient // out
	var _goerr error                                      // out

	_dBusObjectManagerClient = wrapDBusObjectManagerClient(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dBusObjectManagerClient, _goerr
}

// NewDBusObjectManagerClientSync creates a new BusObjectManagerClient object.
//
// This is a synchronous failable constructor - the calling thread is blocked
// until a reply is received. See g_dbus_object_manager_client_new() for the
// asynchronous version.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - connection: BusConnection.
//    - flags: zero or more flags from the BusObjectManagerClientFlags
//    enumeration.
//    - name: owner of the control object (unique or well-known name), or NULL
//    when not using a message bus connection.
//    - objectPath: object path of the control object.
//    - getProxyTypeFunc function or NULL to always construct BusProxy proxies.
//
func NewDBusObjectManagerClientSync(ctx context.Context, connection *DBusConnection, flags DBusObjectManagerClientFlags, name, objectPath string, getProxyTypeFunc DBusProxyTypeFunc) (*DBusObjectManagerClient, error) {
	var _arg8 *C.GCancellable                 // out
	var _arg1 *C.GDBusConnection              // out
	var _arg2 C.GDBusObjectManagerClientFlags // out
	var _arg3 *C.gchar                        // out
	var _arg4 *C.gchar                        // out
	var _arg5 C.GDBusProxyTypeFunc            // out
	var _arg6 C.gpointer
	var _arg7 C.GDestroyNotify
	var _cret *C.GDBusObjectManager // in
	var _cerr *C.GError             // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg8 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))
	_arg2 = C.GDBusObjectManagerClientFlags(flags)
	if name != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg4))
	if getProxyTypeFunc != nil {
		_arg5 = (*[0]byte)(C._gotk4_gio2_DBusProxyTypeFunc)
		_arg6 = C.gpointer(gbox.Assign(getProxyTypeFunc))
		_arg7 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_cret = C.g_dbus_object_manager_client_new_sync(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, &_cerr)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(name)
	runtime.KeepAlive(objectPath)
	runtime.KeepAlive(getProxyTypeFunc)

	var _dBusObjectManagerClient *DBusObjectManagerClient // out
	var _goerr error                                      // out

	_dBusObjectManagerClient = wrapDBusObjectManagerClient(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dBusObjectManagerClient, _goerr
}

// Connection gets the BusConnection used by manager.
func (manager *DBusObjectManagerClient) Connection() *DBusConnection {
	var _arg0 *C.GDBusObjectManagerClient // out
	var _cret *C.GDBusConnection          // in

	_arg0 = (*C.GDBusObjectManagerClient)(unsafe.Pointer(manager.Native()))

	_cret = C.g_dbus_object_manager_client_get_connection(_arg0)
	runtime.KeepAlive(manager)

	var _dBusConnection *DBusConnection // out

	_dBusConnection = wrapDBusConnection(externglib.Take(unsafe.Pointer(_cret)))

	return _dBusConnection
}

// Flags gets the flags that manager was constructed with.
func (manager *DBusObjectManagerClient) Flags() DBusObjectManagerClientFlags {
	var _arg0 *C.GDBusObjectManagerClient     // out
	var _cret C.GDBusObjectManagerClientFlags // in

	_arg0 = (*C.GDBusObjectManagerClient)(unsafe.Pointer(manager.Native()))

	_cret = C.g_dbus_object_manager_client_get_flags(_arg0)
	runtime.KeepAlive(manager)

	var _dBusObjectManagerClientFlags DBusObjectManagerClientFlags // out

	_dBusObjectManagerClientFlags = DBusObjectManagerClientFlags(_cret)

	return _dBusObjectManagerClientFlags
}

// Name gets the name that manager is for, or NULL if not a message bus
// connection.
func (manager *DBusObjectManagerClient) Name() string {
	var _arg0 *C.GDBusObjectManagerClient // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.GDBusObjectManagerClient)(unsafe.Pointer(manager.Native()))

	_cret = C.g_dbus_object_manager_client_get_name(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// NameOwner: unique name that owns the name that manager is for or NULL if
// no-one currently owns that name. You can connect to the #GObject::notify
// signal to track changes to the BusObjectManagerClient:name-owner property.
func (manager *DBusObjectManagerClient) NameOwner() string {
	var _arg0 *C.GDBusObjectManagerClient // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.GDBusObjectManagerClient)(unsafe.Pointer(manager.Native()))

	_cret = C.g_dbus_object_manager_client_get_name_owner(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// ConnectInterfaceProxyPropertiesChanged: emitted when one or more D-Bus
// properties on proxy changes. The local cache has already been updated when
// this signal fires. Note that both changed_properties and
// invalidated_properties are guaranteed to never be NULL (either may be empty
// though).
//
// This signal exists purely as a convenience to avoid having to connect signals
// to all interface proxies managed by manager.
//
// This signal is emitted in the [thread-default main
// context][g-main-context-push-thread-default] that manager was constructed in.
func (manager *DBusObjectManagerClient) ConnectInterfaceProxyPropertiesChanged(f func(objectProxy DBusObjectProxy, interfaceProxy DBusProxy, changedProperties *glib.Variant, invalidatedProperties []string)) externglib.SignalHandle {
	return manager.Connect("interface-proxy-properties-changed", f)
}

// ConnectInterfaceProxySignal: emitted when a D-Bus signal is received on
// interface_proxy.
//
// This signal exists purely as a convenience to avoid having to connect signals
// to all interface proxies managed by manager.
//
// This signal is emitted in the [thread-default main
// context][g-main-context-push-thread-default] that manager was constructed in.
func (manager *DBusObjectManagerClient) ConnectInterfaceProxySignal(f func(objectProxy DBusObjectProxy, interfaceProxy DBusProxy, senderName, signalName string, parameters *glib.Variant)) externglib.SignalHandle {
	return manager.Connect("interface-proxy-signal", f)
}

// NewDBusObjectManagerClient: asynchronously creates a new
// BusObjectManagerClient object.
//
// This is an asynchronous failable constructor. When the result is ready,
// callback will be invoked in the [thread-default main
// context][g-main-context-push-thread-default] of the thread you are calling
// this method from. You can then call g_dbus_object_manager_client_new_finish()
// to get the result. See g_dbus_object_manager_client_new_sync() for the
// synchronous version.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - connection: BusConnection.
//    - flags: zero or more flags from the BusObjectManagerClientFlags
//    enumeration.
//    - name: owner of the control object (unique or well-known name).
//    - objectPath: object path of the control object.
//    - getProxyTypeFunc function or NULL to always construct BusProxy proxies.
//    - callback to call when the request is satisfied.
//
func NewDBusObjectManagerClient(ctx context.Context, connection *DBusConnection, flags DBusObjectManagerClientFlags, name, objectPath string, getProxyTypeFunc DBusProxyTypeFunc, callback AsyncReadyCallback) {
	var _arg8 *C.GCancellable                 // out
	var _arg1 *C.GDBusConnection              // out
	var _arg2 C.GDBusObjectManagerClientFlags // out
	var _arg3 *C.gchar                        // out
	var _arg4 *C.gchar                        // out
	var _arg5 C.GDBusProxyTypeFunc            // out
	var _arg6 C.gpointer
	var _arg7 C.GDestroyNotify
	var _arg9 C.GAsyncReadyCallback // out
	var _arg10 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg8 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))
	_arg2 = C.GDBusObjectManagerClientFlags(flags)
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg4))
	if getProxyTypeFunc != nil {
		_arg5 = (*[0]byte)(C._gotk4_gio2_DBusProxyTypeFunc)
		_arg6 = C.gpointer(gbox.Assign(getProxyTypeFunc))
		_arg7 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}
	if callback != nil {
		_arg9 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg10 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_dbus_object_manager_client_new(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(name)
	runtime.KeepAlive(objectPath)
	runtime.KeepAlive(getProxyTypeFunc)
	runtime.KeepAlive(callback)
}

// NewDBusObjectManagerClientForBus: like g_dbus_object_manager_client_new() but
// takes a Type instead of a BusConnection.
//
// This is an asynchronous failable constructor. When the result is ready,
// callback will be invoked in the [thread-default main
// loop][g-main-context-push-thread-default] of the thread you are calling this
// method from. You can then call
// g_dbus_object_manager_client_new_for_bus_finish() to get the result. See
// g_dbus_object_manager_client_new_for_bus_sync() for the synchronous version.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - busType: Type.
//    - flags: zero or more flags from the BusObjectManagerClientFlags
//    enumeration.
//    - name: owner of the control object (unique or well-known name).
//    - objectPath: object path of the control object.
//    - getProxyTypeFunc function or NULL to always construct BusProxy proxies.
//    - callback to call when the request is satisfied.
//
func NewDBusObjectManagerClientForBus(ctx context.Context, busType BusType, flags DBusObjectManagerClientFlags, name, objectPath string, getProxyTypeFunc DBusProxyTypeFunc, callback AsyncReadyCallback) {
	var _arg8 *C.GCancellable                 // out
	var _arg1 C.GBusType                      // out
	var _arg2 C.GDBusObjectManagerClientFlags // out
	var _arg3 *C.gchar                        // out
	var _arg4 *C.gchar                        // out
	var _arg5 C.GDBusProxyTypeFunc            // out
	var _arg6 C.gpointer
	var _arg7 C.GDestroyNotify
	var _arg9 C.GAsyncReadyCallback // out
	var _arg10 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg8 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GBusType(busType)
	_arg2 = C.GDBusObjectManagerClientFlags(flags)
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg4))
	if getProxyTypeFunc != nil {
		_arg5 = (*[0]byte)(C._gotk4_gio2_DBusProxyTypeFunc)
		_arg6 = C.gpointer(gbox.Assign(getProxyTypeFunc))
		_arg7 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}
	if callback != nil {
		_arg9 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg10 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_dbus_object_manager_client_new_for_bus(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(busType)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(name)
	runtime.KeepAlive(objectPath)
	runtime.KeepAlive(getProxyTypeFunc)
	runtime.KeepAlive(callback)
}
