// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_window_group_get_type()), F: marshalWindowGrouper},
	})
}

// WindowGroup restricts the effect of grabs to windows in the same group,
// thereby making window groups almost behave like separate applications.
//
// A window can be a member in at most one window group at a time. Windows that
// have not been explicitly assigned to a group are implicitly treated like
// windows of the default window group.
//
// GtkWindowGroup objects are referenced by each window in the group, so once
// you have added all windows to a GtkWindowGroup, you can drop the initial
// reference to the window group with g_object_unref(). If the windows in the
// window group are subsequently destroyed, then they will be removed from the
// window group and drop their references on the window group; when all window
// have been removed, the window group will be freed.
type WindowGroup struct {
	*externglib.Object

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ externglib.Objector = (*WindowGroup)(nil)
)

func wrapWindowGroup(obj *externglib.Object) *WindowGroup {
	return &WindowGroup{
		Object: obj,
	}
}

func marshalWindowGrouper(p uintptr) (interface{}, error) {
	return wrapWindowGroup(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewWindowGroup creates a new WindowGroup object. Grabs added with
// gtk_grab_add() only affect windows within the same WindowGroup.
//
// The function returns the following values:
//
//    - windowGroup: new WindowGroup.
//
func NewWindowGroup() *WindowGroup {
	var _cret *C.GtkWindowGroup // in

	_cret = C.gtk_window_group_new()

	var _windowGroup *WindowGroup // out

	_windowGroup = wrapWindowGroup(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _windowGroup
}

// AddWindow adds a window to a WindowGroup.
//
// The function takes the following parameters:
//
//    - window to add.
//
func (windowGroup *WindowGroup) AddWindow(window *Window) {
	var _arg0 *C.GtkWindowGroup // out
	var _arg1 *C.GtkWindow      // out

	_arg0 = (*C.GtkWindowGroup)(unsafe.Pointer(windowGroup.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_window_group_add_window(_arg0, _arg1)
	runtime.KeepAlive(windowGroup)
	runtime.KeepAlive(window)
}

// CurrentDeviceGrab returns the current grab widget for device, or NULL if
// none.
//
// The function takes the following parameters:
//
//    - device: Device.
//
// The function returns the following values:
//
//    - widget (optional): grab widget, or NULL.
//
func (windowGroup *WindowGroup) CurrentDeviceGrab(device gdk.Devicer) Widgetter {
	var _arg0 *C.GtkWindowGroup // out
	var _arg1 *C.GdkDevice      // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkWindowGroup)(unsafe.Pointer(windowGroup.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gtk_window_group_get_current_device_grab(_arg0, _arg1)
	runtime.KeepAlive(windowGroup)
	runtime.KeepAlive(device)

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

// CurrentGrab gets the current grab widget of the given group, see
// gtk_grab_add().
//
// The function returns the following values:
//
//    - widget: current grab widget of the group.
//
func (windowGroup *WindowGroup) CurrentGrab() Widgetter {
	var _arg0 *C.GtkWindowGroup // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkWindowGroup)(unsafe.Pointer(windowGroup.Native()))

	_cret = C.gtk_window_group_get_current_grab(_arg0)
	runtime.KeepAlive(windowGroup)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// ListWindows returns a list of the Windows that belong to window_group.
//
// The function returns the following values:
//
//    - list: a newly-allocated list of windows inside the group.
//
func (windowGroup *WindowGroup) ListWindows() []Window {
	var _arg0 *C.GtkWindowGroup // out
	var _cret *C.GList          // in

	_arg0 = (*C.GtkWindowGroup)(unsafe.Pointer(windowGroup.Native()))

	_cret = C.gtk_window_group_list_windows(_arg0)
	runtime.KeepAlive(windowGroup)

	var _list []Window // out

	_list = make([]Window, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkWindow)(v)
		var dst Window // out
		dst = *wrapWindow(externglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// RemoveWindow removes a window from a WindowGroup.
//
// The function takes the following parameters:
//
//    - window to remove.
//
func (windowGroup *WindowGroup) RemoveWindow(window *Window) {
	var _arg0 *C.GtkWindowGroup // out
	var _arg1 *C.GtkWindow      // out

	_arg0 = (*C.GtkWindowGroup)(unsafe.Pointer(windowGroup.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_window_group_remove_window(_arg0, _arg1)
	runtime.KeepAlive(windowGroup)
	runtime.KeepAlive(window)
}
