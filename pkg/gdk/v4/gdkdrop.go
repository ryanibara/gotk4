// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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

	FinishDrop(action DragAction)

	Actions() DragAction

	Device() Device

	Display() Display

	Drag() Drag

	Formats() *ContentFormats

	Surface() Surface

	StatusDrop(actions DragAction, preferred DragAction)
}

// drop implements the Drop class.
type drop struct {
	gextras.Objector
}

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

func (s drop) FinishDrop(action DragAction) {
	var _arg0 *C.GdkDrop      // out
	var _arg1 C.GdkDragAction // out

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))
	_arg1 = C.GdkDragAction(action)

	C.gdk_drop_finish(_arg0, _arg1)
}

func (s drop) Actions() DragAction {
	var _arg0 *C.GdkDrop      // out
	var _cret C.GdkDragAction // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_actions(_arg0)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

func (s drop) Device() Device {
	var _arg0 *C.GdkDrop   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_device(_arg0)

	var _device Device // out

	_device = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Device)

	return _device
}

func (s drop) Display() Display {
	var _arg0 *C.GdkDrop    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Display)

	return _display
}

func (s drop) Drag() Drag {
	var _arg0 *C.GdkDrop // out
	var _cret *C.GdkDrag // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_drag(_arg0)

	var _drag Drag // out

	_drag = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Drag)

	return _drag
}

func (s drop) Formats() *ContentFormats {
	var _arg0 *C.GdkDrop           // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_formats(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(unsafe.Pointer(_cret))

	return _contentFormats
}

func (s drop) Surface() Surface {
	var _arg0 *C.GdkDrop    // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_surface(_arg0)

	var _surface Surface // out

	_surface = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Surface)

	return _surface
}

func (s drop) StatusDrop(actions DragAction, preferred DragAction) {
	var _arg0 *C.GdkDrop      // out
	var _arg1 C.GdkDragAction // out
	var _arg2 C.GdkDragAction // out

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))
	_arg1 = C.GdkDragAction(actions)
	_arg2 = C.GdkDragAction(preferred)

	C.gdk_drop_status(_arg0, _arg1, _arg2)
}
