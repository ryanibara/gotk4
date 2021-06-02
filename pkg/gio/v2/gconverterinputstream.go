// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_converter_input_stream_get_type()), F: marshalConverterInputStream},
	})
}

// ConverterInputStream: converter input stream implements Stream and allows
// conversion of data of various types during reading.
//
// As of GLib 2.34, InputStream implements InputStream.
type ConverterInputStream interface {
	FilterInputStream
	PollableInputStream

	// Converter gets the #GConverter that is used by @converter_stream.
	Converter() Converter
}

// converterInputStream implements the ConverterInputStream interface.
type converterInputStream struct {
	FilterInputStream
	PollableInputStream
}

var _ ConverterInputStream = (*converterInputStream)(nil)

// WrapConverterInputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapConverterInputStream(obj *externglib.Object) ConverterInputStream {
	return ConverterInputStream{
		FilterInputStream:   WrapFilterInputStream(obj),
		PollableInputStream: WrapPollableInputStream(obj),
	}
}

func marshalConverterInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapConverterInputStream(obj), nil
}

// NewConverterInputStream constructs a class ConverterInputStream.
func NewConverterInputStream(baseStream InputStream, converter Converter) ConverterInputStream {
	var arg1 *C.GInputStream
	var arg2 *C.GConverter

	arg1 = (*C.GInputStream)(baseStream.Native())
	arg2 = (*C.GConverter)(converter.Native())

	ret := C.g_converter_input_stream_new(arg1, arg2)

	var ret0 ConverterInputStream

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(ConverterInputStream)

	return ret0
}

// Converter gets the #GConverter that is used by @converter_stream.
func (c converterInputStream) Converter() Converter {
	var arg0 *C.GConverterInputStream

	arg0 = (*C.GConverterInputStream)(c.Native())

	ret := C.g_converter_input_stream_get_converter(arg0)

	var ret0 Converter

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Converter)

	return ret0
}

type ConverterInputStreamPrivate struct {
	native C.GConverterInputStreamPrivate
}

// WrapConverterInputStreamPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapConverterInputStreamPrivate(ptr unsafe.Pointer) *ConverterInputStreamPrivate {
	if ptr == nil {
		return nil
	}

	return (*ConverterInputStreamPrivate)(ptr)
}

func marshalConverterInputStreamPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapConverterInputStreamPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *ConverterInputStreamPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}
