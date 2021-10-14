// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_unix_mount_monitor_get_type()), F: marshalUnixMountMonitorrer},
		{T: externglib.Type(C.g_unix_mount_entry_get_type()), F: marshalUnixMountEntry},
		{T: externglib.Type(C.g_unix_mount_point_get_type()), F: marshalUnixMountPoint},
	})
}

// UnixIsMountPathSystemInternal determines if mount_path is considered an
// implementation of the OS. This is primarily used for hiding mountable and
// mounted volumes that only are used in the OS and has little to no relevance
// to the casual user.
//
// The function takes the following parameters:
//
//    - mountPath: mount path, e.g. /media/disk or /usr.
//
func UnixIsMountPathSystemInternal(mountPath string) bool {
	var _arg1 *C.char    // out
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(mountPath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_unix_is_mount_path_system_internal(_arg1)
	runtime.KeepAlive(mountPath)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixIsSystemDevicePath determines if device_path is considered a block device
// path which is only used in implementation of the OS. This is primarily used
// for hiding mounted volumes that are intended as APIs for programs to read,
// and system administrators at a shell; rather than something that should, for
// example, appear in a GUI. For example, the Linux /proc filesystem.
//
// The list of device paths considered ‘system’ ones may change over time.
//
// The function takes the following parameters:
//
//    - devicePath: device path, e.g. /dev/loop0 or nfsd.
//
func UnixIsSystemDevicePath(devicePath string) bool {
	var _arg1 *C.char    // out
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(devicePath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_unix_is_system_device_path(_arg1)
	runtime.KeepAlive(devicePath)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixIsSystemFSType determines if fs_type is considered a type of file system
// which is only used in implementation of the OS. This is primarily used for
// hiding mounted volumes that are intended as APIs for programs to read, and
// system administrators at a shell; rather than something that should, for
// example, appear in a GUI. For example, the Linux /proc filesystem.
//
// The list of file system types considered ‘system’ ones may change over time.
//
// The function takes the following parameters:
//
//    - fsType: file system type, e.g. procfs or tmpfs.
//
func UnixIsSystemFSType(fsType string) bool {
	var _arg1 *C.char    // out
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(fsType)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_unix_is_system_fs_type(_arg1)
	runtime.KeepAlive(fsType)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixMountAt gets a MountEntry for a given mount path. If time_read is set, it
// will be filled with a unix timestamp for checking if the mounts have changed
// since with g_unix_mounts_changed_since().
//
// If more mounts have the same mount path, the last matching mount is returned.
//
// The function takes the following parameters:
//
//    - mountPath: path for a possible unix mount.
//
func UnixMountAt(mountPath string) (uint64, *UnixMountEntry) {
	var _arg1 *C.char            // out
	var _arg2 C.guint64          // in
	var _cret *C.GUnixMountEntry // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(mountPath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_unix_mount_at(_arg1, &_arg2)
	runtime.KeepAlive(mountPath)

	var _timeRead uint64                // out
	var _unixMountEntry *UnixMountEntry // out

	_timeRead = uint64(_arg2)
	_unixMountEntry = (*UnixMountEntry)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_unixMountEntry)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _timeRead, _unixMountEntry
}

// UnixMountCompare compares two unix mounts.
//
// The function takes the following parameters:
//
//    - mount1: first MountEntry to compare.
//    - mount2: second MountEntry to compare.
//
func UnixMountCompare(mount1, mount2 *UnixMountEntry) int {
	var _arg1 *C.GUnixMountEntry // out
	var _arg2 *C.GUnixMountEntry // out
	var _cret C.gint             // in

	_arg1 = (*C.GUnixMountEntry)(gextras.StructNative(unsafe.Pointer(mount1)))
	_arg2 = (*C.GUnixMountEntry)(gextras.StructNative(unsafe.Pointer(mount2)))

	_cret = C.g_unix_mount_compare(_arg1, _arg2)
	runtime.KeepAlive(mount1)
	runtime.KeepAlive(mount2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// UnixMountCopy makes a copy of mount_entry.
//
// The function takes the following parameters:
//
//    - mountEntry: MountEntry.
//
func UnixMountCopy(mountEntry *UnixMountEntry) *UnixMountEntry {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.GUnixMountEntry // in

	_arg1 = (*C.GUnixMountEntry)(gextras.StructNative(unsafe.Pointer(mountEntry)))

	_cret = C.g_unix_mount_copy(_arg1)
	runtime.KeepAlive(mountEntry)

	var _unixMountEntry *UnixMountEntry // out

	_unixMountEntry = (*UnixMountEntry)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_unixMountEntry)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _unixMountEntry
}

// UnixMountFor gets a MountEntry for a given file path. If time_read is set, it
// will be filled with a unix timestamp for checking if the mounts have changed
// since with g_unix_mounts_changed_since().
//
// If more mounts have the same mount path, the last matching mount is returned.
//
// The function takes the following parameters:
//
//    - filePath: file path on some unix mount.
//
func UnixMountFor(filePath string) (uint64, *UnixMountEntry) {
	var _arg1 *C.char            // out
	var _arg2 C.guint64          // in
	var _cret *C.GUnixMountEntry // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filePath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_unix_mount_for(_arg1, &_arg2)
	runtime.KeepAlive(filePath)

	var _timeRead uint64                // out
	var _unixMountEntry *UnixMountEntry // out

	_timeRead = uint64(_arg2)
	_unixMountEntry = (*UnixMountEntry)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_unixMountEntry)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _timeRead, _unixMountEntry
}

// UnixMountGetDevicePath gets the device path for a unix mount.
//
// The function takes the following parameters:
//
//    - mountEntry: Mount.
//
func UnixMountGetDevicePath(mountEntry *UnixMountEntry) string {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.char            // in

	_arg1 = (*C.GUnixMountEntry)(gextras.StructNative(unsafe.Pointer(mountEntry)))

	_cret = C.g_unix_mount_get_device_path(_arg1)
	runtime.KeepAlive(mountEntry)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// UnixMountGetFSType gets the filesystem type for the unix mount.
//
// The function takes the following parameters:
//
//    - mountEntry: Mount.
//
func UnixMountGetFSType(mountEntry *UnixMountEntry) string {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.char            // in

	_arg1 = (*C.GUnixMountEntry)(gextras.StructNative(unsafe.Pointer(mountEntry)))

	_cret = C.g_unix_mount_get_fs_type(_arg1)
	runtime.KeepAlive(mountEntry)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UnixMountGetMountPath gets the mount path for a unix mount.
//
// The function takes the following parameters:
//
//    - mountEntry: input MountEntry to get the mount path for.
//
func UnixMountGetMountPath(mountEntry *UnixMountEntry) string {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.char            // in

	_arg1 = (*C.GUnixMountEntry)(gextras.StructNative(unsafe.Pointer(mountEntry)))

	_cret = C.g_unix_mount_get_mount_path(_arg1)
	runtime.KeepAlive(mountEntry)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// UnixMountGetOptions gets a comma-separated list of mount options for the unix
// mount. For example, rw,relatime,seclabel,data=ordered.
//
// This is similar to g_unix_mount_point_get_options(), but it takes a
// MountEntry as an argument.
//
// The function takes the following parameters:
//
//    - mountEntry: MountEntry.
//
func UnixMountGetOptions(mountEntry *UnixMountEntry) string {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.char            // in

	_arg1 = (*C.GUnixMountEntry)(gextras.StructNative(unsafe.Pointer(mountEntry)))

	_cret = C.g_unix_mount_get_options(_arg1)
	runtime.KeepAlive(mountEntry)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// UnixMountGetRootPath gets the root of the mount within the filesystem. This
// is useful e.g. for mounts created by bind operation, or btrfs subvolumes.
//
// For example, the root path is equal to "/" for mount created by "mount
// /dev/sda1 /mnt/foo" and "/bar" for "mount --bind /mnt/foo/bar /mnt/bar".
//
// The function takes the following parameters:
//
//    - mountEntry: MountEntry.
//
func UnixMountGetRootPath(mountEntry *UnixMountEntry) string {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.char            // in

	_arg1 = (*C.GUnixMountEntry)(gextras.StructNative(unsafe.Pointer(mountEntry)))

	_cret = C.g_unix_mount_get_root_path(_arg1)
	runtime.KeepAlive(mountEntry)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// UnixMountGuessCanEject guesses whether a Unix mount can be ejected.
//
// The function takes the following parameters:
//
//    - mountEntry: MountEntry.
//
func UnixMountGuessCanEject(mountEntry *UnixMountEntry) bool {
	var _arg1 *C.GUnixMountEntry // out
	var _cret C.gboolean         // in

	_arg1 = (*C.GUnixMountEntry)(gextras.StructNative(unsafe.Pointer(mountEntry)))

	_cret = C.g_unix_mount_guess_can_eject(_arg1)
	runtime.KeepAlive(mountEntry)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixMountGuessIcon guesses the icon of a Unix mount.
//
// The function takes the following parameters:
//
//    - mountEntry: MountEntry.
//
func UnixMountGuessIcon(mountEntry *UnixMountEntry) Iconner {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.GIcon           // in

	_arg1 = (*C.GUnixMountEntry)(gextras.StructNative(unsafe.Pointer(mountEntry)))

	_cret = C.g_unix_mount_guess_icon(_arg1)
	runtime.KeepAlive(mountEntry)

	var _icon Iconner // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Iconner is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(Iconner)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.Iconner")
		}
		_icon = rv
	}

	return _icon
}

// UnixMountGuessName guesses the name of a Unix mount. The result is a
// translated string.
//
// The function takes the following parameters:
//
//    - mountEntry: MountEntry.
//
func UnixMountGuessName(mountEntry *UnixMountEntry) string {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.char            // in

	_arg1 = (*C.GUnixMountEntry)(gextras.StructNative(unsafe.Pointer(mountEntry)))

	_cret = C.g_unix_mount_guess_name(_arg1)
	runtime.KeepAlive(mountEntry)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// UnixMountGuessShouldDisplay guesses whether a Unix mount should be displayed
// in the UI.
//
// The function takes the following parameters:
//
//    - mountEntry: MountEntry.
//
func UnixMountGuessShouldDisplay(mountEntry *UnixMountEntry) bool {
	var _arg1 *C.GUnixMountEntry // out
	var _cret C.gboolean         // in

	_arg1 = (*C.GUnixMountEntry)(gextras.StructNative(unsafe.Pointer(mountEntry)))

	_cret = C.g_unix_mount_guess_should_display(_arg1)
	runtime.KeepAlive(mountEntry)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixMountGuessSymbolicIcon guesses the symbolic icon of a Unix mount.
//
// The function takes the following parameters:
//
//    - mountEntry: MountEntry.
//
func UnixMountGuessSymbolicIcon(mountEntry *UnixMountEntry) Iconner {
	var _arg1 *C.GUnixMountEntry // out
	var _cret *C.GIcon           // in

	_arg1 = (*C.GUnixMountEntry)(gextras.StructNative(unsafe.Pointer(mountEntry)))

	_cret = C.g_unix_mount_guess_symbolic_icon(_arg1)
	runtime.KeepAlive(mountEntry)

	var _icon Iconner // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Iconner is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(Iconner)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.Iconner")
		}
		_icon = rv
	}

	return _icon
}

// UnixMountIsReadonly checks if a unix mount is mounted read only.
//
// The function takes the following parameters:
//
//    - mountEntry: Mount.
//
func UnixMountIsReadonly(mountEntry *UnixMountEntry) bool {
	var _arg1 *C.GUnixMountEntry // out
	var _cret C.gboolean         // in

	_arg1 = (*C.GUnixMountEntry)(gextras.StructNative(unsafe.Pointer(mountEntry)))

	_cret = C.g_unix_mount_is_readonly(_arg1)
	runtime.KeepAlive(mountEntry)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixMountIsSystemInternal checks if a Unix mount is a system mount. This is
// the Boolean OR of g_unix_is_system_fs_type(), g_unix_is_system_device_path()
// and g_unix_is_mount_path_system_internal() on mount_entry’s properties.
//
// The definition of what a ‘system’ mount entry is may change over time as new
// file system types and device paths are ignored.
//
// The function takes the following parameters:
//
//    - mountEntry: Mount.
//
func UnixMountIsSystemInternal(mountEntry *UnixMountEntry) bool {
	var _arg1 *C.GUnixMountEntry // out
	var _cret C.gboolean         // in

	_arg1 = (*C.GUnixMountEntry)(gextras.StructNative(unsafe.Pointer(mountEntry)))

	_cret = C.g_unix_mount_is_system_internal(_arg1)
	runtime.KeepAlive(mountEntry)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixMountPointsChangedSince checks if the unix mount points have changed
// since a given unix time.
//
// The function takes the following parameters:
//
//    - time: guint64 to contain a timestamp.
//
func UnixMountPointsChangedSince(time uint64) bool {
	var _arg1 C.guint64  // out
	var _cret C.gboolean // in

	_arg1 = C.guint64(time)

	_cret = C.g_unix_mount_points_changed_since(_arg1)
	runtime.KeepAlive(time)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixMountPointsGet gets a #GList of MountPoint containing the unix mount
// points. If time_read is set, it will be filled with the mount timestamp,
// allowing for checking if the mounts have changed with
// g_unix_mount_points_changed_since().
func UnixMountPointsGet() (uint64, []*UnixMountPoint) {
	var _arg1 C.guint64 // in
	var _cret *C.GList  // in

	_cret = C.g_unix_mount_points_get(&_arg1)

	var _timeRead uint64        // out
	var _list []*UnixMountPoint // out

	_timeRead = uint64(_arg1)
	_list = make([]*UnixMountPoint, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GUnixMountPoint)(v)
		var dst *UnixMountPoint // out
		dst = (*UnixMountPoint)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_unix_mount_point_free((*C.GUnixMountPoint)(intern.C))
			},
		)
		_list = append(_list, dst)
	})

	return _timeRead, _list
}

// UnixMountsChangedSince checks if the unix mounts have changed since a given
// unix time.
//
// The function takes the following parameters:
//
//    - time: guint64 to contain a timestamp.
//
func UnixMountsChangedSince(time uint64) bool {
	var _arg1 C.guint64  // out
	var _cret C.gboolean // in

	_arg1 = C.guint64(time)

	_cret = C.g_unix_mounts_changed_since(_arg1)
	runtime.KeepAlive(time)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixMountsGet gets a #GList of MountEntry containing the unix mounts. If
// time_read is set, it will be filled with the mount timestamp, allowing for
// checking if the mounts have changed with g_unix_mounts_changed_since().
func UnixMountsGet() (uint64, []*UnixMountEntry) {
	var _arg1 C.guint64 // in
	var _cret *C.GList  // in

	_cret = C.g_unix_mounts_get(&_arg1)

	var _timeRead uint64        // out
	var _list []*UnixMountEntry // out

	_timeRead = uint64(_arg1)
	_list = make([]*UnixMountEntry, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GUnixMountEntry)(v)
		var dst *UnixMountEntry // out
		dst = (*UnixMountEntry)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
		_list = append(_list, dst)
	})

	return _timeRead, _list
}

// UnixMountMonitor watches Mounts for changes.
type UnixMountMonitor struct {
	*externglib.Object
}

func wrapUnixMountMonitor(obj *externglib.Object) *UnixMountMonitor {
	return &UnixMountMonitor{
		Object: obj,
	}
}

func marshalUnixMountMonitorrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapUnixMountMonitor(obj), nil
}

// NewUnixMountMonitor: deprecated alias for g_unix_mount_monitor_get().
//
// This function was never a true constructor, which is why it was renamed.
//
// Deprecated: Use g_unix_mount_monitor_get() instead.
func NewUnixMountMonitor() *UnixMountMonitor {
	var _cret *C.GUnixMountMonitor // in

	_cret = C.g_unix_mount_monitor_new()

	var _unixMountMonitor *UnixMountMonitor // out

	_unixMountMonitor = wrapUnixMountMonitor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _unixMountMonitor
}

// SetRateLimit: this function does nothing.
//
// Before 2.44, this was a partially-effective way of controlling the rate at
// which events would be reported under some uncommon circumstances. Since
// mount_monitor is a singleton, it also meant that calling this function would
// have side effects for other users of the monitor.
//
// Deprecated: This function does nothing. Don't call it.
//
// The function takes the following parameters:
//
//    - limitMsec: integer with the limit in milliseconds to poll for changes.
//
func (mountMonitor *UnixMountMonitor) SetRateLimit(limitMsec int) {
	var _arg0 *C.GUnixMountMonitor // out
	var _arg1 C.int                // out

	_arg0 = (*C.GUnixMountMonitor)(unsafe.Pointer(mountMonitor.Native()))
	_arg1 = C.int(limitMsec)

	C.g_unix_mount_monitor_set_rate_limit(_arg0, _arg1)
	runtime.KeepAlive(mountMonitor)
	runtime.KeepAlive(limitMsec)
}

// ConnectMountpointsChanged: emitted when the unix mount points have changed.
func (mountMonitor *UnixMountMonitor) ConnectMountpointsChanged(f func()) externglib.SignalHandle {
	return mountMonitor.Connect("mountpoints-changed", f)
}

// ConnectMountsChanged: emitted when the unix mounts have changed.
func (mountMonitor *UnixMountMonitor) ConnectMountsChanged(f func()) externglib.SignalHandle {
	return mountMonitor.Connect("mounts-changed", f)
}

// UnixMountMonitorGet gets the MountMonitor for the current thread-default main
// context.
//
// The mount monitor can be used to monitor for changes to the list of mounted
// filesystems as well as the list of mount points (ie: fstab entries).
//
// You must only call g_object_unref() on the return value from under the same
// main context as you called this function.
func UnixMountMonitorGet() *UnixMountMonitor {
	var _cret *C.GUnixMountMonitor // in

	_cret = C.g_unix_mount_monitor_get()

	var _unixMountMonitor *UnixMountMonitor // out

	_unixMountMonitor = wrapUnixMountMonitor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _unixMountMonitor
}

// UnixMountEntry defines a Unix mount entry (e.g.
// <filename>/media/cdrom</filename>). This corresponds roughly to a mtab entry.
//
// An instance of this type is always passed by reference.
type UnixMountEntry struct {
	*unixMountEntry
}

// unixMountEntry is the struct that's finalized.
type unixMountEntry struct {
	native *C.GUnixMountEntry
}

func marshalUnixMountEntry(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &UnixMountEntry{&unixMountEntry{(*C.GUnixMountEntry)(unsafe.Pointer(b))}}, nil
}

// UnixMountPoint defines a Unix mount point (e.g. <filename>/dev</filename>).
// This corresponds roughly to a fstab entry.
//
// An instance of this type is always passed by reference.
type UnixMountPoint struct {
	*unixMountPoint
}

// unixMountPoint is the struct that's finalized.
type unixMountPoint struct {
	native *C.GUnixMountPoint
}

func marshalUnixMountPoint(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &UnixMountPoint{&unixMountPoint{(*C.GUnixMountPoint)(unsafe.Pointer(b))}}, nil
}

// Compare compares two unix mount points.
func (mount1 *UnixMountPoint) Compare(mount2 *UnixMountPoint) int {
	var _arg0 *C.GUnixMountPoint // out
	var _arg1 *C.GUnixMountPoint // out
	var _cret C.gint             // in

	_arg0 = (*C.GUnixMountPoint)(gextras.StructNative(unsafe.Pointer(mount1)))
	_arg1 = (*C.GUnixMountPoint)(gextras.StructNative(unsafe.Pointer(mount2)))

	_cret = C.g_unix_mount_point_compare(_arg0, _arg1)
	runtime.KeepAlive(mount1)
	runtime.KeepAlive(mount2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Copy makes a copy of mount_point.
func (mountPoint *UnixMountPoint) Copy() *UnixMountPoint {
	var _arg0 *C.GUnixMountPoint // out
	var _cret *C.GUnixMountPoint // in

	_arg0 = (*C.GUnixMountPoint)(gextras.StructNative(unsafe.Pointer(mountPoint)))

	_cret = C.g_unix_mount_point_copy(_arg0)
	runtime.KeepAlive(mountPoint)

	var _unixMountPoint *UnixMountPoint // out

	_unixMountPoint = (*UnixMountPoint)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_unixMountPoint)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_unix_mount_point_free((*C.GUnixMountPoint)(intern.C))
		},
	)

	return _unixMountPoint
}

// DevicePath gets the device path for a unix mount point.
func (mountPoint *UnixMountPoint) DevicePath() string {
	var _arg0 *C.GUnixMountPoint // out
	var _cret *C.char            // in

	_arg0 = (*C.GUnixMountPoint)(gextras.StructNative(unsafe.Pointer(mountPoint)))

	_cret = C.g_unix_mount_point_get_device_path(_arg0)
	runtime.KeepAlive(mountPoint)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// FSType gets the file system type for the mount point.
func (mountPoint *UnixMountPoint) FSType() string {
	var _arg0 *C.GUnixMountPoint // out
	var _cret *C.char            // in

	_arg0 = (*C.GUnixMountPoint)(gextras.StructNative(unsafe.Pointer(mountPoint)))

	_cret = C.g_unix_mount_point_get_fs_type(_arg0)
	runtime.KeepAlive(mountPoint)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// MountPath gets the mount path for a unix mount point.
func (mountPoint *UnixMountPoint) MountPath() string {
	var _arg0 *C.GUnixMountPoint // out
	var _cret *C.char            // in

	_arg0 = (*C.GUnixMountPoint)(gextras.StructNative(unsafe.Pointer(mountPoint)))

	_cret = C.g_unix_mount_point_get_mount_path(_arg0)
	runtime.KeepAlive(mountPoint)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// Options gets the options for the mount point.
func (mountPoint *UnixMountPoint) Options() string {
	var _arg0 *C.GUnixMountPoint // out
	var _cret *C.char            // in

	_arg0 = (*C.GUnixMountPoint)(gextras.StructNative(unsafe.Pointer(mountPoint)))

	_cret = C.g_unix_mount_point_get_options(_arg0)
	runtime.KeepAlive(mountPoint)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// GuessCanEject guesses whether a Unix mount point can be ejected.
func (mountPoint *UnixMountPoint) GuessCanEject() bool {
	var _arg0 *C.GUnixMountPoint // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GUnixMountPoint)(gextras.StructNative(unsafe.Pointer(mountPoint)))

	_cret = C.g_unix_mount_point_guess_can_eject(_arg0)
	runtime.KeepAlive(mountPoint)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GuessIcon guesses the icon of a Unix mount point.
func (mountPoint *UnixMountPoint) GuessIcon() Iconner {
	var _arg0 *C.GUnixMountPoint // out
	var _cret *C.GIcon           // in

	_arg0 = (*C.GUnixMountPoint)(gextras.StructNative(unsafe.Pointer(mountPoint)))

	_cret = C.g_unix_mount_point_guess_icon(_arg0)
	runtime.KeepAlive(mountPoint)

	var _icon Iconner // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Iconner is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(Iconner)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.Iconner")
		}
		_icon = rv
	}

	return _icon
}

// GuessName guesses the name of a Unix mount point. The result is a translated
// string.
func (mountPoint *UnixMountPoint) GuessName() string {
	var _arg0 *C.GUnixMountPoint // out
	var _cret *C.char            // in

	_arg0 = (*C.GUnixMountPoint)(gextras.StructNative(unsafe.Pointer(mountPoint)))

	_cret = C.g_unix_mount_point_guess_name(_arg0)
	runtime.KeepAlive(mountPoint)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// GuessSymbolicIcon guesses the symbolic icon of a Unix mount point.
func (mountPoint *UnixMountPoint) GuessSymbolicIcon() Iconner {
	var _arg0 *C.GUnixMountPoint // out
	var _cret *C.GIcon           // in

	_arg0 = (*C.GUnixMountPoint)(gextras.StructNative(unsafe.Pointer(mountPoint)))

	_cret = C.g_unix_mount_point_guess_symbolic_icon(_arg0)
	runtime.KeepAlive(mountPoint)

	var _icon Iconner // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Iconner is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(Iconner)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.Iconner")
		}
		_icon = rv
	}

	return _icon
}

// IsLoopback checks if a unix mount point is a loopback device.
func (mountPoint *UnixMountPoint) IsLoopback() bool {
	var _arg0 *C.GUnixMountPoint // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GUnixMountPoint)(gextras.StructNative(unsafe.Pointer(mountPoint)))

	_cret = C.g_unix_mount_point_is_loopback(_arg0)
	runtime.KeepAlive(mountPoint)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsReadonly checks if a unix mount point is read only.
func (mountPoint *UnixMountPoint) IsReadonly() bool {
	var _arg0 *C.GUnixMountPoint // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GUnixMountPoint)(gextras.StructNative(unsafe.Pointer(mountPoint)))

	_cret = C.g_unix_mount_point_is_readonly(_arg0)
	runtime.KeepAlive(mountPoint)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsUserMountable checks if a unix mount point is mountable by the user.
func (mountPoint *UnixMountPoint) IsUserMountable() bool {
	var _arg0 *C.GUnixMountPoint // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GUnixMountPoint)(gextras.StructNative(unsafe.Pointer(mountPoint)))

	_cret = C.g_unix_mount_point_is_user_mountable(_arg0)
	runtime.KeepAlive(mountPoint)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnixMountPointAt gets a MountPoint for a given mount path. If time_read is
// set, it will be filled with a unix timestamp for checking if the mount points
// have changed since with g_unix_mount_points_changed_since().
//
// If more mount points have the same mount path, the last matching mount point
// is returned.
//
// The function takes the following parameters:
//
//    - mountPath: path for a possible unix mount point.
//
func UnixMountPointAt(mountPath string) (uint64, *UnixMountPoint) {
	var _arg1 *C.char            // out
	var _arg2 C.guint64          // in
	var _cret *C.GUnixMountPoint // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(mountPath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_unix_mount_point_at(_arg1, &_arg2)
	runtime.KeepAlive(mountPath)

	var _timeRead uint64                // out
	var _unixMountPoint *UnixMountPoint // out

	_timeRead = uint64(_arg2)
	if _cret != nil {
		_unixMountPoint = (*UnixMountPoint)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_unixMountPoint)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_unix_mount_point_free((*C.GUnixMountPoint)(intern.C))
			},
		)
	}

	return _timeRead, _unixMountPoint
}
