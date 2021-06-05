// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
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
	Append(fd int) (gint int, err error)
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
	Get(index_ int) (gint int, err error)
	// Length gets the length of @list (ie: the number of file descriptors
	// contained within).
	Length() int
	// PeekFds returns the array of file descriptors that is contained in this
	// object.
	//
	// After this call, the descriptors remain the property of @list. The caller
	// must not close them and must not free the array. The array is valid only
	// until @list is changed in any way.
	//
	// If @length is non-nil then it is set to the number of file descriptors in
	// the returned array. The returned array is also terminated with -1.
	//
	// This function never returns nil. In case there are no file descriptors
	// contained in @list, an empty array is returned.
	PeekFds() (length int, gints []int)
	// StealFds returns the array of file descriptors that is contained in this
	// object.
	//
	// After this call, the descriptors are no longer contained in @list.
	// Further calls will return an empty list (unless more descriptors have
	// been added).
	//
	// The return result of this function must be freed with g_free(). The
	// caller is also responsible for closing all of the file descriptors. The
	// file descriptors in the array are set to close-on-exec.
	//
	// If @length is non-nil then it is set to the number of file descriptors in
	// the returned array. The returned array is also terminated with -1.
	//
	// This function never returns nil. In case there are no file descriptors
	// contained in @list, an empty array is returned.
	StealFds() (length int, gints []int)
}

// unixFDList implements the UnixFDList interface.
type unixFDList struct {
	gextras.Objector
}

var _ UnixFDList = (*unixFDList)(nil)

// WrapUnixFDList wraps a GObject to the right type. It is
// primarily used internally.
func WrapUnixFDList(obj *externglib.Object) UnixFDList {
	return UnixFDList{
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
	var cret C.GUnixFDList
	var ret1 UnixFDList

	cret = C.g_unix_fd_list_new()

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(UnixFDList)

	return ret1
}

// NewUnixFDListFromArray constructs a class UnixFDList.
func NewUnixFDListFromArray(fds []int) UnixFDList {

	var cret C.GUnixFDList
	var ret1 UnixFDList

	cret = C.g_unix_fd_list_new_from_array(fds, nFds)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(UnixFDList)

	return ret1
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
func (l unixFDList) Append(fd int) (gint int, err error) {
	var arg0 *C.GUnixFDList
	var arg1 C.gint

	arg0 = (*C.GUnixFDList)(unsafe.Pointer(l.Native()))
	arg1 = C.gint(fd)

	var errout *C.GError
	var goerr error
	var cret C.gint
	var ret2 int

	cret = C.g_unix_fd_list_append(arg0, fd, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = C.gint(cret)

	return goerr, ret2
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
func (l unixFDList) Get(index_ int) (gint int, err error) {
	var arg0 *C.GUnixFDList
	var arg1 C.gint

	arg0 = (*C.GUnixFDList)(unsafe.Pointer(l.Native()))
	arg1 = C.gint(index_)

	var errout *C.GError
	var goerr error
	var cret C.gint
	var ret2 int

	cret = C.g_unix_fd_list_get(arg0, index_, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = C.gint(cret)

	return goerr, ret2
}

// Length gets the length of @list (ie: the number of file descriptors
// contained within).
func (l unixFDList) Length() int {
	var arg0 *C.GUnixFDList

	arg0 = (*C.GUnixFDList)(unsafe.Pointer(l.Native()))

	var cret C.gint
	var ret1 int

	cret = C.g_unix_fd_list_get_length(arg0)

	ret1 = C.gint(cret)

	return ret1
}

// PeekFds returns the array of file descriptors that is contained in this
// object.
//
// After this call, the descriptors remain the property of @list. The caller
// must not close them and must not free the array. The array is valid only
// until @list is changed in any way.
//
// If @length is non-nil then it is set to the number of file descriptors in
// the returned array. The returned array is also terminated with -1.
//
// This function never returns nil. In case there are no file descriptors
// contained in @list, an empty array is returned.
func (l unixFDList) PeekFds() (length int, gints []int) {
	var arg0 *C.GUnixFDList

	arg0 = (*C.GUnixFDList)(unsafe.Pointer(l.Native()))

	var cret *C.gint
	var arg1 *C.gint
	var ret2 []int

	cret = C.g_unix_fd_list_peek_fds(arg0, &arg1)

	ret2 = make([]int, arg1)
	for i := 0; i < uintptr(arg1); i++ {
		src := (C.gint)(ptr.Add(unsafe.Pointer(cret), i))
		ret2[i] = C.gint(src)
	}

	return ret1, ret2
}

// StealFds returns the array of file descriptors that is contained in this
// object.
//
// After this call, the descriptors are no longer contained in @list.
// Further calls will return an empty list (unless more descriptors have
// been added).
//
// The return result of this function must be freed with g_free(). The
// caller is also responsible for closing all of the file descriptors. The
// file descriptors in the array are set to close-on-exec.
//
// If @length is non-nil then it is set to the number of file descriptors in
// the returned array. The returned array is also terminated with -1.
//
// This function never returns nil. In case there are no file descriptors
// contained in @list, an empty array is returned.
func (l unixFDList) StealFds() (length int, gints []int) {
	var arg0 *C.GUnixFDList

	arg0 = (*C.GUnixFDList)(unsafe.Pointer(l.Native()))

	var cret *C.gint
	var arg1 *C.gint
	var ret2 []int

	cret = C.g_unix_fd_list_steal_fds(arg0, &arg1)

	ptr.SetSlice(unsafe.Pointer(&ret2), unsafe.Pointer(cret), int(arg1))
	runtime.SetFinalizer(&ret2, func(v *[]int) {
		C.free(ptr.Slice(unsafe.Pointer(v)))
	})

	return ret1, ret2
}

type UnixFDListPrivate struct {
	native C.GUnixFDListPrivate
}

// WrapUnixFDListPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapUnixFDListPrivate(ptr unsafe.Pointer) *UnixFDListPrivate {
	if ptr == nil {
		return nil
	}

	return (*UnixFDListPrivate)(ptr)
}

func marshalUnixFDListPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapUnixFDListPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (u *UnixFDListPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&u.native)
}
