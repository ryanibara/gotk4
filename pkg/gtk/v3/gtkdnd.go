// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// DragCancel cancels an ongoing drag operation on the source side.
//
// If you want to be able to cancel a drag operation in this way, you need to
// keep a pointer to the drag context, either from an explicit call to
// gtk_drag_begin_with_coordinates(), or by connecting to Widget::drag-begin.
//
// If context does not refer to an ongoing drag operation, this function does
// nothing.
//
// If a drag is cancelled in this way, the result argument of
// Widget::drag-failed is set to GTK_DRAG_RESULT_ERROR.
//
// The function takes the following parameters:
//
//    - context as e.g. returned by gtk_drag_begin_with_coordinates().
//
func DragCancel(context *gdk.DragContext) {
	var _arg1 *C.GdkDragContext // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	C.gtk_drag_cancel(_arg1)
	runtime.KeepAlive(context)
}

// DragFinish informs the drag source that the drop is finished, and that the
// data of the drag will no longer be required.
//
// The function takes the following parameters:
//
//    - context: drag context.
//    - success: flag indicating whether the drop was successful.
//    - del: flag indicating whether the source should delete the original
//    data. (This should be TRUE for a move).
//    - time_: timestamp from the Widget::drag-drop signal.
//
func DragFinish(context *gdk.DragContext, success, del bool, time_ uint32) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 C.gboolean        // out
	var _arg3 C.gboolean        // out
	var _arg4 C.guint32         // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	if success {
		_arg2 = C.TRUE
	}
	if del {
		_arg3 = C.TRUE
	}
	_arg4 = C.guint32(time_)

	C.gtk_drag_finish(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(context)
	runtime.KeepAlive(success)
	runtime.KeepAlive(del)
	runtime.KeepAlive(time_)
}

// DragGetSourceWidget determines the source widget for a drag.
//
// The function takes the following parameters:
//
//    - context: (destination side) drag context.
//
func DragGetSourceWidget(context *gdk.DragContext) Widgetter {
	var _arg1 *C.GdkDragContext // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_drag_get_source_widget(_arg1)
	runtime.KeepAlive(context)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// DragSetIconDefault sets the icon for a particular drag to the default icon.
//
// The function takes the following parameters:
//
//    - context for a drag (This must be called with a context for the source
//    side of a drag).
//
func DragSetIconDefault(context *gdk.DragContext) {
	var _arg1 *C.GdkDragContext // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	C.gtk_drag_set_icon_default(_arg1)
	runtime.KeepAlive(context)
}

// DragSetIconGIcon sets the icon for a given drag from the given icon. See the
// documentation for gtk_drag_set_icon_name() for more details about using icons
// in drag and drop.
//
// The function takes the following parameters:
//
//    - context for a drag (This must be called with a context for the source
//    side of a drag).
//    - icon: #GIcon.
//    - hotX: x offset of the hotspot within the icon.
//    - hotY: y offset of the hotspot within the icon.
//
func DragSetIconGIcon(context *gdk.DragContext, icon gio.Iconner, hotX, hotY int) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 *C.GIcon          // out
	var _arg3 C.gint            // out
	var _arg4 C.gint            // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GIcon)(unsafe.Pointer(icon.Native()))
	_arg3 = C.gint(hotX)
	_arg4 = C.gint(hotY)

	C.gtk_drag_set_icon_gicon(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(context)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(hotX)
	runtime.KeepAlive(hotY)
}

// DragSetIconName sets the icon for a given drag from a named themed icon. See
// the docs for IconTheme for more details. Note that the size of the icon
// depends on the icon theme (the icon is loaded at the symbolic size
// K_ICON_SIZE_DND), thus hot_x and hot_y have to be used with care.
//
// The function takes the following parameters:
//
//    - context for a drag (This must be called with a context for the source
//    side of a drag).
//    - iconName: name of icon to use.
//    - hotX: x offset of the hotspot within the icon.
//    - hotY: y offset of the hotspot within the icon.
//
func DragSetIconName(context *gdk.DragContext, iconName string, hotX, hotY int) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 *C.gchar          // out
	var _arg3 C.gint            // out
	var _arg4 C.gint            // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gint(hotX)
	_arg4 = C.gint(hotY)

	C.gtk_drag_set_icon_name(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(context)
	runtime.KeepAlive(iconName)
	runtime.KeepAlive(hotX)
	runtime.KeepAlive(hotY)
}

// DragSetIconPixbuf sets pixbuf as the icon for a given drag.
//
// The function takes the following parameters:
//
//    - context for a drag (This must be called with a context for the source
//    side of a drag).
//    - pixbuf to use as the drag icon.
//    - hotX: x offset within widget of the hotspot.
//    - hotY: y offset within widget of the hotspot.
//
func DragSetIconPixbuf(context *gdk.DragContext, pixbuf *gdkpixbuf.Pixbuf, hotX, hotY int) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 *C.GdkPixbuf      // out
	var _arg3 C.gint            // out
	var _arg4 C.gint            // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	_arg3 = C.gint(hotX)
	_arg4 = C.gint(hotY)

	C.gtk_drag_set_icon_pixbuf(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(context)
	runtime.KeepAlive(pixbuf)
	runtime.KeepAlive(hotX)
	runtime.KeepAlive(hotY)
}

// DragSetIconStock sets the icon for a given drag from a stock ID.
//
// Deprecated: Use gtk_drag_set_icon_name() instead.
//
// The function takes the following parameters:
//
//    - context for a drag (This must be called with a context for the source
//    side of a drag).
//    - stockId: ID of the stock icon to use for the drag.
//    - hotX: x offset within the icon of the hotspot.
//    - hotY: y offset within the icon of the hotspot.
//
func DragSetIconStock(context *gdk.DragContext, stockId string, hotX, hotY int) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 *C.gchar          // out
	var _arg3 C.gint            // out
	var _arg4 C.gint            // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gint(hotX)
	_arg4 = C.gint(hotY)

	C.gtk_drag_set_icon_stock(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(context)
	runtime.KeepAlive(stockId)
	runtime.KeepAlive(hotX)
	runtime.KeepAlive(hotY)
}

// DragSetIconSurface sets surface as the icon for a given drag. GTK+ retains
// references for the arguments, and will release them when they are no longer
// needed.
//
// To position the surface relative to the mouse, use
// cairo_surface_set_device_offset() on surface. The mouse cursor will be
// positioned at the (0,0) coordinate of the surface.
//
// The function takes the following parameters:
//
//    - context for a drag (This must be called with a context for the source
//    side of a drag).
//    - surface to use as icon.
//
func DragSetIconSurface(context *gdk.DragContext, surface *cairo.Surface) {
	var _arg1 *C.GdkDragContext  // out
	var _arg2 *C.cairo_surface_t // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))

	C.gtk_drag_set_icon_surface(_arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(surface)
}

// DragSetIconWidget changes the icon for drag operation to a given widget. GTK+
// will not destroy the widget, so if you don’t want it to persist, you should
// connect to the “drag-end” signal and destroy it yourself.
//
// The function takes the following parameters:
//
//    - context for a drag. (This must be called with a context for the source
//    side of a drag).
//    - widget to use as an icon.
//    - hotX: x offset within widget of the hotspot.
//    - hotY: y offset within widget of the hotspot.
//
func DragSetIconWidget(context *gdk.DragContext, widget Widgetter, hotX, hotY int) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 *C.GtkWidget      // out
	var _arg3 C.gint            // out
	var _arg4 C.gint            // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = C.gint(hotX)
	_arg4 = C.gint(hotY)

	C.gtk_drag_set_icon_widget(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(context)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(hotX)
	runtime.KeepAlive(hotY)
}
