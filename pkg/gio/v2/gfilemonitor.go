// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gio2_FileMonitorClass_cancel(GFileMonitor*);
import "C"

// glib.Type values for gfilemonitor.go.
var GTypeFileMonitor = coreglib.Type(C.g_file_monitor_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
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
}

//export _gotk4_gio2_FileMonitorClass_cancel
func _gotk4_gio2_FileMonitorClass_cancel(arg0 *C.GFileMonitor) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Cancel() bool })

	ok := iface.Cancel()

	if ok {
		cret = C.TRUE
	}

	return cret
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

// Cancel cancels a file monitor.
//
// The function returns the following values:
//
//    - ok always TRUE.
//
func (monitor *FileMonitor) Cancel() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))
	*(**FileMonitor)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "FileMonitor").InvokeMethod("cancel", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsCancelled returns whether the monitor is canceled.
//
// The function returns the following values:
//
//    - ok: TRUE if monitor is canceled. FALSE otherwise.
//
func (monitor *FileMonitor) IsCancelled() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))
	*(**FileMonitor)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "FileMonitor").InvokeMethod("is_cancelled", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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
func (monitor *FileMonitor) SetRateLimit(limitMsecs int32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))
	_arg1 = C.gint(limitMsecs)
	*(**FileMonitor)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gio", "FileMonitor").InvokeMethod("set_rate_limit", args[:], nil)

	runtime.KeepAlive(monitor)
	runtime.KeepAlive(limitMsecs)
}
