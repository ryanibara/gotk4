// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// DBUS_METHOD_INVOCATION_HANDLED: value returned by handlers of the signals
// generated by the gdbus-codegen tool to indicate that a method call has been
// handled by an implementation. It is equal to TRUE, but using this macro is
// sometimes more readable.
//
// In code that needs to be backwards-compatible with older GLib, use TRUE
// instead, often written like this:
//
//    g_dbus_method_invocation_return_error (invocation, ...);
//    return TRUE;    // handled.
const DBUS_METHOD_INVOCATION_HANDLED = true

// DBUS_METHOD_INVOCATION_UNHANDLED: value returned by handlers of the signals
// generated by the gdbus-codegen tool to indicate that a method call has not
// been handled by an implementation. It is equal to FALSE, but using this macro
// is sometimes more readable.
//
// In code that needs to be backwards-compatible with older GLib, use FALSE
// instead.
const DBUS_METHOD_INVOCATION_UNHANDLED = false

// Connection gets the BusConnection the method was invoked on.
//
// The function returns the following values:
//
//    - dBusConnection Do not free, it is owned by invocation.
//
func (invocation *DBusMethodInvocation) Connection() *DBusConnection {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(invocation).Native()))
	*(**DBusMethodInvocation)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMethodInvocation").InvokeMethod("get_connection", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(invocation)

	var _dBusConnection *DBusConnection // out

	_dBusConnection = wrapDBusConnection(coreglib.Take(unsafe.Pointer(_cret)))

	return _dBusConnection
}

// InterfaceName gets the name of the D-Bus interface the method was invoked on.
//
// If this method call is a property Get, Set or GetAll call that has been
// redirected to the method call handler then "org.freedesktop.DBus.Properties"
// will be returned. See BusInterfaceVTable for more information.
//
// The function returns the following values:
//
//    - utf8: string. Do not free, it is owned by invocation.
//
func (invocation *DBusMethodInvocation) InterfaceName() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(invocation).Native()))
	*(**DBusMethodInvocation)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMethodInvocation").InvokeMethod("get_interface_name", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(invocation)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Message gets the BusMessage for the method invocation. This is useful if you
// need to use low-level protocol features, such as UNIX file descriptor
// passing, that cannot be properly expressed in the #GVariant API.
//
// See this [server][gdbus-server] and [client][gdbus-unix-fd-client] for an
// example of how to use this low-level API to send and receive UNIX file
// descriptors.
//
// The function returns the following values:
//
//    - dBusMessage Do not free, it is owned by invocation.
//
func (invocation *DBusMethodInvocation) Message() *DBusMessage {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(invocation).Native()))
	*(**DBusMethodInvocation)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMethodInvocation").InvokeMethod("get_message", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(invocation)

	var _dBusMessage *DBusMessage // out

	_dBusMessage = wrapDBusMessage(coreglib.Take(unsafe.Pointer(_cret)))

	return _dBusMessage
}

// MethodInfo gets information about the method call, if any.
//
// If this method invocation is a property Get, Set or GetAll call that has been
// redirected to the method call handler then NULL will be returned. See
// g_dbus_method_invocation_get_property_info() and BusInterfaceVTable for more
// information.
//
// The function returns the following values:
//
//    - dBusMethodInfo (optional) or NULL. Do not free, it is owned by
//      invocation.
//
func (invocation *DBusMethodInvocation) MethodInfo() *DBusMethodInfo {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(invocation).Native()))
	*(**DBusMethodInvocation)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMethodInvocation").InvokeMethod("get_method_info", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(invocation)

	var _dBusMethodInfo *DBusMethodInfo // out

	if _cret != nil {
		_dBusMethodInfo = (*DBusMethodInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.g_dbus_method_info_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_dBusMethodInfo)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_dbus_method_info_unref((*C.GDBusMethodInfo)(intern.C))
			},
		)
	}

	return _dBusMethodInfo
}

// MethodName gets the name of the method that was invoked.
//
// The function returns the following values:
//
//    - utf8: string. Do not free, it is owned by invocation.
//
func (invocation *DBusMethodInvocation) MethodName() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(invocation).Native()))
	*(**DBusMethodInvocation)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMethodInvocation").InvokeMethod("get_method_name", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(invocation)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ObjectPath gets the object path the method was invoked on.
//
// The function returns the following values:
//
//    - utf8: string. Do not free, it is owned by invocation.
//
func (invocation *DBusMethodInvocation) ObjectPath() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(invocation).Native()))
	*(**DBusMethodInvocation)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMethodInvocation").InvokeMethod("get_object_path", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(invocation)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Parameters gets the parameters of the method invocation. If there are no
// input parameters then this will return a GVariant with 0 children rather than
// NULL.
//
// The function returns the following values:
//
//    - variant tuple. Do not unref this because it is owned by invocation.
//
func (invocation *DBusMethodInvocation) Parameters() *glib.Variant {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(invocation).Native()))
	*(**DBusMethodInvocation)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMethodInvocation").InvokeMethod("get_parameters", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(invocation)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	return _variant
}

// PropertyInfo gets information about the property that this method call is
// for, if any.
//
// This will only be set in the case of an invocation in response to a property
// Get or Set call that has been directed to the method call handler for an
// object on account of its property_get() or property_set() vtable pointers
// being unset.
//
// See BusInterfaceVTable for more information.
//
// If the call was GetAll, NULL will be returned.
//
// The function returns the following values:
//
//    - dBusPropertyInfo (optional) or NULL.
//
func (invocation *DBusMethodInvocation) PropertyInfo() *DBusPropertyInfo {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(invocation).Native()))
	*(**DBusMethodInvocation)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMethodInvocation").InvokeMethod("get_property_info", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(invocation)

	var _dBusPropertyInfo *DBusPropertyInfo // out

	if _cret != nil {
		_dBusPropertyInfo = (*DBusPropertyInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.g_dbus_property_info_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_dBusPropertyInfo)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_dbus_property_info_unref((*C.GDBusPropertyInfo)(intern.C))
			},
		)
	}

	return _dBusPropertyInfo
}

// Sender gets the bus name that invoked the method.
//
// The function returns the following values:
//
//    - utf8: string. Do not free, it is owned by invocation.
//
func (invocation *DBusMethodInvocation) Sender() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(invocation).Native()))
	*(**DBusMethodInvocation)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMethodInvocation").InvokeMethod("get_sender", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(invocation)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ReturnDBusError finishes handling a D-Bus method call by returning an error.
//
// This method will take ownership of invocation. See BusInterfaceVTable for
// more information about the ownership of invocation.
//
// The function takes the following parameters:
//
//    - errorName: valid D-Bus error name.
//    - errorMessage: valid D-Bus error message.
//
func (invocation *DBusMethodInvocation) ReturnDBusError(errorName, errorMessage string) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(invocation).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(invocation).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(errorName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.void)(unsafe.Pointer(C.CString(errorMessage)))
	defer C.free(unsafe.Pointer(_arg2))
	*(**DBusMethodInvocation)(unsafe.Pointer(&args[1])) = _arg1
	*(*string)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gio", "DBusMethodInvocation").InvokeMethod("return_dbus_error", args[:], nil)

	runtime.KeepAlive(invocation)
	runtime.KeepAlive(errorName)
	runtime.KeepAlive(errorMessage)
}

// ReturnGError: like g_dbus_method_invocation_return_error() but takes a
// #GError instead of the error domain, error code and message.
//
// This method will take ownership of invocation. See BusInterfaceVTable for
// more information about the ownership of invocation.
//
// The function takes the following parameters:
//
//    - err: #GError.
//
func (invocation *DBusMethodInvocation) ReturnGError(err error) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(invocation).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(invocation).Native()))
	if err != nil {
		_arg1 = (*C.void)(gerror.New(err))
	}
	*(**DBusMethodInvocation)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gio", "DBusMethodInvocation").InvokeMethod("return_gerror", args[:], nil)

	runtime.KeepAlive(invocation)
	runtime.KeepAlive(err)
}

// ReturnValue finishes handling a D-Bus method call by returning parameters. If
// the parameters GVariant is floating, it is consumed.
//
// It is an error if parameters is not of the right format: it must be a tuple
// containing the out-parameters of the D-Bus method. Even if the method has a
// single out-parameter, it must be contained in a tuple. If the method has no
// out-parameters, parameters may be NULL or an empty tuple.
//
//    GDBusMethodInvocation *invocation = some_invocation;
//    g_autofree gchar *result_string = NULL;
//    g_autoptr (GError) error = NULL;
//
//    result_string = calculate_result (&error);
//
//    if (error != NULL)
//      g_dbus_method_invocation_return_gerror (invocation, error);
//    else
//      g_dbus_method_invocation_return_value (invocation,
//                                             g_variant_new ("(s)", result_string));
//
//    // Do not free invocation here; returning a value does that
//
// This method will take ownership of invocation. See BusInterfaceVTable for
// more information about the ownership of invocation.
//
// Since 2.48, if the method call requested for a reply not to be sent then this
// call will sink parameters and free invocation, but otherwise do nothing (as
// per the recommendations of the D-Bus specification).
//
// The function takes the following parameters:
//
//    - parameters (optional) tuple with out parameters for the method or NULL if
//      not passing any parameters.
//
func (invocation *DBusMethodInvocation) ReturnValue(parameters *glib.Variant) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(invocation).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(invocation).Native()))
	if parameters != nil {
		_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(parameters)))
	}
	*(**DBusMethodInvocation)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gio", "DBusMethodInvocation").InvokeMethod("return_value", args[:], nil)

	runtime.KeepAlive(invocation)
	runtime.KeepAlive(parameters)
}
