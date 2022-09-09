// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern void callbackDelete(gpointer);
// extern void _gotk4_gio2_DBusObjectManagerClient_ConnectInterfaceProxySignal(gpointer, GDBusObjectProxy*, GDBusProxy*, gchar*, gchar*, GVariant*, guintptr);
// extern void _gotk4_gio2_DBusObjectManagerClient_ConnectInterfaceProxyPropertiesChanged(gpointer, GDBusObjectProxy*, GDBusProxy*, GVariant*, gchar**, guintptr);
// extern void _gotk4_gio2_DBusObjectManagerClientClass_interface_proxy_signal(GDBusObjectManagerClient*, GDBusObjectProxy*, GDBusProxy*, gchar*, gchar*, GVariant*);
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
// extern GType _gotk4_gio2_DBusProxyTypeFunc(GDBusObjectManagerClient*, gchar*, gchar*, gpointer);
// void _gotk4_gio2_DBusObjectManagerClient_virtual_interface_proxy_signal(void* fnptr, GDBusObjectManagerClient* arg0, GDBusObjectProxy* arg1, GDBusProxy* arg2, gchar* arg3, gchar* arg4, GVariant* arg5) {
//   ((void (*)(GDBusObjectManagerClient*, GDBusObjectProxy*, GDBusProxy*, gchar*, gchar*, GVariant*))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5);
// };
import "C"

// GType values.
var (
	GTypeDBusObjectManagerClient = coreglib.Type(C.g_dbus_object_manager_client_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDBusObjectManagerClient, F: marshalDBusObjectManagerClient},
	})
}

// DBusObjectManagerClientOverrides contains methods that are overridable.
type DBusObjectManagerClientOverrides struct {
	// The function takes the following parameters:
	//
	//    - objectProxy
	//    - interfaceProxy
	//    - senderName
	//    - signalName
	//    - parameters
	//
	InterfaceProxySignal func(objectProxy *DBusObjectProxy, interfaceProxy *DBusProxy, senderName, signalName string, parameters *glib.Variant)
}

func defaultDBusObjectManagerClientOverrides(v *DBusObjectManagerClient) DBusObjectManagerClientOverrides {
	return DBusObjectManagerClientOverrides{
		InterfaceProxySignal: v.interfaceProxySignal,
	}
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
	_ [0]func() // equal guard
	*coreglib.Object

	AsyncInitable
	DBusObjectManager
	Initable
}

var (
	_ coreglib.Objector = (*DBusObjectManagerClient)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DBusObjectManagerClient, *DBusObjectManagerClientClass, DBusObjectManagerClientOverrides](
		GTypeDBusObjectManagerClient,
		initDBusObjectManagerClientClass,
		wrapDBusObjectManagerClient,
		defaultDBusObjectManagerClientOverrides,
	)
}

func initDBusObjectManagerClientClass(gclass unsafe.Pointer, overrides DBusObjectManagerClientOverrides, classInitFunc func(*DBusObjectManagerClientClass)) {
	pclass := (*C.GDBusObjectManagerClientClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeDBusObjectManagerClient))))

	if overrides.InterfaceProxySignal != nil {
		pclass.interface_proxy_signal = (*[0]byte)(C._gotk4_gio2_DBusObjectManagerClientClass_interface_proxy_signal)
	}

	if classInitFunc != nil {
		class := (*DBusObjectManagerClientClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDBusObjectManagerClient(obj *coreglib.Object) *DBusObjectManagerClient {
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

func marshalDBusObjectManagerClient(p uintptr) (interface{}, error) {
	return wrapDBusObjectManagerClient(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectInterfaceProxyPropertiesChanged is emitted when one or more D-Bus
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
func (manager *DBusObjectManagerClient) ConnectInterfaceProxyPropertiesChanged(f func(objectProxy *DBusObjectProxy, interfaceProxy *DBusProxy, changedProperties *glib.Variant, invalidatedProperties []string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(manager, "interface-proxy-properties-changed", false, unsafe.Pointer(C._gotk4_gio2_DBusObjectManagerClient_ConnectInterfaceProxyPropertiesChanged), f)
}

// ConnectInterfaceProxySignal is emitted when a D-Bus signal is received on
// interface_proxy.
//
// This signal exists purely as a convenience to avoid having to connect signals
// to all interface proxies managed by manager.
//
// This signal is emitted in the [thread-default main
// context][g-main-context-push-thread-default] that manager was constructed in.
func (manager *DBusObjectManagerClient) ConnectInterfaceProxySignal(f func(objectProxy *DBusObjectProxy, interfaceProxy *DBusProxy, senderName, signalName string, parameters *glib.Variant)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(manager, "interface-proxy-signal", false, unsafe.Pointer(C._gotk4_gio2_DBusObjectManagerClient_ConnectInterfaceProxySignal), f)
}

// NewDBusObjectManagerClientFinish finishes an operation started with
// g_dbus_object_manager_client_new().
//
// The function takes the following parameters:
//
//    - res obtained from the ReadyCallback passed to
//      g_dbus_object_manager_client_new().
//
// The function returns the following values:
//
//    - dBusObjectManagerClient: a BusObjectManagerClient object or NULL if error
//      is set. Free with g_object_unref().
//
func NewDBusObjectManagerClientFinish(res AsyncResulter) (*DBusObjectManagerClient, error) {
	var _arg1 *C.GAsyncResult       // out
	var _cret *C.GDBusObjectManager // in
	var _cerr *C.GError             // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(res).Native()))

	_cret = C.g_dbus_object_manager_client_new_finish(_arg1, &_cerr)
	runtime.KeepAlive(res)

	var _dBusObjectManagerClient *DBusObjectManagerClient // out
	var _goerr error                                      // out

	_dBusObjectManagerClient = wrapDBusObjectManagerClient(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
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
//      g_dbus_object_manager_client_new_for_bus().
//
// The function returns the following values:
//
//    - dBusObjectManagerClient: a BusObjectManagerClient object or NULL if error
//      is set. Free with g_object_unref().
//
func NewDBusObjectManagerClientForBusFinish(res AsyncResulter) (*DBusObjectManagerClient, error) {
	var _arg1 *C.GAsyncResult       // out
	var _cret *C.GDBusObjectManager // in
	var _cerr *C.GError             // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(res).Native()))

	_cret = C.g_dbus_object_manager_client_new_for_bus_finish(_arg1, &_cerr)
	runtime.KeepAlive(res)

	var _dBusObjectManagerClient *DBusObjectManagerClient // out
	var _goerr error                                      // out

	_dBusObjectManagerClient = wrapDBusObjectManagerClient(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
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
//    - ctx (optional) or NULL.
//    - busType: Type.
//    - flags: zero or more flags from the BusObjectManagerClientFlags
//      enumeration.
//    - name: owner of the control object (unique or well-known name).
//    - objectPath: object path of the control object.
//    - getProxyTypeFunc (optional) function or NULL to always construct BusProxy
//      proxies.
//
// The function returns the following values:
//
//    - dBusObjectManagerClient: a BusObjectManagerClient object or NULL if error
//      is set. Free with g_object_unref().
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

	_dBusObjectManagerClient = wrapDBusObjectManagerClient(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
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
//    - ctx (optional) or NULL.
//    - connection: BusConnection.
//    - flags: zero or more flags from the BusObjectManagerClientFlags
//      enumeration.
//    - name (optional): owner of the control object (unique or well-known name),
//      or NULL when not using a message bus connection.
//    - objectPath: object path of the control object.
//    - getProxyTypeFunc (optional) function or NULL to always construct BusProxy
//      proxies.
//
// The function returns the following values:
//
//    - dBusObjectManagerClient: a BusObjectManagerClient object or NULL if error
//      is set. Free with g_object_unref().
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
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
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

	_dBusObjectManagerClient = wrapDBusObjectManagerClient(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dBusObjectManagerClient, _goerr
}

// Connection gets the BusConnection used by manager.
//
// The function returns the following values:
//
//    - dBusConnection object. Do not free, the object belongs to manager.
//
func (manager *DBusObjectManagerClient) Connection() *DBusConnection {
	var _arg0 *C.GDBusObjectManagerClient // out
	var _cret *C.GDBusConnection          // in

	_arg0 = (*C.GDBusObjectManagerClient)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	_cret = C.g_dbus_object_manager_client_get_connection(_arg0)
	runtime.KeepAlive(manager)

	var _dBusConnection *DBusConnection // out

	_dBusConnection = wrapDBusConnection(coreglib.Take(unsafe.Pointer(_cret)))

	return _dBusConnection
}

// Flags gets the flags that manager was constructed with.
//
// The function returns the following values:
//
//    - dBusObjectManagerClientFlags: zero of more flags from the
//      BusObjectManagerClientFlags enumeration.
//
func (manager *DBusObjectManagerClient) Flags() DBusObjectManagerClientFlags {
	var _arg0 *C.GDBusObjectManagerClient     // out
	var _cret C.GDBusObjectManagerClientFlags // in

	_arg0 = (*C.GDBusObjectManagerClient)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	_cret = C.g_dbus_object_manager_client_get_flags(_arg0)
	runtime.KeepAlive(manager)

	var _dBusObjectManagerClientFlags DBusObjectManagerClientFlags // out

	_dBusObjectManagerClientFlags = DBusObjectManagerClientFlags(_cret)

	return _dBusObjectManagerClientFlags
}

// Name gets the name that manager is for, or NULL if not a message bus
// connection.
//
// The function returns the following values:
//
//    - utf8: unique or well-known name. Do not free, the string belongs to
//      manager.
//
func (manager *DBusObjectManagerClient) Name() string {
	var _arg0 *C.GDBusObjectManagerClient // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.GDBusObjectManagerClient)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	_cret = C.g_dbus_object_manager_client_get_name(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// NameOwner: unique name that owns the name that manager is for or NULL if
// no-one currently owns that name. You can connect to the #GObject::notify
// signal to track changes to the BusObjectManagerClient:name-owner property.
//
// The function returns the following values:
//
//    - utf8 (optional): name owner or NULL if no name owner exists. Free with
//      g_free().
//
func (manager *DBusObjectManagerClient) NameOwner() string {
	var _arg0 *C.GDBusObjectManagerClient // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.GDBusObjectManagerClient)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	_cret = C.g_dbus_object_manager_client_get_name_owner(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// The function takes the following parameters:
//
//    - objectProxy
//    - interfaceProxy
//    - senderName
//    - signalName
//    - parameters
//
func (manager *DBusObjectManagerClient) interfaceProxySignal(objectProxy *DBusObjectProxy, interfaceProxy *DBusProxy, senderName, signalName string, parameters *glib.Variant) {
	gclass := (*C.GDBusObjectManagerClientClass)(coreglib.PeekParentClass(manager))
	fnarg := gclass.interface_proxy_signal

	var _arg0 *C.GDBusObjectManagerClient // out
	var _arg1 *C.GDBusObjectProxy         // out
	var _arg2 *C.GDBusProxy               // out
	var _arg3 *C.gchar                    // out
	var _arg4 *C.gchar                    // out
	var _arg5 *C.GVariant                 // out

	_arg0 = (*C.GDBusObjectManagerClient)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	_arg1 = (*C.GDBusObjectProxy)(unsafe.Pointer(coreglib.InternObject(objectProxy).Native()))
	_arg2 = (*C.GDBusProxy)(unsafe.Pointer(coreglib.InternObject(interfaceProxy).Native()))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(senderName)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(signalName)))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(parameters)))

	C._gotk4_gio2_DBusObjectManagerClient_virtual_interface_proxy_signal(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(objectProxy)
	runtime.KeepAlive(interfaceProxy)
	runtime.KeepAlive(senderName)
	runtime.KeepAlive(signalName)
	runtime.KeepAlive(parameters)
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
//    - ctx (optional) or NULL.
//    - connection: BusConnection.
//    - flags: zero or more flags from the BusObjectManagerClientFlags
//      enumeration.
//    - name: owner of the control object (unique or well-known name).
//    - objectPath: object path of the control object.
//    - getProxyTypeFunc (optional) function or NULL to always construct BusProxy
//      proxies.
//    - callback (optional) to call when the request is satisfied.
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
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
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
//    - ctx (optional) or NULL.
//    - busType: Type.
//    - flags: zero or more flags from the BusObjectManagerClientFlags
//      enumeration.
//    - name: owner of the control object (unique or well-known name).
//    - objectPath: object path of the control object.
//    - getProxyTypeFunc (optional) function or NULL to always construct BusProxy
//      proxies.
//    - callback (optional) to call when the request is satisfied.
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

// DBusObjectManagerClientClass class structure for BusObjectManagerClient.
//
// An instance of this type is always passed by reference.
type DBusObjectManagerClientClass struct {
	*dBusObjectManagerClientClass
}

// dBusObjectManagerClientClass is the struct that's finalized.
type dBusObjectManagerClientClass struct {
	native *C.GDBusObjectManagerClientClass
}
