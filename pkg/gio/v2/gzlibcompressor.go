// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.g_zlib_compressor_get_type()), F: marshalZlibCompressor},
	})
}

// ZlibCompressor: zlib decompression
type ZlibCompressor interface {
	gextras.Objector
	Converter

	// FileInfo returns the Compressor:file-info property.
	FileInfo() FileInfo
	// SetFileInfo sets @file_info in @compressor. If non-nil, and @compressor's
	// Compressor:format property is G_ZLIB_COMPRESSOR_FORMAT_GZIP, it will be
	// used to set the file name and modification time in the GZIP header of the
	// compressed data.
	//
	// Note: it is an error to call this function while a compression is in
	// progress; it may only be called immediately after creation of
	// @compressor, or after resetting it with g_converter_reset().
	SetFileInfo(fileInfo FileInfo)
}

// zlibCompressor implements the ZlibCompressor interface.
type zlibCompressor struct {
	gextras.Objector
	Converter
}

var _ ZlibCompressor = (*zlibCompressor)(nil)

// WrapZlibCompressor wraps a GObject to the right type. It is
// primarily used internally.
func WrapZlibCompressor(obj *externglib.Object) ZlibCompressor {
	return ZlibCompressor{
		Objector:  obj,
		Converter: WrapConverter(obj),
	}
}

func marshalZlibCompressor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapZlibCompressor(obj), nil
}

// NewZlibCompressor constructs a class ZlibCompressor.
func NewZlibCompressor(format ZlibCompressorFormat, level int) ZlibCompressor {
	var arg1 C.GZlibCompressorFormat
	var arg2 C.int

	arg1 = (C.GZlibCompressorFormat)(format)
	arg2 = C.int(level)

	var cret C.GZlibCompressor
	var ret1 ZlibCompressor

	cret = C.g_zlib_compressor_new(format, level)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ZlibCompressor)

	return ret1
}

// FileInfo returns the Compressor:file-info property.
func (c zlibCompressor) FileInfo() FileInfo {
	var arg0 *C.GZlibCompressor

	arg0 = (*C.GZlibCompressor)(unsafe.Pointer(c.Native()))

	var cret *C.GFileInfo
	var ret1 FileInfo

	cret = C.g_zlib_compressor_get_file_info(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(FileInfo)

	return ret1
}

// SetFileInfo sets @file_info in @compressor. If non-nil, and @compressor's
// Compressor:format property is G_ZLIB_COMPRESSOR_FORMAT_GZIP, it will be
// used to set the file name and modification time in the GZIP header of the
// compressed data.
//
// Note: it is an error to call this function while a compression is in
// progress; it may only be called immediately after creation of
// @compressor, or after resetting it with g_converter_reset().
func (c zlibCompressor) SetFileInfo(fileInfo FileInfo) {
	var arg0 *C.GZlibCompressor
	var arg1 *C.GFileInfo

	arg0 = (*C.GZlibCompressor)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GFileInfo)(unsafe.Pointer(fileInfo.Native()))

	C.g_zlib_compressor_set_file_info(arg0, fileInfo)
}
