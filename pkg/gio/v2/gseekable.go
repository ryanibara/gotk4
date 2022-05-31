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
// extern gboolean _gotk4_gio2_SeekableIface_can_seek(void*);
// extern gboolean _gotk4_gio2_SeekableIface_can_truncate(void*);
import "C"

// glib.Type values for gseekable.go.
var GTypeSeekable = coreglib.Type(C.g_seekable_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeSeekable, F: marshalSeekable},
	})
}

// SeekableOverrider contains methods that are overridable.
type SeekableOverrider interface {
	// CanSeek tests if the stream supports the Iface.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if seekable can be seeked. FALSE otherwise.
	//
	CanSeek() bool
	// CanTruncate tests if the length of the stream can be adjusted with
	// g_seekable_truncate().
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if the stream can be truncated, FALSE otherwise.
	//
	CanTruncate() bool
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
//
// Seekable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Seekable struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Seekable)(nil)
)

// Seekabler describes Seekable's interface methods.
type Seekabler interface {
	coreglib.Objector

	// CanSeek tests if the stream supports the Iface.
	CanSeek() bool
	// CanTruncate tests if the length of the stream can be adjusted with
	// g_seekable_truncate().
	CanTruncate() bool
}

var _ Seekabler = (*Seekable)(nil)

func ifaceInitSeekabler(gifacePtr, data C.gpointer) {
	iface := (*C.GSeekableIface)(unsafe.Pointer(gifacePtr))
	iface.can_seek = (*[0]byte)(C._gotk4_gio2_SeekableIface_can_seek)
	iface.can_truncate = (*[0]byte)(C._gotk4_gio2_SeekableIface_can_truncate)
}

//export _gotk4_gio2_SeekableIface_can_seek
func _gotk4_gio2_SeekableIface_can_seek(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SeekableOverrider)

	ok := iface.CanSeek()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_SeekableIface_can_truncate
func _gotk4_gio2_SeekableIface_can_truncate(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SeekableOverrider)

	ok := iface.CanTruncate()

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapSeekable(obj *coreglib.Object) *Seekable {
	return &Seekable{
		Object: obj,
	}
}

func marshalSeekable(p uintptr) (interface{}, error) {
	return wrapSeekable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CanSeek tests if the stream supports the Iface.
//
// The function returns the following values:
//
//    - ok: TRUE if seekable can be seeked. FALSE otherwise.
//
func (seekable *Seekable) CanSeek() bool {
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(seekable).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(seekable)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanTruncate tests if the length of the stream can be adjusted with
// g_seekable_truncate().
//
// The function returns the following values:
//
//    - ok: TRUE if the stream can be truncated, FALSE otherwise.
//
func (seekable *Seekable) CanTruncate() bool {
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(seekable).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(seekable)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
