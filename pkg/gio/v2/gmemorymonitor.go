// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_MemoryMonitorInterface_low_memory_warning(GMemoryMonitor*, GMemoryMonitorWarningLevel);
// extern void _gotk4_gio2_MemoryMonitor_ConnectLowMemoryWarning(gpointer, GMemoryMonitorWarningLevel, guintptr);
import "C"

// GType values.
var (
	GTypeMemoryMonitor = coreglib.Type(C.g_memory_monitor_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMemoryMonitor, F: marshalMemoryMonitor},
	})
}

// MEMORY_MONITOR_EXTENSION_POINT_NAME: extension point for memory usage
// monitoring functionality. See [Extending GIO][extending-gio].
const MEMORY_MONITOR_EXTENSION_POINT_NAME = "gio-memory-monitor"

// MemoryMonitorOverrider contains methods that are overridable.
type MemoryMonitorOverrider interface {
	// The function takes the following parameters:
	//
	LowMemoryWarning(level MemoryMonitorWarningLevel)
}

// MemoryMonitor will monitor system memory and suggest to the application when
// to free memory so as to leave more room for other applications. It is
// implemented on Linux using the Low Memory Monitor
// (https://gitlab.freedesktop.org/hadess/low-memory-monitor/) (API
// documentation (https://hadess.pages.freedesktop.org/low-memory-monitor/)).
//
// There is also an implementation for use inside Flatpak sandboxes.
//
// Possible actions to take when the signal is received are:
//
// - Free caches
//
// - Save files that haven't been looked at in a while to disk, ready to be
// reopened when needed
//
// - Run a garbage collection cycle
//
// - Try and compress fragmented allocations
//
// - Exit on idle if the process has no reason to stay around
//
// - Call malloc_trim(3) (man:malloc_trim) to return cached heap pages to the
// kernel (if supported by your libc)
//
// Note that some actions may not always improve system performance, and so
// should be profiled for your application. malloc_trim(), for example, may make
// future heap allocations slower (due to releasing cached heap pages back to
// the kernel).
//
// See MonitorWarningLevel for details on the various warning levels.
//
//    static void
//    warning_cb (GMemoryMonitor *m, GMemoryMonitorWarningLevel level)
//    {
//      g_debug ("Warning level: d", level);
//      if (warning_level > G_MEMORY_MONITOR_WARNING_LEVEL_LOW)
//        drop_caches ();
//    }
//
//    static GMemoryMonitor *
//    monitor_low_memory (void)
//    {
//      GMemoryMonitor *m;
//      m = g_memory_monitor_dup_default ();
//      g_signal_connect (G_OBJECT (m), "low-memory-warning",
//                        G_CALLBACK (warning_cb), NULL);
//      return m;
//    }
//
// Don't forget to disconnect the Monitor::low-memory-warning signal, and unref
// the Monitor itself when exiting.
//
// MemoryMonitor wraps an interface. This means the user can get the
// underlying type by calling Cast().
type MemoryMonitor struct {
	_ [0]func() // equal guard
	Initable
}

var ()

// MemoryMonitorrer describes MemoryMonitor's interface methods.
type MemoryMonitorrer interface {
	coreglib.Objector

	baseMemoryMonitor() *MemoryMonitor
}

var _ MemoryMonitorrer = (*MemoryMonitor)(nil)

func ifaceInitMemoryMonitorrer(gifacePtr, data C.gpointer) {
	iface := (*C.GMemoryMonitorInterface)(unsafe.Pointer(gifacePtr))
	iface.low_memory_warning = (*[0]byte)(C._gotk4_gio2_MemoryMonitorInterface_low_memory_warning)
}

//export _gotk4_gio2_MemoryMonitorInterface_low_memory_warning
func _gotk4_gio2_MemoryMonitorInterface_low_memory_warning(arg0 *C.GMemoryMonitor, arg1 C.GMemoryMonitorWarningLevel) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(MemoryMonitorOverrider)

	var _level MemoryMonitorWarningLevel // out

	_level = MemoryMonitorWarningLevel(arg1)

	iface.LowMemoryWarning(_level)
}

func wrapMemoryMonitor(obj *coreglib.Object) *MemoryMonitor {
	return &MemoryMonitor{
		Initable: Initable{
			Object: obj,
		},
	}
}

func marshalMemoryMonitor(p uintptr) (interface{}, error) {
	return wrapMemoryMonitor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (monitor *MemoryMonitor) baseMemoryMonitor() *MemoryMonitor {
	return monitor
}

// BaseMemoryMonitor returns the underlying base object.
func BaseMemoryMonitor(obj MemoryMonitorrer) *MemoryMonitor {
	return obj.baseMemoryMonitor()
}

//export _gotk4_gio2_MemoryMonitor_ConnectLowMemoryWarning
func _gotk4_gio2_MemoryMonitor_ConnectLowMemoryWarning(arg0 C.gpointer, arg1 C.GMemoryMonitorWarningLevel, arg2 C.guintptr) {
	var f func(level MemoryMonitorWarningLevel)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(level MemoryMonitorWarningLevel))
	}

	var _level MemoryMonitorWarningLevel // out

	_level = MemoryMonitorWarningLevel(arg1)

	f(_level)
}

// ConnectLowMemoryWarning is emitted when the system is running low on free
// memory. The signal handler should then take the appropriate action depending
// on the warning level. See the MonitorWarningLevel documentation for details.
func (monitor *MemoryMonitor) ConnectLowMemoryWarning(f func(level MemoryMonitorWarningLevel)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(monitor, "low-memory-warning", false, unsafe.Pointer(C._gotk4_gio2_MemoryMonitor_ConnectLowMemoryWarning), f)
}

// MemoryMonitorDupDefault gets a reference to the default Monitor for the
// system.
//
// The function returns the following values:
//
//    - memoryMonitor: new reference to the default Monitor.
//
func MemoryMonitorDupDefault() *MemoryMonitor {
	var _cret *C.GMemoryMonitor // in

	_cret = C.g_memory_monitor_dup_default()

	var _memoryMonitor *MemoryMonitor // out

	_memoryMonitor = wrapMemoryMonitor(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _memoryMonitor
}
