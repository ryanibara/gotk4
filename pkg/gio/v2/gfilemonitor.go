// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_FileMonitor_ConnectChanged(gpointer, GFile*, GFile*, GFileMonitorEvent, guintptr);
// extern void _gotk4_gio2_FileMonitorClass_changed(GFileMonitor*, GFile*, GFile*, GFileMonitorEvent);
// extern gboolean _gotk4_gio2_FileMonitorClass_cancel(GFileMonitor*);
// gboolean _gotk4_gio2_FileMonitor_virtual_cancel(void* fnptr, GFileMonitor* arg0) {
//   return ((gboolean (*)(GFileMonitor*))(fnptr))(arg0);
// };
// void _gotk4_gio2_FileMonitor_virtual_changed(void* fnptr, GFileMonitor* arg0, GFile* arg1, GFile* arg2, GFileMonitorEvent arg3) {
//   ((void (*)(GFileMonitor*, GFile*, GFile*, GFileMonitorEvent))(fnptr))(arg0, arg1, arg2, arg3);
// };
import "C"

// GType values.
var (
	GTypeFileMonitor = coreglib.Type(C.g_file_monitor_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFileMonitor, F: marshalFileMonitor},
	})
}

// FileMonitorOverrides contains methods that are overridable.
type FileMonitorOverrides struct {
	// Cancel cancels a file monitor.
	//
	// The function returns the following values:
	//
	//    - ok always TRUE.
	//
	Cancel func() bool
	// The function takes the following parameters:
	//
	//    - file
	//    - otherFile
	//    - eventType
	//
	Changed func(file, otherFile Filer, eventType FileMonitorEvent)
}

func defaultFileMonitorOverrides(v *FileMonitor) FileMonitorOverrides {
	return FileMonitorOverrides{
		Cancel:  v.cancel,
		Changed: v.changed,
	}
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
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*FileMonitor)(nil)
)

// FileMonitorrer describes types inherited from class FileMonitor.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type FileMonitorrer interface {
	coreglib.Objector
	baseFileMonitor() *FileMonitor
}

var _ FileMonitorrer = (*FileMonitor)(nil)

func init() {
	coreglib.RegisterClassInfo[*FileMonitor, *FileMonitorClass, FileMonitorOverrides](
		GTypeFileMonitor,
		initFileMonitorClass,
		wrapFileMonitor,
		defaultFileMonitorOverrides,
	)
}

func initFileMonitorClass(gclass unsafe.Pointer, overrides FileMonitorOverrides, classInitFunc func(*FileMonitorClass)) {
	pclass := (*C.GFileMonitorClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeFileMonitor))))

	if overrides.Cancel != nil {
		pclass.cancel = (*[0]byte)(C._gotk4_gio2_FileMonitorClass_cancel)
	}

	if overrides.Changed != nil {
		pclass.changed = (*[0]byte)(C._gotk4_gio2_FileMonitorClass_changed)
	}

	if classInitFunc != nil {
		class := (*FileMonitorClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapFileMonitor(obj *coreglib.Object) *FileMonitor {
	return &FileMonitor{
		Object: obj,
	}
}

func marshalFileMonitor(p uintptr) (interface{}, error) {
	return wrapFileMonitor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (monitor *FileMonitor) baseFileMonitor() *FileMonitor {
	return monitor
}

// BaseFileMonitor returns the underlying base object.
func BaseFileMonitor(obj FileMonitorrer) *FileMonitor {
	return obj.baseFileMonitor()
}

// ConnectChanged is emitted when file has been changed.
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
func (monitor *FileMonitor) ConnectChanged(f func(file, otherFile Filer, eventType FileMonitorEvent)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(monitor, "changed", false, unsafe.Pointer(C._gotk4_gio2_FileMonitor_ConnectChanged), f)
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

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

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

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	_arg2 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(otherFile).Native()))
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

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

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

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))
	_arg1 = C.gint(limitMsecs)

	C.g_file_monitor_set_rate_limit(_arg0, _arg1)
	runtime.KeepAlive(monitor)
	runtime.KeepAlive(limitMsecs)
}

// Cancel cancels a file monitor.
//
// The function returns the following values:
//
//    - ok always TRUE.
//
func (monitor *FileMonitor) cancel() bool {
	gclass := (*C.GFileMonitorClass)(coreglib.PeekParentClass(monitor))
	fnarg := gclass.cancel

	var _arg0 *C.GFileMonitor // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_cret = C._gotk4_gio2_FileMonitor_virtual_cancel(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(monitor)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
//    - file
//    - otherFile
//    - eventType
//
func (monitor *FileMonitor) changed(file, otherFile Filer, eventType FileMonitorEvent) {
	gclass := (*C.GFileMonitorClass)(coreglib.PeekParentClass(monitor))
	fnarg := gclass.changed

	var _arg0 *C.GFileMonitor     // out
	var _arg1 *C.GFile            // out
	var _arg2 *C.GFile            // out
	var _arg3 C.GFileMonitorEvent // out

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))
	_arg2 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(otherFile).Native()))
	_arg3 = C.GFileMonitorEvent(eventType)

	C._gotk4_gio2_FileMonitor_virtual_changed(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(monitor)
	runtime.KeepAlive(file)
	runtime.KeepAlive(otherFile)
	runtime.KeepAlive(eventType)
}

// FileMonitorClass: instance of this type is always passed by reference.
type FileMonitorClass struct {
	*fileMonitorClass
}

// fileMonitorClass is the struct that's finalized.
type fileMonitorClass struct {
	native *C.GFileMonitorClass
}
