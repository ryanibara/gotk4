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

// UnsetPlacement unsets the placement of the contents with respect to the
// scrollbars for the scrolled window. If no window placement is set for a
// scrolled window, it defaults to GTK_CORNER_TOP_LEFT.
//
// See also gtk_scrolled_window_set_placement() and
// gtk_scrolled_window_get_placement().
func (scrolledWindow *ScrolledWindow) UnsetPlacement() {
	var _arg0 *C.GtkScrolledWindow // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	C.gtk_scrolled_window_unset_placement(_arg0)
	runtime.KeepAlive(scrolledWindow)
}
