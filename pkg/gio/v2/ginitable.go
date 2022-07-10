// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gio2_InitableIface_init(void*, void*, GError**);
import "C"

// GTypeInitable returns the GType for the type Initable.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeInitable() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "Initable").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalInitable)
	return gtype
}

// InitableOverrider contains methods that are overridable.
type InitableOverrider interface {
	// Init initializes the object implementing the interface.
	//
	// This method is intended for language bindings. If writing in C,
	// g_initable_new() should typically be used instead.
	//
	// The object must be initialized before any real use after initial
	// construction, either with this function or g_async_initable_init_async().
	//
	// Implementations may also support cancellation. If cancellable is not
	// NULL, then initialization can be cancelled by triggering the cancellable
	// object from another thread. If the operation was cancelled, the error
	// G_IO_ERROR_CANCELLED will be returned. If cancellable is not NULL and the
	// object doesn't support cancellable initialization the error
	// G_IO_ERROR_NOT_SUPPORTED will be returned.
	//
	// If the object is not initialized, or initialization returns with an
	// error, then all operations on the object except g_object_ref() and
	// g_object_unref() are considered to be invalid, and have undefined
	// behaviour. See the [introduction][ginitable] for more details.
	//
	// Callers should not assume that a class which implements #GInitable can be
	// initialized multiple times, unless the class explicitly documents itself
	// as supporting this. Generally, a class’ implementation of init() can
	// assume (and assert) that it will only be called once. Previously, this
	// documentation recommended all #GInitable implementations should be
	// idempotent; that recommendation was relaxed in GLib 2.54.
	//
	// If a class explicitly supports being initialized multiple times, it is
	// recommended that the method is idempotent: multiple calls with the same
	// arguments should return the same results. Only the first call initializes
	// the object; further calls return the result of the first call.
	//
	// One reason why a class might need to support idempotent initialization is
	// if it is designed to be used via the singleton pattern, with a
	// Class.constructor that sometimes returns an existing instance. In this
	// pattern, a caller would expect to be able to call g_initable_init() on
	// the result of g_object_new(), regardless of whether it is in fact a new
	// instance.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable object, NULL to ignore.
	//
	Init(ctx context.Context) error
}

// Initable is implemented by objects that can fail during initialization. If an
// object implements this interface then it must be initialized as the first
// thing after construction, either via g_initable_init() or
// g_async_initable_init_async() (the latter is only available if it also
// implements Initable).
//
// If the object is not initialized, or initialization returns with an error,
// then all operations on the object except g_object_ref() and g_object_unref()
// are considered to be invalid, and have undefined behaviour. They will often
// fail with g_critical() or g_warning(), but this must not be relied on.
//
// Users of objects implementing this are not intended to use the interface
// method directly, instead it will be used automatically in various ways. For C
// applications you generally just call g_initable_new() directly, or indirectly
// via a foo_thing_new() wrapper. This will call g_initable_init() under the
// cover, returning NULL and setting a #GError on failure (at which point the
// instance is unreferenced).
//
// For bindings in languages where the native constructor supports exceptions
// the binding could check for objects implementing GInitable during normal
// construction and automatically initialize them, throwing an exception on
// failure.
//
// Initable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Initable struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Initable)(nil)
)

// Initabler describes Initable's interface methods.
type Initabler interface {
	coreglib.Objector

	// Init initializes the object implementing the interface.
	Init(ctx context.Context) error
}

var _ Initabler = (*Initable)(nil)

func ifaceInitInitabler(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Gio", "InitableIface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("init"))) = unsafe.Pointer(C._gotk4_gio2_InitableIface_init)
}

//export _gotk4_gio2_InitableIface_init
func _gotk4_gio2_InitableIface_init(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(InitableOverrider)

	var _cancellable context.Context // out

	if arg1 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg1))
	}

	_goerr := iface.Init(_cancellable)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

func wrapInitable(obj *coreglib.Object) *Initable {
	return &Initable{
		Object: obj,
	}
}

func marshalInitable(p uintptr) (interface{}, error) {
	return wrapInitable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Init initializes the object implementing the interface.
//
// This method is intended for language bindings. If writing in C,
// g_initable_new() should typically be used instead.
//
// The object must be initialized before any real use after initial
// construction, either with this function or g_async_initable_init_async().
//
// Implementations may also support cancellation. If cancellable is not NULL,
// then initialization can be cancelled by triggering the cancellable object
// from another thread. If the operation was cancelled, the error
// G_IO_ERROR_CANCELLED will be returned. If cancellable is not NULL and the
// object doesn't support cancellable initialization the error
// G_IO_ERROR_NOT_SUPPORTED will be returned.
//
// If the object is not initialized, or initialization returns with an error,
// then all operations on the object except g_object_ref() and g_object_unref()
// are considered to be invalid, and have undefined behaviour. See the
// [introduction][ginitable] for more details.
//
// Callers should not assume that a class which implements #GInitable can be
// initialized multiple times, unless the class explicitly documents itself as
// supporting this. Generally, a class’ implementation of init() can assume (and
// assert) that it will only be called once. Previously, this documentation
// recommended all #GInitable implementations should be idempotent; that
// recommendation was relaxed in GLib 2.54.
//
// If a class explicitly supports being initialized multiple times, it is
// recommended that the method is idempotent: multiple calls with the same
// arguments should return the same results. Only the first call initializes the
// object; further calls return the result of the first call.
//
// One reason why a class might need to support idempotent initialization is if
// it is designed to be used via the singleton pattern, with a Class.constructor
// that sometimes returns an existing instance. In this pattern, a caller would
// expect to be able to call g_initable_init() on the result of g_object_new(),
// regardless of whether it is in fact a new instance.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//
func (initable *Initable) Init(ctx context.Context) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(initable).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[1] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}

	_info := girepository.MustFind("Gio", "Initable")
	_info.InvokeIfaceMethod("init", _args[:], nil)

	runtime.KeepAlive(initable)
	runtime.KeepAlive(ctx)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cerr))))
	}

	return _goerr
}
