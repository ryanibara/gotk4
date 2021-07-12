// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_socket_service_get_type()), F: marshalSocketServicer},
	})
}

// SocketServiceOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SocketServiceOverrider interface {
	Incoming(connection SocketConnectioner, sourceObject gextras.Objector) bool
}

// SocketServicer describes SocketService's methods.
type SocketServicer interface {
	// IsActive: check whether the service is active or not.
	IsActive() bool
	// Start restarts the service, i.e.
	Start()
	// Stop stops the service, i.e.
	Stop()
}

// SocketService is an object that represents a service that is provided to the
// network or over local sockets. When a new connection is made to the service
// the Service::incoming signal is emitted.
//
// A Service is a subclass of Listener and you need to add the addresses you
// want to accept connections on with the Listener APIs.
//
// There are two options for implementing a network service based on Service.
// The first is to create the service using g_socket_service_new() and to
// connect to the Service::incoming signal. The second is to subclass Service
// and override the default signal handler implementation.
//
// In either case, the handler must immediately return, or else it will block
// additional incoming connections from being serviced. If you are interested in
// writing connection handlers that contain blocking code then see
// SocketService.
//
// The socket service runs on the main loop of the [thread-default
// context][g-main-context-push-thread-default-context] of the thread it is
// created in, and is not threadsafe in general. However, the calls to start and
// stop the service are thread-safe so these can be used from threads that
// handle incoming clients.
type SocketService struct {
	SocketListener
}

var (
	_ SocketServicer  = (*SocketService)(nil)
	_ gextras.Nativer = (*SocketService)(nil)
)

func wrapSocketService(obj *externglib.Object) *SocketService {
	return &SocketService{
		SocketListener: SocketListener{
			Object: obj,
		},
	}
}

func marshalSocketServicer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSocketService(obj), nil
}

// NewSocketService creates a new Service with no sockets to listen for. New
// listeners can be added with e.g. g_socket_listener_add_address() or
// g_socket_listener_add_inet_port().
//
// New services are created active, there is no need to call
// g_socket_service_start(), unless g_socket_service_stop() has been called
// before.
func NewSocketService() *SocketService {
	var _cret *C.GSocketService // in

	_cret = C.g_socket_service_new()

	var _socketService *SocketService // out

	_socketService = wrapSocketService(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _socketService
}

// IsActive: check whether the service is active or not. An active service will
// accept new clients that connect, while a non-active service will let
// connecting clients queue up until the service is started.
func (service *SocketService) IsActive() bool {
	var _arg0 *C.GSocketService // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GSocketService)(unsafe.Pointer(service.Native()))

	_cret = C.g_socket_service_is_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Start restarts the service, i.e. start accepting connections from the added
// sockets when the mainloop runs. This only needs to be called after the
// service has been stopped from g_socket_service_stop().
//
// This call is thread-safe, so it may be called from a thread handling an
// incoming client request.
func (service *SocketService) Start() {
	var _arg0 *C.GSocketService // out

	_arg0 = (*C.GSocketService)(unsafe.Pointer(service.Native()))

	C.g_socket_service_start(_arg0)
}

// Stop stops the service, i.e. stops accepting connections from the added
// sockets when the mainloop runs.
//
// This call is thread-safe, so it may be called from a thread handling an
// incoming client request.
//
// Note that this only stops accepting new connections; it does not close the
// listening sockets, and you can call g_socket_service_start() again later to
// begin listening again. To close the listening sockets, call
// g_socket_listener_close(). (This will happen automatically when the Service
// is finalized.)
//
// This must be called before calling g_socket_listener_close() as the socket
// service will start accepting connections immediately when a new socket is
// added.
func (service *SocketService) Stop() {
	var _arg0 *C.GSocketService // out

	_arg0 = (*C.GSocketService)(unsafe.Pointer(service.Native()))

	C.g_socket_service_stop(_arg0)
}
