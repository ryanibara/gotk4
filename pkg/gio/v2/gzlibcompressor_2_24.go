// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

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

	_zlibCompressor = wrapZlibCompressor(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _zlibCompressor
}
