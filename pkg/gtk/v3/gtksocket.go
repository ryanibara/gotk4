// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern gboolean _gotk4_gtk3_SocketClass_plug_removed(GtkSocket*);
// extern gboolean _gotk4_gtk3_Socket_ConnectPlugRemoved(gpointer, guintptr);
// extern void _gotk4_gtk3_SocketClass_plug_added(GtkSocket*);
// extern void _gotk4_gtk3_Socket_ConnectPlugAdded(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeSocket = coreglib.Type(C.gtk_socket_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSocket, F: marshalSocket},
	})
}

// SocketOverrider contains methods that are overridable.
type SocketOverrider interface {
	PlugAdded()
	// The function returns the following values:
	//
	PlugRemoved() bool
}

// Socket: together with Plug, Socket provides the ability to embed widgets from
// one process into another process in a fashion that is transparent to the
// user. One process creates a Socket widget and passes that widget’s window ID
// to the other process, which then creates a Plug with that window ID. Any
// widgets contained in the Plug then will appear inside the first application’s
// window.
//
// The socket’s window ID is obtained by using gtk_socket_get_id(). Before using
// this function, the socket must have been realized, and for hence, have been
// added to its parent.
//
// Obtaining the window ID of a socket.
//
//    GtkWidget *socket = gtk_socket_new ();
//    gtk_widget_show (socket);
//    gtk_container_add (GTK_CONTAINER (parent), socket);
//
//    // The following call is only necessary if one of
//    // the ancestors of the socket is not yet visible.
//    gtk_widget_realize (socket);
//    g_print ("The ID of the sockets window is %#x\n",
//             gtk_socket_get_id (socket));
//
// Note that if you pass the window ID of the socket to another process that
// will create a plug in the socket, you must make sure that the socket widget
// is not destroyed until that plug is created. Violating this rule will cause
// unpredictable consequences, the most likely consequence being that the plug
// will appear as a separate toplevel window. You can check if the plug has been
// created by using gtk_socket_get_plug_window(). If it returns a non-NULL
// value, then the plug has been successfully created inside of the socket.
//
// When GTK+ is notified that the embedded window has been destroyed, then it
// will destroy the socket as well. You should always, therefore, be prepared
// for your sockets to be destroyed at any time when the main event loop is
// running. To prevent this from happening, you can connect to the
// Socket::plug-removed signal.
//
// The communication between a Socket and a Plug follows the XEmbed Protocol
// (http://www.freedesktop.org/Standards/xembed-spec). This protocol has also
// been implemented in other toolkits, e.g. Qt, allowing the same level of
// integration when embedding a Qt widget in GTK or vice versa.
//
// The Plug and Socket widgets are only available when GTK+ is compiled for the
// X11 platform and GDK_WINDOWING_X11 is defined. They can only be used on a
// X11Display. To use Plug and Socket, you need to include the gtk/gtkx.h
// header.
type Socket struct {
	_ [0]func() // equal guard
	Container
}

var (
	_ Containerer = (*Socket)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeSocket,
		GoType:        reflect.TypeOf((*Socket)(nil)),
		InitClass:     initClassSocket,
		FinalizeClass: finalizeClassSocket,
	})
}

func initClassSocket(gclass unsafe.Pointer, goval any) {

	pclass := (*C.GtkSocketClass)(unsafe.Pointer(gclass))

	if _, ok := goval.(interface{ PlugAdded() }); ok {
		pclass.plug_added = (*[0]byte)(C._gotk4_gtk3_SocketClass_plug_added)
	}

	if _, ok := goval.(interface{ PlugRemoved() bool }); ok {
		pclass.plug_removed = (*[0]byte)(C._gotk4_gtk3_SocketClass_plug_removed)
	}
	if goval, ok := goval.(interface{ InitSocket(*SocketClass) }); ok {
		klass := (*SocketClass)(gextras.NewStructNative(gclass))
		goval.InitSocket(klass)
	}
}

func finalizeClassSocket(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ FinalizeSocket(*SocketClass) }); ok {
		klass := (*SocketClass)(gextras.NewStructNative(gclass))
		goval.FinalizeSocket(klass)
	}
}

//export _gotk4_gtk3_SocketClass_plug_added
func _gotk4_gtk3_SocketClass_plug_added(arg0 *C.GtkSocket) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface{ PlugAdded() })

	iface.PlugAdded()
}

//export _gotk4_gtk3_SocketClass_plug_removed
func _gotk4_gtk3_SocketClass_plug_removed(arg0 *C.GtkSocket) (cret C.gboolean) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface{ PlugRemoved() bool })

	ok := iface.PlugRemoved()

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapSocket(obj *coreglib.Object) *Socket {
	return &Socket{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
	}
}

func marshalSocket(p uintptr) (interface{}, error) {
	return wrapSocket(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Socket_ConnectPlugAdded
func _gotk4_gtk3_Socket_ConnectPlugAdded(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectPlugAdded: this signal is emitted when a client is successfully added
// to the socket.
func (socket_ *Socket) ConnectPlugAdded(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(socket_, "plug-added", false, unsafe.Pointer(C._gotk4_gtk3_Socket_ConnectPlugAdded), f)
}

//export _gotk4_gtk3_Socket_ConnectPlugRemoved
func _gotk4_gtk3_Socket_ConnectPlugRemoved(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (ok bool))
	}

	ok := f()

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectPlugRemoved: this signal is emitted when a client is removed from the
// socket. The default action is to destroy the Socket widget, so if you want to
// reuse it you must add a signal handler that returns TRUE.
func (socket_ *Socket) ConnectPlugRemoved(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(socket_, "plug-removed", false, unsafe.Pointer(C._gotk4_gtk3_Socket_ConnectPlugRemoved), f)
}

// NewSocket: create a new empty Socket.
//
// The function returns the following values:
//
//    - socket: new Socket.
//
func NewSocket() *Socket {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_socket_new()

	var _socket *Socket // out

	_socket = wrapSocket(coreglib.Take(unsafe.Pointer(_cret)))

	return _socket
}

// PlugWindow retrieves the window of the plug. Use this to check if the plug
// has been created inside of the socket.
//
// The function returns the following values:
//
//    - window (optional) of the plug if available, or NULL.
//
func (socket_ *Socket) PlugWindow() gdk.Windower {
	var _arg0 *C.GtkSocket // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GtkSocket)(unsafe.Pointer(coreglib.InternObject(socket_).Native()))

	_cret = C.gtk_socket_get_plug_window(_arg0)
	runtime.KeepAlive(socket_)

	var _window gdk.Windower // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gdk.Windower)
				return ok
			})
			rv, ok := casted.(gdk.Windower)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
			}
			_window = rv
		}
	}

	return _window
}

// SocketClass: instance of this type is always passed by reference.
type SocketClass struct {
	*socketClass
}

// socketClass is the struct that's finalized.
type socketClass struct {
	native *C.GtkSocketClass
}

func (s *SocketClass) ParentClass() *ContainerClass {
	valptr := &s.native.parent_class
	var _v *ContainerClass // out
	_v = (*ContainerClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
