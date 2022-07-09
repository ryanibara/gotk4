// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gdk4_AsyncReadyCallback(void*, void*, gpointer);
// extern void _gotk4_gio2_AsyncReadyCallback(void*, void*, gpointer);
import "C"

// GTypeContentSerializer returns the GType for the type ContentSerializer.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeContentSerializer() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "ContentSerializer").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalContentSerializer)
	return gtype
}

// ContentSerializeAsync: serialize content and write it to the given output
// stream, asynchronously.
//
// The default I/O priority is G_PRIORITY_DEFAULT (i.e. 0), and lower numbers
// indicate a higher priority.
//
// When the operation is finished, callback will be called. You must then call
// content_serialize_finish to get the result of the operation.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object.
//    - stream: GOutputStream to write the serialized content to.
//    - mimeType: mime type to serialize to.
//    - value: content to serialize.
//    - ioPriority: i/O priority of the operation.
//    - callback (optional) to call when the operation is done.
//
func ContentSerializeAsync(ctx context.Context, stream gio.OutputStreamer, mimeType string, value *coreglib.Value, ioPriority int32, callback gio.AsyncReadyCallback) {
	var _args [7]girepository.Argument

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[4] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(value.Native()))
	*(*C.int)(unsafe.Pointer(&_args[3])) = C.int(ioPriority)
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[5])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[6] = C.gpointer(gbox.AssignOnce(callback))
	}

	girepository.MustFind("Gdk", "content_serialize_async").Invoke(_args[:], nil)

	runtime.KeepAlive(ctx)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(mimeType)
	runtime.KeepAlive(value)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// ContentSerializeFinish finishes a content serialization operation.
//
// The function takes the following parameters:
//
//    - result: GAsyncResult.
//
func ContentSerializeFinish(result gio.AsyncResulter) error {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	girepository.MustFind("Gdk", "content_serialize_finish").Invoke(_args[:], nil)

	runtime.KeepAlive(result)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// ContentSerializer: GdkContentSerializer is used to serialize content for
// inter-application data transfers.
//
// The GdkContentSerializer transforms an object that is identified by a GType
// into a serialized form (i.e. a byte stream) that is identified by a mime
// type.
//
// GTK provides serializers and deserializers for common data types such as
// text, colors, images or file lists. To register your own serialization
// functions, use content_register_serializer.
//
// Also see gdk.ContentDeserializer.
type ContentSerializer struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.AsyncResult
}

var (
	_ coreglib.Objector = (*ContentSerializer)(nil)
)

func wrapContentSerializer(obj *coreglib.Object) *ContentSerializer {
	return &ContentSerializer{
		Object: obj,
		AsyncResult: gio.AsyncResult{
			Object: obj,
		},
	}
}

func marshalContentSerializer(p uintptr) (interface{}, error) {
	return wrapContentSerializer(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Cancellable gets the cancellable for the current operation.
//
// This is the GCancellable that was passed to [content_serialize_async].
//
// The function returns the following values:
//
//    - cancellable for the current operation.
//
func (serializer *ContentSerializer) Cancellable() *gio.Cancellable {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(serializer).Native()))

	_gret := girepository.MustFind("Gdk", "ContentSerializer").InvokeMethod("get_cancellable", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(serializer)

	var _cancellable *gio.Cancellable // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_cancellable = &gio.Cancellable{
			Object: obj,
		}
	}

	return _cancellable
}

// MIMEType gets the mime type to serialize to.
//
// The function returns the following values:
//
//    - utf8: mime type for the current operation.
//
func (serializer *ContentSerializer) MIMEType() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(serializer).Native()))

	_gret := girepository.MustFind("Gdk", "ContentSerializer").InvokeMethod("get_mime_type", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(serializer)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// OutputStream gets the output stream for the current operation.
//
// This is the stream that was passed to content_serialize_async.
//
// The function returns the following values:
//
//    - outputStream: output stream for the current operation.
//
func (serializer *ContentSerializer) OutputStream() gio.OutputStreamer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(serializer).Native()))

	_gret := girepository.MustFind("Gdk", "ContentSerializer").InvokeMethod("get_output_stream", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(serializer)

	var _outputStream gio.OutputStreamer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.OutputStreamer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gio.OutputStreamer)
			return ok
		})
		rv, ok := casted.(gio.OutputStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.OutputStreamer")
		}
		_outputStream = rv
	}

	return _outputStream
}

// Priority gets the I/O priority for the current operation.
//
// This is the priority that was passed to content_serialize_async.
//
// The function returns the following values:
//
//    - gint: i/O priority for the current operation.
//
func (serializer *ContentSerializer) Priority() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(serializer).Native()))

	_gret := girepository.MustFind("Gdk", "ContentSerializer").InvokeMethod("get_priority", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(serializer)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// TaskData gets the data that was associated with the current operation.
//
// See gdk.ContentSerializer.SetTaskData().
//
// The function returns the following values:
//
//    - gpointer (optional): task data for serializer.
//
func (serializer *ContentSerializer) TaskData() unsafe.Pointer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(serializer).Native()))

	_gret := girepository.MustFind("Gdk", "ContentSerializer").InvokeMethod("get_task_data", _args[:], nil)
	_cret = *(*C.gpointer)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(serializer)

	var _gpointer unsafe.Pointer // out

	_gpointer = (unsafe.Pointer)(unsafe.Pointer(_cret))

	return _gpointer
}

// UserData gets the user data that was passed when the serializer was
// registered.
//
// The function returns the following values:
//
//    - gpointer (optional): user data for this serializer.
//
func (serializer *ContentSerializer) UserData() unsafe.Pointer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(serializer).Native()))

	_gret := girepository.MustFind("Gdk", "ContentSerializer").InvokeMethod("get_user_data", _args[:], nil)
	_cret = *(*C.gpointer)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(serializer)

	var _gpointer unsafe.Pointer // out

	_gpointer = (unsafe.Pointer)(unsafe.Pointer(_cret))

	return _gpointer
}

// Value gets the GValue to read the object to serialize from.
//
// The function returns the following values:
//
//    - value: GValue for the current operation.
//
func (serializer *ContentSerializer) Value() *coreglib.Value {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(serializer).Native()))

	_gret := girepository.MustFind("Gdk", "ContentSerializer").InvokeMethod("get_value", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(serializer)

	var _value *coreglib.Value // out

	_value = coreglib.ValueFromNative(unsafe.Pointer(_cret))

	return _value
}

// ReturnError: indicate that the serialization has ended with an error.
//
// This function consumes error.
//
// The function takes the following parameters:
//
//    - err: GError.
//
func (serializer *ContentSerializer) ReturnError(err error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(serializer).Native()))
	if err != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gerror.New(err))
	}

	girepository.MustFind("Gdk", "ContentSerializer").InvokeMethod("return_error", _args[:], nil)

	runtime.KeepAlive(serializer)
	runtime.KeepAlive(err)
}

// ReturnSuccess: indicate that the serialization has been successfully
// completed.
func (serializer *ContentSerializer) ReturnSuccess() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(serializer).Native()))

	girepository.MustFind("Gdk", "ContentSerializer").InvokeMethod("return_success", _args[:], nil)

	runtime.KeepAlive(serializer)
}
