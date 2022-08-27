// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// BuildFilenamev behaves exactly like g_build_filename(), but takes the path
// elements as a string array, instead of varargs. This function is mainly meant
// for language bindings.
//
// The function takes the following parameters:
//
//    - args: NULL-terminated array of strings containing the path elements.
//
// The function returns the following values:
//
//    - filename: newly-allocated string that must be freed with g_free().
//
func BuildFilenamev(args []string) string {
	var _arg1 **C.gchar // out
	var _cret *C.gchar  // in

	{
		_arg1 = (**C.gchar)(C.calloc(C.size_t((len(args) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(args)+1)
			var zero *C.gchar
			out[len(args)] = zero
			for i := range args {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(args[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.g_build_filenamev(_arg1)
	runtime.KeepAlive(args)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// BuildPathv behaves exactly like g_build_path(), but takes the path elements
// as a string array, instead of varargs. This function is mainly meant for
// language bindings.
//
// The function takes the following parameters:
//
//    - separator: string used to separator the elements of the path.
//    - args: NULL-terminated array of strings containing the path elements.
//
// The function returns the following values:
//
//    - filename: newly-allocated string that must be freed with g_free().
//
func BuildPathv(separator string, args []string) string {
	var _arg1 *C.gchar  // out
	var _arg2 **C.gchar // out
	var _cret *C.gchar  // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(separator)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		_arg2 = (**C.gchar)(C.calloc(C.size_t((len(args) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(args)+1)
			var zero *C.gchar
			out[len(args)] = zero
			for i := range args {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(args[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.g_build_pathv(_arg1, _arg2)
	runtime.KeepAlive(separator)
	runtime.KeepAlive(args)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// FileSetContents writes all of contents to a file named filename. This is a
// convenience wrapper around calling g_file_set_contents_full() with flags set
// to G_FILE_SET_CONTENTS_CONSISTENT | G_FILE_SET_CONTENTS_ONLY_EXISTING and
// mode set to 0666.
//
// The function takes the following parameters:
//
//    - filename: name of a file to write contents to, in the GLib file name
//      encoding.
//    - contents: string to write to the file.
//
func FileSetContents(filename, contents string) error {
	var _arg1 *C.gchar // out
	var _arg2 *C.gchar // out
	var _arg3 C.gssize
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg3 = (C.gssize)(len(contents))
	_arg2 = (*C.gchar)(C.calloc(C.size_t((len(contents) + 1)), C.size_t(C.sizeof_gchar)))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(_arg2)), len(contents)), contents)
	defer C.free(unsafe.Pointer(_arg2))

	C.g_file_set_contents(_arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(contents)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// MkdirWithParents: create a directory if it doesn't already exist. Create
// intermediate parent directories as needed, too.
//
// The function takes the following parameters:
//
//    - pathname in the GLib file name encoding.
//    - mode permissions to use for newly created directories.
//
// The function returns the following values:
//
//    - gint: 0 if the directory already exists, or was successfully created.
//      Returns -1 if an error occurred, with errno set.
//
func MkdirWithParents(pathname string, mode int) int {
	var _arg1 *C.gchar // out
	var _arg2 C.gint   // out
	var _cret C.gint   // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(pathname)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(mode)

	_cret = C.g_mkdir_with_parents(_arg1, _arg2)
	runtime.KeepAlive(pathname)
	runtime.KeepAlive(mode)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
