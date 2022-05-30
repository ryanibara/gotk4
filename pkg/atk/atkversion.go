// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// BINARY_AGE: like atk_get_binary_age(), but from the headers used at
// application compile time, rather than from the library linked against at
// application run time.
const BINARY_AGE = 23610

// INTERFACE_AGE: like atk_get_interface_age(), but from the headers used at
// application compile time, rather than from the library linked against at
// application run time.
const INTERFACE_AGE = 1

// MAJOR_VERSION: like atk_get_major_version(), but from the headers used at
// application compile time, rather than from the library linked against at
// application run time.
const MAJOR_VERSION = 2

// MICRO_VERSION: like atk_get_micro_version(), but from the headers used at
// application compile time, rather than from the library linked against at
// application run time.
const MICRO_VERSION = 0

// MINOR_VERSION: like atk_get_minor_version(), but from the headers used at
// application compile time, rather than from the library linked against at
// application run time.
const MINOR_VERSION = 36

// VERSION_MIN_REQUIRED: macro that should be defined by the user prior to
// including the atk/atk.h header. The definition should be one of the
// predefined ATK version macros: ATK_VERSION_2_12, ATK_VERSION_2_14,...
//
// This macro defines the earliest version of ATK that the package is required
// to be able to compile against.
//
// If the compiler is configured to warn about the use of deprecated functions,
// then using functions that were deprecated in version ATK_VERSION_MIN_REQUIRED
// or earlier will cause warnings (but using functions deprecated in later
// releases will not).
const VERSION_MIN_REQUIRED = 2

// GetBinaryAge returns the binary age as passed to libtool when building the
// ATK library the process is running against.
//
// The function returns the following values:
//
//    - guint: binary age of the ATK library.
//
func GetBinaryAge() uint {
	var _cret C.guint // in

	_gret := girepository.MustFind("Atk", "get_binary_age").Invoke(nil, nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// GetInterfaceAge returns the interface age as passed to libtool when building
// the ATK library the process is running against.
//
// The function returns the following values:
//
//    - guint: interface age of the ATK library.
//
func GetInterfaceAge() uint {
	var _cret C.guint // in

	_gret := girepository.MustFind("Atk", "get_interface_age").Invoke(nil, nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// GetMajorVersion returns the major version number of the ATK library. (e.g. in
// ATK version 2.7.4 this is 2.)
//
// This function is in the library, so it represents the ATK library your code
// is running against. In contrast, the K_MAJOR_VERSION macro represents the
// major version of the ATK headers you have included when compiling your code.
//
// The function returns the following values:
//
//    - guint: major version number of the ATK library.
//
func GetMajorVersion() uint {
	var _cret C.guint // in

	_gret := girepository.MustFind("Atk", "get_major_version").Invoke(nil, nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// GetMicroVersion returns the micro version number of the ATK library. (e.g. in
// ATK version 2.7.4 this is 4.)
//
// This function is in the library, so it represents the ATK library your code
// is are running against. In contrast, the K_MICRO_VERSION macro represents the
// micro version of the ATK headers you have included when compiling your code.
//
// The function returns the following values:
//
//    - guint: micro version number of the ATK library.
//
func GetMicroVersion() uint {
	var _cret C.guint // in

	_gret := girepository.MustFind("Atk", "get_micro_version").Invoke(nil, nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// GetMinorVersion returns the minor version number of the ATK library. (e.g. in
// ATK version 2.7.4 this is 7.)
//
// This function is in the library, so it represents the ATK library your code
// is are running against. In contrast, the K_MINOR_VERSION macro represents the
// minor version of the ATK headers you have included when compiling your code.
//
// The function returns the following values:
//
//    - guint: minor version number of the ATK library.
//
func GetMinorVersion() uint {
	var _cret C.guint // in

	_gret := girepository.MustFind("Atk", "get_minor_version").Invoke(nil, nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}
