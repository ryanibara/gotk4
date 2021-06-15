// Code generated by girgen. DO NOT EDIT.

package gtk

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

// CheckVersion checks that the GTK library in use is compatible with the given
// version.
//
// Generally you would pass in the constants GTK_MAJOR_VERSION,
// GTK_MINOR_VERSION, GTK_MICRO_VERSION as the three arguments to this function;
// that produces a check that the library in use is compatible with the version
// of GTK the application or module was compiled against.
//
// Compatibility is defined by two things: first the version of the running
// library is newer than the version
// @required_major.required_minor.@required_micro. Second the running library
// must be binary compatible with the version
// @required_major.required_minor.@required_micro (same major version.)
//
// This function is primarily for GTK modules; the module can call this function
// to check that it wasn’t loaded into an incompatible version of GTK. However,
// such a check isn’t completely reliable, since the module may be linked
// against an old version of GTK and calling the old version of
// gtk_check_version(), but still get loaded into an application using a newer
// version of GTK.
func CheckVersion(requiredMajor uint, requiredMinor uint, requiredMicro uint) string {
	var _arg1 C.guint // out
	var _arg2 C.guint // out
	var _arg3 C.guint // out
	var _cret *C.char // in

	_arg1 = (C.guint)(requiredMajor)
	_arg2 = (C.guint)(requiredMinor)
	_arg3 = (C.guint)(requiredMicro)

	_cret = C.gtk_check_version(_arg1, _arg2, _arg3)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// GetBinaryAge returns the binary age as passed to `libtool`.
//
// If `libtool` means nothing to you, don't worry about it.
func GetBinaryAge() uint {
	var _cret C.guint // in

	_cret = C.gtk_get_binary_age()

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// GetInterfaceAge returns the interface age as passed to `libtool`.
//
// If `libtool` means nothing to you, don't worry about it.
func GetInterfaceAge() uint {
	var _cret C.guint // in

	_cret = C.gtk_get_interface_age()

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// GetMajorVersion returns the major version number of the GTK library.
//
// For example, in GTK version 3.1.5 this is 3.
//
// This function is in the library, so it represents the GTK library your code
// is running against. Contrast with the GTK_MAJOR_VERSION macro, which
// represents the major version of the GTK headers you have included when
// compiling your code.
func GetMajorVersion() uint {
	var _cret C.guint // in

	_cret = C.gtk_get_major_version()

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// GetMicroVersion returns the micro version number of the GTK library.
//
// For example, in GTK version 3.1.5 this is 5.
//
// This function is in the library, so it represents the GTK library your code
// is are running against. Contrast with the GTK_MICRO_VERSION macro, which
// represents the micro version of the GTK headers you have included when
// compiling your code.
func GetMicroVersion() uint {
	var _cret C.guint // in

	_cret = C.gtk_get_micro_version()

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// GetMinorVersion returns the minor version number of the GTK library.
//
// For example, in GTK version 3.1.5 this is 1.
//
// This function is in the library, so it represents the GTK library your code
// is are running against. Contrast with the GTK_MINOR_VERSION macro, which
// represents the minor version of the GTK headers you have included when
// compiling your code.
func GetMinorVersion() uint {
	var _cret C.guint // in

	_cret = C.gtk_get_minor_version()

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}
