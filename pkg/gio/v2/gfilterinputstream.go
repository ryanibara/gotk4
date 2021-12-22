// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_filter_input_stream_get_type()), F: marshalFilterInputStreamer},
	})
}

// FilterInputStream: base class for input stream implementations that perform
// some kind of filtering operation on a base stream. Typical examples of
// filtering operations are character set conversion, compression and byte order
// flipping.
type FilterInputStream struct {
	_ [0]func() // equal guard
	InputStream
}

var (
	_ InputStreamer = (*FilterInputStream)(nil)
)

// FilterInputStreamer describes types inherited from class FilterInputStream.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type FilterInputStreamer interface {
	externglib.Objector
	baseFilterInputStream() *FilterInputStream
}

var _ FilterInputStreamer = (*FilterInputStream)(nil)

func wrapFilterInputStream(obj *externglib.Object) *FilterInputStream {
	return &FilterInputStream{
		InputStream: InputStream{
			Object: obj,
		},
	}
}

func marshalFilterInputStreamer(p uintptr) (interface{}, error) {
	return wrapFilterInputStream(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (stream *FilterInputStream) baseFilterInputStream() *FilterInputStream {
	return stream
}

// BaseFilterInputStream returns the underlying base object.
func BaseFilterInputStream(obj FilterInputStreamer) *FilterInputStream {
	return obj.baseFilterInputStream()
}

// BaseStream gets the base stream for the filter stream.
//
// The function returns the following values:
//
//    - inputStream: Stream.
//
func (stream *FilterInputStream) BaseStream() InputStreamer {
	var _arg0 *C.GFilterInputStream // out
	var _cret *C.GInputStream       // in

	_arg0 = (*C.GFilterInputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_filter_input_stream_get_base_stream(_arg0)
	runtime.KeepAlive(stream)

	var _inputStream InputStreamer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.InputStreamer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(InputStreamer)
			return ok
		})
		rv, ok := casted.(InputStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.InputStreamer")
		}
		_inputStream = rv
	}

	return _inputStream
}

// CloseBaseStream returns whether the base stream will be closed when stream is
// closed.
//
// The function returns the following values:
//
//    - ok: TRUE if the base stream will be closed.
//
func (stream *FilterInputStream) CloseBaseStream() bool {
	var _arg0 *C.GFilterInputStream // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GFilterInputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_filter_input_stream_get_close_base_stream(_arg0)
	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetCloseBaseStream sets whether the base stream will be closed when stream is
// closed.
//
// The function takes the following parameters:
//
//    - closeBase: TRUE to close the base stream.
//
func (stream *FilterInputStream) SetCloseBaseStream(closeBase bool) {
	var _arg0 *C.GFilterInputStream // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GFilterInputStream)(unsafe.Pointer(stream.Native()))
	if closeBase {
		_arg1 = C.TRUE
	}

	C.g_filter_input_stream_set_close_base_stream(_arg0, _arg1)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(closeBase)
}
