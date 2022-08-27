// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_Mount_ConnectUnmounted(gpointer, guintptr);
// extern void _gotk4_gio2_Mount_ConnectPreUnmount(gpointer, guintptr);
// extern void _gotk4_gio2_Mount_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
// GDrive* _gotk4_gio2_Mount_virtual_get_drive(void* fnptr, GMount* arg0) {
//   return ((GDrive* (*)(GMount*))(fnptr))(arg0);
// };
// GFile* _gotk4_gio2_Mount_virtual_get_default_location(void* fnptr, GMount* arg0) {
//   return ((GFile* (*)(GMount*))(fnptr))(arg0);
// };
// GFile* _gotk4_gio2_Mount_virtual_get_root(void* fnptr, GMount* arg0) {
//   return ((GFile* (*)(GMount*))(fnptr))(arg0);
// };
// GIcon* _gotk4_gio2_Mount_virtual_get_icon(void* fnptr, GMount* arg0) {
//   return ((GIcon* (*)(GMount*))(fnptr))(arg0);
// };
// GVolume* _gotk4_gio2_Mount_virtual_get_volume(void* fnptr, GMount* arg0) {
//   return ((GVolume* (*)(GMount*))(fnptr))(arg0);
// };
// char* _gotk4_gio2_Mount_virtual_get_name(void* fnptr, GMount* arg0) {
//   return ((char* (*)(GMount*))(fnptr))(arg0);
// };
// char* _gotk4_gio2_Mount_virtual_get_uuid(void* fnptr, GMount* arg0) {
//   return ((char* (*)(GMount*))(fnptr))(arg0);
// };
// gboolean _gotk4_gio2_Mount_virtual_can_eject(void* fnptr, GMount* arg0) {
//   return ((gboolean (*)(GMount*))(fnptr))(arg0);
// };
// gboolean _gotk4_gio2_Mount_virtual_can_unmount(void* fnptr, GMount* arg0) {
//   return ((gboolean (*)(GMount*))(fnptr))(arg0);
// };
// gboolean _gotk4_gio2_Mount_virtual_eject_finish(void* fnptr, GMount* arg0, GAsyncResult* arg1, GError** arg2) {
//   return ((gboolean (*)(GMount*, GAsyncResult*, GError**))(fnptr))(arg0, arg1, arg2);
// };
// gboolean _gotk4_gio2_Mount_virtual_remount_finish(void* fnptr, GMount* arg0, GAsyncResult* arg1, GError** arg2) {
//   return ((gboolean (*)(GMount*, GAsyncResult*, GError**))(fnptr))(arg0, arg1, arg2);
// };
// gboolean _gotk4_gio2_Mount_virtual_unmount_finish(void* fnptr, GMount* arg0, GAsyncResult* arg1, GError** arg2) {
//   return ((gboolean (*)(GMount*, GAsyncResult*, GError**))(fnptr))(arg0, arg1, arg2);
// };
// void _gotk4_gio2_Mount_virtual_changed(void* fnptr, GMount* arg0) {
//   ((void (*)(GMount*))(fnptr))(arg0);
// };
// void _gotk4_gio2_Mount_virtual_eject(void* fnptr, GMount* arg0, GMountUnmountFlags arg1, GCancellable* arg2, GAsyncReadyCallback arg3, gpointer arg4) {
//   ((void (*)(GMount*, GMountUnmountFlags, GCancellable*, GAsyncReadyCallback, gpointer))(fnptr))(arg0, arg1, arg2, arg3, arg4);
// };
// void _gotk4_gio2_Mount_virtual_pre_unmount(void* fnptr, GMount* arg0) {
//   ((void (*)(GMount*))(fnptr))(arg0);
// };
// void _gotk4_gio2_Mount_virtual_remount(void* fnptr, GMount* arg0, GMountMountFlags arg1, GMountOperation* arg2, GCancellable* arg3, GAsyncReadyCallback arg4, gpointer arg5) {
//   ((void (*)(GMount*, GMountMountFlags, GMountOperation*, GCancellable*, GAsyncReadyCallback, gpointer))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5);
// };
// void _gotk4_gio2_Mount_virtual_unmount(void* fnptr, GMount* arg0, GMountUnmountFlags arg1, GCancellable* arg2, GAsyncReadyCallback arg3, gpointer arg4) {
//   ((void (*)(GMount*, GMountUnmountFlags, GCancellable*, GAsyncReadyCallback, gpointer))(fnptr))(arg0, arg1, arg2, arg3, arg4);
// };
// void _gotk4_gio2_Mount_virtual_unmounted(void* fnptr, GMount* arg0) {
//   ((void (*)(GMount*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeMount = coreglib.Type(C.g_mount_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMount, F: marshalMount},
	})
}

// Mount interface represents user-visible mounts. Note, when porting from
// GnomeVFS, #GMount is the moral equivalent of VFSVolume.
//
// #GMount is a "mounted" filesystem that you can access. Mounted is in quotes
// because it's not the same as a unix mount, it might be a gvfs mount, but you
// can still access the files on it if you use GIO. Might or might not be
// related to a volume object.
//
// Unmounting a #GMount instance is an asynchronous operation. For more
// information about asynchronous operations, see Result and #GTask. To unmount
// a #GMount instance, first call g_mount_unmount_with_operation() with (at
// least) the #GMount instance and a ReadyCallback. The callback will be fired
// when the operation has resolved (either with success or failure), and a
// Result structure will be passed to the callback. That callback should then
// call g_mount_unmount_with_operation_finish() with the #GMount and the Result
// data to see if the operation was completed successfully. If an error is
// present when g_mount_unmount_with_operation_finish() is called, then it will
// be filled with any error information.
//
// Mount wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Mount struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Mount)(nil)
)

// Mounter describes Mount's interface methods.
type Mounter interface {
	coreglib.Objector

	// CanEject checks if mount can be ejected.
	CanEject() bool
	// CanUnmount checks if mount can be unmounted.
	CanUnmount() bool
	// Eject ejects a mount.
	Eject(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback)
	// EjectFinish finishes ejecting a mount.
	EjectFinish(result AsyncResulter) error
	// EjectWithOperation ejects a mount.
	EjectWithOperation(ctx context.Context, flags MountUnmountFlags, mountOperation *MountOperation, callback AsyncReadyCallback)
	// EjectWithOperationFinish finishes ejecting a mount.
	EjectWithOperationFinish(result AsyncResulter) error
	// DefaultLocation gets the default location of mount.
	DefaultLocation() *File
	// Drive gets the drive for the mount.
	Drive() *Drive
	// Icon gets the icon for mount.
	Icon() *Icon
	// Name gets the name of mount.
	Name() string
	// Root gets the root directory on mount.
	Root() *File
	// SortKey gets the sort key for mount, if any.
	SortKey() string
	// SymbolicIcon gets the symbolic icon for mount.
	SymbolicIcon() *Icon
	// UUID gets the UUID for the mount.
	UUID() string
	// Volume gets the volume for the mount.
	Volume() *Volume
	// GuessContentType tries to guess the type of content stored on mount.
	GuessContentType(ctx context.Context, forceRescan bool, callback AsyncReadyCallback)
	// GuessContentTypeFinish finishes guessing content types of mount.
	GuessContentTypeFinish(result AsyncResulter) ([]string, error)
	// GuessContentTypeSync tries to guess the type of content stored on mount.
	GuessContentTypeSync(ctx context.Context, forceRescan bool) ([]string, error)
	// IsShadowed determines if mount is shadowed.
	IsShadowed() bool
	// Remount remounts a mount.
	Remount(ctx context.Context, flags MountMountFlags, mountOperation *MountOperation, callback AsyncReadyCallback)
	// RemountFinish finishes remounting a mount.
	RemountFinish(result AsyncResulter) error
	// Shadow increments the shadow count on mount.
	Shadow()
	// Unmount unmounts a mount.
	Unmount(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback)
	// UnmountFinish finishes unmounting a mount.
	UnmountFinish(result AsyncResulter) error
	// UnmountWithOperation unmounts a mount.
	UnmountWithOperation(ctx context.Context, flags MountUnmountFlags, mountOperation *MountOperation, callback AsyncReadyCallback)
	// UnmountWithOperationFinish finishes unmounting a mount.
	UnmountWithOperationFinish(result AsyncResulter) error
	// Unshadow decrements the shadow count on mount.
	Unshadow()

	// Changed is emitted when the mount has been changed.
	ConnectChanged(func()) coreglib.SignalHandle
	// Pre-unmount: this signal may be emitted when the #GMount is about to be
	// unmounted.
	ConnectPreUnmount(func()) coreglib.SignalHandle
	// Unmounted: this signal is emitted when the #GMount have been unmounted.
	ConnectUnmounted(func()) coreglib.SignalHandle
}

var _ Mounter = (*Mount)(nil)

func wrapMount(obj *coreglib.Object) *Mount {
	return &Mount{
		Object: obj,
	}
}

func marshalMount(p uintptr) (interface{}, error) {
	return wrapMount(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChanged is emitted when the mount has been changed.
func (mount *Mount) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(mount, "changed", false, unsafe.Pointer(C._gotk4_gio2_Mount_ConnectChanged), f)
}

// ConnectPreUnmount: this signal may be emitted when the #GMount is about to be
// unmounted.
//
// This signal depends on the backend and is only emitted if GIO was used to
// unmount.
func (mount *Mount) ConnectPreUnmount(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(mount, "pre-unmount", false, unsafe.Pointer(C._gotk4_gio2_Mount_ConnectPreUnmount), f)
}

// ConnectUnmounted: this signal is emitted when the #GMount have been
// unmounted. If the recipient is holding references to the object they should
// release them so the object can be finalized.
func (mount *Mount) ConnectUnmounted(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(mount, "unmounted", false, unsafe.Pointer(C._gotk4_gio2_Mount_ConnectUnmounted), f)
}

// CanEject checks if mount can be ejected.
//
// The function returns the following values:
//
//    - ok: TRUE if the mount can be ejected.
//
func (mount *Mount) CanEject() bool {
	var _arg0 *C.GMount  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C.g_mount_can_eject(_arg0)
	runtime.KeepAlive(mount)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanUnmount checks if mount can be unmounted.
//
// The function returns the following values:
//
//    - ok: TRUE if the mount can be unmounted.
//
func (mount *Mount) CanUnmount() bool {
	var _arg0 *C.GMount  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C.g_mount_can_unmount(_arg0)
	runtime.KeepAlive(mount)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Eject ejects a mount. This is an asynchronous operation, and is finished by
// calling g_mount_eject_finish() with the mount and Result data returned in the
// callback.
//
// Deprecated: Use g_mount_eject_with_operation() instead.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - flags affecting the unmount if required for eject.
//    - callback (optional) or NULL.
//
func (mount *Mount) Eject(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GMountUnmountFlags(flags)
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_mount_eject(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(callback)
}

// EjectFinish finishes ejecting a mount. If any errors occurred during the
// operation, error will be set to contain the errors and FALSE will be
// returned.
//
// Deprecated: Use g_mount_eject_with_operation_finish() instead.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (mount *Mount) EjectFinish(result AsyncResulter) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C.g_mount_eject_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// DefaultLocation gets the default location of mount. The default location of
// the given mount is a path that reflects the main entry point for the user
// (e.g. the home directory, or the root of the volume).
//
// The function returns the following values:
//
//    - file: #GFile. The returned object should be unreffed with
//      g_object_unref() when no longer needed.
//
func (mount *Mount) DefaultLocation() *File {
	var _arg0 *C.GMount // out
	var _cret *C.GFile  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C.g_mount_get_default_location(_arg0)
	runtime.KeepAlive(mount)

	var _file *File // out

	_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _file
}

// Drive gets the drive for the mount.
//
// This is a convenience method for getting the #GVolume and then using that
// object to get the #GDrive.
//
// The function returns the following values:
//
//    - drive (optional) or NULL if mount is not associated with a volume or a
//      drive. The returned object should be unreffed with g_object_unref() when
//      no longer needed.
//
func (mount *Mount) Drive() *Drive {
	var _arg0 *C.GMount // out
	var _cret *C.GDrive // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C.g_mount_get_drive(_arg0)
	runtime.KeepAlive(mount)

	var _drive *Drive // out

	if _cret != nil {
		_drive = wrapDrive(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _drive
}

// Icon gets the icon for mount.
//
// The function returns the following values:
//
//    - icon: #GIcon. The returned object should be unreffed with
//      g_object_unref() when no longer needed.
//
func (mount *Mount) Icon() *Icon {
	var _arg0 *C.GMount // out
	var _cret *C.GIcon  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C.g_mount_get_icon(_arg0)
	runtime.KeepAlive(mount)

	var _icon *Icon // out

	_icon = wrapIcon(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _icon
}

// Name gets the name of mount.
//
// The function returns the following values:
//
//    - utf8: name for the given mount. The returned string should be freed with
//      g_free() when no longer needed.
//
func (mount *Mount) Name() string {
	var _arg0 *C.GMount // out
	var _cret *C.char   // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C.g_mount_get_name(_arg0)
	runtime.KeepAlive(mount)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Root gets the root directory on mount.
//
// The function returns the following values:
//
//    - file: #GFile. The returned object should be unreffed with
//      g_object_unref() when no longer needed.
//
func (mount *Mount) Root() *File {
	var _arg0 *C.GMount // out
	var _cret *C.GFile  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C.g_mount_get_root(_arg0)
	runtime.KeepAlive(mount)

	var _file *File // out

	_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _file
}

// UUID gets the UUID for the mount. The reference is typically based on the
// file system UUID for the mount in question and should be considered an opaque
// string. Returns NULL if there is no UUID available.
//
// The function returns the following values:
//
//    - utf8 (optional): UUID for mount or NULL if no UUID can be computed. The
//      returned string should be freed with g_free() when no longer needed.
//
func (mount *Mount) UUID() string {
	var _arg0 *C.GMount // out
	var _cret *C.char   // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C.g_mount_get_uuid(_arg0)
	runtime.KeepAlive(mount)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Volume gets the volume for the mount.
//
// The function returns the following values:
//
//    - volume (optional) or NULL if mount is not associated with a volume. The
//      returned object should be unreffed with g_object_unref() when no longer
//      needed.
//
func (mount *Mount) Volume() *Volume {
	var _arg0 *C.GMount  // out
	var _cret *C.GVolume // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C.g_mount_get_volume(_arg0)
	runtime.KeepAlive(mount)

	var _volume *Volume // out

	if _cret != nil {
		_volume = wrapVolume(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _volume
}

// Remount remounts a mount. This is an asynchronous operation, and is finished
// by calling g_mount_remount_finish() with the mount and Results data returned
// in the callback.
//
// Remounting is useful when some setting affecting the operation of the volume
// has been changed, as these may need a remount to take affect. While this is
// semantically equivalent with unmounting and then remounting not all backends
// might need to actually be unmounted.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - flags affecting the operation.
//    - mountOperation (optional) or NULL to avoid user interaction.
//    - callback (optional) or NULL.
//
func (mount *Mount) Remount(ctx context.Context, flags MountMountFlags, mountOperation *MountOperation, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg3 *C.GCancellable       // out
	var _arg1 C.GMountMountFlags    // out
	var _arg2 *C.GMountOperation    // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GMountMountFlags(flags)
	if mountOperation != nil {
		_arg2 = (*C.GMountOperation)(unsafe.Pointer(coreglib.InternObject(mountOperation).Native()))
	}
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_mount_remount(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(mountOperation)
	runtime.KeepAlive(callback)
}

// RemountFinish finishes remounting a mount. If any errors occurred during the
// operation, error will be set to contain the errors and FALSE will be
// returned.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (mount *Mount) RemountFinish(result AsyncResulter) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C.g_mount_remount_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Unmount unmounts a mount. This is an asynchronous operation, and is finished
// by calling g_mount_unmount_finish() with the mount and Result data returned
// in the callback.
//
// Deprecated: Use g_mount_unmount_with_operation() instead.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - flags affecting the operation.
//    - callback (optional) or NULL.
//
func (mount *Mount) Unmount(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GMountUnmountFlags(flags)
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_mount_unmount(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(callback)
}

// UnmountFinish finishes unmounting a mount. If any errors occurred during the
// operation, error will be set to contain the errors and FALSE will be
// returned.
//
// Deprecated: Use g_mount_unmount_with_operation_finish() instead.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (mount *Mount) UnmountFinish(result AsyncResulter) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C.g_mount_unmount_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// canEject checks if mount can be ejected.
//
// The function returns the following values:
//
//    - ok: TRUE if the mount can be ejected.
//
func (mount *Mount) canEject() bool {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.can_eject

	var _arg0 *C.GMount  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C._gotk4_gio2_Mount_virtual_can_eject(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(mount)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// canUnmount checks if mount can be unmounted.
//
// The function returns the following values:
//
//    - ok: TRUE if the mount can be unmounted.
//
func (mount *Mount) canUnmount() bool {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.can_unmount

	var _arg0 *C.GMount  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C._gotk4_gio2_Mount_virtual_can_unmount(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(mount)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (mount *Mount) changed() {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.changed

	var _arg0 *C.GMount // out

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	C._gotk4_gio2_Mount_virtual_changed(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(mount)
}

// Eject ejects a mount. This is an asynchronous operation, and is finished by
// calling g_mount_eject_finish() with the mount and Result data returned in the
// callback.
//
// Deprecated: Use g_mount_eject_with_operation() instead.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - flags affecting the unmount if required for eject.
//    - callback (optional) or NULL.
//
func (mount *Mount) eject(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback) {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.eject

	var _arg0 *C.GMount             // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GMountUnmountFlags(flags)
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C._gotk4_gio2_Mount_virtual_eject(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(callback)
}

// ejectFinish finishes ejecting a mount. If any errors occurred during the
// operation, error will be set to contain the errors and FALSE will be
// returned.
//
// Deprecated: Use g_mount_eject_with_operation_finish() instead.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (mount *Mount) ejectFinish(result AsyncResulter) error {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.eject_finish

	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C._gotk4_gio2_Mount_virtual_eject_finish(unsafe.Pointer(fnarg), _arg0, _arg1, &_cerr)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// defaultLocation gets the default location of mount. The default location of
// the given mount is a path that reflects the main entry point for the user
// (e.g. the home directory, or the root of the volume).
//
// The function returns the following values:
//
//    - file: #GFile. The returned object should be unreffed with
//      g_object_unref() when no longer needed.
//
func (mount *Mount) defaultLocation() *File {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.get_default_location

	var _arg0 *C.GMount // out
	var _cret *C.GFile  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C._gotk4_gio2_Mount_virtual_get_default_location(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(mount)

	var _file *File // out

	_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _file
}

// Drive gets the drive for the mount.
//
// This is a convenience method for getting the #GVolume and then using that
// object to get the #GDrive.
//
// The function returns the following values:
//
//    - drive (optional) or NULL if mount is not associated with a volume or a
//      drive. The returned object should be unreffed with g_object_unref() when
//      no longer needed.
//
func (mount *Mount) drive() *Drive {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.get_drive

	var _arg0 *C.GMount // out
	var _cret *C.GDrive // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C._gotk4_gio2_Mount_virtual_get_drive(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(mount)

	var _drive *Drive // out

	if _cret != nil {
		_drive = wrapDrive(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _drive
}

// Icon gets the icon for mount.
//
// The function returns the following values:
//
//    - icon: #GIcon. The returned object should be unreffed with
//      g_object_unref() when no longer needed.
//
func (mount *Mount) icon() *Icon {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.get_icon

	var _arg0 *C.GMount // out
	var _cret *C.GIcon  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C._gotk4_gio2_Mount_virtual_get_icon(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(mount)

	var _icon *Icon // out

	_icon = wrapIcon(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _icon
}

// Name gets the name of mount.
//
// The function returns the following values:
//
//    - utf8: name for the given mount. The returned string should be freed with
//      g_free() when no longer needed.
//
func (mount *Mount) name() string {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.get_name

	var _arg0 *C.GMount // out
	var _cret *C.char   // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C._gotk4_gio2_Mount_virtual_get_name(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(mount)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Root gets the root directory on mount.
//
// The function returns the following values:
//
//    - file: #GFile. The returned object should be unreffed with
//      g_object_unref() when no longer needed.
//
func (mount *Mount) root() *File {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.get_root

	var _arg0 *C.GMount // out
	var _cret *C.GFile  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C._gotk4_gio2_Mount_virtual_get_root(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(mount)

	var _file *File // out

	_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _file
}

// uuiD gets the UUID for the mount. The reference is typically based on the
// file system UUID for the mount in question and should be considered an opaque
// string. Returns NULL if there is no UUID available.
//
// The function returns the following values:
//
//    - utf8 (optional): UUID for mount or NULL if no UUID can be computed. The
//      returned string should be freed with g_free() when no longer needed.
//
func (mount *Mount) uuiD() string {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.get_uuid

	var _arg0 *C.GMount // out
	var _cret *C.char   // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C._gotk4_gio2_Mount_virtual_get_uuid(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(mount)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Volume gets the volume for the mount.
//
// The function returns the following values:
//
//    - volume (optional) or NULL if mount is not associated with a volume. The
//      returned object should be unreffed with g_object_unref() when no longer
//      needed.
//
func (mount *Mount) volume() *Volume {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.get_volume

	var _arg0 *C.GMount  // out
	var _cret *C.GVolume // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C._gotk4_gio2_Mount_virtual_get_volume(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(mount)

	var _volume *Volume // out

	if _cret != nil {
		_volume = wrapVolume(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _volume
}

func (mount *Mount) preUnmount() {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.pre_unmount

	var _arg0 *C.GMount // out

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	C._gotk4_gio2_Mount_virtual_pre_unmount(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(mount)
}

// Remount remounts a mount. This is an asynchronous operation, and is finished
// by calling g_mount_remount_finish() with the mount and Results data returned
// in the callback.
//
// Remounting is useful when some setting affecting the operation of the volume
// has been changed, as these may need a remount to take affect. While this is
// semantically equivalent with unmounting and then remounting not all backends
// might need to actually be unmounted.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - flags affecting the operation.
//    - mountOperation (optional) or NULL to avoid user interaction.
//    - callback (optional) or NULL.
//
func (mount *Mount) remount(ctx context.Context, flags MountMountFlags, mountOperation *MountOperation, callback AsyncReadyCallback) {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.remount

	var _arg0 *C.GMount             // out
	var _arg3 *C.GCancellable       // out
	var _arg1 C.GMountMountFlags    // out
	var _arg2 *C.GMountOperation    // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GMountMountFlags(flags)
	if mountOperation != nil {
		_arg2 = (*C.GMountOperation)(unsafe.Pointer(coreglib.InternObject(mountOperation).Native()))
	}
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C._gotk4_gio2_Mount_virtual_remount(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(mountOperation)
	runtime.KeepAlive(callback)
}

// remountFinish finishes remounting a mount. If any errors occurred during the
// operation, error will be set to contain the errors and FALSE will be
// returned.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (mount *Mount) remountFinish(result AsyncResulter) error {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.remount_finish

	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C._gotk4_gio2_Mount_virtual_remount_finish(unsafe.Pointer(fnarg), _arg0, _arg1, &_cerr)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Unmount unmounts a mount. This is an asynchronous operation, and is finished
// by calling g_mount_unmount_finish() with the mount and Result data returned
// in the callback.
//
// Deprecated: Use g_mount_unmount_with_operation() instead.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - flags affecting the operation.
//    - callback (optional) or NULL.
//
func (mount *Mount) unmount(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback) {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.unmount

	var _arg0 *C.GMount             // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GMountUnmountFlags(flags)
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C._gotk4_gio2_Mount_virtual_unmount(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(callback)
}

// unmountFinish finishes unmounting a mount. If any errors occurred during the
// operation, error will be set to contain the errors and FALSE will be
// returned.
//
// Deprecated: Use g_mount_unmount_with_operation_finish() instead.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (mount *Mount) unmountFinish(result AsyncResulter) error {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.unmount_finish

	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C._gotk4_gio2_Mount_virtual_unmount_finish(unsafe.Pointer(fnarg), _arg0, _arg1, &_cerr)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

func (mount *Mount) unmounted() {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.unmounted

	var _arg0 *C.GMount // out

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	C._gotk4_gio2_Mount_virtual_unmounted(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(mount)
}

// MountIface: interface for implementing operations for mounts.
//
// An instance of this type is always passed by reference.
type MountIface struct {
	*mountIface
}

// mountIface is the struct that's finalized.
type mountIface struct {
	native *C.GMountIface
}
