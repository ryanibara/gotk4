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
		{T: externglib.Type(C.g_simple_proxy_resolver_get_type()), F: marshalSimpleProxyResolverer},
	})
}

// SimpleProxyResolverer describes SimpleProxyResolver's methods.
type SimpleProxyResolverer interface {
	// SetDefaultProxy sets the default proxy on @resolver, to be used for any
	// URIs that don't match ProxyResolver:ignore-hosts or a proxy set via
	// g_simple_proxy_resolver_set_uri_proxy().
	SetDefaultProxy(defaultProxy string)
	// SetURIProxy adds a URI-scheme-specific proxy to @resolver; URIs whose
	// scheme matches @uri_scheme (and which don't match
	// ProxyResolver:ignore-hosts) will be proxied via @proxy.
	SetURIProxy(uriScheme string, proxy string)
}

// SimpleProxyResolver is a simple Resolver implementation that handles a single
// default proxy, multiple URI-scheme-specific proxies, and a list of hosts that
// proxies should not be used for.
//
// ProxyResolver is never the default proxy resolver, but it can be used as the
// base class for another proxy resolver implementation, or it can be created
// and used manually, such as with g_socket_client_set_proxy_resolver().
type SimpleProxyResolver struct {
	*externglib.Object

	ProxyResolver
}

var (
	_ SimpleProxyResolverer = (*SimpleProxyResolver)(nil)
	_ gextras.Nativer       = (*SimpleProxyResolver)(nil)
)

func wrapSimpleProxyResolver(obj *externglib.Object) SimpleProxyResolverer {
	return &SimpleProxyResolver{
		Object: obj,
		ProxyResolver: ProxyResolver{
			Object: obj,
		},
	}
}

func marshalSimpleProxyResolverer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSimpleProxyResolver(obj), nil
}

// SetDefaultProxy sets the default proxy on @resolver, to be used for any URIs
// that don't match ProxyResolver:ignore-hosts or a proxy set via
// g_simple_proxy_resolver_set_uri_proxy().
//
// If @default_proxy starts with "socks://", ProxyResolver will treat it as
// referring to all three of the socks5, socks4a, and socks4 proxy types.
func (resolver *SimpleProxyResolver) SetDefaultProxy(defaultProxy string) {
	var _arg0 *C.GSimpleProxyResolver // out
	var _arg1 *C.gchar                // out

	_arg0 = (*C.GSimpleProxyResolver)(unsafe.Pointer(resolver.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(defaultProxy)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_simple_proxy_resolver_set_default_proxy(_arg0, _arg1)
}

// SetURIProxy adds a URI-scheme-specific proxy to @resolver; URIs whose scheme
// matches @uri_scheme (and which don't match ProxyResolver:ignore-hosts) will
// be proxied via @proxy.
//
// As with ProxyResolver:default-proxy, if @proxy starts with "socks://",
// ProxyResolver will treat it as referring to all three of the socks5, socks4a,
// and socks4 proxy types.
func (resolver *SimpleProxyResolver) SetURIProxy(uriScheme string, proxy string) {
	var _arg0 *C.GSimpleProxyResolver // out
	var _arg1 *C.gchar                // out
	var _arg2 *C.gchar                // out

	_arg0 = (*C.GSimpleProxyResolver)(unsafe.Pointer(resolver.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uriScheme)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(proxy)))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_simple_proxy_resolver_set_uri_proxy(_arg0, _arg1, _arg2)
}
