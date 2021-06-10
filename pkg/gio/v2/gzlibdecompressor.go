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
		{T: externglib.Type(C.g_zlib_decompressor_get_type()), F: marshalZlibDecompressor},
	})
}

// ZlibDecompressor: zlib decompression
type ZlibDecompressor interface {
	gextras.Objector
	Converter
}

// zlibDecompressor implements the ZlibDecompressor interface.
type zlibDecompressor struct {
	gextras.Objector
	Converter
}

var _ ZlibDecompressor = (*zlibDecompressor)(nil)

// WrapZlibDecompressor wraps a GObject to the right type. It is
// primarily used internally.
func WrapZlibDecompressor(obj *externglib.Object) ZlibDecompressor {
	return ZlibDecompressor{
		Objector:  obj,
		Converter: WrapConverter(obj),
	}
}

func marshalZlibDecompressor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapZlibDecompressor(obj), nil
}
