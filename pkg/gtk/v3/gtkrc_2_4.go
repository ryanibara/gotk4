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

// RCResetStyles: this function recomputes the styles for all widgets that use
// a particular Settings object. (There is one Settings object per Screen,
// see gtk_settings_get_for_screen()); It is useful when some global parameter
// has changed that affects the appearance of all widgets, because when a widget
// gets a new style, it will both redraw and recompute any cached information
// about its appearance. As an example, it is used when the default font size
// set by the operating system changes. Note that this function doesn’t affect
// widgets that have a style set explicitly on them with gtk_widget_set_style().
//
// Deprecated: Use CssProvider instead.
//
// The function takes the following parameters:
//
//   - settings: Settings.
//
func RCResetStyles(settings *Settings) {
	var _arg1 *C.GtkSettings // out

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))

	C.gtk_rc_reset_styles(_arg1)
	runtime.KeepAlive(settings)
}
