// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdbool.h>
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
		{T: externglib.Type(C.g_unix_output_stream_get_type()), F: marshalUnixOutputStream},
	})
}

type UnixOutputStreamPrivate struct {
	native C.GUnixOutputStreamPrivate
}

// WrapUnixOutputStreamPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapUnixOutputStreamPrivate(ptr unsafe.Pointer) *UnixOutputStreamPrivate {
	if ptr == nil {
		return nil
	}

	return (*UnixOutputStreamPrivate)(ptr)
}

func marshalUnixOutputStreamPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapUnixOutputStreamPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (u *UnixOutputStreamPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&u.native)
}

// UnixOutputStream implements Stream for writing to a UNIX file descriptor,
// including asynchronous operations. (If the file descriptor refers to a socket
// or pipe, this will use poll() to do asynchronous I/O. If it refers to a
// regular file, it will fall back to doing asynchronous I/O in another thread.)
//
// Note that `<gio/gunixoutputstream.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file when
// using it.
type UnixOutputStream interface {
	OutputStream
	FileDescriptorBased
	PollableOutputStream

	// CloseFd returns whether the file descriptor of @stream will be closed
	// when the stream is closed.
	CloseFd() bool
	// Fd: return the UNIX file descriptor that the stream writes to.
	Fd() int
	// SetCloseFd sets whether the file descriptor of @stream shall be closed
	// when the stream is closed.
	SetCloseFd(closeFd bool)
}

// unixOutputStream implements the UnixOutputStream interface.
type unixOutputStream struct {
	OutputStream
	FileDescriptorBased
	PollableOutputStream
}

var _ UnixOutputStream = (*unixOutputStream)(nil)

// WrapUnixOutputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapUnixOutputStream(obj *externglib.Object) UnixOutputStream {
	return UnixOutputStream{
		OutputStream:         WrapOutputStream(obj),
		FileDescriptorBased:  WrapFileDescriptorBased(obj),
		PollableOutputStream: WrapPollableOutputStream(obj),
	}
}

func marshalUnixOutputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapUnixOutputStream(obj), nil
}

// NewUnixOutputStream constructs a class UnixOutputStream.
func NewUnixOutputStream(fd int, closeFd bool) UnixOutputStream {
	var arg1 C.gint
	var arg2 C.gboolean

	arg1 = C.gint(fd)
	if closeFd {
		arg2 = C.TRUE
	}

	ret := C.g_unix_output_stream_new(arg1, arg2)

	var ret0 UnixOutputStream

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(UnixOutputStream)

	return ret0
}

// CloseFd returns whether the file descriptor of @stream will be closed
// when the stream is closed.
func (s unixOutputStream) CloseFd() bool {
	var arg0 *C.GUnixOutputStream

	arg0 = (*C.GUnixOutputStream)(s.Native())

	ret := C.g_unix_output_stream_get_close_fd(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Fd: return the UNIX file descriptor that the stream writes to.
func (s unixOutputStream) Fd() int {
	var arg0 *C.GUnixOutputStream

	arg0 = (*C.GUnixOutputStream)(s.Native())

	ret := C.g_unix_output_stream_get_fd(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// SetCloseFd sets whether the file descriptor of @stream shall be closed
// when the stream is closed.
func (s unixOutputStream) SetCloseFd(closeFd bool) {
	var arg0 *C.GUnixOutputStream
	var arg1 C.gboolean

	arg0 = (*C.GUnixOutputStream)(s.Native())
	if closeFd {
		arg1 = C.TRUE
	}

	C.g_unix_output_stream_set_close_fd(arg0, arg1)
}
