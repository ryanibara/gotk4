// Code generated by girgen. DO NOT EDIT.

package gio

import (
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
		{T: externglib.Type(C.g_network_service_get_type()), F: marshalNetworkService},
	})
}

// NetworkService: like Address does with hostnames, Service provides an easy
// way to resolve a SRV record, and then attempt to connect to one of the hosts
// that implements that service, handling service priority/weighting, multiple
// IP addresses, and multiple address families.
//
// See Target for more information about SRV records, and see Connectable for an
// example of using the connectable interface.
type NetworkService interface {
	gextras.Objector
	SocketConnectable

	// Domain gets the domain that @srv serves. This might be either UTF-8 or
	// ASCII-encoded, depending on what @srv was created with.
	Domain() string
	// Protocol gets @srv's protocol name (eg, "tcp").
	Protocol() string
	// Scheme gets the URI scheme used to resolve proxies. By default, the
	// service name is used as scheme.
	Scheme() string
	// Service gets @srv's service name (eg, "ldap").
	Service() string
	// SetScheme set's the URI scheme used to resolve proxies. By default, the
	// service name is used as scheme.
	SetScheme(scheme string)
}

// networkService implements the NetworkService interface.
type networkService struct {
	gextras.Objector
	SocketConnectable
}

var _ NetworkService = (*networkService)(nil)

// WrapNetworkService wraps a GObject to the right type. It is
// primarily used internally.
func WrapNetworkService(obj *externglib.Object) NetworkService {
	return NetworkService{
		Objector:          obj,
		SocketConnectable: WrapSocketConnectable(obj),
	}
}

func marshalNetworkService(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNetworkService(obj), nil
}

// NewNetworkService constructs a class NetworkService.
func NewNetworkService(service string, protocol string, domain string) NetworkService {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar

	arg1 = (*C.gchar)(C.CString(service))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(protocol))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(arg3))

	ret := C.g_network_service_new(arg1, arg2, arg3)

	var ret0 NetworkService

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(NetworkService)

	return ret0
}

// Domain gets the domain that @srv serves. This might be either UTF-8 or
// ASCII-encoded, depending on what @srv was created with.
func (s networkService) Domain() string {
	var arg0 *C.GNetworkService

	arg0 = (*C.GNetworkService)(s.Native())

	ret := C.g_network_service_get_domain(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// Protocol gets @srv's protocol name (eg, "tcp").
func (s networkService) Protocol() string {
	var arg0 *C.GNetworkService

	arg0 = (*C.GNetworkService)(s.Native())

	ret := C.g_network_service_get_protocol(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// Scheme gets the URI scheme used to resolve proxies. By default, the
// service name is used as scheme.
func (s networkService) Scheme() string {
	var arg0 *C.GNetworkService

	arg0 = (*C.GNetworkService)(s.Native())

	ret := C.g_network_service_get_scheme(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// Service gets @srv's service name (eg, "ldap").
func (s networkService) Service() string {
	var arg0 *C.GNetworkService

	arg0 = (*C.GNetworkService)(s.Native())

	ret := C.g_network_service_get_service(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// SetScheme set's the URI scheme used to resolve proxies. By default, the
// service name is used as scheme.
func (s networkService) SetScheme(scheme string) {
	var arg0 *C.GNetworkService
	var arg1 *C.gchar

	arg0 = (*C.GNetworkService)(s.Native())
	arg1 = (*C.gchar)(C.CString(scheme))
	defer C.free(unsafe.Pointer(arg1))

	C.g_network_service_set_scheme(arg0, arg1)
}

type NetworkServicePrivate struct {
	native C.GNetworkServicePrivate
}

// WrapNetworkServicePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapNetworkServicePrivate(ptr unsafe.Pointer) *NetworkServicePrivate {
	if ptr == nil {
		return nil
	}

	return (*NetworkServicePrivate)(ptr)
}

func marshalNetworkServicePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapNetworkServicePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (n *NetworkServicePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&n.native)
}
