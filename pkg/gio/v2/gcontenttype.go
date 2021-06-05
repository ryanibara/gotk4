// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
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
import "C"

// ContentTypeCanBeExecutable checks if a content type can be executable. Note
// that for instance things like text files can be executables (i.e. scripts and
// batch files).
func ContentTypeCanBeExecutable(typ string) bool {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ret1 bool

	cret = C.g_content_type_can_be_executable(typ)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// ContentTypeEquals compares two content types for equality.
func ContentTypeEquals(type1 string, type2 string) bool {
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg1 = (*C.gchar)(C.CString(type1))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(type2))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.gboolean
	var ret1 bool

	cret = C.g_content_type_equals(type1, type2)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// ContentTypeFromMIMEType tries to find a content type based on the mime type
// name.
func ContentTypeFromMIMEType(mimeType string) string {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var ret1 string

	cret = C.g_content_type_from_mime_type(mimeType)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// ContentTypeGetDescription gets the human readable description of the content
// type.
func ContentTypeGetDescription(typ string) string {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var ret1 string

	cret = C.g_content_type_get_description(typ)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// ContentTypeGetGenericIconName gets the generic icon name for a content type.
//
// See the shared-mime-info
// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
// specification for more on the generic icon name.
func ContentTypeGetGenericIconName(typ string) string {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var ret1 string

	cret = C.g_content_type_get_generic_icon_name(typ)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// ContentTypeGetIcon gets the icon for a content type.
func ContentTypeGetIcon(typ string) Icon {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GIcon
	var ret1 Icon

	cret = C.g_content_type_get_icon(typ)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Icon)

	return ret1
}

// ContentTypeGetMIMEDirs: get the list of directories which MIME data is loaded
// from. See g_content_type_set_mime_dirs() for details.
func ContentTypeGetMIMEDirs() []string {
	var cret **C.gchar
	var ret1 []string

	cret = C.g_content_type_get_mime_dirs()

	{
		var length int
		for p := cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		ret1 = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
			ret1[i] = C.GoString(src)
		}
	}

	return ret1
}

// ContentTypeGetMIMEType gets the mime type for the content type, if one is
// registered.
func ContentTypeGetMIMEType(typ string) string {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var ret1 string

	cret = C.g_content_type_get_mime_type(typ)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// ContentTypeGetSymbolicIcon gets the symbolic icon for a content type.
func ContentTypeGetSymbolicIcon(typ string) Icon {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GIcon
	var ret1 Icon

	cret = C.g_content_type_get_symbolic_icon(typ)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Icon)

	return ret1
}

// ContentTypeGuessForTree tries to guess the type of the tree with root @root,
// by looking at the files it contains. The result is an array of content types,
// with the best guess coming first.
//
// The types returned all have the form x-content/foo, e.g. x-content/audio-cdda
// (for audio CDs) or x-content/image-dcf (for a camera memory card). See the
// shared-mime-info
// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
// specification for more on x-content types.
//
// This function is useful in the implementation of
// g_mount_guess_content_type().
func ContentTypeGuessForTree(root File) []string {
	var arg1 *C.GFile

	arg1 = (*C.GFile)(unsafe.Pointer(root.Native()))

	var cret **C.gchar
	var ret1 []string

	cret = C.g_content_type_guess_for_tree(root)

	{
		var length int
		for p := cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		ret1 = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
			ret1[i] = C.GoString(src)
			defer C.free(unsafe.Pointer(src))
		}
	}

	return ret1
}

// ContentTypeIsA determines if @type is a subset of @supertype.
func ContentTypeIsA(typ string, supertype string) bool {
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(supertype))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.gboolean
	var ret1 bool

	cret = C.g_content_type_is_a(typ, supertype)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// ContentTypeIsMIMEType determines if @type is a subset of @mime_type.
// Convenience wrapper around g_content_type_is_a().
func ContentTypeIsMIMEType(typ string, mimeType string) bool {
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.gboolean
	var ret1 bool

	cret = C.g_content_type_is_mime_type(typ, mimeType)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// ContentTypeIsUnknown checks if the content type is the generic "unknown"
// type. On UNIX this is the "application/octet-stream" mimetype, while on win32
// it is "*" and on OSX it is a dynamic type or octet-stream.
func ContentTypeIsUnknown(typ string) bool {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(typ))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ret1 bool

	cret = C.g_content_type_is_unknown(typ)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// ContentTypeSetMIMEDirs: set the list of directories used by GIO to load the
// MIME database. If @dirs is nil, the directories used are the default:
//
//    - the `mime` subdirectory of the directory in `$XDG_DATA_HOME`
//    - the `mime` subdirectory of every directory in `$XDG_DATA_DIRS`
//
// This function is intended to be used when writing tests that depend on
// information stored in the MIME database, in order to control the data.
//
// Typically, in case your tests use G_TEST_OPTION_ISOLATE_DIRS, but they depend
// on the system’s MIME database, you should call this function with @dirs set
// to nil before calling g_test_init(), for instance:
//
//      // Load MIME data from the system
//      g_content_type_set_mime_dirs (NULL);
//      // Isolate the environment
//      g_test_init (&argc, &argv, G_TEST_OPTION_ISOLATE_DIRS, NULL);
//
//      …
//
//      return g_test_run ();
func ContentTypeSetMIMEDirs(dirs []string) {
	var arg1 **C.gchar

	arg1 = C.malloc(len(dirs) * (unsafe.Sizeof(int(0)) + 1))
	defer C.free(unsafe.Pointer(arg1))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg1), int(len(dirs)))

		for i := range dirs {
			out[i] = (*C.gchar)(C.CString(dirs[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.g_content_type_set_mime_dirs(dirs)
}

// ContentTypesGetRegistered gets a list of strings containing all the
// registered content types known to the system. The list and its data should be
// freed using `g_list_free_full (list, g_free)`.
func ContentTypesGetRegistered() *glib.List {
	var cret *C.GList
	var ret1 *glib.List

	cret = C.g_content_types_get_registered()

	ret1 = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}
