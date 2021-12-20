// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_dbus_object_skeleton_get_type()), F: marshalDBusObjectSkeletonner},
	})
}

// DBusObjectSkeletonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DBusObjectSkeletonOverrider interface {
	// The function takes the following parameters:
	//
	//    - interface_
	//    - invocation
	//
	// The function returns the following values:
	//
	AuthorizeMethod(interface_ DBusInterfaceSkeletonner, invocation *DBusMethodInvocation) bool
}

// DBusObjectSkeleton instance is essentially a group of D-Bus interfaces. The
// set of exported interfaces on the object may be dynamic and change at
// runtime.
//
// This type is intended to be used with BusObjectManager.
type DBusObjectSkeleton struct {
	*externglib.Object

	DBusObject
}

var (
	_ externglib.Objector = (*DBusObjectSkeleton)(nil)
)

func wrapDBusObjectSkeleton(obj *externglib.Object) *DBusObjectSkeleton {
	return &DBusObjectSkeleton{
		Object: obj,
		DBusObject: DBusObject{
			Object: obj,
		},
	}
}

func marshalDBusObjectSkeletonner(p uintptr) (interface{}, error) {
	return wrapDBusObjectSkeleton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectAuthorizeMethod: emitted when a method is invoked by a remote caller
// and used to determine if the method call is authorized.
//
// This signal is like BusInterfaceSkeleton's
// BusInterfaceSkeleton::g-authorize-method signal, except that it is for the
// enclosing object.
//
// The default class handler just returns TRUE.
func (object *DBusObjectSkeleton) ConnectAuthorizeMethod(f func(iface DBusInterfaceSkeletonner, invocation DBusMethodInvocation) bool) externglib.SignalHandle {
	return object.Connect("authorize-method", f)
}

// NewDBusObjectSkeleton creates a new BusObjectSkeleton.
//
// The function takes the following parameters:
//
//    - objectPath: object path.
//
// The function returns the following values:
//
//    - dBusObjectSkeleton Free with g_object_unref().
//
func NewDBusObjectSkeleton(objectPath string) *DBusObjectSkeleton {
	var _arg1 *C.gchar               // out
	var _cret *C.GDBusObjectSkeleton // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_object_skeleton_new(_arg1)
	runtime.KeepAlive(objectPath)

	var _dBusObjectSkeleton *DBusObjectSkeleton // out

	_dBusObjectSkeleton = wrapDBusObjectSkeleton(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusObjectSkeleton
}

// AddInterface adds interface_ to object.
//
// If object already contains a BusInterfaceSkeleton with the same interface
// name, it is removed before interface_ is added.
//
// Note that object takes its own reference on interface_ and holds it until
// removed.
//
// The function takes the following parameters:
//
//    - interface_: BusInterfaceSkeleton.
//
func (object *DBusObjectSkeleton) AddInterface(interface_ DBusInterfaceSkeletonner) {
	var _arg0 *C.GDBusObjectSkeleton    // out
	var _arg1 *C.GDBusInterfaceSkeleton // out

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))
	_arg1 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	C.g_dbus_object_skeleton_add_interface(_arg0, _arg1)
	runtime.KeepAlive(object)
	runtime.KeepAlive(interface_)
}

// Flush: this method simply calls g_dbus_interface_skeleton_flush() on all
// interfaces belonging to object. See that method for when flushing is useful.
func (object *DBusObjectSkeleton) Flush() {
	var _arg0 *C.GDBusObjectSkeleton // out

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))

	C.g_dbus_object_skeleton_flush(_arg0)
	runtime.KeepAlive(object)
}

// RemoveInterface removes interface_ from object.
//
// The function takes the following parameters:
//
//    - interface_: BusInterfaceSkeleton.
//
func (object *DBusObjectSkeleton) RemoveInterface(interface_ DBusInterfaceSkeletonner) {
	var _arg0 *C.GDBusObjectSkeleton    // out
	var _arg1 *C.GDBusInterfaceSkeleton // out

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))
	_arg1 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	C.g_dbus_object_skeleton_remove_interface(_arg0, _arg1)
	runtime.KeepAlive(object)
	runtime.KeepAlive(interface_)
}

// RemoveInterfaceByName removes the BusInterface with interface_name from
// object.
//
// If no D-Bus interface of the given interface exists, this function does
// nothing.
//
// The function takes the following parameters:
//
//    - interfaceName d-Bus interface name.
//
func (object *DBusObjectSkeleton) RemoveInterfaceByName(interfaceName string) {
	var _arg0 *C.GDBusObjectSkeleton // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(interfaceName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_dbus_object_skeleton_remove_interface_by_name(_arg0, _arg1)
	runtime.KeepAlive(object)
	runtime.KeepAlive(interfaceName)
}

// SetObjectPath sets the object path for object.
//
// The function takes the following parameters:
//
//    - objectPath: valid D-Bus object path.
//
func (object *DBusObjectSkeleton) SetObjectPath(objectPath string) {
	var _arg0 *C.GDBusObjectSkeleton // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_dbus_object_skeleton_set_object_path(_arg0, _arg1)
	runtime.KeepAlive(object)
	runtime.KeepAlive(objectPath)
}
