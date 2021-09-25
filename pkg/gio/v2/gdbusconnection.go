// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// DBusInterfaceGetPropertyFunc: type of the get_property function in
// BusInterfaceVTable.
type DBusInterfaceGetPropertyFunc func(connection *DBusConnection, sender string, objectPath string, interfaceName string, propertyName string) (err error, variant *glib.Variant)

//export _gotk4_gio2_DBusInterfaceGetPropertyFunc
func _gotk4_gio2_DBusInterfaceGetPropertyFunc(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 *C.gchar, arg4 *C.gchar, arg5 **C.GError, arg6 C.gpointer) (cret *C.GVariant) {
	v := gbox.Get(uintptr(arg6))
	if v == nil {
		panic(`callback not found`)
	}

	var connection *DBusConnection // out
	var sender string              // out
	var objectPath string          // out
	var interfaceName string       // out
	var propertyName string        // out

	connection = wrapDBusConnection(externglib.Take(unsafe.Pointer(arg0)))
	sender = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	objectPath = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	interfaceName = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	propertyName = C.GoString((*C.gchar)(unsafe.Pointer(arg4)))

	fn := v.(DBusInterfaceGetPropertyFunc)
	err, variant := fn(connection, sender, objectPath, interfaceName, propertyName)

	*arg5 = (*C.GError)(gerror.New(err))
	cret = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(variant)))

	return cret
}

// DBusInterfaceMethodCallFunc: type of the method_call function in
// BusInterfaceVTable.
type DBusInterfaceMethodCallFunc func(connection *DBusConnection, sender string, objectPath string, interfaceName string, methodName string, parameters *glib.Variant, invocation *DBusMethodInvocation)

//export _gotk4_gio2_DBusInterfaceMethodCallFunc
func _gotk4_gio2_DBusInterfaceMethodCallFunc(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 *C.gchar, arg4 *C.gchar, arg5 *C.GVariant, arg6 *C.GDBusMethodInvocation, arg7 C.gpointer) {
	v := gbox.Get(uintptr(arg7))
	if v == nil {
		panic(`callback not found`)
	}

	var connection *DBusConnection       // out
	var sender string                    // out
	var objectPath string                // out
	var interfaceName string             // out
	var methodName string                // out
	var parameters *glib.Variant         // out
	var invocation *DBusMethodInvocation // out

	connection = wrapDBusConnection(externglib.Take(unsafe.Pointer(arg0)))
	sender = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	objectPath = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	interfaceName = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	methodName = C.GoString((*C.gchar)(unsafe.Pointer(arg4)))
	parameters = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg5)))
	C.g_variant_ref(arg5)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(parameters)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)
	invocation = wrapDBusMethodInvocation(externglib.AssumeOwnership(unsafe.Pointer(arg6)))

	fn := v.(DBusInterfaceMethodCallFunc)
	fn(connection, sender, objectPath, interfaceName, methodName, parameters, invocation)
}

// DBusInterfaceSetPropertyFunc: type of the set_property function in
// BusInterfaceVTable.
type DBusInterfaceSetPropertyFunc func(connection *DBusConnection, sender string, objectPath string, interfaceName string, propertyName string, value *glib.Variant) (err error, ok bool)

//export _gotk4_gio2_DBusInterfaceSetPropertyFunc
func _gotk4_gio2_DBusInterfaceSetPropertyFunc(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 *C.gchar, arg4 *C.gchar, arg5 *C.GVariant, arg6 **C.GError, arg7 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg7))
	if v == nil {
		panic(`callback not found`)
	}

	var connection *DBusConnection // out
	var sender string              // out
	var objectPath string          // out
	var interfaceName string       // out
	var propertyName string        // out
	var value *glib.Variant        // out

	connection = wrapDBusConnection(externglib.Take(unsafe.Pointer(arg0)))
	sender = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	objectPath = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	interfaceName = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	propertyName = C.GoString((*C.gchar)(unsafe.Pointer(arg4)))
	value = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg5)))
	C.g_variant_ref(arg5)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(value)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	fn := v.(DBusInterfaceSetPropertyFunc)
	err, ok := fn(connection, sender, objectPath, interfaceName, propertyName, value)

	*arg6 = (*C.GError)(gerror.New(err))
	if ok {
		cret = C.TRUE
	}

	return cret
}

// DBusMessageFilterFunction: signature for function used in
// g_dbus_connection_add_filter().
//
// A filter function is passed a BusMessage and expected to return a BusMessage
// too. Passive filter functions that don't modify the message can simply return
// the message object:
//
//    static GDBusMessage *
//    passive_filter (GDBusConnection *connection
//                    GDBusMessage    *message,
//                    gboolean         incoming,
//                    gpointer         user_data)
//    {
//      // inspect message
//      return message;
//    }
//
// Filter functions that wants to drop a message can simply return NULL:
//
//    static GDBusMessage *
//    drop_filter (GDBusConnection *connection
//                 GDBusMessage    *message,
//                 gboolean         incoming,
//                 gpointer         user_data)
//    {
//      if (should_drop_message)
//        {
//          g_object_unref (message);
//          message = NULL;
//        }
//      return message;
//    }
//
// Finally, a filter function may modify a message by copying it:
//
//    static GDBusMessage *
//    modifying_filter (GDBusConnection *connection
//                      GDBusMessage    *message,
//                      gboolean         incoming,
//                      gpointer         user_data)
//    {
//      GDBusMessage *copy;
//      GError *error;
//
//      error = NULL;
//      copy = g_dbus_message_copy (message, &error);
//      // handle error being set
//      g_object_unref (message);
//
//      // modify copy
//
//      return copy;
//    }
//
// If the returned BusMessage is different from message and cannot be sent on
// connection (it could use features, such as file descriptors, not compatible
// with connection), then a warning is logged to standard error. Applications
// can check this ahead of time using g_dbus_message_to_blob() passing a
// BusCapabilityFlags value obtained from connection.
type DBusMessageFilterFunction func(connection *DBusConnection, message *DBusMessage, incoming bool) (dBusMessage *DBusMessage)

//export _gotk4_gio2_DBusMessageFilterFunction
func _gotk4_gio2_DBusMessageFilterFunction(arg0 *C.GDBusConnection, arg1 *C.GDBusMessage, arg2 C.gboolean, arg3 C.gpointer) (cret *C.GDBusMessage) {
	v := gbox.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var connection *DBusConnection // out
	var message *DBusMessage       // out
	var incoming bool              // out

	connection = wrapDBusConnection(externglib.Take(unsafe.Pointer(arg0)))
	message = wrapDBusMessage(externglib.AssumeOwnership(unsafe.Pointer(arg1)))
	if arg2 != 0 {
		incoming = true
	}

	fn := v.(DBusMessageFilterFunction)
	dBusMessage := fn(connection, message, incoming)

	if dBusMessage != nil {
		cret = (*C.GDBusMessage)(unsafe.Pointer(dBusMessage.Native()))
		C.g_object_ref(C.gpointer(dBusMessage.Native()))
	}

	return cret
}

// DBusSignalCallback: signature for callback function used in
// g_dbus_connection_signal_subscribe().
type DBusSignalCallback func(connection *DBusConnection, senderName string, objectPath string, interfaceName string, signalName string, parameters *glib.Variant)

//export _gotk4_gio2_DBusSignalCallback
func _gotk4_gio2_DBusSignalCallback(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 *C.gchar, arg4 *C.gchar, arg5 *C.GVariant, arg6 C.gpointer) {
	v := gbox.Get(uintptr(arg6))
	if v == nil {
		panic(`callback not found`)
	}

	var connection *DBusConnection // out
	var senderName string          // out
	var objectPath string          // out
	var interfaceName string       // out
	var signalName string          // out
	var parameters *glib.Variant   // out

	connection = wrapDBusConnection(externglib.Take(unsafe.Pointer(arg0)))
	if arg1 != nil {
		senderName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	}
	objectPath = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	interfaceName = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	signalName = C.GoString((*C.gchar)(unsafe.Pointer(arg4)))
	parameters = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg5)))
	C.g_variant_ref(arg5)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(parameters)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	fn := v.(DBusSignalCallback)
	fn(connection, senderName, objectPath, interfaceName, signalName, parameters)
}

// DBusSubtreeDispatchFunc: type of the dispatch function in BusSubtreeVTable.
//
// Subtrees are flat. node, if non-NULL, is always exactly one segment of the
// object path (ie: it never contains a slash).
type DBusSubtreeDispatchFunc func(connection *DBusConnection, sender string, objectPath string, interfaceName string, node string) (outUserData cgo.Handle, dBusInterfaceVTable *DBusInterfaceVTable)

//export _gotk4_gio2_DBusSubtreeDispatchFunc
func _gotk4_gio2_DBusSubtreeDispatchFunc(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 *C.gchar, arg4 *C.gchar, arg5 *C.gpointer, arg6 C.gpointer) (cret *C.GDBusInterfaceVTable) {
	v := gbox.Get(uintptr(arg6))
	if v == nil {
		panic(`callback not found`)
	}

	var connection *DBusConnection // out
	var sender string              // out
	var objectPath string          // out
	var interfaceName string       // out
	var node string                // out

	connection = wrapDBusConnection(externglib.Take(unsafe.Pointer(arg0)))
	sender = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	objectPath = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	interfaceName = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	node = C.GoString((*C.gchar)(unsafe.Pointer(arg4)))

	fn := v.(DBusSubtreeDispatchFunc)
	outUserData, dBusInterfaceVTable := fn(connection, sender, objectPath, interfaceName, node)

	*arg5 = (C.gpointer)(unsafe.Pointer(outUserData))
	if dBusInterfaceVTable != nil {
		cret = (*C.GDBusInterfaceVTable)(gextras.StructNative(unsafe.Pointer(dBusInterfaceVTable)))
	}

	return cret
}

// DBusSubtreeEnumerateFunc: type of the enumerate function in BusSubtreeVTable.
//
// This function is called when generating introspection data and also when
// preparing to dispatch incoming messages in the event that the
// G_DBUS_SUBTREE_FLAGS_DISPATCH_TO_UNENUMERATED_NODES flag is not specified
// (ie: to verify that the object path is valid).
//
// Hierarchies are not supported; the items that you return should not contain
// the / character.
//
// The return value will be freed with g_strfreev().
type DBusSubtreeEnumerateFunc func(connection *DBusConnection, sender string, objectPath string) (utf8s []string)

//export _gotk4_gio2_DBusSubtreeEnumerateFunc
func _gotk4_gio2_DBusSubtreeEnumerateFunc(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 C.gpointer) (cret **C.gchar) {
	v := gbox.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var connection *DBusConnection // out
	var sender string              // out
	var objectPath string          // out

	connection = wrapDBusConnection(externglib.Take(unsafe.Pointer(arg0)))
	sender = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	objectPath = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	fn := v.(DBusSubtreeEnumerateFunc)
	utf8s := fn(connection, sender, objectPath)

	{
		cret = (**C.gchar)(C.malloc(C.ulong(len(utf8s)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		{
			out := unsafe.Slice(cret, len(utf8s)+1)
			var zero *C.gchar
			out[len(utf8s)] = zero
			for i := range utf8s {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(utf8s[i])))
			}
		}
	}

	return cret
}

// DBusSubtreeIntrospectFunc: type of the introspect function in
// BusSubtreeVTable.
//
// Subtrees are flat. node, if non-NULL, is always exactly one segment of the
// object path (ie: it never contains a slash).
//
// This function should return NULL to indicate that there is no object at this
// node.
//
// If this function returns non-NULL, the return value is expected to be a
// NULL-terminated array of pointers to BusInterfaceInfo structures describing
// the interfaces implemented by node. This array will have
// g_dbus_interface_info_unref() called on each item before being freed with
// g_free().
//
// The difference between returning NULL and an array containing zero items is
// that the standard DBus interfaces will returned to the remote introspector in
// the empty array case, but not in the NULL case.
type DBusSubtreeIntrospectFunc func(connection *DBusConnection, sender string, objectPath string, node string) (dBusInterfaceInfos []*DBusInterfaceInfo)

//export _gotk4_gio2_DBusSubtreeIntrospectFunc
func _gotk4_gio2_DBusSubtreeIntrospectFunc(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 *C.gchar, arg4 C.gpointer) (cret **C.GDBusInterfaceInfo) {
	v := gbox.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	var connection *DBusConnection // out
	var sender string              // out
	var objectPath string          // out
	var node string                // out

	connection = wrapDBusConnection(externglib.Take(unsafe.Pointer(arg0)))
	sender = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	objectPath = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	node = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))

	fn := v.(DBusSubtreeIntrospectFunc)
	dBusInterfaceInfos := fn(connection, sender, objectPath, node)

	if dBusInterfaceInfos != nil {
		{
			cret = (**C.GDBusInterfaceInfo)(C.malloc(C.ulong(len(dBusInterfaceInfos)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
			{
				out := unsafe.Slice(cret, len(dBusInterfaceInfos)+1)
				var zero *C.GDBusInterfaceInfo
				out[len(dBusInterfaceInfos)] = zero
				for i := range dBusInterfaceInfos {
					out[i] = (*C.GDBusInterfaceInfo)(gextras.StructNative(unsafe.Pointer(dBusInterfaceInfos[i])))
				}
			}
		}
	}

	return cret
}

// BusGet: asynchronously connects to the message bus specified by bus_type.
//
// When the operation is finished, callback will be invoked. You can then call
// g_bus_get_finish() to get the result of the operation.
//
// This is an asynchronous failable function. See g_bus_get_sync() for the
// synchronous version.
func BusGet(ctx context.Context, busType BusType, callback AsyncReadyCallback) {
	var _arg2 *C.GCancellable       // out
	var _arg1 C.GBusType            // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GBusType(busType)
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_bus_get(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(busType)
	runtime.KeepAlive(callback)
}

// BusGetFinish finishes an operation started with g_bus_get().
//
// The returned object is a singleton, that is, shared with other callers of
// g_bus_get() and g_bus_get_sync() for bus_type. In the event that you need a
// private message bus connection, use g_dbus_address_get_for_bus_sync() and
// g_dbus_connection_new_for_address().
//
// Note that the returned BusConnection object will (usually) have the
// BusConnection:exit-on-close property set to TRUE.
func BusGetFinish(res AsyncResulter) (*DBusConnection, error) {
	var _arg1 *C.GAsyncResult    // out
	var _cret *C.GDBusConnection // in
	var _cerr *C.GError          // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	_cret = C.g_bus_get_finish(_arg1, &_cerr)
	runtime.KeepAlive(res)

	var _dBusConnection *DBusConnection // out
	var _goerr error                    // out

	_dBusConnection = wrapDBusConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dBusConnection, _goerr
}

// BusGetSync: synchronously connects to the message bus specified by bus_type.
// Note that the returned object may shared with other callers, e.g. if two
// separate parts of a process calls this function with the same bus_type, they
// will share the same object.
//
// This is a synchronous failable function. See g_bus_get() and
// g_bus_get_finish() for the asynchronous version.
//
// The returned object is a singleton, that is, shared with other callers of
// g_bus_get() and g_bus_get_sync() for bus_type. In the event that you need a
// private message bus connection, use g_dbus_address_get_for_bus_sync() and
// g_dbus_connection_new_for_address().
//
// Note that the returned BusConnection object will (usually) have the
// BusConnection:exit-on-close property set to TRUE.
func BusGetSync(ctx context.Context, busType BusType) (*DBusConnection, error) {
	var _arg2 *C.GCancellable    // out
	var _arg1 C.GBusType         // out
	var _cret *C.GDBusConnection // in
	var _cerr *C.GError          // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GBusType(busType)

	_cret = C.g_bus_get_sync(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(busType)

	var _dBusConnection *DBusConnection // out
	var _goerr error                    // out

	_dBusConnection = wrapDBusConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dBusConnection, _goerr
}

// NewDBusConnection: asynchronously sets up a D-Bus connection for exchanging
// D-Bus messages with the end represented by stream.
//
// If stream is a Connection, then the corresponding #GSocket will be put into
// non-blocking mode.
//
// The D-Bus connection will interact with stream from a worker thread. As a
// result, the caller should not interact with stream after this method has been
// called, except by calling g_object_unref() on it.
//
// If observer is not NULL it may be used to control the authentication process.
//
// When the operation is finished, callback will be invoked. You can then call
// g_dbus_connection_new_finish() to get the result of the operation.
//
// This is an asynchronous failable constructor. See
// g_dbus_connection_new_sync() for the synchronous version.
func NewDBusConnection(ctx context.Context, stream IOStreamer, guid string, flags DBusConnectionFlags, observer *DBusAuthObserver, callback AsyncReadyCallback) {
	var _arg5 *C.GCancellable        // out
	var _arg1 *C.GIOStream           // out
	var _arg2 *C.gchar               // out
	var _arg3 C.GDBusConnectionFlags // out
	var _arg4 *C.GDBusAuthObserver   // out
	var _arg6 C.GAsyncReadyCallback  // out
	var _arg7 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GIOStream)(unsafe.Pointer(stream.Native()))
	if guid != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(guid)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	_arg3 = C.GDBusConnectionFlags(flags)
	if observer != nil {
		_arg4 = (*C.GDBusAuthObserver)(unsafe.Pointer(observer.Native()))
	}
	if callback != nil {
		_arg6 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg7 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_dbus_connection_new(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(guid)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(observer)
	runtime.KeepAlive(callback)
}

// NewDBusConnectionForAddress: asynchronously connects and sets up a D-Bus
// client connection for exchanging D-Bus messages with an endpoint specified by
// address which must be in the D-Bus address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// This constructor can only be used to initiate client-side connections - use
// g_dbus_connection_new() if you need to act as the server. In particular,
// flags cannot contain the G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER,
// G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS or
// G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_REQUIRE_SAME_USER flags.
//
// When the operation is finished, callback will be invoked. You can then call
// g_dbus_connection_new_for_address_finish() to get the result of the
// operation.
//
// If observer is not NULL it may be used to control the authentication process.
//
// This is an asynchronous failable constructor. See
// g_dbus_connection_new_for_address_sync() for the synchronous version.
func NewDBusConnectionForAddress(ctx context.Context, address string, flags DBusConnectionFlags, observer *DBusAuthObserver, callback AsyncReadyCallback) {
	var _arg4 *C.GCancellable        // out
	var _arg1 *C.gchar               // out
	var _arg2 C.GDBusConnectionFlags // out
	var _arg3 *C.GDBusAuthObserver   // out
	var _arg5 C.GAsyncReadyCallback  // out
	var _arg6 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(address)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GDBusConnectionFlags(flags)
	if observer != nil {
		_arg3 = (*C.GDBusAuthObserver)(unsafe.Pointer(observer.Native()))
	}
	if callback != nil {
		_arg5 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg6 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_dbus_connection_new_for_address(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(address)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(observer)
	runtime.KeepAlive(callback)
}

// DBusInterfaceVTable: virtual table for handling properties and method calls
// for a D-Bus interface.
//
// Since 2.38, if you want to handle getting/setting D-Bus properties
// asynchronously, give NULL as your get_property() or set_property() function.
// The D-Bus call will be directed to your method_call function, with the
// provided interface_name set to "org.freedesktop.DBus.Properties".
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
// The usual checks on the validity of the calls is performed. For Get calls, an
// error is automatically returned if the property does not exist or the
// permissions do not allow access. The same checks are performed for Set calls,
// and the provided value is also checked for being the correct type.
//
// For both Get and Set calls, the BusMethodInvocation passed to the method_call
// handler can be queried with g_dbus_method_invocation_get_property_info() to
// get a pointer to the BusPropertyInfo of the property.
//
// If you have readable properties specified in your interface info, you must
// ensure that you either provide a non-NULL get_property() function or provide
// implementations of both the Get and GetAll methods on
// org.freedesktop.DBus.Properties interface in your method_call function. Note
// that the required return type of the Get call is (v), not the type of the
// property. GetAll expects a return value of type a{sv}.
//
// If you have writable properties specified in your interface info, you must
// ensure that you either provide a non-NULL set_property() function or provide
// an implementation of the Set call. If implementing the call, you must return
// the value of type G_VARIANT_TYPE_UNIT.
//
// An instance of this type is always passed by reference.
type DBusInterfaceVTable struct {
	*dBusInterfaceVTable
}

// dBusInterfaceVTable is the struct that's finalized.
type dBusInterfaceVTable struct {
	native *C.GDBusInterfaceVTable
}

// DBusSubtreeVTable: virtual table for handling subtrees registered with
// g_dbus_connection_register_subtree().
//
// An instance of this type is always passed by reference.
type DBusSubtreeVTable struct {
	*dBusSubtreeVTable
}

// dBusSubtreeVTable is the struct that's finalized.
type dBusSubtreeVTable struct {
	native *C.GDBusSubtreeVTable
}
