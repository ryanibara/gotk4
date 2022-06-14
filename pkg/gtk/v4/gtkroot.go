// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkroot.go.
var GTypeRoot = coreglib.Type(C.gtk_root_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeRoot, F: marshalRoot},
	})
}

// Root: GtkRoot is the interface implemented by all widgets that can act as a
// toplevel widget.
//
// The root widget takes care of providing the connection to the windowing
// system and manages layout, drawing and event delivery for its widget
// hierarchy.
//
// The obvious example of a GtkRoot is GtkWindow.
//
// To get the display to which a GtkRoot belongs, use gtk.Root.GetDisplay().
//
// GtkRoot also maintains the location of keyboard focus inside its widget
// hierarchy, with gtk.Root.SetFocus() and gtk.Root.GetFocus().
//
// Root wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Root struct {
	_ [0]func() // equal guard
	NativeSurface
}

var ()

// Rooter describes Root's interface methods.
type Rooter interface {
	coreglib.Objector

	// Display returns the display that this GtkRoot is on.
	Display() *gdk.Display
	// Focus retrieves the current focused widget within the root.
	Focus() Widgetter
	// SetFocus: if focus is not the current focus widget, and is focusable,
	// sets it as the focus widget for the root.
	SetFocus(focus Widgetter)
}

var _ Rooter = (*Root)(nil)

func wrapRoot(obj *coreglib.Object) *Root {
	return &Root{
		NativeSurface: NativeSurface{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
		},
	}
}

func marshalRoot(p uintptr) (interface{}, error) {
	return wrapRoot(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Display returns the display that this GtkRoot is on.
//
// The function returns the following values:
//
//    - display of root.
//
func (self *Root) Display() *gdk.Display {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _display *gdk.Display // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_display = &gdk.Display{
			Object: obj,
		}
	}

	return _display
}

// Focus retrieves the current focused widget within the root.
//
// Note that this is the widget that would have the focus if the root is active;
// if the root is not focused then gtk_widget_has_focus (widget) will be FALSE
// for the widget.
//
// The function returns the following values:
//
//    - widget (optional): currently focused widget, or NULL if there is none.
//
func (self *Root) Focus() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// SetFocus: if focus is not the current focus widget, and is focusable, sets it
// as the focus widget for the root.
//
// If focus is NULL, unsets the focus widget for the root.
//
// To set the focus to a particular widget in the root, it is usually more
// convenient to use gtk.Widget.GrabFocus() instead of this function.
//
// The function takes the following parameters:
//
//    - focus (optional): widget to be the new focus widget, or NULL to unset the
//      focus widget.
//
func (self *Root) SetFocus(focus Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if focus != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(focus).Native()))
	}

	runtime.KeepAlive(self)
	runtime.KeepAlive(focus)
}
