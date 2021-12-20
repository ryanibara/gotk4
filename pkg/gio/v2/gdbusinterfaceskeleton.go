// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_dbus_interface_skeleton_get_type()), F: marshalDBusInterfaceSkeletonner},
	})
}

// DBusInterfaceSkeletonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DBusInterfaceSkeletonOverrider interface {
	// Flush: if interface_ has outstanding changes, request for these changes
	// to be emitted immediately.
	//
	// For example, an exported D-Bus interface may queue up property changes
	// and emit the org.freedesktop.DBus.Properties.PropertiesChanged signal
	// later (e.g. in an idle handler). This technique is useful for collapsing
	// multiple property changes into one.
	Flush()
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	GAuthorizeMethod(invocation *DBusMethodInvocation) bool
	// Info gets D-Bus introspection information for the D-Bus interface
	// implemented by interface_.
	//
	// The function returns the following values:
	//
	//    - dBusInterfaceInfo (never NULL). Do not free.
	//
	Info() *DBusInterfaceInfo
	// Properties gets all D-Bus properties for interface_.
	//
	// The function returns the following values:
	//
	//    - variant of type ['a{sv}'][G-VARIANT-TYPE-VARDICT:CAPS]. Free with
	//      g_variant_unref().
	//
	Properties() *glib.Variant
}

// DBusInterfaceSkeleton: abstract base class for D-Bus interfaces on the
// service side.
type DBusInterfaceSkeleton struct {
	_ [0]func() // equal guard
	*externglib.Object

	DBusInterface
}

var (
	_ externglib.Objector = (*DBusInterfaceSkeleton)(nil)
)

// DBusInterfaceSkeletonner describes types inherited from class DBusInterfaceSkeleton.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type DBusInterfaceSkeletonner interface {
	externglib.Objector
	baseDBusInterfaceSkeleton() *DBusInterfaceSkeleton
}

var _ DBusInterfaceSkeletonner = (*DBusInterfaceSkeleton)(nil)

func wrapDBusInterfaceSkeleton(obj *externglib.Object) *DBusInterfaceSkeleton {
	return &DBusInterfaceSkeleton{
		Object: obj,
		DBusInterface: DBusInterface{
			Object: obj,
		},
	}
}

func marshalDBusInterfaceSkeletonner(p uintptr) (interface{}, error) {
	return wrapDBusInterfaceSkeleton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (interface_ *DBusInterfaceSkeleton) baseDBusInterfaceSkeleton() *DBusInterfaceSkeleton {
	return interface_
}

// BaseDBusInterfaceSkeleton returns the underlying base object.
func BaseDBusInterfaceSkeleton(obj DBusInterfaceSkeletonner) *DBusInterfaceSkeleton {
	return obj.baseDBusInterfaceSkeleton()
}

// ConnectGAuthorizeMethod: emitted when a method is invoked by a remote caller
// and used to determine if the method call is authorized.
//
// Note that this signal is emitted in a thread dedicated to handling the method
// call so handlers are allowed to perform blocking IO. This means that it is
// appropriate to call e.g. polkit_authority_check_authorization_sync()
// (http://hal.freedesktop.org/docs/polkit/PolkitAuthority.html#polkit-authority-check-authorization-sync)
// with the POLKIT_CHECK_AUTHORIZATION_FLAGS_ALLOW_USER_INTERACTION
// (http://hal.freedesktop.org/docs/polkit/PolkitAuthority.htmlLKIT-CHECK-AUTHORIZATION-FLAGS-ALLOW-USER-INTERACTION:CAPS)
// flag set.
//
// If FALSE is returned then no further handlers are run and the signal handler
// must take a reference to invocation and finish handling the call (e.g. return
// an error via g_dbus_method_invocation_return_error()).
//
// Otherwise, if TRUE is returned, signal emission continues. If no handlers
// return FALSE, then the method is dispatched. If interface has an enclosing
// BusObjectSkeleton, then the BusObjectSkeleton::authorize-method signal
// handlers run before the handlers for this signal.
//
// The default class handler just returns TRUE.
//
// Please note that the common case is optimized: if no signals handlers are
// connected and the default class handler isn't overridden (for both interface
// and the enclosing BusObjectSkeleton, if any) and BusInterfaceSkeleton:g-flags
// does not have the
// G_DBUS_INTERFACE_SKELETON_FLAGS_HANDLE_METHOD_INVOCATIONS_IN_THREAD flags
// set, no dedicated thread is ever used and the call will be handled in the
// same thread as the object that interface belongs to was exported in.
func (interface_ *DBusInterfaceSkeleton) ConnectGAuthorizeMethod(f func(invocation DBusMethodInvocation) bool) externglib.SignalHandle {
	return interface_.Connect("g-authorize-method", f)
}

// Export exports interface_ at object_path on connection.
//
// This can be called multiple times to export the same interface_ onto multiple
// connections however the object_path provided must be the same for all
// connections.
//
// Use g_dbus_interface_skeleton_unexport() to unexport the object.
//
// The function takes the following parameters:
//
//    - connection to export interface_ on.
//    - objectPath: path to export the interface at.
//
func (interface_ *DBusInterfaceSkeleton) Export(connection *DBusConnection, objectPath string) error {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _arg1 *C.GDBusConnection        // out
	var _arg2 *C.gchar                  // out
	var _cerr *C.GError                 // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_dbus_interface_skeleton_export(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(interface_)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(objectPath)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Flush: if interface_ has outstanding changes, request for these changes to be
// emitted immediately.
//
// For example, an exported D-Bus interface may queue up property changes and
// emit the org.freedesktop.DBus.Properties.PropertiesChanged signal later (e.g.
// in an idle handler). This technique is useful for collapsing multiple
// property changes into one.
func (interface_ *DBusInterfaceSkeleton) Flush() {
	var _arg0 *C.GDBusInterfaceSkeleton // out

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	C.g_dbus_interface_skeleton_flush(_arg0)
	runtime.KeepAlive(interface_)
}

// Connection gets the first connection that interface_ is exported on, if any.
//
// The function returns the following values:
//
//    - dBusConnection (optional) or NULL if interface_ is not exported anywhere.
//      Do not free, the object belongs to interface_.
//
func (interface_ *DBusInterfaceSkeleton) Connection() *DBusConnection {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _cret *C.GDBusConnection        // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	_cret = C.g_dbus_interface_skeleton_get_connection(_arg0)
	runtime.KeepAlive(interface_)

	var _dBusConnection *DBusConnection // out

	if _cret != nil {
		_dBusConnection = wrapDBusConnection(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _dBusConnection
}

// Connections gets a list of the connections that interface_ is exported on.
//
// The function returns the following values:
//
//    - list of all the connections that interface_ is exported on. The returned
//      list should be freed with g_list_free() after each element has been freed
//      with g_object_unref().
//
func (interface_ *DBusInterfaceSkeleton) Connections() []DBusConnection {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _cret *C.GList                  // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	_cret = C.g_dbus_interface_skeleton_get_connections(_arg0)
	runtime.KeepAlive(interface_)

	var _list []DBusConnection // out

	_list = make([]DBusConnection, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GDBusConnection)(v)
		var dst DBusConnection // out
		dst = *wrapDBusConnection(externglib.AssumeOwnership(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// Flags gets the BusInterfaceSkeletonFlags that describes what the behavior of
// interface_.
//
// The function returns the following values:
//
//    - dBusInterfaceSkeletonFlags: one or more flags from the
//      BusInterfaceSkeletonFlags enumeration.
//
func (interface_ *DBusInterfaceSkeleton) Flags() DBusInterfaceSkeletonFlags {
	var _arg0 *C.GDBusInterfaceSkeleton     // out
	var _cret C.GDBusInterfaceSkeletonFlags // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	_cret = C.g_dbus_interface_skeleton_get_flags(_arg0)
	runtime.KeepAlive(interface_)

	var _dBusInterfaceSkeletonFlags DBusInterfaceSkeletonFlags // out

	_dBusInterfaceSkeletonFlags = DBusInterfaceSkeletonFlags(_cret)

	return _dBusInterfaceSkeletonFlags
}

// Info gets D-Bus introspection information for the D-Bus interface implemented
// by interface_.
//
// The function returns the following values:
//
//    - dBusInterfaceInfo (never NULL). Do not free.
//
func (interface_ *DBusInterfaceSkeleton) Info() *DBusInterfaceInfo {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _cret *C.GDBusInterfaceInfo     // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	_cret = C.g_dbus_interface_skeleton_get_info(_arg0)
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

// ObjectPath gets the object path that interface_ is exported on, if any.
//
// The function returns the following values:
//
//    - utf8 (optional): string owned by interface_ or NULL if interface_ is not
//      exported anywhere. Do not free, the string belongs to interface_.
//
func (interface_ *DBusInterfaceSkeleton) ObjectPath() string {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _cret *C.gchar                  // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	_cret = C.g_dbus_interface_skeleton_get_object_path(_arg0)
	runtime.KeepAlive(interface_)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Properties gets all D-Bus properties for interface_.
//
// The function returns the following values:
//
//    - variant of type ['a{sv}'][G-VARIANT-TYPE-VARDICT:CAPS]. Free with
//      g_variant_unref().
//
func (interface_ *DBusInterfaceSkeleton) Properties() *glib.Variant {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _cret *C.GVariant               // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	_cret = C.g_dbus_interface_skeleton_get_properties(_arg0)
	runtime.KeepAlive(interface_)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	return _variant
}

// HasConnection checks if interface_ is exported on connection.
//
// The function takes the following parameters:
//
//    - connection: BusConnection.
//
// The function returns the following values:
//
//    - ok: TRUE if interface_ is exported on connection, FALSE otherwise.
//
func (interface_ *DBusInterfaceSkeleton) HasConnection(connection *DBusConnection) bool {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _arg1 *C.GDBusConnection        // out
	var _cret C.gboolean                // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))

	_cret = C.g_dbus_interface_skeleton_has_connection(_arg0, _arg1)
	runtime.KeepAlive(interface_)
	runtime.KeepAlive(connection)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetFlags sets flags describing what the behavior of skeleton should be.
//
// The function takes the following parameters:
//
//    - flags flags from the BusInterfaceSkeletonFlags enumeration.
//
func (interface_ *DBusInterfaceSkeleton) SetFlags(flags DBusInterfaceSkeletonFlags) {
	var _arg0 *C.GDBusInterfaceSkeleton     // out
	var _arg1 C.GDBusInterfaceSkeletonFlags // out

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))
	_arg1 = C.GDBusInterfaceSkeletonFlags(flags)

	C.g_dbus_interface_skeleton_set_flags(_arg0, _arg1)
	runtime.KeepAlive(interface_)
	runtime.KeepAlive(flags)
}

// Unexport stops exporting interface_ on all connections it is exported on.
//
// To unexport interface_ from only a single connection, use
// g_dbus_interface_skeleton_unexport_from_connection().
func (interface_ *DBusInterfaceSkeleton) Unexport() {
	var _arg0 *C.GDBusInterfaceSkeleton // out

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	C.g_dbus_interface_skeleton_unexport(_arg0)
	runtime.KeepAlive(interface_)
}

// UnexportFromConnection stops exporting interface_ on connection.
//
// To stop exporting on all connections the interface is exported on, use
// g_dbus_interface_skeleton_unexport().
//
// The function takes the following parameters:
//
//    - connection: BusConnection.
//
func (interface_ *DBusInterfaceSkeleton) UnexportFromConnection(connection *DBusConnection) {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _arg1 *C.GDBusConnection        // out

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))

	C.g_dbus_interface_skeleton_unexport_from_connection(_arg0, _arg1)
	runtime.KeepAlive(interface_)
	runtime.KeepAlive(connection)
}
