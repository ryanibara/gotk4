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
		{T: externglib.Type(C.g_network_monitor_get_type()), F: marshalNetworkMonitor},
	})
}

// NetworkMonitorGetDefault gets the default Monitor for the system.
func NetworkMonitorGetDefault() NetworkMonitor {
	ret := C.g_network_monitor_get_default()

	var ret0 NetworkMonitor

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(NetworkMonitor)

	return ret0
}

// NetworkMonitorOverrider contains methods that are overridable. This
// interface is a subset of the interface NetworkMonitor.
type NetworkMonitorOverrider interface {
	// CanReach attempts to determine whether or not the host pointed to by
	// @connectable can be reached, without actually trying to connect to it.
	//
	// This may return true even when Monitor:network-available is false, if,
	// for example, @monitor can determine that @connectable refers to a host on
	// a local network.
	//
	// If @monitor believes that an attempt to connect to @connectable will
	// succeed, it will return true. Otherwise, it will return false and set
	// @error to an appropriate error (such as G_IO_ERROR_HOST_UNREACHABLE).
	//
	// Note that although this does not attempt to connect to @connectable, it
	// may still block for a brief period of time (eg, trying to do multicast
	// DNS on the local network), so if you do not want to block, you should use
	// g_network_monitor_can_reach_async().
	CanReach(connectable SocketConnectable, cancellable Cancellable) error
	// CanReachAsync: asynchronously attempts to determine whether or not the
	// host pointed to by @connectable can be reached, without actually trying
	// to connect to it.
	//
	// For more details, see g_network_monitor_can_reach().
	//
	// When the operation is finished, @callback will be called. You can then
	// call g_network_monitor_can_reach_finish() to get the result of the
	// operation.
	CanReachAsync(connectable SocketConnectable, cancellable Cancellable, callback AsyncReadyCallback)
	// CanReachFinish finishes an async network connectivity test. See
	// g_network_monitor_can_reach_async().
	CanReachFinish(result AsyncResult) error

	NetworkChanged(networkAvailable bool)
}

// NetworkMonitor provides an easy-to-use cross-platform API for monitoring
// network connectivity. On Linux, the available implementations are based on
// the kernel's netlink interface and on NetworkManager.
//
// There is also an implementation for use inside Flatpak sandboxes.
type NetworkMonitor interface {
	Initable
	NetworkMonitorOverrider

	// Connectivity gets a more detailed networking state than
	// g_network_monitor_get_network_available().
	//
	// If Monitor:network-available is false, then the connectivity state will
	// be G_NETWORK_CONNECTIVITY_LOCAL.
	//
	// If Monitor:network-available is true, then the connectivity state will be
	// G_NETWORK_CONNECTIVITY_FULL (if there is full Internet connectivity),
	// G_NETWORK_CONNECTIVITY_LIMITED (if the host has a default route, but
	// appears to be unable to actually reach the full Internet), or
	// G_NETWORK_CONNECTIVITY_PORTAL (if the host is trapped behind a "captive
	// portal" that requires some sort of login or acknowledgement before
	// allowing full Internet access).
	//
	// Note that in the case of G_NETWORK_CONNECTIVITY_LIMITED and
	// G_NETWORK_CONNECTIVITY_PORTAL, it is possible that some sites are
	// reachable but others are not. In this case, applications can attempt to
	// connect to remote servers, but should gracefully fall back to their
	// "offline" behavior if the connection attempt fails.
	Connectivity() NetworkConnectivity
	// NetworkAvailable checks if the network is available. "Available" here
	// means that the system has a default route available for at least one of
	// IPv4 or IPv6. It does not necessarily imply that the public Internet is
	// reachable. See Monitor:network-available for more details.
	NetworkAvailable() bool
	// NetworkMetered checks if the network is metered. See
	// Monitor:network-metered for more details.
	NetworkMetered() bool
}

// networkMonitor implements the NetworkMonitor interface.
type networkMonitor struct {
	Initable
}

var _ NetworkMonitor = (*networkMonitor)(nil)

// WrapNetworkMonitor wraps a GObject to a type that implements interface
// NetworkMonitor. It is primarily used internally.
func WrapNetworkMonitor(obj *externglib.Object) NetworkMonitor {
	return NetworkMonitor{
		Initable: WrapInitable(obj),
	}
}

func marshalNetworkMonitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNetworkMonitor(obj), nil
}

// CanReach attempts to determine whether or not the host pointed to by
// @connectable can be reached, without actually trying to connect to it.
//
// This may return true even when Monitor:network-available is false, if,
// for example, @monitor can determine that @connectable refers to a host on
// a local network.
//
// If @monitor believes that an attempt to connect to @connectable will
// succeed, it will return true. Otherwise, it will return false and set
// @error to an appropriate error (such as G_IO_ERROR_HOST_UNREACHABLE).
//
// Note that although this does not attempt to connect to @connectable, it
// may still block for a brief period of time (eg, trying to do multicast
// DNS on the local network), so if you do not want to block, you should use
// g_network_monitor_can_reach_async().
func (m networkMonitor) CanReach(connectable SocketConnectable, cancellable Cancellable) error {
	var arg0 *C.GNetworkMonitor
	var arg1 *C.GSocketConnectable
	var arg2 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GNetworkMonitor)(m.Native())
	arg1 = (*C.GSocketConnectable)(connectable.Native())
	arg2 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_network_monitor_can_reach(arg0, arg1, arg2, &gError)

	var goError error

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return goError
}

// CanReachAsync: asynchronously attempts to determine whether or not the
// host pointed to by @connectable can be reached, without actually trying
// to connect to it.
//
// For more details, see g_network_monitor_can_reach().
//
// When the operation is finished, @callback will be called. You can then
// call g_network_monitor_can_reach_finish() to get the result of the
// operation.
func (m networkMonitor) CanReachAsync(connectable SocketConnectable, cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GNetworkMonitor
	var arg1 *C.GSocketConnectable
	var arg2 *C.GCancellable
	var arg3 C.GAsyncReadyCallback
	var arg4 C.gpointer

	arg0 = (*C.GNetworkMonitor)(m.Native())
	arg1 = (*C.GSocketConnectable)(connectable.Native())
	arg2 = (*C.GCancellable)(cancellable.Native())
	arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	arg4 = C.gpointer(box.Assign(callback))

	C.g_network_monitor_can_reach_async(arg0, arg1, arg2, arg3, arg4)
}

// CanReachFinish finishes an async network connectivity test. See
// g_network_monitor_can_reach_async().
func (m networkMonitor) CanReachFinish(result AsyncResult) error {
	var arg0 *C.GNetworkMonitor
	var arg1 *C.GAsyncResult
	var gError *C.GError

	arg0 = (*C.GNetworkMonitor)(m.Native())
	arg1 = (*C.GAsyncResult)(result.Native())

	ret := C.g_network_monitor_can_reach_finish(arg0, arg1, &gError)

	var goError error

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return goError
}

// Connectivity gets a more detailed networking state than
// g_network_monitor_get_network_available().
//
// If Monitor:network-available is false, then the connectivity state will
// be G_NETWORK_CONNECTIVITY_LOCAL.
//
// If Monitor:network-available is true, then the connectivity state will be
// G_NETWORK_CONNECTIVITY_FULL (if there is full Internet connectivity),
// G_NETWORK_CONNECTIVITY_LIMITED (if the host has a default route, but
// appears to be unable to actually reach the full Internet), or
// G_NETWORK_CONNECTIVITY_PORTAL (if the host is trapped behind a "captive
// portal" that requires some sort of login or acknowledgement before
// allowing full Internet access).
//
// Note that in the case of G_NETWORK_CONNECTIVITY_LIMITED and
// G_NETWORK_CONNECTIVITY_PORTAL, it is possible that some sites are
// reachable but others are not. In this case, applications can attempt to
// connect to remote servers, but should gracefully fall back to their
// "offline" behavior if the connection attempt fails.
func (m networkMonitor) Connectivity() NetworkConnectivity {
	var arg0 *C.GNetworkMonitor

	arg0 = (*C.GNetworkMonitor)(m.Native())

	ret := C.g_network_monitor_get_connectivity(arg0)

	var ret0 NetworkConnectivity

	ret0 = NetworkConnectivity(ret)

	return ret0
}

// NetworkAvailable checks if the network is available. "Available" here
// means that the system has a default route available for at least one of
// IPv4 or IPv6. It does not necessarily imply that the public Internet is
// reachable. See Monitor:network-available for more details.
func (m networkMonitor) NetworkAvailable() bool {
	var arg0 *C.GNetworkMonitor

	arg0 = (*C.GNetworkMonitor)(m.Native())

	ret := C.g_network_monitor_get_network_available(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != C.false

	return ret0
}

// NetworkMetered checks if the network is metered. See
// Monitor:network-metered for more details.
func (m networkMonitor) NetworkMetered() bool {
	var arg0 *C.GNetworkMonitor

	arg0 = (*C.GNetworkMonitor)(m.Native())

	ret := C.g_network_monitor_get_network_metered(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != C.false

	return ret0
}
