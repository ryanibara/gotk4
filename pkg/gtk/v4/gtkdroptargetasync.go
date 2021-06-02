// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_drop_target_async_get_type()), F: marshalDropTargetAsync},
	})
}

// DropTargetAsync: `GtkDropTargetAsync` is an event controller to receive
// Drag-and-Drop operations, asynchronously.
//
// It is the more complete but also more complex method of handling drop
// operations compared to [class@Gtk.DropTarget], and you should only use it if
// `GtkDropTarget` doesn't provide all the features you need.
//
// To use a `GtkDropTargetAsync` to receive drops on a widget, you create a
// `GtkDropTargetAsync` object, configure which data formats and actions you
// support, connect to its signals, and then attach it to the widget with
// [method@Gtk.Widget.add_controller].
//
// During a drag operation, the first signal that a `GtkDropTargetAsync` emits
// is [signal@Gtk.DropTargetAsync::accept], which is meant to determine whether
// the target is a possible drop site for the ongoing drop. The default handler
// for the ::accept signal accepts the drop if it finds a compatible data format
// and an action that is supported on both sides.
//
// If it is, and the widget becomes a target, you will receive a
// [signal@Gtk.DropTargetAsync::drag-enter] signal, followed by
// [signal@Gtk.DropTargetAsync::drag-motion] signals as the pointer moves,
// optionally a [signal@Gtk.DropTargetAsync::drop] signal when a drop happens,
// and finally a [signal@Gtk.DropTargetAsync::drag-leave] signal when the
// pointer moves off the widget.
//
// The ::drag-enter and ::drag-motion handler return a `GdkDragAction` to update
// the status of the ongoing operation. The ::drop handler should decide if it
// ultimately accepts the drop and if it does, it should initiate the data
// transfer and finish the operation by calling [method@Gdk.Drop.finish].
//
// Between the ::drag-enter and ::drag-leave signals the widget is a current
// drop target, and will receive the GTK_STATE_FLAG_DROP_ACTIVE state, which can
// be used by themes to style the widget as a drop target.
type DropTargetAsync interface {
	EventController

	// Actions gets the actions that this drop target supports.
	Actions() gdk.DragAction
	// Formats gets the data formats that this drop target accepts.
	//
	// If the result is nil, all formats are expected to be supported.
	Formats() *gdk.ContentFormats
	// RejectDrop sets the @drop as not accepted on this drag site.
	//
	// This function should be used when delaying the decision on whether to
	// accept a drag or not until after reading the data.
	RejectDrop(drop gdk.Drop)
	// SetActions sets the actions that this drop target supports.
	SetActions(actions gdk.DragAction)
	// SetFormats sets the data formats that this drop target will accept.
	SetFormats(formats *gdk.ContentFormats)
}

// dropTargetAsync implements the DropTargetAsync interface.
type dropTargetAsync struct {
	EventController
}

var _ DropTargetAsync = (*dropTargetAsync)(nil)

// WrapDropTargetAsync wraps a GObject to the right type. It is
// primarily used internally.
func WrapDropTargetAsync(obj *externglib.Object) DropTargetAsync {
	return DropTargetAsync{
		EventController: WrapEventController(obj),
	}
}

func marshalDropTargetAsync(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDropTargetAsync(obj), nil
}

// NewDropTargetAsync constructs a class DropTargetAsync.
func NewDropTargetAsync(formats *gdk.ContentFormats, actions gdk.DragAction) DropTargetAsync {
	var arg1 *C.GdkContentFormats
	var arg2 C.GdkDragAction

	arg1 = (*C.GdkContentFormats)(formats.Native())
	arg2 = (C.GdkDragAction)(actions)

	ret := C.gtk_drop_target_async_new(arg1, arg2)

	var ret0 DropTargetAsync

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(DropTargetAsync)

	return ret0
}

// Actions gets the actions that this drop target supports.
func (s dropTargetAsync) Actions() gdk.DragAction {
	var arg0 *C.GtkDropTargetAsync

	arg0 = (*C.GtkDropTargetAsync)(s.Native())

	ret := C.gtk_drop_target_async_get_actions(arg0)

	var ret0 gdk.DragAction

	ret0 = gdk.DragAction(ret)

	return ret0
}

// Formats gets the data formats that this drop target accepts.
//
// If the result is nil, all formats are expected to be supported.
func (s dropTargetAsync) Formats() *gdk.ContentFormats {
	var arg0 *C.GtkDropTargetAsync

	arg0 = (*C.GtkDropTargetAsync)(s.Native())

	ret := C.gtk_drop_target_async_get_formats(arg0)

	var ret0 *gdk.ContentFormats

	{
		ret0 = gdk.WrapContentFormats(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *gdk.ContentFormats) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// RejectDrop sets the @drop as not accepted on this drag site.
//
// This function should be used when delaying the decision on whether to
// accept a drag or not until after reading the data.
func (s dropTargetAsync) RejectDrop(drop gdk.Drop) {
	var arg0 *C.GtkDropTargetAsync
	var arg1 *C.GdkDrop

	arg0 = (*C.GtkDropTargetAsync)(s.Native())
	arg1 = (*C.GdkDrop)(drop.Native())

	C.gtk_drop_target_async_reject_drop(arg0, arg1)
}

// SetActions sets the actions that this drop target supports.
func (s dropTargetAsync) SetActions(actions gdk.DragAction) {
	var arg0 *C.GtkDropTargetAsync
	var arg1 C.GdkDragAction

	arg0 = (*C.GtkDropTargetAsync)(s.Native())
	arg1 = (C.GdkDragAction)(actions)

	C.gtk_drop_target_async_set_actions(arg0, arg1)
}

// SetFormats sets the data formats that this drop target will accept.
func (s dropTargetAsync) SetFormats(formats *gdk.ContentFormats) {
	var arg0 *C.GtkDropTargetAsync
	var arg1 *C.GdkContentFormats

	arg0 = (*C.GtkDropTargetAsync)(s.Native())
	arg1 = (*C.GdkContentFormats)(formats.Native())

	C.gtk_drop_target_async_set_formats(arg0, arg1)
}
