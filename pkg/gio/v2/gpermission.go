// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
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
//
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_permission_get_type()), F: marshalPermission},
	})
}

// Permission: a #GPermission represents the status of the caller's permission
// to perform a certain action.
//
// You can query if the action is currently allowed and if it is possible to
// acquire the permission so that the action will be allowed in the future.
//
// There is also an API to actually acquire the permission and one to release
// it.
//
// As an example, a #GPermission might represent the ability for the user to
// write to a #GSettings object. This #GPermission object could then be used to
// decide if it is appropriate to show a "Click here to unlock" button in a
// dialog and to provide the mechanism to invoke when that button is clicked.
type Permission interface {
	gextras.Objector

	AcquirePermission(cancellable Cancellable) error

	AcquireAsyncPermission(cancellable Cancellable, callback AsyncReadyCallback)

	AcquireFinishPermission(result AsyncResult) error

	Allowed() bool

	CanAcquire() bool

	CanRelease() bool

	ImplUpdatePermission(allowed bool, canAcquire bool, canRelease bool)

	ReleasePermission(cancellable Cancellable) error

	ReleaseAsyncPermission(cancellable Cancellable, callback AsyncReadyCallback)

	ReleaseFinishPermission(result AsyncResult) error
}

// permission implements the Permission class.
type permission struct {
	gextras.Objector
}

// WrapPermission wraps a GObject to the right type. It is
// primarily used internally.
func WrapPermission(obj *externglib.Object) Permission {
	return permission{
		Objector: obj,
	}
}

func marshalPermission(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPermission(obj), nil
}

func (p permission) AcquirePermission(cancellable Cancellable) error {
	var _arg0 *C.GPermission  // out
	var _arg1 *C.GCancellable // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_permission_acquire(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (p permission) AcquireAsyncPermission(cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GPermission        // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg2 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg3 = C.gpointer(box.Assign(callback))

	C.g_permission_acquire_async(_arg0, _arg1, _arg2, _arg3)
}

func (p permission) AcquireFinishPermission(result AsyncResult) error {
	var _arg0 *C.GPermission  // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_permission_acquire_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (p permission) Allowed() bool {
	var _arg0 *C.GPermission // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))

	_cret = C.g_permission_get_allowed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p permission) CanAcquire() bool {
	var _arg0 *C.GPermission // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))

	_cret = C.g_permission_get_can_acquire(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p permission) CanRelease() bool {
	var _arg0 *C.GPermission // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))

	_cret = C.g_permission_get_can_release(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p permission) ImplUpdatePermission(allowed bool, canAcquire bool, canRelease bool) {
	var _arg0 *C.GPermission // out
	var _arg1 C.gboolean     // out
	var _arg2 C.gboolean     // out
	var _arg3 C.gboolean     // out

	_arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))
	if allowed {
		_arg1 = C.TRUE
	}
	if canAcquire {
		_arg2 = C.TRUE
	}
	if canRelease {
		_arg3 = C.TRUE
	}

	C.g_permission_impl_update(_arg0, _arg1, _arg2, _arg3)
}

func (p permission) ReleasePermission(cancellable Cancellable) error {
	var _arg0 *C.GPermission  // out
	var _arg1 *C.GCancellable // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_permission_release(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (p permission) ReleaseAsyncPermission(cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GPermission        // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg2 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg3 = C.gpointer(box.Assign(callback))

	C.g_permission_release_async(_arg0, _arg1, _arg2, _arg3)
}

func (p permission) ReleaseFinishPermission(result AsyncResult) error {
	var _arg0 *C.GPermission  // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_permission_release_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
