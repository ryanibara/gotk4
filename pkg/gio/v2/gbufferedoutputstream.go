// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeBufferedOutputStream = coreglib.Type(C.g_buffered_output_stream_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeBufferedOutputStream, F: marshalBufferedOutputStream},
	})
}

// BufferedOutputStreamOverrider contains methods that are overridable.
type BufferedOutputStreamOverrider interface {
}

// BufferedOutputStream: buffered output stream implements OutputStream and
// provides for buffered writes.
//
// By default, OutputStream's buffer size is set at 4 kilobytes.
//
// To create a buffered output stream, use g_buffered_output_stream_new(), or
// g_buffered_output_stream_new_sized() to specify the buffer's size at
// construction.
//
// To get the size of a buffer within a buffered input stream, use
// g_buffered_output_stream_get_buffer_size(). To change the size of a buffered
// output stream's buffer, use g_buffered_output_stream_set_buffer_size(). Note
// that the buffer's size cannot be reduced below the size of the data within
// the buffer.
type BufferedOutputStream struct {
	_ [0]func() // equal guard
	FilterOutputStream

	Seekable
}

var (
	_ FilterOutputStreamer = (*BufferedOutputStream)(nil)
)

func initClassBufferedOutputStream(gclass unsafe.Pointer, goval any) {
}

func wrapBufferedOutputStream(obj *coreglib.Object) *BufferedOutputStream {
	return &BufferedOutputStream{
		FilterOutputStream: FilterOutputStream{
			OutputStream: OutputStream{
				Object: obj,
			},
		},
		Seekable: Seekable{
			Object: obj,
		},
	}
}

func marshalBufferedOutputStream(p uintptr) (interface{}, error) {
	return wrapBufferedOutputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBufferedOutputStream creates a new buffered output stream for a base
// stream.
//
// The function takes the following parameters:
//
//    - baseStream: Stream.
//
// The function returns the following values:
//
//    - bufferedOutputStream for the given base_stream.
//
func NewBufferedOutputStream(baseStream OutputStreamer) *BufferedOutputStream {
	var _arg1 *C.GOutputStream // out
	var _cret *C.GOutputStream // in

	_arg1 = (*C.GOutputStream)(unsafe.Pointer(coreglib.InternObject(baseStream).Native()))

	_cret = C.g_buffered_output_stream_new(_arg1)
	runtime.KeepAlive(baseStream)

	var _bufferedOutputStream *BufferedOutputStream // out

	_bufferedOutputStream = wrapBufferedOutputStream(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _bufferedOutputStream
}

// NewBufferedOutputStreamSized creates a new buffered output stream with a
// given buffer size.
//
// The function takes the following parameters:
//
//    - baseStream: Stream.
//    - size: #gsize.
//
// The function returns the following values:
//
//    - bufferedOutputStream with an internal buffer set to size.
//
func NewBufferedOutputStreamSized(baseStream OutputStreamer, size uint) *BufferedOutputStream {
	var _arg1 *C.GOutputStream // out
	var _arg2 C.gsize          // out
	var _cret *C.GOutputStream // in

	_arg1 = (*C.GOutputStream)(unsafe.Pointer(coreglib.InternObject(baseStream).Native()))
	_arg2 = C.gsize(size)

	_cret = C.g_buffered_output_stream_new_sized(_arg1, _arg2)
	runtime.KeepAlive(baseStream)
	runtime.KeepAlive(size)

	var _bufferedOutputStream *BufferedOutputStream // out

	_bufferedOutputStream = wrapBufferedOutputStream(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _bufferedOutputStream
}

// AutoGrow checks if the buffer automatically grows as data is added.
//
// The function returns the following values:
//
//    - ok: TRUE if the stream's buffer automatically grows, FALSE otherwise.
//
func (stream *BufferedOutputStream) AutoGrow() bool {
	var _arg0 *C.GBufferedOutputStream // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GBufferedOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_cret = C.g_buffered_output_stream_get_auto_grow(_arg0)
	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// BufferSize gets the size of the buffer in the stream.
//
// The function returns the following values:
//
//    - gsize: current size of the buffer.
//
func (stream *BufferedOutputStream) BufferSize() uint {
	var _arg0 *C.GBufferedOutputStream // out
	var _cret C.gsize                  // in

	_arg0 = (*C.GBufferedOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_cret = C.g_buffered_output_stream_get_buffer_size(_arg0)
	runtime.KeepAlive(stream)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// SetAutoGrow sets whether or not the stream's buffer should automatically
// grow. If auto_grow is true, then each write will just make the buffer larger,
// and you must manually flush the buffer to actually write out the data to the
// underlying stream.
//
// The function takes the following parameters:
//
//    - autoGrow: #gboolean.
//
func (stream *BufferedOutputStream) SetAutoGrow(autoGrow bool) {
	var _arg0 *C.GBufferedOutputStream // out
	var _arg1 C.gboolean               // out

	_arg0 = (*C.GBufferedOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	if autoGrow {
		_arg1 = C.TRUE
	}

	C.g_buffered_output_stream_set_auto_grow(_arg0, _arg1)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(autoGrow)
}

// SetBufferSize sets the size of the internal buffer to size.
//
// The function takes the following parameters:
//
//    - size: #gsize.
//
func (stream *BufferedOutputStream) SetBufferSize(size uint) {
	var _arg0 *C.GBufferedOutputStream // out
	var _arg1 C.gsize                  // out

	_arg0 = (*C.GBufferedOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	_arg1 = C.gsize(size)

	C.g_buffered_output_stream_set_buffer_size(_arg0, _arg1)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(size)
}
