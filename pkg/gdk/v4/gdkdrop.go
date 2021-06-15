// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_drop_get_type()), F: marshalDrop},
	})
}

// Drop: the `GdkDrop` object represents the target of an ongoing DND operation.
//
// Possible drop sites get informed about the status of the ongoing drag
// operation with events of type GDK_DRAG_ENTER, GDK_DRAG_LEAVE, GDK_DRAG_MOTION
// and GDK_DROP_START. The `GdkDrop` object can be obtained from these
// [class@Gdk.Event] types using [method@Gdk.DNDEvent.get_drop].
//
// The actual data transfer is initiated from the target side via an async read,
// using one of the `GdkDrop` methods for this purpose:
// [method@Gdk.Drop.read_async] or [method@Gdk.Drop.read_value_async].
//
// GTK provides a higher level abstraction based on top of these functions, and
// so they are not normally needed in GTK applications. See the "Drag and Drop"
// section of the GTK documentation for more information.
type Drop interface {
	gextras.Objector

	// Finish ends the drag operation after a drop.
	//
	// The @action must be a single action selected from the actions available
	// via [method@Gdk.Drop.get_actions].
	Finish(action DragAction)
	// Actions returns the possible actions for this `GdkDrop`.
	//
	// If this value contains multiple actions - i.e.
	// [func@Gdk.DragAction.is_unique] returns false for the result -
	// [method@Gdk.Drop.finish] must choose the action to use when accepting the
	// drop. This will only happen if you passed GDK_ACTION_ASK as one of the
	// possible actions in [method@Gdk.Drop.status]. GDK_ACTION_ASK itself will
	// not be included in the actions returned by this function.
	//
	// This value may change over the lifetime of the [class@Gdk.Drop] both as a
	// response to source side actions as well as to calls to
	// [method@Gdk.Drop.status] or [method@Gdk.Drop.finish]. The source side
	// will not change this value anymore once a drop has started.
	Actions() DragAction
	// Device returns the `GdkDevice` performing the drop.
	Device() Device
	// Display gets the `GdkDisplay` that @self was created for.
	Display() Display
	// Drag: if this is an in-app drag-and-drop operation, returns the `GdkDrag`
	// that corresponds to this drop.
	//
	// If it is not, nil is returned.
	Drag() Drag
	// Formats returns the `GdkContentFormats` that the drop offers the data to
	// be read in.
	Formats() *ContentFormats
	// Surface returns the `GdkSurface` performing the drop.
	Surface() Surface
	// ReadFinish finishes an async drop read operation.
	//
	// Note that you must not use blocking read calls on the returned stream in
	// the GTK thread, since some platforms might require communication with GTK
	// to complete the data transfer. You can use async APIs such as
	// g_input_stream_read_bytes_async().
	//
	// See [method@Gdk.Drop.read_async].
	ReadFinish(result gio.AsyncResult) (string, gio.InputStream, error)
	// ReadValueFinish finishes an async drop read.
	//
	// See [method@Gdk.Drop.read_value_async].
	ReadValueFinish(result gio.AsyncResult) (**externglib.Value, error)
	// Status selects all actions that are potentially supported by the
	// destination.
	//
	// When calling this function, do not restrict the passed in actions to the
	// ones provided by [method@Gdk.Drop.get_actions]. Those actions may change
	// in the future, even depending on the actions you provide here.
	//
	// The @preferred action is a hint to the drag'n'drop mechanism about which
	// action to use when multiple actions are possible.
	//
	// This function should be called by drag destinations in response to
	// GDK_DRAG_ENTER or GDK_DRAG_MOTION events. If the destination does not yet
	// know the exact actions it supports, it should set any possible actions
	// first and then later call this function again.
	Status(actions DragAction, preferred DragAction)
}

// drop implements the Drop class.
type drop struct {
	gextras.Objector
}

var _ Drop = (*drop)(nil)

// WrapDrop wraps a GObject to the right type. It is
// primarily used internally.
func WrapDrop(obj *externglib.Object) Drop {
	return drop{
		Objector: obj,
	}
}

func marshalDrop(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDrop(obj), nil
}

// Finish ends the drag operation after a drop.
//
// The @action must be a single action selected from the actions available
// via [method@Gdk.Drop.get_actions].
func (s drop) Finish(action DragAction) {
	var _arg0 *C.GdkDrop      // out
	var _arg1 C.GdkDragAction // out

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GdkDragAction)(action)

	C.gdk_drop_finish(_arg0, _arg1)
}

// Actions returns the possible actions for this `GdkDrop`.
//
// If this value contains multiple actions - i.e.
// [func@Gdk.DragAction.is_unique] returns false for the result -
// [method@Gdk.Drop.finish] must choose the action to use when accepting the
// drop. This will only happen if you passed GDK_ACTION_ASK as one of the
// possible actions in [method@Gdk.Drop.status]. GDK_ACTION_ASK itself will
// not be included in the actions returned by this function.
//
// This value may change over the lifetime of the [class@Gdk.Drop] both as a
// response to source side actions as well as to calls to
// [method@Gdk.Drop.status] or [method@Gdk.Drop.finish]. The source side
// will not change this value anymore once a drop has started.
func (s drop) Actions() DragAction {
	var _arg0 *C.GdkDrop      // out
	var _cret C.GdkDragAction // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_actions(_arg0)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

// Device returns the `GdkDevice` performing the drop.
func (s drop) Device() Device {
	var _arg0 *C.GdkDrop   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_device(_arg0)

	var _device Device // out

	_device = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Device)

	return _device
}

// Display gets the `GdkDisplay` that @self was created for.
func (s drop) Display() Display {
	var _arg0 *C.GdkDrop    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Display)

	return _display
}

// Drag: if this is an in-app drag-and-drop operation, returns the `GdkDrag`
// that corresponds to this drop.
//
// If it is not, nil is returned.
func (s drop) Drag() Drag {
	var _arg0 *C.GdkDrop // out
	var _cret *C.GdkDrag // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_drag(_arg0)

	var _drag Drag // out

	_drag = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Drag)

	return _drag
}

// Formats returns the `GdkContentFormats` that the drop offers the data to
// be read in.
func (s drop) Formats() *ContentFormats {
	var _arg0 *C.GdkDrop           // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_formats(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = WrapContentFormats(unsafe.Pointer(_cret))

	return _contentFormats
}

// Surface returns the `GdkSurface` performing the drop.
func (s drop) Surface() Surface {
	var _arg0 *C.GdkDrop    // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_surface(_arg0)

	var _surface Surface // out

	_surface = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Surface)

	return _surface
}

// ReadFinish finishes an async drop read operation.
//
// Note that you must not use blocking read calls on the returned stream in
// the GTK thread, since some platforms might require communication with GTK
// to complete the data transfer. You can use async APIs such as
// g_input_stream_read_bytes_async().
//
// See [method@Gdk.Drop.read_async].
func (s drop) ReadFinish(result gio.AsyncResult) (string, gio.InputStream, error) {
	var _arg0 *C.GdkDrop      // out
	var _arg1 *C.GAsyncResult // out
	var _arg2 *C.char         // in
	var _cret *C.GInputStream // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.gdk_drop_read_finish(_arg0, _arg1, &_arg2, &_cerr)

	var _outMimeType string          // out
	var _inputStream gio.InputStream // out
	var _goerr error                 // out

	_outMimeType = C.GoString(_arg2)
	defer C.free(unsafe.Pointer(_arg2))
	_inputStream = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(gio.InputStream)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _outMimeType, _inputStream, _goerr
}

// ReadValueFinish finishes an async drop read.
//
// See [method@Gdk.Drop.read_value_async].
func (s drop) ReadValueFinish(result gio.AsyncResult) (**externglib.Value, error) {
	var _arg0 *C.GdkDrop      // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GValue       // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.gdk_drop_read_value_finish(_arg0, _arg1, &_cerr)

	var _value **externglib.Value // out
	var _goerr error              // out

	_value = externglib.ValueFromNative(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _value, _goerr
}

// Status selects all actions that are potentially supported by the
// destination.
//
// When calling this function, do not restrict the passed in actions to the
// ones provided by [method@Gdk.Drop.get_actions]. Those actions may change
// in the future, even depending on the actions you provide here.
//
// The @preferred action is a hint to the drag'n'drop mechanism about which
// action to use when multiple actions are possible.
//
// This function should be called by drag destinations in response to
// GDK_DRAG_ENTER or GDK_DRAG_MOTION events. If the destination does not yet
// know the exact actions it supports, it should set any possible actions
// first and then later call this function again.
func (s drop) Status(actions DragAction, preferred DragAction) {
	var _arg0 *C.GdkDrop      // out
	var _arg1 C.GdkDragAction // out
	var _arg2 C.GdkDragAction // out

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GdkDragAction)(actions)
	_arg2 = (C.GdkDragAction)(preferred)

	C.gdk_drop_status(_arg0, _arg1, _arg2)
}
