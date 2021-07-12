// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_proxy_get_type()), F: marshalProxier},
	})
}

// ProxyOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ProxyOverrider interface {
	// ConnectProxier: given @connection to communicate with a proxy (eg, a
	// Connection that is connected to the proxy server), this does the
	// necessary handshake to connect to @proxy_address, and if required, wraps
	// the OStream to handle proxy payload.
	ConnectProxier(connection IOStreamer, proxyAddress ProxyAddresser, cancellable Cancellabler) (*IOStream, error)
	// ConnectAsync asynchronous version of g_proxy_connect().
	ConnectAsync(connection IOStreamer, proxyAddress ProxyAddresser, cancellable Cancellabler, callback AsyncReadyCallback)
	// ConnectFinish: see g_proxy_connect().
	ConnectFinish(result AsyncResulter) (*IOStream, error)
	// SupportsHostname: some proxy protocols expect to be passed a hostname,
	// which they will resolve to an IP address themselves. Others, like SOCKS4,
	// do not allow this. This function will return false if @proxy is
	// implementing such a protocol. When false is returned, the caller should
	// resolve the destination hostname first, and then pass a Address
	// containing the stringified IP address to g_proxy_connect() or
	// g_proxy_connect_async().
	SupportsHostname() bool
}

// Proxier describes Proxy's methods.
type Proxier interface {
	// ConnectProxier: given @connection to communicate with a proxy (eg, a
	// Connection that is connected to the proxy server), this does the
	// necessary handshake to connect to @proxy_address, and if required, wraps
	// the OStream to handle proxy payload.
	ConnectProxier(connection IOStreamer, proxyAddress ProxyAddresser, cancellable Cancellabler) (*IOStream, error)
	// ConnectAsync asynchronous version of g_proxy_connect().
	ConnectAsync(connection IOStreamer, proxyAddress ProxyAddresser, cancellable Cancellabler, callback AsyncReadyCallback)
	// ConnectFinish: see g_proxy_connect().
	ConnectFinish(result AsyncResulter) (*IOStream, error)
	// SupportsHostname: some proxy protocols expect to be passed a hostname,
	// which they will resolve to an IP address themselves.
	SupportsHostname() bool
}

// Proxy handles connecting to a remote host via a given type of proxy server.
// It is implemented by the 'gio-proxy' extension point. The extensions are
// named after their proxy protocol name. As an example, a SOCKS5 proxy
// implementation can be retrieved with the name 'socks5' using the function
// g_io_extension_point_get_extension_by_name().
type Proxy struct {
	*externglib.Object
}

var (
	_ Proxier         = (*Proxy)(nil)
	_ gextras.Nativer = (*Proxy)(nil)
)

func wrapProxy(obj *externglib.Object) *Proxy {
	return &Proxy{
		Object: obj,
	}
}

func marshalProxier(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapProxy(obj), nil
}

// ConnectProxier: given @connection to communicate with a proxy (eg, a
// Connection that is connected to the proxy server), this does the necessary
// handshake to connect to @proxy_address, and if required, wraps the OStream to
// handle proxy payload.
func (proxy *Proxy) ConnectProxier(connection IOStreamer, proxyAddress ProxyAddresser, cancellable Cancellabler) (*IOStream, error) {
	var _arg0 *C.GProxy        // out
	var _arg1 *C.GIOStream     // out
	var _arg2 *C.GProxyAddress // out
	var _arg3 *C.GCancellable  // out
	var _cret *C.GIOStream     // in
	var _cerr *C.GError        // in

	_arg0 = (*C.GProxy)(unsafe.Pointer(proxy.Native()))
	_arg1 = (*C.GIOStream)(unsafe.Pointer((connection).(gextras.Nativer).Native()))
	_arg2 = (*C.GProxyAddress)(unsafe.Pointer((proxyAddress).(gextras.Nativer).Native()))
	_arg3 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_proxy_connect(_arg0, _arg1, _arg2, _arg3, &_cerr)

	var _ioStream *IOStream // out
	var _goerr error        // out

	_ioStream = wrapIOStream(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _ioStream, _goerr
}

// ConnectAsync asynchronous version of g_proxy_connect().
func (proxy *Proxy) ConnectAsync(connection IOStreamer, proxyAddress ProxyAddresser, cancellable Cancellabler, callback AsyncReadyCallback) {
	var _arg0 *C.GProxy             // out
	var _arg1 *C.GIOStream          // out
	var _arg2 *C.GProxyAddress      // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GProxy)(unsafe.Pointer(proxy.Native()))
	_arg1 = (*C.GIOStream)(unsafe.Pointer((connection).(gextras.Nativer).Native()))
	_arg2 = (*C.GProxyAddress)(unsafe.Pointer((proxyAddress).(gextras.Nativer).Native()))
	_arg3 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(gbox.Assign(callback))

	C.g_proxy_connect_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// ConnectFinish: see g_proxy_connect().
func (proxy *Proxy) ConnectFinish(result AsyncResulter) (*IOStream, error) {
	var _arg0 *C.GProxy       // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GIOStream    // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GProxy)(unsafe.Pointer(proxy.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.g_proxy_connect_finish(_arg0, _arg1, &_cerr)

	var _ioStream *IOStream // out
	var _goerr error        // out

	_ioStream = wrapIOStream(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _ioStream, _goerr
}

// SupportsHostname: some proxy protocols expect to be passed a hostname, which
// they will resolve to an IP address themselves. Others, like SOCKS4, do not
// allow this. This function will return false if @proxy is implementing such a
// protocol. When false is returned, the caller should resolve the destination
// hostname first, and then pass a Address containing the stringified IP address
// to g_proxy_connect() or g_proxy_connect_async().
func (proxy *Proxy) SupportsHostname() bool {
	var _arg0 *C.GProxy  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GProxy)(unsafe.Pointer(proxy.Native()))

	_cret = C.g_proxy_supports_hostname(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ProxyGetDefaultForProtocol: find the `gio-proxy` extension point for a proxy
// implementation that supports the specified protocol.
func ProxyGetDefaultForProtocol(protocol string) *Proxy {
	var _arg1 *C.gchar  // out
	var _cret *C.GProxy // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(protocol)))

	_cret = C.g_proxy_get_default_for_protocol(_arg1)

	var _proxy *Proxy // out

	_proxy = wrapProxy(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _proxy
}
