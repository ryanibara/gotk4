// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern GDriveStartStopType _gotk4_gio2_DriveIface_get_start_stop_type(GDrive*);
// extern GIcon* _gotk4_gio2_DriveIface_get_icon(GDrive*);
// extern GIcon* _gotk4_gio2_DriveIface_get_symbolic_icon(GDrive*);
// extern GList* _gotk4_gio2_DriveIface_get_volumes(GDrive*);
// extern char* _gotk4_gio2_DriveIface_get_identifier(GDrive*, char*);
// extern char* _gotk4_gio2_DriveIface_get_name(GDrive*);
// extern char** _gotk4_gio2_DriveIface_enumerate_identifiers(GDrive*);
// extern gboolean _gotk4_gio2_DriveIface_can_eject(GDrive*);
// extern gboolean _gotk4_gio2_DriveIface_can_poll_for_media(GDrive*);
// extern gboolean _gotk4_gio2_DriveIface_can_start(GDrive*);
// extern gboolean _gotk4_gio2_DriveIface_can_start_degraded(GDrive*);
// extern gboolean _gotk4_gio2_DriveIface_can_stop(GDrive*);
// extern gboolean _gotk4_gio2_DriveIface_eject_finish(GDrive*, GAsyncResult*, GError**);
// extern gboolean _gotk4_gio2_DriveIface_eject_with_operation_finish(GDrive*, GAsyncResult*, GError**);
// extern gboolean _gotk4_gio2_DriveIface_has_media(GDrive*);
// extern gboolean _gotk4_gio2_DriveIface_has_volumes(GDrive*);
// extern gboolean _gotk4_gio2_DriveIface_is_media_check_automatic(GDrive*);
// extern gboolean _gotk4_gio2_DriveIface_is_media_removable(GDrive*);
// extern gboolean _gotk4_gio2_DriveIface_is_removable(GDrive*);
// extern gboolean _gotk4_gio2_DriveIface_poll_for_media_finish(GDrive*, GAsyncResult*, GError**);
// extern gboolean _gotk4_gio2_DriveIface_start_finish(GDrive*, GAsyncResult*, GError**);
// extern gboolean _gotk4_gio2_DriveIface_stop_finish(GDrive*, GAsyncResult*, GError**);
// extern gchar* _gotk4_gio2_DriveIface_get_sort_key(GDrive*);
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
// extern void _gotk4_gio2_DriveIface_changed(GDrive*);
// extern void _gotk4_gio2_DriveIface_disconnected(GDrive*);
// extern void _gotk4_gio2_DriveIface_eject_button(GDrive*);
// extern void _gotk4_gio2_DriveIface_stop_button(GDrive*);
// extern void _gotk4_gio2_Drive_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gio2_Drive_ConnectDisconnected(gpointer, guintptr);
// extern void _gotk4_gio2_Drive_ConnectEjectButton(gpointer, guintptr);
// extern void _gotk4_gio2_Drive_ConnectStopButton(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeDrive = coreglib.Type(C.g_drive_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDrive, F: marshalDrive},
	})
}

// DRIVE_IDENTIFIER_KIND_UNIX_DEVICE: string used to obtain a Unix device path
// with g_drive_get_identifier().
const DRIVE_IDENTIFIER_KIND_UNIX_DEVICE = "unix-device"

// Drive - this represent a piece of hardware connected to the machine. It's
// generally only created for removable hardware or hardware with removable
// media.
//
// #GDrive is a container class for #GVolume objects that stem from the same
// piece of media. As such, #GDrive abstracts a drive with (or without)
// removable media and provides operations for querying whether media is
// available, determining whether media change is automatically detected and
// ejecting the media.
//
// If the #GDrive reports that media isn't automatically detected, one can poll
// for media; typically one should not do this periodically as a poll for media
// operation is potentially expensive and may spin up the drive creating noise.
//
// #GDrive supports starting and stopping drives with authentication support for
// the former. This can be used to support a diverse set of use cases including
// connecting/disconnecting iSCSI devices, powering down external disk
// enclosures and starting/stopping multi-disk devices such as RAID devices.
// Note that the actual semantics and side-effects of starting/stopping a
// #GDrive may vary according to implementation. To choose the correct verbs in
// e.g. a file manager, use g_drive_get_start_stop_type().
//
// For porting from GnomeVFS note that there is no equivalent of #GDrive in that
// API.
//
// Drive wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Drive struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Drive)(nil)
)

// Driver describes Drive's interface methods.
type Driver interface {
	coreglib.Objector

	// CanEject checks if a drive can be ejected.
	CanEject() bool
	// CanPollForMedia checks if a drive can be polled for media changes.
	CanPollForMedia() bool
	// CanStart checks if a drive can be started.
	CanStart() bool
	// CanStartDegraded checks if a drive can be started degraded.
	CanStartDegraded() bool
	// CanStop checks if a drive can be stopped.
	CanStop() bool
	// Eject: asynchronously ejects a drive.
	Eject(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback)
	// EjectFinish finishes ejecting a drive.
	EjectFinish(result AsyncResulter) error
	// EjectWithOperation ejects a drive.
	EjectWithOperation(ctx context.Context, flags MountUnmountFlags, mountOperation *MountOperation, callback AsyncReadyCallback)
	// EjectWithOperationFinish finishes ejecting a drive.
	EjectWithOperationFinish(result AsyncResulter) error
	// EnumerateIdentifiers gets the kinds of identifiers that drive has.
	EnumerateIdentifiers() []string
	// Icon gets the icon for drive.
	Icon() *Icon
	// Identifier gets the identifier of the given kind for drive.
	Identifier(kind string) string
	// Name gets the name of drive.
	Name() string
	// SortKey gets the sort key for drive, if any.
	SortKey() string
	// StartStopType gets a hint about how a drive can be started/stopped.
	StartStopType() DriveStartStopType
	// SymbolicIcon gets the icon for drive.
	SymbolicIcon() *Icon
	// Volumes: get a list of mountable volumes for drive.
	Volumes() []*Volume
	// HasMedia checks if the drive has media.
	HasMedia() bool
	// HasVolumes: check if drive has any mountable volumes.
	HasVolumes() bool
	// IsMediaCheckAutomatic checks if drive is capable of automatically
	// detecting media changes.
	IsMediaCheckAutomatic() bool
	// IsMediaRemovable checks if the drive supports removable media.
	IsMediaRemovable() bool
	// IsRemovable checks if the #GDrive and/or its media is considered
	// removable by the user.
	IsRemovable() bool
	// PollForMedia: asynchronously polls drive to see if media has been
	// inserted or removed.
	PollForMedia(ctx context.Context, callback AsyncReadyCallback)
	// PollForMediaFinish finishes an operation started with
	// g_drive_poll_for_media() on a drive.
	PollForMediaFinish(result AsyncResulter) error
	// Start: asynchronously starts a drive.
	Start(ctx context.Context, flags DriveStartFlags, mountOperation *MountOperation, callback AsyncReadyCallback)
	// StartFinish finishes starting a drive.
	StartFinish(result AsyncResulter) error
	// Stop: asynchronously stops a drive.
	Stop(ctx context.Context, flags MountUnmountFlags, mountOperation *MountOperation, callback AsyncReadyCallback)
	// StopFinish finishes stopping a drive.
	StopFinish(result AsyncResulter) error

	// Changed is emitted when the drive's state has changed.
	ConnectChanged(func()) coreglib.SignalHandle
	// Disconnected: this signal is emitted when the #GDrive have been
	// disconnected.
	ConnectDisconnected(func()) coreglib.SignalHandle
	// Eject-button is emitted when the physical eject button (if any) of a
	// drive has been pressed.
	ConnectEjectButton(func()) coreglib.SignalHandle
	// Stop-button is emitted when the physical stop button (if any) of a drive
	// has been pressed.
	ConnectStopButton(func()) coreglib.SignalHandle
}

var _ Driver = (*Drive)(nil)

func wrapDrive(obj *coreglib.Object) *Drive {
	return &Drive{
		Object: obj,
	}
}

func marshalDrive(p uintptr) (interface{}, error) {
	return wrapDrive(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_Drive_ConnectChanged
func _gotk4_gio2_Drive_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectChanged is emitted when the drive's state has changed.
func (drive *Drive) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(drive, "changed", false, unsafe.Pointer(C._gotk4_gio2_Drive_ConnectChanged), f)
}

//export _gotk4_gio2_Drive_ConnectDisconnected
func _gotk4_gio2_Drive_ConnectDisconnected(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectDisconnected: this signal is emitted when the #GDrive have been
// disconnected. If the recipient is holding references to the object they
// should release them so the object can be finalized.
func (drive *Drive) ConnectDisconnected(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(drive, "disconnected", false, unsafe.Pointer(C._gotk4_gio2_Drive_ConnectDisconnected), f)
}

//export _gotk4_gio2_Drive_ConnectEjectButton
func _gotk4_gio2_Drive_ConnectEjectButton(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectEjectButton is emitted when the physical eject button (if any) of a
// drive has been pressed.
func (drive *Drive) ConnectEjectButton(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(drive, "eject-button", false, unsafe.Pointer(C._gotk4_gio2_Drive_ConnectEjectButton), f)
}

//export _gotk4_gio2_Drive_ConnectStopButton
func _gotk4_gio2_Drive_ConnectStopButton(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectStopButton is emitted when the physical stop button (if any) of a
// drive has been pressed.
func (drive *Drive) ConnectStopButton(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(drive, "stop-button", false, unsafe.Pointer(C._gotk4_gio2_Drive_ConnectStopButton), f)
}

// CanEject checks if a drive can be ejected.
//
// The function returns the following values:
//
//    - ok: TRUE if the drive can be ejected, FALSE otherwise.
//
func (drive *Drive) CanEject() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_can_eject(_arg0)
	runtime.KeepAlive(drive)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanPollForMedia checks if a drive can be polled for media changes.
//
// The function returns the following values:
//
//    - ok: TRUE if the drive can be polled for media changes, FALSE otherwise.
//
func (drive *Drive) CanPollForMedia() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_can_poll_for_media(_arg0)
	runtime.KeepAlive(drive)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanStart checks if a drive can be started.
//
// The function returns the following values:
//
//    - ok: TRUE if the drive can be started, FALSE otherwise.
//
func (drive *Drive) CanStart() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_can_start(_arg0)
	runtime.KeepAlive(drive)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanStartDegraded checks if a drive can be started degraded.
//
// The function returns the following values:
//
//    - ok: TRUE if the drive can be started degraded, FALSE otherwise.
//
func (drive *Drive) CanStartDegraded() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_can_start_degraded(_arg0)
	runtime.KeepAlive(drive)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanStop checks if a drive can be stopped.
//
// The function returns the following values:
//
//    - ok: TRUE if the drive can be stopped, FALSE otherwise.
//
func (drive *Drive) CanStop() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_can_stop(_arg0)
	runtime.KeepAlive(drive)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Eject: asynchronously ejects a drive.
//
// When the operation is finished, callback will be called. You can then call
// g_drive_eject_finish() to obtain the result of the operation.
//
// Deprecated: Use g_drive_eject_with_operation() instead.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - flags affecting the unmount if required for eject.
//    - callback (optional) or NULL.
//
func (drive *Drive) Eject(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback) {
	var _arg0 *C.GDrive             // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))
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

	C.g_drive_eject(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(drive)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(callback)
}

// EjectFinish finishes ejecting a drive.
//
// Deprecated: Use g_drive_eject_with_operation_finish() instead.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (drive *Drive) EjectFinish(result AsyncResulter) error {
	var _arg0 *C.GDrive       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C.g_drive_eject_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(drive)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// EjectWithOperation ejects a drive. This is an asynchronous operation, and is
// finished by calling g_drive_eject_with_operation_finish() with the drive and
// Result data returned in the callback.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - flags affecting the unmount if required for eject.
//    - mountOperation (optional) or NULL to avoid user interaction.
//    - callback (optional) or NULL.
//
func (drive *Drive) EjectWithOperation(ctx context.Context, flags MountUnmountFlags, mountOperation *MountOperation, callback AsyncReadyCallback) {
	var _arg0 *C.GDrive             // out
	var _arg3 *C.GCancellable       // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg2 *C.GMountOperation    // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GMountUnmountFlags(flags)
	if mountOperation != nil {
		_arg2 = (*C.GMountOperation)(unsafe.Pointer(coreglib.InternObject(mountOperation).Native()))
	}
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_drive_eject_with_operation(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(drive)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(mountOperation)
	runtime.KeepAlive(callback)
}

// EjectWithOperationFinish finishes ejecting a drive. If any errors occurred
// during the operation, error will be set to contain the errors and FALSE will
// be returned.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (drive *Drive) EjectWithOperationFinish(result AsyncResulter) error {
	var _arg0 *C.GDrive       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C.g_drive_eject_with_operation_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(drive)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// EnumerateIdentifiers gets the kinds of identifiers that drive has. Use
// g_drive_get_identifier() to obtain the identifiers themselves.
//
// The function returns the following values:
//
//    - utf8s: NULL-terminated array of strings containing kinds of identifiers.
//      Use g_strfreev() to free.
//
func (drive *Drive) EnumerateIdentifiers() []string {
	var _arg0 *C.GDrive // out
	var _cret **C.char  // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_enumerate_identifiers(_arg0)
	runtime.KeepAlive(drive)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// Icon gets the icon for drive.
//
// The function returns the following values:
//
//    - icon for the drive. Free the returned object with g_object_unref().
//
func (drive *Drive) Icon() *Icon {
	var _arg0 *C.GDrive // out
	var _cret *C.GIcon  // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_get_icon(_arg0)
	runtime.KeepAlive(drive)

	var _icon *Icon // out

	_icon = wrapIcon(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _icon
}

// Identifier gets the identifier of the given kind for drive. The only
// identifier currently available is DRIVE_IDENTIFIER_KIND_UNIX_DEVICE.
//
// The function takes the following parameters:
//
//    - kind of identifier to return.
//
// The function returns the following values:
//
//    - utf8 (optional): newly allocated string containing the requested
//      identifier, or NULL if the #GDrive doesn't have this kind of identifier.
//
func (drive *Drive) Identifier(kind string) string {
	var _arg0 *C.GDrive // out
	var _arg1 *C.char   // out
	var _cret *C.char   // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(kind)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_drive_get_identifier(_arg0, _arg1)
	runtime.KeepAlive(drive)
	runtime.KeepAlive(kind)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Name gets the name of drive.
//
// The function returns the following values:
//
//    - utf8: string containing drive's name. The returned string should be freed
//      when no longer needed.
//
func (drive *Drive) Name() string {
	var _arg0 *C.GDrive // out
	var _cret *C.char   // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_get_name(_arg0)
	runtime.KeepAlive(drive)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SortKey gets the sort key for drive, if any.
//
// The function returns the following values:
//
//    - utf8 (optional): sorting key for drive or NULL if no such key is
//      available.
//
func (drive *Drive) SortKey() string {
	var _arg0 *C.GDrive // out
	var _cret *C.gchar  // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_get_sort_key(_arg0)
	runtime.KeepAlive(drive)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// StartStopType gets a hint about how a drive can be started/stopped.
//
// The function returns the following values:
//
//    - driveStartStopType: value from the StartStopType enumeration.
//
func (drive *Drive) StartStopType() DriveStartStopType {
	var _arg0 *C.GDrive             // out
	var _cret C.GDriveStartStopType // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_get_start_stop_type(_arg0)
	runtime.KeepAlive(drive)

	var _driveStartStopType DriveStartStopType // out

	_driveStartStopType = DriveStartStopType(_cret)

	return _driveStartStopType
}

// SymbolicIcon gets the icon for drive.
//
// The function returns the following values:
//
//    - icon: symbolic #GIcon for the drive. Free the returned object with
//      g_object_unref().
//
func (drive *Drive) SymbolicIcon() *Icon {
	var _arg0 *C.GDrive // out
	var _cret *C.GIcon  // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_get_symbolic_icon(_arg0)
	runtime.KeepAlive(drive)

	var _icon *Icon // out

	_icon = wrapIcon(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _icon
}

// Volumes: get a list of mountable volumes for drive.
//
// The returned list should be freed with g_list_free(), after its elements have
// been unreffed with g_object_unref().
//
// The function returns the following values:
//
//    - list containing any #GVolume objects on the given drive.
//
func (drive *Drive) Volumes() []*Volume {
	var _arg0 *C.GDrive // out
	var _cret *C.GList  // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_get_volumes(_arg0)
	runtime.KeepAlive(drive)

	var _list []*Volume // out

	_list = make([]*Volume, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GVolume)(v)
		var dst *Volume // out
		dst = wrapVolume(coreglib.AssumeOwnership(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// HasMedia checks if the drive has media. Note that the OS may not be polling
// the drive for media changes; see g_drive_is_media_check_automatic() for more
// details.
//
// The function returns the following values:
//
//    - ok: TRUE if drive has media, FALSE otherwise.
//
func (drive *Drive) HasMedia() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_has_media(_arg0)
	runtime.KeepAlive(drive)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasVolumes: check if drive has any mountable volumes.
//
// The function returns the following values:
//
//    - ok: TRUE if the drive contains volumes, FALSE otherwise.
//
func (drive *Drive) HasVolumes() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_has_volumes(_arg0)
	runtime.KeepAlive(drive)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMediaCheckAutomatic checks if drive is capable of automatically detecting
// media changes.
//
// The function returns the following values:
//
//    - ok: TRUE if the drive is capable of automatically detecting media
//      changes, FALSE otherwise.
//
func (drive *Drive) IsMediaCheckAutomatic() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_is_media_check_automatic(_arg0)
	runtime.KeepAlive(drive)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMediaRemovable checks if the drive supports removable media.
//
// The function returns the following values:
//
//    - ok: TRUE if drive supports removable media, FALSE otherwise.
//
func (drive *Drive) IsMediaRemovable() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_is_media_removable(_arg0)
	runtime.KeepAlive(drive)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsRemovable checks if the #GDrive and/or its media is considered removable by
// the user. See g_drive_is_media_removable().
//
// The function returns the following values:
//
//    - ok: TRUE if drive and/or its media is considered removable, FALSE
//      otherwise.
//
func (drive *Drive) IsRemovable() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))

	_cret = C.g_drive_is_removable(_arg0)
	runtime.KeepAlive(drive)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PollForMedia: asynchronously polls drive to see if media has been inserted or
// removed.
//
// When the operation is finished, callback will be called. You can then call
// g_drive_poll_for_media_finish() to obtain the result of the operation.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - callback (optional) or NULL.
//
func (drive *Drive) PollForMedia(ctx context.Context, callback AsyncReadyCallback) {
	var _arg0 *C.GDrive             // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_drive_poll_for_media(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(drive)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// PollForMediaFinish finishes an operation started with
// g_drive_poll_for_media() on a drive.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (drive *Drive) PollForMediaFinish(result AsyncResulter) error {
	var _arg0 *C.GDrive       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C.g_drive_poll_for_media_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(drive)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Start: asynchronously starts a drive.
//
// When the operation is finished, callback will be called. You can then call
// g_drive_start_finish() to obtain the result of the operation.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - flags affecting the start operation.
//    - mountOperation (optional) or NULL to avoid user interaction.
//    - callback (optional) or NULL.
//
func (drive *Drive) Start(ctx context.Context, flags DriveStartFlags, mountOperation *MountOperation, callback AsyncReadyCallback) {
	var _arg0 *C.GDrive             // out
	var _arg3 *C.GCancellable       // out
	var _arg1 C.GDriveStartFlags    // out
	var _arg2 *C.GMountOperation    // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GDriveStartFlags(flags)
	if mountOperation != nil {
		_arg2 = (*C.GMountOperation)(unsafe.Pointer(coreglib.InternObject(mountOperation).Native()))
	}
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_drive_start(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(drive)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(mountOperation)
	runtime.KeepAlive(callback)
}

// StartFinish finishes starting a drive.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (drive *Drive) StartFinish(result AsyncResulter) error {
	var _arg0 *C.GDrive       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C.g_drive_start_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(drive)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Stop: asynchronously stops a drive.
//
// When the operation is finished, callback will be called. You can then call
// g_drive_stop_finish() to obtain the result of the operation.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - flags affecting the unmount if required for stopping.
//    - mountOperation (optional) or NULL to avoid user interaction.
//    - callback (optional) or NULL.
//
func (drive *Drive) Stop(ctx context.Context, flags MountUnmountFlags, mountOperation *MountOperation, callback AsyncReadyCallback) {
	var _arg0 *C.GDrive             // out
	var _arg3 *C.GCancellable       // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg2 *C.GMountOperation    // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GMountUnmountFlags(flags)
	if mountOperation != nil {
		_arg2 = (*C.GMountOperation)(unsafe.Pointer(coreglib.InternObject(mountOperation).Native()))
	}
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_drive_stop(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(drive)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(mountOperation)
	runtime.KeepAlive(callback)
}

// StopFinish finishes stopping a drive.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (drive *Drive) StopFinish(result AsyncResulter) error {
	var _arg0 *C.GDrive       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(coreglib.InternObject(drive).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C.g_drive_stop_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(drive)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
