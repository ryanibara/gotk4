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
		{T: externglib.Type(C.g_threaded_socket_service_get_type()), F: marshalThreadedSocketServicer},
	})
}

// ThreadedSocketServiceOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ThreadedSocketServiceOverrider interface {
	Run(connection SocketConnectioner, sourceObject gextras.Objector) bool
}

// ThreadedSocketServicer describes ThreadedSocketService's methods.
type ThreadedSocketServicer interface {
	privateThreadedSocketService()
}

// ThreadedSocketService is a simple subclass of Service that handles incoming
// connections by creating a worker thread and dispatching the connection to it
// by emitting the SocketService::run signal in the new thread.
//
// The signal handler may perform blocking IO and need not return until the
// connection is closed.
//
// The service is implemented using a thread pool, so there is a limited amount
// of threads available to serve incoming requests. The service automatically
// stops the Service from accepting new connections when all threads are busy.
//
// As with Service, you may connect to SocketService::run, or subclass and
// override the default handler.
type ThreadedSocketService struct {
	SocketService
}

var (
	_ ThreadedSocketServicer = (*ThreadedSocketService)(nil)
	_ gextras.Nativer        = (*ThreadedSocketService)(nil)
)

func wrapThreadedSocketService(obj *externglib.Object) ThreadedSocketServicer {
	return &ThreadedSocketService{
		SocketService: SocketService{
			SocketListener: SocketListener{
				Object: obj,
			},
		},
	}
}

func marshalThreadedSocketServicer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapThreadedSocketService(obj), nil
}

// NewThreadedSocketService creates a new SocketService with no listeners.
// Listeners must be added with one of the Listener "add" methods.
func NewThreadedSocketService(maxThreads int) *ThreadedSocketService {
	var _arg1 C.int             // out
	var _cret *C.GSocketService // in

	_arg1 = C.int(maxThreads)

	_cret = C.g_threaded_socket_service_new(_arg1)

	var _threadedSocketService *ThreadedSocketService // out

	_threadedSocketService = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*ThreadedSocketService)

	return _threadedSocketService
}

func (*ThreadedSocketService) privateThreadedSocketService() {}
