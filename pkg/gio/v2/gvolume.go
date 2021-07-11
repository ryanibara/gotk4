// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
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
		{T: externglib.Type(C.g_volume_get_type()), F: marshalVolumer},
	})
}

// VolumeOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type VolumeOverrider interface {
	// CanEject checks if a volume can be ejected.
	CanEject() bool
	// CanMount checks if a volume can be mounted.
	CanMount() bool
	//
	Changed()
	// EjectFinish finishes ejecting a volume. If any errors occurred during the
	// operation, @error will be set to contain the errors and false will be
	// returned.
	//
	// Deprecated: Use g_volume_eject_with_operation_finish() instead.
	EjectFinish(result AsyncResulter) error
	// EjectWithOperationFinish finishes ejecting a volume. If any errors
	// occurred during the operation, @error will be set to contain the errors
	// and false will be returned.
	EjectWithOperationFinish(result AsyncResulter) error
	// EnumerateIdentifiers gets the kinds of [identifiers][volume-identifier]
	// that @volume has. Use g_volume_get_identifier() to obtain the identifiers
	// themselves.
	EnumerateIdentifiers() []string
	// ActivationRoot gets the activation root for a #GVolume if it is known
	// ahead of mount time. Returns nil otherwise. If not nil and if @volume is
	// mounted, then the result of g_mount_get_root() on the #GMount object
	// obtained from g_volume_get_mount() will always either be equal or a
	// prefix of what this function returns. In other words, in code
	//
	//    (g_file_has_prefix (volume_activation_root, mount_root) ||
	//     g_file_equal (volume_activation_root, mount_root))
	//
	// will always be true.
	//
	// Activation roots are typically used in Monitor implementations to find
	// the underlying mount to shadow, see g_mount_is_shadowed() for more
	// details.
	ActivationRoot() *File
	// Drive gets the drive for the @volume.
	Drive() *Drive
	// Icon gets the icon for @volume.
	Icon() *Icon
	// Identifier gets the identifier of the given kind for @volume. See the
	// [introduction][volume-identifier] for more information about volume
	// identifiers.
	Identifier(kind string) string
	// Mount gets the mount for the @volume.
	Mount() *Mount
	// Name gets the name of @volume.
	Name() string
	// SortKey gets the sort key for @volume, if any.
	SortKey() string
	// SymbolicIcon gets the symbolic icon for @volume.
	SymbolicIcon() *Icon
	// UUID gets the UUID for the @volume. The reference is typically based on
	// the file system UUID for the volume in question and should be considered
	// an opaque string. Returns nil if there is no UUID available.
	UUID() string
	// MountFinish finishes mounting a volume. If any errors occurred during the
	// operation, @error will be set to contain the errors and false will be
	// returned.
	//
	// If the mount operation succeeded, g_volume_get_mount() on @volume is
	// guaranteed to return the mount right after calling this function; there's
	// no need to listen for the 'mount-added' signal on Monitor.
	MountFinish(result AsyncResulter) error
	//
	Removed()
	// ShouldAutomount returns whether the volume should be automatically
	// mounted.
	ShouldAutomount() bool
}

// Volumer describes Volume's methods.
type Volumer interface {
	// CanEject checks if a volume can be ejected.
	CanEject() bool
	// CanMount checks if a volume can be mounted.
	CanMount() bool
	// EjectFinish finishes ejecting a volume.
	EjectFinish(result AsyncResulter) error
	// EjectWithOperationFinish finishes ejecting a volume.
	EjectWithOperationFinish(result AsyncResulter) error
	// EnumerateIdentifiers gets the kinds of [identifiers][volume-identifier]
	// that @volume has.
	EnumerateIdentifiers() []string
	// ActivationRoot gets the activation root for a #GVolume if it is known
	// ahead of mount time.
	ActivationRoot() *File
	// Drive gets the drive for the @volume.
	Drive() *Drive
	// Icon gets the icon for @volume.
	Icon() *Icon
	// Identifier gets the identifier of the given kind for @volume.
	Identifier(kind string) string
	// Mount gets the mount for the @volume.
	Mount() *Mount
	// Name gets the name of @volume.
	Name() string
	// SortKey gets the sort key for @volume, if any.
	SortKey() string
	// SymbolicIcon gets the symbolic icon for @volume.
	SymbolicIcon() *Icon
	// UUID gets the UUID for the @volume.
	UUID() string
	// MountFinish finishes mounting a volume.
	MountFinish(result AsyncResulter) error
	// ShouldAutomount returns whether the volume should be automatically
	// mounted.
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
// Typically, one will only want to pass nil for the Operation if automounting
// all volumes when a desktop session starts since it's not desirable to put up
// a lot of dialogs asking for credentials.
//
// The callback will be fired when the operation has resolved (either with
// success or failure), and a Result instance will be passed to the callback.
// That callback should then call g_volume_mount_finish() with the #GVolume
// instance and the Result data to see if the operation was completed
// successfully. If an @error is present when g_volume_mount_finish() is called,
// then it will be filled with any error information.
//
//
// Volume Identifiers
//
// It is sometimes necessary to directly access the underlying operating system
// object behind a volume (e.g. for passing a volume to an application via the
// commandline). For this purpose, GIO allows to obtain an 'identifier' for the
// volume. There can be different kinds of identifiers, such as Hal UDIs,
// filesystem labels, traditional Unix devices (e.g. `/dev/sda2`), UUIDs. GIO
// uses predefined strings as names for the different kinds of identifiers:
// VOLUME_IDENTIFIER_KIND_UUID, VOLUME_IDENTIFIER_KIND_LABEL, etc. Use
// g_volume_get_identifier() to obtain an identifier for a volume.
//
//    Note that VOLUME_IDENTIFIER_KIND_HAL_UDI will only be available when the gvfs hal volume monitor is in use. Other volume monitors will generally be able to provide the VOLUME_IDENTIFIER_KIND_UNIX_DEVICE identifier, which can be used to obtain a hal device by means of libhal_manager_find_device_string_match().
type Volume struct {
	*externglib.Object
}

var (
	_ Volumer         = (*Volume)(nil)
	_ gextras.Nativer = (*Volume)(nil)
)

func wrapVolume(obj *externglib.Object) Volumer {
	return &Volume{
		Object: obj,
	}
}

func marshalVolumer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapVolume(obj), nil
}

// CanEject checks if a volume can be ejected.
func (volume *Volume) CanEject() bool {
	var _arg0 *C.GVolume // out
	var _cret C.gboolean // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_can_eject(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanMount checks if a volume can be mounted.
func (volume *Volume) CanMount() bool {
	var _arg0 *C.GVolume // out
	var _cret C.gboolean // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_can_mount(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// EjectFinish finishes ejecting a volume. If any errors occurred during the
// operation, @error will be set to contain the errors and false will be
// returned.
//
// Deprecated: Use g_volume_eject_with_operation_finish() instead.
func (volume *Volume) EjectFinish(result AsyncResulter) error {
	var _arg0 *C.GVolume      // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	C.g_volume_eject_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// EjectWithOperationFinish finishes ejecting a volume. If any errors occurred
// during the operation, @error will be set to contain the errors and false will
// be returned.
func (volume *Volume) EjectWithOperationFinish(result AsyncResulter) error {
	var _arg0 *C.GVolume      // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	C.g_volume_eject_with_operation_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// EnumerateIdentifiers gets the kinds of [identifiers][volume-identifier] that
// @volume has. Use g_volume_get_identifier() to obtain the identifiers
// themselves.
func (volume *Volume) EnumerateIdentifiers() []string {
	var _arg0 *C.GVolume // out
	var _cret **C.char

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_enumerate_identifiers(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
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
// of mount time. Returns nil otherwise. If not nil and if @volume is mounted,
// then the result of g_mount_get_root() on the #GMount object obtained from
// g_volume_get_mount() will always either be equal or a prefix of what this
// function returns. In other words, in code
//
//    (g_file_has_prefix (volume_activation_root, mount_root) ||
//     g_file_equal (volume_activation_root, mount_root))
//
// will always be true.
//
// Activation roots are typically used in Monitor implementations to find the
// underlying mount to shadow, see g_mount_is_shadowed() for more details.
func (volume *Volume) ActivationRoot() *File {
	var _arg0 *C.GVolume // out
	var _cret *C.GFile   // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_get_activation_root(_arg0)

	var _file *File // out

	_file = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*File)

	return _file
}

// Drive gets the drive for the @volume.
func (volume *Volume) Drive() *Drive {
	var _arg0 *C.GVolume // out
	var _cret *C.GDrive  // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_get_drive(_arg0)

	var _drive *Drive // out

	_drive = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*Drive)

	return _drive
}

// Icon gets the icon for @volume.
func (volume *Volume) Icon() *Icon {
	var _arg0 *C.GVolume // out
	var _cret *C.GIcon   // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_get_icon(_arg0)

	var _icon *Icon // out

	_icon = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*Icon)

	return _icon
}

// Identifier gets the identifier of the given kind for @volume. See the
// [introduction][volume-identifier] for more information about volume
// identifiers.
func (volume *Volume) Identifier(kind string) string {
	var _arg0 *C.GVolume // out
	var _arg1 *C.char    // out
	var _cret *C.char    // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(kind)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_volume_get_identifier(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Mount gets the mount for the @volume.
func (volume *Volume) Mount() *Mount {
	var _arg0 *C.GVolume // out
	var _cret *C.GMount  // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_get_mount(_arg0)

	var _mount *Mount // out

	_mount = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*Mount)

	return _mount
}

// Name gets the name of @volume.
func (volume *Volume) Name() string {
	var _arg0 *C.GVolume // out
	var _cret *C.char    // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SortKey gets the sort key for @volume, if any.
func (volume *Volume) SortKey() string {
	var _arg0 *C.GVolume // out
	var _cret *C.gchar   // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_get_sort_key(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SymbolicIcon gets the symbolic icon for @volume.
func (volume *Volume) SymbolicIcon() *Icon {
	var _arg0 *C.GVolume // out
	var _cret *C.GIcon   // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_get_symbolic_icon(_arg0)

	var _icon *Icon // out

	_icon = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*Icon)

	return _icon
}

// UUID gets the UUID for the @volume. The reference is typically based on the
// file system UUID for the volume in question and should be considered an
// opaque string. Returns nil if there is no UUID available.
func (volume *Volume) UUID() string {
	var _arg0 *C.GVolume // out
	var _cret *C.char    // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_get_uuid(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// MountFinish finishes mounting a volume. If any errors occurred during the
// operation, @error will be set to contain the errors and false will be
// returned.
//
// If the mount operation succeeded, g_volume_get_mount() on @volume is
// guaranteed to return the mount right after calling this function; there's no
// need to listen for the 'mount-added' signal on Monitor.
func (volume *Volume) MountFinish(result AsyncResulter) error {
	var _arg0 *C.GVolume      // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	C.g_volume_mount_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// ShouldAutomount returns whether the volume should be automatically mounted.
func (volume *Volume) ShouldAutomount() bool {
	var _arg0 *C.GVolume // out
	var _cret C.gboolean // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_should_automount(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
