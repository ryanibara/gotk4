// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// DBusActionGroupGet obtains a BusActionGroup for the action group which is
// exported at the given bus_name and object_path.
//
// The thread default main context is taken at the time of this call. All
// signals on the menu model (and any linked models) are reported with respect
// to this context. All calls on the returned menu model (and linked models)
// must also originate from this same context, with the thread default main
// context unchanged.
//
// This call is non-blocking. The returned action group may or may not already
// be filled in. The correct thing to do is connect the signals for the action
// group to monitor for changes and then to call g_action_group_list_actions()
// to get the initial list.
//
// The function takes the following parameters:
//
//    - connection: BusConnection.
//    - busName (optional) bus name which exports the action group or NULL if
//      connection is not a message bus connection.
//    - objectPath: object path at which the action group is exported.
//
// The function returns the following values:
//
//    - dBusActionGroup: BusActionGroup.
//
func DBusActionGroupGet(connection *DBusConnection, busName, objectPath string) *DBusActionGroup {
	var _arg1 *C.GDBusConnection  // out
	var _arg2 *C.gchar            // out
	var _arg3 *C.gchar            // out
	var _cret *C.GDBusActionGroup // in

	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(externglib.InternObject(connection).Native()))
	if busName != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(busName)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.g_dbus_action_group_get(_arg1, _arg2, _arg3)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(busName)
	runtime.KeepAlive(objectPath)

	var _dBusActionGroup *DBusActionGroup // out

	_dBusActionGroup = wrapDBusActionGroup(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusActionGroup
}
