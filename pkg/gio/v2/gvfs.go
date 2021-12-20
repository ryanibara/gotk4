// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// GFile* _gotk4_gio2_VFSFileLookupFunc(GVfs*, char*, gpointer);
// extern void callbackDelete(gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_vfs_get_type()), F: marshalVFSer},
	})
}

// VFS_EXTENSION_POINT_NAME: extension point for #GVfs functionality. See
// [Extending GIO][extending-gio].
const VFS_EXTENSION_POINT_NAME = "gio-vfs"

// VFSFileLookupFunc: this function type is used by g_vfs_register_uri_scheme()
// to make it possible for a client to associate an URI scheme to a different
// #GFile implementation.
//
// The client should return a reference to the new file that has been created
// for uri, or NULL to continue with the default implementation.
type VFSFileLookupFunc func(vfs *VFS, identifier string) (file Filer)

//export _gotk4_gio2_VFSFileLookupFunc
func _gotk4_gio2_VFSFileLookupFunc(arg0 *C.GVfs, arg1 *C.char, arg2 C.gpointer) (cret *C.GFile) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var vfs *VFS          // out
	var identifier string // out

	vfs = wrapVFS(externglib.Take(unsafe.Pointer(arg0)))
	identifier = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	fn := v.(VFSFileLookupFunc)
	file := fn(vfs, identifier)

	cret = (*C.GFile)(unsafe.Pointer(file.Native()))
	C.g_object_ref(C.gpointer(file.Native()))

	return cret
}

// VFSOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type VFSOverrider interface {
	// The function takes the following parameters:
	//
	AddWritableNamespaces(list *FileAttributeInfoList)
	// FileForPath gets a #GFile for path.
	//
	// The function takes the following parameters:
	//
	//    - path: string containing a VFS path.
	//
	// The function returns the following values:
	//
	//    - file: #GFile. Free the returned object with g_object_unref().
	//
	FileForPath(path string) Filer
	// FileForURI gets a #GFile for uri.
	//
	// This operation never fails, but the returned object might not support any
	// I/O operation if the URI is malformed or if the URI scheme is not
	// supported.
	//
	// The function takes the following parameters:
	//
	//    - uri: string containing a URI.
	//
	// The function returns the following values:
	//
	//    - file: #GFile. Free the returned object with g_object_unref().
	//
	FileForURI(uri string) Filer
	// SupportedURISchemes gets a list of URI schemes supported by vfs.
	//
	// The function returns the following values:
	//
	//    - utf8s: NULL-terminated array of strings. The returned array belongs
	//      to GIO and must not be freed or modified.
	//
	SupportedURISchemes() []string
	// IsActive checks if the VFS is active.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if construction of the vfs was successful and it is now
	//      active.
	//
	IsActive() bool
	// The function takes the following parameters:
	//
	//    - source
	//    - dest
	//
	LocalFileMoved(source, dest string)
	// The function takes the following parameters:
	//
	LocalFileRemoved(filename string)
	// The function takes the following parameters:
	//
	//    - ctx (optional)
	//    - filename
	//    - info
	//    - flags
	//
	LocalFileSetAttributes(ctx context.Context, filename string, info *FileInfo, flags FileQueryInfoFlags) error
	// ParseName: this operation never fails, but the returned object might not
	// support any I/O operations if the parse_name cannot be parsed by the
	// #GVfs module.
	//
	// The function takes the following parameters:
	//
	//    - parseName: string to be parsed by the VFS module.
	//
	// The function returns the following values:
	//
	//    - file for the given parse_name. Free the returned object with
	//      g_object_unref().
	//
	ParseName(parseName string) Filer
}

// VFS: entry point for using GIO functionality.
type VFS struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*VFS)(nil)
)

func wrapVFS(obj *externglib.Object) *VFS {
	return &VFS{
		Object: obj,
	}
}

func marshalVFSer(p uintptr) (interface{}, error) {
	return wrapVFS(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// FileForPath gets a #GFile for path.
//
// The function takes the following parameters:
//
//    - path: string containing a VFS path.
//
// The function returns the following values:
//
//    - file: #GFile. Free the returned object with g_object_unref().
//
func (vfs *VFS) FileForPath(path string) Filer {
	var _arg0 *C.GVfs  // out
	var _arg1 *C.char  // out
	var _cret *C.GFile // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(vfs.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_vfs_get_file_for_path(_arg0, _arg1)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(path)

	var _file Filer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Filer is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.Cast()
		rv, ok := casted.(Filer)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.Filer")
		}
		_file = rv
	}

	return _file
}

// FileForURI gets a #GFile for uri.
//
// This operation never fails, but the returned object might not support any I/O
// operation if the URI is malformed or if the URI scheme is not supported.
//
// The function takes the following parameters:
//
//    - uri: string containing a URI.
//
// The function returns the following values:
//
//    - file: #GFile. Free the returned object with g_object_unref().
//
func (vfs *VFS) FileForURI(uri string) Filer {
	var _arg0 *C.GVfs  // out
	var _arg1 *C.char  // out
	var _cret *C.GFile // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(vfs.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_vfs_get_file_for_uri(_arg0, _arg1)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(uri)

	var _file Filer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Filer is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.Cast()
		rv, ok := casted.(Filer)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.Filer")
		}
		_file = rv
	}

	return _file
}

// SupportedURISchemes gets a list of URI schemes supported by vfs.
//
// The function returns the following values:
//
//    - utf8s: NULL-terminated array of strings. The returned array belongs to
//      GIO and must not be freed or modified.
//
func (vfs *VFS) SupportedURISchemes() []string {
	var _arg0 *C.GVfs   // out
	var _cret **C.gchar // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(vfs.Native()))

	_cret = C.g_vfs_get_supported_uri_schemes(_arg0)
	runtime.KeepAlive(vfs)

	var _utf8s []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// IsActive checks if the VFS is active.
//
// The function returns the following values:
//
//    - ok: TRUE if construction of the vfs was successful and it is now active.
//
func (vfs *VFS) IsActive() bool {
	var _arg0 *C.GVfs    // out
	var _cret C.gboolean // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(vfs.Native()))

	_cret = C.g_vfs_is_active(_arg0)
	runtime.KeepAlive(vfs)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ParseName: this operation never fails, but the returned object might not
// support any I/O operations if the parse_name cannot be parsed by the #GVfs
// module.
//
// The function takes the following parameters:
//
//    - parseName: string to be parsed by the VFS module.
//
// The function returns the following values:
//
//    - file for the given parse_name. Free the returned object with
//      g_object_unref().
//
func (vfs *VFS) ParseName(parseName string) Filer {
	var _arg0 *C.GVfs  // out
	var _arg1 *C.char  // out
	var _cret *C.GFile // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(vfs.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(parseName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_vfs_parse_name(_arg0, _arg1)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(parseName)

	var _file Filer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Filer is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.Cast()
		rv, ok := casted.(Filer)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.Filer")
		}
		_file = rv
	}

	return _file
}

// RegisterURIScheme registers uri_func and parse_name_func as the #GFile URI
// and parse name lookup functions for URIs with a scheme matching scheme. Note
// that scheme is registered only within the running application, as opposed to
// desktop-wide as it happens with GVfs backends.
//
// When a #GFile is requested with an URI containing scheme (e.g. through
// g_file_new_for_uri()), uri_func will be called to allow a custom constructor.
// The implementation of uri_func should not be blocking, and must not call
// g_vfs_register_uri_scheme() or g_vfs_unregister_uri_scheme().
//
// When g_file_parse_name() is called with a parse name obtained from such file,
// parse_name_func will be called to allow the #GFile to be created again. In
// that case, it's responsibility of parse_name_func to make sure the parse name
// matches what the custom #GFile implementation returned when
// g_file_get_parse_name() was previously called. The implementation of
// parse_name_func should not be blocking, and must not call
// g_vfs_register_uri_scheme() or g_vfs_unregister_uri_scheme().
//
// It's an error to call this function twice with the same scheme. To unregister
// a custom URI scheme, use g_vfs_unregister_uri_scheme().
//
// The function takes the following parameters:
//
//    - scheme: URI scheme, e.g. "http".
//    - uriFunc (optional): FileLookupFunc.
//    - parseNameFunc (optional): FileLookupFunc.
//
// The function returns the following values:
//
//    - ok: TRUE if scheme was successfully registered, or FALSE if a handler for
//      scheme already exists.
//
func (vfs *VFS) RegisterURIScheme(scheme string, uriFunc, parseNameFunc VFSFileLookupFunc) bool {
	var _arg0 *C.GVfs              // out
	var _arg1 *C.char              // out
	var _arg2 C.GVfsFileLookupFunc // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify
	var _arg5 C.GVfsFileLookupFunc // out
	var _arg6 C.gpointer
	var _arg7 C.GDestroyNotify
	var _cret C.gboolean // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(vfs.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(scheme)))
	defer C.free(unsafe.Pointer(_arg1))
	if uriFunc != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_VFSFileLookupFunc)
		_arg3 = C.gpointer(gbox.Assign(uriFunc))
		_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}
	if parseNameFunc != nil {
		_arg5 = (*[0]byte)(C._gotk4_gio2_VFSFileLookupFunc)
		_arg6 = C.gpointer(gbox.Assign(parseNameFunc))
		_arg7 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_cret = C.g_vfs_register_uri_scheme(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(scheme)
	runtime.KeepAlive(uriFunc)
	runtime.KeepAlive(parseNameFunc)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnregisterURIScheme unregisters the URI handler for scheme previously
// registered with g_vfs_register_uri_scheme().
//
// The function takes the following parameters:
//
//    - scheme: URI scheme, e.g. "http".
//
// The function returns the following values:
//
//    - ok: TRUE if scheme was successfully unregistered, or FALSE if a handler
//      for scheme does not exist.
//
func (vfs *VFS) UnregisterURIScheme(scheme string) bool {
	var _arg0 *C.GVfs    // out
	var _arg1 *C.char    // out
	var _cret C.gboolean // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(vfs.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(scheme)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_vfs_unregister_uri_scheme(_arg0, _arg1)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(scheme)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VFSGetDefault gets the default #GVfs for the system.
//
// The function returns the following values:
//
//    - vfs which will be the local file system #GVfs if no other implementation
//      is available.
//
func VFSGetDefault() *VFS {
	var _cret *C.GVfs // in

	_cret = C.g_vfs_get_default()

	var _vfs *VFS // out

	_vfs = wrapVFS(externglib.Take(unsafe.Pointer(_cret)))

	return _vfs
}

// VFSGetLocal gets the local #GVfs for the system.
//
// The function returns the following values:
//
//    - vfs: #GVfs.
//
func VFSGetLocal() *VFS {
	var _cret *C.GVfs // in

	_cret = C.g_vfs_get_local()

	var _vfs *VFS // out

	_vfs = wrapVFS(externglib.Take(unsafe.Pointer(_cret)))

	return _vfs
}
