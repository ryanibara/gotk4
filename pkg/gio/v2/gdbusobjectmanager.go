// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
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

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_dbus_object_manager_get_type()), F: marshalDBusObjectManager},
	})
}

// DBusObjectManagerOverrider contains methods that are overridable. This
// interface is a subset of the interface DBusObjectManager.
type DBusObjectManagerOverrider interface {
	// Interface gets the interface proxy for @interface_name at @object_path,
	// if any.
	Interface(objectPath string, interfaceName string) DBusInterface
	// Object gets the BusObjectProxy at @object_path, if any.
	Object(objectPath string) DBusObject
	// ObjectPath gets the object path that @manager is for.
	ObjectPath() string
	// Objects gets all BusObject objects known to @manager.
	Objects() *glib.List

	InterfaceAdded(object DBusObject, interface_ DBusInterface)

	InterfaceRemoved(object DBusObject, interface_ DBusInterface)

	ObjectAdded(object DBusObject)

	ObjectRemoved(object DBusObject)
}

// DBusObjectManager: the BusObjectManager type is the base type for service-
// and client-side implementations of the standardized
// org.freedesktop.DBus.ObjectManager
// (http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-objectmanager)
// interface.
//
// See BusObjectManagerClient for the client-side implementation and
// BusObjectManagerServer for the service-side implementation.
type DBusObjectManager interface {
	gextras.Objector
	DBusObjectManagerOverrider
}

// dBusObjectManager implements the DBusObjectManager interface.
type dBusObjectManager struct {
	gextras.Objector
}

var _ DBusObjectManager = (*dBusObjectManager)(nil)

// WrapDBusObjectManager wraps a GObject to a type that implements interface
// DBusObjectManager. It is primarily used internally.
func WrapDBusObjectManager(obj *externglib.Object) DBusObjectManager {
	return DBusObjectManager{
		Objector: obj,
	}
}

func marshalDBusObjectManager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDBusObjectManager(obj), nil
}

// Interface gets the interface proxy for @interface_name at @object_path,
// if any.
func (m dBusObjectManager) Interface(objectPath string, interfaceName string) DBusInterface {
	var arg0 *C.GDBusObjectManager
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg0 = (*C.GDBusObjectManager)(unsafe.Pointer(m.Native()))
	arg1 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(interfaceName))
	defer C.free(unsafe.Pointer(arg2))

	var cret *C.GDBusInterface
	var ret1 DBusInterface

	cret = C.g_dbus_object_manager_get_interface(arg0, objectPath, interfaceName)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DBusInterface)

	return ret1
}

// Object gets the BusObjectProxy at @object_path, if any.
func (m dBusObjectManager) Object(objectPath string) DBusObject {
	var arg0 *C.GDBusObjectManager
	var arg1 *C.gchar

	arg0 = (*C.GDBusObjectManager)(unsafe.Pointer(m.Native()))
	arg1 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GDBusObject
	var ret1 DBusObject

	cret = C.g_dbus_object_manager_get_object(arg0, objectPath)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DBusObject)

	return ret1
}

// ObjectPath gets the object path that @manager is for.
func (m dBusObjectManager) ObjectPath() string {
	var arg0 *C.GDBusObjectManager

	arg0 = (*C.GDBusObjectManager)(unsafe.Pointer(m.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.g_dbus_object_manager_get_object_path(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Objects gets all BusObject objects known to @manager.
func (m dBusObjectManager) Objects() *glib.List {
	var arg0 *C.GDBusObjectManager

	arg0 = (*C.GDBusObjectManager)(unsafe.Pointer(m.Native()))

	var cret *C.GList
	var ret1 *glib.List

	cret = C.g_dbus_object_manager_get_objects(arg0)

	ret1 = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}
