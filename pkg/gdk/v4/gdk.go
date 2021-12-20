// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

// The function returns the following values:
//
func GLErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gdk_gl_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func VulkanErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gdk_vulkan_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}
