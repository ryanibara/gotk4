// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_unix_mount_monitor_get_type()), F: marshalUnixMountMonitor},
		{T: externglib.Type(C.g_unix_mount_entry_get_type()), F: marshalUnixMountEntry},
		{T: externglib.Type(C.g_unix_mount_point_get_type()), F: marshalUnixMountPoint},
	})
}

// UnixIsMountPathSystemInternal determines if @mount_path is considered an
// implementation of the OS. This is primarily used for hiding mountable and
// mounted volumes that only are used in the OS and has little to no relevance
// to the casual user.
func UnixIsMountPathSystemInternal(mountPath string) bool {
	var _arg1 *C.char    // out
	var _cret C.gboolean // in

	_arg1 = (*C.char)(C.CString(mountPath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_unix_is_mount_path_system_internal(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixIsSystemDevicePath determines if @device_path is considered a block
// device path which is only used in implementation of the OS. This is primarily
// used for hiding mounted volumes that are intended as APIs for programs to
// read, and system administrators at a shell; rather than something that
// should, for example, appear in a GUI. For example, the Linux `/proc`
// filesystem.
//
// The list of device paths considered ‘system’ ones may change over time.
func UnixIsSystemDevicePath(devicePath string) bool {
	var _arg1 *C.char    // out
	var _cret C.gboolean // in

	_arg1 = (*C.char)(C.CString(devicePath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_unix_is_system_device_path(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixIsSystemFSType determines if @fs_type is considered a type of file system
// which is only used in implementation of the OS. This is primarily used for
// hiding mounted volumes that are intended as APIs for programs to read, and
// system administrators at a shell; rather than something that should, for
// example, appear in a GUI. For example, the Linux `/proc` filesystem.
//
// The list of file system types considered ‘system’ ones may change over time.
func UnixIsSystemFSType(fsType string) bool {
	var _arg1 *C.char    // out
	var _cret C.gboolean // in

	_arg1 = (*C.char)(C.CString(fsType))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_unix_is_system_fs_type(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixMountAt gets a MountEntry for a given mount path. If @time_read is set,
// it will be filled with a unix timestamp for checking if the mounts have
// changed since with g_unix_mounts_changed_since().
//
// If more mounts have the same mount path, the last matching mount is returned.
func UnixMountAt(mountPath string) (uint64, UnixMountEntry) {
	var _arg1 *C.char            // out
	var _arg2 *C.guint64         // in
	var _cret *C.GUnixMountEntry // in

	_arg1 = (*C.char)(C.CString(mountPath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_unix_mount_at(_arg1, &_arg2)

	var _timeRead uint64               // out
	var _unixMountEntry UnixMountEntry // out

	_timeRead = uint64(_arg2)
	_unixMountEntry = (UnixMountEntry)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_unixMountEntry, func(v UnixMountEntry) {
		C.free(unsafe.Pointer(v))
	})

	return _timeRead, _unixMountEntry
}

// UnixMountCompare compares two unix mounts.
func UnixMountCompare(mount1 UnixMountEntry, mount2 UnixMountEntry) int {
	var _arg1 *C.GUnixMountEntry // out
	var _arg2 *C.GUnixMountEntry // out
	var _cret C.gint             // in

	_arg1 = (*C.GUnixMountEntry)(unsafe.Pointer(mount1))
	_arg2 = (*C.GUnixMountEntry)(unsafe.Pointer(mount2))

	_cret = C.g_unix_mount_compare(_arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// UnixMountCopy makes a copy of @mount_entry.
func UnixMountCopy(mountEntry UnixMountEntry) UnixMountEntry {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.GUnixMountEntry // in

	_arg1 = (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	_cret = C.g_unix_mount_copy(_arg1)

	var _unixMountEntry UnixMountEntry // out

	_unixMountEntry = (UnixMountEntry)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_unixMountEntry, func(v UnixMountEntry) {
		C.free(unsafe.Pointer(v))
	})

	return _unixMountEntry
}

// UnixMountFor gets a MountEntry for a given file path. If @time_read is set,
// it will be filled with a unix timestamp for checking if the mounts have
// changed since with g_unix_mounts_changed_since().
//
// If more mounts have the same mount path, the last matching mount is returned.
func UnixMountFor(filePath string) (uint64, UnixMountEntry) {
	var _arg1 *C.char            // out
	var _arg2 *C.guint64         // in
	var _cret *C.GUnixMountEntry // in

	_arg1 = (*C.char)(C.CString(filePath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_unix_mount_for(_arg1, &_arg2)

	var _timeRead uint64               // out
	var _unixMountEntry UnixMountEntry // out

	_timeRead = uint64(_arg2)
	_unixMountEntry = (UnixMountEntry)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_unixMountEntry, func(v UnixMountEntry) {
		C.free(unsafe.Pointer(v))
	})

	return _timeRead, _unixMountEntry
}

// UnixMountFree frees a unix mount.
func UnixMountFree(mountEntry UnixMountEntry) {
	var _arg1 *C.GUnixMountEntry // out

	_arg1 = (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	C.g_unix_mount_free(_arg1)
}

// UnixMountGetDevicePath gets the device path for a unix mount.
func UnixMountGetDevicePath(mountEntry UnixMountEntry) string {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.char            // in

	_arg1 = (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	_cret = C.g_unix_mount_get_device_path(_arg1)

	var _filename string // out

	_filename = C.GoString(_cret)

	return _filename
}

// UnixMountGetFSType gets the filesystem type for the unix mount.
func UnixMountGetFSType(mountEntry UnixMountEntry) string {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.char            // in

	_arg1 = (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	_cret = C.g_unix_mount_get_fs_type(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// UnixMountGetMountPath gets the mount path for a unix mount.
func UnixMountGetMountPath(mountEntry UnixMountEntry) string {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.char            // in

	_arg1 = (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	_cret = C.g_unix_mount_get_mount_path(_arg1)

	var _filename string // out

	_filename = C.GoString(_cret)

	return _filename
}

// UnixMountGetOptions gets a comma-separated list of mount options for the unix
// mount. For example, `rw,relatime,seclabel,data=ordered`.
//
// This is similar to g_unix_mount_point_get_options(), but it takes a
// MountEntry as an argument.
func UnixMountGetOptions(mountEntry UnixMountEntry) string {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.char            // in

	_arg1 = (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	_cret = C.g_unix_mount_get_options(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// UnixMountGetRootPath gets the root of the mount within the filesystem. This
// is useful e.g. for mounts created by bind operation, or btrfs subvolumes.
//
// For example, the root path is equal to "/" for mount created by "mount
// /dev/sda1 /mnt/foo" and "/bar" for "mount --bind /mnt/foo/bar /mnt/bar".
func UnixMountGetRootPath(mountEntry UnixMountEntry) string {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.char            // in

	_arg1 = (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	_cret = C.g_unix_mount_get_root_path(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// UnixMountGuessCanEject guesses whether a Unix mount can be ejected.
func UnixMountGuessCanEject(mountEntry UnixMountEntry) bool {
	var _arg1 *C.GUnixMountEntry // out
	var _cret C.gboolean         // in

	_arg1 = (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	_cret = C.g_unix_mount_guess_can_eject(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixMountGuessIcon guesses the icon of a Unix mount.
func UnixMountGuessIcon(mountEntry UnixMountEntry) Icon {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.GIcon           // in

	_arg1 = (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	_cret = C.g_unix_mount_guess_icon(_arg1)

	var _icon Icon // out

	_icon = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Icon)

	return _icon
}

// UnixMountGuessName guesses the name of a Unix mount. The result is a
// translated string.
func UnixMountGuessName(mountEntry UnixMountEntry) string {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.char            // in

	_arg1 = (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	_cret = C.g_unix_mount_guess_name(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// UnixMountGuessShouldDisplay guesses whether a Unix mount should be displayed
// in the UI.
func UnixMountGuessShouldDisplay(mountEntry UnixMountEntry) bool {
	var _arg1 *C.GUnixMountEntry // out
	var _cret C.gboolean         // in

	_arg1 = (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	_cret = C.g_unix_mount_guess_should_display(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixMountGuessSymbolicIcon guesses the symbolic icon of a Unix mount.
func UnixMountGuessSymbolicIcon(mountEntry UnixMountEntry) Icon {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.GIcon           // in

	_arg1 = (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	_cret = C.g_unix_mount_guess_symbolic_icon(_arg1)

	var _icon Icon // out

	_icon = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Icon)

	return _icon
}

// UnixMountIsReadonly checks if a unix mount is mounted read only.
func UnixMountIsReadonly(mountEntry UnixMountEntry) bool {
	var _arg1 *C.GUnixMountEntry // out
	var _cret C.gboolean         // in

	_arg1 = (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	_cret = C.g_unix_mount_is_readonly(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixMountIsSystemInternal checks if a Unix mount is a system mount. This is
// the Boolean OR of g_unix_is_system_fs_type(), g_unix_is_system_device_path()
// and g_unix_is_mount_path_system_internal() on @mount_entry’s properties.
//
// The definition of what a ‘system’ mount entry is may change over time as new
// file system types and device paths are ignored.
func UnixMountIsSystemInternal(mountEntry UnixMountEntry) bool {
	var _arg1 *C.GUnixMountEntry // out
	var _cret C.gboolean         // in

	_arg1 = (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	_cret = C.g_unix_mount_is_system_internal(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixMountPointsChangedSince checks if the unix mount points have changed
// since a given unix time.
func UnixMountPointsChangedSince(time uint64) bool {
	var _arg1 C.guint64  // out
	var _cret C.gboolean // in

	_arg1 = C.guint64(time)

	_cret = C.g_unix_mount_points_changed_since(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixMountsChangedSince checks if the unix mounts have changed since a given
// unix time.
func UnixMountsChangedSince(time uint64) bool {
	var _arg1 C.guint64  // out
	var _cret C.gboolean // in

	_arg1 = C.guint64(time)

	_cret = C.g_unix_mounts_changed_since(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixMountMonitor watches Mounts for changes.
type UnixMountMonitor interface {
	gextras.Objector

	// SetRateLimitUnixMountMonitor: this function does nothing.
	//
	// Before 2.44, this was a partially-effective way of controlling the rate
	// at which events would be reported under some uncommon circumstances.
	// Since @mount_monitor is a singleton, it also meant that calling this
	// function would have side effects for other users of the monitor.
	//
	// Deprecated: since version 2.44.
	SetRateLimitUnixMountMonitor(limitMsec int)
}

// unixMountMonitor implements the UnixMountMonitor class.
type unixMountMonitor struct {
	gextras.Objector
}

// WrapUnixMountMonitor wraps a GObject to the right type. It is
// primarily used internally.
func WrapUnixMountMonitor(obj *externglib.Object) UnixMountMonitor {
	return unixMountMonitor{
		Objector: obj,
	}
}

func marshalUnixMountMonitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapUnixMountMonitor(obj), nil
}

// NewUnixMountMonitor: deprecated alias for g_unix_mount_monitor_get().
//
// This function was never a true constructor, which is why it was renamed.
//
// Deprecated: since version 2.44.
func NewUnixMountMonitor() UnixMountMonitor {
	var _cret *C.GUnixMountMonitor // in

	_cret = C.g_unix_mount_monitor_new()

	var _unixMountMonitor UnixMountMonitor // out

	_unixMountMonitor = WrapUnixMountMonitor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _unixMountMonitor
}

func (m unixMountMonitor) SetRateLimitUnixMountMonitor(limitMsec int) {
	var _arg0 *C.GUnixMountMonitor // out
	var _arg1 C.int                // out

	_arg0 = (*C.GUnixMountMonitor)(unsafe.Pointer(m.Native()))
	_arg1 = C.int(limitMsec)

	C.g_unix_mount_monitor_set_rate_limit(_arg0, _arg1)
}

// UnixMountEntry defines a Unix mount entry (e.g.
// <filename>/media/cdrom</filename>). This corresponds roughly to a mtab entry.
type UnixMountEntry struct {
	native C.GUnixMountEntry
}

// WrapUnixMountEntry wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapUnixMountEntry(ptr unsafe.Pointer) *UnixMountEntry {
	return (*UnixMountEntry)(ptr)
}

func marshalUnixMountEntry(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*UnixMountEntry)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (u *UnixMountEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&u.native)
}

// UnixMountPoint defines a Unix mount point (e.g. <filename>/dev</filename>).
// This corresponds roughly to a fstab entry.
type UnixMountPoint struct {
	native C.GUnixMountPoint
}

// WrapUnixMountPoint wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapUnixMountPoint(ptr unsafe.Pointer) *UnixMountPoint {
	return (*UnixMountPoint)(ptr)
}

func marshalUnixMountPoint(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*UnixMountPoint)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (u *UnixMountPoint) Native() unsafe.Pointer {
	return unsafe.Pointer(&u.native)
}

// Compare compares two unix mount points.
func (m *UnixMountPoint) Compare(mount2 UnixMountPoint) int {
	var _arg0 *C.GUnixMountPoint // out
	var _arg1 *C.GUnixMountPoint // out
	var _cret C.gint             // in

	_arg0 = (*C.GUnixMountPoint)(unsafe.Pointer(m))
	_arg1 = (*C.GUnixMountPoint)(unsafe.Pointer(mount2))

	_cret = C.g_unix_mount_point_compare(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Copy makes a copy of @mount_point.
func (m *UnixMountPoint) Copy() UnixMountPoint {
	var _arg0 *C.GUnixMountPoint // out
	var _cret *C.GUnixMountPoint // in

	_arg0 = (*C.GUnixMountPoint)(unsafe.Pointer(m))

	_cret = C.g_unix_mount_point_copy(_arg0)

	var _unixMountPoint UnixMountPoint // out

	_unixMountPoint = (UnixMountPoint)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_unixMountPoint, func(v UnixMountPoint) {
		C.g_unix_mount_point_free((*C.GUnixMountPoint)(unsafe.Pointer(v)))
	})

	return _unixMountPoint
}

// Free frees a unix mount point.
func (m *UnixMountPoint) Free() {
	var _arg0 *C.GUnixMountPoint // out

	_arg0 = (*C.GUnixMountPoint)(unsafe.Pointer(m))

	C.g_unix_mount_point_free(_arg0)
}

// DevicePath gets the device path for a unix mount point.
func (m *UnixMountPoint) DevicePath() string {
	var _arg0 *C.GUnixMountPoint // out
	var _cret *C.char            // in

	_arg0 = (*C.GUnixMountPoint)(unsafe.Pointer(m))

	_cret = C.g_unix_mount_point_get_device_path(_arg0)

	var _filename string // out

	_filename = C.GoString(_cret)

	return _filename
}

// FSType gets the file system type for the mount point.
func (m *UnixMountPoint) FSType() string {
	var _arg0 *C.GUnixMountPoint // out
	var _cret *C.char            // in

	_arg0 = (*C.GUnixMountPoint)(unsafe.Pointer(m))

	_cret = C.g_unix_mount_point_get_fs_type(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// MountPath gets the mount path for a unix mount point.
func (m *UnixMountPoint) MountPath() string {
	var _arg0 *C.GUnixMountPoint // out
	var _cret *C.char            // in

	_arg0 = (*C.GUnixMountPoint)(unsafe.Pointer(m))

	_cret = C.g_unix_mount_point_get_mount_path(_arg0)

	var _filename string // out

	_filename = C.GoString(_cret)

	return _filename
}

// Options gets the options for the mount point.
func (m *UnixMountPoint) Options() string {
	var _arg0 *C.GUnixMountPoint // out
	var _cret *C.char            // in

	_arg0 = (*C.GUnixMountPoint)(unsafe.Pointer(m))

	_cret = C.g_unix_mount_point_get_options(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// GuessCanEject guesses whether a Unix mount point can be ejected.
func (m *UnixMountPoint) GuessCanEject() bool {
	var _arg0 *C.GUnixMountPoint // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GUnixMountPoint)(unsafe.Pointer(m))

	_cret = C.g_unix_mount_point_guess_can_eject(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GuessIcon guesses the icon of a Unix mount point.
func (m *UnixMountPoint) GuessIcon() Icon {
	var _arg0 *C.GUnixMountPoint // out
	var _cret *C.GIcon           // in

	_arg0 = (*C.GUnixMountPoint)(unsafe.Pointer(m))

	_cret = C.g_unix_mount_point_guess_icon(_arg0)

	var _icon Icon // out

	_icon = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Icon)

	return _icon
}

// GuessName guesses the name of a Unix mount point. The result is a translated
// string.
func (m *UnixMountPoint) GuessName() string {
	var _arg0 *C.GUnixMountPoint // out
	var _cret *C.char            // in

	_arg0 = (*C.GUnixMountPoint)(unsafe.Pointer(m))

	_cret = C.g_unix_mount_point_guess_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// GuessSymbolicIcon guesses the symbolic icon of a Unix mount point.
func (m *UnixMountPoint) GuessSymbolicIcon() Icon {
	var _arg0 *C.GUnixMountPoint // out
	var _cret *C.GIcon           // in

	_arg0 = (*C.GUnixMountPoint)(unsafe.Pointer(m))

	_cret = C.g_unix_mount_point_guess_symbolic_icon(_arg0)

	var _icon Icon // out

	_icon = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Icon)

	return _icon
}

// IsLoopback checks if a unix mount point is a loopback device.
func (m *UnixMountPoint) IsLoopback() bool {
	var _arg0 *C.GUnixMountPoint // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GUnixMountPoint)(unsafe.Pointer(m))

	_cret = C.g_unix_mount_point_is_loopback(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsReadonly checks if a unix mount point is read only.
func (m *UnixMountPoint) IsReadonly() bool {
	var _arg0 *C.GUnixMountPoint // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GUnixMountPoint)(unsafe.Pointer(m))

	_cret = C.g_unix_mount_point_is_readonly(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsUserMountable checks if a unix mount point is mountable by the user.
func (m *UnixMountPoint) IsUserMountable() bool {
	var _arg0 *C.GUnixMountPoint // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GUnixMountPoint)(unsafe.Pointer(m))

	_cret = C.g_unix_mount_point_is_user_mountable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
