// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_seekable_get_type()), F: marshalSeekable},
	})
}

// Seekable is implemented by streams (implementations of Stream or Stream) that
// support seeking.
//
// Seekable streams largely fall into two categories: resizable and fixed-size.
//
// #GSeekable on fixed-sized streams is approximately the same as POSIX lseek()
// on a block device (for example: attempting to seek past the end of the device
// is an error). Fixed streams typically cannot be truncated.
//
// #GSeekable on resizable streams is approximately the same as POSIX lseek() on
// a normal file. Seeking past the end and writing data will usually cause the
// stream to resize by introducing zero bytes.
type Seekable interface {
	gextras.Objector

	// CanSeek tests if the stream supports the Iface.
	CanSeek() bool
	// CanTruncate tests if the length of the stream can be adjusted with
	// g_seekable_truncate().
	CanTruncate() bool
	// Seek seeks in the stream by the given @offset, modified by @type.
	//
	// Attempting to seek past the end of the stream will have different results
	// depending on if the stream is fixed-sized or resizable. If the stream is
	// resizable then seeking past the end and then writing will result in zeros
	// filling the empty space. Seeking past the end of a resizable stream and
	// reading will result in EOF. Seeking past the end of a fixed-sized stream
	// will fail.
	//
	// Any operation that would result in a negative offset will fail.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	Seek(offset int64, typ glib.SeekType, cancellable Cancellable) error
	// Tell tells the current position within the stream.
	Tell() int64
	// Truncate sets the length of the stream to @offset. If the stream was
	// previously larger than @offset, the extra data is discarded. If the
	// stream was previously shorter than @offset, it is extended with NUL
	// ('\0') bytes.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	Truncate(offset int64, cancellable Cancellable) error
}

// seekable implements the Seekable interface.
type seekable struct {
	gextras.Objector
}

var _ Seekable = (*seekable)(nil)

// WrapSeekable wraps a GObject to a type that implements
// interface Seekable. It is primarily used internally.
func WrapSeekable(obj *externglib.Object) Seekable {
	return seekable{
		Objector: obj,
	}
}

func marshalSeekable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSeekable(obj), nil
}

func (s seekable) CanSeek() bool {
	var _arg0 *C.GSeekable // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSeekable)(unsafe.Pointer(s.Native()))

	_cret = C.g_seekable_can_seek(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s seekable) CanTruncate() bool {
	var _arg0 *C.GSeekable // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSeekable)(unsafe.Pointer(s.Native()))

	_cret = C.g_seekable_can_truncate(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s seekable) Seek(offset int64, typ glib.SeekType, cancellable Cancellable) error {
	var _arg0 *C.GSeekable    // out
	var _arg1 C.goffset       // out
	var _arg2 C.GSeekType     // out
	var _arg3 *C.GCancellable // out
	var _cerr **C.GError      // in

	_arg0 = (*C.GSeekable)(unsafe.Pointer(s.Native()))
	_arg1 = C.goffset(offset)
	_arg2 = C.GSeekType(typ)
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_seekable_seek(_arg0, _arg1, _arg2, _arg3, &_cerr)

	var _goerr error // out

	{
		var refTmpIn *C.GError
		var refTmpOut error

		refTmpIn = *_cerr

		refTmpOut = gerror.Take(unsafe.Pointer(refTmpIn))

		_goerr = refTmpOut
	}

	return _goerr
}

func (s seekable) Tell() int64 {
	var _arg0 *C.GSeekable // out
	var _cret C.goffset    // in

	_arg0 = (*C.GSeekable)(unsafe.Pointer(s.Native()))

	_cret = C.g_seekable_tell(_arg0)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

func (s seekable) Truncate(offset int64, cancellable Cancellable) error {
	var _arg0 *C.GSeekable    // out
	var _arg1 C.goffset       // out
	var _arg2 *C.GCancellable // out
	var _cerr **C.GError      // in

	_arg0 = (*C.GSeekable)(unsafe.Pointer(s.Native()))
	_arg1 = C.goffset(offset)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_seekable_truncate(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	{
		var refTmpIn *C.GError
		var refTmpOut error

		refTmpIn = *_cerr

		refTmpOut = gerror.Take(unsafe.Pointer(refTmpIn))

		_goerr = refTmpOut
	}

	return _goerr
}
