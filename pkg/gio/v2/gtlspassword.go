// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"sync"
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
		{T: externglib.Type(C.g_tls_password_get_type()), F: marshalTLSPassworder},
	})
}

// TLSPasswordOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TLSPasswordOverrider interface {
	// The function returns the following values:
	//
	DefaultWarning() string
	// Value: get the password value. If length is not NULL then it will be
	// filled in with the length of the password value. (Note that the password
	// value is not nul-terminated, so you can only pass NULL for length in
	// contexts where you know the password will have a certain fixed length.).
	//
	// The function takes the following parameters:
	//
	//    - length (optional): location to place the length of the password.
	//
	// The function returns the following values:
	//
	//    - guint8: password value (owned by the password object).
	//
	Value(length *uint) *byte
}

// TLSPassword holds a password used in TLS.
type TLSPassword struct {
	*externglib.Object

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ externglib.Objector = (*TLSPassword)(nil)
)

func wrapTLSPassword(obj *externglib.Object) *TLSPassword {
	return &TLSPassword{
		Object: obj,
	}
}

func marshalTLSPassworder(p uintptr) (interface{}, error) {
	return wrapTLSPassword(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTLSPassword: create a new Password object.
//
// The function takes the following parameters:
//
//    - flags: password flags.
//    - description of what the password is for.
//
// The function returns the following values:
//
//    - tlsPassword: newly allocated password object.
//
func NewTLSPassword(flags TLSPasswordFlags, description string) *TLSPassword {
	var _arg1 C.GTlsPasswordFlags // out
	var _arg2 *C.gchar            // out
	var _cret *C.GTlsPassword     // in

	_arg1 = C.GTlsPasswordFlags(flags)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(description)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_tls_password_new(_arg1, _arg2)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(description)

	var _tlsPassword *TLSPassword // out

	_tlsPassword = wrapTLSPassword(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _tlsPassword
}

// Description: get a description string about what the password will be used
// for.
//
// The function returns the following values:
//
//    - utf8: description of the password.
//
func (password *TLSPassword) Description() string {
	var _arg0 *C.GTlsPassword // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))

	_cret = C.g_tls_password_get_description(_arg0)
	runtime.KeepAlive(password)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Flags: get flags about the password.
//
// The function returns the following values:
//
//    - tlsPasswordFlags flags about the password.
//
func (password *TLSPassword) Flags() TLSPasswordFlags {
	var _arg0 *C.GTlsPassword     // out
	var _cret C.GTlsPasswordFlags // in

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))

	_cret = C.g_tls_password_get_flags(_arg0)
	runtime.KeepAlive(password)

	var _tlsPasswordFlags TLSPasswordFlags // out

	_tlsPasswordFlags = TLSPasswordFlags(_cret)

	return _tlsPasswordFlags
}

// Value: get the password value. If length is not NULL then it will be filled
// in with the length of the password value. (Note that the password value is
// not nul-terminated, so you can only pass NULL for length in contexts where
// you know the password will have a certain fixed length.).
//
// The function takes the following parameters:
//
//    - length (optional): location to place the length of the password.
//
// The function returns the following values:
//
//    - guint8: password value (owned by the password object).
//
func (password *TLSPassword) Value(length *uint) *byte {
	var _arg0 *C.GTlsPassword // out
	var _arg1 *C.gsize        // out
	var _cret *C.guchar       // in

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))
	if length != nil {
		_arg1 = (*C.gsize)(unsafe.Pointer(length))
	}

	_cret = C.g_tls_password_get_value(_arg0, _arg1)
	runtime.KeepAlive(password)
	runtime.KeepAlive(length)

	var _guint8 *byte // out

	_guint8 = (*byte)(unsafe.Pointer(_cret))

	return _guint8
}

// Warning: get a user readable translated warning. Usually this warning is a
// representation of the password flags returned from
// g_tls_password_get_flags().
//
// The function returns the following values:
//
//    - utf8: warning.
//
func (password *TLSPassword) Warning() string {
	var _arg0 *C.GTlsPassword // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))

	_cret = C.g_tls_password_get_warning(_arg0)
	runtime.KeepAlive(password)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetDescription: set a description string about what the password will be used
// for.
//
// The function takes the following parameters:
//
//    - description of the password.
//
func (password *TLSPassword) SetDescription(description string) {
	var _arg0 *C.GTlsPassword // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(description)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_tls_password_set_description(_arg0, _arg1)
	runtime.KeepAlive(password)
	runtime.KeepAlive(description)
}

// SetFlags: set flags about the password.
//
// The function takes the following parameters:
//
//    - flags about the password.
//
func (password *TLSPassword) SetFlags(flags TLSPasswordFlags) {
	var _arg0 *C.GTlsPassword     // out
	var _arg1 C.GTlsPasswordFlags // out

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))
	_arg1 = C.GTlsPasswordFlags(flags)

	C.g_tls_password_set_flags(_arg0, _arg1)
	runtime.KeepAlive(password)
	runtime.KeepAlive(flags)
}

// SetValue: set the value for this password. The value will be copied by the
// password object.
//
// Specify the length, for a non-nul-terminated password. Pass -1 as length if
// using a nul-terminated password, and length will be calculated automatically.
// (Note that the terminating nul is not considered part of the password in this
// case.).
//
// The function takes the following parameters:
//
//    - value: new password value.
//
func (password *TLSPassword) SetValue(value []byte) {
	var _arg0 *C.GTlsPassword // out
	var _arg1 *C.guchar       // out
	var _arg2 C.gssize

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))
	_arg2 = (C.gssize)(len(value))
	if len(value) > 0 {
		_arg1 = (*C.guchar)(unsafe.Pointer(&value[0]))
	}

	C.g_tls_password_set_value(_arg0, _arg1, _arg2)
	runtime.KeepAlive(password)
	runtime.KeepAlive(value)
}

// SetWarning: set a user readable translated warning. Usually this warning is a
// representation of the password flags returned from
// g_tls_password_get_flags().
//
// The function takes the following parameters:
//
//    - warning: user readable warning.
//
func (password *TLSPassword) SetWarning(warning string) {
	var _arg0 *C.GTlsPassword // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(warning)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_tls_password_set_warning(_arg0, _arg1)
	runtime.KeepAlive(password)
	runtime.KeepAlive(warning)
}
