// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
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
		{T: externglib.Type(C.g_proxy_get_type()), F: marshalProXY},
	})
}

// ProXYOverrider contains methods that are overridable. This
// interface is a subset of the interface ProXY.
type ProXYOverrider interface {
	// Connect: given @connection to communicate with a proxy (eg, a Connection
	// that is connected to the proxy server), this does the necessary handshake
	// to connect to @proxy_address, and if required, wraps the OStream to
	// handle proxy payload.
	Connect(connection IOStream, proxyAddress ProXYAddress, cancellable Cancellable) (IOStream, error)
	// ConnectAsync asynchronous version of g_proxy_connect().
	ConnectAsync()
	// ConnectFinish: see g_proxy_connect().
	ConnectFinish(result AsyncResult) (IOStream, error)
	// SupportsHostname: some proxy protocols expect to be passed a hostname,
	// which they will resolve to an IP address themselves. Others, like SOCKS4,
	// do not allow this. This function will return false if @proxy is
	// implementing such a protocol. When false is returned, the caller should
	// resolve the destination hostname first, and then pass a Address
	// containing the stringified IP address to g_proxy_connect() or
	// g_proxy_connect_async().
	SupportsHostname() bool
}

// ProXY: a #GProxy handles connecting to a remote host via a given type of
// proxy server. It is implemented by the 'gio-proxy' extension point. The
// extensions are named after their proxy protocol name. As an example, a SOCKS5
// proxy implementation can be retrieved with the name 'socks5' using the
// function g_io_extension_point_get_extension_by_name().
type ProXY interface {
	gextras.Objector
	ProXYOverrider
}

// proXY implements the ProXY interface.
type proXY struct {
	gextras.Objector
}

var _ ProXY = (*proXY)(nil)

// WrapProXY wraps a GObject to a type that implements interface
// ProXY. It is primarily used internally.
func WrapProXY(obj *externglib.Object) ProXY {
	return ProXY{
		Objector: obj,
	}
}

func marshalProXY(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapProXY(obj), nil
}

// Connect: given @connection to communicate with a proxy (eg, a Connection
// that is connected to the proxy server), this does the necessary handshake
// to connect to @proxy_address, and if required, wraps the OStream to
// handle proxy payload.
func (p proXY) Connect(connection IOStream, proxyAddress ProXYAddress, cancellable Cancellable) (IOStream, error) {
	var _arg0 *C.GProxy
	var _arg1 *C.GIOStream
	var _arg2 *C.GProxyAddress
	var _arg3 *C.GCancellable

	_arg0 = (*C.GProxy)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GIOStream)(unsafe.Pointer(connection.Native()))
	_arg2 = (*C.GProxyAddress)(unsafe.Pointer(proxyAddress.Native()))
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cret *C.GIOStream
	var _cerr *C.GError

	cret = C.g_proxy_connect(_arg0, _arg1, _arg2, _arg3, _cerr)

	var _ioStream IOStream
	var _goerr error

	_ioStream = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(IOStream)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _ioStream, _goerr
}

// ConnectAsync asynchronous version of g_proxy_connect().
func (p proXY) ConnectAsync() {
	var _arg0 *C.GProxy

	_arg0 = (*C.GProxy)(unsafe.Pointer(p.Native()))

	C.g_proxy_connect_async(_arg0)
}

// ConnectFinish: see g_proxy_connect().
func (p proXY) ConnectFinish(result AsyncResult) (IOStream, error) {
	var _arg0 *C.GProxy
	var _arg1 *C.GAsyncResult

	_arg0 = (*C.GProxy)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var _cret *C.GIOStream
	var _cerr *C.GError

	cret = C.g_proxy_connect_finish(_arg0, _arg1, _cerr)

	var _ioStream IOStream
	var _goerr error

	_ioStream = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(IOStream)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _ioStream, _goerr
}

// SupportsHostname: some proxy protocols expect to be passed a hostname,
// which they will resolve to an IP address themselves. Others, like SOCKS4,
// do not allow this. This function will return false if @proxy is
// implementing such a protocol. When false is returned, the caller should
// resolve the destination hostname first, and then pass a Address
// containing the stringified IP address to g_proxy_connect() or
// g_proxy_connect_async().
func (p proXY) SupportsHostname() bool {
	var _arg0 *C.GProxy

	_arg0 = (*C.GProxy)(unsafe.Pointer(p.Native()))

	var _cret C.gboolean

	cret = C.g_proxy_supports_hostname(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}