// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// GetLocaleVariants returns a list of derived variants of locale, which can be
// used to e.g. construct locale-dependent filenames or search paths. The
// returned list is sorted from most desirable to least desirable. This function
// handles territory, charset and extra locale modifiers. See setlocale(3)
// (man:setlocale) for information about locales and their format.
//
// locale itself is guaranteed to be returned in the output.
//
// For example, if locale is fr_BE, then the returned list is fr_BE, fr. If
// locale is en_GB.UTF-8euro, then the returned list is en_GB.UTF-8euro,
// en_GB.UTF-8, en_GBeuro, en_GB, en.UTF-8euro, en.UTF-8, eneuro, en.
//
// If you need the list of variants for the current locale, use
// g_get_language_names().
//
// The function takes the following parameters:
//
//    - locale identifier.
//
// The function returns the following values:
//
//    - utf8s: newly allocated array of newly allocated strings with the locale
//      variants. Free with g_strfreev().
//
func GetLocaleVariants(locale string) []string {
	var _arg1 *C.gchar  // out
	var _cret **C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(locale)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_get_locale_variants(_arg1)
	runtime.KeepAlive(locale)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
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
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}
