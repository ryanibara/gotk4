// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

// ContentDeserializeAsync: read content from the given input stream and
// deserialize it, asynchronously.
//
// The default I/O priority is G_PRIORITY_DEFAULT (i.e. 0), and lower numbers
// indicate a higher priority.
//
// When the operation is finished, @callback will be called. You must then call
// [func@content_deserialize_finish] to get the result of the operation.
func ContentDeserializeAsync(stream gio.InputStream, mimeType string, typ externglib.Type, ioPriority int, cancellable gio.Cancellable, callback gio.AsyncReadyCallback) {
	var _arg1 *C.GInputStream
	var _arg2 *C.char
	var _arg3 C.GType
	var _arg4 C.int
	var _arg5 *C.GCancellable
	var _arg6 C.GAsyncReadyCallback
	var _arg7 C.gpointer

	_arg1 = (*C.GInputStream)(unsafe.Pointer(stream.Native()))
	_arg2 = (*C.char)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.GType(typ)
	_arg4 = C.int(ioPriority)
	_arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg6 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg7 = C.gpointer(box.Assign(callback))

	C.gdk_content_deserialize_async(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
}

// ContentDeserializeFinish finishes a content deserialization operation.
func ContentDeserializeFinish(result gio.AsyncResult, value **externglib.Value) error {
	var _arg1 *C.GAsyncResult
	var _arg2 *C.GValue

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))
	_arg2 = (*C.GValue)(value.GValue)

	var _cerr *C.GError

	C.gdk_content_deserialize_finish(_arg1, _arg2, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
