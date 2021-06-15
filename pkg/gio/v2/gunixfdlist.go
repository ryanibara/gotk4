// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 glib-2.0 gobject-introspection-1.0
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
		{T: externglib.Type(C.g_unix_fd_list_get_type()), F: marshalUnixFDList},
	})
}

// UnixFDList: a FDList contains a list of file descriptors. It owns the file
// descriptors that it contains, closing them when finalized.
//
// It may be wrapped in a FDMessage and sent over a #GSocket in the
// G_SOCKET_FAMILY_UNIX family by using g_socket_send_message() and received
// using g_socket_receive_message().
//
// Note that `<gio/gunixfdlist.h>` belongs to the UNIX-specific GIO interfaces,
// thus you have to use the `gio-unix-2.0.pc` pkg-config file when using it.
type UnixFDList interface {
	gextras.Objector

	// Append adds a file descriptor to @list.
	//
	// The file descriptor is duplicated using dup(). You keep your copy of the
	// descriptor and the copy contained in @list will be closed when @list is
	// finalized.
	//
	// A possible cause of failure is exceeding the per-process or system-wide
	// file descriptor limit.
	//
	// The index of the file descriptor in the list is returned. If you use this
	// index with g_unix_fd_list_get() then you will receive back a duplicated
	// copy of the same file descriptor.
	Append(fd int) (int, error)
	// Get gets a file descriptor out of @list.
	//
	// @index_ specifies the index of the file descriptor to get. It is a
	// programmer error for @index_ to be out of range; see
	// g_unix_fd_list_get_length().
	//
	// The file descriptor is duplicated using dup() and set as close-on-exec
	// before being returned. You must call close() on it when you are done.
	//
	// A possible cause of failure is exceeding the per-process or system-wide
	// file descriptor limit.
	Get(index_ int) (int, error)
	// Length gets the length of @list (ie: the number of file descriptors
	// contained within).
	Length() int
}

// unixFDList implements the UnixFDList class.
type unixFDList struct {
	gextras.Objector
}

var _ UnixFDList = (*unixFDList)(nil)

// WrapUnixFDList wraps a GObject to the right type. It is
// primarily used internally.
func WrapUnixFDList(obj *externglib.Object) UnixFDList {
	return unixFDList{
		Objector: obj,
	}
}

func marshalUnixFDList(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapUnixFDList(obj), nil
}

// NewUnixFDList constructs a class UnixFDList.
func NewUnixFDList() UnixFDList {
	var _cret C.GUnixFDList // in

	_cret = C.g_unix_fd_list_new()

	var _unixFDList UnixFDList // out

	_unixFDList = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(UnixFDList)

	return _unixFDList
}

// NewUnixFDListFromArray constructs a class UnixFDList.
func NewUnixFDListFromArray(fds []int) UnixFDList {
	var _arg1 *C.gint
	var _arg2 C.gint
	var _cret C.GUnixFDList // in

	_arg2 = C.gint(len(fds))
	_arg1 = (*C.gint)(unsafe.Pointer(&fds[0]))

	_cret = C.g_unix_fd_list_new_from_array(_arg1, _arg2)

	var _unixFDList UnixFDList // out

	_unixFDList = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(UnixFDList)

	return _unixFDList
}

// Append adds a file descriptor to @list.
//
// The file descriptor is duplicated using dup(). You keep your copy of the
// descriptor and the copy contained in @list will be closed when @list is
// finalized.
//
// A possible cause of failure is exceeding the per-process or system-wide
// file descriptor limit.
//
// The index of the file descriptor in the list is returned. If you use this
// index with g_unix_fd_list_get() then you will receive back a duplicated
// copy of the same file descriptor.
func (l unixFDList) Append(fd int) (int, error) {
	var _arg0 *C.GUnixFDList // out
	var _arg1 C.gint         // out
	var _cret C.gint         // in
	var _cerr *C.GError      // in

	_arg0 = (*C.GUnixFDList)(unsafe.Pointer(l.Native()))
	_arg1 = (C.gint)(fd)

	_cret = C.g_unix_fd_list_append(_arg0, _arg1, &_cerr)

	var _gint int    // out
	var _goerr error // out

	_gint = (int)(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gint, _goerr
}

// Get gets a file descriptor out of @list.
//
// @index_ specifies the index of the file descriptor to get. It is a
// programmer error for @index_ to be out of range; see
// g_unix_fd_list_get_length().
//
// The file descriptor is duplicated using dup() and set as close-on-exec
// before being returned. You must call close() on it when you are done.
//
// A possible cause of failure is exceeding the per-process or system-wide
// file descriptor limit.
func (l unixFDList) Get(index_ int) (int, error) {
	var _arg0 *C.GUnixFDList // out
	var _arg1 C.gint         // out
	var _cret C.gint         // in
	var _cerr *C.GError      // in

	_arg0 = (*C.GUnixFDList)(unsafe.Pointer(l.Native()))
	_arg1 = (C.gint)(index_)

	_cret = C.g_unix_fd_list_get(_arg0, _arg1, &_cerr)

	var _gint int    // out
	var _goerr error // out

	_gint = (int)(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gint, _goerr
}

// Length gets the length of @list (ie: the number of file descriptors
// contained within).
func (l unixFDList) Length() int {
	var _arg0 *C.GUnixFDList // out
	var _cret C.gint         // in

	_arg0 = (*C.GUnixFDList)(unsafe.Pointer(l.Native()))

	_cret = C.g_unix_fd_list_get_length(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}
