// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_simple_async_result_get_type()), F: marshalSimpleAsyncResulter},
	})
}

// SimpleAsyncReportGErrorInIdle reports an error in an idle function. Similar
// to g_simple_async_report_error_in_idle(), but takes a #GError rather than
// building a new one.
//
// Deprecated: Use g_task_report_error().
//
// The function takes the following parameters:
//
//    - object or NULL.
//    - callback: ReadyCallback.
//    - err to report.
//
func SimpleAsyncReportGErrorInIdle(object *externglib.Object, callback AsyncReadyCallback, err error) {
	var _arg1 *C.GObject            // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer
	var _arg4 *C.GError // out

	if object != nil {
		_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}
	_arg4 = (*C.GError)(gerror.New(err))

	C.g_simple_async_report_gerror_in_idle(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(object)
	runtime.KeepAlive(callback)
	runtime.KeepAlive(err)
}

// SimpleAsyncResult as of GLib 2.46, AsyncResult is deprecated in favor of
// #GTask, which provides a simpler API.
//
// AsyncResult implements Result.
//
// GSimpleAsyncResult handles ReadyCallbacks, error reporting, operation
// cancellation and the final state of an operation, completely transparent to
// the application. Results can be returned as a pointer e.g. for functions that
// return data that is collected asynchronously, a boolean value for checking
// the success or failure of an operation, or a #gssize for operations which
// return the number of bytes modified by the operation; all of the simple
// return cases are covered.
//
// Most of the time, an application will not need to know of the details of this
// API; it is handled transparently, and any necessary operations are handled by
// Result's interface. However, if implementing a new GIO module, for writing
// language bindings, or for complex applications that need better control of
// how asynchronous operations are completed, it is important to understand this
// functionality.
//
// GSimpleAsyncResults are tagged with the calling function to ensure that
// asynchronous functions and their finishing functions are used together
// correctly.
//
// To create a new AsyncResult, call g_simple_async_result_new(). If the result
// needs to be created for a #GError, use g_simple_async_result_new_from_error()
// or g_simple_async_result_new_take_error(). If a #GError is not available
// (e.g. the asynchronous operation's doesn't take a #GError argument), but the
// result still needs to be created for an error condition, use
// g_simple_async_result_new_error() (or g_simple_async_result_set_error_va() if
// your application or binding requires passing a variable argument list
// directly), and the error can then be propagated through the use of
// g_simple_async_result_propagate_error().
//
// An asynchronous operation can be made to ignore a cancellation event by
// calling g_simple_async_result_set_handle_cancellation() with a AsyncResult
// for the operation and FALSE. This is useful for operations that are dangerous
// to cancel, such as close (which would cause a leak if cancelled before being
// run).
//
// GSimpleAsyncResult can integrate into GLib's event loop, Loop, or it can use
// #GThreads. g_simple_async_result_complete() will finish an I/O task directly
// from the point where it is called. g_simple_async_result_complete_in_idle()
// will finish it from an idle handler in the [thread-default main
// context][g-main-context-push-thread-default] where the AsyncResult was
// created. g_simple_async_result_run_in_thread() will run the job in a separate
// thread and then use g_simple_async_result_complete_in_idle() to deliver the
// result.
//
// To set the results of an asynchronous function,
// g_simple_async_result_set_op_res_gpointer(),
// g_simple_async_result_set_op_res_gboolean(), and
// g_simple_async_result_set_op_res_gssize() are provided, setting the
// operation's result to a gpointer, gboolean, or gssize, respectively.
//
// Likewise, to get the result of an asynchronous function,
// g_simple_async_result_get_op_res_gpointer(),
// g_simple_async_result_get_op_res_gboolean(), and
// g_simple_async_result_get_op_res_gssize() are provided, getting the
// operation's result as a gpointer, gboolean, and gssize, respectively.
//
// For the details of the requirements implementations must respect, see Result.
// A typical implementation of an asynchronous operation using
// GSimpleAsyncResult looks something like this:
//
//    static void
//    baked_cb (Cake    *cake,
//              gpointer user_data)
//    {
//      // In this example, this callback is not given a reference to the cake,
//      // so the GSimpleAsyncResult has to take a reference to it.
//      GSimpleAsyncResult *result = user_data;
//
//      if (cake == NULL)
//        g_simple_async_result_set_error (result,
//                                         BAKER_ERRORS,
//                                         BAKER_ERROR_NO_FLOUR,
//                                         "Go to the supermarket");
//      else
//        g_simple_async_result_set_op_res_gpointer (result,
//                                                   g_object_ref (cake),
//                                                   g_object_unref);
//
//
//      // In this example, we assume that baked_cb is called as a callback from
//      // the mainloop, so it's safe to complete the operation synchronously here.
//      // If, however, _baker_prepare_cake () might call its callback without
//      // first returning to the mainloop — inadvisable, but some APIs do so —
//      // we would need to use g_simple_async_result_complete_in_idle().
//      g_simple_async_result_complete (result);
//      g_object_unref (result);
//    }
//
//    void
//    baker_bake_cake_async (Baker              *self,
//                           guint               radius,
//                           GAsyncReadyCallback callback,
//                           gpointer            user_data)
//    {
//      GSimpleAsyncResult *simple;
//      Cake               *cake;
//
//      if (radius < 3)
//        {
//          g_simple_async_report_error_in_idle (G_OBJECT (self),
//                                               callback,
//                                               user_data,
//                                               BAKER_ERRORS,
//                                               BAKER_ERROR_TOO_SMALL,
//                                               "ucm radius cakes are silly",
//                                               radius);
//          return;
//        }
//
//      simple = g_simple_async_result_new (G_OBJECT (self),
//                                          callback,
//                                          user_data,
//                                          baker_bake_cake_async);
//      cake = _baker_get_cached_cake (self, radius);
//
//      if (cake != NULL)
//        {
//          g_simple_async_result_set_op_res_gpointer (simple,
//                                                     g_object_ref (cake),
//                                                     g_object_unref);
//          g_simple_async_result_complete_in_idle (simple);
//          g_object_unref (simple);
//          // Drop the reference returned by _baker_get_cached_cake();
//          // the GSimpleAsyncResult has taken its own reference.
//          g_object_unref (cake);
//          return;
//        }
//
//      _baker_prepare_cake (self, radius, baked_cb, simple);
//    }
//
//    Cake *
//    baker_bake_cake_finish (Baker        *self,
//                            GAsyncResult *result,
//                            GError      **error)
//    {
//      GSimpleAsyncResult *simple;
//      Cake               *cake;
//
//      g_return_val_if_fail (g_simple_async_result_is_valid (result,
//                                                            G_OBJECT (self),
//                                                            baker_bake_cake_async),
//                            NULL);
//
//      simple = (GSimpleAsyncResult *) result;
//
//      if (g_simple_async_result_propagate_error (simple, error))
//        return NULL;
//
//      cake = CAKE (g_simple_async_result_get_op_res_gpointer (simple));
//      return g_object_ref (cake);
//    }.
type SimpleAsyncResult struct {
	*externglib.Object

	AsyncResult
}

func wrapSimpleAsyncResult(obj *externglib.Object) *SimpleAsyncResult {
	return &SimpleAsyncResult{
		Object: obj,
		AsyncResult: AsyncResult{
			Object: obj,
		},
	}
}

func marshalSimpleAsyncResulter(p uintptr) (interface{}, error) {
	return wrapSimpleAsyncResult(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSimpleAsyncResult creates a AsyncResult.
//
// The common convention is to create the AsyncResult in the function that
// starts the asynchronous operation and use that same function as the
// source_tag.
//
// If your operation supports cancellation with #GCancellable (which it probably
// should) then you should provide the user's cancellable to
// g_simple_async_result_set_check_cancellable() immediately after this function
// returns.
//
// Deprecated: Use g_task_new() instead.
//
// The function takes the following parameters:
//
//    - sourceObject or NULL.
//    - callback: ReadyCallback.
//    - sourceTag asynchronous function.
//
func NewSimpleAsyncResult(sourceObject *externglib.Object, callback AsyncReadyCallback, sourceTag cgo.Handle) *SimpleAsyncResult {
	var _arg1 *C.GObject            // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer
	var _arg4 C.gpointer            // out
	var _cret *C.GSimpleAsyncResult // in

	if sourceObject != nil {
		_arg1 = (*C.GObject)(unsafe.Pointer(sourceObject.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}
	_arg4 = (C.gpointer)(unsafe.Pointer(sourceTag))

	_cret = C.g_simple_async_result_new(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(sourceObject)
	runtime.KeepAlive(callback)
	runtime.KeepAlive(sourceTag)

	var _simpleAsyncResult *SimpleAsyncResult // out

	_simpleAsyncResult = wrapSimpleAsyncResult(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _simpleAsyncResult
}

// NewSimpleAsyncResultFromError creates a AsyncResult from an error condition.
//
// Deprecated: Use g_task_new() and g_task_return_error() instead.
//
// The function takes the following parameters:
//
//    - sourceObject or NULL.
//    - callback: ReadyCallback.
//    - err: #GError.
//
func NewSimpleAsyncResultFromError(sourceObject *externglib.Object, callback AsyncReadyCallback, err error) *SimpleAsyncResult {
	var _arg1 *C.GObject            // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer
	var _arg4 *C.GError             // out
	var _cret *C.GSimpleAsyncResult // in

	if sourceObject != nil {
		_arg1 = (*C.GObject)(unsafe.Pointer(sourceObject.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}
	_arg4 = (*C.GError)(gerror.New(err))

	_cret = C.g_simple_async_result_new_from_error(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(sourceObject)
	runtime.KeepAlive(callback)
	runtime.KeepAlive(err)

	var _simpleAsyncResult *SimpleAsyncResult // out

	_simpleAsyncResult = wrapSimpleAsyncResult(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _simpleAsyncResult
}

// Complete completes an asynchronous I/O job immediately. Must be called in the
// thread where the asynchronous result was to be delivered, as it invokes the
// callback directly. If you are in a different thread use
// g_simple_async_result_complete_in_idle().
//
// Calling this function takes a reference to simple for as long as is needed to
// complete the call.
//
// Deprecated: Use #GTask instead.
func (simple *SimpleAsyncResult) Complete() {
	var _arg0 *C.GSimpleAsyncResult // out

	_arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(simple.Native()))

	C.g_simple_async_result_complete(_arg0)
	runtime.KeepAlive(simple)
}

// CompleteInIdle completes an asynchronous function in an idle handler in the
// [thread-default main context][g-main-context-push-thread-default] of the
// thread that simple was initially created in (and re-pushes that context
// around the invocation of the callback).
//
// Calling this function takes a reference to simple for as long as is needed to
// complete the call.
//
// Deprecated: Use #GTask instead.
func (simple *SimpleAsyncResult) CompleteInIdle() {
	var _arg0 *C.GSimpleAsyncResult // out

	_arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(simple.Native()))

	C.g_simple_async_result_complete_in_idle(_arg0)
	runtime.KeepAlive(simple)
}

// OpResGboolean gets the operation result boolean from within the asynchronous
// result.
//
// Deprecated: Use #GTask and g_task_propagate_boolean() instead.
func (simple *SimpleAsyncResult) OpResGboolean() bool {
	var _arg0 *C.GSimpleAsyncResult // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(simple.Native()))

	_cret = C.g_simple_async_result_get_op_res_gboolean(_arg0)
	runtime.KeepAlive(simple)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// OpResGssize gets a gssize from the asynchronous result.
//
// Deprecated: Use #GTask and g_task_propagate_int() instead.
func (simple *SimpleAsyncResult) OpResGssize() int {
	var _arg0 *C.GSimpleAsyncResult // out
	var _cret C.gssize              // in

	_arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(simple.Native()))

	_cret = C.g_simple_async_result_get_op_res_gssize(_arg0)
	runtime.KeepAlive(simple)

	var _gssize int // out

	_gssize = int(_cret)

	return _gssize
}

// PropagateError propagates an error from within the simple asynchronous result
// to a given destination.
//
// If the #GCancellable given to a prior call to
// g_simple_async_result_set_check_cancellable() is cancelled then this function
// will return TRUE with dest set appropriately.
//
// Deprecated: Use #GTask instead.
func (simple *SimpleAsyncResult) PropagateError() error {
	var _arg0 *C.GSimpleAsyncResult // out
	var _cerr *C.GError             // in

	_arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(simple.Native()))

	C.g_simple_async_result_propagate_error(_arg0, &_cerr)
	runtime.KeepAlive(simple)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetCheckCancellable sets a #GCancellable to check before dispatching results.
//
// This function has one very specific purpose: the provided cancellable is
// checked at the time of g_simple_async_result_propagate_error() If it is
// cancelled, these functions will return an "Operation was cancelled" error
// (G_IO_ERROR_CANCELLED).
//
// Implementors of cancellable asynchronous functions should use this in order
// to provide a guarantee to their callers that cancelling an async operation
// will reliably result in an error being returned for that operation (even if a
// positive result for the operation has already been sent as an idle to the
// main context to be dispatched).
//
// The checking described above is done regardless of any call to the unrelated
// g_simple_async_result_set_handle_cancellation() function.
//
// Deprecated: Use #GTask instead.
//
// The function takes the following parameters:
//
//    - ctx to check, or NULL to unset.
//
func (simple *SimpleAsyncResult) SetCheckCancellable(ctx context.Context) {
	var _arg0 *C.GSimpleAsyncResult // out
	var _arg1 *C.GCancellable       // out

	_arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(simple.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	C.g_simple_async_result_set_check_cancellable(_arg0, _arg1)
	runtime.KeepAlive(simple)
	runtime.KeepAlive(ctx)
}

// SetFromError sets the result from a #GError.
//
// Deprecated: Use #GTask and g_task_return_error() instead.
//
// The function takes the following parameters:
//
//    - err: #GError.
//
func (simple *SimpleAsyncResult) SetFromError(err error) {
	var _arg0 *C.GSimpleAsyncResult // out
	var _arg1 *C.GError             // out

	_arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(simple.Native()))
	_arg1 = (*C.GError)(gerror.New(err))

	C.g_simple_async_result_set_from_error(_arg0, _arg1)
	runtime.KeepAlive(simple)
	runtime.KeepAlive(err)
}

// SetHandleCancellation sets whether to handle cancellation within the
// asynchronous operation.
//
// This function has nothing to do with
// g_simple_async_result_set_check_cancellable(). It only refers to the
// #GCancellable passed to g_simple_async_result_run_in_thread().
//
// Deprecated: since version 2.46.
//
// The function takes the following parameters:
//
//    - handleCancellation: #gboolean.
//
func (simple *SimpleAsyncResult) SetHandleCancellation(handleCancellation bool) {
	var _arg0 *C.GSimpleAsyncResult // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(simple.Native()))
	if handleCancellation {
		_arg1 = C.TRUE
	}

	C.g_simple_async_result_set_handle_cancellation(_arg0, _arg1)
	runtime.KeepAlive(simple)
	runtime.KeepAlive(handleCancellation)
}

// SetOpResGboolean sets the operation result to a boolean within the
// asynchronous result.
//
// Deprecated: Use #GTask and g_task_return_boolean() instead.
//
// The function takes the following parameters:
//
//    - opRes: #gboolean.
//
func (simple *SimpleAsyncResult) SetOpResGboolean(opRes bool) {
	var _arg0 *C.GSimpleAsyncResult // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(simple.Native()))
	if opRes {
		_arg1 = C.TRUE
	}

	C.g_simple_async_result_set_op_res_gboolean(_arg0, _arg1)
	runtime.KeepAlive(simple)
	runtime.KeepAlive(opRes)
}

// SetOpResGssize sets the operation result within the asynchronous result to
// the given op_res.
//
// Deprecated: Use #GTask and g_task_return_int() instead.
//
// The function takes the following parameters:
//
//    - opRes: #gssize.
//
func (simple *SimpleAsyncResult) SetOpResGssize(opRes int) {
	var _arg0 *C.GSimpleAsyncResult // out
	var _arg1 C.gssize              // out

	_arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(simple.Native()))
	_arg1 = C.gssize(opRes)

	C.g_simple_async_result_set_op_res_gssize(_arg0, _arg1)
	runtime.KeepAlive(simple)
	runtime.KeepAlive(opRes)
}

// SimpleAsyncResultIsValid ensures that the data passed to the _finish function
// of an async operation is consistent. Three checks are performed.
//
// First, result is checked to ensure that it is really a AsyncResult. Second,
// source is checked to ensure that it matches the source object of result.
// Third, source_tag is checked to ensure that it is equal to the source_tag
// argument given to g_simple_async_result_new() (which, by convention, is a
// pointer to the _async function corresponding to the _finish function from
// which this function is called). (Alternatively, if either source_tag or
// result's source tag is NULL, then the source tag check is skipped.)
//
// Deprecated: Use #GTask and g_task_is_valid() instead.
//
// The function takes the following parameters:
//
//    - result passed to the _finish function.
//    - source passed to the _finish function.
//    - sourceTag asynchronous function.
//
func SimpleAsyncResultIsValid(result AsyncResulter, source *externglib.Object, sourceTag cgo.Handle) bool {
	var _arg1 *C.GAsyncResult // out
	var _arg2 *C.GObject      // out
	var _arg3 C.gpointer      // out
	var _cret C.gboolean      // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))
	if source != nil {
		_arg2 = (*C.GObject)(unsafe.Pointer(source.Native()))
	}
	_arg3 = (C.gpointer)(unsafe.Pointer(sourceTag))

	_cret = C.g_simple_async_result_is_valid(_arg1, _arg2, _arg3)
	runtime.KeepAlive(result)
	runtime.KeepAlive(source)
	runtime.KeepAlive(sourceTag)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
