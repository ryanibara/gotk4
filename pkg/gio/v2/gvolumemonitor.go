// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_volume_monitor_get_type()), F: marshalVolumeMonitorrer},
	})
}

// VolumeMonitorOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type VolumeMonitorOverrider interface {
	DriveChanged(drive Driver)
	DriveConnected(drive Driver)
	DriveDisconnected(drive Driver)
	DriveEjectButton(drive Driver)
	DriveStopButton(drive Driver)
	// ConnectedDrives gets a list of drives connected to the system.
	//
	// The returned list should be freed with g_list_free(), after its elements
	// have been unreffed with g_object_unref().
	ConnectedDrives() []Driver
	// MountForUUID finds a #GMount object by its UUID (see g_mount_get_uuid())
	MountForUUID(uuid string) Mounter
	// Mounts gets a list of the mounts on the system.
	//
	// The returned list should be freed with g_list_free(), after its elements
	// have been unreffed with g_object_unref().
	Mounts() []Mounter
	// VolumeForUUID finds a #GVolume object by its UUID (see
	// g_volume_get_uuid())
	VolumeForUUID(uuid string) Volumer
	// Volumes gets a list of the volumes on the system.
	//
	// The returned list should be freed with g_list_free(), after its elements
	// have been unreffed with g_object_unref().
	Volumes() []Volumer
	MountAdded(mount Mounter)
	MountChanged(mount Mounter)
	MountPreUnmount(mount Mounter)
	MountRemoved(mount Mounter)
	VolumeAdded(volume Volumer)
	VolumeChanged(volume Volumer)
	VolumeRemoved(volume Volumer)
}

// VolumeMonitor is for listing the user interesting devices and volumes on the
// computer. In other words, what a file selector or file manager would show in
// a sidebar.
//
// Monitor is not [thread-default-context
// aware][g-main-context-push-thread-default], and so should not be used other
// than from the main thread, with no thread-default-context active.
//
// In order to receive updates about volumes and mounts monitored through GVFS,
// a main loop must be running.
type VolumeMonitor struct {
	*externglib.Object
}

func wrapVolumeMonitor(obj *externglib.Object) *VolumeMonitor {
	return &VolumeMonitor{
		Object: obj,
	}
}

func marshalVolumeMonitorrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapVolumeMonitor(obj), nil
}

// ConnectedDrives gets a list of drives connected to the system.
//
// The returned list should be freed with g_list_free(), after its elements have
// been unreffed with g_object_unref().
func (volumeMonitor *VolumeMonitor) ConnectedDrives() []Driver {
	var _arg0 *C.GVolumeMonitor // out
	var _cret *C.GList          // in

	_arg0 = (*C.GVolumeMonitor)(unsafe.Pointer(volumeMonitor.Native()))

	_cret = C.g_volume_monitor_get_connected_drives(_arg0)

	var _list []Driver // out

	_list = make([]Driver, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GDrive)(v)
		var dst Driver // out
		dst = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(src)))).(Driver)
		_list = append(_list, dst)
	})

	return _list
}

// MountForUUID finds a #GMount object by its UUID (see g_mount_get_uuid())
func (volumeMonitor *VolumeMonitor) MountForUUID(uuid string) Mounter {
	var _arg0 *C.GVolumeMonitor // out
	var _arg1 *C.char           // out
	var _cret *C.GMount         // in

	_arg0 = (*C.GVolumeMonitor)(unsafe.Pointer(volumeMonitor.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uuid)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_volume_monitor_get_mount_for_uuid(_arg0, _arg1)

	var _mount Mounter // out

	if _cret != nil {
		_mount = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(Mounter)
	}

	return _mount
}

// Mounts gets a list of the mounts on the system.
//
// The returned list should be freed with g_list_free(), after its elements have
// been unreffed with g_object_unref().
func (volumeMonitor *VolumeMonitor) Mounts() []Mounter {
	var _arg0 *C.GVolumeMonitor // out
	var _cret *C.GList          // in

	_arg0 = (*C.GVolumeMonitor)(unsafe.Pointer(volumeMonitor.Native()))

	_cret = C.g_volume_monitor_get_mounts(_arg0)

	var _list []Mounter // out

	_list = make([]Mounter, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GMount)(v)
		var dst Mounter // out
		dst = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(src)))).(Mounter)
		_list = append(_list, dst)
	})

	return _list
}

// VolumeForUUID finds a #GVolume object by its UUID (see g_volume_get_uuid())
func (volumeMonitor *VolumeMonitor) VolumeForUUID(uuid string) Volumer {
	var _arg0 *C.GVolumeMonitor // out
	var _arg1 *C.char           // out
	var _cret *C.GVolume        // in

	_arg0 = (*C.GVolumeMonitor)(unsafe.Pointer(volumeMonitor.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uuid)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_volume_monitor_get_volume_for_uuid(_arg0, _arg1)

	var _volume Volumer // out

	if _cret != nil {
		_volume = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(Volumer)
	}

	return _volume
}

// Volumes gets a list of the volumes on the system.
//
// The returned list should be freed with g_list_free(), after its elements have
// been unreffed with g_object_unref().
func (volumeMonitor *VolumeMonitor) Volumes() []Volumer {
	var _arg0 *C.GVolumeMonitor // out
	var _cret *C.GList          // in

	_arg0 = (*C.GVolumeMonitor)(unsafe.Pointer(volumeMonitor.Native()))

	_cret = C.g_volume_monitor_get_volumes(_arg0)

	var _list []Volumer // out

	_list = make([]Volumer, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GVolume)(v)
		var dst Volumer // out
		dst = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(src)))).(Volumer)
		_list = append(_list, dst)
	})

	return _list
}

// VolumeMonitorAdoptOrphanMount: this function should be called by any Monitor
// implementation when a new #GMount object is created that is not associated
// with a #GVolume object. It must be called just before emitting the
// mount_added signal.
//
// If the return value is not NULL, the caller must associate the returned
// #GVolume object with the #GMount. This involves returning it in its
// g_mount_get_volume() implementation. The caller must also listen for the
// "removed" signal on the returned object and give up its reference when
// handling that signal
//
// Similarly, if implementing g_volume_monitor_adopt_orphan_mount(), the
// implementor must take a reference to mount and return it in its
// g_volume_get_mount() implemented. Also, the implementor must listen for the
// "unmounted" signal on mount and give up its reference upon handling that
// signal.
//
// There are two main use cases for this function.
//
// One is when implementing a user space file system driver that reads blocks of
// a block device that is already represented by the native volume monitor (for
// example a CD Audio file system driver). Such a driver will generate its own
// #GMount object that needs to be associated with the #GVolume object that
// represents the volume.
//
// The other is for implementing a Monitor whose sole purpose is to return
// #GVolume objects representing entries in the users "favorite servers" list or
// similar.
//
// Deprecated: Instead of using this function, Monitor implementations should
// instead create shadow mounts with the URI of the mount they intend to adopt.
// See the proxy volume monitor in gvfs for an example of this. Also see
// g_mount_is_shadowed(), g_mount_shadow() and g_mount_unshadow() functions.
func VolumeMonitorAdoptOrphanMount(mount Mounter) Volumer {
	var _arg1 *C.GMount  // out
	var _cret *C.GVolume // in

	_arg1 = (*C.GMount)(unsafe.Pointer(mount.Native()))

	_cret = C.g_volume_monitor_adopt_orphan_mount(_arg1)

	var _volume Volumer // out

	_volume = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(Volumer)

	return _volume
}

// VolumeMonitorGet gets the volume monitor used by gio.
func VolumeMonitorGet() *VolumeMonitor {
	var _cret *C.GVolumeMonitor // in

	_cret = C.g_volume_monitor_get()

	var _volumeMonitor *VolumeMonitor // out

	_volumeMonitor = wrapVolumeMonitor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _volumeMonitor
}
