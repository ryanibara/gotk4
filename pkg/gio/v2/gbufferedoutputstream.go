// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
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
		{T: externglib.Type(C.g_buffered_output_stream_get_type()), F: marshalBufferedOutputStream},
	})
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
type BufferedOutputStream interface {
	FilterOutputStream
	Seekable

	// AutoGrow checks if the buffer automatically grows as data is added.
	AutoGrow() bool
	// BufferSize gets the size of the buffer in the @stream.
	BufferSize() uint
	// SetAutoGrow sets whether or not the @stream's buffer should automatically
	// grow. If @auto_grow is true, then each write will just make the buffer
	// larger, and you must manually flush the buffer to actually write out the
	// data to the underlying stream.
	SetAutoGrow(autoGrow bool)
	// SetBufferSize sets the size of the internal buffer to @size.
	SetBufferSize(size uint)
}

// bufferedOutputStream implements the BufferedOutputStream interface.
type bufferedOutputStream struct {
	FilterOutputStream
	Seekable
}

var _ BufferedOutputStream = (*bufferedOutputStream)(nil)

// WrapBufferedOutputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapBufferedOutputStream(obj *externglib.Object) BufferedOutputStream {
	return BufferedOutputStream{
		FilterOutputStream: WrapFilterOutputStream(obj),
		Seekable:           WrapSeekable(obj),
	}
}

func marshalBufferedOutputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBufferedOutputStream(obj), nil
}

// NewBufferedOutputStream constructs a class BufferedOutputStream.
func NewBufferedOutputStream(baseStream OutputStream) BufferedOutputStream {
	var arg1 *C.GOutputStream

	arg1 = (*C.GOutputStream)(unsafe.Pointer(baseStream.Native()))

	var cret C.GBufferedOutputStream
	var ret1 BufferedOutputStream

	cret = C.g_buffered_output_stream_new(baseStream)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(BufferedOutputStream)

	return ret1
}

// NewBufferedOutputStreamSized constructs a class BufferedOutputStream.
func NewBufferedOutputStreamSized(baseStream OutputStream, size uint) BufferedOutputStream {
	var arg1 *C.GOutputStream
	var arg2 C.gsize

	arg1 = (*C.GOutputStream)(unsafe.Pointer(baseStream.Native()))
	arg2 = C.gsize(size)

	var cret C.GBufferedOutputStream
	var ret1 BufferedOutputStream

	cret = C.g_buffered_output_stream_new_sized(baseStream, size)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(BufferedOutputStream)

	return ret1
}

// AutoGrow checks if the buffer automatically grows as data is added.
func (s bufferedOutputStream) AutoGrow() bool {
	var arg0 *C.GBufferedOutputStream

	arg0 = (*C.GBufferedOutputStream)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.g_buffered_output_stream_get_auto_grow(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// BufferSize gets the size of the buffer in the @stream.
func (s bufferedOutputStream) BufferSize() uint {
	var arg0 *C.GBufferedOutputStream

	arg0 = (*C.GBufferedOutputStream)(unsafe.Pointer(s.Native()))

	var cret C.gsize
	var ret1 uint

	cret = C.g_buffered_output_stream_get_buffer_size(arg0)

	ret1 = C.gsize(cret)

	return ret1
}

// SetAutoGrow sets whether or not the @stream's buffer should automatically
// grow. If @auto_grow is true, then each write will just make the buffer
// larger, and you must manually flush the buffer to actually write out the
// data to the underlying stream.
func (s bufferedOutputStream) SetAutoGrow(autoGrow bool) {
	var arg0 *C.GBufferedOutputStream
	var arg1 C.gboolean

	arg0 = (*C.GBufferedOutputStream)(unsafe.Pointer(s.Native()))
	if autoGrow {
		arg1 = C.gboolean(1)
	}

	C.g_buffered_output_stream_set_auto_grow(arg0, autoGrow)
}

// SetBufferSize sets the size of the internal buffer to @size.
func (s bufferedOutputStream) SetBufferSize(size uint) {
	var arg0 *C.GBufferedOutputStream
	var arg1 C.gsize

	arg0 = (*C.GBufferedOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = C.gsize(size)

	C.g_buffered_output_stream_set_buffer_size(arg0, size)
}

type BufferedOutputStreamPrivate struct {
	native C.GBufferedOutputStreamPrivate
}

// WrapBufferedOutputStreamPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBufferedOutputStreamPrivate(ptr unsafe.Pointer) *BufferedOutputStreamPrivate {
	if ptr == nil {
		return nil
	}

	return (*BufferedOutputStreamPrivate)(ptr)
}

func marshalBufferedOutputStreamPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapBufferedOutputStreamPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (b *BufferedOutputStreamPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}
