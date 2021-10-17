// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_converter_input_stream_get_type()), F: marshalConverterInputStreamer},
	})
}

// ConverterInputStream: converter input stream implements Stream and allows
// conversion of data of various types during reading.
//
// As of GLib 2.34, InputStream implements InputStream.
type ConverterInputStream struct {
	FilterInputStream

	PollableInputStream
	*externglib.Object
}

func wrapConverterInputStream(obj *externglib.Object) *ConverterInputStream {
	return &ConverterInputStream{
		FilterInputStream: FilterInputStream{
			InputStream: InputStream{
				Object: obj,
			},
		},
		PollableInputStream: PollableInputStream{
			InputStream: InputStream{
				Object: obj,
			},
		},
		Object: obj,
	}
}

func marshalConverterInputStreamer(p uintptr) (interface{}, error) {
	return wrapConverterInputStream(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewConverterInputStream creates a new converter input stream for the
// base_stream.
//
// The function takes the following parameters:
//
//    - baseStream: Stream.
//    - converter: #GConverter.
//
func NewConverterInputStream(baseStream InputStreamer, converter Converterer) *ConverterInputStream {
	var _arg1 *C.GInputStream // out
	var _arg2 *C.GConverter   // out
	var _cret *C.GInputStream // in

	_arg1 = (*C.GInputStream)(unsafe.Pointer(baseStream.Native()))
	_arg2 = (*C.GConverter)(unsafe.Pointer(converter.Native()))

	_cret = C.g_converter_input_stream_new(_arg1, _arg2)
	runtime.KeepAlive(baseStream)
	runtime.KeepAlive(converter)

	var _converterInputStream *ConverterInputStream // out

	_converterInputStream = wrapConverterInputStream(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _converterInputStream
}

// Converter gets the #GConverter that is used by converter_stream.
func (converterStream *ConverterInputStream) Converter() Converterer {
	var _arg0 *C.GConverterInputStream // out
	var _cret *C.GConverter            // in

	_arg0 = (*C.GConverterInputStream)(unsafe.Pointer(converterStream.Native()))

	_cret = C.g_converter_input_stream_get_converter(_arg0)
	runtime.KeepAlive(converterStream)

	var _converter Converterer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Converterer is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Converterer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.Converterer")
		}
		_converter = rv
	}

	return _converter
}
