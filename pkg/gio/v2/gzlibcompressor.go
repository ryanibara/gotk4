// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gzlibcompressor.go.
var GTypeZlibCompressor = externglib.Type(C.g_zlib_compressor_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeZlibCompressor, F: marshalZlibCompressor},
	})
}

// ZlibCompressorOverrider contains methods that are overridable.
type ZlibCompressorOverrider interface {
}

// ZlibCompressor: zlib decompression.
type ZlibCompressor struct {
	_ [0]func() // equal guard
	*externglib.Object

	Converter
}

var (
	_ externglib.Objector = (*ZlibCompressor)(nil)
)

func classInitZlibCompressorrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapZlibCompressor(obj *externglib.Object) *ZlibCompressor {
	return &ZlibCompressor{
		Object: obj,
		Converter: Converter{
			Object: obj,
		},
	}
}

func marshalZlibCompressor(p uintptr) (interface{}, error) {
	return wrapZlibCompressor(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewZlibCompressor creates a new Compressor.
//
// The function takes the following parameters:
//
//    - format to use for the compressed data.
//    - level: compression level (0-9), -1 for default.
//
// The function returns the following values:
//
//    - zlibCompressor: new Compressor.
//
func NewZlibCompressor(format ZlibCompressorFormat, level int) *ZlibCompressor {
	var _arg1 C.GZlibCompressorFormat // out
	var _arg2 C.int                   // out
	var _cret *C.GZlibCompressor      // in

	_arg1 = C.GZlibCompressorFormat(format)
	_arg2 = C.int(level)

	_cret = C.g_zlib_compressor_new(_arg1, _arg2)
	runtime.KeepAlive(format)
	runtime.KeepAlive(level)

	var _zlibCompressor *ZlibCompressor // out

	_zlibCompressor = wrapZlibCompressor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _zlibCompressor
}

// FileInfo returns the Compressor:file-info property.
//
// The function returns the following values:
//
//    - fileInfo (optional) or NULL.
//
func (compressor *ZlibCompressor) FileInfo() *FileInfo {
	var _arg0 *C.GZlibCompressor // out
	var _cret *C.GFileInfo       // in

	_arg0 = (*C.GZlibCompressor)(unsafe.Pointer(compressor.Native()))

	_cret = C.g_zlib_compressor_get_file_info(_arg0)
	runtime.KeepAlive(compressor)

	var _fileInfo *FileInfo // out

	if _cret != nil {
		_fileInfo = wrapFileInfo(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _fileInfo
}

// SetFileInfo sets file_info in compressor. If non-NULL, and compressor's
// Compressor:format property is G_ZLIB_COMPRESSOR_FORMAT_GZIP, it will be used
// to set the file name and modification time in the GZIP header of the
// compressed data.
//
// Note: it is an error to call this function while a compression is in
// progress; it may only be called immediately after creation of compressor, or
// after resetting it with g_converter_reset().
//
// The function takes the following parameters:
//
//    - fileInfo (optional): Info.
//
func (compressor *ZlibCompressor) SetFileInfo(fileInfo *FileInfo) {
	var _arg0 *C.GZlibCompressor // out
	var _arg1 *C.GFileInfo       // out

	_arg0 = (*C.GZlibCompressor)(unsafe.Pointer(compressor.Native()))
	if fileInfo != nil {
		_arg1 = (*C.GFileInfo)(unsafe.Pointer(fileInfo.Native()))
	}

	C.g_zlib_compressor_set_file_info(_arg0, _arg1)
	runtime.KeepAlive(compressor)
	runtime.KeepAlive(fileInfo)
}
