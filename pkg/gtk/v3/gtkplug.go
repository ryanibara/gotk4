// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_PlugClass_embedded(GtkPlug*);
// extern void _gotk4_gtk3_Plug_ConnectEmbedded(gpointer, guintptr);
import "C"

// glib.Type values for gtkplug.go.
var GTypePlug = externglib.Type(C.gtk_plug_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypePlug, F: marshalPlug},
	})
}

// PlugOverrider contains methods that are overridable.
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
	_ [0]func() // equal guard
	Window
}

var (
	_ Binner = (*Plug)(nil)
)

func classInitPlugger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkPlugClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkPlugClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Embedded() }); ok {
		pclass.embedded = (*[0]byte)(C._gotk4_gtk3_PlugClass_embedded)
	}
}

//export _gotk4_gtk3_PlugClass_embedded
func _gotk4_gtk3_PlugClass_embedded(arg0 *C.GtkPlug) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Embedded() })

	iface.Embedded()
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
						Object: obj,
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
					},
				},
			},
		},
	}
}

func marshalPlug(p uintptr) (interface{}, error) {
	return wrapPlug(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Plug_ConnectEmbedded
func _gotk4_gtk3_Plug_ConnectEmbedded(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectEmbedded gets emitted when the plug becomes embedded in a socket.
func (plug *Plug) ConnectEmbedded(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(plug, "embedded", false, unsafe.Pointer(C._gotk4_gtk3_Plug_ConnectEmbedded), f)
}

// Embedded determines whether the plug is embedded in a socket.
//
// The function returns the following values:
//
//    - ok: TRUE if the plug is embedded in a socket.
//
func (plug *Plug) Embedded() bool {
	var _arg0 *C.GtkPlug // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkPlug)(unsafe.Pointer(externglib.InternObject(plug).Native()))

	_cret = C.gtk_plug_get_embedded(_arg0)
	runtime.KeepAlive(plug)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SocketWindow retrieves the socket the plug is embedded in.
//
// The function returns the following values:
//
//    - window (optional) of the socket, or NULL.
//
func (plug *Plug) SocketWindow() gdk.Windower {
	var _arg0 *C.GtkPlug   // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GtkPlug)(unsafe.Pointer(externglib.InternObject(plug).Native()))

	_cret = C.gtk_plug_get_socket_window(_arg0)
	runtime.KeepAlive(plug)

	var _window gdk.Windower // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
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
