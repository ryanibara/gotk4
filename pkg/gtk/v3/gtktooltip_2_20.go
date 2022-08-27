// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// SetIconFromGIcon sets the icon of the tooltip (which is in front of the text)
// to be the icon indicated by gicon with the size indicated by size. If gicon
// is NULL, the image will be hidden.
//
// The function takes the following parameters:
//
//    - gicon (optional) representing the icon, or NULL.
//    - size: stock icon size (IconSize).
//
func (tooltip *Tooltip) SetIconFromGIcon(gicon gio.Iconner, size int) {
	var _arg0 *C.GtkTooltip // out
	var _arg1 *C.GIcon      // out
	var _arg2 C.GtkIconSize // out

	_arg0 = (*C.GtkTooltip)(unsafe.Pointer(coreglib.InternObject(tooltip).Native()))
	if gicon != nil {
		_arg1 = (*C.GIcon)(unsafe.Pointer(coreglib.InternObject(gicon).Native()))
	}
	_arg2 = C.GtkIconSize(size)

	C.gtk_tooltip_set_icon_from_gicon(_arg0, _arg1, _arg2)
	runtime.KeepAlive(tooltip)
	runtime.KeepAlive(gicon)
	runtime.KeepAlive(size)
}
