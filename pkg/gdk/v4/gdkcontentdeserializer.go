// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// ContentDeserializeAsync: read content from the given input stream and
// deserialize it, asynchronously. When the operation is finished, @callback
// will be called. You can then call gdk_content_deserialize_finish() to get the
// result of the operation.
func ContentDeserializeAsync(stream gio.InputStream, mimeType string, typ externglib.Type, ioPriority int, cancellable gio.Cancellable, callback gio.AsyncReadyCallback) {

	C.gdk_content_deserialize_async(stream, mimeType, typ, ioPriority, cancellable, callback, userData)
}

// ContentDeserializeFinish finishes a content deserialization operation.
func ContentDeserializeFinish(result gio.AsyncResult, value *externglib.Value) error {
	var arg1 *C.GAsyncResult
	var arg2 *C.GValue

	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))
	arg2 = (*C.GValue)(value.GValue)

	var errout *C.GError
	var goerr error

	C.gdk_content_deserialize_finish(result, value, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))

	return goerr
}

// ContentRegisterDeserializer registers a function to create objects of a given
// @type from a serialized representation with the given mime type.
func ContentRegisterDeserializer(mimeType string, typ externglib.Type, deserialize ContentDeserializeFunc) {

	C.gdk_content_register_deserializer(mimeType, typ, deserialize, data, notify)
}
