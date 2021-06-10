// Code generated by girgen. DO NOT EDIT.

package gio

import (
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
		{T: externglib.Type(C.g_dbus_object_get_type()), F: marshalDBusObject},
	})
}

// DBusObjectOverrider contains methods that are overridable. This
// interface is a subset of the interface DBusObject.
type DBusObjectOverrider interface {
	// ObjectPath gets the object path for @object.
	ObjectPath() string

	InterfaceAdded(interface_ DBusInterface)

	InterfaceRemoved(interface_ DBusInterface)
}

// DBusObject: the BusObject type is the base type for D-Bus objects on both the
// service side (see BusObjectSkeleton) and the client side (see
// BusObjectProxy). It is essentially just a container of interfaces.
type DBusObject interface {
	gextras.Objector
	DBusObjectOverrider
}

// dBusObject implements the DBusObject interface.
type dBusObject struct {
	gextras.Objector
}

var _ DBusObject = (*dBusObject)(nil)

// WrapDBusObject wraps a GObject to a type that implements interface
// DBusObject. It is primarily used internally.
func WrapDBusObject(obj *externglib.Object) DBusObject {
	return DBusObject{
		Objector: obj,
	}
}

func marshalDBusObject(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDBusObject(obj), nil
}

// ObjectPath gets the object path for @object.
func (o dBusObject) ObjectPath() string {
	var _arg0 *C.GDBusObject

	_arg0 = (*C.GDBusObject)(unsafe.Pointer(o.Native()))

	var _cret *C.gchar

	_cret = C.g_dbus_object_get_object_path(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}
