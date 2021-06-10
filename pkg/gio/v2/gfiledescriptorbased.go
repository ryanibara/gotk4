// Code generated by girgen. DO NOT EDIT.

package gio

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
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

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_file_descriptor_based_get_type()), F: marshalFileDescriptorBased},
	})
}

// FileDescriptorBasedOverrider contains methods that are overridable. This
// interface is a subset of the interface FileDescriptorBased.
type FileDescriptorBasedOverrider interface {
	// Fd gets the underlying file descriptor.
	Fd() int
}

// FileDescriptorBased is implemented by streams (implementations of Stream or
// Stream) that are based on file descriptors.
//
// Note that `<gio/gfiledescriptorbased.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file when
// using it.
type FileDescriptorBased interface {
	gextras.Objector
	FileDescriptorBasedOverrider
}

// fileDescriptorBased implements the FileDescriptorBased interface.
type fileDescriptorBased struct {
	gextras.Objector
}

var _ FileDescriptorBased = (*fileDescriptorBased)(nil)

// WrapFileDescriptorBased wraps a GObject to a type that implements interface
// FileDescriptorBased. It is primarily used internally.
func WrapFileDescriptorBased(obj *externglib.Object) FileDescriptorBased {
	return FileDescriptorBased{
		Objector: obj,
	}
}

func marshalFileDescriptorBased(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileDescriptorBased(obj), nil
}

// Fd gets the underlying file descriptor.
func (f fileDescriptorBased) Fd() int {
	var _arg0 *C.GFileDescriptorBased

	_arg0 = (*C.GFileDescriptorBased)(unsafe.Pointer(f.Native()))

	var _cret C.int

	_cret = C.g_file_descriptor_based_get_fd(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}
