// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// ShowURI: this function launches the default application for showing a given
// uri, or shows an error dialog if that fails.
//
// The function takes the following parameters:
//
//    - parent (optional) window.
//    - uri to show.
//    - timestamp from the event that triggered this call, or GDK_CURRENT_TIME.
//
func ShowURI(parent *Window, uri string, timestamp uint32) {
	var _arg1 *C.GtkWindow // out
	var _arg2 *C.char      // out
	var _arg3 C.guint32    // out

	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(externglib.InternObject(parent).Native()))
	}
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.guint32(timestamp)

	C.gtk_show_uri(_arg1, _arg2, _arg3)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(timestamp)
}

// ShowURIFull: this function launches the default application for showing a
// given uri.
//
// The callback will be called when the launch is completed. It should call
// gtk_show_uri_full_finish() to obtain the result.
//
// This is the recommended call to be used as it passes information necessary
// for sandbox helpers to parent their dialogs properly.
//
// The function takes the following parameters:
//
//    - ctx (optional) to cancel the launch.
//    - parent (optional) window.
//    - uri to show.
//    - timestamp from the event that triggered this call, or GDK_CURRENT_TIME.
//    - callback (optional) to call when the action is complete.
//
func ShowURIFull(ctx context.Context, parent *Window, uri string, timestamp uint32, callback gio.AsyncReadyCallback) {
	var _arg4 *C.GCancellable       // out
	var _arg1 *C.GtkWindow          // out
	var _arg2 *C.char               // out
	var _arg3 C.guint32             // out
	var _arg5 C.GAsyncReadyCallback // out
	var _arg6 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(externglib.InternObject(parent).Native()))
	}
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.guint32(timestamp)
	if callback != nil {
		_arg5 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg6 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.gtk_show_uri_full(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(timestamp)
	runtime.KeepAlive(callback)
}

// ShowURIFullFinish finishes the gtk_show_uri() call and returns the result of
// the operation.
//
// The function takes the following parameters:
//
//    - parent passed to gtk_show_uri().
//    - result that was passed to callback.
//
func ShowURIFullFinish(parent *Window, result gio.AsyncResulter) error {
	var _arg1 *C.GtkWindow    // out
	var _arg2 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(externglib.InternObject(parent).Native()))
	_arg2 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(result).Native()))

	C.gtk_show_uri_full_finish(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
