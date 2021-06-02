// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0
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
//
// extern void gotk4_BusNameAppearedCallback(GDBusConnection* _0, const gchar* _1, const gchar* _2, gpointer _3);
// extern void gotk4_BusNameVanishedCallback(GDBusConnection* _0, const gchar* _1, gpointer _2);
import "C"

// BusNameAppearedCallback: invoked when the name being watched is known to have
// to have an owner.
type BusNameAppearedCallback func(connection DBusConnection, name string, nameOwner string)

//export gotk4_BusNameAppearedCallback
func gotk4_BusNameAppearedCallback(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 C.gpointer) {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var connection DBusConnection
	var name string
	var nameOwner string

	connection = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0.Native()))).(DBusConnection)

	name = C.GoString(arg1)

	nameOwner = C.GoString(arg2)

	v.(BusNameAppearedCallback)(connection, name, nameOwner)
}

// BusNameVanishedCallback: invoked when the name being watched is known not to
// have to have an owner.
//
// This is also invoked when the BusConnection on which the watch was
// established has been closed. In that case, @connection will be nil.
type BusNameVanishedCallback func(connection DBusConnection, name string)

//export gotk4_BusNameVanishedCallback
func gotk4_BusNameVanishedCallback(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var connection DBusConnection
	var name string

	connection = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0.Native()))).(DBusConnection)

	name = C.GoString(arg1)

	v.(BusNameVanishedCallback)(connection, name)
}

// BusUnwatchName stops watching a name.
//
// Note that there may still be D-Bus traffic to process (relating to watching
// and unwatching the name) in the current thread-default Context after this
// function has returned. You should continue to iterate the Context until the
// Notify function passed to g_bus_watch_name() is called, in order to avoid
// memory leaks through callbacks queued on the Context after it’s stopped being
// iterated.
func BusUnwatchName(watcherID uint) {
	var arg1 C.guint

	arg1 = C.guint(watcherID)

	C.g_bus_unwatch_name(arg1)
}
