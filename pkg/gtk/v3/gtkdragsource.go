// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

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

// DragSourceAddImageTargets: add the writable image targets supported by
// SelectionData to the target list of the drag source. The targets are added
// with info = 0. If you need another value, use
// gtk_target_list_add_image_targets() and gtk_drag_source_set_target_list().
func (widget *Widget) DragSourceAddImageTargets() {
	var _arg0 *C.GtkWidget // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))

	C.gtk_drag_source_add_image_targets(_arg0)
	runtime.KeepAlive(widget)
}

// DragSourceAddTextTargets: add the text targets supported by SelectionData to
// the target list of the drag source. The targets are added with info = 0. If
// you need another value, use gtk_target_list_add_text_targets() and
// gtk_drag_source_set_target_list().
func (widget *Widget) DragSourceAddTextTargets() {
	var _arg0 *C.GtkWidget // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))

	C.gtk_drag_source_add_text_targets(_arg0)
	runtime.KeepAlive(widget)
}

// DragSourceAddURITargets: add the URI targets supported by SelectionData to
// the target list of the drag source. The targets are added with info = 0. If
// you need another value, use gtk_target_list_add_uri_targets() and
// gtk_drag_source_set_target_list().
func (widget *Widget) DragSourceAddURITargets() {
	var _arg0 *C.GtkWidget // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))

	C.gtk_drag_source_add_uri_targets(_arg0)
	runtime.KeepAlive(widget)
}

// DragSourceGetTargetList gets the list of targets this widget can provide for
// drag-and-drop.
//
// The function returns the following values:
//
//    - targetList (optional) or NULL if none.
//
func (widget *Widget) DragSourceGetTargetList() *TargetList {
	var _arg0 *C.GtkWidget     // out
	var _cret *C.GtkTargetList // in

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))

	_cret = C.gtk_drag_source_get_target_list(_arg0)
	runtime.KeepAlive(widget)

	var _targetList *TargetList // out

	if _cret != nil {
		_targetList = (*TargetList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.gtk_target_list_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_targetList)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gtk_target_list_unref((*C.GtkTargetList)(intern.C))
			},
		)
	}

	return _targetList
}

// DragSourceSet sets up a widget so that GTK+ will start a drag operation when
// the user clicks and drags on the widget. The widget must have a window.
//
// The function takes the following parameters:
//
//    - startButtonMask: bitmask of buttons that can start the drag.
//    - targets (optional): table of targets that the drag will support, may be
//      NULL.
//    - actions: bitmask of possible actions for a drag from this widget.
//
func (widget *Widget) DragSourceSet(startButtonMask gdk.ModifierType, targets []TargetEntry, actions gdk.DragAction) {
	var _arg0 *C.GtkWidget      // out
	var _arg1 C.GdkModifierType // out
	var _arg2 *C.GtkTargetEntry // out
	var _arg3 C.gint
	var _arg4 C.GdkDragAction // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))
	_arg1 = C.GdkModifierType(startButtonMask)
	_arg3 = (C.gint)(len(targets))
	_arg2 = (*C.GtkTargetEntry)(C.calloc(C.size_t(len(targets)), C.size_t(C.sizeof_GtkTargetEntry)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GtkTargetEntry)(_arg2), len(targets))
		for i := range targets {
			out[i] = *(*C.GtkTargetEntry)(gextras.StructNative(unsafe.Pointer((&targets[i]))))
		}
	}
	_arg4 = C.GdkDragAction(actions)

	C.gtk_drag_source_set(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(startButtonMask)
	runtime.KeepAlive(targets)
	runtime.KeepAlive(actions)
}

// DragSourceSetIconGIcon sets the icon that will be used for drags from a
// particular source to icon. See the docs for IconTheme for more details.
//
// The function takes the following parameters:
//
//    - icon: #GIcon.
//
func (widget *Widget) DragSourceSetIconGIcon(icon gio.IconOverrider) {
	var _arg0 *C.GtkWidget // out
	var _arg1 *C.GIcon     // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(externglib.InternObject(icon).Native()))

	C.gtk_drag_source_set_icon_gicon(_arg0, _arg1)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(icon)
}

// DragSourceSetIconName sets the icon that will be used for drags from a
// particular source to a themed icon. See the docs for IconTheme for more
// details.
//
// The function takes the following parameters:
//
//    - iconName: name of icon to use.
//
func (widget *Widget) DragSourceSetIconName(iconName string) {
	var _arg0 *C.GtkWidget // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_drag_source_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(iconName)
}

// DragSourceSetIconPixbuf sets the icon that will be used for drags from a
// particular widget from a Pixbuf. GTK+ retains a reference for pixbuf and will
// release it when it is no longer needed.
//
// The function takes the following parameters:
//
//    - pixbuf for the drag icon.
//
func (widget *Widget) DragSourceSetIconPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkWidget // out
	var _arg1 *C.GdkPixbuf // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(externglib.InternObject(pixbuf).Native()))

	C.gtk_drag_source_set_icon_pixbuf(_arg0, _arg1)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(pixbuf)
}

// DragSourceSetIconStock sets the icon that will be used for drags from a
// particular source to a stock icon.
//
// Deprecated: Use gtk_drag_source_set_icon_name() instead.
//
// The function takes the following parameters:
//
//    - stockId: ID of the stock icon to use.
//
func (widget *Widget) DragSourceSetIconStock(stockId string) {
	var _arg0 *C.GtkWidget // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_drag_source_set_icon_stock(_arg0, _arg1)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(stockId)
}

// DragSourceSetTargetList changes the target types that this widget offers for
// drag-and-drop. The widget must first be made into a drag source with
// gtk_drag_source_set().
//
// The function takes the following parameters:
//
//    - targetList (optional): list of draggable targets, or NULL for none.
//
func (widget *Widget) DragSourceSetTargetList(targetList *TargetList) {
	var _arg0 *C.GtkWidget     // out
	var _arg1 *C.GtkTargetList // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))
	if targetList != nil {
		_arg1 = (*C.GtkTargetList)(gextras.StructNative(unsafe.Pointer(targetList)))
	}

	C.gtk_drag_source_set_target_list(_arg0, _arg1)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(targetList)
}

// DragSourceUnset undoes the effects of gtk_drag_source_set().
func (widget *Widget) DragSourceUnset() {
	var _arg0 *C.GtkWidget // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))

	C.gtk_drag_source_unset(_arg0)
	runtime.KeepAlive(widget)
}
