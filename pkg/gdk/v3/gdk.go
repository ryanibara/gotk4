// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_status_get_type()), F: marshalStatus},
		{T: externglib.Type(C.gdk_device_tool_get_type()), F: marshalDeviceTooler},
		{T: externglib.Type(C.gdk_drag_context_get_type()), F: marshalDragContexter},
	})
}

// The function returns the following values:
//
func GLErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gdk_gl_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

type Status C.gint

const (
	OK         Status = 0
	Error      Status = -1
	ErrorParam Status = -2
	ErrorFile  Status = -3
	ErrorMem   Status = -4
)

func marshalStatus(p uintptr) (interface{}, error) {
	return Status(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for Status.
func (s Status) String() string {
	switch s {
	case OK:
		return "OK"
	case Error:
		return "Error"
	case ErrorParam:
		return "ErrorParam"
	case ErrorFile:
		return "ErrorFile"
	case ErrorMem:
		return "ErrorMem"
	default:
		return fmt.Sprintf("Status(%d)", s)
	}
}

type DeviceTool struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*DeviceTool)(nil)
)

func wrapDeviceTool(obj *externglib.Object) *DeviceTool {
	return &DeviceTool{
		Object: obj,
	}
}

func marshalDeviceTooler(p uintptr) (interface{}, error) {
	return wrapDeviceTool(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

type DragContext struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*DragContext)(nil)
)

func wrapDragContext(obj *externglib.Object) *DragContext {
	return &DragContext{
		Object: obj,
	}
}

func marshalDragContexter(p uintptr) (interface{}, error) {
	return wrapDragContext(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectActionChanged: new action is being chosen for the drag and drop
// operation.
//
// This signal will only be emitted if the DragContext manages the drag and drop
// operation. See gdk_drag_context_manage_dnd() for more information.
func (context *DragContext) ConnectActionChanged(f func(action DragAction)) externglib.SignalHandle {
	return context.Connect("action-changed", f)
}

// ConnectCancel: drag and drop operation was cancelled.
//
// This signal will only be emitted if the DragContext manages the drag and drop
// operation. See gdk_drag_context_manage_dnd() for more information.
func (context *DragContext) ConnectCancel(f func(reason DragCancelReason)) externglib.SignalHandle {
	return context.Connect("cancel", f)
}

// ConnectDNDFinished: drag and drop operation was finished, the drag
// destination finished reading all data. The drag source can now free all
// miscellaneous data.
//
// This signal will only be emitted if the DragContext manages the drag and drop
// operation. See gdk_drag_context_manage_dnd() for more information.
func (context *DragContext) ConnectDNDFinished(f func()) externglib.SignalHandle {
	return context.Connect("dnd-finished", f)
}

// ConnectDropPerformed: drag and drop operation was performed on an accepting
// client.
//
// This signal will only be emitted if the DragContext manages the drag and drop
// operation. See gdk_drag_context_manage_dnd() for more information.
func (context *DragContext) ConnectDropPerformed(f func(time int)) externglib.SignalHandle {
	return context.Connect("drop-performed", f)
}
