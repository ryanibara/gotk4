// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
// #include <glib-object.h>
//
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tls_database_get_type()), F: marshalTLSDatabase},
	})
}

// TLSDatabase is used to look up certificates and other information from a
// certificate or key store. It is an abstract base class which TLS library
// specific subtypes override.
//
// A Database may be accessed from multiple threads by the TLS backend. All
// implementations are required to be fully thread-safe.
//
// Most common client applications will not directly interact with Database. It
// is used internally by Connection.
type TLSDatabase interface {
	gextras.Objector

	CreateCertificateHandleTLSDatabase(certificate TLSCertificate) string

	LookupCertificateForHandleTLSDatabase(handle string, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable) (TLSCertificate, error)

	LookupCertificateForHandleAsyncTLSDatabase(handle string, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable, callback AsyncReadyCallback)

	LookupCertificateForHandleFinishTLSDatabase(result AsyncResult) (TLSCertificate, error)

	LookupCertificateIssuerTLSDatabase(certificate TLSCertificate, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable) (TLSCertificate, error)

	LookupCertificateIssuerAsyncTLSDatabase(certificate TLSCertificate, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable, callback AsyncReadyCallback)

	LookupCertificateIssuerFinishTLSDatabase(result AsyncResult) (TLSCertificate, error)

	LookupCertificatesIssuedByAsyncTLSDatabase(issuerRawDn []byte, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable, callback AsyncReadyCallback)

	VerifyChainTLSDatabase(chain TLSCertificate, purpose string, identity SocketConnectable, interaction TLSInteraction, flags TLSDatabaseVerifyFlags, cancellable Cancellable) (TLSCertificateFlags, error)

	VerifyChainAsyncTLSDatabase(chain TLSCertificate, purpose string, identity SocketConnectable, interaction TLSInteraction, flags TLSDatabaseVerifyFlags, cancellable Cancellable, callback AsyncReadyCallback)

	VerifyChainFinishTLSDatabase(result AsyncResult) (TLSCertificateFlags, error)
}

// tlsDatabase implements the TLSDatabase class.
type tlsDatabase struct {
	gextras.Objector
}

// WrapTLSDatabase wraps a GObject to the right type. It is
// primarily used internally.
func WrapTLSDatabase(obj *externglib.Object) TLSDatabase {
	return tlsDatabase{
		Objector: obj,
	}
}

func marshalTLSDatabase(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTLSDatabase(obj), nil
}

func (s tlsDatabase) CreateCertificateHandleTLSDatabase(certificate TLSCertificate) string {
	var _arg0 *C.GTlsDatabase    // out
	var _arg1 *C.GTlsCertificate // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(certificate.Native()))

	_cret = C.g_tls_database_create_certificate_handle(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (s tlsDatabase) LookupCertificateForHandleTLSDatabase(handle string, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable) (TLSCertificate, error) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg1 *C.gchar                  // out
	var _arg2 *C.GTlsInteraction        // out
	var _arg3 C.GTlsDatabaseLookupFlags // out
	var _arg4 *C.GCancellable           // out
	var _cret *C.GTlsCertificate        // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(handle))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	_arg3 = C.GTlsDatabaseLookupFlags(flags)
	_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_tls_database_lookup_certificate_for_handle(_arg0, _arg1, _arg2, _arg3, _arg4, &_cerr)

	var _tlsCertificate TLSCertificate // out
	var _goerr error                   // out

	_tlsCertificate = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(TLSCertificate)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

func (s tlsDatabase) LookupCertificateForHandleAsyncTLSDatabase(handle string, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg1 *C.gchar                  // out
	var _arg2 *C.GTlsInteraction        // out
	var _arg3 C.GTlsDatabaseLookupFlags // out
	var _arg4 *C.GCancellable           // out
	var _arg5 C.GAsyncReadyCallback     // out
	var _arg6 C.gpointer

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(handle))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	_arg3 = C.GTlsDatabaseLookupFlags(flags)
	_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg5 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg6 = C.gpointer(box.Assign(callback))

	C.g_tls_database_lookup_certificate_for_handle_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

func (s tlsDatabase) LookupCertificateForHandleFinishTLSDatabase(result AsyncResult) (TLSCertificate, error) {
	var _arg0 *C.GTlsDatabase    // out
	var _arg1 *C.GAsyncResult    // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_tls_database_lookup_certificate_for_handle_finish(_arg0, _arg1, &_cerr)

	var _tlsCertificate TLSCertificate // out
	var _goerr error                   // out

	_tlsCertificate = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(TLSCertificate)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

func (s tlsDatabase) LookupCertificateIssuerTLSDatabase(certificate TLSCertificate, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable) (TLSCertificate, error) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg1 *C.GTlsCertificate        // out
	var _arg2 *C.GTlsInteraction        // out
	var _arg3 C.GTlsDatabaseLookupFlags // out
	var _arg4 *C.GCancellable           // out
	var _cret *C.GTlsCertificate        // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(certificate.Native()))
	_arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	_arg3 = C.GTlsDatabaseLookupFlags(flags)
	_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_tls_database_lookup_certificate_issuer(_arg0, _arg1, _arg2, _arg3, _arg4, &_cerr)

	var _tlsCertificate TLSCertificate // out
	var _goerr error                   // out

	_tlsCertificate = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(TLSCertificate)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

func (s tlsDatabase) LookupCertificateIssuerAsyncTLSDatabase(certificate TLSCertificate, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg1 *C.GTlsCertificate        // out
	var _arg2 *C.GTlsInteraction        // out
	var _arg3 C.GTlsDatabaseLookupFlags // out
	var _arg4 *C.GCancellable           // out
	var _arg5 C.GAsyncReadyCallback     // out
	var _arg6 C.gpointer

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(certificate.Native()))
	_arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	_arg3 = C.GTlsDatabaseLookupFlags(flags)
	_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg5 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg6 = C.gpointer(box.Assign(callback))

	C.g_tls_database_lookup_certificate_issuer_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

func (s tlsDatabase) LookupCertificateIssuerFinishTLSDatabase(result AsyncResult) (TLSCertificate, error) {
	var _arg0 *C.GTlsDatabase    // out
	var _arg1 *C.GAsyncResult    // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_tls_database_lookup_certificate_issuer_finish(_arg0, _arg1, &_cerr)

	var _tlsCertificate TLSCertificate // out
	var _goerr error                   // out

	_tlsCertificate = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(TLSCertificate)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

func (s tlsDatabase) LookupCertificatesIssuedByAsyncTLSDatabase(issuerRawDn []byte, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GTlsDatabase // out
	var _arg1 *C.GByteArray
	var _arg2 *C.GTlsInteraction        // out
	var _arg3 C.GTlsDatabaseLookupFlags // out
	var _arg4 *C.GCancellable           // out
	var _arg5 C.GAsyncReadyCallback     // out
	var _arg6 C.gpointer

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = C.g_byte_array_new_take((*C.guint8)(&issuerRawDn[0]), C.size(len(issuerRawDn)))
	defer C.g_byte_array_steal(issuerRawDn, nil)
	_arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	_arg3 = C.GTlsDatabaseLookupFlags(flags)
	_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg5 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg6 = C.gpointer(box.Assign(callback))

	C.g_tls_database_lookup_certificates_issued_by_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

func (s tlsDatabase) VerifyChainTLSDatabase(chain TLSCertificate, purpose string, identity SocketConnectable, interaction TLSInteraction, flags TLSDatabaseVerifyFlags, cancellable Cancellable) (TLSCertificateFlags, error) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg1 *C.GTlsCertificate        // out
	var _arg2 *C.gchar                  // out
	var _arg3 *C.GSocketConnectable     // out
	var _arg4 *C.GTlsInteraction        // out
	var _arg5 C.GTlsDatabaseVerifyFlags // out
	var _arg6 *C.GCancellable           // out
	var _cret C.GTlsCertificateFlags    // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(chain.Native()))
	_arg2 = (*C.gchar)(C.CString(purpose))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GSocketConnectable)(unsafe.Pointer(identity.Native()))
	_arg4 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	_arg5 = C.GTlsDatabaseVerifyFlags(flags)
	_arg6 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_tls_database_verify_chain(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, &_cerr)

	var _tlsCertificateFlags TLSCertificateFlags // out
	var _goerr error                             // out

	_tlsCertificateFlags = TLSCertificateFlags(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificateFlags, _goerr
}

func (s tlsDatabase) VerifyChainAsyncTLSDatabase(chain TLSCertificate, purpose string, identity SocketConnectable, interaction TLSInteraction, flags TLSDatabaseVerifyFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg1 *C.GTlsCertificate        // out
	var _arg2 *C.gchar                  // out
	var _arg3 *C.GSocketConnectable     // out
	var _arg4 *C.GTlsInteraction        // out
	var _arg5 C.GTlsDatabaseVerifyFlags // out
	var _arg6 *C.GCancellable           // out
	var _arg7 C.GAsyncReadyCallback     // out
	var _arg8 C.gpointer

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(chain.Native()))
	_arg2 = (*C.gchar)(C.CString(purpose))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GSocketConnectable)(unsafe.Pointer(identity.Native()))
	_arg4 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	_arg5 = C.GTlsDatabaseVerifyFlags(flags)
	_arg6 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg7 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg8 = C.gpointer(box.Assign(callback))

	C.g_tls_database_verify_chain_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
}

func (s tlsDatabase) VerifyChainFinishTLSDatabase(result AsyncResult) (TLSCertificateFlags, error) {
	var _arg0 *C.GTlsDatabase        // out
	var _arg1 *C.GAsyncResult        // out
	var _cret C.GTlsCertificateFlags // in
	var _cerr *C.GError              // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_tls_database_verify_chain_finish(_arg0, _arg1, &_cerr)

	var _tlsCertificateFlags TLSCertificateFlags // out
	var _goerr error                             // out

	_tlsCertificateFlags = TLSCertificateFlags(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificateFlags, _goerr
}
