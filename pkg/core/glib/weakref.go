package glib

// #include <glib.h>
// #include <glib-object.h>
//
// extern void _gotk4_glib_weak_notify(gpointer, GObject*);
import "C"

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/intern"
)

// WeakRefObject is like SetFinalizer, except it's not thread-safe (so notify
// SHOULD NOT REFERENCE OBJECT). It is best that you just do not use this at
// all.
func WeakRefObject(obj Objector, notify func()) {
	data := gbox.AssignOnce(notify)
	C.g_object_weak_ref(
		BaseObject(obj).native(),
		C.GWeakNotify(C._gotk4_glib_weak_notify),
		C.gpointer(data))
}

//export _gotk4_glib_weak_notify
func _gotk4_glib_weak_notify(data C.gpointer, _ *C.GObject) {
	notify := gbox.Get(uintptr(data)).(func())
	notify()
}

// WeakRef wraps GWeakRef. It provides a container that allows the user to
// obtain a weak reference to a CGo GObject. The weak reference is thread-safe
// and will be cleared when the object is finalized.
type WeakRef struct {
	weak C.GWeakRef
	gtyp Type
}

// NewWeakRef creates a new weak reference on the Go heap to the given GObject's
// C pointer. The fact that the returned WeakRef is in Go-allocated memory does
// not actually add a reference to the object, which is the default behavior.
func NewWeakRef(obj Objector) *WeakRef {
	wk := new(WeakRef)
	wk.gtyp = BaseObject(obj).Type()
	C.g_weak_ref_init(&wk.weak, C.gpointer(BaseObject(obj).native()))

	// Unsure if calling clear is needed, but we'd rather be careful.
	runtime.SetFinalizer(wk, (*WeakRef).clear)
	runtime.KeepAlive(obj)

	return wk
}

func (r *WeakRef) clear() {
	C.g_weak_ref_clear(&r.weak)
}

// Get acquires a strong reference to the object if the weak reference is still
// valid. If the weak reference is no longer valid, Get returns nil.
func (r *WeakRef) Get() Objector {
	// The thread safetyness of this is actually debatable. I haven't confirmed
	// this, but it should still work fine.
	gobjectPtr := C.g_weak_ref_get(&r.weak)
	if gobjectPtr == nil {
		return nil
	}

	// weak_ref_get actually takes a strong reference atomically. With the rest
	// of this function, we either obtain a strong Go reference (mapping to the
	// toggle reference) or we cannot get a reference at all. In either case, we
	// can safely unref the object after this function.
	defer C.g_object_unref(C.gpointer(gobjectPtr))

	// Try to see if we have the object in our GObject intern pool. If not,
	// don't try to construct a new one.
	box := intern.TryGet(unsafe.Pointer(gobjectPtr))
	if box == nil {
		return nil
	}

	// Construct a new Object pointer from the box. Keep in mind our equality
	// guarantees.
	return (&Object{box: box}).CastType(r.gtyp)
}
