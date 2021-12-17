// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_root_get_type()), F: marshalRooter},
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
type Root struct {
	NativeSurface
}

var ()

// Rooter describes Root's interface methods.
type Rooter interface {
	externglib.Objector

	// Display returns the display that this GtkRoot is on.
	Display() *gdk.Display
	// Focus retrieves the current focused widget within the root.
	Focus() Widgetter
	// SetFocus: if focus is not the current focus widget, and is focusable,
	// sets it as the focus widget for the root.
	SetFocus(focus Widgetter)
}

var _ Rooter = (*Root)(nil)

func wrapRoot(obj *externglib.Object) *Root {
	return &Root{
		NativeSurface: NativeSurface{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
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

func marshalRooter(p uintptr) (interface{}, error) {
	return wrapRoot(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Display returns the display that this GtkRoot is on.
func (self *Root) Display() *gdk.Display {
	var _arg0 *C.GtkRoot    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GtkRoot)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_root_get_display(_arg0)
	runtime.KeepAlive(self)

	var _display *gdk.Display // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
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
func (self *Root) Focus() Widgetter {
	var _arg0 *C.GtkRoot   // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkRoot)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_root_get_focus(_arg0)
	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
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
//    - focus: widget to be the new focus widget, or NULL to unset the focus
//    widget.
//
func (self *Root) SetFocus(focus Widgetter) {
	var _arg0 *C.GtkRoot   // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkRoot)(unsafe.Pointer(self.Native()))
	if focus != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(focus.Native()))
	}

	C.gtk_root_set_focus(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(focus)
}
