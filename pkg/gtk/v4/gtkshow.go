// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

// ShowURI: this function launches the default application for showing a given
// uri, or shows an error dialog if that fails.
func ShowURI(parent *Window, uri string, timestamp uint32) {
	var _arg1 *C.GtkWindow // out
	var _arg2 *C.char      // out
	var _arg3 C.guint32    // out

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	_arg3 = C.guint32(timestamp)

	C.gtk_show_uri(_arg1, _arg2, _arg3)
}

// ShowURIFullFinish finishes the gtk_show_uri() call and returns the result of
// the operation.
func ShowURIFullFinish(parent *Window, result gio.AsyncResulter) error {
	var _arg1 *C.GtkWindow    // out
	var _arg2 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	_arg2 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	C.gtk_show_uri_full_finish(_arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
