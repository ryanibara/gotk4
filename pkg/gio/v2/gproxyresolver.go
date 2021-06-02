// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
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
		{T: externglib.Type(C.g_proxy_resolver_get_type()), F: marshalProxyResolver},
	})
}

// ProxyResolverGetDefault gets the default Resolver for the system.
func ProxyResolverGetDefault() ProxyResolver {
	ret := C.g_proxy_resolver_get_default()

	var ret0 ProxyResolver

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(ProxyResolver)

	return ret0
}

// ProxyResolverOverrider contains methods that are overridable. This
// interface is a subset of the interface ProxyResolver.
type ProxyResolverOverrider interface {
	// IsSupported checks if @resolver can be used on this system. (This is used
	// internally; g_proxy_resolver_get_default() will only return a proxy
	// resolver that returns true for this method.)
	IsSupported() bool
	// Lookup looks into the system proxy configuration to determine what proxy,
	// if any, to use to connect to @uri. The returned proxy URIs are of the
	// form `<protocol>://[user[:password]@]host:port` or `direct://`, where
	// <protocol> could be http, rtsp, socks or other proxying protocol.
	//
	// If you don't know what network protocol is being used on the socket, you
	// should use `none` as the URI protocol. In this case, the resolver might
	// still return a generic proxy type (such as SOCKS), but would not return
	// protocol-specific proxy types (such as http).
	//
	// `direct://` is used when no proxy is needed. Direct connection should not
	// be attempted unless it is part of the returned array of proxies.
	Lookup(uri string, cancellable Cancellable) (utf8s []string, err error)
	// LookupAsync asynchronous lookup of proxy. See g_proxy_resolver_lookup()
	// for more details.
	LookupAsync(uri string, cancellable Cancellable, callback AsyncReadyCallback)
	// LookupFinish: call this function to obtain the array of proxy URIs when
	// g_proxy_resolver_lookup_async() is complete. See
	// g_proxy_resolver_lookup() for more details.
	LookupFinish(result AsyncResult) (utf8s []string, err error)
}

// ProxyResolver provides synchronous and asynchronous network proxy resolution.
// Resolver is used within Client through the method
// g_socket_connectable_proxy_enumerate().
//
// Implementations of Resolver based on libproxy and GNOME settings can be found
// in glib-networking. GIO comes with an implementation for use inside Flatpak
// portals.
type ProxyResolver interface {
	gextras.Objector
	ProxyResolverOverrider
}

// proxyResolver implements the ProxyResolver interface.
type proxyResolver struct {
	gextras.Objector
}

var _ ProxyResolver = (*proxyResolver)(nil)

// WrapProxyResolver wraps a GObject to a type that implements interface
// ProxyResolver. It is primarily used internally.
func WrapProxyResolver(obj *externglib.Object) ProxyResolver {
	return ProxyResolver{
		Objector: obj,
	}
}

func marshalProxyResolver(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapProxyResolver(obj), nil
}

// IsSupported checks if @resolver can be used on this system. (This is used
// internally; g_proxy_resolver_get_default() will only return a proxy
// resolver that returns true for this method.)
func (r proxyResolver) IsSupported() bool {
	var arg0 *C.GProxyResolver

	arg0 = (*C.GProxyResolver)(r.Native())

	ret := C.g_proxy_resolver_is_supported(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Lookup looks into the system proxy configuration to determine what proxy,
// if any, to use to connect to @uri. The returned proxy URIs are of the
// form `<protocol>://[user[:password]@]host:port` or `direct://`, where
// <protocol> could be http, rtsp, socks or other proxying protocol.
//
// If you don't know what network protocol is being used on the socket, you
// should use `none` as the URI protocol. In this case, the resolver might
// still return a generic proxy type (such as SOCKS), but would not return
// protocol-specific proxy types (such as http).
//
// `direct://` is used when no proxy is needed. Direct connection should not
// be attempted unless it is part of the returned array of proxies.
func (r proxyResolver) Lookup(uri string, cancellable Cancellable) (utf8s []string, err error) {
	var arg0 *C.GProxyResolver
	var arg1 *C.gchar
	var arg2 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GProxyResolver)(r.Native())
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_proxy_resolver_lookup(arg0, arg1, arg2, &gError)

	var ret0 []string
	var goError error

	{
		var length uint
		for p := unsafe.Pointer(ret); *p != 0; p = unsafe.Pointer(uintptr(p) + 1) {
			length++
		}

		ret0 = make([]string, length)
		for i := 0; i < length; i++ {
			src := (*C.gchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ret)) + i))
			ret0[i] = C.GoString(src)
			C.free(unsafe.Pointer(src))
		}
	}

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// LookupAsync asynchronous lookup of proxy. See g_proxy_resolver_lookup()
// for more details.
func (r proxyResolver) LookupAsync(uri string, cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GProxyResolver
	var arg1 *C.gchar
	var arg2 *C.GCancellable
	var arg3 C.GAsyncReadyCallback
	var arg4 C.gpointer

	arg0 = (*C.GProxyResolver)(r.Native())
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GCancellable)(cancellable.Native())
	arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	arg4 = C.gpointer(box.Assign(callback))

	C.g_proxy_resolver_lookup_async(arg0, arg1, arg2, arg3, arg4)
}

// LookupFinish: call this function to obtain the array of proxy URIs when
// g_proxy_resolver_lookup_async() is complete. See
// g_proxy_resolver_lookup() for more details.
func (r proxyResolver) LookupFinish(result AsyncResult) (utf8s []string, err error) {
	var arg0 *C.GProxyResolver
	var arg1 *C.GAsyncResult
	var gError *C.GError

	arg0 = (*C.GProxyResolver)(r.Native())
	arg1 = (*C.GAsyncResult)(result.Native())

	ret := C.g_proxy_resolver_lookup_finish(arg0, arg1, &gError)

	var ret0 []string
	var goError error

	{
		var length uint
		for p := unsafe.Pointer(ret); *p != 0; p = unsafe.Pointer(uintptr(p) + 1) {
			length++
		}

		ret0 = make([]string, length)
		for i := 0; i < length; i++ {
			src := (*C.gchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ret)) + i))
			ret0[i] = C.GoString(src)
			C.free(unsafe.Pointer(src))
		}
	}

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}
