// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gchar* _gotk4_gio2_TlsPasswordClass_get_default_warning(GTlsPassword*);
// extern guchar* _gotk4_gio2_TlsPasswordClass_get_value(GTlsPassword*, gsize*);
import "C"

// glib.Type values for gtlspassword.go.
var GTypeTLSPassword = coreglib.Type(C.g_tls_password_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeTLSPassword, F: marshalTLSPassword},
	})
}

// TLSPasswordOverrider contains methods that are overridable.
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
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TLSPassword)(nil)
)

func classInitTLSPassworder(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GTlsPasswordClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GTlsPasswordClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ DefaultWarning() string }); ok {
		pclass.get_default_warning = (*[0]byte)(C._gotk4_gio2_TlsPasswordClass_get_default_warning)
	}

	if _, ok := goval.(interface{ Value(length *uint) *byte }); ok {
		pclass.get_value = (*[0]byte)(C._gotk4_gio2_TlsPasswordClass_get_value)
	}
}

//export _gotk4_gio2_TlsPasswordClass_get_default_warning
func _gotk4_gio2_TlsPasswordClass_get_default_warning(arg0 *C.GTlsPassword) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ DefaultWarning() string })

	utf8 := iface.DefaultWarning()

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_gio2_TlsPasswordClass_get_value
func _gotk4_gio2_TlsPasswordClass_get_value(arg0 *C.GTlsPassword, arg1 *C.gsize) (cret *C.guchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Value(length *uint) *byte })

	var _length *uint // out

	if arg1 != nil {
		_length = (*uint)(unsafe.Pointer(arg1))
	}

	guint8 := iface.Value(_length)

	cret = (*C.void)(unsafe.Pointer(guint8))

	return cret
}

func wrapTLSPassword(obj *coreglib.Object) *TLSPassword {
	return &TLSPassword{
		Object: obj,
	}
}

func marshalTLSPassword(p uintptr) (interface{}, error) {
	return wrapTLSPassword(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Description: get a description string about what the password will be used
// for.
//
// The function returns the following values:
//
//    - utf8: description of the password.
//
func (password *TLSPassword) Description() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(password).Native()))
	*(**TLSPassword)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "TlsPassword").InvokeMethod("get_description", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(password)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
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
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(password).Native()))
	if length != nil {
		_arg1 = (*C.void)(unsafe.Pointer(length))
	}
	*(**TLSPassword)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gio", "TlsPassword").InvokeMethod("get_value", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(password).Native()))
	*(**TLSPassword)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "TlsPassword").InvokeMethod("get_warning", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(password).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(description)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**TLSPassword)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gio", "TlsPassword").InvokeMethod("set_description", args[:], nil)

	runtime.KeepAlive(password)
	runtime.KeepAlive(description)
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
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 C.gssize

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(password).Native()))
	_arg2 = (C.gssize)(len(value))
	if len(value) > 0 {
		_arg1 = (*C.void)(unsafe.Pointer(&value[0]))
	}
	*(**TLSPassword)(unsafe.Pointer(&args[1])) = _arg1
	*(*[]byte)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gio", "TlsPassword").InvokeMethod("set_value", args[:], nil)

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
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(password).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(warning)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**TLSPassword)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gio", "TlsPassword").InvokeMethod("set_warning", args[:], nil)

	runtime.KeepAlive(password)
	runtime.KeepAlive(warning)
}
