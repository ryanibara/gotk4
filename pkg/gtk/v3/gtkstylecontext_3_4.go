// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// RenderInsertionCursor draws a text caret on cr at the specified index of
// layout.
//
// The function takes the following parameters:
//
//   - context: StyleContext.
//   - cr: #cairo_t.
//   - x: x origin.
//   - y: y origin.
//   - layout of the text.
//   - index in the Layout.
//   - direction of the text.
//
func RenderInsertionCursor(context *StyleContext, cr *cairo.Context, x, y float64, layout *pango.Layout, index int, direction pango.Direction) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 *C.PangoLayout     // out
	var _arg6 C.int              // out
	var _arg7 C.PangoDirection   // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = (*C.PangoLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))
	_arg6 = C.int(index)
	_arg7 = C.PangoDirection(direction)

	C.gtk_render_insertion_cursor(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(context)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(index)
	runtime.KeepAlive(direction)
}
