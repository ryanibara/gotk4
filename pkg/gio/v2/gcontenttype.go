// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// ContentTypeCanBeExecutable checks if a content type can be executable.
// Note that for instance things like text files can be executables (i.e.
// scripts and batch files).
//
// The function takes the following parameters:
//
//   - typ: content type string.
//
// The function returns the following values:
//
//   - ok: TRUE if the file type corresponds to a type that can be executable,
//     FALSE otherwise.
//
func ContentTypeCanBeExecutable(typ string) bool {
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_content_type_can_be_executable(_arg1)
	runtime.KeepAlive(typ)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContentTypeEquals compares two content types for equality.
//
// The function takes the following parameters:
//
//   - type1: content type string.
//   - type2: content type string.
//
// The function returns the following values:
//
//   - ok: TRUE if the two strings are identical or equivalent, FALSE otherwise.
//
func ContentTypeEquals(type1, type2 string) bool {
	var _arg1 *C.gchar   // out
	var _arg2 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(type1)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(type2)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_content_type_equals(_arg1, _arg2)
	runtime.KeepAlive(type1)
	runtime.KeepAlive(type2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContentTypeGetDescription gets the human readable description of the content
// type.
//
// The function takes the following parameters:
//
//   - typ: content type string.
//
// The function returns the following values:
//
//   - utf8: short description of the content type type. Free the returned
//     string with g_free().
//
func ContentTypeGetDescription(typ string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_content_type_get_description(_arg1)
	runtime.KeepAlive(typ)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ContentTypeGetIcon gets the icon for a content type.
//
// The function takes the following parameters:
//
//   - typ: content type string.
//
// The function returns the following values:
//
//   - icon corresponding to the content type. Free the returned object with
//     g_object_unref().
//
func ContentTypeGetIcon(typ string) *Icon {
	var _arg1 *C.gchar // out
	var _cret *C.GIcon // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_content_type_get_icon(_arg1)
	runtime.KeepAlive(typ)

	var _icon *Icon // out

	_icon = wrapIcon(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _icon
}

// ContentTypeGetMIMEType gets the mime type for the content type, if one is
// registered.
//
// The function takes the following parameters:
//
//   - typ: content type string.
//
// The function returns the following values:
//
//   - utf8 (optional): registered mime type for the given type, or NULL if
//     unknown; free with g_free().
//
func ContentTypeGetMIMEType(typ string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_content_type_get_mime_type(_arg1)
	runtime.KeepAlive(typ)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// ContentTypeGuess guesses the content type based on example data. If the
// function is uncertain, result_uncertain will be set to TRUE. Either filename
// or data may be NULL, in which case the guess will be based solely on the
// other argument.
//
// The function takes the following parameters:
//
//   - filename (optional): string, or NULL.
//   - data (optional): stream of data, or NULL.
//
// The function returns the following values:
//
//   - resultUncertain (optional): return location for the certainty of the
//     result, or NULL.
//   - utf8: string indicating a guessed content type for the given data.
//     Free with g_free().
//
func ContentTypeGuess(filename string, data []byte) (bool, string) {
	var _arg1 *C.gchar  // out
	var _arg2 *C.guchar // out
	var _arg3 C.gsize
	var _arg4 C.gboolean // in
	var _cret *C.gchar   // in

	if filename != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg3 = (C.gsize)(len(data))
	if len(data) > 0 {
		_arg2 = (*C.guchar)(unsafe.Pointer(&data[0]))
	}

	_cret = C.g_content_type_guess(_arg1, _arg2, _arg3, &_arg4)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(data)

	var _resultUncertain bool // out
	var _utf8 string          // out

	if _arg4 != 0 {
		_resultUncertain = true
	}
	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _resultUncertain, _utf8
}

// ContentTypeIsA determines if type is a subset of supertype.
//
// The function takes the following parameters:
//
//   - typ: content type string.
//   - supertype: content type string.
//
// The function returns the following values:
//
//   - ok: TRUE if type is a kind of supertype, FALSE otherwise.
//
func ContentTypeIsA(typ, supertype string) bool {
	var _arg1 *C.gchar   // out
	var _arg2 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(supertype)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_content_type_is_a(_arg1, _arg2)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(supertype)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContentTypeIsUnknown checks if the content type is the generic "unknown"
// type. On UNIX this is the "application/octet-stream" mimetype, while on win32
// it is "*" and on OSX it is a dynamic type or octet-stream.
//
// The function takes the following parameters:
//
//   - typ: content type string.
//
// The function returns the following values:
//
//   - ok: TRUE if the type is the unknown type.
//
func ContentTypeIsUnknown(typ string) bool {
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_content_type_is_unknown(_arg1)
	runtime.KeepAlive(typ)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContentTypesGetRegistered gets a list of strings containing all the
// registered content types known to the system. The list and its data should be
// freed using g_list_free_full (list, g_free).
//
// The function returns the following values:
//
//   - list of the registered content types.
//
func ContentTypesGetRegistered() []string {
	var _cret *C.GList // in

	_cret = C.g_content_types_get_registered()

	var _list []string // out

	_list = make([]string, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.gchar)(v)
		var dst string // out
		dst = C.GoString((*C.gchar)(unsafe.Pointer(src)))
		defer C.free(unsafe.Pointer(src))
		_list = append(_list, dst)
	})

	return _list
}
