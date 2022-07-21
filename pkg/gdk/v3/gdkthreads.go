// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// extern gboolean _gotk4_gdk3_SourceFunc(gpointer);
// extern gboolean _gotk4_glib2_SourceFunc(gpointer);
// extern void callbackDelete(gpointer);
import "C"

// ThreadsAddIdle adds a function to be called whenever there are no higher
// priority events pending. If the function returns FALSE it is automatically
// removed from the list of event sources and will not be called again.
//
// This variant of g_idle_add_full() calls function with the GDK lock held. It
// can be thought of a MT-safe version for GTK+ widgets for the following use
// case, where you have to worry about idle_callback() running in thread A and
// accessing self after it has been finalized in thread B:
//
//    static gboolean
//    idle_callback (gpointer data)
//    {
//       // gdk_threads_enter(); would be needed for g_idle_add()
//
//       SomeWidget *self = data;
//       // do stuff with self
//
//       self->idle_id = 0;
//
//       // gdk_threads_leave(); would be needed for g_idle_add()
//       return FALSE;
//    }
//
//    static void
//    some_widget_do_stuff_later (SomeWidget *self)
//    {
//       self->idle_id = gdk_threads_add_idle (idle_callback, self)
//       // using g_idle_add() here would require thread protection in the callback
//    }
//
//    static void
//    some_widget_finalize (GObject *object)
//    {
//       SomeWidget *self = SOME_WIDGET (object);
//       if (self->idle_id)
//         g_source_remove (self->idle_id);
//       G_OBJECT_CLASS (parent_class)->finalize (object);
//    }.
//
// The function takes the following parameters:
//
//    - priority of the idle source. Typically this will be in the range between
//      PRIORITY_DEFAULT_IDLE and PRIORITY_HIGH_IDLE.
//    - function to call.
//
// The function returns the following values:
//
//    - guint: ID (greater than 0) of the event source.
//
func ThreadsAddIdle(priority int32, function glib.SourceFunc) uint32 {
	var _arg1 C.gint        // out
	var _arg2 C.GSourceFunc // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify
	var _cret C.guint // in

	_arg1 = C.gint(priority)
	_arg2 = (*[0]byte)(C._gotk4_glib2_SourceFunc)
	_arg3 = C.gpointer(gbox.Assign(function))
	_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	_cret = C.gdk_threads_add_idle_full(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(priority)
	runtime.KeepAlive(function)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// ThreadsAddTimeout sets a function to be called at regular intervals holding
// the GDK lock, with the given priority. The function is called repeatedly
// until it returns FALSE, at which point the timeout is automatically destroyed
// and the function will not be called again. The notify function is called when
// the timeout is destroyed. The first call to the function will be at the end
// of the first interval.
//
// Note that timeout functions may be delayed, due to the processing of other
// event sources. Thus they should not be relied on for precise timing. After
// each call to the timeout function, the time of the next timeout is
// recalculated based on the current time and the given interval (it does not
// try to “catch up” time lost in delays).
//
// This variant of g_timeout_add_full() can be thought of a MT-safe version for
// GTK+ widgets for the following use case:
//
//    static gboolean timeout_callback (gpointer data)
//    {
//       SomeWidget *self = data;
//
//       // do stuff with self
//
//       self->timeout_id = 0;
//
//       return G_SOURCE_REMOVE;
//    }
//
//    static void some_widget_do_stuff_later (SomeWidget *self)
//    {
//       self->timeout_id = g_timeout_add (timeout_callback, self)
//    }
//
//    static void some_widget_finalize (GObject *object)
//    {
//       SomeWidget *self = SOME_WIDGET (object);
//
//       if (self->timeout_id)
//         g_source_remove (self->timeout_id);
//
//       G_OBJECT_CLASS (parent_class)->finalize (object);
//    }.
//
// The function takes the following parameters:
//
//    - priority of the timeout source. Typically this will be in the range
//      between PRIORITY_DEFAULT_IDLE and PRIORITY_HIGH_IDLE.
//    - interval: time between calls to the function, in milliseconds (1/1000ths
//      of a second).
//    - function to call.
//
// The function returns the following values:
//
//    - guint: ID (greater than 0) of the event source.
//
func ThreadsAddTimeout(priority int32, interval uint32, function glib.SourceFunc) uint32 {
	var _arg1 C.gint        // out
	var _arg2 C.guint       // out
	var _arg3 C.GSourceFunc // out
	var _arg4 C.gpointer
	var _arg5 C.GDestroyNotify
	var _cret C.guint // in

	_arg1 = C.gint(priority)
	_arg2 = C.guint(interval)
	_arg3 = (*[0]byte)(C._gotk4_glib2_SourceFunc)
	_arg4 = C.gpointer(gbox.Assign(function))
	_arg5 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	_cret = C.gdk_threads_add_timeout_full(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(priority)
	runtime.KeepAlive(interval)
	runtime.KeepAlive(function)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// ThreadsAddTimeoutSeconds: variant of gdk_threads_add_timeout_full() with
// second-granularity. See g_timeout_add_seconds_full() for a discussion of why
// it is a good idea to use this function if you don’t need finer granularity.
//
// The function takes the following parameters:
//
//    - priority of the timeout source. Typically this will be in the range
//      between PRIORITY_DEFAULT_IDLE and PRIORITY_HIGH_IDLE.
//    - interval: time between calls to the function, in seconds.
//    - function to call.
//
// The function returns the following values:
//
//    - guint: ID (greater than 0) of the event source.
//
func ThreadsAddTimeoutSeconds(priority int32, interval uint32, function glib.SourceFunc) uint32 {
	var _arg1 C.gint        // out
	var _arg2 C.guint       // out
	var _arg3 C.GSourceFunc // out
	var _arg4 C.gpointer
	var _arg5 C.GDestroyNotify
	var _cret C.guint // in

	_arg1 = C.gint(priority)
	_arg2 = C.guint(interval)
	_arg3 = (*[0]byte)(C._gotk4_glib2_SourceFunc)
	_arg4 = C.gpointer(gbox.Assign(function))
	_arg5 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	_cret = C.gdk_threads_add_timeout_seconds_full(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(priority)
	runtime.KeepAlive(interval)
	runtime.KeepAlive(function)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// ThreadsEnter: this function marks the beginning of a critical section in
// which GDK and GTK+ functions can be called safely and without causing race
// conditions. Only one thread at a time can be in such a critial section.
//
// Deprecated: All GDK and GTK+ calls should be made from the main thread.
func ThreadsEnter() {
	C.gdk_threads_enter()
}

// ThreadsInit initializes GDK so that it can be used from multiple threads in
// conjunction with gdk_threads_enter() and gdk_threads_leave().
//
// This call must be made before any use of the main loop from GTK+; to be safe,
// call it before gtk_init().
//
// Deprecated: All GDK and GTK+ calls should be made from the main thread.
func ThreadsInit() {
	C.gdk_threads_init()
}

// ThreadsLeave leaves a critical region begun with gdk_threads_enter().
//
// Deprecated: All GDK and GTK+ calls should be made from the main thread.
func ThreadsLeave() {
	C.gdk_threads_leave()
}
