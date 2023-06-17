// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// DBusErrorEncodeGError creates a D-Bus error name to use for error.
// If error matches a registered error (cf. g_dbus_error_register_error()),
// the corresponding D-Bus error name will be returned.
//
// Otherwise the a name of the form
// org.gtk.GDBus.UnmappedGError.Quark._ESCAPED_QUARK_NAME.Code_ERROR_CODE will
// be used. This allows other GDBus applications to map the error on the wire
// back to a #GError using g_dbus_error_new_for_dbus_error().
//
// This function is typically only used in object mappings to put a #GError on
// the wire. Regular applications should not use it.
//
// The function takes the following parameters:
//
//   - err: #GError.
//
// The function returns the following values:
//
//   - utf8 d-Bus error name (never NULL). Free with g_free().
//
func DBusErrorEncodeGError(err error) string {
	var _arg1 *C.GError // out
	var _cret *C.gchar  // in

	if err != nil {
		_arg1 = (*C.GError)(gerror.New(err))
	}

	_cret = C.g_dbus_error_encode_gerror(_arg1)
	runtime.KeepAlive(err)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// DBusErrorGetRemoteError gets the D-Bus error name used for error, if any.
//
// This function is guaranteed to return a D-Bus error name for all
// #GErrors returned from functions handling remote method calls (e.g.
// g_dbus_connection_call_finish()) unless g_dbus_error_strip_remote_error() has
// been used on error.
//
// The function takes the following parameters:
//
//   - err: #GError.
//
// The function returns the following values:
//
//   - utf8 (optional): allocated string or NULL if the D-Bus error name could
//     not be found. Free with g_free().
//
func DBusErrorGetRemoteError(err error) string {
	var _arg1 *C.GError // out
	var _cret *C.gchar  // in

	if err != nil {
		_arg1 = (*C.GError)(gerror.New(err))
	}

	_cret = C.g_dbus_error_get_remote_error(_arg1)
	runtime.KeepAlive(err)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// DBusErrorIsRemoteError checks if error represents an error received via D-Bus
// from a remote peer. If so, use g_dbus_error_get_remote_error() to get the
// name of the error.
//
// The function takes the following parameters:
//
//   - err: #GError.
//
// The function returns the following values:
//
//   - ok: TRUE if error represents an error from a remote peer, FALSE
//     otherwise.
//
func DBusErrorIsRemoteError(err error) bool {
	var _arg1 *C.GError  // out
	var _cret C.gboolean // in

	if err != nil {
		_arg1 = (*C.GError)(gerror.New(err))
	}

	_cret = C.g_dbus_error_is_remote_error(_arg1)
	runtime.KeepAlive(err)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NewDBusErrorForDBusError creates a #GError based on the contents of
// dbus_error_name and dbus_error_message.
//
// Errors registered with g_dbus_error_register_error() will be looked up
// using dbus_error_name and if a match is found, the error domain and code
// is used. Applications can use g_dbus_error_get_remote_error() to recover
// dbus_error_name.
//
// If a match against a registered error is not found and the D-Bus error
// name is in a form as returned by g_dbus_error_encode_gerror() the error
// domain and code encoded in the name is used to create the #GError. Also,
// dbus_error_name is added to the error message such that it can be recovered
// with g_dbus_error_get_remote_error().
//
// Otherwise, a #GError with the error code G_IO_ERROR_DBUS_ERROR
// in the IO_ERROR error domain is returned. Also, dbus_error_name
// is added to the error message such that it can be recovered with
// g_dbus_error_get_remote_error().
//
// In all three cases, dbus_error_name can always be recovered from the
// returned #GError using the g_dbus_error_get_remote_error() function (unless
// g_dbus_error_strip_remote_error() hasn't been used on the returned error).
//
// This function is typically only used in object mappings to prepare #GError
// instances for applications. Regular applications should not use it.
//
// The function takes the following parameters:
//
//   - dbusErrorName d-Bus error name.
//   - dbusErrorMessage d-Bus error message.
//
// The function returns the following values:
//
//   - err: allocated #GError. Free with g_error_free().
//
func NewDBusErrorForDBusError(dbusErrorName, dbusErrorMessage string) error {
	var _arg1 *C.gchar  // out
	var _arg2 *C.gchar  // out
	var _cret *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(dbusErrorName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(dbusErrorMessage)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_dbus_error_new_for_dbus_error(_arg1, _arg2)
	runtime.KeepAlive(dbusErrorName)
	runtime.KeepAlive(dbusErrorMessage)

	var _err error // out

	_err = gerror.Take(unsafe.Pointer(_cret))

	return _err
}

// DBusErrorRegisterError creates an association to map between dbus_error_name
// and #GErrors specified by error_domain and error_code.
//
// This is typically done in the routine that returns the #GQuark for an error
// domain.
//
// The function takes the following parameters:
//
//   - errorDomain for an error domain.
//   - errorCode: error code.
//   - dbusErrorName d-Bus error name.
//
// The function returns the following values:
//
//   - ok: TRUE if the association was created, FALSE if it already exists.
//
func DBusErrorRegisterError(errorDomain glib.Quark, errorCode int, dbusErrorName string) bool {
	var _arg1 C.GQuark   // out
	var _arg2 C.gint     // out
	var _arg3 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = C.guint32(errorDomain)
	type _ = glib.Quark
	type _ = uint32
	_arg2 = C.gint(errorCode)
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(dbusErrorName)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.g_dbus_error_register_error(_arg1, _arg2, _arg3)
	runtime.KeepAlive(errorDomain)
	runtime.KeepAlive(errorCode)
	runtime.KeepAlive(dbusErrorName)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DBusErrorRegisterErrorDomain: helper function for associating a #GError error
// domain with D-Bus error names.
//
// While quark_volatile has a volatile qualifier, this is a historical artifact
// and the argument passed to it should not be volatile.
//
// The function takes the following parameters:
//
//   - errorDomainQuarkName: error domain name.
//   - quarkVolatile: pointer where to store the #GQuark.
//   - entries: pointer to num_entries BusErrorEntry struct items.
//
func DBusErrorRegisterErrorDomain(errorDomainQuarkName string, quarkVolatile *uint, entries []DBusErrorEntry) {
	var _arg1 *C.gchar           // out
	var _arg2 *C.gsize           // out
	var _arg3 *C.GDBusErrorEntry // out
	var _arg4 C.guint

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(errorDomainQuarkName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gsize)(unsafe.Pointer(quarkVolatile))
	_arg4 = (C.guint)(len(entries))
	_arg3 = (*C.GDBusErrorEntry)(C.calloc(C.size_t(len(entries)), C.size_t(C.sizeof_GDBusErrorEntry)))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice((*C.GDBusErrorEntry)(_arg3), len(entries))
		for i := range entries {
			out[i] = *(*C.GDBusErrorEntry)(gextras.StructNative(unsafe.Pointer((&entries[i]))))
		}
	}

	C.g_dbus_error_register_error_domain(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(errorDomainQuarkName)
	runtime.KeepAlive(quarkVolatile)
	runtime.KeepAlive(entries)
}

// DBusErrorStripRemoteError looks for extra information in the error message
// used to recover the D-Bus error name and strips it if found. If stripped,
// the message field in error will correspond exactly to what was received on
// the wire.
//
// This is typically used when presenting errors to the end user.
//
// The function takes the following parameters:
//
//   - err: #GError.
//
// The function returns the following values:
//
//   - ok: TRUE if information was stripped, FALSE otherwise.
//
func DBusErrorStripRemoteError(err error) bool {
	var _arg1 *C.GError  // out
	var _cret C.gboolean // in

	if err != nil {
		_arg1 = (*C.GError)(gerror.New(err))
	}

	_cret = C.g_dbus_error_strip_remote_error(_arg1)
	runtime.KeepAlive(err)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DBusErrorUnregisterError destroys an association previously set up with
// g_dbus_error_register_error().
//
// The function takes the following parameters:
//
//   - errorDomain for an error domain.
//   - errorCode: error code.
//   - dbusErrorName d-Bus error name.
//
// The function returns the following values:
//
//   - ok: TRUE if the association was destroyed, FALSE if it wasn't found.
//
func DBusErrorUnregisterError(errorDomain glib.Quark, errorCode int, dbusErrorName string) bool {
	var _arg1 C.GQuark   // out
	var _arg2 C.gint     // out
	var _arg3 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = C.guint32(errorDomain)
	type _ = glib.Quark
	type _ = uint32
	_arg2 = C.gint(errorCode)
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(dbusErrorName)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.g_dbus_error_unregister_error(_arg1, _arg2, _arg3)
	runtime.KeepAlive(errorDomain)
	runtime.KeepAlive(errorCode)
	runtime.KeepAlive(dbusErrorName)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DBusErrorEntry: struct used in g_dbus_error_register_error_domain().
//
// An instance of this type is always passed by reference.
type DBusErrorEntry struct {
	*dBusErrorEntry
}

// dBusErrorEntry is the struct that's finalized.
type dBusErrorEntry struct {
	native *C.GDBusErrorEntry
}

// ErrorCode: error code.
func (d *DBusErrorEntry) ErrorCode() int {
	valptr := &d.native.error_code
	var _v int // out
	_v = int(*valptr)
	return _v
}

// DBusErrorName d-Bus error name to associate with error_code.
func (d *DBusErrorEntry) DBusErrorName() string {
	valptr := &d.native.dbus_error_name
	var _v string // out
	_v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return _v
}

// ErrorCode: error code.
func (d *DBusErrorEntry) SetErrorCode(errorCode int) {
	valptr := &d.native.error_code
	*valptr = C.gint(errorCode)
}
