// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gio2_FileMonitorClass_cancel(GFileMonitor*);
// extern void _gotk4_gio2_FileMonitorClass_changed(GFileMonitor*, GFile*, GFile*, GFileMonitorEvent);
// extern void _gotk4_gio2_FileMonitor_ConnectChanged(gpointer, GFile*, GFile*, GFileMonitorEvent, guintptr);
import "C"

// glib.Type values for gfilemonitor.go.
var GTypeFileMonitor = externglib.Type(C.g_file_monitor_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeFileMonitor, F: marshalFileMonitor},
	})
}

// FileMonitorOverrider contains methods that are overridable.
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

func classInitFileMonitorrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GFileMonitorClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GFileMonitorClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Cancel() bool }); ok {
		pclass.cancel = (*[0]byte)(C._gotk4_gio2_FileMonitorClass_cancel)
	}

	if _, ok := goval.(interface {
		Changed(file, otherFile Filer, eventType FileMonitorEvent)
	}); ok {
		pclass.changed = (*[0]byte)(C._gotk4_gio2_FileMonitorClass_changed)
	}
}

//export _gotk4_gio2_FileMonitorClass_cancel
func _gotk4_gio2_FileMonitorClass_cancel(arg0 *C.GFileMonitor) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Cancel() bool })

	ok := iface.Cancel()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_FileMonitorClass_changed
func _gotk4_gio2_FileMonitorClass_changed(arg0 *C.GFileMonitor, arg1 *C.GFile, arg2 *C.GFile, arg3 C.GFileMonitorEvent) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Changed(file, otherFile Filer, eventType FileMonitorEvent)
	})

	var _file Filer                 // out
	var _otherFile Filer            // out
	var _eventType FileMonitorEvent // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.Filer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Filer)
			return ok
		})
		rv, ok := casted.(Filer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Filer")
		}
		_file = rv
	}
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gio.Filer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Filer)
			return ok
		})
		rv, ok := casted.(Filer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Filer")
		}
		_otherFile = rv
	}
	_eventType = FileMonitorEvent(arg3)

	iface.Changed(_file, _otherFile, _eventType)
}

func wrapFileMonitor(obj *externglib.Object) *FileMonitor {
	return &FileMonitor{
		Object: obj,
	}
}

func marshalFileMonitor(p uintptr) (interface{}, error) {
	return wrapFileMonitor(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (monitor *FileMonitor) baseFileMonitor() *FileMonitor {
	return monitor
}

// BaseFileMonitor returns the underlying base object.
func BaseFileMonitor(obj FileMonitorrer) *FileMonitor {
	return obj.baseFileMonitor()
}

//export _gotk4_gio2_FileMonitor_ConnectChanged
func _gotk4_gio2_FileMonitor_ConnectChanged(arg0 C.gpointer, arg1 *C.GFile, arg2 *C.GFile, arg3 C.GFileMonitorEvent, arg4 C.guintptr) {
	var f func(file, otherFile Filer, eventType FileMonitorEvent)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(file, otherFile Filer, eventType FileMonitorEvent))
	}

	var _file Filer                 // out
	var _otherFile Filer            // out
	var _eventType FileMonitorEvent // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.Filer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Filer)
			return ok
		})
		rv, ok := casted.(Filer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Filer")
		}
		_file = rv
	}
	if arg2 != nil {
		{
			objptr := unsafe.Pointer(arg2)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Filer)
				return ok
			})
			rv, ok := casted.(Filer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Filer")
			}
			_otherFile = rv
		}
	}
	_eventType = FileMonitorEvent(arg3)

	f(_file, _otherFile, _eventType)
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
	return externglib.ConnectGeneratedClosure(monitor, "changed", false, unsafe.Pointer(C._gotk4_gio2_FileMonitor_ConnectChanged), f)
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
