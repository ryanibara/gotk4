// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// DBusEscapeObjectPath: this is a language binding friendly version of
// g_dbus_escape_object_path_bytestring().
//
// The function takes the following parameters:
//
//   - s: string to escape.
//
// The function returns the following values:
//
//   - utf8: escaped version of s. Free with g_free().
//
func DBusEscapeObjectPath(s string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(s)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_escape_object_path(_arg1)
	runtime.KeepAlive(s)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// DBusEscapeObjectPathBytestring escapes bytes for use in a D-Bus object path
// component. bytes is an array of zero or more nonzero bytes in an unspecified
// encoding, followed by a single zero byte.
//
// The escaping method consists of replacing all non-alphanumeric characters
// (see g_ascii_isalnum()) with their hexadecimal value preceded by an
// underscore (_). For example: foo.bar.baz will become foo_2ebar_2ebaz.
//
// This method is appropriate to use when the input is nearly a valid object
// path component but is not when your input is far from being a valid object
// path component. Other escaping algorithms are also valid to use with D-Bus
// object paths.
//
// This can be reversed with g_dbus_unescape_object_path().
//
// The function takes the following parameters:
//
//   - bytes: string of bytes to escape.
//
// The function returns the following values:
//
//   - utf8: escaped version of bytes. Free with g_free().
//
func DBusEscapeObjectPathBytestring(bytes []byte) string {
	var _arg1 *C.guint8 // out
	var _cret *C.gchar  // in

	{
		var zero byte
		bytes = append(bytes, zero)
		_arg1 = (*C.guint8)(unsafe.Pointer(&bytes[0]))
	}

	_cret = C.g_dbus_escape_object_path_bytestring(_arg1)
	runtime.KeepAlive(bytes)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// DBusUnescapeObjectPath unescapes an string that was previously escaped with
// g_dbus_escape_object_path(). If the string is in a format that could not have
// been returned by g_dbus_escape_object_path(), this function returns NULL.
//
// Encoding alphanumeric characters which do not need to be encoded is not
// allowed (e.g _63 is not valid, the string should contain c instead).
//
// The function takes the following parameters:
//
//   - s: string to unescape.
//
// The function returns the following values:
//
//   - guint8s (optional): an unescaped version of s, or NULL if s is not a
//     string returned from g_dbus_escape_object_path(). Free with g_free().
//
func DBusUnescapeObjectPath(s string) []byte {
	var _arg1 *C.gchar  // out
	var _cret *C.guint8 // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(s)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_unescape_object_path(_arg1)
	runtime.KeepAlive(s)

	var _guint8s []byte // out

	if _cret != nil {
		{
			var i int
			var z C.guint8
			for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
				i++
			}

			src := unsafe.Slice(_cret, i)
			_guint8s = make([]byte, i)
			for i := range src {
				_guint8s[i] = byte(src[i])
			}
		}
	}

	return _guint8s
}
