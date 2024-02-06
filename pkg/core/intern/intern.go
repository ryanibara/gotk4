// Package intern implements value interning for Cgo sharing.
package intern

// #cgo pkg-config: gobject-2.0
// #include <glib-object.h>
//
// extern void goToggleNotify(gpointer, GObject*, gboolean);
// static const gchar* gotk4_object_type_name(gpointer obj) { return G_OBJECT_TYPE_NAME(obj); };
import "C"

import (
	"fmt"
	"log"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gdebug"

	// Require a non-moving GC for heap pointers. Current GC is moving only by
	// the stack. See https://github.com/go4org/intern.
	_ "go4.org/unsafe/assume-no-moving-gc"
)

const maxTypesAllowed = 2

var knownTypes uint32 = 0

type BoxedType[T any] struct {
	ctor func(*Box) *T
	id   uint32
}

func RegisterType[T any](ctor func(*Box) *T) BoxedType[T] {
	t := BoxedType[T]{
		ctor: ctor,
		id:   atomic.AddUint32(&knownTypes, 1) - 1,
	}
	if t.id > maxTypesAllowed {
		panic("BoxedType ID overflow")
	}
	return t
}

func (t *BoxedType[T]) Get(box *Box) *T {
	old := atomic.LoadPointer(&box.data[t.id])
	if old != nil {
		return (*T)(old)
	}

	if t.ctor == nil {
		return nil
	}

	new := t.ctor(box)

	if atomic.CompareAndSwapPointer(&box.data[t.id], nil, unsafe.Pointer(new)) {
		return new
	}

	ptr := atomic.LoadPointer(&box.data[t.id])
	if ptr != nil {
		return (*T)(ptr)
	}

	panic("Load returned nil after CompareAndSwap(old = nil) failed")
}

func (t *BoxedType[T]) Set(box *Box, v *T) {
	if t.ctor != nil {
		panic("bug: Set not permitted if t.ctor != nil")
	}
	atomic.StorePointer(&box.data[t.id], unsafe.Pointer(v))
}

func (t *BoxedType[T]) Delete(box *Box) {
	atomic.StorePointer(&box.data[t.id], nil)
}

// Box is an opaque type holding extra data.
type Box struct {
	gobject unsafe.Pointer
	dummy   *boxDummy
	data    [maxTypesAllowed]unsafe.Pointer
}

type boxDummy struct {
	gobject unsafe.Pointer
}

// Object returns Box's C GObject pointer.
func (b *Box) GObject() unsafe.Pointer {
	return b.gobject
}

// Hack to force an object on the heap.
var never bool
var sink_ interface{}

//go:nosplit
func sink(v interface{}) {
	if never {
		sink_ = v
	}
}

var (
	traceObjects  = gdebug.NewDebugLoggerNullable("trace-objects")
	toggleRefs    = gdebug.NewDebugLoggerNullable("toggle-refs")
	objectProfile *pprof.Profile
)

func init() {
	if gdebug.HasKey("profile-objects") {
		objectProfile = pprof.NewProfile("gotk4-object-box")
	}
}

func objInfo(obj unsafe.Pointer) string {
	return fmt.Sprintf("%p (%s):", obj, C.GoString(C.gotk4_object_type_name(C.gpointer(obj))))
}

func objRefCount(obj unsafe.Pointer) int {
	return int(C.g_atomic_int_get((*C.gint)(unsafe.Pointer(&(*C.GObject)(obj).ref_count))))
}

// newBox creates a zero-value instance of Box.
func newBox(obj unsafe.Pointer) *Box {
	box := &Box{}
	box.gobject = obj

	// Cheat Go's GC by adding a finalizer to a dummy pointer that is inside Box
	// but is not Box itself.
	box.dummy = &boxDummy{gobject: obj}
	sink(box.dummy)
	runtime.SetFinalizer(box.dummy, finalizeBox)

	if objectProfile != nil {
		objectProfile.Add(obj, 3)
	}

	if traceObjects != nil {
		traceObjects.Printf("%p: %s", obj, debug.Stack())
	}

	// Force box on the heap. Objects on the stack can move, but not objects on
	// the heap. At least not for now; the assume-no-moving-gc import will
	// guard against that.
	sink(box)

	return box
}

// shared contains shared closure data.
var shared = struct {
	mu sync.RWMutex
	// weak stores *Box while the object is in Go's heap. The finalizer will
	// move *Box to strong if the reference is toggled. This is only the case,
	// because the finalizer will not run otherwise.
	weak map[unsafe.Pointer]uintptr
	// strong stores *Box while the object is still referenced by C but not Go.
	strong map[unsafe.Pointer]*Box
}{
	weak:   make(map[unsafe.Pointer]uintptr, 1024),
	strong: make(map[unsafe.Pointer]*Box, 1024),
}

// TryGet gets the Box associated with the GObject or nil if it's gone. The
// caller must not retain the Box pointer anywhere.
//
//go:nosplit
func TryGet(gobject unsafe.Pointer) *Box {
	shared.mu.RLock()
	box, _ := gets(gobject)
	shared.mu.RUnlock()
	return box
}

// Get gets the interned box for the given GObject C pointer. If the object is
// new or unknown, then a new box is made. If the intern box already exists for
// a given C pointer, then that box is weakly referenced and returned. The box
// will be reference-counted; the caller must use ShouldFree to unreference it.
//
//go:nosplit
func Get(gobject unsafe.Pointer, take bool) *Box {
	// If the registry does not exist, then we'll have to globally register it.
	// If the registry is currently strongly referenced, then we must move it to
	// a weak reference.

	box := TryGet(gobject)
	if box != nil {
		return box
	}

	shared.mu.Lock()

	box, _ = gets(gobject)
	if box != nil {
		shared.mu.Unlock()
		return box
	}

	box = newBox(gobject)

	// add_toggle_ref's documentation states:
	//
	//    Since a (normal) reference must be held to the object before
	//    calling g_object_add_toggle_ref(), the initial state of the
	//    reverse link is always strong.
	//
	shared.strong[gobject] = box

	if toggleRefs != nil {
		toggleRefs.Println(objInfo(gobject),
			"Get: will introduce new box, current ref =", objRefCount(gobject))
	}

	shared.mu.Unlock()

	C.g_object_add_toggle_ref(
		(*C.GObject)(gobject),
		(*[0]byte)(C.goToggleNotify), nil,
	)

	// We should already have a strong reference. Sink the object in case. This
	// will force the reference to be truly strong.
	if C.g_object_is_floating(C.gpointer(gobject)) != C.FALSE {
		// First, we need to ref_sink the object to convert the floating
		// reference to a strong reference.
		C.g_object_ref_sink(C.gpointer(gobject))
		// Then, we need to unref it to balance the ref_sink.
		C.g_object_unref(C.gpointer(gobject))

		if toggleRefs != nil {
			toggleRefs.Println(objInfo(gobject),
				"Get: ref_sink'd the object, current ref =", objRefCount(gobject))
		}
	}

	// If we're "not taking," then we can assume our ownership over the object,
	// meaning the strong reference is now ours. That means we need to replace
	// it, not add.
	if !take {
		if toggleRefs != nil {
			toggleRefs.Println(objInfo(gobject),
				"Get: not taking, so unrefing the object, current ref =", objRefCount(gobject))
		}
		C.g_object_unref(C.gpointer(gobject))
	}

	if toggleRefs != nil {
		toggleRefs.Println(objInfo(gobject),
			"Get: introduced new box, current ref =", objRefCount(gobject))
	}

	// Undo the initial ref_sink.
	// C.g_object_unref(C.gpointer(gobject))

	return box
}

// Free explicitly frees the box permanently. It must not be resurrected after
// this.
//
// Deprecated: this function is no longer needed.
func Free(box *Box) {
	panic("not implemented")
}

// finalizeBox only delays its finalization until GLib notifies us a toggle. It
// does so for as long as an object is stored only in the Go heap. Once the
// object is also shared, the toggle notifier will strongly reference the Box.
//
//go:nosplit
func finalizeBox(dummy *boxDummy) {
	if dummy == nil {
		panic("bug: finalizeBox called with nil dummy")
	}

	shared.mu.Lock()

	box, strong := gets(dummy.gobject)
	if box == nil {
		log.Print("gotk4: intern: finalizer got unknown gobject ", dummy.gobject, ", ignoring")
		shared.mu.Unlock()
		return
	}

	var objInfoRes string
	if toggleRefs != nil {
		objInfoRes = objInfo(dummy.gobject)
		toggleRefs.Println(objInfoRes, "finalizeBox: acquiring lock...")
	}

	if strong {
		// If the closures are strong-referenced, then they might still be
		// referenced from the C side, and those closures might access this
		// object. Don't free.

		// Delegate finalizing to the next cycle.
		runtime.SetFinalizer(dummy, finalizeBox)

		shared.mu.Unlock()

		if toggleRefs != nil {
			toggleRefs.Println(objInfoRes, "finalizeBox: moving finalize to next GC cycle")
		}
	} else {
		// If the closures are weak-referenced, then the object reference hasn't
		// been toggled yet. Since the object is going away and we're still
		// weakly referenced, we can wipe the closures away.
		delete(shared.weak, dummy.gobject)

		shared.mu.Unlock()

		// Unreference the object. This will potentially free the object as
		// well. The closures are definitely gone at this point.
		C.g_object_remove_toggle_ref(
			(*C.GObject)(unsafe.Pointer(dummy.gobject)),
			(*[0]byte)(C.goToggleNotify), nil,
		)

		if toggleRefs != nil {
			toggleRefs.Println(objInfoRes, "finalizeBox: removed toggle ref during GC")
		}

		if objectProfile != nil {
			objectProfile.Remove(dummy.gobject)
		}
	}
}

// goToggleNotify is called by GLib on each toggle notification. It doesn't
// actually free anything and relies on Box's finalizer to free both the box and
// the C GObject.
//
//go:nosplit
//export goToggleNotify
func goToggleNotify(_ C.gpointer, obj *C.GObject, isLastInt C.gboolean) {
	gobject := unsafe.Pointer(obj)
	isLast := isLastInt != C.FALSE

	shared.mu.Lock()

	if isLast {
		// delete(shared.sharing, gobject)
		makeWeak(gobject)
	} else {
		// shared.sharing[gobject] = struct{}{}
		makeStrong(gobject)
	}

	shared.mu.Unlock()

	if toggleRefs != nil {
		toggleRefs.Println(objInfo(unsafe.Pointer(obj)), "goToggleNotify: is last =", isLast)
	}
}

//go:nocheckptr
//go:nosplit
func gets(gobject unsafe.Pointer) (b *Box, strong bool) {
	if strong, ok := shared.strong[gobject]; ok {
		return strong, true
	}

	if weak, ok := shared.weak[gobject]; ok {
		// If forObject is false, then that probably means this was called
		// inside goMarshal while the Go object is still alive, otherwise
		// toggleNotify would've moved it over. We don't have to worry about
		// this being freed as long as we acquire the mutex.
		//
		// TODO: does this actually resurrect the value properly? We have a
		// mutex to guard this which is also used in the finalizer, so it
		// shouldn't explode, but still.
		return (*Box)(unsafe.Pointer(weak)), false
	}

	return nil, false
}

// makeStrong forces the Box instance associated with the given object to be
// strongly referenced.
//
//go:nosplit
func makeStrong(gobject unsafe.Pointer) {
	// TODO: double mutex check, similar to ShouldFree.

	box, strong := gets(gobject)
	if toggleRefs != nil {
		toggleRefs.Println(objInfo(gobject), "makeStrong: obtained box", box, "strong =", strong)
	}
	if box == nil {
		return
	}

	if !strong {
		shared.strong[gobject] = box
		delete(shared.weak, gobject)
	}
}

// makeWeak forces the Box intsance associated with the given object to be
// weakly referenced.
//
//go:nosplit
func makeWeak(gobject unsafe.Pointer) {
	box, strong := gets(gobject)
	if toggleRefs != nil {
		toggleRefs.Println(objInfo(gobject), "makeWeak: obtained box", box, "strong =", strong)
	}
	if box == nil {
		return
	}

	if strong {
		shared.weak[gobject] = uintptr(unsafe.Pointer(box))
		delete(shared.strong, gobject)
	}
}
