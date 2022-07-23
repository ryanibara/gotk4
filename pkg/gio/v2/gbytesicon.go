// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// NewBytesIcon creates a new icon for a bytes.
//
// This cannot fail, but loading and interpreting the bytes may fail later on
// (for example, if g_loadable_icon_load() is called) if the image is invalid.
//
// The function takes the following parameters:
//
//    - bytes: #GBytes.
//
// The function returns the following values:
//
//    - bytesIcon for the given bytes.
//
func NewBytesIcon(bytes *glib.Bytes) *BytesIcon {
	var _arg1 *C.GBytes // out
	var _cret *C.GIcon  // in

	_arg1 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(bytes)))

	_cret = C.g_bytes_icon_new(_arg1)
	runtime.KeepAlive(bytes)

	var _bytesIcon *BytesIcon // out

	_bytesIcon = wrapBytesIcon(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _bytesIcon
}

// Bytes gets the #GBytes associated with the given icon.
//
// The function returns the following values:
//
//    - bytes: #GBytes.
//
func (icon *BytesIcon) Bytes() *glib.Bytes {
	var _arg0 *C.GBytesIcon // out
	var _cret *C.GBytes     // in

	_arg0 = (*C.GBytesIcon)(unsafe.Pointer(coreglib.InternObject(icon).Native()))

	_cret = C.g_bytes_icon_get_bytes(_arg0)
	runtime.KeepAlive(icon)

	var _bytes *glib.Bytes // out

	_bytes = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_bytes_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	return _bytes
}
