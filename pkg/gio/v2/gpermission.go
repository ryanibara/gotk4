// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gio2_PermissionClass_acquire(GPermission*, GCancellable*, GError**);
// extern gboolean _gotk4_gio2_PermissionClass_acquire_finish(GPermission*, GAsyncResult*, GError**);
// extern gboolean _gotk4_gio2_PermissionClass_release(GPermission*, GCancellable*, GError**);
// extern gboolean _gotk4_gio2_PermissionClass_release_finish(GPermission*, GAsyncResult*, GError**);
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// GType values.
var (
	GTypePermission = coreglib.Type(C.g_permission_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePermission, F: marshalPermission},
	})
}

// PermissionOverrider contains methods that are overridable.
type PermissionOverrider interface {
	// Acquire attempts to acquire the permission represented by permission.
	//
	// The precise method by which this happens depends on the permission and
	// the underlying authentication mechanism. A simple example is that a
	// dialog may appear asking the user to enter their password.
	//
	// You should check with g_permission_get_can_acquire() before calling this
	// function.
	//
	// If the permission is acquired then TRUE is returned. Otherwise, FALSE is
	// returned and error is set appropriately.
	//
	// This call is blocking, likely for a very long time (in the case that user
	// interaction is required). See g_permission_acquire_async() for the
	// non-blocking version.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional) or NULL.
	//
	Acquire(ctx context.Context) error
	// AcquireFinish collects the result of attempting to acquire the permission
	// represented by permission.
	//
	// This is the second half of the asynchronous version of
	// g_permission_acquire().
	//
	// The function takes the following parameters:
	//
	//    - result given to the ReadyCallback.
	//
	AcquireFinish(result AsyncResulter) error
	// Release attempts to release the permission represented by permission.
	//
	// The precise method by which this happens depends on the permission and
	// the underlying authentication mechanism. In most cases the permission
	// will be dropped immediately without further action.
	//
	// You should check with g_permission_get_can_release() before calling this
	// function.
	//
	// If the permission is released then TRUE is returned. Otherwise, FALSE is
	// returned and error is set appropriately.
	//
	// This call is blocking, likely for a very long time (in the case that user
	// interaction is required). See g_permission_release_async() for the
	// non-blocking version.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional) or NULL.
	//
	Release(ctx context.Context) error
	// ReleaseFinish collects the result of attempting to release the permission
	// represented by permission.
	//
	// This is the second half of the asynchronous version of
	// g_permission_release().
	//
	// The function takes the following parameters:
	//
	//    - result given to the ReadyCallback.
	//
	ReleaseFinish(result AsyncResulter) error
}

// Permission represents the status of the caller's permission to perform a
// certain action.
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
type Permission struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Permission)(nil)
)

// Permissioner describes types inherited from class Permission.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Permissioner interface {
	coreglib.Objector
	basePermission() *Permission
}

var _ Permissioner = (*Permission)(nil)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:        GTypePermission,
		GoType:       reflect.TypeOf((*Permission)(nil)),
		InitClass:    initClassPermission,
		ClassSize:    uint32(unsafe.Sizeof(C.GPermission{})),
		InstanceSize: uint32(unsafe.Sizeof(C.GPermissionClass{})),
	})
}

func initClassPermission(gclass unsafe.Pointer, goval any) {

	pclass := (*C.GPermissionClass)(unsafe.Pointer(gclass))

	if _, ok := goval.(interface {
		Acquire(ctx context.Context) error
	}); ok {
		pclass.acquire = (*[0]byte)(C._gotk4_gio2_PermissionClass_acquire)
	}

	if _, ok := goval.(interface {
		AcquireFinish(result AsyncResulter) error
	}); ok {
		pclass.acquire_finish = (*[0]byte)(C._gotk4_gio2_PermissionClass_acquire_finish)
	}

	if _, ok := goval.(interface {
		Release(ctx context.Context) error
	}); ok {
		pclass.release = (*[0]byte)(C._gotk4_gio2_PermissionClass_release)
	}

	if _, ok := goval.(interface {
		ReleaseFinish(result AsyncResulter) error
	}); ok {
		pclass.release_finish = (*[0]byte)(C._gotk4_gio2_PermissionClass_release_finish)
	}
}

//export _gotk4_gio2_PermissionClass_acquire
func _gotk4_gio2_PermissionClass_acquire(arg0 *C.GPermission, arg1 *C.GCancellable, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Acquire(ctx context.Context) error
	})

	var _cancellable context.Context // out

	if arg1 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg1))
	}

	_goerr := iface.Acquire(_cancellable)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_PermissionClass_acquire_finish
func _gotk4_gio2_PermissionClass_acquire_finish(arg0 *C.GPermission, arg1 *C.GAsyncResult, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface {
		AcquireFinish(result AsyncResulter) error
	})

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	_goerr := iface.AcquireFinish(_result)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_PermissionClass_release
func _gotk4_gio2_PermissionClass_release(arg0 *C.GPermission, arg1 *C.GCancellable, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Release(ctx context.Context) error
	})

	var _cancellable context.Context // out

	if arg1 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg1))
	}

	_goerr := iface.Release(_cancellable)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_PermissionClass_release_finish
func _gotk4_gio2_PermissionClass_release_finish(arg0 *C.GPermission, arg1 *C.GAsyncResult, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface {
		ReleaseFinish(result AsyncResulter) error
	})

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	_goerr := iface.ReleaseFinish(_result)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

func wrapPermission(obj *coreglib.Object) *Permission {
	return &Permission{
		Object: obj,
	}
}

func marshalPermission(p uintptr) (interface{}, error) {
	return wrapPermission(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (permission *Permission) basePermission() *Permission {
	return permission
}

// BasePermission returns the underlying base object.
func BasePermission(obj Permissioner) *Permission {
	return obj.basePermission()
}

// Acquire attempts to acquire the permission represented by permission.
//
// The precise method by which this happens depends on the permission and the
// underlying authentication mechanism. A simple example is that a dialog may
// appear asking the user to enter their password.
//
// You should check with g_permission_get_can_acquire() before calling this
// function.
//
// If the permission is acquired then TRUE is returned. Otherwise, FALSE is
// returned and error is set appropriately.
//
// This call is blocking, likely for a very long time (in the case that user
// interaction is required). See g_permission_acquire_async() for the
// non-blocking version.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//
func (permission *Permission) Acquire(ctx context.Context) error {
	var _arg0 *C.GPermission  // out
	var _arg1 *C.GCancellable // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(coreglib.InternObject(permission).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	C.g_permission_acquire(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(permission)
	runtime.KeepAlive(ctx)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// AcquireAsync attempts to acquire the permission represented by permission.
//
// This is the first half of the asynchronous version of g_permission_acquire().
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - callback (optional) to call when done.
//
func (permission *Permission) AcquireAsync(ctx context.Context, callback AsyncReadyCallback) {
	var _arg0 *C.GPermission        // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GPermission)(unsafe.Pointer(coreglib.InternObject(permission).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_permission_acquire_async(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(permission)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// AcquireFinish collects the result of attempting to acquire the permission
// represented by permission.
//
// This is the second half of the asynchronous version of
// g_permission_acquire().
//
// The function takes the following parameters:
//
//    - result given to the ReadyCallback.
//
func (permission *Permission) AcquireFinish(result AsyncResulter) error {
	var _arg0 *C.GPermission  // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(coreglib.InternObject(permission).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C.g_permission_acquire_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(permission)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Allowed gets the value of the 'allowed' property. This property is TRUE if
// the caller currently has permission to perform the action that permission
// represents the permission to perform.
//
// The function returns the following values:
//
//    - ok: value of the 'allowed' property.
//
func (permission *Permission) Allowed() bool {
	var _arg0 *C.GPermission // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(coreglib.InternObject(permission).Native()))

	_cret = C.g_permission_get_allowed(_arg0)
	runtime.KeepAlive(permission)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanAcquire gets the value of the 'can-acquire' property. This property is
// TRUE if it is generally possible to acquire the permission by calling
// g_permission_acquire().
//
// The function returns the following values:
//
//    - ok: value of the 'can-acquire' property.
//
func (permission *Permission) CanAcquire() bool {
	var _arg0 *C.GPermission // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(coreglib.InternObject(permission).Native()))

	_cret = C.g_permission_get_can_acquire(_arg0)
	runtime.KeepAlive(permission)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanRelease gets the value of the 'can-release' property. This property is
// TRUE if it is generally possible to release the permission by calling
// g_permission_release().
//
// The function returns the following values:
//
//    - ok: value of the 'can-release' property.
//
func (permission *Permission) CanRelease() bool {
	var _arg0 *C.GPermission // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(coreglib.InternObject(permission).Native()))

	_cret = C.g_permission_get_can_release(_arg0)
	runtime.KeepAlive(permission)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ImplUpdate: this function is called by the #GPermission implementation to
// update the properties of the permission. You should never call this function
// except from a #GPermission implementation.
//
// GObject notify signals are generated, as appropriate.
//
// The function takes the following parameters:
//
//    - allowed: new value for the 'allowed' property.
//    - canAcquire: new value for the 'can-acquire' property.
//    - canRelease: new value for the 'can-release' property.
//
func (permission *Permission) ImplUpdate(allowed, canAcquire, canRelease bool) {
	var _arg0 *C.GPermission // out
	var _arg1 C.gboolean     // out
	var _arg2 C.gboolean     // out
	var _arg3 C.gboolean     // out

	_arg0 = (*C.GPermission)(unsafe.Pointer(coreglib.InternObject(permission).Native()))
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
	runtime.KeepAlive(permission)
	runtime.KeepAlive(allowed)
	runtime.KeepAlive(canAcquire)
	runtime.KeepAlive(canRelease)
}

// Release attempts to release the permission represented by permission.
//
// The precise method by which this happens depends on the permission and the
// underlying authentication mechanism. In most cases the permission will be
// dropped immediately without further action.
//
// You should check with g_permission_get_can_release() before calling this
// function.
//
// If the permission is released then TRUE is returned. Otherwise, FALSE is
// returned and error is set appropriately.
//
// This call is blocking, likely for a very long time (in the case that user
// interaction is required). See g_permission_release_async() for the
// non-blocking version.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//
func (permission *Permission) Release(ctx context.Context) error {
	var _arg0 *C.GPermission  // out
	var _arg1 *C.GCancellable // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(coreglib.InternObject(permission).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	C.g_permission_release(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(permission)
	runtime.KeepAlive(ctx)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// ReleaseAsync attempts to release the permission represented by permission.
//
// This is the first half of the asynchronous version of g_permission_release().
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - callback (optional) to call when done.
//
func (permission *Permission) ReleaseAsync(ctx context.Context, callback AsyncReadyCallback) {
	var _arg0 *C.GPermission        // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GPermission)(unsafe.Pointer(coreglib.InternObject(permission).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_permission_release_async(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(permission)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// ReleaseFinish collects the result of attempting to release the permission
// represented by permission.
//
// This is the second half of the asynchronous version of
// g_permission_release().
//
// The function takes the following parameters:
//
//    - result given to the ReadyCallback.
//
func (permission *Permission) ReleaseFinish(result AsyncResulter) error {
	var _arg0 *C.GPermission  // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(coreglib.InternObject(permission).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C.g_permission_release_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(permission)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
