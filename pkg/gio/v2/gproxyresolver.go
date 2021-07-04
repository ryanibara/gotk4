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
		{T: externglib.Type(C.g_proxy_resolver_get_type()), F: marshalProxyResolver},
	})
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

	// IsSupported: call this function to obtain the array of proxy URIs when
	// g_proxy_resolver_lookup_async() is complete. See
	// g_proxy_resolver_lookup() for more details.
	IsSupported() bool
	// Lookup: call this function to obtain the array of proxy URIs when
	// g_proxy_resolver_lookup_async() is complete. See
	// g_proxy_resolver_lookup() for more details.
	Lookup(uri string, cancellable Cancellable) ([]string, error)
	// LookupAsync: call this function to obtain the array of proxy URIs when
	// g_proxy_resolver_lookup_async() is complete. See
	// g_proxy_resolver_lookup() for more details.
	LookupAsync(uri string, cancellable Cancellable, callback AsyncReadyCallback)
	// LookupFinish: call this function to obtain the array of proxy URIs when
	// g_proxy_resolver_lookup_async() is complete. See
	// g_proxy_resolver_lookup() for more details.
	LookupFinish(result AsyncResult) ([]string, error)
}

// proxyResolver implements the ProxyResolver interface.
type proxyResolver struct {
	gextras.Objector
}

var _ ProxyResolver = (*proxyResolver)(nil)

// WrapProxyResolver wraps a GObject to a type that implements
// interface ProxyResolver. It is primarily used internally.
func WrapProxyResolver(obj *externglib.Object) ProxyResolver {
	return proxyResolver{
		Objector: obj,
	}
}

func marshalProxyResolver(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapProxyResolver(obj), nil
}

func (r proxyResolver) IsSupported() bool {
	var _arg0 *C.GProxyResolver // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GProxyResolver)(unsafe.Pointer(r.Native()))

	_cret = C.g_proxy_resolver_is_supported(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (r proxyResolver) Lookup(uri string, cancellable Cancellable) ([]string, error) {
	var _arg0 *C.GProxyResolver // out
	var _arg1 *C.gchar          // out
	var _arg2 *C.GCancellable   // out
	var _cret **C.gchar
	var _cerr *C.GError // in

	_arg0 = (*C.GProxyResolver)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_proxy_resolver_lookup(_arg0, _arg1, _arg2, &_cerr)

	var _utf8s []string
	var _goerr error // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8s, _goerr
}

func (r proxyResolver) LookupAsync(uri string, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GProxyResolver     // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GProxyResolver)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_proxy_resolver_lookup_async(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (r proxyResolver) LookupFinish(result AsyncResult) ([]string, error) {
	var _arg0 *C.GProxyResolver // out
	var _arg1 *C.GAsyncResult   // out
	var _cret **C.gchar
	var _cerr *C.GError // in

	_arg0 = (*C.GProxyResolver)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_proxy_resolver_lookup_finish(_arg0, _arg1, &_cerr)

	var _utf8s []string
	var _goerr error // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8s, _goerr
}
