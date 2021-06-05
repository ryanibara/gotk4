// Code generated by girgen. DO NOT EDIT.

package glib

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// CheckVersion checks that the GLib library in use is compatible with the given
// version. Generally you would pass in the constants IB_MAJOR_VERSION,
// IB_MINOR_VERSION, IB_MICRO_VERSION as the three arguments to this function;
// that produces a check that the library in use is compatible with the version
// of GLib the application or module was compiled against.
//
// Compatibility is defined by two things: first the version of the running
// library is newer than the version
// @required_major.required_minor.@required_micro. Second the running library
// must be binary compatible with the version
// @required_major.required_minor.@required_micro (same major version.)
func CheckVersion(requiredMajor uint, requiredMinor uint, requiredMicro uint) string {
	var arg1 C.guint
	var arg2 C.guint
	var arg3 C.guint

	arg1 = C.guint(requiredMajor)
	arg2 = C.guint(requiredMinor)
	arg3 = C.guint(requiredMicro)

	var cret *C.gchar
	var ret1 string

	cret = C.glib_check_version(requiredMajor, requiredMinor, requiredMicro)

	ret1 = C.GoString(cret)

	return ret1
}
