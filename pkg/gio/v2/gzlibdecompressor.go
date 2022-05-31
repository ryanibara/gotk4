// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gzlibdecompressor.go.
var GTypeZlibDecompressor = coreglib.Type(C.g_zlib_decompressor_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeZlibDecompressor, F: marshalZlibDecompressor},
	})
}

// ZlibDecompressorOverrider contains methods that are overridable.
type ZlibDecompressorOverrider interface {
}

// ZlibDecompressor: zlib decompression.
type ZlibDecompressor struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Converter
}

var (
	_ coreglib.Objector = (*ZlibDecompressor)(nil)
)

func classInitZlibDecompressorrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapZlibDecompressor(obj *coreglib.Object) *ZlibDecompressor {
	return &ZlibDecompressor{
		Object: obj,
		Converter: Converter{
			Object: obj,
		},
	}
}

func marshalZlibDecompressor(p uintptr) (interface{}, error) {
	return wrapZlibDecompressor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// FileInfo retrieves the Info constructed from the GZIP header data of
// compressed data processed by compressor, or NULL if decompressor's
// Decompressor:format property is not G_ZLIB_COMPRESSOR_FORMAT_GZIP, or the
// header data was not fully processed yet, or it not present in the data stream
// at all.
//
// The function returns the following values:
//
//    - fileInfo (optional) or NULL.
//
func (decompressor *ZlibDecompressor) FileInfo() *FileInfo {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(decompressor).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "ZlibDecompressor").InvokeMethod("get_file_info", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(decompressor)

	var _fileInfo *FileInfo // out

	if _cret != nil {
		_fileInfo = wrapFileInfo(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _fileInfo
}
