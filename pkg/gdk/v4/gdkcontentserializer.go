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

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_content_serializer_get_type()), F: marshalContentSerializer},
	})
}

// ContentSerializeFinish finishes a content serialization operation.
func ContentSerializeFinish(result gio.AsyncResult) error {
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.gdk_content_serialize_finish(_arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// ContentSerializer: a `GdkContentSerializer` is used to serialize content for
// inter-application data transfers.
//
// The `GdkContentSerializer` transforms an object that is identified by a GType
// into a serialized form (i.e. a byte stream) that is identified by a mime
// type.
//
// GTK provides serializers and deserializers for common data types such as
// text, colors, images or file lists. To register your own serialization
// functions, use [func@content_register_serializer].
//
// Also see [class@Gdk.ContentDeserializer].
type ContentSerializer interface {
	gextras.Objector
	gio.AsyncResult

	// Cancellable gets the cancellable for the current operation.
	//
	// This is the `GCancellable` that was passed to [content_serialize_async].
	Cancellable() gio.Cancellable
	// GType gets the `GType` to of the object to serialize.
	GType() externglib.Type
	// MIMEType gets the mime type to serialize to.
	MIMEType() string
	// OutputStream gets the output stream for the current operation.
	//
	// This is the stream that was passed to [func@content_serialize_async].
	OutputStream() gio.OutputStream
	// Priority gets the I/O priority for the current operation.
	//
	// This is the priority that was passed to [func@content_serialize_async].
	Priority() int
	// Value gets the `GValue` to read the object to serialize from.
	Value() **externglib.Value
	// ReturnError: indicate that the serialization has ended with an error.
	//
	// This function consumes @error.
	ReturnError(err error)
	// ReturnSuccess: indicate that the serialization has been successfully
	// completed.
	ReturnSuccess()
}

// contentSerializer implements the ContentSerializer class.
type contentSerializer struct {
	gextras.Objector
	gio.AsyncResult
}

var _ ContentSerializer = (*contentSerializer)(nil)

// WrapContentSerializer wraps a GObject to the right type. It is
// primarily used internally.
func WrapContentSerializer(obj *externglib.Object) ContentSerializer {
	return contentSerializer{
		Objector:        obj,
		gio.AsyncResult: gio.WrapAsyncResult(obj),
	}
}

func marshalContentSerializer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapContentSerializer(obj), nil
}

// Cancellable gets the cancellable for the current operation.
//
// This is the `GCancellable` that was passed to [content_serialize_async].
func (s contentSerializer) Cancellable() gio.Cancellable {
	var _arg0 *C.GdkContentSerializer // out
	var _cret *C.GCancellable         // in

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_content_serializer_get_cancellable(_arg0)

	var _cancellable gio.Cancellable // out

	_cancellable = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.Cancellable)

	return _cancellable
}

// GType gets the `GType` to of the object to serialize.
func (s contentSerializer) GType() externglib.Type {
	var _arg0 *C.GdkContentSerializer // out
	var _cret C.GType                 // in

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_content_serializer_get_gtype(_arg0)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// MIMEType gets the mime type to serialize to.
func (s contentSerializer) MIMEType() string {
	var _arg0 *C.GdkContentSerializer // out
	var _cret *C.char                 // in

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_content_serializer_get_mime_type(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// OutputStream gets the output stream for the current operation.
//
// This is the stream that was passed to [func@content_serialize_async].
func (s contentSerializer) OutputStream() gio.OutputStream {
	var _arg0 *C.GdkContentSerializer // out
	var _cret *C.GOutputStream        // in

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_content_serializer_get_output_stream(_arg0)

	var _outputStream gio.OutputStream // out

	_outputStream = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.OutputStream)

	return _outputStream
}

// Priority gets the I/O priority for the current operation.
//
// This is the priority that was passed to [func@content_serialize_async].
func (s contentSerializer) Priority() int {
	var _arg0 *C.GdkContentSerializer // out
	var _cret C.int                   // in

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_content_serializer_get_priority(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Value gets the `GValue` to read the object to serialize from.
func (s contentSerializer) Value() **externglib.Value {
	var _arg0 *C.GdkContentSerializer // out
	var _cret *C.GValue               // in

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_content_serializer_get_value(_arg0)

	var _value **externglib.Value // out

	_value = externglib.ValueFromNative(unsafe.Pointer(_cret))

	return _value
}

// ReturnError: indicate that the serialization has ended with an error.
//
// This function consumes @error.
func (s contentSerializer) ReturnError(err error) {
	var _arg0 *C.GdkContentSerializer // out
	var _arg1 *C.GError               // out

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GError)(gerror.New(unsafe.Pointer(err)))
	defer C.g_error_free(_arg1)

	C.gdk_content_serializer_return_error(_arg0, _arg1)
}

// ReturnSuccess: indicate that the serialization has been successfully
// completed.
func (s contentSerializer) ReturnSuccess() {
	var _arg0 *C.GdkContentSerializer // out

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(s.Native()))

	C.gdk_content_serializer_return_success(_arg0)
}
