// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0
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
		{T: externglib.Type(C.g_tls_password_get_type()), F: marshalTLSPassword},
	})
}

type TLSPasswordPrivate struct {
	native C.GTlsPasswordPrivate
}

// WrapTLSPasswordPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTLSPasswordPrivate(ptr unsafe.Pointer) *TLSPasswordPrivate {
	if ptr == nil {
		return nil
	}

	return (*TLSPasswordPrivate)(ptr)
}

func marshalTLSPasswordPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTLSPasswordPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TLSPasswordPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// TLSPassword holds a password used in TLS.
type TLSPassword interface {
	gextras.Objector

	// Description: get a description string about what the password will be
	// used for.
	Description() string
	// Flags: get flags about the password.
	Flags() TLSPasswordFlags
	// Value: get the password value. If @length is not nil then it will be
	// filled in with the length of the password value. (Note that the password
	// value is not nul-terminated, so you can only pass nil for @length in
	// contexts where you know the password will have a certain fixed length.)
	Value(length uint) byte
	// Warning: get a user readable translated warning. Usually this warning is
	// a representation of the password flags returned from
	// g_tls_password_get_flags().
	Warning() string
	// SetDescription: set a description string about what the password will be
	// used for.
	SetDescription(description string)
	// SetFlags: set flags about the password.
	SetFlags(flags TLSPasswordFlags)
	// SetValue: set the value for this password. The @value will be copied by
	// the password object.
	//
	// Specify the @length, for a non-nul-terminated password. Pass -1 as
	// @length if using a nul-terminated password, and @length will be
	// calculated automatically. (Note that the terminating nul is not
	// considered part of the password in this case.)
	SetValue(value []byte)
	// SetValueFull: provide the value for this password.
	//
	// The @value will be owned by the password object, and later freed using
	// the @destroy function callback.
	//
	// Specify the @length, for a non-nul-terminated password. Pass -1 as
	// @length if using a nul-terminated password, and @length will be
	// calculated automatically. (Note that the terminating nul is not
	// considered part of the password in this case.)
	SetValueFull(value []byte)
	// SetWarning: set a user readable translated warning. Usually this warning
	// is a representation of the password flags returned from
	// g_tls_password_get_flags().
	SetWarning(warning string)
}

// tlsPassword implements the TLSPassword interface.
type tlsPassword struct {
	gextras.Objector
}

var _ TLSPassword = (*tlsPassword)(nil)

// WrapTLSPassword wraps a GObject to the right type. It is
// primarily used internally.
func WrapTLSPassword(obj *externglib.Object) TLSPassword {
	return TLSPassword{
		Objector: obj,
	}
}

func marshalTLSPassword(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTLSPassword(obj), nil
}

// NewTLSPassword constructs a class TLSPassword.
func NewTLSPassword(flags TLSPasswordFlags, description string) TLSPassword {
	var arg1 C.GTlsPasswordFlags
	var arg2 *C.gchar

	arg1 = (C.GTlsPasswordFlags)(flags)
	arg2 = (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(arg2))

	ret := C.g_tls_password_new(arg1, arg2)

	var ret0 TLSPassword

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(TLSPassword)

	return ret0
}

// Description: get a description string about what the password will be
// used for.
func (p tlsPassword) Description() string {
	var arg0 *C.GTlsPassword

	arg0 = (*C.GTlsPassword)(p.Native())

	ret := C.g_tls_password_get_description(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// Flags: get flags about the password.
func (p tlsPassword) Flags() TLSPasswordFlags {
	var arg0 *C.GTlsPassword

	arg0 = (*C.GTlsPassword)(p.Native())

	ret := C.g_tls_password_get_flags(arg0)

	var ret0 TLSPasswordFlags

	ret0 = TLSPasswordFlags(ret)

	return ret0
}

// Value: get the password value. If @length is not nil then it will be
// filled in with the length of the password value. (Note that the password
// value is not nul-terminated, so you can only pass nil for @length in
// contexts where you know the password will have a certain fixed length.)
func (p tlsPassword) Value(length uint) byte {
	var arg0 *C.GTlsPassword
	var arg1 *C.gsize

	arg0 = (*C.GTlsPassword)(p.Native())
	arg1 = (*C.gsize)(length)

	ret := C.g_tls_password_get_value(arg0, arg1)

	var ret0 byte

	ret0 = byte(ret)

	return ret0
}

// Warning: get a user readable translated warning. Usually this warning is
// a representation of the password flags returned from
// g_tls_password_get_flags().
func (p tlsPassword) Warning() string {
	var arg0 *C.GTlsPassword

	arg0 = (*C.GTlsPassword)(p.Native())

	ret := C.g_tls_password_get_warning(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// SetDescription: set a description string about what the password will be
// used for.
func (p tlsPassword) SetDescription(description string) {
	var arg0 *C.GTlsPassword
	var arg1 *C.gchar

	arg0 = (*C.GTlsPassword)(p.Native())
	arg1 = (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(arg1))

	C.g_tls_password_set_description(arg0, arg1)
}

// SetFlags: set flags about the password.
func (p tlsPassword) SetFlags(flags TLSPasswordFlags) {
	var arg0 *C.GTlsPassword
	var arg1 C.GTlsPasswordFlags

	arg0 = (*C.GTlsPassword)(p.Native())
	arg1 = (C.GTlsPasswordFlags)(flags)

	C.g_tls_password_set_flags(arg0, arg1)
}

// SetValue: set the value for this password. The @value will be copied by
// the password object.
//
// Specify the @length, for a non-nul-terminated password. Pass -1 as
// @length if using a nul-terminated password, and @length will be
// calculated automatically. (Note that the terminating nul is not
// considered part of the password in this case.)
func (p tlsPassword) SetValue(value []byte) {
	var arg0 *C.GTlsPassword
	var arg1 *C.guchar
	var arg2 C.gssize

	arg0 = (*C.GTlsPassword)(p.Native())
	arg1 = (*C.guchar)(unsafe.Pointer(&value[0]))
	arg2 = len(value)
	defer runtime.KeepAlive(value)

	C.g_tls_password_set_value(arg0, arg1, arg2)
}

// SetValueFull: provide the value for this password.
//
// The @value will be owned by the password object, and later freed using
// the @destroy function callback.
//
// Specify the @length, for a non-nul-terminated password. Pass -1 as
// @length if using a nul-terminated password, and @length will be
// calculated automatically. (Note that the terminating nul is not
// considered part of the password in this case.)
func (p tlsPassword) SetValueFull(value []byte) {
	var arg0 *C.GTlsPassword
	var arg1 *C.guchar
	var arg2 C.gssize
	var arg3 C.GDestroyNotify

	arg0 = (*C.GTlsPassword)(p.Native())
	arg1 = (*C.guchar)(unsafe.Pointer(&value[0]))
	arg2 = len(value)
	defer runtime.KeepAlive(value)

	C.g_tls_password_set_value_full(arg0, arg1, arg2, arg3)
}

// SetWarning: set a user readable translated warning. Usually this warning
// is a representation of the password flags returned from
// g_tls_password_get_flags().
func (p tlsPassword) SetWarning(warning string) {
	var arg0 *C.GTlsPassword
	var arg1 *C.gchar

	arg0 = (*C.GTlsPassword)(p.Native())
	arg1 = (*C.gchar)(C.CString(warning))
	defer C.free(unsafe.Pointer(arg1))

	C.g_tls_password_set_warning(arg0, arg1)
}
