// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gerror"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

// ThreadPool: the Pool struct represents a thread pool. It has three public
// read-only members, but the underlying struct is bigger, so you must not copy
// this struct.
type ThreadPool struct {
	native C.GThreadPool
}

// WrapThreadPool wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapThreadPool(ptr unsafe.Pointer) *ThreadPool {
	if ptr == nil {
		return nil
	}

	return (*ThreadPool)(ptr)
}

func marshalThreadPool(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapThreadPool(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *ThreadPool) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// UserData gets the field inside the struct.
func (t *ThreadPool) UserData() interface{} {
	var v interface{}
	v = (interface{})(t.native.user_data)
	return v
}

// Exclusive gets the field inside the struct.
func (t *ThreadPool) Exclusive() bool {
	var v bool
	if t.native.exclusive {
		v = true
	}
	return v
}

// Free frees all resources allocated for @pool.
//
// If @immediate is true, no new task is processed for @pool. Otherwise @pool is
// not freed before the last task is processed. Note however, that no thread of
// this pool is interrupted while processing a task. Instead at least all still
// running threads can finish their tasks before the @pool is freed.
//
// If @wait_ is true, this function does not return before all tasks to be
// processed (dependent on @immediate, whether all or only the currently
// running) are ready. Otherwise this function returns immediately.
//
// After calling this function @pool must not be used anymore.
func (p *ThreadPool) Free(immediate bool, wait_ bool) {
	var _arg0 *C.GThreadPool
	var _arg1 C.gboolean
	var _arg2 C.gboolean

	_arg0 = (*C.GThreadPool)(unsafe.Pointer(p.Native()))
	if immediate {
		_arg1 = C.gboolean(1)
	}
	if wait_ {
		_arg2 = C.gboolean(1)
	}

	C.g_thread_pool_free(_arg0, _arg1, _arg2)
}

// MaxThreads returns the maximal number of threads for @pool.
func (p *ThreadPool) MaxThreads() int {
	var _arg0 *C.GThreadPool

	_arg0 = (*C.GThreadPool)(unsafe.Pointer(p.Native()))

	var _cret C.gint

	_cret = C.g_thread_pool_get_max_threads(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// NumThreads returns the number of threads currently running in @pool.
func (p *ThreadPool) NumThreads() uint {
	var _arg0 *C.GThreadPool

	_arg0 = (*C.GThreadPool)(unsafe.Pointer(p.Native()))

	var _cret C.guint

	_cret = C.g_thread_pool_get_num_threads(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// MoveToFront moves the item to the front of the queue of unprocessed items, so
// that it will be processed next.
func (p *ThreadPool) MoveToFront(data interface{}) bool {
	var _arg0 *C.GThreadPool
	var _arg1 C.gpointer

	_arg0 = (*C.GThreadPool)(unsafe.Pointer(p.Native()))
	_arg1 = C.gpointer(data)

	var _cret C.gboolean

	_cret = C.g_thread_pool_move_to_front(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Push inserts @data into the list of tasks to be executed by @pool.
//
// When the number of currently running threads is lower than the maximal
// allowed number of threads, a new thread is started (or reused) with the
// properties given to g_thread_pool_new(). Otherwise, @data stays in the queue
// until a thread in this pool finishes its previous task and processes @data.
//
// @error can be nil to ignore errors, or non-nil to report errors. An error can
// only occur when a new thread couldn't be created. In that case @data is
// simply appended to the queue of work to do.
//
// Before version 2.32, this function did not return a success status.
func (p *ThreadPool) Push(data interface{}) error {
	var _arg0 *C.GThreadPool
	var _arg1 C.gpointer

	_arg0 = (*C.GThreadPool)(unsafe.Pointer(p.Native()))
	_arg1 = C.gpointer(data)

	var _cerr *C.GError

	C.g_thread_pool_push(_arg0, _arg1, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetMaxThreads sets the maximal allowed number of threads for @pool. A value
// of -1 means that the maximal number of threads is unlimited. If @pool is an
// exclusive thread pool, setting the maximal number of threads to -1 is not
// allowed.
//
// Setting @max_threads to 0 means stopping all work for @pool. It is
// effectively frozen until @max_threads is set to a non-zero value again.
//
// A thread is never terminated while calling @func, as supplied by
// g_thread_pool_new(). Instead the maximal number of threads only has effect
// for the allocation of new threads in g_thread_pool_push(). A new thread is
// allocated, whenever the number of currently running threads in @pool is
// smaller than the maximal number.
//
// @error can be nil to ignore errors, or non-nil to report errors. An error can
// only occur when a new thread couldn't be created.
//
// Before version 2.32, this function did not return a success status.
func (p *ThreadPool) SetMaxThreads(maxThreads int) error {
	var _arg0 *C.GThreadPool
	var _arg1 C.gint

	_arg0 = (*C.GThreadPool)(unsafe.Pointer(p.Native()))
	_arg1 = C.gint(maxThreads)

	var _cerr *C.GError

	C.g_thread_pool_set_max_threads(_arg0, _arg1, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetSortFunction sets the function used to sort the list of tasks. This allows
// the tasks to be processed by a priority determined by @func, and not just in
// the order in which they were added to the pool.
//
// Note, if the maximum number of threads is more than 1, the order that threads
// are executed cannot be guaranteed 100%. Threads are scheduled by the
// operating system and are executed at random. It cannot be assumed that
// threads are executed in the order they are created.
func (p *ThreadPool) SetSortFunction(fn CompareDataFunc) {
	var _arg0 *C.GThreadPool
	var _arg1 C.GCompareDataFunc
	var _arg2 C.gpointer

	_arg0 = (*C.GThreadPool)(unsafe.Pointer(p.Native()))
	_arg1 = (*[0]byte)(C.gotk4_CompareDataFunc)
	_arg2 = C.gpointer(box.Assign(fn))

	C.g_thread_pool_set_sort_function(_arg0, _arg1, _arg2)
}

// Unprocessed returns the number of tasks still unprocessed in @pool.
func (p *ThreadPool) Unprocessed() uint {
	var _arg0 *C.GThreadPool

	_arg0 = (*C.GThreadPool)(unsafe.Pointer(p.Native()))

	var _cret C.guint

	_cret = C.g_thread_pool_unprocessed(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}
