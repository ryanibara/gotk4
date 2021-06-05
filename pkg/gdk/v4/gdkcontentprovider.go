// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_content_provider_get_type()), F: marshalContentProvider},
	})
}

// ContentProvider: a GdkContentProvider is used to provide content for the
// clipboard in a number of formats.
//
// To create a GdkContentProvider, use gdk_content_provider_new_for_value() or
// gdk_content_provider_new_for_bytes().
//
// GDK knows how to handle common text and image formats out-of-the-box. See
// ContentSerializer and ContentDeserializer if you want to add support for
// application-specific data formats.
type ContentProvider interface {
	gextras.Objector

	// ContentChanged emits the ContentProvider::content-changed signal.
	ContentChanged()
	// Value gets the contents of @provider stored in @value.
	//
	// The @value will have been initialized to the #GType the value should be
	// provided in. This given #GType does not need to be listed in the formats
	// returned by gdk_content_provider_ref_formats(). However, if the given
	// #GType is not supported, this operation can fail and
	// IO_ERROR_NOT_SUPPORTED will be reported.
	Value(value *externglib.Value) error
	// RefFormats gets the formats that the provider can provide its current
	// contents in.
	RefFormats() *ContentFormats
	// RefStorableFormats gets the formats that the provider suggests other
	// applications to store the data in. An example of such an application
	// would be a clipboard manager.
	//
	// This can be assumed to be a subset of gdk_content_provider_ref_formats().
	RefStorableFormats() *ContentFormats
	// WriteMIMETypeAsync: asynchronously writes the contents of @provider to
	// @stream in the given @mime_type. When the operation is finished @callback
	// will be called. You can then call
	// gdk_content_provider_write_mime_type_finish() to get the result of the
	// operation.
	//
	// The given mime type does not need to be listed in the formats returned by
	// gdk_content_provider_ref_formats(). However, if the given #GType is not
	// supported, IO_ERROR_NOT_SUPPORTED will be reported.
	//
	// The given @stream will not be closed.
	WriteMIMETypeAsync(mimeType string, stream gio.OutputStream, ioPriority int, cancellable gio.Cancellable, callback gio.AsyncReadyCallback)
	// WriteMIMETypeFinish finishes an asynchronous write operation started with
	// gdk_content_provider_write_mime_type_async().
	WriteMIMETypeFinish(result gio.AsyncResult) error
}

// contentProvider implements the ContentProvider interface.
type contentProvider struct {
	gextras.Objector
}

var _ ContentProvider = (*contentProvider)(nil)

// WrapContentProvider wraps a GObject to the right type. It is
// primarily used internally.
func WrapContentProvider(obj *externglib.Object) ContentProvider {
	return ContentProvider{
		Objector: obj,
	}
}

func marshalContentProvider(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapContentProvider(obj), nil
}

// NewContentProviderForBytes constructs a class ContentProvider.
func NewContentProviderForBytes(mimeType string, bytes *glib.Bytes) ContentProvider {
	var arg1 *C.char
	var arg2 *C.GBytes

	arg1 = (*C.char)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GBytes)(unsafe.Pointer(bytes.Native()))

	var cret C.GdkContentProvider
	var ret1 ContentProvider

	cret = C.gdk_content_provider_new_for_bytes(mimeType, bytes)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ContentProvider)

	return ret1
}

// NewContentProviderForValue constructs a class ContentProvider.
func NewContentProviderForValue(value *externglib.Value) ContentProvider {
	var arg1 *C.GValue

	arg1 = (*C.GValue)(value.GValue)

	var cret C.GdkContentProvider
	var ret1 ContentProvider

	cret = C.gdk_content_provider_new_for_value(value)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ContentProvider)

	return ret1
}

// NewContentProviderUnion constructs a class ContentProvider.
func NewContentProviderUnion(providers []ContentProvider) ContentProvider {

	var cret C.GdkContentProvider
	var ret1 ContentProvider

	cret = C.gdk_content_provider_new_union(providers, nProviders)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ContentProvider)

	return ret1
}

// ContentChanged emits the ContentProvider::content-changed signal.
func (p contentProvider) ContentChanged() {
	var arg0 *C.GdkContentProvider

	arg0 = (*C.GdkContentProvider)(unsafe.Pointer(p.Native()))

	C.gdk_content_provider_content_changed(arg0)
}

// Value gets the contents of @provider stored in @value.
//
// The @value will have been initialized to the #GType the value should be
// provided in. This given #GType does not need to be listed in the formats
// returned by gdk_content_provider_ref_formats(). However, if the given
// #GType is not supported, this operation can fail and
// IO_ERROR_NOT_SUPPORTED will be reported.
func (p contentProvider) Value(value *externglib.Value) error {
	var arg0 *C.GdkContentProvider
	var arg1 *C.GValue
	var errout *C.GError

	arg0 = (*C.GdkContentProvider)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GValue)(value.GValue)

	var goerr error

	C.gdk_content_provider_get_value(arg0, value, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// RefFormats gets the formats that the provider can provide its current
// contents in.
func (p contentProvider) RefFormats() *ContentFormats {
	var arg0 *C.GdkContentProvider

	arg0 = (*C.GdkContentProvider)(unsafe.Pointer(p.Native()))

	var cret *C.GdkContentFormats
	var ret1 *ContentFormats

	cret = C.gdk_content_provider_ref_formats(arg0)

	ret1 = WrapContentFormats(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *ContentFormats) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// RefStorableFormats gets the formats that the provider suggests other
// applications to store the data in. An example of such an application
// would be a clipboard manager.
//
// This can be assumed to be a subset of gdk_content_provider_ref_formats().
func (p contentProvider) RefStorableFormats() *ContentFormats {
	var arg0 *C.GdkContentProvider

	arg0 = (*C.GdkContentProvider)(unsafe.Pointer(p.Native()))

	var cret *C.GdkContentFormats
	var ret1 *ContentFormats

	cret = C.gdk_content_provider_ref_storable_formats(arg0)

	ret1 = WrapContentFormats(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *ContentFormats) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// WriteMIMETypeAsync: asynchronously writes the contents of @provider to
// @stream in the given @mime_type. When the operation is finished @callback
// will be called. You can then call
// gdk_content_provider_write_mime_type_finish() to get the result of the
// operation.
//
// The given mime type does not need to be listed in the formats returned by
// gdk_content_provider_ref_formats(). However, if the given #GType is not
// supported, IO_ERROR_NOT_SUPPORTED will be reported.
//
// The given @stream will not be closed.
func (p contentProvider) WriteMIMETypeAsync(mimeType string, stream gio.OutputStream, ioPriority int, cancellable gio.Cancellable, callback gio.AsyncReadyCallback) {
	var arg0 *C.GdkContentProvider

	arg0 = (*C.GdkContentProvider)(unsafe.Pointer(p.Native()))

	C.gdk_content_provider_write_mime_type_async(arg0, mimeType, stream, ioPriority, cancellable, callback, userData)
}

// WriteMIMETypeFinish finishes an asynchronous write operation started with
// gdk_content_provider_write_mime_type_async().
func (p contentProvider) WriteMIMETypeFinish(result gio.AsyncResult) error {
	var arg0 *C.GdkContentProvider
	var arg1 *C.GAsyncResult
	var errout *C.GError

	arg0 = (*C.GdkContentProvider)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var goerr error

	C.gdk_content_provider_write_mime_type_finish(arg0, result, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}
