// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// gboolean _gotk4_gio2_DTLSConnection_virtual_accept_certificate(void* fnptr, GDtlsConnection* arg0, GTlsCertificate* arg1, GTlsCertificateFlags arg2) {
//   return ((gboolean (*)(GDtlsConnection*, GTlsCertificate*, GTlsCertificateFlags))(fnptr))(arg0, arg1, arg2);
// };
// gboolean _gotk4_gio2_DTLSConnection_virtual_get_binding_data(void* fnptr, GDtlsConnection* arg0, GTlsChannelBindingType arg1, GByteArray* arg2, GError** arg3) {
//   return ((gboolean (*)(GDtlsConnection*, GTlsChannelBindingType, GByteArray*, GError**))(fnptr))(arg0, arg1, arg2, arg3);
// };
import "C"

// The function takes the following parameters:
//
//    - peerCert
//    - errors
//
// The function returns the following values:
//
func (connection *DTLSConnection) acceptCertificate(peerCert TLSCertificater, errors TLSCertificateFlags) bool {
	gclass := (*C.GDtlsConnectionInterface)(coreglib.PeekParentClass(connection))
	fnarg := gclass.accept_certificate

	var _arg0 *C.GDtlsConnection     // out
	var _arg1 *C.GTlsCertificate     // out
	var _arg2 C.GTlsCertificateFlags // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(coreglib.InternObject(peerCert).Native()))
	_arg2 = C.GTlsCertificateFlags(errors)

	_cret = C._gotk4_gio2_DTLSConnection_virtual_accept_certificate(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(peerCert)
	runtime.KeepAlive(errors)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
//    - typ
//    - data
//
func (conn *DTLSConnection) bindingData(typ TLSChannelBindingType, data []byte) error {
	gclass := (*C.GDtlsConnectionInterface)(coreglib.PeekParentClass(conn))
	fnarg := gclass.get_binding_data

	var _arg0 *C.GDtlsConnection       // out
	var _arg1 C.GTlsChannelBindingType // out
	var _arg2 *C.GByteArray            // out
	var _cerr *C.GError                // in

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	_arg1 = C.GTlsChannelBindingType(typ)
	_arg2 = C.g_byte_array_sized_new(C.guint(len(data)))
	if len(data) > 0 {
		_arg2 = C.g_byte_array_append(_arg2, (*C.guint8)(&data[0]), C.guint(len(data)))
	}
	defer C.g_byte_array_unref(_arg2)

	C._gotk4_gio2_DTLSConnection_virtual_get_binding_data(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(data)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
