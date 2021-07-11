// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
		{T: externglib.Type(C.g_dbus_interface_skeleton_get_type()), F: marshalDBusInterfaceSkeletonner},
	})
}

// DBusInterfaceSkeletonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DBusInterfaceSkeletonOverrider interface {
	// Flush: if @interface_ has outstanding changes, request for these changes
	// to be emitted immediately.
	//
	// For example, an exported D-Bus interface may queue up property changes
	// and emit the `org.freedesktop.DBus.Properties.PropertiesChanged` signal
	// later (e.g. in an idle handler). This technique is useful for collapsing
	// multiple property changes into one.
	Flush()

	GAuthorizeMethod(invocation DBusMethodInvocationer) bool
	// Info gets D-Bus introspection information for the D-Bus interface
	// implemented by @interface_.
	Info() *DBusInterfaceInfo
	// Properties gets all D-Bus properties for @interface_.
	Properties() *glib.Variant
}

// DBusInterfaceSkeletonner describes DBusInterfaceSkeleton's methods.
type DBusInterfaceSkeletonner interface {
	// Export exports @interface_ at @object_path on @connection.
	Export(connection DBusConnectioner, objectPath string) error
	// Flush: if @interface_ has outstanding changes, request for these changes
	// to be emitted immediately.
	Flush()
	// Connection gets the first connection that @interface_ is exported on, if
	// any.
	Connection() *DBusConnection
	// Flags gets the BusInterfaceSkeletonFlags that describes what the behavior
	// of @interface_
	Flags() DBusInterfaceSkeletonFlags
	// Info gets D-Bus introspection information for the D-Bus interface
	// implemented by @interface_.
	Info() *DBusInterfaceInfo
	// ObjectPath gets the object path that @interface_ is exported on, if any.
	ObjectPath() string
	// Properties gets all D-Bus properties for @interface_.
	Properties() *glib.Variant
	// HasConnection checks if @interface_ is exported on @connection.
	HasConnection(connection DBusConnectioner) bool
	// Unexport stops exporting @interface_ on all connections it is exported
	// on.
	Unexport()
	// UnexportFromConnection stops exporting @interface_ on @connection.
	UnexportFromConnection(connection DBusConnectioner)
}

// DBusInterfaceSkeleton: abstract base class for D-Bus interfaces on the
// service side.
type DBusInterfaceSkeleton struct {
	*externglib.Object

	DBusInterface
}

var (
	_ DBusInterfaceSkeletonner = (*DBusInterfaceSkeleton)(nil)
	_ gextras.Nativer          = (*DBusInterfaceSkeleton)(nil)
)

func wrapDBusInterfaceSkeleton(obj *externglib.Object) DBusInterfaceSkeletonner {
	return &DBusInterfaceSkeleton{
		Object: obj,
		DBusInterface: DBusInterface{
			Object: obj,
		},
	}
}

func marshalDBusInterfaceSkeletonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDBusInterfaceSkeleton(obj), nil
}

// Export exports @interface_ at @object_path on @connection.
//
// This can be called multiple times to export the same @interface_ onto
// multiple connections however the @object_path provided must be the same for
// all connections.
//
// Use g_dbus_interface_skeleton_unexport() to unexport the object.
func (interface_ *DBusInterfaceSkeleton) Export(connection DBusConnectioner, objectPath string) error {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _arg1 *C.GDBusConnection        // out
	var _arg2 *C.gchar                  // out
	var _cerr *C.GError                 // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer((connection).(gextras.Nativer).Native()))
	_arg2 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_dbus_interface_skeleton_export(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// Flush: if @interface_ has outstanding changes, request for these changes to
// be emitted immediately.
//
// For example, an exported D-Bus interface may queue up property changes and
// emit the `org.freedesktop.DBus.Properties.PropertiesChanged` signal later
// (e.g. in an idle handler). This technique is useful for collapsing multiple
// property changes into one.
func (interface_ *DBusInterfaceSkeleton) Flush() {
	var _arg0 *C.GDBusInterfaceSkeleton // out

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	C.g_dbus_interface_skeleton_flush(_arg0)
}

// Connection gets the first connection that @interface_ is exported on, if any.
func (interface_ *DBusInterfaceSkeleton) Connection() *DBusConnection {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _cret *C.GDBusConnection        // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	_cret = C.g_dbus_interface_skeleton_get_connection(_arg0)

	var _dBusConnection *DBusConnection // out

	_dBusConnection = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*DBusConnection)

	return _dBusConnection
}

// Flags gets the BusInterfaceSkeletonFlags that describes what the behavior of
// @interface_
func (interface_ *DBusInterfaceSkeleton) Flags() DBusInterfaceSkeletonFlags {
	var _arg0 *C.GDBusInterfaceSkeleton     // out
	var _cret C.GDBusInterfaceSkeletonFlags // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	_cret = C.g_dbus_interface_skeleton_get_flags(_arg0)

	var _dBusInterfaceSkeletonFlags DBusInterfaceSkeletonFlags // out

	_dBusInterfaceSkeletonFlags = (DBusInterfaceSkeletonFlags)(_cret)

	return _dBusInterfaceSkeletonFlags
}

// Info gets D-Bus introspection information for the D-Bus interface implemented
// by @interface_.
func (interface_ *DBusInterfaceSkeleton) Info() *DBusInterfaceInfo {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _cret *C.GDBusInterfaceInfo     // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	_cret = C.g_dbus_interface_skeleton_get_info(_arg0)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	_dBusInterfaceInfo = (*DBusInterfaceInfo)(unsafe.Pointer(_cret))
	C.g_dbus_interface_info_ref(_cret)
	runtime.SetFinalizer(_dBusInterfaceInfo, func(v *DBusInterfaceInfo) {
		C.g_dbus_interface_info_unref((*C.GDBusInterfaceInfo)(unsafe.Pointer(v)))
	})

	return _dBusInterfaceInfo
}

// ObjectPath gets the object path that @interface_ is exported on, if any.
func (interface_ *DBusInterfaceSkeleton) ObjectPath() string {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _cret *C.gchar                  // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	_cret = C.g_dbus_interface_skeleton_get_object_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Properties gets all D-Bus properties for @interface_.
func (interface_ *DBusInterfaceSkeleton) Properties() *glib.Variant {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _cret *C.GVariant               // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	_cret = C.g_dbus_interface_skeleton_get_properties(_arg0)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(unsafe.Pointer(v)))
	})

	return _variant
}

// HasConnection checks if @interface_ is exported on @connection.
func (interface_ *DBusInterfaceSkeleton) HasConnection(connection DBusConnectioner) bool {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _arg1 *C.GDBusConnection        // out
	var _cret C.gboolean                // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer((connection).(gextras.Nativer).Native()))

	_cret = C.g_dbus_interface_skeleton_has_connection(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Unexport stops exporting @interface_ on all connections it is exported on.
//
// To unexport @interface_ from only a single connection, use
// g_dbus_interface_skeleton_unexport_from_connection()
func (interface_ *DBusInterfaceSkeleton) Unexport() {
	var _arg0 *C.GDBusInterfaceSkeleton // out

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	C.g_dbus_interface_skeleton_unexport(_arg0)
}

// UnexportFromConnection stops exporting @interface_ on @connection.
//
// To stop exporting on all connections the interface is exported on, use
// g_dbus_interface_skeleton_unexport().
func (interface_ *DBusInterfaceSkeleton) UnexportFromConnection(connection DBusConnectioner) {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _arg1 *C.GDBusConnection        // out

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer((connection).(gextras.Nativer).Native()))

	C.g_dbus_interface_skeleton_unexport_from_connection(_arg0, _arg1)
}
