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
// extern GTlsDatabase* _gotk4_gio2_TlsBackendInterface_get_default_database(void*);
// extern gboolean _gotk4_gio2_TlsBackendInterface_supports_dtls(void*);
// extern gboolean _gotk4_gio2_TlsBackendInterface_supports_tls(void*);
import "C"

// glib.Type values for gtlsbackend.go.
var GTypeTLSBackend = coreglib.Type(C.g_tls_backend_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeTLSBackend, F: marshalTLSBackend},
	})
}

// TLS_BACKEND_EXTENSION_POINT_NAME: extension point for TLS functionality via
// Backend. See [Extending GIO][extending-gio].
const TLS_BACKEND_EXTENSION_POINT_NAME = "gio-tls-backend"

// TLSBackendOverrider contains methods that are overridable.
type TLSBackendOverrider interface {
	// DefaultDatabase gets the default Database used to verify TLS connections.
	//
	// The function returns the following values:
	//
	//    - tlsDatabase: default database, which should be unreffed when done.
	//
	DefaultDatabase() TLSDatabaser
	// SupportsDTLS checks if DTLS is supported. DTLS support may not be
	// available even if TLS support is available, and vice-versa.
	//
	// The function returns the following values:
	//
	//    - ok: whether DTLS is supported.
	//
	SupportsDTLS() bool
	// SupportsTLS checks if TLS is supported; if this returns FALSE for the
	// default Backend, it means no "real" TLS backend is available.
	//
	// The function returns the following values:
	//
	//    - ok: whether or not TLS is supported.
	//
	SupportsTLS() bool
}

// TLSBackend: TLS (Transport Layer Security, aka SSL) and DTLS backend.
//
// TLSBackend wraps an interface. This means the user can get the
// underlying type by calling Cast().
type TLSBackend struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TLSBackend)(nil)
)

// TLSBackender describes TLSBackend's interface methods.
type TLSBackender interface {
	coreglib.Objector

	// DefaultDatabase gets the default Database used to verify TLS connections.
	DefaultDatabase() TLSDatabaser
	// SetDefaultDatabase: set the default Database used to verify TLS
	// connections Any subsequent call to g_tls_backend_get_default_database()
	// will return the database set in this call.
	SetDefaultDatabase(database TLSDatabaser)
	// SupportsDTLS checks if DTLS is supported.
	SupportsDTLS() bool
	// SupportsTLS checks if TLS is supported; if this returns FALSE for the
	// default Backend, it means no "real" TLS backend is available.
	SupportsTLS() bool
}

var _ TLSBackender = (*TLSBackend)(nil)

func ifaceInitTLSBackender(gifacePtr, data C.gpointer) {
	iface := (*C.GTlsBackendInterface)(unsafe.Pointer(gifacePtr))
	iface.get_default_database = (*[0]byte)(C._gotk4_gio2_TlsBackendInterface_get_default_database)
	iface.supports_dtls = (*[0]byte)(C._gotk4_gio2_TlsBackendInterface_supports_dtls)
	iface.supports_tls = (*[0]byte)(C._gotk4_gio2_TlsBackendInterface_supports_tls)
}

//export _gotk4_gio2_TlsBackendInterface_get_default_database
func _gotk4_gio2_TlsBackendInterface_get_default_database(arg0 *C.void) (cret *C.GTlsDatabase) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TLSBackendOverrider)

	tlsDatabase := iface.DefaultDatabase()

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(tlsDatabase).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(tlsDatabase).Native()))

	return cret
}

//export _gotk4_gio2_TlsBackendInterface_supports_dtls
func _gotk4_gio2_TlsBackendInterface_supports_dtls(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TLSBackendOverrider)

	ok := iface.SupportsDTLS()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_TlsBackendInterface_supports_tls
func _gotk4_gio2_TlsBackendInterface_supports_tls(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TLSBackendOverrider)

	ok := iface.SupportsTLS()

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapTLSBackend(obj *coreglib.Object) *TLSBackend {
	return &TLSBackend{
		Object: obj,
	}
}

func marshalTLSBackend(p uintptr) (interface{}, error) {
	return wrapTLSBackend(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DefaultDatabase gets the default Database used to verify TLS connections.
//
// The function returns the following values:
//
//    - tlsDatabase: default database, which should be unreffed when done.
//
func (backend *TLSBackend) DefaultDatabase() TLSDatabaser {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(backend).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(backend)

	var _tlsDatabase TLSDatabaser // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.TLSDatabaser is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(TLSDatabaser)
			return ok
		})
		rv, ok := casted.(TLSDatabaser)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.TLSDatabaser")
		}
		_tlsDatabase = rv
	}

	return _tlsDatabase
}

// SetDefaultDatabase: set the default Database used to verify TLS connections
//
// Any subsequent call to g_tls_backend_get_default_database() will return the
// database set in this call. Existing databases and connections are not
// modified.
//
// Setting a NULL default database will reset to using the system default
// database as if g_tls_backend_set_default_database() had never been called.
//
// The function takes the following parameters:
//
//    - database (optional): Database.
//
func (backend *TLSBackend) SetDefaultDatabase(database TLSDatabaser) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(backend).Native()))
	if database != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(database).Native()))
	}

	runtime.KeepAlive(backend)
	runtime.KeepAlive(database)
}

// SupportsDTLS checks if DTLS is supported. DTLS support may not be available
// even if TLS support is available, and vice-versa.
//
// The function returns the following values:
//
//    - ok: whether DTLS is supported.
//
func (backend *TLSBackend) SupportsDTLS() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(backend).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(backend)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SupportsTLS checks if TLS is supported; if this returns FALSE for the default
// Backend, it means no "real" TLS backend is available.
//
// The function returns the following values:
//
//    - ok: whether or not TLS is supported.
//
func (backend *TLSBackend) SupportsTLS() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(backend).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(backend)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// TLSBackendGetDefault gets the default Backend for the system.
//
// The function returns the following values:
//
//    - tlsBackend which will be a dummy object if no TLS backend is available.
//
func TLSBackendGetDefault() *TLSBackend {
	_gret := girepository.MustFind("Gio", "get_default").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _tlsBackend *TLSBackend // out

	_tlsBackend = wrapTLSBackend(coreglib.Take(unsafe.Pointer(_cret)))

	return _tlsBackend
}
