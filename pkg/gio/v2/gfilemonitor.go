// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_file_monitor_get_type()), F: marshalFileMonitorrer},
	})
}

// FileMonitorOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FileMonitorOverrider interface {
	// Cancel cancels a file monitor.
	//
	// The function returns the following values:
	//
	//    - ok always TRUE.
	//
	Cancel() bool
	// The function takes the following parameters:
	//
	//    - file
	//    - otherFile
	//    - eventType
	//
	Changed(file, otherFile Filer, eventType FileMonitorEvent)
}

// FileMonitor monitors a file or directory for changes.
//
// To obtain a Monitor for a file or directory, use g_file_monitor(),
// g_file_monitor_file(), or g_file_monitor_directory().
//
// To get informed about changes to the file or directory you are monitoring,
// connect to the Monitor::changed signal. The signal will be emitted in the
// [thread-default main context][g-main-context-push-thread-default] of the
// thread that the monitor was created in (though if the global default main
// context is blocked, this may cause notifications to be blocked even if the
// thread-default context is still running).
type FileMonitor struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*FileMonitor)(nil)
)

// FileMonitorrer describes types inherited from class FileMonitor.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type FileMonitorrer interface {
	externglib.Objector
	baseFileMonitor() *FileMonitor
}

var _ FileMonitorrer = (*FileMonitor)(nil)

func wrapFileMonitor(obj *externglib.Object) *FileMonitor {
	return &FileMonitor{
		Object: obj,
	}
}

func marshalFileMonitorrer(p uintptr) (interface{}, error) {
	return wrapFileMonitor(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (monitor *FileMonitor) baseFileMonitor() *FileMonitor {
	return monitor
}

// BaseFileMonitor returns the underlying base object.
func BaseFileMonitor(obj FileMonitorrer) *FileMonitor {
	return obj.baseFileMonitor()
}

// ConnectChanged: emitted when file has been changed.
//
// If using G_FILE_MONITOR_WATCH_MOVES on a directory monitor, and the
// information is available (and if supported by the backend), event_type may be
// G_FILE_MONITOR_EVENT_RENAMED, G_FILE_MONITOR_EVENT_MOVED_IN or
// G_FILE_MONITOR_EVENT_MOVED_OUT.
//
// In all cases file will be a child of the monitored directory. For renames,
// file will be the old name and other_file is the new name. For "moved in"
// events, file is the name of the file that appeared and other_file is the old
// name that it was moved from (in another directory). For "moved out" events,
// file is the name of the file that used to be in this directory and other_file
// is the name of the file at its new location.
//
// It makes sense to treat G_FILE_MONITOR_EVENT_MOVED_IN as equivalent to
// G_FILE_MONITOR_EVENT_CREATED and G_FILE_MONITOR_EVENT_MOVED_OUT as equivalent
// to G_FILE_MONITOR_EVENT_DELETED, with extra information.
// G_FILE_MONITOR_EVENT_RENAMED is equivalent to a delete/create pair. This is
// exactly how the events will be reported in the case that the
// G_FILE_MONITOR_WATCH_MOVES flag is not in use.
//
// If using the deprecated flag G_FILE_MONITOR_SEND_MOVED flag and event_type is
// FILE_MONITOR_EVENT_MOVED, file will be set to a #GFile containing the old
// path, and other_file will be set to a #GFile containing the new path.
//
// In all the other cases, other_file will be set to LL.
func (monitor *FileMonitor) ConnectChanged(f func(file, otherFile Filer, eventType FileMonitorEvent)) externglib.SignalHandle {
	return monitor.Connect("changed", f)
}

// Cancel cancels a file monitor.
//
// The function returns the following values:
//
//    - ok always TRUE.
//
func (monitor *FileMonitor) Cancel() bool {
	var _arg0 *C.GFileMonitor // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.g_file_monitor_cancel(_arg0)
	runtime.KeepAlive(monitor)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// EmitEvent emits the Monitor::changed signal if a change has taken place.
// Should be called from file monitor implementations only.
//
// Implementations are responsible to call this method from the [thread-default
// main context][g-main-context-push-thread-default] of the thread that the
// monitor was created in.
//
// The function takes the following parameters:
//
//    - child: #GFile.
//    - otherFile: #GFile.
//    - eventType: set of MonitorEvent flags.
//
func (monitor *FileMonitor) EmitEvent(child, otherFile Filer, eventType FileMonitorEvent) {
	var _arg0 *C.GFileMonitor     // out
	var _arg1 *C.GFile            // out
	var _arg2 *C.GFile            // out
	var _arg3 C.GFileMonitorEvent // out

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(monitor.Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GFile)(unsafe.Pointer(otherFile.Native()))
	_arg3 = C.GFileMonitorEvent(eventType)

	C.g_file_monitor_emit_event(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(monitor)
	runtime.KeepAlive(child)
	runtime.KeepAlive(otherFile)
	runtime.KeepAlive(eventType)
}

// IsCancelled returns whether the monitor is canceled.
//
// The function returns the following values:
//
//    - ok: TRUE if monitor is canceled. FALSE otherwise.
//
func (monitor *FileMonitor) IsCancelled() bool {
	var _arg0 *C.GFileMonitor // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.g_file_monitor_is_cancelled(_arg0)
	runtime.KeepAlive(monitor)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetRateLimit sets the rate limit to which the monitor will report consecutive
// change events to the same file.
//
// The function takes the following parameters:
//
//    - limitMsecs: non-negative integer with the limit in milliseconds to poll
//      for changes.
//
func (monitor *FileMonitor) SetRateLimit(limitMsecs int) {
	var _arg0 *C.GFileMonitor // out
	var _arg1 C.gint          // out

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(monitor.Native()))
	_arg1 = C.gint(limitMsecs)

	C.g_file_monitor_set_rate_limit(_arg0, _arg1)
	runtime.KeepAlive(monitor)
	runtime.KeepAlive(limitMsecs)
}
