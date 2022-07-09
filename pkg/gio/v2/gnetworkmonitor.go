// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gio2_NetworkMonitorInterface_can_reach(void*, void*, void*, GError**);
// extern gboolean _gotk4_gio2_NetworkMonitorInterface_can_reach_finish(void*, void*, GError**);
// extern void _gotk4_gio2_AsyncReadyCallback(void*, void*, gpointer);
// extern void _gotk4_gio2_NetworkMonitorInterface_network_changed(void*, gboolean);
// extern void _gotk4_gio2_NetworkMonitor_ConnectNetworkChanged(gpointer, gboolean, guintptr);
import "C"

// GTypeNetworkMonitor returns the GType for the type NetworkMonitor.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeNetworkMonitor() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "NetworkMonitor").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalNetworkMonitor)
	return gtype
}

// NETWORK_MONITOR_EXTENSION_POINT_NAME: extension point for network status
// monitoring functionality. See [Extending GIO][extending-gio].
const NETWORK_MONITOR_EXTENSION_POINT_NAME = "gio-network-monitor"

// NetworkMonitor provides an easy-to-use cross-platform API for monitoring
// network connectivity. On Linux, the available implementations are based on
// the kernel's netlink interface and on NetworkManager.
//
// There is also an implementation for use inside Flatpak sandboxes.
//
// NetworkMonitor wraps an interface. This means the user can get the
// underlying type by calling Cast().
type NetworkMonitor struct {
	_ [0]func() // equal guard
	Initable
}

var ()

// NetworkMonitorrer describes NetworkMonitor's interface methods.
type NetworkMonitorrer interface {
	coreglib.Objector

	// CanReach attempts to determine whether or not the host pointed to by
	// connectable can be reached, without actually trying to connect to it.
	CanReach(ctx context.Context, connectable SocketConnectabler) error
	// CanReachAsync: asynchronously attempts to determine whether or not the
	// host pointed to by connectable can be reached, without actually trying to
	// connect to it.
	CanReachAsync(ctx context.Context, connectable SocketConnectabler, callback AsyncReadyCallback)
	// CanReachFinish finishes an async network connectivity test.
	CanReachFinish(result AsyncResulter) error
	// NetworkAvailable checks if the network is available.
	NetworkAvailable() bool
	// NetworkMetered checks if the network is metered.
	NetworkMetered() bool

	// Network-changed is emitted when the network configuration changes.
	ConnectNetworkChanged(func(networkAvailable bool)) coreglib.SignalHandle
}

var _ NetworkMonitorrer = (*NetworkMonitor)(nil)

func wrapNetworkMonitor(obj *coreglib.Object) *NetworkMonitor {
	return &NetworkMonitor{
		Initable: Initable{
			Object: obj,
		},
	}
}

func marshalNetworkMonitor(p uintptr) (interface{}, error) {
	return wrapNetworkMonitor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_NetworkMonitor_ConnectNetworkChanged
func _gotk4_gio2_NetworkMonitor_ConnectNetworkChanged(arg0 C.gpointer, arg1 C.gboolean, arg2 C.guintptr) {
	var f func(networkAvailable bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(networkAvailable bool))
	}

	var _networkAvailable bool // out

	if arg1 != 0 {
		_networkAvailable = true
	}

	f(_networkAvailable)
}

// ConnectNetworkChanged is emitted when the network configuration changes.
func (monitor *NetworkMonitor) ConnectNetworkChanged(f func(networkAvailable bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(monitor, "network-changed", false, unsafe.Pointer(C._gotk4_gio2_NetworkMonitor_ConnectNetworkChanged), f)
}

// CanReach attempts to determine whether or not the host pointed to by
// connectable can be reached, without actually trying to connect to it.
//
// This may return TRUE even when Monitor:network-available is FALSE, if, for
// example, monitor can determine that connectable refers to a host on a local
// network.
//
// If monitor believes that an attempt to connect to connectable will succeed,
// it will return TRUE. Otherwise, it will return FALSE and set error to an
// appropriate error (such as G_IO_ERROR_HOST_UNREACHABLE).
//
// Note that although this does not attempt to connect to connectable, it may
// still block for a brief period of time (eg, trying to do multicast DNS on the
// local network), so if you do not want to block, you should use
// g_network_monitor_can_reach_async().
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - connectable: Connectable.
//
func (monitor *NetworkMonitor) CanReach(ctx context.Context, connectable SocketConnectabler) error {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(connectable).Native()))

	runtime.KeepAlive(monitor)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connectable)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// CanReachAsync: asynchronously attempts to determine whether or not the host
// pointed to by connectable can be reached, without actually trying to connect
// to it.
//
// For more details, see g_network_monitor_can_reach().
//
// When the operation is finished, callback will be called. You can then call
// g_network_monitor_can_reach_finish() to get the result of the operation.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - connectable: Connectable.
//    - callback (optional) to call when the request is satisfied.
//
func (monitor *NetworkMonitor) CanReachAsync(ctx context.Context, connectable SocketConnectabler, callback AsyncReadyCallback) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(connectable).Native()))
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[3])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[4] = C.gpointer(gbox.AssignOnce(callback))
	}

	runtime.KeepAlive(monitor)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connectable)
	runtime.KeepAlive(callback)
}

// CanReachFinish finishes an async network connectivity test. See
// g_network_monitor_can_reach_async().
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (monitor *NetworkMonitor) CanReachFinish(result AsyncResulter) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	runtime.KeepAlive(monitor)
	runtime.KeepAlive(result)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// NetworkAvailable checks if the network is available. "Available" here means
// that the system has a default route available for at least one of IPv4 or
// IPv6. It does not necessarily imply that the public Internet is reachable.
// See Monitor:network-available for more details.
//
// The function returns the following values:
//
//    - ok: whether the network is available.
//
func (monitor *NetworkMonitor) NetworkAvailable() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// NetworkMetered checks if the network is metered. See Monitor:network-metered
// for more details.
//
// The function returns the following values:
//
//    - ok: whether the connection is metered.
//
func (monitor *NetworkMonitor) NetworkMetered() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// NetworkMonitorGetDefault gets the default Monitor for the system.
//
// The function returns the following values:
//
//    - networkMonitor which will be a dummy object if no network monitor is
//      available.
//
func NetworkMonitorGetDefault() *NetworkMonitor {
	_gret := girepository.MustFind("Gio", "get_default").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _networkMonitor *NetworkMonitor // out

	_networkMonitor = wrapNetworkMonitor(coreglib.Take(unsafe.Pointer(_cret)))

	return _networkMonitor
}
