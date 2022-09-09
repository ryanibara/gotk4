// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeWindowGroup = coreglib.Type(C.gtk_window_group_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWindowGroup, F: marshalWindowGroup},
	})
}

// WindowGroupOverrides contains methods that are overridable.
type WindowGroupOverrides struct {
}

func defaultWindowGroupOverrides(v *WindowGroup) WindowGroupOverrides {
	return WindowGroupOverrides{}
}

// WindowGroup: GtkWindowGroup makes group of windows behave like separate
// applications.
//
// It achieves this by limiting the effect of GTK grabs and modality to windows
// in the same group.
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
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*WindowGroup)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*WindowGroup, *WindowGroupClass, WindowGroupOverrides](
		GTypeWindowGroup,
		initWindowGroupClass,
		wrapWindowGroup,
		defaultWindowGroupOverrides,
	)
}

func initWindowGroupClass(gclass unsafe.Pointer, overrides WindowGroupOverrides, classInitFunc func(*WindowGroupClass)) {
	if classInitFunc != nil {
		class := (*WindowGroupClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapWindowGroup(obj *coreglib.Object) *WindowGroup {
	return &WindowGroup{
		Object: obj,
	}
}

func marshalWindowGroup(p uintptr) (interface{}, error) {
	return wrapWindowGroup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewWindowGroup creates a new GtkWindowGroup object.
//
// Modality of windows only affects windows within the same GtkWindowGroup.
//
// The function returns the following values:
//
//    - windowGroup: new GtkWindowGroup.
//
func NewWindowGroup() *WindowGroup {
	var _cret *C.GtkWindowGroup // in

	_cret = C.gtk_window_group_new()

	var _windowGroup *WindowGroup // out

	_windowGroup = wrapWindowGroup(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _windowGroup
}

// AddWindow adds a window to a GtkWindowGroup.
//
// The function takes the following parameters:
//
//    - window: GtkWindow to add.
//
func (windowGroup *WindowGroup) AddWindow(window *Window) {
	var _arg0 *C.GtkWindowGroup // out
	var _arg1 *C.GtkWindow      // out

	_arg0 = (*C.GtkWindowGroup)(unsafe.Pointer(coreglib.InternObject(windowGroup).Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	C.gtk_window_group_add_window(_arg0, _arg1)
	runtime.KeepAlive(windowGroup)
	runtime.KeepAlive(window)
}

// ListWindows returns a list of the GtkWindows that belong to window_group.
//
// The function returns the following values:
//
//    - list: a newly-allocated list of windows inside the group.
//
func (windowGroup *WindowGroup) ListWindows() []*Window {
	var _arg0 *C.GtkWindowGroup // out
	var _cret *C.GList          // in

	_arg0 = (*C.GtkWindowGroup)(unsafe.Pointer(coreglib.InternObject(windowGroup).Native()))

	_cret = C.gtk_window_group_list_windows(_arg0)
	runtime.KeepAlive(windowGroup)

	var _list []*Window // out

	_list = make([]*Window, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkWindow)(v)
		var dst *Window // out
		dst = wrapWindow(coreglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// RemoveWindow removes a window from a GtkWindowGroup.
//
// The function takes the following parameters:
//
//    - window: GtkWindow to remove.
//
func (windowGroup *WindowGroup) RemoveWindow(window *Window) {
	var _arg0 *C.GtkWindowGroup // out
	var _arg1 *C.GtkWindow      // out

	_arg0 = (*C.GtkWindowGroup)(unsafe.Pointer(coreglib.InternObject(windowGroup).Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	C.gtk_window_group_remove_window(_arg0, _arg1)
	runtime.KeepAlive(windowGroup)
	runtime.KeepAlive(window)
}

// WindowGroupClass: instance of this type is always passed by reference.
type WindowGroupClass struct {
	*windowGroupClass
}

// windowGroupClass is the struct that's finalized.
type windowGroupClass struct {
	native *C.GtkWindowGroupClass
}
