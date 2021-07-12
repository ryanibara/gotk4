// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_accel_map_get_type()), F: marshalAccelMaper},
	})
}

// AccelMaper describes AccelMap's methods.
type AccelMaper interface {
	privateAccelMap()
}

// AccelMap: accelerator maps are used to define runtime configurable
// accelerators. Functions for manipulating them are are usually used by higher
// level convenience mechanisms like UIManager and are thus considered
// “low-level”. You’ll want to use them if you’re manually creating menus that
// should have user-configurable accelerators.
//
// An accelerator is uniquely defined by: - accelerator path - accelerator key -
// accelerator modifiers
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
	*externglib.Object
}

var (
	_ AccelMaper      = (*AccelMap)(nil)
	_ gextras.Nativer = (*AccelMap)(nil)
)

func wrapAccelMap(obj *externglib.Object) *AccelMap {
	return &AccelMap{
		Object: obj,
	}
}

func marshalAccelMaper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAccelMap(obj), nil
}

func (*AccelMap) privateAccelMap() {}

// AccelMapAddEntry registers a new accelerator with the global accelerator map.
// This function should only be called once per @accel_path with the canonical
// @accel_key and @accel_mods for this path. To change the accelerator during
// runtime programatically, use gtk_accel_map_change_entry().
//
// Set @accel_key and @accel_mods to 0 to request a removal of the accelerator.
//
// Note that @accel_path string will be stored in a #GQuark. Therefore, if you
// pass a static string, you can save some memory by interning it first with
// g_intern_static_string().
func AccelMapAddEntry(accelPath string, accelKey uint, accelMods gdk.ModifierType) {
	var _arg1 *C.gchar          // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelPath)))
	_arg2 = C.guint(accelKey)
	_arg3 = C.GdkModifierType(accelMods)

	C.gtk_accel_map_add_entry(_arg1, _arg2, _arg3)
}

// AccelMapAddFilter adds a filter to the global list of accel path filters.
//
// Accel map entries whose accel path matches one of the filters are skipped by
// gtk_accel_map_foreach().
//
// This function is intended for GTK+ modules that create their own menus, but
// don’t want them to be saved into the applications accelerator map dump.
func AccelMapAddFilter(filterPattern string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filterPattern)))

	C.gtk_accel_map_add_filter(_arg1)
}

// AccelMapChangeEntry changes the @accel_key and @accel_mods currently
// associated with @accel_path. Due to conflicts with other accelerators, a
// change may not always be possible, @replace indicates whether other
// accelerators may be deleted to resolve such conflicts. A change will only
// occur if all conflicts could be resolved (which might not be the case if
// conflicting accelerators are locked). Successful changes are indicated by a
// true return value.
//
// Note that @accel_path string will be stored in a #GQuark. Therefore, if you
// pass a static string, you can save some memory by interning it first with
// g_intern_static_string().
func AccelMapChangeEntry(accelPath string, accelKey uint, accelMods gdk.ModifierType, replace bool) bool {
	var _arg1 *C.gchar          // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out
	var _arg4 C.gboolean        // out
	var _cret C.gboolean        // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelPath)))
	_arg2 = C.guint(accelKey)
	_arg3 = C.GdkModifierType(accelMods)
	if replace {
		_arg4 = C.TRUE
	}

	_cret = C.gtk_accel_map_change_entry(_arg1, _arg2, _arg3, _arg4)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AccelMapGet gets the singleton global AccelMap object. This object is useful
// only for notification of changes to the accelerator map via the ::changed
// signal; it isn’t a parameter to the other accelerator map functions.
func AccelMapGet() *AccelMap {
	var _cret *C.GtkAccelMap // in

	_cret = C.gtk_accel_map_get()

	var _accelMap *AccelMap // out

	_accelMap = wrapAccelMap(externglib.Take(unsafe.Pointer(_cret)))

	return _accelMap
}

// AccelMapLoad parses a file previously saved with gtk_accel_map_save() for
// accelerator specifications, and propagates them accordingly.
func AccelMapLoad(fileName string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fileName)))

	C.gtk_accel_map_load(_arg1)
}

// AccelMapLoadFd: filedescriptor variant of gtk_accel_map_load().
//
// Note that the file descriptor will not be closed by this function.
func AccelMapLoadFd(fd int) {
	var _arg1 C.gint // out

	_arg1 = C.gint(fd)

	C.gtk_accel_map_load_fd(_arg1)
}

// AccelMapLoadScanner variant of gtk_accel_map_load().
func AccelMapLoadScanner(scanner *glib.Scanner) {
	var _arg1 *C.GScanner // out

	_arg1 = (*C.GScanner)(unsafe.Pointer(scanner))

	C.gtk_accel_map_load_scanner(_arg1)
}

// AccelMapLockPath locks the given accelerator path. If the accelerator map
// doesn’t yet contain an entry for @accel_path, a new one is created.
//
// Locking an accelerator path prevents its accelerator from being changed
// during runtime. A locked accelerator path can be unlocked by
// gtk_accel_map_unlock_path(). Refer to gtk_accel_map_change_entry() for
// information about runtime accelerator changes.
//
// If called more than once, @accel_path remains locked until
// gtk_accel_map_unlock_path() has been called an equivalent number of times.
//
// Note that locking of individual accelerator paths is independent from locking
// the AccelGroup containing them. For runtime accelerator changes to be
// possible, both the accelerator path and its AccelGroup have to be unlocked.
func AccelMapLockPath(accelPath string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelPath)))

	C.gtk_accel_map_lock_path(_arg1)
}

// AccelMapLookupEntry looks up the accelerator entry for @accel_path and fills
// in @key.
func AccelMapLookupEntry(accelPath string) (AccelKey, bool) {
	var _arg1 *C.gchar // out
	var _key AccelKey
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelPath)))

	_cret = C.gtk_accel_map_lookup_entry(_arg1, (*C.GtkAccelKey)(unsafe.Pointer(&_key)))

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _key, _ok
}

// AccelMapSave saves current accelerator specifications (accelerator path, key
// and modifiers) to @file_name. The file is written in a format suitable to be
// read back in by gtk_accel_map_load().
func AccelMapSave(fileName string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fileName)))

	C.gtk_accel_map_save(_arg1)
}

// AccelMapSaveFd: filedescriptor variant of gtk_accel_map_save().
//
// Note that the file descriptor will not be closed by this function.
func AccelMapSaveFd(fd int) {
	var _arg1 C.gint // out

	_arg1 = C.gint(fd)

	C.gtk_accel_map_save_fd(_arg1)
}

// AccelMapUnlockPath undoes the last call to gtk_accel_map_lock_path() on this
// @accel_path. Refer to gtk_accel_map_lock_path() for information about
// accelerator path locking.
func AccelMapUnlockPath(accelPath string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelPath)))

	C.gtk_accel_map_unlock_path(_arg1)
}
