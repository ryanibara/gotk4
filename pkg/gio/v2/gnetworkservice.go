// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeNetworkService = coreglib.Type(C.g_network_service_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeNetworkService, F: marshalNetworkService},
	})
}

// NetworkServiceOverrider contains methods that are overridable.
type NetworkServiceOverrider interface {
}

// NetworkService: like Address does with hostnames, Service provides an easy
// way to resolve a SRV record, and then attempt to connect to one of the hosts
// that implements that service, handling service priority/weighting, multiple
// IP addresses, and multiple address families.
//
// See Target for more information about SRV records, and see Connectable for an
// example of using the connectable interface.
type NetworkService struct {
	_ [0]func() // equal guard
	*coreglib.Object

	SocketConnectable
}

var (
	_ coreglib.Objector = (*NetworkService)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeNetworkService,
		GoType:        reflect.TypeOf((*NetworkService)(nil)),
		InitClass:     initClassNetworkService,
		FinalizeClass: finalizeClassNetworkService,
	})
}

func initClassNetworkService(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ InitNetworkService(*NetworkServiceClass) }); ok {
		klass := (*NetworkServiceClass)(gextras.NewStructNative(gclass))
		goval.InitNetworkService(klass)
	}
}

func finalizeClassNetworkService(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ FinalizeNetworkService(*NetworkServiceClass) }); ok {
		klass := (*NetworkServiceClass)(gextras.NewStructNative(gclass))
		goval.FinalizeNetworkService(klass)
	}
}

func wrapNetworkService(obj *coreglib.Object) *NetworkService {
	return &NetworkService{
		Object: obj,
		SocketConnectable: SocketConnectable{
			Object: obj,
		},
	}
}

func marshalNetworkService(p uintptr) (interface{}, error) {
	return wrapNetworkService(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewNetworkService creates a new Service representing the given service,
// protocol, and domain. This will initially be unresolved; use the Connectable
// interface to resolve it.
//
// The function takes the following parameters:
//
//    - service type to look up (eg, "ldap").
//    - protocol: networking protocol to use for service (eg, "tcp").
//    - domain: DNS domain to look up the service in.
//
// The function returns the following values:
//
//    - networkService: new Service.
//
func NewNetworkService(service, protocol, domain string) *NetworkService {
	var _arg1 *C.gchar              // out
	var _arg2 *C.gchar              // out
	var _arg3 *C.gchar              // out
	var _cret *C.GSocketConnectable // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(service)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(protocol)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.g_network_service_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(service)
	runtime.KeepAlive(protocol)
	runtime.KeepAlive(domain)

	var _networkService *NetworkService // out

	_networkService = wrapNetworkService(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _networkService
}

// Domain gets the domain that srv serves. This might be either UTF-8 or
// ASCII-encoded, depending on what srv was created with.
//
// The function returns the following values:
//
//    - utf8 srv's domain name.
//
func (srv *NetworkService) Domain() string {
	var _arg0 *C.GNetworkService // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkService)(unsafe.Pointer(coreglib.InternObject(srv).Native()))

	_cret = C.g_network_service_get_domain(_arg0)
	runtime.KeepAlive(srv)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Protocol gets srv's protocol name (eg, "tcp").
//
// The function returns the following values:
//
//    - utf8 srv's protocol name.
//
func (srv *NetworkService) Protocol() string {
	var _arg0 *C.GNetworkService // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkService)(unsafe.Pointer(coreglib.InternObject(srv).Native()))

	_cret = C.g_network_service_get_protocol(_arg0)
	runtime.KeepAlive(srv)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Scheme gets the URI scheme used to resolve proxies. By default, the service
// name is used as scheme.
//
// The function returns the following values:
//
//    - utf8 srv's scheme name.
//
func (srv *NetworkService) Scheme() string {
	var _arg0 *C.GNetworkService // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkService)(unsafe.Pointer(coreglib.InternObject(srv).Native()))

	_cret = C.g_network_service_get_scheme(_arg0)
	runtime.KeepAlive(srv)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Service gets srv's service name (eg, "ldap").
//
// The function returns the following values:
//
//    - utf8 srv's service name.
//
func (srv *NetworkService) Service() string {
	var _arg0 *C.GNetworkService // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkService)(unsafe.Pointer(coreglib.InternObject(srv).Native()))

	_cret = C.g_network_service_get_service(_arg0)
	runtime.KeepAlive(srv)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetScheme set's the URI scheme used to resolve proxies. By default, the
// service name is used as scheme.
//
// The function takes the following parameters:
//
//    - scheme: URI scheme.
//
func (srv *NetworkService) SetScheme(scheme string) {
	var _arg0 *C.GNetworkService // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GNetworkService)(unsafe.Pointer(coreglib.InternObject(srv).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_network_service_set_scheme(_arg0, _arg1)
	runtime.KeepAlive(srv)
	runtime.KeepAlive(scheme)
}

// NetworkServiceClass: instance of this type is always passed by reference.
type NetworkServiceClass struct {
	*networkServiceClass
}

// networkServiceClass is the struct that's finalized.
type networkServiceClass struct {
	native *C.GNetworkServiceClass
}
