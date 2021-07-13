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
		{T: externglib.Type(C.g_dbus_object_manager_server_get_type()), F: marshalDBusObjectManagerServerer},
	})
}

// DBusObjectManagerServerer describes DBusObjectManagerServer's methods.
type DBusObjectManagerServerer interface {
	// Export exports object on manager.
	Export(object *DBusObjectSkeleton)
	// ExportUniquely: like g_dbus_object_manager_server_export() but appends a
	// string of the form _N (with N being a natural number) to object's object
	// path if an object with the given path already exists.
	ExportUniquely(object *DBusObjectSkeleton)
	// Connection gets the BusConnection used by manager.
	Connection() *DBusConnection
	// IsExported returns whether object is currently exported on manager.
	IsExported(object *DBusObjectSkeleton) bool
	// SetConnection exports all objects managed by manager on connection.
	SetConnection(connection *DBusConnection)
	// Unexport: if manager has an object at path, removes the object.
	Unexport(objectPath string) bool
}

// DBusObjectManagerServer is used to export BusObject instances using the
// standardized org.freedesktop.DBus.ObjectManager
// (http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-objectmanager)
// interface. For example, remote D-Bus clients can get all objects and
// properties in a single call. Additionally, any change in the object hierarchy
// is broadcast using signals. This means that D-Bus clients can keep caches up
// to date by only listening to D-Bus signals.
//
// The recommended path to export an object manager at is the path form of the
// well-known name of a D-Bus service, or below. For example, if a D-Bus service
// is available at the well-known name net.example.ExampleService1, the object
// manager should typically be exported at /net/example/ExampleService1, or
// below (to allow for multiple object managers in a service).
//
// It is supported, but not recommended, to export an object manager at the root
// path, /.
//
// See BusObjectManagerClient for the client-side code that is intended to be
// used with BusObjectManagerServer or any D-Bus object implementing the
// org.freedesktop.DBus.ObjectManager interface.
type DBusObjectManagerServer struct {
	*externglib.Object

	DBusObjectManager
}

var (
	_ DBusObjectManagerServerer = (*DBusObjectManagerServer)(nil)
	_ gextras.Nativer           = (*DBusObjectManagerServer)(nil)
)

func wrapDBusObjectManagerServer(obj *externglib.Object) *DBusObjectManagerServer {
	return &DBusObjectManagerServer{
		Object: obj,
		DBusObjectManager: DBusObjectManager{
			Object: obj,
		},
	}
}

func marshalDBusObjectManagerServerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDBusObjectManagerServer(obj), nil
}

// NewDBusObjectManagerServer creates a new BusObjectManagerServer object.
//
// The returned server isn't yet exported on any connection. To do so, use
// g_dbus_object_manager_server_set_connection(). Normally you want to export
// all of your objects before doing so to avoid InterfacesAdded
// (http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-objectmanager)
// signals being emitted.
func NewDBusObjectManagerServer(objectPath string) *DBusObjectManagerServer {
	var _arg1 *C.gchar                    // out
	var _cret *C.GDBusObjectManagerServer // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))

	_cret = C.g_dbus_object_manager_server_new(_arg1)

	var _dBusObjectManagerServer *DBusObjectManagerServer // out

	_dBusObjectManagerServer = wrapDBusObjectManagerServer(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusObjectManagerServer
}

// Export exports object on manager.
//
// If there is already a BusObject exported at the object path, then the old
// object is removed.
//
// The object path for object must be in the hierarchy rooted by the object path
// for manager.
//
// Note that manager will take a reference on object for as long as it is
// exported.
func (manager *DBusObjectManagerServer) Export(object *DBusObjectSkeleton) {
	var _arg0 *C.GDBusObjectManagerServer // out
	var _arg1 *C.GDBusObjectSkeleton      // out

	_arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))

	C.g_dbus_object_manager_server_export(_arg0, _arg1)
}

// ExportUniquely: like g_dbus_object_manager_server_export() but appends a
// string of the form _N (with N being a natural number) to object's object path
// if an object with the given path already exists. As such, the
// BusObjectProxy:g-object-path property of object may be modified.
func (manager *DBusObjectManagerServer) ExportUniquely(object *DBusObjectSkeleton) {
	var _arg0 *C.GDBusObjectManagerServer // out
	var _arg1 *C.GDBusObjectSkeleton      // out

	_arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))

	C.g_dbus_object_manager_server_export_uniquely(_arg0, _arg1)
}

// Connection gets the BusConnection used by manager.
func (manager *DBusObjectManagerServer) Connection() *DBusConnection {
	var _arg0 *C.GDBusObjectManagerServer // out
	var _cret *C.GDBusConnection          // in

	_arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(manager.Native()))

	_cret = C.g_dbus_object_manager_server_get_connection(_arg0)

	var _dBusConnection *DBusConnection // out

	_dBusConnection = wrapDBusConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusConnection
}

// IsExported returns whether object is currently exported on manager.
func (manager *DBusObjectManagerServer) IsExported(object *DBusObjectSkeleton) bool {
	var _arg0 *C.GDBusObjectManagerServer // out
	var _arg1 *C.GDBusObjectSkeleton      // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))

	_cret = C.g_dbus_object_manager_server_is_exported(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetConnection exports all objects managed by manager on connection. If
// connection is NULL, stops exporting objects.
func (manager *DBusObjectManagerServer) SetConnection(connection *DBusConnection) {
	var _arg0 *C.GDBusObjectManagerServer // out
	var _arg1 *C.GDBusConnection          // out

	_arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))

	C.g_dbus_object_manager_server_set_connection(_arg0, _arg1)
}

// Unexport: if manager has an object at path, removes the object. Otherwise
// does nothing.
//
// Note that object_path must be in the hierarchy rooted by the object path for
// manager.
func (manager *DBusObjectManagerServer) Unexport(objectPath string) bool {
	var _arg0 *C.GDBusObjectManagerServer // out
	var _arg1 *C.gchar                    // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))

	_cret = C.g_dbus_object_manager_server_unexport(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
