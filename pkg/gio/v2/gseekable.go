// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_seekable_get_type()), F: marshalSeekabler},
	})
}

// SeekableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SeekableOverrider interface {
	// CanSeek tests if the stream supports the Iface.
	CanSeek() bool
	// CanTruncate tests if the length of the stream can be adjusted with
	// g_seekable_truncate().
	CanTruncate() bool
	// Seek seeks in the stream by the given offset, modified by type.
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
	// If cancellable is not NULL, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	Seek(ctx context.Context, offset int64, typ glib.SeekType) error
	// Tell tells the current position within the stream.
	Tell() int64
	// TruncateFn sets the length of the stream to offset. If the stream was
	// previously larger than offset, the extra data is discarded. If the stream
	// was previously shorter than offset, it is extended with NUL ('\0') bytes.
	//
	// If cancellable is not NULL, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	TruncateFn(ctx context.Context, offset int64) error
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
type Seekable struct {
	*externglib.Object
}

// Seekabler describes Seekable's abstract methods.
type Seekabler interface {
	externglib.Objector

	// CanSeek tests if the stream supports the Iface.
	CanSeek() bool
	// CanTruncate tests if the length of the stream can be adjusted with
	// g_seekable_truncate().
	CanTruncate() bool
	// Seek seeks in the stream by the given offset, modified by type.
	Seek(ctx context.Context, offset int64, typ glib.SeekType) error
	// Tell tells the current position within the stream.
	Tell() int64
	// Truncate sets the length of the stream to offset.
	Truncate(ctx context.Context, offset int64) error
}

var _ Seekabler = (*Seekable)(nil)

func wrapSeekable(obj *externglib.Object) *Seekable {
	return &Seekable{
		Object: obj,
	}
}

func marshalSeekabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSeekable(obj), nil
}

// CanSeek tests if the stream supports the Iface.
func (seekable *Seekable) CanSeek() bool {
	var _arg0 *C.GSeekable // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSeekable)(unsafe.Pointer(seekable.Native()))

	_cret = C.g_seekable_can_seek(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanTruncate tests if the length of the stream can be adjusted with
// g_seekable_truncate().
func (seekable *Seekable) CanTruncate() bool {
	var _arg0 *C.GSeekable // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSeekable)(unsafe.Pointer(seekable.Native()))

	_cret = C.g_seekable_can_truncate(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Seek seeks in the stream by the given offset, modified by type.
//
// Attempting to seek past the end of the stream will have different results
// depending on if the stream is fixed-sized or resizable. If the stream is
// resizable then seeking past the end and then writing will result in zeros
// filling the empty space. Seeking past the end of a resizable stream and
// reading will result in EOF. Seeking past the end of a fixed-sized stream will
// fail.
//
// Any operation that would result in a negative offset will fail.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned.
func (seekable *Seekable) Seek(ctx context.Context, offset int64, typ glib.SeekType) error {
	var _arg0 *C.GSeekable    // out
	var _arg3 *C.GCancellable // out
	var _arg1 C.goffset       // out
	var _arg2 C.GSeekType     // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GSeekable)(unsafe.Pointer(seekable.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.goffset(offset)
	_arg2 = C.GSeekType(typ)

	C.g_seekable_seek(_arg0, _arg1, _arg2, _arg3, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Tell tells the current position within the stream.
func (seekable *Seekable) Tell() int64 {
	var _arg0 *C.GSeekable // out
	var _cret C.goffset    // in

	_arg0 = (*C.GSeekable)(unsafe.Pointer(seekable.Native()))

	_cret = C.g_seekable_tell(_arg0)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// Truncate sets the length of the stream to offset. If the stream was
// previously larger than offset, the extra data is discarded. If the stream was
// previously shorter than offset, it is extended with NUL ('\0') bytes.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned. If an operation was
// partially finished when the operation was cancelled the partial result will
// be returned, without an error.
func (seekable *Seekable) Truncate(ctx context.Context, offset int64) error {
	var _arg0 *C.GSeekable    // out
	var _arg2 *C.GCancellable // out
	var _arg1 C.goffset       // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GSeekable)(unsafe.Pointer(seekable.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.goffset(offset)

	C.g_seekable_truncate(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
