// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gcredentials.go.
var GTypeCredentials = externglib.Type(C.g_credentials_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeCredentials, F: marshalCredentials},
	})
}

// CredentialsOverrider contains methods that are overridable.
type CredentialsOverrider interface {
}

// Credentials type is a reference-counted wrapper for native credentials. This
// information is typically used for identifying, authenticating and authorizing
// other processes.
//
// Some operating systems supports looking up the credentials of the remote peer
// of a communication endpoint - see e.g. g_socket_get_credentials().
//
// Some operating systems supports securely sending and receiving credentials
// over a Unix Domain Socket, see CredentialsMessage,
// g_unix_connection_send_credentials() and
// g_unix_connection_receive_credentials() for details.
//
// On Linux, the native credential type is a struct ucred - see the unix(7) man
// page for details. This corresponds to G_CREDENTIALS_TYPE_LINUX_UCRED.
//
// On Apple operating systems (including iOS, tvOS, and macOS), the native
// credential type is a struct xucred. This corresponds to
// G_CREDENTIALS_TYPE_APPLE_XUCRED.
//
// On FreeBSD, Debian GNU/kFreeBSD, and GNU/Hurd, the native credential type is
// a struct cmsgcred. This corresponds to G_CREDENTIALS_TYPE_FREEBSD_CMSGCRED.
//
// On NetBSD, the native credential type is a struct unpcbid. This corresponds
// to G_CREDENTIALS_TYPE_NETBSD_UNPCBID.
//
// On OpenBSD, the native credential type is a struct sockpeercred. This
// corresponds to G_CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED.
//
// On Solaris (including OpenSolaris and its derivatives), the native credential
// type is a ucred_t. This corresponds to G_CREDENTIALS_TYPE_SOLARIS_UCRED.
type Credentials struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Credentials)(nil)
)

func classInitCredentialser(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapCredentials(obj *externglib.Object) *Credentials {
	return &Credentials{
		Object: obj,
	}
}

func marshalCredentials(p uintptr) (interface{}, error) {
	return wrapCredentials(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCredentials creates a new #GCredentials object with credentials matching
// the the current process.
//
// The function returns the following values:
//
//    - credentials Free with g_object_unref().
//
func NewCredentials() *Credentials {
	var _cret *C.GCredentials // in

	_cret = C.g_credentials_new()

	var _credentials *Credentials // out

	_credentials = wrapCredentials(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _credentials
}

// IsSameUser checks if credentials and other_credentials is the same user.
//
// This operation can fail if #GCredentials is not supported on the the OS.
//
// The function takes the following parameters:
//
//    - otherCredentials: #GCredentials.
//
func (credentials *Credentials) IsSameUser(otherCredentials *Credentials) error {
	var _arg0 *C.GCredentials // out
	var _arg1 *C.GCredentials // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GCredentials)(unsafe.Pointer(externglib.InternObject(credentials).Native()))
	_arg1 = (*C.GCredentials)(unsafe.Pointer(externglib.InternObject(otherCredentials).Native()))

	C.g_credentials_is_same_user(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(credentials)
	runtime.KeepAlive(otherCredentials)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetNative copies the native credentials of type native_type from native into
// credentials.
//
// It is a programming error (which will cause a warning to be logged) to use
// this method if there is no #GCredentials support for the OS or if native_type
// isn't supported by the OS.
//
// The function takes the following parameters:
//
//    - nativeType: type of native credentials to set.
//    - native: pointer to native credentials.
//
func (credentials *Credentials) SetNative(nativeType CredentialsType, native cgo.Handle) {
	var _arg0 *C.GCredentials    // out
	var _arg1 C.GCredentialsType // out
	var _arg2 C.gpointer         // out

	_arg0 = (*C.GCredentials)(unsafe.Pointer(externglib.InternObject(credentials).Native()))
	_arg1 = C.GCredentialsType(nativeType)
	_arg2 = (C.gpointer)(unsafe.Pointer(native))

	C.g_credentials_set_native(_arg0, _arg1, _arg2)
	runtime.KeepAlive(credentials)
	runtime.KeepAlive(nativeType)
	runtime.KeepAlive(native)
}

// String creates a human-readable textual representation of credentials that
// can be used in logging and debug messages. The format of the returned string
// may change in future GLib release.
//
// The function returns the following values:
//
//    - utf8: string that should be freed with g_free().
//
func (credentials *Credentials) String() string {
	var _arg0 *C.GCredentials // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GCredentials)(unsafe.Pointer(externglib.InternObject(credentials).Native()))

	_cret = C.g_credentials_to_string(_arg0)
	runtime.KeepAlive(credentials)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
