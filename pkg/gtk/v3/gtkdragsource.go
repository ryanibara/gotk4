// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// DragSourceAddImageTargets: add the writable image targets supported by
// SelectionData to the target list of the drag source. The targets are added
// with info = 0. If you need another value, use
// gtk_target_list_add_image_targets() and gtk_drag_source_set_target_list().
func (widget *Widget) DragSourceAddImageTargets() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_source_add_image_targets", _args[:], nil)

	runtime.KeepAlive(widget)
}

// DragSourceAddTextTargets: add the text targets supported by SelectionData to
// the target list of the drag source. The targets are added with info = 0. If
// you need another value, use gtk_target_list_add_text_targets() and
// gtk_drag_source_set_target_list().
func (widget *Widget) DragSourceAddTextTargets() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_source_add_text_targets", _args[:], nil)

	runtime.KeepAlive(widget)
}

// DragSourceAddURITargets: add the URI targets supported by SelectionData to
// the target list of the drag source. The targets are added with info = 0. If
// you need another value, use gtk_target_list_add_uri_targets() and
// gtk_drag_source_set_target_list().
func (widget *Widget) DragSourceAddURITargets() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_source_add_uri_targets", _args[:], nil)

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
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_source_get_target_list", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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

// DragSourceSetIconGIcon sets the icon that will be used for drags from a
// particular source to icon. See the docs for IconTheme for more details.
//
// The function takes the following parameters:
//
//    - icon: #GIcon.
//
func (widget *Widget) DragSourceSetIconGIcon(icon gio.Iconner) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(icon).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_source_set_icon_gicon", _args[:], nil)

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
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_source_set_icon_name", _args[:], nil)

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
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_source_set_icon_pixbuf", _args[:], nil)

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
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg1))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_source_set_icon_stock", _args[:], nil)

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
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	if targetList != nil {
		_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(targetList)))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_source_set_target_list", _args[:], nil)

	runtime.KeepAlive(widget)
	runtime.KeepAlive(targetList)
}

// DragSourceUnset undoes the effects of gtk_drag_source_set().
func (widget *Widget) DragSourceUnset() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_source_unset", _args[:], nil)

	runtime.KeepAlive(widget)
}
