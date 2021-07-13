// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
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

// BusAcquiredCallback: invoked when a connection to a message bus has been
// obtained.
type BusAcquiredCallback func(connection *DBusConnection, name string, userData cgo.Handle)

//export gotk4_BusAcquiredCallback
func gotk4_BusAcquiredCallback(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var connection *DBusConnection // out
	var name string                // out
	var userData cgo.Handle        // out

	connection = wrapDBusConnection(externglib.Take(unsafe.Pointer(arg0)))
	name = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	defer C.free(unsafe.Pointer(arg1))
	userData = (cgo.Handle)(unsafe.Pointer(arg2))

	fn := v.(BusAcquiredCallback)
	fn(connection, name, userData)
}

// BusNameAcquiredCallback: invoked when the name is acquired.
type BusNameAcquiredCallback func(connection *DBusConnection, name string, userData cgo.Handle)

//export gotk4_BusNameAcquiredCallback
func gotk4_BusNameAcquiredCallback(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var connection *DBusConnection // out
	var name string                // out
	var userData cgo.Handle        // out

	connection = wrapDBusConnection(externglib.Take(unsafe.Pointer(arg0)))
	name = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	defer C.free(unsafe.Pointer(arg1))
	userData = (cgo.Handle)(unsafe.Pointer(arg2))

	fn := v.(BusNameAcquiredCallback)
	fn(connection, name, userData)
}

// BusNameLostCallback: invoked when the name is lost or connection has been
// closed.
type BusNameLostCallback func(connection *DBusConnection, name string, userData cgo.Handle)

//export gotk4_BusNameLostCallback
func gotk4_BusNameLostCallback(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var connection *DBusConnection // out
	var name string                // out
	var userData cgo.Handle        // out

	connection = wrapDBusConnection(externglib.Take(unsafe.Pointer(arg0)))
	name = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	defer C.free(unsafe.Pointer(arg1))
	userData = (cgo.Handle)(unsafe.Pointer(arg2))

	fn := v.(BusNameLostCallback)
	fn(connection, name, userData)
}

// BusUnownName stops owning a name.
//
// Note that there may still be D-Bus traffic to process (relating to owning and
// unowning the name) in the current thread-default Context after this function
// has returned. You should continue to iterate the Context until the Notify
// function passed to g_bus_own_name() is called, in order to avoid memory leaks
// through callbacks queued on the Context after it’s stopped being iterated.
func BusUnownName(ownerId uint) {
	var _arg1 C.guint // out

	_arg1 = C.guint(ownerId)

	C.g_bus_unown_name(_arg1)
}
