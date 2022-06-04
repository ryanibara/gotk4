// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gdkcontentformats.go.
var GTypeContentFormatsBuilder = coreglib.Type(C.gdk_content_formats_builder_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeContentFormatsBuilder, F: marshalContentFormatsBuilder},
	})
}

// InternMIMEType canonicalizes the given mime type and interns the result.
//
// If string is not a valid mime type, NULL is returned instead. See RFC 2048
// for the syntax if mime types.
//
// The function takes the following parameters:
//
//    - str: string of a potential mime type.
//
// The function returns the following values:
//
//    - utf8: interned string for the canonicalized mime type or NULL if the
//      string wasn't a valid mime type.
//
func InternMIMEType(str string) string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_args[0]))

	_gret := girepository.MustFind("Gdk", "intern_mime_type").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(str)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ContentFormatsBuilder: GdkContentFormatsBuilder is an auxiliary struct used
// to create new GdkContentFormats, and should not be kept around.
//
// An instance of this type is always passed by reference.
type ContentFormatsBuilder struct {
	*contentFormatsBuilder
}

// contentFormatsBuilder is the struct that's finalized.
type contentFormatsBuilder struct {
	native unsafe.Pointer
}

func marshalContentFormatsBuilder(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &ContentFormatsBuilder{&contentFormatsBuilder{(unsafe.Pointer)(b)}}, nil
}

// NewContentFormatsBuilder constructs a struct ContentFormatsBuilder.
func NewContentFormatsBuilder() *ContentFormatsBuilder {
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _contentFormatsBuilder *ContentFormatsBuilder // out

	_contentFormatsBuilder = (*ContentFormatsBuilder)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormatsBuilder)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_content_formats_builder_unref((*C.GdkContentFormatsBuilder)(intern.C))
		},
	)

	return _contentFormatsBuilder
}

// AddFormats appends all formats from formats to builder, skipping those that
// already exist.
//
// The function takes the following parameters:
//
//    - formats to add.
//
func (builder *ContentFormatsBuilder) AddFormats(formats *ContentFormats) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(builder)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(formats)))

	runtime.KeepAlive(builder)
	runtime.KeepAlive(formats)
}

// AddMIMEType appends mime_type to builder if it has not already been added.
//
// The function takes the following parameters:
//
//    - mimeType: mime type.
//
func (builder *ContentFormatsBuilder) AddMIMEType(mimeType string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(builder)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_args[1]))

	runtime.KeepAlive(builder)
	runtime.KeepAlive(mimeType)
}

// ToFormats creates a new GdkContentFormats from the given builder.
//
// The given GdkContentFormatsBuilder is reset once this function returns; you
// cannot call this function multiple times on the same builder instance.
//
// This function is intended primarily for bindings. C code should use
// gdk.ContentFormatsBuilder.FreeToFormats().
//
// The function returns the following values:
//
//    - contentFormats: newly created GdkContentFormats with all the formats
//      added to builder.
//
func (builder *ContentFormatsBuilder) ToFormats() *ContentFormats {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(builder)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(builder)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormats)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
		},
	)

	return _contentFormats
}
