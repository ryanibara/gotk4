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
	GTypeSimpleProxyResolver = coreglib.Type(C.g_simple_proxy_resolver_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSimpleProxyResolver, F: marshalSimpleProxyResolver},
	})
}

// SimpleProxyResolverOverrides contains methods that are overridable.
type SimpleProxyResolverOverrides struct {
}

func defaultSimpleProxyResolverOverrides(v *SimpleProxyResolver) SimpleProxyResolverOverrides {
	return SimpleProxyResolverOverrides{}
}

// SimpleProxyResolver is a simple Resolver implementation that handles a single
// default proxy, multiple URI-scheme-specific proxies, and a list of hosts that
// proxies should not be used for.
//
// ProxyResolver is never the default proxy resolver, but it can be used as the
// base class for another proxy resolver implementation, or it can be created
// and used manually, such as with g_socket_client_set_proxy_resolver().
type SimpleProxyResolver struct {
	_ [0]func() // equal guard
	*coreglib.Object

	ProxyResolver
}

var (
	_ coreglib.Objector = (*SimpleProxyResolver)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*SimpleProxyResolver, *SimpleProxyResolverClass, SimpleProxyResolverOverrides](
		GTypeSimpleProxyResolver,
		initSimpleProxyResolverClass,
		wrapSimpleProxyResolver,
		defaultSimpleProxyResolverOverrides,
	)
}

func initSimpleProxyResolverClass(gclass unsafe.Pointer, overrides SimpleProxyResolverOverrides, classInitFunc func(*SimpleProxyResolverClass)) {
	if classInitFunc != nil {
		class := (*SimpleProxyResolverClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSimpleProxyResolver(obj *coreglib.Object) *SimpleProxyResolver {
	return &SimpleProxyResolver{
		Object: obj,
		ProxyResolver: ProxyResolver{
			Object: obj,
		},
	}
}

func marshalSimpleProxyResolver(p uintptr) (interface{}, error) {
	return wrapSimpleProxyResolver(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// SetDefaultProxy sets the default proxy on resolver, to be used for any URIs
// that don't match ProxyResolver:ignore-hosts or a proxy set via
// g_simple_proxy_resolver_set_uri_proxy().
//
// If default_proxy starts with "socks://", ProxyResolver will treat it as
// referring to all three of the socks5, socks4a, and socks4 proxy types.
//
// The function takes the following parameters:
//
//    - defaultProxy: default proxy to use.
//
func (resolver *SimpleProxyResolver) SetDefaultProxy(defaultProxy string) {
	var _arg0 *C.GSimpleProxyResolver // out
	var _arg1 *C.gchar                // out

	_arg0 = (*C.GSimpleProxyResolver)(unsafe.Pointer(coreglib.InternObject(resolver).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(defaultProxy)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_simple_proxy_resolver_set_default_proxy(_arg0, _arg1)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(defaultProxy)
}

// SetURIProxy adds a URI-scheme-specific proxy to resolver; URIs whose scheme
// matches uri_scheme (and which don't match ProxyResolver:ignore-hosts) will be
// proxied via proxy.
//
// As with ProxyResolver:default-proxy, if proxy starts with "socks://",
// ProxyResolver will treat it as referring to all three of the socks5, socks4a,
// and socks4 proxy types.
//
// The function takes the following parameters:
//
//    - uriScheme: URI scheme to add a proxy for.
//    - proxy to use for uri_scheme.
//
func (resolver *SimpleProxyResolver) SetURIProxy(uriScheme, proxy string) {
	var _arg0 *C.GSimpleProxyResolver // out
	var _arg1 *C.gchar                // out
	var _arg2 *C.gchar                // out

	_arg0 = (*C.GSimpleProxyResolver)(unsafe.Pointer(coreglib.InternObject(resolver).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uriScheme)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(proxy)))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_simple_proxy_resolver_set_uri_proxy(_arg0, _arg1, _arg2)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(uriScheme)
	runtime.KeepAlive(proxy)
}
