// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0
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
		{T: externglib.Type(C.g_unix_input_stream_get_type()), F: marshalUnixInputStream},
	})
}

// UnixInputStream implements Stream for reading from a UNIX file descriptor,
// including asynchronous operations. (If the file descriptor refers to a socket
// or pipe, this will use poll() to do asynchronous I/O. If it refers to a
// regular file, it will fall back to doing asynchronous I/O in another thread.)
//
// Note that `<gio/gunixinputstream.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file when
// using it.
type UnixInputStream interface {
	InputStream
	FileDescriptorBased
	PollableInputStream

	// CloseFd returns whether the file descriptor of @stream will be closed
	// when the stream is closed.
	CloseFd() bool
	// Fd: return the UNIX file descriptor that the stream reads from.
	Fd() int
	// SetCloseFd sets whether the file descriptor of @stream shall be closed
	// when the stream is closed.
	SetCloseFd(closeFd bool)
}

// unixInputStream implements the UnixInputStream interface.
type unixInputStream struct {
	InputStream
	FileDescriptorBased
	PollableInputStream
}

var _ UnixInputStream = (*unixInputStream)(nil)

// WrapUnixInputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapUnixInputStream(obj *externglib.Object) UnixInputStream {
	return UnixInputStream{
		InputStream:         WrapInputStream(obj),
		FileDescriptorBased: WrapFileDescriptorBased(obj),
		PollableInputStream: WrapPollableInputStream(obj),
	}
}

func marshalUnixInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapUnixInputStream(obj), nil
}

// NewUnixInputStream constructs a class UnixInputStream.
func NewUnixInputStream(fd int, closeFd bool) UnixInputStream {
	var arg1 C.gint
	var arg2 C.gboolean

	arg1 = C.gint(fd)
	if closeFd {
		arg2 = C.TRUE
	}

	ret := C.g_unix_input_stream_new(arg1, arg2)

	var ret0 UnixInputStream

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(UnixInputStream)

	return ret0
}

// CloseFd returns whether the file descriptor of @stream will be closed
// when the stream is closed.
func (s unixInputStream) CloseFd() bool {
	var arg0 *C.GUnixInputStream

	arg0 = (*C.GUnixInputStream)(s.Native())

	ret := C.g_unix_input_stream_get_close_fd(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != C.false

	return ret0
}

// Fd: return the UNIX file descriptor that the stream reads from.
func (s unixInputStream) Fd() int {
	var arg0 *C.GUnixInputStream

	arg0 = (*C.GUnixInputStream)(s.Native())

	ret := C.g_unix_input_stream_get_fd(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// SetCloseFd sets whether the file descriptor of @stream shall be closed
// when the stream is closed.
func (s unixInputStream) SetCloseFd(closeFd bool) {
	var arg0 *C.GUnixInputStream
	var arg1 C.gboolean

	arg0 = (*C.GUnixInputStream)(s.Native())
	if closeFd {
		arg1 = C.TRUE
	}

	C.g_unix_input_stream_set_close_fd(arg0, arg1)
}

type UnixInputStreamPrivate struct {
	native C.GUnixInputStreamPrivate
}

// WrapUnixInputStreamPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapUnixInputStreamPrivate(ptr unsafe.Pointer) *UnixInputStreamPrivate {
	if ptr == nil {
		return nil
	}

	return (*UnixInputStreamPrivate)(ptr)
}

func marshalUnixInputStreamPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapUnixInputStreamPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (u *UnixInputStreamPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&u.native)
}
