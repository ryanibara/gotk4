// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0
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
		{T: externglib.Type(C.g_memory_output_stream_get_type()), F: marshalMemoryOutputStream},
	})
}

// MemoryOutputStream is a class for using arbitrary memory chunks as output for
// GIO streaming output operations.
//
// As of GLib 2.34, OutputStream trivially implements OutputStream: it always
// polls as ready.
type MemoryOutputStream interface {
	OutputStream
	PollableOutputStream
	Seekable

	// Data gets any loaded data from the @ostream.
	//
	// Note that the returned pointer may become invalid on the next write or
	// truncate operation on the stream.
	Data() interface{}
	// DataSize returns the number of bytes from the start up to including the
	// last byte written in the stream that has not been truncated away.
	DataSize() uint
	// Size gets the size of the currently allocated data area (available from
	// g_memory_output_stream_get_data()).
	//
	// You probably don't want to use this function on resizable streams. See
	// g_memory_output_stream_get_data_size() instead. For resizable streams the
	// size returned by this function is an implementation detail and may be
	// change at any time in response to operations on the stream.
	//
	// If the stream is fixed-sized (ie: no realloc was passed to
	// g_memory_output_stream_new()) then this is the maximum size of the stream
	// and further writes will return G_IO_ERROR_NO_SPACE.
	//
	// In any case, if you want the number of bytes currently written to the
	// stream, use g_memory_output_stream_get_data_size().
	Size() uint
	// StealAsBytes returns data from the @ostream as a #GBytes. @ostream must
	// be closed before calling this function.
	StealAsBytes() *glib.Bytes
	// StealData gets any loaded data from the @ostream. Ownership of the data
	// is transferred to the caller; when no longer needed it must be freed
	// using the free function set in @ostream's OutputStream:destroy-function
	// property.
	//
	// @ostream must be closed before calling this function.
	StealData() interface{}
}

// memoryOutputStream implements the MemoryOutputStream interface.
type memoryOutputStream struct {
	OutputStream
	PollableOutputStream
	Seekable
}

var _ MemoryOutputStream = (*memoryOutputStream)(nil)

// WrapMemoryOutputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapMemoryOutputStream(obj *externglib.Object) MemoryOutputStream {
	return MemoryOutputStream{
		OutputStream:         WrapOutputStream(obj),
		PollableOutputStream: WrapPollableOutputStream(obj),
		Seekable:             WrapSeekable(obj),
	}
}

func marshalMemoryOutputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMemoryOutputStream(obj), nil
}

// NewMemoryOutputStream constructs a class MemoryOutputStream.
func NewMemoryOutputStream(data interface{}, size uint, reallocFunction ReallocFunc) MemoryOutputStream {
	var arg1 C.gpointer
	var arg2 C.gsize
	var arg3 C.GReallocFunc
	var arg4 C.GDestroyNotify

	arg1 = C.gpointer(box.Assign(data))
	arg2 = C.gsize(size)

	ret := C.g_memory_output_stream_new(arg1, arg2, arg3, arg4)

	var ret0 MemoryOutputStream

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(MemoryOutputStream)

	return ret0
}

// NewMemoryOutputStreamResizable constructs a class MemoryOutputStream.
func NewMemoryOutputStreamResizable() MemoryOutputStream {
	ret := C.g_memory_output_stream_new_resizable()

	var ret0 MemoryOutputStream

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(MemoryOutputStream)

	return ret0
}

// Data gets any loaded data from the @ostream.
//
// Note that the returned pointer may become invalid on the next write or
// truncate operation on the stream.
func (o memoryOutputStream) Data() interface{} {
	var arg0 *C.GMemoryOutputStream

	arg0 = (*C.GMemoryOutputStream)(o.Native())

	ret := C.g_memory_output_stream_get_data(arg0)

	var ret0 interface{}

	ret0 = box.Get(uintptr(ret)).(interface{})

	return ret0
}

// DataSize returns the number of bytes from the start up to including the
// last byte written in the stream that has not been truncated away.
func (o memoryOutputStream) DataSize() uint {
	var arg0 *C.GMemoryOutputStream

	arg0 = (*C.GMemoryOutputStream)(o.Native())

	ret := C.g_memory_output_stream_get_data_size(arg0)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// Size gets the size of the currently allocated data area (available from
// g_memory_output_stream_get_data()).
//
// You probably don't want to use this function on resizable streams. See
// g_memory_output_stream_get_data_size() instead. For resizable streams the
// size returned by this function is an implementation detail and may be
// change at any time in response to operations on the stream.
//
// If the stream is fixed-sized (ie: no realloc was passed to
// g_memory_output_stream_new()) then this is the maximum size of the stream
// and further writes will return G_IO_ERROR_NO_SPACE.
//
// In any case, if you want the number of bytes currently written to the
// stream, use g_memory_output_stream_get_data_size().
func (o memoryOutputStream) Size() uint {
	var arg0 *C.GMemoryOutputStream

	arg0 = (*C.GMemoryOutputStream)(o.Native())

	ret := C.g_memory_output_stream_get_size(arg0)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// StealAsBytes returns data from the @ostream as a #GBytes. @ostream must
// be closed before calling this function.
func (o memoryOutputStream) StealAsBytes() *glib.Bytes {
	var arg0 *C.GMemoryOutputStream

	arg0 = (*C.GMemoryOutputStream)(o.Native())

	ret := C.g_memory_output_stream_steal_as_bytes(arg0)

	var ret0 *glib.Bytes

	{
		ret0 = glib.WrapBytes(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *glib.Bytes) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// StealData gets any loaded data from the @ostream. Ownership of the data
// is transferred to the caller; when no longer needed it must be freed
// using the free function set in @ostream's OutputStream:destroy-function
// property.
//
// @ostream must be closed before calling this function.
func (o memoryOutputStream) StealData() interface{} {
	var arg0 *C.GMemoryOutputStream

	arg0 = (*C.GMemoryOutputStream)(o.Native())

	ret := C.g_memory_output_stream_steal_data(arg0)

	var ret0 interface{}

	ret0 = box.Get(uintptr(ret)).(interface{})

	return ret0
}

type MemoryOutputStreamPrivate struct {
	native C.GMemoryOutputStreamPrivate
}

// WrapMemoryOutputStreamPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapMemoryOutputStreamPrivate(ptr unsafe.Pointer) *MemoryOutputStreamPrivate {
	if ptr == nil {
		return nil
	}

	return (*MemoryOutputStreamPrivate)(ptr)
}

func marshalMemoryOutputStreamPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapMemoryOutputStreamPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (m *MemoryOutputStreamPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&m.native)
}
