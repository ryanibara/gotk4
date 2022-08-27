// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// ChildNotify emits a Widget::child-notify signal for the [child
// property][child-properties] child_property on the child.
//
// This is an analogue of g_object_notify() for child properties.
//
// Also see gtk_widget_child_notify().
//
// The function takes the following parameters:
//
//    - child widget.
//    - childProperty: name of a child property installed on the class of
//      container.
//
func (container *Container) ChildNotify(child Widgetter, childProperty string) {
	var _arg0 *C.GtkContainer // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.gchar        // out

	_arg0 = (*C.GtkContainer)(unsafe.Pointer(coreglib.InternObject(container).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(childProperty)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_container_child_notify(_arg0, _arg1, _arg2)
	runtime.KeepAlive(container)
	runtime.KeepAlive(child)
	runtime.KeepAlive(childProperty)
}
