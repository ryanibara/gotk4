// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <glib.h>
import "C"

// AsyncQueue: the GAsyncQueue struct is an opaque data structure which
// represents an asynchronous queue. It should only be accessed through the
// g_async_queue_* functions.
type AsyncQueue struct {
	native C.GAsyncQueue
}

// WrapAsyncQueue wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAsyncQueue(ptr unsafe.Pointer) *AsyncQueue {
	if ptr == nil {
		return nil
	}

	return (*AsyncQueue)(ptr)
}

func marshalAsyncQueue(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAsyncQueue(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AsyncQueue) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// Length returns the length of the queue.
//
// Actually this function returns the number of data items in the queue minus
// the number of waiting threads, so a negative value means waiting threads, and
// a positive value means available entries in the @queue. A return value of 0
// could mean n entries in the queue and n threads waiting. This can happen due
// to locking of the queue or due to scheduling.
func (q *AsyncQueue) Length() int {
	var arg0 *C.GAsyncQueue

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))

	var cret C.gint
	var ret1 int

	cret = C.g_async_queue_length(arg0)

	ret1 = C.gint(cret)

	return ret1
}

// LengthUnlocked returns the length of the queue.
//
// Actually this function returns the number of data items in the queue minus
// the number of waiting threads, so a negative value means waiting threads, and
// a positive value means available entries in the @queue. A return value of 0
// could mean n entries in the queue and n threads waiting. This can happen due
// to locking of the queue or due to scheduling.
//
// This function must be called while holding the @queue's lock.
func (q *AsyncQueue) LengthUnlocked() int {
	var arg0 *C.GAsyncQueue

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))

	var cret C.gint
	var ret1 int

	cret = C.g_async_queue_length_unlocked(arg0)

	ret1 = C.gint(cret)

	return ret1
}

// Lock acquires the @queue's lock. If another thread is already holding the
// lock, this call will block until the lock becomes available.
//
// Call g_async_queue_unlock() to drop the lock again.
//
// While holding the lock, you can only call the g_async_queue_*_unlocked()
// functions on @queue. Otherwise, deadlock may occur.
func (q *AsyncQueue) Lock() {
	var arg0 *C.GAsyncQueue

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))

	C.g_async_queue_lock(arg0)
}

// Pop pops data from the @queue. If @queue is empty, this function blocks until
// data becomes available.
func (q *AsyncQueue) Pop() interface{} {
	var arg0 *C.GAsyncQueue

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))

	var cret C.gpointer
	var ret1 interface{}

	cret = C.g_async_queue_pop(arg0)

	ret1 = C.gpointer(cret)

	return ret1
}

// PopUnlocked pops data from the @queue. If @queue is empty, this function
// blocks until data becomes available.
//
// This function must be called while holding the @queue's lock.
func (q *AsyncQueue) PopUnlocked() interface{} {
	var arg0 *C.GAsyncQueue

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))

	var cret C.gpointer
	var ret1 interface{}

	cret = C.g_async_queue_pop_unlocked(arg0)

	ret1 = C.gpointer(cret)

	return ret1
}

// Push pushes the @data into the @queue. @data must not be nil.
func (q *AsyncQueue) Push(data interface{}) {
	var arg0 *C.GAsyncQueue
	var arg1 C.gpointer

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))
	arg1 = C.gpointer(data)

	C.g_async_queue_push(arg0, data)
}

// PushFront pushes the @item into the @queue. @item must not be nil. In
// contrast to g_async_queue_push(), this function pushes the new item ahead of
// the items already in the queue, so that it will be the next one to be popped
// off the queue.
func (q *AsyncQueue) PushFront(item interface{}) {
	var arg0 *C.GAsyncQueue
	var arg1 C.gpointer

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))
	arg1 = C.gpointer(item)

	C.g_async_queue_push_front(arg0, item)
}

// PushFrontUnlocked pushes the @item into the @queue. @item must not be nil. In
// contrast to g_async_queue_push_unlocked(), this function pushes the new item
// ahead of the items already in the queue, so that it will be the next one to
// be popped off the queue.
//
// This function must be called while holding the @queue's lock.
func (q *AsyncQueue) PushFrontUnlocked(item interface{}) {
	var arg0 *C.GAsyncQueue
	var arg1 C.gpointer

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))
	arg1 = C.gpointer(item)

	C.g_async_queue_push_front_unlocked(arg0, item)
}

// PushSorted inserts @data into @queue using @func to determine the new
// position.
//
// This function requires that the @queue is sorted before pushing on new
// elements, see g_async_queue_sort().
//
// This function will lock @queue before it sorts the queue and unlock it when
// it is finished.
//
// For an example of @func see g_async_queue_sort().
func (q *AsyncQueue) PushSorted(data interface{}, fn CompareDataFunc) {
	var arg0 *C.GAsyncQueue

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))

	C.g_async_queue_push_sorted(arg0, data, fn, userData)
}

// PushSortedUnlocked inserts @data into @queue using @func to determine the new
// position.
//
// The sort function @func is passed two elements of the @queue. It should
// return 0 if they are equal, a negative value if the first element should be
// higher in the @queue or a positive value if the first element should be lower
// in the @queue than the second element.
//
// This function requires that the @queue is sorted before pushing on new
// elements, see g_async_queue_sort().
//
// This function must be called while holding the @queue's lock.
//
// For an example of @func see g_async_queue_sort().
func (q *AsyncQueue) PushSortedUnlocked(data interface{}, fn CompareDataFunc) {
	var arg0 *C.GAsyncQueue

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))

	C.g_async_queue_push_sorted_unlocked(arg0, data, fn, userData)
}

// PushUnlocked pushes the @data into the @queue. @data must not be nil.
//
// This function must be called while holding the @queue's lock.
func (q *AsyncQueue) PushUnlocked(data interface{}) {
	var arg0 *C.GAsyncQueue
	var arg1 C.gpointer

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))
	arg1 = C.gpointer(data)

	C.g_async_queue_push_unlocked(arg0, data)
}

// Ref increases the reference count of the asynchronous @queue by 1. You do not
// need to hold the lock to call this function.
func (q *AsyncQueue) Ref() *AsyncQueue {
	var arg0 *C.GAsyncQueue

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))

	var cret *C.GAsyncQueue
	var ret1 *AsyncQueue

	cret = C.g_async_queue_ref(arg0)

	ret1 = WrapAsyncQueue(unsafe.Pointer(cret))

	return ret1
}

// RefUnlocked increases the reference count of the asynchronous @queue by 1.
func (q *AsyncQueue) RefUnlocked() {
	var arg0 *C.GAsyncQueue

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))

	C.g_async_queue_ref_unlocked(arg0)
}

// Remove: remove an item from the queue.
func (q *AsyncQueue) Remove(item interface{}) bool {
	var arg0 *C.GAsyncQueue
	var arg1 C.gpointer

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))
	arg1 = C.gpointer(item)

	var cret C.gboolean
	var ret1 bool

	cret = C.g_async_queue_remove(arg0, item)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// RemoveUnlocked: remove an item from the queue.
//
// This function must be called while holding the @queue's lock.
func (q *AsyncQueue) RemoveUnlocked(item interface{}) bool {
	var arg0 *C.GAsyncQueue
	var arg1 C.gpointer

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))
	arg1 = C.gpointer(item)

	var cret C.gboolean
	var ret1 bool

	cret = C.g_async_queue_remove_unlocked(arg0, item)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Sort sorts @queue using @func.
//
// The sort function @func is passed two elements of the @queue. It should
// return 0 if they are equal, a negative value if the first element should be
// higher in the @queue or a positive value if the first element should be lower
// in the @queue than the second element.
//
// This function will lock @queue before it sorts the queue and unlock it when
// it is finished.
//
// If you were sorting a list of priority numbers to make sure the lowest
// priority would be at the top of the queue, you could use:
//
//     gint32 id1;
//     gint32 id2;
//
//     id1 = GPOINTER_TO_INT (element1);
//     id2 = GPOINTER_TO_INT (element2);
//
//     return (id1 > id2 ? +1 : id1 == id2 ? 0 : -1);
func (q *AsyncQueue) Sort(fn CompareDataFunc) {
	var arg0 *C.GAsyncQueue

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))

	C.g_async_queue_sort(arg0, fn, userData)
}

// SortUnlocked sorts @queue using @func.
//
// The sort function @func is passed two elements of the @queue. It should
// return 0 if they are equal, a negative value if the first element should be
// higher in the @queue or a positive value if the first element should be lower
// in the @queue than the second element.
//
// This function must be called while holding the @queue's lock.
func (q *AsyncQueue) SortUnlocked(fn CompareDataFunc) {
	var arg0 *C.GAsyncQueue

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))

	C.g_async_queue_sort_unlocked(arg0, fn, userData)
}

// TimedPop pops data from the @queue. If the queue is empty, blocks until
// @end_time or until data becomes available.
//
// If no data is received before @end_time, nil is returned.
//
// To easily calculate @end_time, a combination of g_get_real_time() and
// g_time_val_add() can be used.
func (q *AsyncQueue) TimedPop(endTime *TimeVal) interface{} {
	var arg0 *C.GAsyncQueue
	var arg1 *C.GTimeVal

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))
	arg1 = (*C.GTimeVal)(unsafe.Pointer(endTime.Native()))

	var cret C.gpointer
	var ret1 interface{}

	cret = C.g_async_queue_timed_pop(arg0, endTime)

	ret1 = C.gpointer(cret)

	return ret1
}

// TimedPopUnlocked pops data from the @queue. If the queue is empty, blocks
// until @end_time or until data becomes available.
//
// If no data is received before @end_time, nil is returned.
//
// To easily calculate @end_time, a combination of g_get_real_time() and
// g_time_val_add() can be used.
//
// This function must be called while holding the @queue's lock.
func (q *AsyncQueue) TimedPopUnlocked(endTime *TimeVal) interface{} {
	var arg0 *C.GAsyncQueue
	var arg1 *C.GTimeVal

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))
	arg1 = (*C.GTimeVal)(unsafe.Pointer(endTime.Native()))

	var cret C.gpointer
	var ret1 interface{}

	cret = C.g_async_queue_timed_pop_unlocked(arg0, endTime)

	ret1 = C.gpointer(cret)

	return ret1
}

// TimeoutPop pops data from the @queue. If the queue is empty, blocks for
// @timeout microseconds, or until data becomes available.
//
// If no data is received before the timeout, nil is returned.
func (q *AsyncQueue) TimeoutPop(timeout uint64) interface{} {
	var arg0 *C.GAsyncQueue
	var arg1 C.guint64

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))
	arg1 = C.guint64(timeout)

	var cret C.gpointer
	var ret1 interface{}

	cret = C.g_async_queue_timeout_pop(arg0, timeout)

	ret1 = C.gpointer(cret)

	return ret1
}

// TimeoutPopUnlocked pops data from the @queue. If the queue is empty, blocks
// for @timeout microseconds, or until data becomes available.
//
// If no data is received before the timeout, nil is returned.
//
// This function must be called while holding the @queue's lock.
func (q *AsyncQueue) TimeoutPopUnlocked(timeout uint64) interface{} {
	var arg0 *C.GAsyncQueue
	var arg1 C.guint64

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))
	arg1 = C.guint64(timeout)

	var cret C.gpointer
	var ret1 interface{}

	cret = C.g_async_queue_timeout_pop_unlocked(arg0, timeout)

	ret1 = C.gpointer(cret)

	return ret1
}

// TryPop tries to pop data from the @queue. If no data is available, nil is
// returned.
func (q *AsyncQueue) TryPop() interface{} {
	var arg0 *C.GAsyncQueue

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))

	var cret C.gpointer
	var ret1 interface{}

	cret = C.g_async_queue_try_pop(arg0)

	ret1 = C.gpointer(cret)

	return ret1
}

// TryPopUnlocked tries to pop data from the @queue. If no data is available,
// nil is returned.
//
// This function must be called while holding the @queue's lock.
func (q *AsyncQueue) TryPopUnlocked() interface{} {
	var arg0 *C.GAsyncQueue

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))

	var cret C.gpointer
	var ret1 interface{}

	cret = C.g_async_queue_try_pop_unlocked(arg0)

	ret1 = C.gpointer(cret)

	return ret1
}

// Unlock releases the queue's lock.
//
// Calling this function when you have not acquired the with
// g_async_queue_lock() leads to undefined behaviour.
func (q *AsyncQueue) Unlock() {
	var arg0 *C.GAsyncQueue

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))

	C.g_async_queue_unlock(arg0)
}

// Unref decreases the reference count of the asynchronous @queue by 1.
//
// If the reference count went to 0, the @queue will be destroyed and the memory
// allocated will be freed. So you are not allowed to use the @queue afterwards,
// as it might have disappeared. You do not need to hold the lock to call this
// function.
func (q *AsyncQueue) Unref() {
	var arg0 *C.GAsyncQueue

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))

	C.g_async_queue_unref(arg0)
}

// UnrefAndUnlock decreases the reference count of the asynchronous @queue by 1
// and releases the lock. This function must be called while holding the
// @queue's lock. If the reference count went to 0, the @queue will be destroyed
// and the memory allocated will be freed.
func (q *AsyncQueue) UnrefAndUnlock() {
	var arg0 *C.GAsyncQueue

	arg0 = (*C.GAsyncQueue)(unsafe.Pointer(q.Native()))

	C.g_async_queue_unref_and_unlock(arg0)
}
