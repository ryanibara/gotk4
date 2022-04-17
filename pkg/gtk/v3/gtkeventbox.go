// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtkeventbox.go.
var GTypeEventBox = externglib.Type(C.gtk_event_box_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeEventBox, F: marshalEventBox},
	})
}

// EventBoxOverrider contains methods that are overridable.
type EventBoxOverrider interface {
	externglib.Objector
}

// WrapEventBoxOverrider wraps the EventBoxOverrider
// interface implementation to access the instance methods.
func WrapEventBoxOverrider(obj EventBoxOverrider) *EventBox {
	return wrapEventBox(externglib.BaseObject(obj))
}

// EventBox widget is a subclass of Bin which also has its own window. It is
// useful since it allows you to catch events for widgets which do not have
// their own window.
type EventBox struct {
	_ [0]func() // equal guard
	Bin
}

var (
	_ Binner = (*EventBox)(nil)
)

func classInitEventBoxer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapEventBox(obj *externglib.Object) *EventBox {
	return &EventBox{
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
	}
}

func marshalEventBox(p uintptr) (interface{}, error) {
	return wrapEventBox(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewEventBox creates a new EventBox.
//
// The function returns the following values:
//
//    - eventBox: new EventBox.
//
func NewEventBox() *EventBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_event_box_new()

	var _eventBox *EventBox // out

	_eventBox = wrapEventBox(externglib.Take(unsafe.Pointer(_cret)))

	return _eventBox
}

// AboveChild returns whether the event box window is above or below the windows
// of its child. See gtk_event_box_set_above_child() for details.
//
// The function returns the following values:
//
//    - ok: TRUE if the event box window is above the window of its child.
//
func (eventBox *EventBox) AboveChild() bool {
	var _arg0 *C.GtkEventBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkEventBox)(unsafe.Pointer(externglib.InternObject(eventBox).Native()))

	_cret = C.gtk_event_box_get_above_child(_arg0)
	runtime.KeepAlive(eventBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VisibleWindow returns whether the event box has a visible window. See
// gtk_event_box_set_visible_window() for details.
//
// The function returns the following values:
//
//    - ok: TRUE if the event box window is visible.
//
func (eventBox *EventBox) VisibleWindow() bool {
	var _arg0 *C.GtkEventBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkEventBox)(unsafe.Pointer(externglib.InternObject(eventBox).Native()))

	_cret = C.gtk_event_box_get_visible_window(_arg0)
	runtime.KeepAlive(eventBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAboveChild: set whether the event box window is positioned above the
// windows of its child, as opposed to below it. If the window is above, all
// events inside the event box will go to the event box. If the window is below,
// events in windows of child widgets will first got to that widget, and then to
// its parents.
//
// The default is to keep the window below the child.
//
// The function takes the following parameters:
//
//    - aboveChild: TRUE if the event box window is above its child.
//
func (eventBox *EventBox) SetAboveChild(aboveChild bool) {
	var _arg0 *C.GtkEventBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkEventBox)(unsafe.Pointer(externglib.InternObject(eventBox).Native()))
	if aboveChild {
		_arg1 = C.TRUE
	}

	C.gtk_event_box_set_above_child(_arg0, _arg1)
	runtime.KeepAlive(eventBox)
	runtime.KeepAlive(aboveChild)
}

// SetVisibleWindow: set whether the event box uses a visible or invisible child
// window. The default is to use visible windows.
//
// In an invisible window event box, the window that the event box creates is a
// GDK_INPUT_ONLY window, which means that it is invisible and only serves to
// receive events.
//
// A visible window event box creates a visible (GDK_INPUT_OUTPUT) window that
// acts as the parent window for all the widgets contained in the event box.
//
// You should generally make your event box invisible if you just want to trap
// events. Creating a visible window may cause artifacts that are visible to the
// user, especially if the user is using a theme with gradients or pixmaps.
//
// The main reason to create a non input-only event box is if you want to set
// the background to a different color or draw on it.
//
// There is one unexpected issue for an invisible event box that has its window
// below the child. (See gtk_event_box_set_above_child().) Since the input-only
// window is not an ancestor window of any windows that descendent widgets of
// the event box create, events on these windows aren’t propagated up by the
// windowing system, but only by GTK+. The practical effect of this is if an
// event isn’t in the event mask for the descendant window (see
// gtk_widget_add_events()), it won’t be received by the event box.
//
// This problem doesn’t occur for visible event boxes, because in that case, the
// event box window is actually the ancestor of the descendant windows, not just
// at the same place on the screen.
//
// The function takes the following parameters:
//
//    - visibleWindow: TRUE to make the event box have a visible window.
//
func (eventBox *EventBox) SetVisibleWindow(visibleWindow bool) {
	var _arg0 *C.GtkEventBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkEventBox)(unsafe.Pointer(externglib.InternObject(eventBox).Native()))
	if visibleWindow {
		_arg1 = C.TRUE
	}

	C.gtk_event_box_set_visible_window(_arg0, _arg1)
	runtime.KeepAlive(eventBox)
	runtime.KeepAlive(visibleWindow)
}
