// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void callbackDelete(gpointer);
// extern GtkWidget* _gotk4_gtk3_ListBoxCreateWidgetFunc(gpointer, gpointer);
import "C"

// ListBoxCreateWidgetFunc: called for list boxes that are bound to a Model with
// gtk_list_box_bind_model() for each item that gets added to the model.
//
// Versions of GTK+ prior to 3.18 called gtk_widget_show_all() on the rows
// created by the GtkListBoxCreateWidgetFunc, but this forced all widgets inside
// the row to be shown, and is no longer the case. Applications should be
// updated to show the desired row widgets.
type ListBoxCreateWidgetFunc func(item *coreglib.Object) (widget Widgetter)

// BindModel binds model to box.
//
// If box was already bound to a model, that previous binding is destroyed.
//
// The contents of box are cleared and then filled with widgets that represent
// items from model. box is updated whenever model changes. If model is NULL,
// box is left empty.
//
// It is undefined to add or remove widgets directly (for example, with
// gtk_list_box_insert() or gtk_container_add()) while box is bound to a model.
//
// Note that using a model is incompatible with the filtering and sorting
// functionality in GtkListBox. When using a model, filtering and sorting should
// be implemented by the model.
//
// The function takes the following parameters:
//
//    - model (optional) to be bound to box.
//    - createWidgetFunc (optional): function that creates widgets for items or
//      NULL in case you also passed NULL as model.
//
func (box *ListBox) BindModel(model gio.ListModeller, createWidgetFunc ListBoxCreateWidgetFunc) {
	var _arg0 *C.GtkListBox                // out
	var _arg1 *C.GListModel                // out
	var _arg2 C.GtkListBoxCreateWidgetFunc // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}
	if createWidgetFunc != nil {
		_arg2 = (*[0]byte)(C._gotk4_gtk3_ListBoxCreateWidgetFunc)
		_arg3 = C.gpointer(gbox.Assign(createWidgetFunc))
		_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_list_box_bind_model(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(box)
	runtime.KeepAlive(model)
	runtime.KeepAlive(createWidgetFunc)
}
