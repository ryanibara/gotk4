// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
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

// BusNameAppearedCallback: invoked when the name being watched is known to have
// to have an owner.
type BusNameAppearedCallback func(connection *DBusConnection, name string, nameOwner string)

//export _gotk4_gio2_BusNameAppearedCallback
func _gotk4_gio2_BusNameAppearedCallback(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 C.gpointer) {
	v := gbox.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var connection *DBusConnection // out
	var name string                // out
	var nameOwner string           // out

	connection = wrapDBusConnection(externglib.Take(unsafe.Pointer(arg0)))
	name = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	defer C.free(unsafe.Pointer(arg1))
	nameOwner = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	defer C.free(unsafe.Pointer(arg2))

	fn := v.(BusNameAppearedCallback)
	fn(connection, name, nameOwner)
}

// BusNameVanishedCallback: invoked when the name being watched is known not to
// have to have an owner.
//
// This is also invoked when the BusConnection on which the watch was
// established has been closed. In that case, connection will be NULL.
type BusNameVanishedCallback func(connection *DBusConnection, name string)

//export _gotk4_gio2_BusNameVanishedCallback
func _gotk4_gio2_BusNameVanishedCallback(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var connection *DBusConnection // out
	var name string                // out

	connection = wrapDBusConnection(externglib.Take(unsafe.Pointer(arg0)))
	name = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	defer C.free(unsafe.Pointer(arg1))

	fn := v.(BusNameVanishedCallback)
	fn(connection, name)
}

// BusUnwatchName stops watching a name.
//
// Note that there may still be D-Bus traffic to process (relating to watching
// and unwatching the name) in the current thread-default Context after this
// function has returned. You should continue to iterate the Context until the
// Notify function passed to g_bus_watch_name() is called, in order to avoid
// memory leaks through callbacks queued on the Context after it’s stopped being
// iterated.
func BusUnwatchName(watcherId uint) {
	var _arg1 C.guint // out

	_arg1 = C.guint(watcherId)

	C.g_bus_unwatch_name(_arg1)
}
