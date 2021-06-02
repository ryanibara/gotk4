// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
//
// extern void callbackDelete(gpointer);
import "C"

//export callbackDelete
func callbackDelete(ptr C.gpointer) {
	box.Delete(box.Callback, uintptr(ptr))
}

// ContentRegisterSerializer registers a function to serialize objects of a
// given type.
func ContentRegisterSerializer(typ externglib.Type, mimeType string, serialize ContentSerializeFunc) {
	var arg1 C.GType
	var arg2 *C.char
	var arg3 C.GdkContentSerializeFunc
	var arg4 C.gpointer
	var arg5 C.GDestroyNotify

	arg1 = C.GType(typ)
	arg2 = (*C.gchar)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*[0]byte)(C.gotk4_ContentSerializeFunc)
	arg4 = C.gpointer(box.Assign(serialize))
	arg5 = (*[0]byte)(C.callbackDelete)

	C.gdk_content_register_serializer(arg1, arg2, arg3, arg4, arg5)
}

// ContentSerializeAsync: serialize content and write it to the given output
// stream, asynchronously.
//
// The default I/O priority is G_PRIORITY_DEFAULT (i.e. 0), and lower numbers
// indicate a higher priority.
//
// When the operation is finished, @callback will be called. You must then call
// [func@content_serialize_finish] to get the result of the operation.
func ContentSerializeAsync(stream gio.OutputStream, mimeType string, value *externglib.Value, ioPriority int, cancellable gio.Cancellable, callback gio.AsyncReadyCallback) {
	var arg1 *C.GOutputStream
	var arg2 *C.char
	var arg3 *C.GValue
	var arg4 C.int
	var arg5 *C.GCancellable
	var arg6 C.GAsyncReadyCallback
	var arg7 C.gpointer

	arg1 = (*C.GOutputStream)(stream.Native())
	arg2 = (*C.gchar)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.GValue)(value.GValue)
	arg4 = C.int(ioPriority)
	arg5 = (*C.GCancellable)(cancellable.Native())
	arg6 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	arg7 = C.gpointer(box.Assign(callback))

	C.gdk_content_serialize_async(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// ContentSerializeFinish finishes a content serialization operation.
func ContentSerializeFinish(result gio.AsyncResult) error {
	var arg1 *C.GAsyncResult
	var gError *C.GError

	arg1 = (*C.GAsyncResult)(result.Native())

	ret := C.gdk_content_serialize_finish(arg1, &gError)

	var goError error

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return goError
}
