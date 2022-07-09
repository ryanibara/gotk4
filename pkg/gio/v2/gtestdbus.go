// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// AddServiceDir: add a path where dbus-daemon will look up .service files. This
// can't be called after g_test_dbus_up().
//
// The function takes the following parameters:
//
//    - path to a directory containing .service files.
//
func (self *TestDBus) AddServiceDir(path string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gio", "TestDBus").InvokeMethod("add_service_dir", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(path)
}

// Down: stop the session bus started by g_test_dbus_up().
//
// This will wait for the singleton returned by g_bus_get() or g_bus_get_sync()
// to be destroyed. This is done to ensure that the next unit test won't get a
// leaked singleton from this test.
func (self *TestDBus) Down() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	girepository.MustFind("Gio", "TestDBus").InvokeMethod("down", _args[:], nil)

	runtime.KeepAlive(self)
}

// BusAddress: get the address on which dbus-daemon is running. If
// g_test_dbus_up() has not been called yet, NULL is returned. This can be used
// with g_dbus_connection_new_for_address().
//
// The function returns the following values:
//
//    - utf8 (optional) address of the bus, or NULL.
//
func (self *TestDBus) BusAddress() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gio", "TestDBus").InvokeMethod("get_bus_address", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Stop the session bus started by g_test_dbus_up().
//
// Unlike g_test_dbus_down(), this won't verify the BusConnection singleton
// returned by g_bus_get() or g_bus_get_sync() is destroyed. Unit tests wanting
// to verify behaviour after the session bus has been stopped can use this
// function but should still call g_test_dbus_down() when done.
func (self *TestDBus) Stop() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	girepository.MustFind("Gio", "TestDBus").InvokeMethod("stop", _args[:], nil)

	runtime.KeepAlive(self)
}

// Up: start a dbus-daemon instance and set DBUS_SESSION_BUS_ADDRESS. After this
// call, it is safe for unit tests to start sending messages on the session bus.
//
// If this function is called from setup callback of g_test_add(),
// g_test_dbus_down() must be called in its teardown callback.
//
// If this function is called from unit test's main(), then g_test_dbus_down()
// must be called after g_test_run().
func (self *TestDBus) Up() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	girepository.MustFind("Gio", "TestDBus").InvokeMethod("up", _args[:], nil)

	runtime.KeepAlive(self)
}

// TestDBusUnset: unset DISPLAY and DBUS_SESSION_BUS_ADDRESS env variables to
// ensure the test won't use user's session bus.
//
// This is useful for unit tests that want to verify behaviour when no session
// bus is running. It is not necessary to call this if unit test already calls
// g_test_dbus_up() before acquiring the session bus.
func TestDBusUnset() {
	girepository.MustFind("Gio", "unset").Invoke(nil, nil)
}
