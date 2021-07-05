// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_status_get_type()), F: marshalStatus},
		{T: externglib.Type(C.gdk_device_tool_get_type()), F: marshalDeviceTool},
		{T: externglib.Type(C.gdk_drag_context_get_type()), F: marshalDragContext},
	})
}

type Status int

const (
	StatusOk         Status = 0
	StatusError      Status = -1
	StatusErrorParam Status = -2
	StatusErrorFile  Status = -3
	StatusErrorMem   Status = -4
)

func marshalStatus(p uintptr) (interface{}, error) {
	return Status(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

type DeviceTool interface {
	gextras.Objector

	// HardwareID gets the hardware ID of this tool, or 0 if it's not known.
	// When non-zero, the identificator is unique for the given tool model,
	// meaning that two identical tools will share the same @hardware_id, but
	// will have different serial numbers (see gdk_device_tool_get_serial()).
	//
	// This is a more concrete (and device specific) method to identify a
	// DeviceTool than gdk_device_tool_get_tool_type(), as a tablet may support
	// multiple devices with the same DeviceToolType, but having different
	// hardware identificators.
	HardwareID() uint64
	// Serial gets the serial of this tool, this value can be used to identify a
	// physical tool (eg. a tablet pen) across program executions.
	Serial() uint64
	// ToolType gets the DeviceToolType of the tool.
	ToolType() DeviceToolType
}

// deviceTool implements the DeviceTool class.
type deviceTool struct {
	gextras.Objector
}

// WrapDeviceTool wraps a GObject to the right type. It is
// primarily used internally.
func WrapDeviceTool(obj *externglib.Object) DeviceTool {
	return deviceTool{
		Objector: obj,
	}
}

func marshalDeviceTool(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDeviceTool(obj), nil
}

func (t deviceTool) HardwareID() uint64 {
	var _arg0 *C.GdkDeviceTool // out
	var _cret C.guint64        // in

	_arg0 = (*C.GdkDeviceTool)(unsafe.Pointer(t.Native()))

	_cret = C.gdk_device_tool_get_hardware_id(_arg0)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

func (t deviceTool) Serial() uint64 {
	var _arg0 *C.GdkDeviceTool // out
	var _cret C.guint64        // in

	_arg0 = (*C.GdkDeviceTool)(unsafe.Pointer(t.Native()))

	_cret = C.gdk_device_tool_get_serial(_arg0)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

func (t deviceTool) ToolType() DeviceToolType {
	var _arg0 *C.GdkDeviceTool    // out
	var _cret C.GdkDeviceToolType // in

	_arg0 = (*C.GdkDeviceTool)(unsafe.Pointer(t.Native()))

	_cret = C.gdk_device_tool_get_tool_type(_arg0)

	var _deviceToolType DeviceToolType // out

	_deviceToolType = DeviceToolType(_cret)

	return _deviceToolType
}

type DragContext interface {
	gextras.Objector

	// Actions determines the bitmask of actions proposed by the source if
	// gdk_drag_context_get_suggested_action() returns GDK_ACTION_ASK.
	Actions() DragAction
	// DestWindow returns the destination window for the DND operation.
	DestWindow() Window
	// Device returns the Device associated to the drag context.
	Device() Device
	// DragWindow returns the window on which the drag icon should be rendered
	// during the drag operation. Note that the window may not be available
	// until the drag operation has begun. GDK will move the window in
	// accordance with the ongoing drag operation. The window is owned by
	// @context and will be destroyed when the drag operation is over.
	DragWindow() Window
	// Protocol returns the drag protocol that is used by this context.
	Protocol() DragProtocol
	// SelectedAction determines the action chosen by the drag destination.
	SelectedAction() DragAction
	// SourceWindow returns the Window where the DND operation started.
	SourceWindow() Window
	// SuggestedAction determines the suggested drag action of the context.
	SuggestedAction() DragAction
	// ManageDndDragContext requests the drag and drop operation to be managed
	// by @context. When a drag and drop operation becomes managed, the
	// DragContext will internally handle all input and source-side EventDND
	// events as required by the windowing system.
	//
	// Once the drag and drop operation is managed, the drag context will emit
	// the following signals: - The DragContext::action-changed signal whenever
	// the final action to be performed by the drag and drop operation changes.
	// - The DragContext::drop-performed signal after the user performs the drag
	// and drop gesture (typically by releasing the mouse button). - The
	// DragContext::dnd-finished signal after the drag and drop operation
	// concludes (after all Selection transfers happen). - The
	// DragContext::cancel signal if the drag and drop operation is finished but
	// doesn't happen over an accepting destination, or is cancelled through
	// other means.
	ManageDndDragContext(ipcWindow Window, actions DragAction) bool
	// SetDeviceDragContext associates a Device to @context, so all Drag and
	// Drop events for @context are emitted as if they came from this device.
	SetDeviceDragContext(device Device)
	// SetHotspotDragContext sets the position of the drag window that will be
	// kept under the cursor hotspot. Initially, the hotspot is at the top left
	// corner of the drag window.
	SetHotspotDragContext(hotX int, hotY int)
}

// dragContext implements the DragContext class.
type dragContext struct {
	gextras.Objector
}

// WrapDragContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapDragContext(obj *externglib.Object) DragContext {
	return dragContext{
		Objector: obj,
	}
}

func marshalDragContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDragContext(obj), nil
}

func (c dragContext) Actions() DragAction {
	var _arg0 *C.GdkDragContext // out
	var _cret C.GdkDragAction   // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_drag_context_get_actions(_arg0)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

func (c dragContext) DestWindow() Window {
	var _arg0 *C.GdkDragContext // out
	var _cret *C.GdkWindow      // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_drag_context_get_dest_window(_arg0)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _window
}

func (c dragContext) Device() Device {
	var _arg0 *C.GdkDragContext // out
	var _cret *C.GdkDevice      // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_drag_context_get_device(_arg0)

	var _device Device // out

	_device = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Device)

	return _device
}

func (c dragContext) DragWindow() Window {
	var _arg0 *C.GdkDragContext // out
	var _cret *C.GdkWindow      // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_drag_context_get_drag_window(_arg0)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _window
}

func (c dragContext) Protocol() DragProtocol {
	var _arg0 *C.GdkDragContext // out
	var _cret C.GdkDragProtocol // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_drag_context_get_protocol(_arg0)

	var _dragProtocol DragProtocol // out

	_dragProtocol = DragProtocol(_cret)

	return _dragProtocol
}

func (c dragContext) SelectedAction() DragAction {
	var _arg0 *C.GdkDragContext // out
	var _cret C.GdkDragAction   // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_drag_context_get_selected_action(_arg0)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

func (c dragContext) SourceWindow() Window {
	var _arg0 *C.GdkDragContext // out
	var _cret *C.GdkWindow      // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_drag_context_get_source_window(_arg0)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _window
}

func (c dragContext) SuggestedAction() DragAction {
	var _arg0 *C.GdkDragContext // out
	var _cret C.GdkDragAction   // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_drag_context_get_suggested_action(_arg0)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

func (c dragContext) ManageDndDragContext(ipcWindow Window, actions DragAction) bool {
	var _arg0 *C.GdkDragContext // out
	var _arg1 *C.GdkWindow      // out
	var _arg2 C.GdkDragAction   // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(ipcWindow.Native()))
	_arg2 = C.GdkDragAction(actions)

	_cret = C.gdk_drag_context_manage_dnd(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c dragContext) SetDeviceDragContext(device Device) {
	var _arg0 *C.GdkDragContext // out
	var _arg1 *C.GdkDevice      // out

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	C.gdk_drag_context_set_device(_arg0, _arg1)
}

func (c dragContext) SetHotspotDragContext(hotX int, hotY int) {
	var _arg0 *C.GdkDragContext // out
	var _arg1 C.gint            // out
	var _arg2 C.gint            // out

	_arg0 = (*C.GdkDragContext)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(hotX)
	_arg2 = C.gint(hotY)

	C.gdk_drag_context_set_hotspot(_arg0, _arg1, _arg2)
}
