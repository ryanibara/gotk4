// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeDragCancelReason returns the GType for the type DragCancelReason.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDragCancelReason() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "DragCancelReason").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDragCancelReason)
	return gtype
}

// GTypeDragProtocol returns the GType for the type DragProtocol.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDragProtocol() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "DragProtocol").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDragProtocol)
	return gtype
}

// GTypeDragAction returns the GType for the type DragAction.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDragAction() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "DragAction").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDragAction)
	return gtype
}

// DragCancelReason: used in DragContext to the reason of a cancelled DND
// operation.
type DragCancelReason C.gint

const (
	// DragCancelNoTarget: there is no suitable drop target.
	DragCancelNoTarget DragCancelReason = iota
	// DragCancelUserCancelled: drag cancelled by the user.
	DragCancelUserCancelled
	// DragCancelError: unspecified error.
	DragCancelError
)

func marshalDragCancelReason(p uintptr) (interface{}, error) {
	return DragCancelReason(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for DragCancelReason.
func (d DragCancelReason) String() string {
	switch d {
	case DragCancelNoTarget:
		return "NoTarget"
	case DragCancelUserCancelled:
		return "UserCancelled"
	case DragCancelError:
		return "Error"
	default:
		return fmt.Sprintf("DragCancelReason(%d)", d)
	}
}

// DragProtocol: used in DragContext to indicate the protocol according to which
// DND is done.
type DragProtocol C.gint

const (
	// DragProtoNone: no protocol.
	DragProtoNone DragProtocol = iota
	// DragProtoMotif: motif DND protocol. No longer supported.
	DragProtoMotif
	// DragProtoXdnd: xdnd protocol.
	DragProtoXdnd
	// DragProtoRootwin: extension to the Xdnd protocol for unclaimed root
	// window drops.
	DragProtoRootwin
	// DragProtoWin32Dropfiles: simple WM_DROPFILES protocol.
	DragProtoWin32Dropfiles
	// DragProtoOle2: complex OLE2 DND protocol (not implemented).
	DragProtoOle2
	// DragProtoLocal: intra-application DND.
	DragProtoLocal
	// DragProtoWayland: wayland DND protocol.
	DragProtoWayland
)

func marshalDragProtocol(p uintptr) (interface{}, error) {
	return DragProtocol(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for DragProtocol.
func (d DragProtocol) String() string {
	switch d {
	case DragProtoNone:
		return "None"
	case DragProtoMotif:
		return "Motif"
	case DragProtoXdnd:
		return "Xdnd"
	case DragProtoRootwin:
		return "Rootwin"
	case DragProtoWin32Dropfiles:
		return "Win32Dropfiles"
	case DragProtoOle2:
		return "Ole2"
	case DragProtoLocal:
		return "Local"
	case DragProtoWayland:
		return "Wayland"
	default:
		return fmt.Sprintf("DragProtocol(%d)", d)
	}
}

// DragAction: used in DragContext to indicate what the destination should do
// with the dropped data.
type DragAction C.guint

const (
	// ActionDefault means nothing, and should not be used.
	ActionDefault DragAction = 0b1
	// ActionCopy: copy the data.
	ActionCopy DragAction = 0b10
	// ActionMove: move the data, i.e. first copy it, then delete it from the
	// source using the DELETE target of the X selection protocol.
	ActionMove DragAction = 0b100
	// ActionLink: add a link to the data. Note that this is only useful if
	// source and destination agree on what it means.
	ActionLink DragAction = 0b1000
	// ActionPrivate: special action which tells the source that the destination
	// will do something that the source doesn’t understand.
	ActionPrivate DragAction = 0b10000
	// ActionAsk: ask the user what to do with the data.
	ActionAsk DragAction = 0b100000
)

func marshalDragAction(p uintptr) (interface{}, error) {
	return DragAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for DragAction.
func (d DragAction) String() string {
	if d == 0 {
		return "DragAction(0)"
	}

	var builder strings.Builder
	builder.Grow(70)

	for d != 0 {
		next := d & (d - 1)
		bit := d - next

		switch bit {
		case ActionDefault:
			builder.WriteString("Default|")
		case ActionCopy:
			builder.WriteString("Copy|")
		case ActionMove:
			builder.WriteString("Move|")
		case ActionLink:
			builder.WriteString("Link|")
		case ActionPrivate:
			builder.WriteString("Private|")
		case ActionAsk:
			builder.WriteString("Ask|")
		default:
			builder.WriteString(fmt.Sprintf("DragAction(0b%b)|", bit))
		}

		d = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if d contains other.
func (d DragAction) Has(other DragAction) bool {
	return (d & other) == other
}

// DragAbort aborts a drag without dropping.
//
// This function is called by the drag source.
//
// This function does not need to be called in managed drag and drop operations.
// See gdk_drag_context_manage_dnd() for more information.
//
// The function takes the following parameters:
//
//    - context: DragContext.
//    - time_: timestamp for this operation.
//
func DragAbort(context *DragContext, time_ uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(*C.guint32)(unsafe.Pointer(&_args[1])) = C.guint32(time_)

	girepository.MustFind("Gdk", "drag_abort").Invoke(_args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(time_)
}

// DragDrop drops on the current destination.
//
// This function is called by the drag source.
//
// This function does not need to be called in managed drag and drop operations.
// See gdk_drag_context_manage_dnd() for more information.
//
// The function takes the following parameters:
//
//    - context: DragContext.
//    - time_: timestamp for this operation.
//
func DragDrop(context *DragContext, time_ uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(*C.guint32)(unsafe.Pointer(&_args[1])) = C.guint32(time_)

	girepository.MustFind("Gdk", "drag_drop").Invoke(_args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(time_)
}

// DragDropDone: inform GDK if the drop ended successfully. Passing FALSE for
// success may trigger a drag cancellation animation.
//
// This function is called by the drag source, and should be the last call
// before dropping the reference to the context.
//
// The DragContext will only take the first gdk_drag_drop_done() call as
// effective, if this function is called multiple times, all subsequent calls
// will be ignored.
//
// The function takes the following parameters:
//
//    - context: DragContext.
//    - success: whether the drag was ultimatively successful.
//
func DragDropDone(context *DragContext, success bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	if success {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gdk", "drag_drop_done").Invoke(_args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(success)
}

// DragDropSucceeded returns whether the dropped data has been successfully
// transferred. This function is intended to be used while handling a
// GDK_DROP_FINISHED event, its return value is meaningless at other times.
//
// The function takes the following parameters:
//
//    - context: DragContext.
//
// The function returns the following values:
//
//    - ok: TRUE if the drop was successful.
//
func DragDropSucceeded(context *DragContext) bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gdk", "drag_drop_succeeded").Invoke(_args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// DragFindWindowForScreen finds the destination window and DND protocol to use
// at the given pointer position.
//
// This function is called by the drag source to obtain the dest_window and
// protocol parameters for gdk_drag_motion().
//
// The function takes the following parameters:
//
//    - context: DragContext.
//    - dragWindow: window which may be at the pointer position, but should be
//      ignored, since it is put up by the drag source as an icon.
//    - screen where the destination window is sought.
//    - xRoot: x position of the pointer in root coordinates.
//    - yRoot: y position of the pointer in root coordinates.
//
// The function returns the following values:
//
//    - destWindow: location to store the destination window in.
//    - protocol: location to store the DND protocol in.
//
func DragFindWindowForScreen(context *DragContext, dragWindow Windower, screen *Screen, xRoot, yRoot int32) (Windower, DragProtocol) {
	var _args [5]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(dragWindow).Native()))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[3])) = C.gint(xRoot)
	*(*C.gint)(unsafe.Pointer(&_args[4])) = C.gint(yRoot)

	girepository.MustFind("Gdk", "drag_find_window_for_screen").Invoke(_args[:], _outs[:])

	runtime.KeepAlive(context)
	runtime.KeepAlive(dragWindow)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(xRoot)
	runtime.KeepAlive(yRoot)

	var _destWindow Windower   // out
	var _protocol DragProtocol // out

	{
		objptr := unsafe.Pointer(_outs[0])
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Windower)
			return ok
		})
		rv, ok := casted.(Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_destWindow = rv
	}
	_protocol = *(*DragProtocol)(unsafe.Pointer(_outs[1]))

	return _destWindow, _protocol
}

// DropFinish ends the drag operation after a drop.
//
// This function is called by the drag destination.
//
// The function takes the following parameters:
//
//    - context: DragContext.
//    - success: TRUE if the data was successfully received.
//    - time_: timestamp for this operation.
//
func DropFinish(context *DragContext, success bool, time_ uint32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	if success {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}
	*(*C.guint32)(unsafe.Pointer(&_args[2])) = C.guint32(time_)

	girepository.MustFind("Gdk", "drop_finish").Invoke(_args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(success)
	runtime.KeepAlive(time_)
}

// DropReply accepts or rejects a drop.
//
// This function is called by the drag destination in response to a drop
// initiated by the drag source.
//
// The function takes the following parameters:
//
//    - context: DragContext.
//    - accepted: TRUE if the drop is accepted.
//    - time_: timestamp for this operation.
//
func DropReply(context *DragContext, accepted bool, time_ uint32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	if accepted {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}
	*(*C.guint32)(unsafe.Pointer(&_args[2])) = C.guint32(time_)

	girepository.MustFind("Gdk", "drop_reply").Invoke(_args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(accepted)
	runtime.KeepAlive(time_)
}

// DestWindow returns the destination window for the DND operation.
//
// The function returns the following values:
//
//    - window: Window.
//
func (context *DragContext) DestWindow() Windower {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gdk", "DragContext").InvokeMethod("get_dest_window", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _window Windower // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Windower)
			return ok
		})
		rv, ok := casted.(Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_window = rv
	}

	return _window
}

// Device returns the Device associated to the drag context.
//
// The function returns the following values:
//
//    - device associated to context.
//
func (context *DragContext) Device() Devicer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gdk", "DragContext").InvokeMethod("get_device", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _device Devicer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Devicer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Devicer)
			return ok
		})
		rv, ok := casted.(Devicer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
		}
		_device = rv
	}

	return _device
}

// DragWindow returns the window on which the drag icon should be rendered
// during the drag operation. Note that the window may not be available until
// the drag operation has begun. GDK will move the window in accordance with the
// ongoing drag operation. The window is owned by context and will be destroyed
// when the drag operation is over.
//
// The function returns the following values:
//
//    - window (optional): drag window, or NULL.
//
func (context *DragContext) DragWindow() Windower {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gdk", "DragContext").InvokeMethod("get_drag_window", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _window Windower // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Windower)
				return ok
			})
			rv, ok := casted.(Windower)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
			}
			_window = rv
		}
	}

	return _window
}

// SourceWindow returns the Window where the DND operation started.
//
// The function returns the following values:
//
//    - window: Window.
//
func (context *DragContext) SourceWindow() Windower {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gdk", "DragContext").InvokeMethod("get_source_window", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _window Windower // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Windower)
			return ok
		})
		rv, ok := casted.(Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_window = rv
	}

	return _window
}

// SetDevice associates a Device to context, so all Drag and Drop events for
// context are emitted as if they came from this device.
//
// The function takes the following parameters:
//
//    - device: Device.
//
func (context *DragContext) SetDevice(device Devicer) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	girepository.MustFind("Gdk", "DragContext").InvokeMethod("set_device", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(device)
}

// SetHotspot sets the position of the drag window that will be kept under the
// cursor hotspot. Initially, the hotspot is at the top left corner of the drag
// window.
//
// The function takes the following parameters:
//
//    - hotX: x coordinate of the drag window hotspot.
//    - hotY: y coordinate of the drag window hotspot.
//
func (context *DragContext) SetHotspot(hotX, hotY int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(hotX)
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(hotY)

	girepository.MustFind("Gdk", "DragContext").InvokeMethod("set_hotspot", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(hotX)
	runtime.KeepAlive(hotY)
}
