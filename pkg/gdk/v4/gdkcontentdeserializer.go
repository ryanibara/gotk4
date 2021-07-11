// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_content_deserializer_get_type()), F: marshalContentDeserializerer},
	})
}

// ContentDeserializeFinish finishes a content deserialization operation.
func ContentDeserializeFinish(result gio.AsyncResulter, value *externglib.Value) error {
	var _arg1 *C.GAsyncResult // out
	var _arg2 *C.GValue       // out
	var _cerr *C.GError       // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))
	_arg2 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	C.gdk_content_deserialize_finish(_arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// ContentDeserializerer describes ContentDeserializer's methods.
type ContentDeserializerer interface {
	// Cancellable gets the cancellable for the current operation.
	Cancellable() *gio.Cancellable
	// GType gets the GType to create an instance of.
	GType() externglib.Type
	// InputStream gets the input stream for the current operation.
	InputStream() *gio.InputStream
	// MIMEType gets the mime type to deserialize from.
	MIMEType() string
	// Priority gets the I/O priority for the current operation.
	Priority() int
	// TaskData gets the data that was associated with the current operation.
	TaskData() interface{}
	// UserData gets the user data that was passed when the deserializer was
	// registered.
	UserData() interface{}
	// Value gets the `GValue` to store the deserialized object in.
	Value() *externglib.Value
	// ReturnError: indicate that the deserialization has ended with an error.
	ReturnError(err error)
	// ReturnSuccess: indicate that the deserialization has been successfully
	// completed.
	ReturnSuccess()
}

// ContentDeserializer: `GdkContentDeserializer` is used to deserialize content
// received via inter-application data transfers.
//
// The `GdkContentDeserializer` transforms serialized content that is identified
// by a mime type into an object identified by a GType.
//
// GTK provides serializers and deserializers for common data types such as
// text, colors, images or file lists. To register your own deserialization
// functions, use [func@content_register_deserializer].
//
// Also see [class@Gdk.ContentSerializer].
type ContentDeserializer struct {
	*externglib.Object

	gio.AsyncResult
}

var (
	_ ContentDeserializerer = (*ContentDeserializer)(nil)
	_ gextras.Nativer       = (*ContentDeserializer)(nil)
)

func wrapContentDeserializer(obj *externglib.Object) ContentDeserializerer {
	return &ContentDeserializer{
		Object: obj,
		AsyncResult: gio.AsyncResult{
			Object: obj,
		},
	}
}

func marshalContentDeserializerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapContentDeserializer(obj), nil
}

// Cancellable gets the cancellable for the current operation.
//
// This is the `GCancellable` that was passed to
// [func@content_deserialize_async].
func (deserializer *ContentDeserializer) Cancellable() *gio.Cancellable {
	var _arg0 *C.GdkContentDeserializer // out
	var _cret *C.GCancellable           // in

	_arg0 = (*C.GdkContentDeserializer)(unsafe.Pointer(deserializer.Native()))

	_cret = C.gdk_content_deserializer_get_cancellable(_arg0)

	var _cancellable *gio.Cancellable // out

	_cancellable = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gio.Cancellable)

	return _cancellable
}

// GType gets the GType to create an instance of.
func (deserializer *ContentDeserializer) GType() externglib.Type {
	var _arg0 *C.GdkContentDeserializer // out
	var _cret C.GType                   // in

	_arg0 = (*C.GdkContentDeserializer)(unsafe.Pointer(deserializer.Native()))

	_cret = C.gdk_content_deserializer_get_gtype(_arg0)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// InputStream gets the input stream for the current operation.
//
// This is the stream that was passed to [func@content_deserialize_async].
func (deserializer *ContentDeserializer) InputStream() *gio.InputStream {
	var _arg0 *C.GdkContentDeserializer // out
	var _cret *C.GInputStream           // in

	_arg0 = (*C.GdkContentDeserializer)(unsafe.Pointer(deserializer.Native()))

	_cret = C.gdk_content_deserializer_get_input_stream(_arg0)

	var _inputStream *gio.InputStream // out

	_inputStream = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gio.InputStream)

	return _inputStream
}

// MIMEType gets the mime type to deserialize from.
func (deserializer *ContentDeserializer) MIMEType() string {
	var _arg0 *C.GdkContentDeserializer // out
	var _cret *C.char                   // in

	_arg0 = (*C.GdkContentDeserializer)(unsafe.Pointer(deserializer.Native()))

	_cret = C.gdk_content_deserializer_get_mime_type(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Priority gets the I/O priority for the current operation.
//
// This is the priority that was passed to [funccontent_deserialize_async].
func (deserializer *ContentDeserializer) Priority() int {
	var _arg0 *C.GdkContentDeserializer // out
	var _cret C.int                     // in

	_arg0 = (*C.GdkContentDeserializer)(unsafe.Pointer(deserializer.Native()))

	_cret = C.gdk_content_deserializer_get_priority(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// TaskData gets the data that was associated with the current operation.
//
// See [method@Gdk.ContentDeserializer.set_task_data].
func (deserializer *ContentDeserializer) TaskData() interface{} {
	var _arg0 *C.GdkContentDeserializer // out
	var _cret C.gpointer                // in

	_arg0 = (*C.GdkContentDeserializer)(unsafe.Pointer(deserializer.Native()))

	_cret = C.gdk_content_deserializer_get_task_data(_arg0)

	var _gpointer interface{} // out

	_gpointer = box.Get(uintptr(_cret))

	return _gpointer
}

// UserData gets the user data that was passed when the deserializer was
// registered.
func (deserializer *ContentDeserializer) UserData() interface{} {
	var _arg0 *C.GdkContentDeserializer // out
	var _cret C.gpointer                // in

	_arg0 = (*C.GdkContentDeserializer)(unsafe.Pointer(deserializer.Native()))

	_cret = C.gdk_content_deserializer_get_user_data(_arg0)

	var _gpointer interface{} // out

	_gpointer = box.Get(uintptr(_cret))

	return _gpointer
}

// Value gets the `GValue` to store the deserialized object in.
func (deserializer *ContentDeserializer) Value() *externglib.Value {
	var _arg0 *C.GdkContentDeserializer // out
	var _cret *C.GValue                 // in

	_arg0 = (*C.GdkContentDeserializer)(unsafe.Pointer(deserializer.Native()))

	_cret = C.gdk_content_deserializer_get_value(_arg0)

	var _value *externglib.Value // out

	_value = externglib.ValueFromNative(unsafe.Pointer(_cret))

	return _value
}

// ReturnError: indicate that the deserialization has ended with an error.
//
// This function consumes @error.
func (deserializer *ContentDeserializer) ReturnError(err error) {
	var _arg0 *C.GdkContentDeserializer // out
	var _arg1 *C.GError                 // out

	_arg0 = (*C.GdkContentDeserializer)(unsafe.Pointer(deserializer.Native()))
	_arg1 = (*C.GError)(gerror.New(err))

	C.gdk_content_deserializer_return_error(_arg0, _arg1)
}

// ReturnSuccess: indicate that the deserialization has been successfully
// completed.
func (deserializer *ContentDeserializer) ReturnSuccess() {
	var _arg0 *C.GdkContentDeserializer // out

	_arg0 = (*C.GdkContentDeserializer)(unsafe.Pointer(deserializer.Native()))

	C.gdk_content_deserializer_return_success(_arg0)
}
