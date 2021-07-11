// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_root_get_type()), F: marshalRooter},
	})
}

// Rooter describes Root's methods.
type Rooter interface {
	// Display returns the display that this `GtkRoot` is on.
	Display() *gdk.Display
	// Focus retrieves the current focused widget within the root.
	Focus() *Widget
	// SetFocus: if @focus is not the current focus widget, and is focusable,
	// sets it as the focus widget for the root.
	SetFocus(focus Widgetter)
}

// Root: `GtkRoot` is the interface implemented by all widgets that can act as a
// toplevel widget.
//
// The root widget takes care of providing the connection to the windowing
// system and manages layout, drawing and event delivery for its widget
// hierarchy.
//
// The obvious example of a `GtkRoot` is `GtkWindow`.
//
// To get the display to which a `GtkRoot` belongs, use
// [method@Gtk.Root.get_display].
//
// `GtkRoot` also maintains the location of keyboard focus inside its widget
// hierarchy, with [method@Gtk.Root.set_focus] and [method@Gtk.Root.get_focus].
type Root struct {
	Native
}

var (
	_ Rooter          = (*Root)(nil)
	_ gextras.Nativer = (*Root)(nil)
)

func wrapRoot(obj *externglib.Object) Rooter {
	return &Root{
		Native: Native{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
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
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRoot(obj), nil
}

// Display returns the display that this `GtkRoot` is on.
func (self *Root) Display() *gdk.Display {
	var _arg0 *C.GtkRoot    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GtkRoot)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_root_get_display(_arg0)

	var _display *gdk.Display // out

	_display = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gdk.Display)

	return _display
}

// Focus retrieves the current focused widget within the root.
//
// Note that this is the widget that would have the focus if the root is active;
// if the root is not focused then `gtk_widget_has_focus (widget)` will be false
// for the widget.
func (self *Root) Focus() *Widget {
	var _arg0 *C.GtkRoot   // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkRoot)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_root_get_focus(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// SetFocus: if @focus is not the current focus widget, and is focusable, sets
// it as the focus widget for the root.
//
// If @focus is nil, unsets the focus widget for the root.
//
// To set the focus to a particular widget in the root, it is usually more
// convenient to use [method@Gtk.Widget.grab_focus] instead of this function.
func (self *Root) SetFocus(focus Widgetter) {
	var _arg0 *C.GtkRoot   // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkRoot)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((focus).(gextras.Nativer).Native()))

	C.gtk_root_set_focus(_arg0, _arg1)
}
