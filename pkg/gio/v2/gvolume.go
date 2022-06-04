// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern GDrive* _gotk4_gio2_VolumeIface_get_drive(void*);
// extern GFile* _gotk4_gio2_VolumeIface_get_activation_root(void*);
// extern GIcon* _gotk4_gio2_VolumeIface_get_icon(void*);
// extern GIcon* _gotk4_gio2_VolumeIface_get_symbolic_icon(void*);
// extern GMount* _gotk4_gio2_VolumeIface_get_mount(void*);
// extern char* _gotk4_gio2_VolumeIface_get_identifier(void*, void*);
// extern char* _gotk4_gio2_VolumeIface_get_name(void*);
// extern char* _gotk4_gio2_VolumeIface_get_uuid(void*);
// extern char** _gotk4_gio2_VolumeIface_enumerate_identifiers(void*);
// extern gboolean _gotk4_gio2_VolumeIface_can_eject(void*);
// extern gboolean _gotk4_gio2_VolumeIface_can_mount(void*);
// extern gboolean _gotk4_gio2_VolumeIface_eject_finish(void*, void*, GError**);
// extern gboolean _gotk4_gio2_VolumeIface_eject_with_operation_finish(void*, void*, GError**);
// extern gboolean _gotk4_gio2_VolumeIface_mount_finish(void*, void*, GError**);
// extern gboolean _gotk4_gio2_VolumeIface_should_automount(void*);
// extern gchar* _gotk4_gio2_VolumeIface_get_sort_key(void*);
// extern void _gotk4_gio2_VolumeIface_changed(void*);
// extern void _gotk4_gio2_VolumeIface_removed(void*);
// extern void _gotk4_gio2_Volume_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gio2_Volume_ConnectRemoved(gpointer, guintptr);
import "C"

// glib.Type values for gvolume.go.
var GTypeVolume = coreglib.Type(C.g_volume_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeVolume, F: marshalVolume},
	})
}

// VOLUME_IDENTIFIER_KIND_CLASS: string used to obtain the volume class with
// g_volume_get_identifier().
//
// Known volume classes include device, network, and loop. Other classes may be
// added in the future.
//
// This is intended to be used by applications to classify #GVolume instances
// into different sections - for example a file manager or file chooser can use
// this information to show network volumes under a "Network" heading and device
// volumes under a "Devices" heading.
const VOLUME_IDENTIFIER_KIND_CLASS = "class"

// VOLUME_IDENTIFIER_KIND_HAL_UDI: string used to obtain a Hal UDI with
// g_volume_get_identifier().
//
// Deprecated: Do not use, HAL is deprecated.
const VOLUME_IDENTIFIER_KIND_HAL_UDI = "hal-udi"

// VOLUME_IDENTIFIER_KIND_LABEL: string used to obtain a filesystem label with
// g_volume_get_identifier().
const VOLUME_IDENTIFIER_KIND_LABEL = "label"

// VOLUME_IDENTIFIER_KIND_NFS_MOUNT: string used to obtain a NFS mount with
// g_volume_get_identifier().
const VOLUME_IDENTIFIER_KIND_NFS_MOUNT = "nfs-mount"

// VOLUME_IDENTIFIER_KIND_UNIX_DEVICE: string used to obtain a Unix device path
// with g_volume_get_identifier().
const VOLUME_IDENTIFIER_KIND_UNIX_DEVICE = "unix-device"

// VOLUME_IDENTIFIER_KIND_UUID: string used to obtain a UUID with
// g_volume_get_identifier().
const VOLUME_IDENTIFIER_KIND_UUID = "uuid"

// VolumeOverrider contains methods that are overridable.
type VolumeOverrider interface {
	// CanEject checks if a volume can be ejected.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if the volume can be ejected. FALSE otherwise.
	//
	CanEject() bool
	// CanMount checks if a volume can be mounted.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if the volume can be mounted. FALSE otherwise.
	//
	CanMount() bool
	Changed()
	// EjectFinish finishes ejecting a volume. If any errors occurred during the
	// operation, error will be set to contain the errors and FALSE will be
	// returned.
	//
	// Deprecated: Use g_volume_eject_with_operation_finish() instead.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	EjectFinish(result AsyncResulter) error
	// EjectWithOperationFinish finishes ejecting a volume. If any errors
	// occurred during the operation, error will be set to contain the errors
	// and FALSE will be returned.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	EjectWithOperationFinish(result AsyncResulter) error
	// EnumerateIdentifiers gets the kinds of [identifiers][volume-identifier]
	// that volume has. Use g_volume_get_identifier() to obtain the identifiers
	// themselves.
	//
	// The function returns the following values:
	//
	//    - utf8s: NULL-terminated array of strings containing kinds of
	//      identifiers. Use g_strfreev() to free.
	//
	EnumerateIdentifiers() []string
	// ActivationRoot gets the activation root for a #GVolume if it is known
	// ahead of mount time. Returns NULL otherwise. If not NULL and if volume is
	// mounted, then the result of g_mount_get_root() on the #GMount object
	// obtained from g_volume_get_mount() will always either be equal or a
	// prefix of what this function returns. In other words, in code
	//
	//    (g_file_has_prefix (volume_activation_root, mount_root) ||
	//     g_file_equal (volume_activation_root, mount_root))
	//
	// will always be TRUE.
	//
	// Activation roots are typically used in Monitor implementations to find
	// the underlying mount to shadow, see g_mount_is_shadowed() for more
	// details.
	//
	// The function returns the following values:
	//
	//    - file (optional): activation root of volume or NULL. Use
	//      g_object_unref() to free.
	//
	ActivationRoot() *File
	// Drive gets the drive for the volume.
	//
	// The function returns the following values:
	//
	//    - drive (optional) or NULL if volume is not associated with a drive.
	//      The returned object should be unreffed with g_object_unref() when no
	//      longer needed.
	//
	Drive() *Drive
	// Icon gets the icon for volume.
	//
	// The function returns the following values:
	//
	//    - icon: #GIcon. The returned object should be unreffed with
	//      g_object_unref() when no longer needed.
	//
	Icon() *Icon
	// Identifier gets the identifier of the given kind for volume. See the
	// [introduction][volume-identifier] for more information about volume
	// identifiers.
	//
	// The function takes the following parameters:
	//
	//    - kind of identifier to return.
	//
	// The function returns the following values:
	//
	//    - utf8 (optional): newly allocated string containing the requested
	//      identifier, or NULL if the #GVolume doesn't have this kind of
	//      identifier.
	//
	Identifier(kind string) string
	// Mount gets the mount for the volume.
	//
	// The function returns the following values:
	//
	//    - mount (optional) or NULL if volume isn't mounted. The returned object
	//      should be unreffed with g_object_unref() when no longer needed.
	//
	Mount() *Mount
	// Name gets the name of volume.
	//
	// The function returns the following values:
	//
	//    - utf8: name for the given volume. The returned string should be freed
	//      with g_free() when no longer needed.
	//
	Name() string
	// SortKey gets the sort key for volume, if any.
	//
	// The function returns the following values:
	//
	//    - utf8 (optional): sorting key for volume or NULL if no such key is
	//      available.
	//
	SortKey() string
	// SymbolicIcon gets the symbolic icon for volume.
	//
	// The function returns the following values:
	//
	//    - icon: #GIcon. The returned object should be unreffed with
	//      g_object_unref() when no longer needed.
	//
	SymbolicIcon() *Icon
	// UUID gets the UUID for the volume. The reference is typically based on
	// the file system UUID for the volume in question and should be considered
	// an opaque string. Returns NULL if there is no UUID available.
	//
	// The function returns the following values:
	//
	//    - utf8 (optional): UUID for volume or NULL if no UUID can be computed.
	//      The returned string should be freed with g_free() when no longer
	//      needed.
	//
	UUID() string
	// MountFinish finishes mounting a volume. If any errors occurred during the
	// operation, error will be set to contain the errors and FALSE will be
	// returned.
	//
	// If the mount operation succeeded, g_volume_get_mount() on volume is
	// guaranteed to return the mount right after calling this function; there's
	// no need to listen for the 'mount-added' signal on Monitor.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	MountFinish(result AsyncResulter) error
	Removed()
	// ShouldAutomount returns whether the volume should be automatically
	// mounted.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if the volume should be automatically mounted.
	//
	ShouldAutomount() bool
}

// Volume interface represents user-visible objects that can be mounted. Note,
// when porting from GnomeVFS, #GVolume is the moral equivalent of VFSDrive.
//
// Mounting a #GVolume instance is an asynchronous operation. For more
// information about asynchronous operations, see Result and #GTask. To mount a
// #GVolume, first call g_volume_mount() with (at least) the #GVolume instance,
// optionally a Operation object and a ReadyCallback.
//
// Typically, one will only want to pass NULL for the Operation if automounting
// all volumes when a desktop session starts since it's not desirable to put up
// a lot of dialogs asking for credentials.
//
// The callback will be fired when the operation has resolved (either with
// success or failure), and a Result instance will be passed to the callback.
// That callback should then call g_volume_mount_finish() with the #GVolume
// instance and the Result data to see if the operation was completed
// successfully. If an error is present when g_volume_mount_finish() is called,
// then it will be filled with any error information.
//
//
// Volume Identifiers
//
// It is sometimes necessary to directly access the underlying operating system
// object behind a volume (e.g. for passing a volume to an application via the
// commandline). For this purpose, GIO allows to obtain an 'identifier' for the
// volume. There can be different kinds of identifiers, such as Hal UDIs,
// filesystem labels, traditional Unix devices (e.g. /dev/sda2), UUIDs. GIO uses
// predefined strings as names for the different kinds of identifiers:
// VOLUME_IDENTIFIER_KIND_UUID, VOLUME_IDENTIFIER_KIND_LABEL, etc. Use
// g_volume_get_identifier() to obtain an identifier for a volume.
//
//    Note that VOLUME_IDENTIFIER_KIND_HAL_UDI will only be available when the gvfs hal volume monitor is in use. Other volume monitors will generally be able to provide the VOLUME_IDENTIFIER_KIND_UNIX_DEVICE identifier, which can be used to obtain a hal device by means of libhal_manager_find_device_string_match().
//
// Volume wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Volume struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Volume)(nil)
)

// Volumer describes Volume's interface methods.
type Volumer interface {
	coreglib.Objector

	// CanEject checks if a volume can be ejected.
	CanEject() bool
	// CanMount checks if a volume can be mounted.
	CanMount() bool
	// EjectFinish finishes ejecting a volume.
	EjectFinish(result AsyncResulter) error
	// EjectWithOperationFinish finishes ejecting a volume.
	EjectWithOperationFinish(result AsyncResulter) error
	// EnumerateIdentifiers gets the kinds of [identifiers][volume-identifier]
	// that volume has.
	EnumerateIdentifiers() []string
	// ActivationRoot gets the activation root for a #GVolume if it is known
	// ahead of mount time.
	ActivationRoot() *File
	// Drive gets the drive for the volume.
	Drive() *Drive
	// Icon gets the icon for volume.
	Icon() *Icon
	// Identifier gets the identifier of the given kind for volume.
	Identifier(kind string) string
	// Mount gets the mount for the volume.
	Mount() *Mount
	// Name gets the name of volume.
	Name() string
	// SortKey gets the sort key for volume, if any.
	SortKey() string
	// SymbolicIcon gets the symbolic icon for volume.
	SymbolicIcon() *Icon
	// UUID gets the UUID for the volume.
	UUID() string
	// MountFinish finishes mounting a volume.
	MountFinish(result AsyncResulter) error
	// ShouldAutomount returns whether the volume should be automatically
	// mounted.
	ShouldAutomount() bool

	// Changed is emitted when the volume has been changed.
	ConnectChanged(func()) coreglib.SignalHandle
	// Removed: this signal is emitted when the #GVolume have been removed.
	ConnectRemoved(func()) coreglib.SignalHandle
}

var _ Volumer = (*Volume)(nil)

func ifaceInitVolumer(gifacePtr, data C.gpointer) {
	iface := (*C.GVolumeIface)(unsafe.Pointer(gifacePtr))
	iface.can_eject = (*[0]byte)(C._gotk4_gio2_VolumeIface_can_eject)
	iface.can_mount = (*[0]byte)(C._gotk4_gio2_VolumeIface_can_mount)
	iface.changed = (*[0]byte)(C._gotk4_gio2_VolumeIface_changed)
	iface.eject_finish = (*[0]byte)(C._gotk4_gio2_VolumeIface_eject_finish)
	iface.eject_with_operation_finish = (*[0]byte)(C._gotk4_gio2_VolumeIface_eject_with_operation_finish)
	iface.enumerate_identifiers = (*[0]byte)(C._gotk4_gio2_VolumeIface_enumerate_identifiers)
	iface.get_activation_root = (*[0]byte)(C._gotk4_gio2_VolumeIface_get_activation_root)
	iface.get_drive = (*[0]byte)(C._gotk4_gio2_VolumeIface_get_drive)
	iface.get_icon = (*[0]byte)(C._gotk4_gio2_VolumeIface_get_icon)
	iface.get_identifier = (*[0]byte)(C._gotk4_gio2_VolumeIface_get_identifier)
	iface.get_mount = (*[0]byte)(C._gotk4_gio2_VolumeIface_get_mount)
	iface.get_name = (*[0]byte)(C._gotk4_gio2_VolumeIface_get_name)
	iface.get_sort_key = (*[0]byte)(C._gotk4_gio2_VolumeIface_get_sort_key)
	iface.get_symbolic_icon = (*[0]byte)(C._gotk4_gio2_VolumeIface_get_symbolic_icon)
	iface.get_uuid = (*[0]byte)(C._gotk4_gio2_VolumeIface_get_uuid)
	iface.mount_finish = (*[0]byte)(C._gotk4_gio2_VolumeIface_mount_finish)
	iface.removed = (*[0]byte)(C._gotk4_gio2_VolumeIface_removed)
	iface.should_automount = (*[0]byte)(C._gotk4_gio2_VolumeIface_should_automount)
}

//export _gotk4_gio2_VolumeIface_can_eject
func _gotk4_gio2_VolumeIface_can_eject(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	ok := iface.CanEject()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_VolumeIface_can_mount
func _gotk4_gio2_VolumeIface_can_mount(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	ok := iface.CanMount()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_VolumeIface_changed
func _gotk4_gio2_VolumeIface_changed(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	iface.Changed()
}

//export _gotk4_gio2_VolumeIface_eject_finish
func _gotk4_gio2_VolumeIface_eject_finish(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	_goerr := iface.EjectFinish(_result)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_VolumeIface_eject_with_operation_finish
func _gotk4_gio2_VolumeIface_eject_with_operation_finish(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	_goerr := iface.EjectWithOperationFinish(_result)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_VolumeIface_enumerate_identifiers
func _gotk4_gio2_VolumeIface_enumerate_identifiers(arg0 *C.void) (cret **C.char) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	utf8s := iface.EnumerateIdentifiers()

	{
		cret = (**C.void)(C.calloc(C.size_t((len(utf8s) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		{
			out := unsafe.Slice(cret, len(utf8s)+1)
			var zero *C.void
			out[len(utf8s)] = zero
			for i := range utf8s {
				out[i] = (*C.void)(unsafe.Pointer(C.CString(utf8s[i])))
			}
		}
	}

	return cret
}

//export _gotk4_gio2_VolumeIface_get_activation_root
func _gotk4_gio2_VolumeIface_get_activation_root(arg0 *C.void) (cret *C.GFile) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	file := iface.ActivationRoot()

	if file != nil {
		cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(file).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(file).Native()))
	}

	return cret
}

//export _gotk4_gio2_VolumeIface_get_drive
func _gotk4_gio2_VolumeIface_get_drive(arg0 *C.void) (cret *C.GDrive) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	drive := iface.Drive()

	if drive != nil {
		cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(drive).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(drive).Native()))
	}

	return cret
}

//export _gotk4_gio2_VolumeIface_get_icon
func _gotk4_gio2_VolumeIface_get_icon(arg0 *C.void) (cret *C.GIcon) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	icon := iface.Icon()

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(icon).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(icon).Native()))

	return cret
}

//export _gotk4_gio2_VolumeIface_get_identifier
func _gotk4_gio2_VolumeIface_get_identifier(arg0 *C.void, arg1 *C.void) (cret *C.char) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	var _kind string // out

	_kind = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	utf8 := iface.Identifier(_kind)

	if utf8 != "" {
		cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
	}

	return cret
}

//export _gotk4_gio2_VolumeIface_get_mount
func _gotk4_gio2_VolumeIface_get_mount(arg0 *C.void) (cret *C.GMount) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	mount := iface.Mount()

	if mount != nil {
		cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(mount).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(mount).Native()))
	}

	return cret
}

//export _gotk4_gio2_VolumeIface_get_name
func _gotk4_gio2_VolumeIface_get_name(arg0 *C.void) (cret *C.char) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	utf8 := iface.Name()

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

//export _gotk4_gio2_VolumeIface_get_sort_key
func _gotk4_gio2_VolumeIface_get_sort_key(arg0 *C.void) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	utf8 := iface.SortKey()

	if utf8 != "" {
		cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
		defer C.free(unsafe.Pointer(cret))
	}

	return cret
}

//export _gotk4_gio2_VolumeIface_get_symbolic_icon
func _gotk4_gio2_VolumeIface_get_symbolic_icon(arg0 *C.void) (cret *C.GIcon) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	icon := iface.SymbolicIcon()

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(icon).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(icon).Native()))

	return cret
}

//export _gotk4_gio2_VolumeIface_get_uuid
func _gotk4_gio2_VolumeIface_get_uuid(arg0 *C.void) (cret *C.char) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	utf8 := iface.UUID()

	if utf8 != "" {
		cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
	}

	return cret
}

//export _gotk4_gio2_VolumeIface_mount_finish
func _gotk4_gio2_VolumeIface_mount_finish(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	_goerr := iface.MountFinish(_result)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_VolumeIface_removed
func _gotk4_gio2_VolumeIface_removed(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	iface.Removed()
}

//export _gotk4_gio2_VolumeIface_should_automount
func _gotk4_gio2_VolumeIface_should_automount(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(VolumeOverrider)

	ok := iface.ShouldAutomount()

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapVolume(obj *coreglib.Object) *Volume {
	return &Volume{
		Object: obj,
	}
}

func marshalVolume(p uintptr) (interface{}, error) {
	return wrapVolume(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_Volume_ConnectChanged
func _gotk4_gio2_Volume_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectChanged is emitted when the volume has been changed.
func (volume *Volume) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(volume, "changed", false, unsafe.Pointer(C._gotk4_gio2_Volume_ConnectChanged), f)
}

//export _gotk4_gio2_Volume_ConnectRemoved
func _gotk4_gio2_Volume_ConnectRemoved(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectRemoved: this signal is emitted when the #GVolume have been removed.
// If the recipient is holding references to the object they should release them
// so the object can be finalized.
func (volume *Volume) ConnectRemoved(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(volume, "removed", false, unsafe.Pointer(C._gotk4_gio2_Volume_ConnectRemoved), f)
}

// CanEject checks if a volume can be ejected.
//
// The function returns the following values:
//
//    - ok: TRUE if the volume can be ejected. FALSE otherwise.
//
func (volume *Volume) CanEject() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(volume).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(volume)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// CanMount checks if a volume can be mounted.
//
// The function returns the following values:
//
//    - ok: TRUE if the volume can be mounted. FALSE otherwise.
//
func (volume *Volume) CanMount() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(volume).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(volume)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// EjectFinish finishes ejecting a volume. If any errors occurred during the
// operation, error will be set to contain the errors and FALSE will be
// returned.
//
// Deprecated: Use g_volume_eject_with_operation_finish() instead.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (volume *Volume) EjectFinish(result AsyncResulter) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(volume).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	runtime.KeepAlive(volume)
	runtime.KeepAlive(result)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// EjectWithOperationFinish finishes ejecting a volume. If any errors occurred
// during the operation, error will be set to contain the errors and FALSE will
// be returned.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (volume *Volume) EjectWithOperationFinish(result AsyncResulter) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(volume).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	runtime.KeepAlive(volume)
	runtime.KeepAlive(result)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// EnumerateIdentifiers gets the kinds of [identifiers][volume-identifier] that
// volume has. Use g_volume_get_identifier() to obtain the identifiers
// themselves.
//
// The function returns the following values:
//
//    - utf8s: NULL-terminated array of strings containing kinds of identifiers.
//      Use g_strfreev() to free.
//
func (volume *Volume) EnumerateIdentifiers() []string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(volume).Native()))

	_cret = *(***C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(volume)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.void
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

// ActivationRoot gets the activation root for a #GVolume if it is known ahead
// of mount time. Returns NULL otherwise. If not NULL and if volume is mounted,
// then the result of g_mount_get_root() on the #GMount object obtained from
// g_volume_get_mount() will always either be equal or a prefix of what this
// function returns. In other words, in code
//
//    (g_file_has_prefix (volume_activation_root, mount_root) ||
//     g_file_equal (volume_activation_root, mount_root))
//
// will always be TRUE.
//
// Activation roots are typically used in Monitor implementations to find the
// underlying mount to shadow, see g_mount_is_shadowed() for more details.
//
// The function returns the following values:
//
//    - file (optional): activation root of volume or NULL. Use g_object_unref()
//      to free.
//
func (volume *Volume) ActivationRoot() *File {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(volume).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(volume)

	var _file *File // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _file
}

// Drive gets the drive for the volume.
//
// The function returns the following values:
//
//    - drive (optional) or NULL if volume is not associated with a drive. The
//      returned object should be unreffed with g_object_unref() when no longer
//      needed.
//
func (volume *Volume) Drive() *Drive {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(volume).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(volume)

	var _drive *Drive // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_drive = wrapDrive(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _drive
}

// Icon gets the icon for volume.
//
// The function returns the following values:
//
//    - icon: #GIcon. The returned object should be unreffed with
//      g_object_unref() when no longer needed.
//
func (volume *Volume) Icon() *Icon {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(volume).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(volume)

	var _icon *Icon // out

	_icon = wrapIcon(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _icon
}

// Identifier gets the identifier of the given kind for volume. See the
// [introduction][volume-identifier] for more information about volume
// identifiers.
//
// The function takes the following parameters:
//
//    - kind of identifier to return.
//
// The function returns the following values:
//
//    - utf8 (optional): newly allocated string containing the requested
//      identifier, or NULL if the #GVolume doesn't have this kind of identifier.
//
func (volume *Volume) Identifier(kind string) string {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(volume).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(kind)))
	defer C.free(unsafe.Pointer(_args[1]))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(volume)
	runtime.KeepAlive(kind)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Mount gets the mount for the volume.
//
// The function returns the following values:
//
//    - mount (optional) or NULL if volume isn't mounted. The returned object
//      should be unreffed with g_object_unref() when no longer needed.
//
func (volume *Volume) Mount() *Mount {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(volume).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(volume)

	var _mount *Mount // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_mount = wrapMount(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _mount
}

// Name gets the name of volume.
//
// The function returns the following values:
//
//    - utf8: name for the given volume. The returned string should be freed with
//      g_free() when no longer needed.
//
func (volume *Volume) Name() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(volume).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(volume)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SortKey gets the sort key for volume, if any.
//
// The function returns the following values:
//
//    - utf8 (optional): sorting key for volume or NULL if no such key is
//      available.
//
func (volume *Volume) SortKey() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(volume).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(volume)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SymbolicIcon gets the symbolic icon for volume.
//
// The function returns the following values:
//
//    - icon: #GIcon. The returned object should be unreffed with
//      g_object_unref() when no longer needed.
//
func (volume *Volume) SymbolicIcon() *Icon {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(volume).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(volume)

	var _icon *Icon // out

	_icon = wrapIcon(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _icon
}

// UUID gets the UUID for the volume. The reference is typically based on the
// file system UUID for the volume in question and should be considered an
// opaque string. Returns NULL if there is no UUID available.
//
// The function returns the following values:
//
//    - utf8 (optional): UUID for volume or NULL if no UUID can be computed. The
//      returned string should be freed with g_free() when no longer needed.
//
func (volume *Volume) UUID() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(volume).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(volume)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// MountFinish finishes mounting a volume. If any errors occurred during the
// operation, error will be set to contain the errors and FALSE will be
// returned.
//
// If the mount operation succeeded, g_volume_get_mount() on volume is
// guaranteed to return the mount right after calling this function; there's no
// need to listen for the 'mount-added' signal on Monitor.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (volume *Volume) MountFinish(result AsyncResulter) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(volume).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	runtime.KeepAlive(volume)
	runtime.KeepAlive(result)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// ShouldAutomount returns whether the volume should be automatically mounted.
//
// The function returns the following values:
//
//    - ok: TRUE if the volume should be automatically mounted.
//
func (volume *Volume) ShouldAutomount() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(volume).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(volume)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}
