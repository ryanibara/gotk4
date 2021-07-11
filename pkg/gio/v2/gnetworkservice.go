// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_network_service_get_type()), F: marshalNetworkServicer},
	})
}

// NetworkServicer describes NetworkService's methods.
type NetworkServicer interface {
	// Domain gets the domain that @srv serves.
	Domain() string
	// Protocol gets @srv's protocol name (eg, "tcp").
	Protocol() string
	// Scheme gets the URI scheme used to resolve proxies.
	Scheme() string
	// Service gets @srv's service name (eg, "ldap").
	Service() string
	// SetScheme set's the URI scheme used to resolve proxies.
	SetScheme(scheme string)
}

// NetworkService: like Address does with hostnames, Service provides an easy
// way to resolve a SRV record, and then attempt to connect to one of the hosts
// that implements that service, handling service priority/weighting, multiple
// IP addresses, and multiple address families.
//
// See Target for more information about SRV records, and see Connectable for an
// example of using the connectable interface.
type NetworkService struct {
	*externglib.Object

	SocketConnectable
}

var (
	_ NetworkServicer = (*NetworkService)(nil)
	_ gextras.Nativer = (*NetworkService)(nil)
)

func wrapNetworkService(obj *externglib.Object) NetworkServicer {
	return &NetworkService{
		Object: obj,
		SocketConnectable: SocketConnectable{
			Object: obj,
		},
	}
}

func marshalNetworkServicer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNetworkService(obj), nil
}

// NewNetworkService creates a new Service representing the given @service,
// @protocol, and @domain. This will initially be unresolved; use the
// Connectable interface to resolve it.
func NewNetworkService(service string, protocol string, domain string) *NetworkService {
	var _arg1 *C.gchar              // out
	var _arg2 *C.gchar              // out
	var _arg3 *C.gchar              // out
	var _cret *C.GSocketConnectable // in

	_arg1 = (*C.gchar)(C.CString(service))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(protocol))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.g_network_service_new(_arg1, _arg2, _arg3)

	var _networkService *NetworkService // out

	_networkService = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*NetworkService)

	return _networkService
}

// Domain gets the domain that @srv serves. This might be either UTF-8 or
// ASCII-encoded, depending on what @srv was created with.
func (srv *NetworkService) Domain() string {
	var _arg0 *C.GNetworkService // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkService)(unsafe.Pointer(srv.Native()))

	_cret = C.g_network_service_get_domain(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Protocol gets @srv's protocol name (eg, "tcp").
func (srv *NetworkService) Protocol() string {
	var _arg0 *C.GNetworkService // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkService)(unsafe.Pointer(srv.Native()))

	_cret = C.g_network_service_get_protocol(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Scheme gets the URI scheme used to resolve proxies. By default, the service
// name is used as scheme.
func (srv *NetworkService) Scheme() string {
	var _arg0 *C.GNetworkService // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkService)(unsafe.Pointer(srv.Native()))

	_cret = C.g_network_service_get_scheme(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Service gets @srv's service name (eg, "ldap").
func (srv *NetworkService) Service() string {
	var _arg0 *C.GNetworkService // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkService)(unsafe.Pointer(srv.Native()))

	_cret = C.g_network_service_get_service(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// SetScheme set's the URI scheme used to resolve proxies. By default, the
// service name is used as scheme.
func (srv *NetworkService) SetScheme(scheme string) {
	var _arg0 *C.GNetworkService // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GNetworkService)(unsafe.Pointer(srv.Native()))
	_arg1 = (*C.gchar)(C.CString(scheme))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_network_service_set_scheme(_arg0, _arg1)
}
