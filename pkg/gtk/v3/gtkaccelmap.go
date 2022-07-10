// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeAccelMap returns the GType for the type AccelMap.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeAccelMap() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "AccelMap").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalAccelMap)
	return gtype
}

// AccelMap: accelerator maps are used to define runtime configurable
// accelerators. Functions for manipulating them are are usually used by higher
// level convenience mechanisms like UIManager and are thus considered
// “low-level”. You’ll want to use them if you’re manually creating menus that
// should have user-configurable accelerators.
//
// An accelerator is uniquely defined by:
//
// - accelerator path
//
// - accelerator key
//
// - accelerator modifiers
//
// The accelerator path must consist of
// “<WINDOWTYPE>/Category1/Category2/.../Action”, where WINDOWTYPE should be a
// unique application-specific identifier that corresponds to the kind of window
// the accelerator is being used in, e.g. “Gimp-Image”, “Abiword-Document” or
// “Gnumeric-Settings”. The “Category1/.../Action” portion is most appropriately
// chosen by the action the accelerator triggers, i.e. for accelerators on menu
// items, choose the item’s menu path, e.g. “File/Save As”, “Image/View/Zoom” or
// “Edit/Select All”. So a full valid accelerator path may look like:
// “<Gimp-Toolbox>/File/Dialogs/Tool Options...”.
//
// All accelerators are stored inside one global AccelMap that can be obtained
// using gtk_accel_map_get(). See [Monitoring changes][monitoring-changes] for
// additional details.
//
//
// Manipulating accelerators
//
// New accelerators can be added using gtk_accel_map_add_entry(). To search for
// specific accelerator, use gtk_accel_map_lookup_entry(). Modifications of
// existing accelerators should be done using gtk_accel_map_change_entry().
//
// In order to avoid having some accelerators changed, they can be locked using
// gtk_accel_map_lock_path(). Unlocking is done using
// gtk_accel_map_unlock_path().
//
//
// Saving and loading accelerator maps
//
// Accelerator maps can be saved to and loaded from some external resource. For
// simple saving and loading from file, gtk_accel_map_save() and
// gtk_accel_map_load() are provided. Saving and loading can also be done by
// providing file descriptor to gtk_accel_map_save_fd() and
// gtk_accel_map_load_fd().
//
//
// Monitoring changes
//
// AccelMap object is only useful for monitoring changes of accelerators. By
// connecting to AccelMap::changed signal, one can monitor changes of all
// accelerators. It is also possible to monitor only single accelerator path by
// using it as a detail of the AccelMap::changed signal.
type AccelMap struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*AccelMap)(nil)
)

func wrapAccelMap(obj *coreglib.Object) *AccelMap {
	return &AccelMap{
		Object: obj,
	}
}

func marshalAccelMap(p uintptr) (interface{}, error) {
	return wrapAccelMap(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AccelMapAddFilter adds a filter to the global list of accel path filters.
//
// Accel map entries whose accel path matches one of the filters are skipped by
// gtk_accel_map_foreach().
//
// This function is intended for GTK+ modules that create their own menus, but
// don’t want them to be saved into the applications accelerator map dump.
//
// The function takes the following parameters:
//
//    - filterPattern: pattern (see Spec).
//
func AccelMapAddFilter(filterPattern string) {
	var _args [1]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(filterPattern)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "add_filter")
	_info.InvokeFunction(_args[:], nil)

	runtime.KeepAlive(filterPattern)
}

// AccelMapGet gets the singleton global AccelMap object. This object is useful
// only for notification of changes to the accelerator map via the ::changed
// signal; it isn’t a parameter to the other accelerator map functions.
//
// The function returns the following values:
//
//    - accelMap: global AccelMap object.
//
func AccelMapGet() *AccelMap {
	_info := girepository.MustFind("Gtk", "get")
	_gret := _info.InvokeFunction(nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _accelMap *AccelMap // out

	_accelMap = wrapAccelMap(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _accelMap
}

// AccelMapLoad parses a file previously saved with gtk_accel_map_save() for
// accelerator specifications, and propagates them accordingly.
//
// The function takes the following parameters:
//
//    - fileName: file containing accelerator specifications, in the GLib file
//      name encoding.
//
func AccelMapLoad(fileName string) {
	var _args [1]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "load")
	_info.InvokeFunction(_args[:], nil)

	runtime.KeepAlive(fileName)
}

// AccelMapLoadFd: filedescriptor variant of gtk_accel_map_load().
//
// Note that the file descriptor will not be closed by this function.
//
// The function takes the following parameters:
//
//    - fd: valid readable file descriptor.
//
func AccelMapLoadFd(fd int32) {
	var _args [1]girepository.Argument

	*(*C.gint)(unsafe.Pointer(&_args[0])) = C.gint(fd)

	_info := girepository.MustFind("Gtk", "load_fd")
	_info.InvokeFunction(_args[:], nil)

	runtime.KeepAlive(fd)
}

// AccelMapLoadScanner variant of gtk_accel_map_load().
//
// The function takes the following parameters:
//
//    - scanner which has already been provided with an input file.
//
func AccelMapLoadScanner(scanner *glib.Scanner) {
	var _args [1]girepository.Argument

	*(**C.GScanner)(unsafe.Pointer(&_args[0])) = (*C.GScanner)(gextras.StructNative(unsafe.Pointer(scanner)))

	_info := girepository.MustFind("Gtk", "load_scanner")
	_info.InvokeFunction(_args[:], nil)

	runtime.KeepAlive(scanner)
}

// AccelMapLockPath locks the given accelerator path. If the accelerator map
// doesn’t yet contain an entry for accel_path, a new one is created.
//
// Locking an accelerator path prevents its accelerator from being changed
// during runtime. A locked accelerator path can be unlocked by
// gtk_accel_map_unlock_path(). Refer to gtk_accel_map_change_entry() for
// information about runtime accelerator changes.
//
// If called more than once, accel_path remains locked until
// gtk_accel_map_unlock_path() has been called an equivalent number of times.
//
// Note that locking of individual accelerator paths is independent from locking
// the AccelGroup containing them. For runtime accelerator changes to be
// possible, both the accelerator path and its AccelGroup have to be unlocked.
//
// The function takes the following parameters:
//
//    - accelPath: valid accelerator path.
//
func AccelMapLockPath(accelPath string) {
	var _args [1]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(accelPath)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "lock_path")
	_info.InvokeFunction(_args[:], nil)

	runtime.KeepAlive(accelPath)
}

// AccelMapLookupEntry looks up the accelerator entry for accel_path and fills
// in key.
//
// The function takes the following parameters:
//
//    - accelPath: valid accelerator path.
//
// The function returns the following values:
//
//    - key (optional): accelerator key to be filled in (optional).
//    - ok: TRUE if accel_path is known, FALSE otherwise.
//
func AccelMapLookupEntry(accelPath string) (*AccelKey, bool) {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(accelPath)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "lookup_entry")
	_gret := _info.InvokeFunction(_args[:], _outs[:])
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(accelPath)

	var _key *AccelKey // out
	var _ok bool       // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_key = (*AccelKey)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[0])))))
	}
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _key, _ok
}

// AccelMapSave saves current accelerator specifications (accelerator path, key
// and modifiers) to file_name. The file is written in a format suitable to be
// read back in by gtk_accel_map_load().
//
// The function takes the following parameters:
//
//    - fileName: name of the file to contain accelerator specifications, in the
//      GLib file name encoding.
//
func AccelMapSave(fileName string) {
	var _args [1]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "save")
	_info.InvokeFunction(_args[:], nil)

	runtime.KeepAlive(fileName)
}

// AccelMapSaveFd: filedescriptor variant of gtk_accel_map_save().
//
// Note that the file descriptor will not be closed by this function.
//
// The function takes the following parameters:
//
//    - fd: valid writable file descriptor.
//
func AccelMapSaveFd(fd int32) {
	var _args [1]girepository.Argument

	*(*C.gint)(unsafe.Pointer(&_args[0])) = C.gint(fd)

	_info := girepository.MustFind("Gtk", "save_fd")
	_info.InvokeFunction(_args[:], nil)

	runtime.KeepAlive(fd)
}

// AccelMapUnlockPath undoes the last call to gtk_accel_map_lock_path() on this
// accel_path. Refer to gtk_accel_map_lock_path() for information about
// accelerator path locking.
//
// The function takes the following parameters:
//
//    - accelPath: valid accelerator path.
//
func AccelMapUnlockPath(accelPath string) {
	var _args [1]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(accelPath)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "unlock_path")
	_info.InvokeFunction(_args[:], nil)

	runtime.KeepAlive(accelPath)
}
