// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gdk_drag_cancel_reason_get_type()), F: marshalDragCancelReason},
		{T: externglib.Type(C.gdk_drag_get_type()), F: marshalDrag},
	})
}

// DragCancelReason: used in `GdkDrag` to the reason of a cancelled DND
// operation.
type DragCancelReason int

const (
	// NoTarget: there is no suitable drop target.
	DragCancelReasonNoTarget DragCancelReason = 0
	// UserCancelled: drag cancelled by the user
	DragCancelReasonUserCancelled DragCancelReason = 1
	// error: unspecified error.
	DragCancelReasonError DragCancelReason = 2
)

func marshalDragCancelReason(p uintptr) (interface{}, error) {
	return DragCancelReason(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Drag: the `GdkDrag` object represents the source of an ongoing DND operation.
//
// A `GdkDrag` is created when a drag is started, and stays alive for duration
// of the DND operation. After a drag has been started with
// [func@Gdk.Drag.begin], the caller gets informed about the status of the
// ongoing drag operation with signals on the `GdkDrag` object.
//
// GTK provides a higher level abstraction based on top of these functions, and
// so they are not normally needed in GTK applications. See the "Drag and Drop"
// section of the GTK documentation for more information.
type Drag interface {
	gextras.Objector

	// DropDoneDrag informs GDK that the drop ended.
	//
	// Passing false for @success may trigger a drag cancellation animation.
	//
	// This function is called by the drag source, and should be the last call
	// before dropping the reference to the @drag.
	//
	// The `GdkDrag` will only take the first [method@Gdk.Drag.drop_done] call
	// as effective, if this function is called multiple times, all subsequent
	// calls will be ignored.
	DropDoneDrag(success bool)
	// Actions determines the bitmask of possible actions proposed by the
	// source.
	Actions() DragAction
	// Content returns the `GdkContentProvider` associated to the `GdkDrag`
	// object.
	Content() ContentProvider
	// Device returns the `GdkDevice` associated to the `GdkDrag` object.
	Device() Device
	// Display gets the `GdkDisplay` that the drag object was created for.
	Display() Display
	// DragSurface returns the surface on which the drag icon should be rendered
	// during the drag operation.
	//
	// Note that the surface may not be available until the drag operation has
	// begun. GDK will move the surface in accordance with the ongoing drag
	// operation. The surface is owned by @drag and will be destroyed when the
	// drag operation is over.
	DragSurface() Surface
	// Formats retrieves the formats supported by this `GdkDrag` object.
	Formats() ContentFormats
	// SelectedAction determines the action chosen by the drag destination.
	SelectedAction() DragAction
	// Surface returns the `GdkSurface` where the drag originates.
	Surface() Surface
	// SetHotspotDrag sets the position of the drag surface that will be kept
	// under the cursor hotspot.
	//
	// Initially, the hotspot is at the top left corner of the drag surface.
	SetHotspotDrag(hotX int, hotY int)
}

// drag implements the Drag class.
type drag struct {
	gextras.Objector
}

// WrapDrag wraps a GObject to the right type. It is
// primarily used internally.
func WrapDrag(obj *externglib.Object) Drag {
	return drag{
		Objector: obj,
	}
}

func marshalDrag(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDrag(obj), nil
}

func (d drag) DropDoneDrag(success bool) {
	var _arg0 *C.GdkDrag // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(d.Native()))
	if success {
		_arg1 = C.TRUE
	}

	C.gdk_drag_drop_done(_arg0, _arg1)
}

func (d drag) Actions() DragAction {
	var _arg0 *C.GdkDrag      // out
	var _cret C.GdkDragAction // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_drag_get_actions(_arg0)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

func (d drag) Content() ContentProvider {
	var _arg0 *C.GdkDrag            // out
	var _cret *C.GdkContentProvider // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_drag_get_content(_arg0)

	var _contentProvider ContentProvider // out

	_contentProvider = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ContentProvider)

	return _contentProvider
}

func (d drag) Device() Device {
	var _arg0 *C.GdkDrag   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_drag_get_device(_arg0)

	var _device Device // out

	_device = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Device)

	return _device
}

func (d drag) Display() Display {
	var _arg0 *C.GdkDrag    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_drag_get_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Display)

	return _display
}

func (d drag) DragSurface() Surface {
	var _arg0 *C.GdkDrag    // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_drag_get_drag_surface(_arg0)

	var _surface Surface // out

	_surface = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Surface)

	return _surface
}

func (d drag) Formats() ContentFormats {
	var _arg0 *C.GdkDrag           // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_drag_get_formats(_arg0)

	var _contentFormats ContentFormats // out

	_contentFormats = (ContentFormats)(unsafe.Pointer(_cret))
	C.gdk_content_formats_ref(_cret)

	return _contentFormats
}

func (d drag) SelectedAction() DragAction {
	var _arg0 *C.GdkDrag      // out
	var _cret C.GdkDragAction // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_drag_get_selected_action(_arg0)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

func (d drag) Surface() Surface {
	var _arg0 *C.GdkDrag    // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_drag_get_surface(_arg0)

	var _surface Surface // out

	_surface = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Surface)

	return _surface
}

func (d drag) SetHotspotDrag(hotX int, hotY int) {
	var _arg0 *C.GdkDrag // out
	var _arg1 C.int      // out
	var _arg2 C.int      // out

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(d.Native()))
	_arg1 = C.int(hotX)
	_arg2 = C.int(hotY)

	C.gdk_drag_set_hotspot(_arg0, _arg1, _arg2)
}
