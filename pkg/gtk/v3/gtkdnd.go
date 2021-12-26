// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
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
//    - del: flag indicating whether the source should delete the original data.
//      (This should be TRUE for a move).
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
// The function returns the following values:
//
//    - widget (optional): if the drag is occurring within a single application,
//      a pointer to the source widget. Otherwise, NULL.
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
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
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
//      side of a drag).
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
//      side of a drag).
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
//      side of a drag).
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
//      side of a drag).
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
//      side of a drag).
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
//      side of a drag).
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
//      side of a drag).
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

// DragBegin: this function is equivalent to gtk_drag_begin_with_coordinates(),
// passing -1, -1 as coordinates.
//
// Deprecated: Use gtk_drag_begin_with_coordinates() instead.
//
// The function takes the following parameters:
//
//    - targets (data formats) in which the source can provide the data.
//    - actions: bitmask of the allowed drag actions for this drag.
//    - button the user clicked to start the drag.
//    - event (optional) that triggered the start of the drag, or NULL if none
//      can be obtained.
//
// The function returns the following values:
//
//    - dragContext: context for this drag.
//
func (widget *Widget) DragBegin(targets *TargetList, actions gdk.DragAction, button int, event *gdk.Event) *gdk.DragContext {
	var _arg0 *C.GtkWidget      // out
	var _arg1 *C.GtkTargetList  // out
	var _arg2 C.GdkDragAction   // out
	var _arg3 C.gint            // out
	var _arg4 *C.GdkEvent       // out
	var _cret *C.GdkDragContext // in

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg1 = (*C.GtkTargetList)(gextras.StructNative(unsafe.Pointer(targets)))
	_arg2 = C.GdkDragAction(actions)
	_arg3 = C.gint(button)
	if event != nil {
		_arg4 = (*C.GdkEvent)(gextras.StructNative(unsafe.Pointer(event)))
	}

	_cret = C.gtk_drag_begin(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(targets)
	runtime.KeepAlive(actions)
	runtime.KeepAlive(button)
	runtime.KeepAlive(event)

	var _dragContext *gdk.DragContext // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_dragContext = &gdk.DragContext{
			Object: obj,
		}
	}

	return _dragContext
}

// DragBeginWithCoordinates initiates a drag on the source side. The function
// only needs to be used when the application is starting drags itself, and is
// not needed when gtk_drag_source_set() is used.
//
// The event is used to retrieve the timestamp that will be used internally to
// grab the pointer. If event is NULL, then GDK_CURRENT_TIME will be used.
// However, you should try to pass a real event in all cases, since that can be
// used to get information about the drag.
//
// Generally there are three cases when you want to start a drag by hand by
// calling this function:
//
// 1. During a Widget::button-press-event handler, if you want to start a drag
// immediately when the user presses the mouse button. Pass the event that you
// have in your Widget::button-press-event handler.
//
// 2. During a Widget::motion-notify-event handler, if you want to start a drag
// when the mouse moves past a certain threshold distance after a button-press.
// Pass the event that you have in your Widget::motion-notify-event handler.
//
// 3. During a timeout handler, if you want to start a drag after the mouse
// button is held down for some time. Try to save the last event that you got
// from the mouse, using gdk_event_copy(), and pass it to this function
// (remember to free the event with gdk_event_free() when you are done). If you
// really cannot pass a real event, pass NULL instead.
//
// The function takes the following parameters:
//
//    - targets (data formats) in which the source can provide the data.
//    - actions: bitmask of the allowed drag actions for this drag.
//    - button the user clicked to start the drag.
//    - event (optional) that triggered the start of the drag, or NULL if none
//      can be obtained.
//    - x: initial x coordinate to start dragging from, in the coordinate space
//      of widget. If -1 is passed, the coordinates are retrieved from event or
//      the current pointer position.
//    - y: initial y coordinate to start dragging from, in the coordinate space
//      of widget. If -1 is passed, the coordinates are retrieved from event or
//      the current pointer position.
//
// The function returns the following values:
//
//    - dragContext: context for this drag.
//
func (widget *Widget) DragBeginWithCoordinates(targets *TargetList, actions gdk.DragAction, button int, event *gdk.Event, x, y int) *gdk.DragContext {
	var _arg0 *C.GtkWidget      // out
	var _arg1 *C.GtkTargetList  // out
	var _arg2 C.GdkDragAction   // out
	var _arg3 C.gint            // out
	var _arg4 *C.GdkEvent       // out
	var _arg5 C.gint            // out
	var _arg6 C.gint            // out
	var _cret *C.GdkDragContext // in

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg1 = (*C.GtkTargetList)(gextras.StructNative(unsafe.Pointer(targets)))
	_arg2 = C.GdkDragAction(actions)
	_arg3 = C.gint(button)
	if event != nil {
		_arg4 = (*C.GdkEvent)(gextras.StructNative(unsafe.Pointer(event)))
	}
	_arg5 = C.gint(x)
	_arg6 = C.gint(y)

	_cret = C.gtk_drag_begin_with_coordinates(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(targets)
	runtime.KeepAlive(actions)
	runtime.KeepAlive(button)
	runtime.KeepAlive(event)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _dragContext *gdk.DragContext // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_dragContext = &gdk.DragContext{
			Object: obj,
		}
	}

	return _dragContext
}

// DragCheckThreshold checks to see if a mouse drag starting at (start_x,
// start_y) and ending at (current_x, current_y) has passed the GTK+ drag
// threshold, and thus should trigger the beginning of a drag-and-drop
// operation.
//
// The function takes the following parameters:
//
//    - startX: x coordinate of start of drag.
//    - startY: y coordinate of start of drag.
//    - currentX: current X coordinate.
//    - currentY: current Y coordinate.
//
// The function returns the following values:
//
//    - ok: TRUE if the drag threshold has been passed.
//
func (widget *Widget) DragCheckThreshold(startX, startY, currentX, currentY int) bool {
	var _arg0 *C.GtkWidget // out
	var _arg1 C.gint       // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out
	var _arg4 C.gint       // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg1 = C.gint(startX)
	_arg2 = C.gint(startY)
	_arg3 = C.gint(currentX)
	_arg4 = C.gint(currentY)

	_cret = C.gtk_drag_check_threshold(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(startX)
	runtime.KeepAlive(startY)
	runtime.KeepAlive(currentX)
	runtime.KeepAlive(currentY)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DragHighlight highlights a widget as a currently hovered drop target. To end
// the highlight, call gtk_drag_unhighlight(). GTK+ calls this automatically if
// GTK_DEST_DEFAULT_HIGHLIGHT is set.
func (widget *Widget) DragHighlight() {
	var _arg0 *C.GtkWidget // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_drag_highlight(_arg0)
	runtime.KeepAlive(widget)
}

// DragUnhighlight removes a highlight set by gtk_drag_highlight() from a
// widget.
func (widget *Widget) DragUnhighlight() {
	var _arg0 *C.GtkWidget // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_drag_unhighlight(_arg0)
	runtime.KeepAlive(widget)
}
