// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_plug_get_type()), F: marshalPlugger},
	})
}

// PlugOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PlugOverrider interface {
	Embedded()
}

// Plug: together with Socket, Plug provides the ability to embed widgets from
// one process into another process in a fashion that is transparent to the
// user. One process creates a Socket widget and passes the ID of that widget’s
// window to the other process, which then creates a Plug with that window ID.
// Any widgets contained in the Plug then will appear inside the first
// application’s window.
//
// The communication between a Socket and a Plug follows the XEmbed Protocol
// (http://www.freedesktop.org/Standards/xembed-spec). This protocol has also
// been implemented in other toolkits, e.g. Qt, allowing the same level of
// integration when embedding a Qt widget in GTK+ or vice versa.
//
// The Plug and Socket widgets are only available when GTK+ is compiled for the
// X11 platform and GDK_WINDOWING_X11 is defined. They can only be used on a
// X11Display. To use Plug and Socket, you need to include the gtk/gtkx.h
// header.
type Plug struct {
	Window
}

func wrapPlug(obj *externglib.Object) *Plug {
	return &Plug{
		Window: Window{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalPlugger(p uintptr) (interface{}, error) {
	return wrapPlug(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Embedded determines whether the plug is embedded in a socket.
func (plug *Plug) Embedded() bool {
	var _arg0 *C.GtkPlug // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkPlug)(unsafe.Pointer(plug.Native()))

	_cret = C.gtk_plug_get_embedded(_arg0)
	runtime.KeepAlive(plug)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SocketWindow retrieves the socket the plug is embedded in.
func (plug *Plug) SocketWindow() gdk.Windower {
	var _arg0 *C.GtkPlug   // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GtkPlug)(unsafe.Pointer(plug.Native()))

	_cret = C.gtk_plug_get_socket_window(_arg0)
	runtime.KeepAlive(plug)

	var _window gdk.Windower // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(gdk.Windower)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Windower")
			}
			_window = rv
		}
	}

	return _window
}

// ConnectEmbedded gets emitted when the plug becomes embedded in a socket.
func (plug *Plug) ConnectEmbedded(f func()) externglib.SignalHandle {
	return plug.Connect("embedded", f)
}
