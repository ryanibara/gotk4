// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gobject/v2"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// ThreadsAddIdleFull adds a function to be called whenever there are no higher
// priority events pending. If the function returns false it is automatically
// removed from the list of event sources and will not be called again.
//
// This variant of g_idle_add_full() calls @function with the GDK lock held. It
// can be thought of a MT-safe version for GTK+ widgets for the following use
// case, where you have to worry about idle_callback() running in thread A and
// accessing @self after it has been finalized in thread B:
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
//    }
func ThreadsAddIdleFull(priority int, function glib.SourceFunc) uint {

	var cret C.guint
	var ret1 uint

	cret = C.gdk_threads_add_idle_full(priority, function, data, notify)

	ret1 = C.guint(cret)

	return ret1
}

// ThreadsAddTimeoutFull sets a function to be called at regular intervals
// holding the GDK lock, with the given priority. The function is called
// repeatedly until it returns false, at which point the timeout is
// automatically destroyed and the function will not be called again. The
// @notify function is called when the timeout is destroyed. The first call to
// the function will be at the end of the first @interval.
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
//    }
func ThreadsAddTimeoutFull(priority int, interval uint, function glib.SourceFunc) uint {

	var cret C.guint
	var ret1 uint

	cret = C.gdk_threads_add_timeout_full(priority, interval, function, data, notify)

	ret1 = C.guint(cret)

	return ret1
}

// ThreadsAddTimeoutSecondsFull: a variant of gdk_threads_add_timeout_full()
// with second-granularity. See g_timeout_add_seconds_full() for a discussion of
// why it is a good idea to use this function if you don’t need finer
// granularity.
func ThreadsAddTimeoutSecondsFull(priority int, interval uint, function glib.SourceFunc) uint {

	var cret C.guint
	var ret1 uint

	cret = C.gdk_threads_add_timeout_seconds_full(priority, interval, function, data, notify)

	ret1 = C.guint(cret)

	return ret1
}

// ThreadsEnter: this function marks the beginning of a critical section in
// which GDK and GTK+ functions can be called safely and without causing race
// conditions. Only one thread at a time can be in such a critial section.
func ThreadsEnter() {
	C.gdk_threads_enter()
}

// ThreadsInit initializes GDK so that it can be used from multiple threads in
// conjunction with gdk_threads_enter() and gdk_threads_leave().
//
// This call must be made before any use of the main loop from GTK+; to be safe,
// call it before gtk_init().
func ThreadsInit() {
	C.gdk_threads_init()
}

// ThreadsLeave leaves a critical region begun with gdk_threads_enter().
func ThreadsLeave() {
	C.gdk_threads_leave()
}
