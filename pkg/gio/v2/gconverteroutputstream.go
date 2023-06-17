// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeConverterOutputStream = coreglib.Type(C.g_converter_output_stream_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeConverterOutputStream, F: marshalConverterOutputStream},
	})
}

// ConverterOutputStreamOverrides contains methods that are overridable.
type ConverterOutputStreamOverrides struct {
}

func defaultConverterOutputStreamOverrides(v *ConverterOutputStream) ConverterOutputStreamOverrides {
	return ConverterOutputStreamOverrides{}
}

// ConverterOutputStream: converter output stream implements Stream and allows
// conversion of data of various types during reading.
//
// As of GLib 2.34, OutputStream implements OutputStream.
type ConverterOutputStream struct {
	_ [0]func() // equal guard
	FilterOutputStream

	*coreglib.Object
	OutputStream
	PollableOutputStream
}

var (
	_ FilterOutputStreamer = (*ConverterOutputStream)(nil)
	_ coreglib.Objector    = (*ConverterOutputStream)(nil)
	_ OutputStreamer       = (*ConverterOutputStream)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ConverterOutputStream, *ConverterOutputStreamClass, ConverterOutputStreamOverrides](
		GTypeConverterOutputStream,
		initConverterOutputStreamClass,
		wrapConverterOutputStream,
		defaultConverterOutputStreamOverrides,
	)
}

func initConverterOutputStreamClass(gclass unsafe.Pointer, overrides ConverterOutputStreamOverrides, classInitFunc func(*ConverterOutputStreamClass)) {
	if classInitFunc != nil {
		class := (*ConverterOutputStreamClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapConverterOutputStream(obj *coreglib.Object) *ConverterOutputStream {
	return &ConverterOutputStream{
		FilterOutputStream: FilterOutputStream{
			OutputStream: OutputStream{
				Object: obj,
			},
		},
		Object: obj,
		OutputStream: OutputStream{
			Object: obj,
		},
		PollableOutputStream: PollableOutputStream{
			OutputStream: OutputStream{
				Object: obj,
			},
		},
	}
}

func marshalConverterOutputStream(p uintptr) (interface{}, error) {
	return wrapConverterOutputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewConverterOutputStream creates a new converter output stream for the
// base_stream.
//
// The function takes the following parameters:
//
//   - baseStream: Stream.
//   - converter: #GConverter.
//
// The function returns the following values:
//
//   - converterOutputStream: new Stream.
//
func NewConverterOutputStream(baseStream OutputStreamer, converter Converterer) *ConverterOutputStream {
	var _arg1 *C.GOutputStream // out
	var _arg2 *C.GConverter    // out
	var _cret *C.GOutputStream // in

	_arg1 = (*C.GOutputStream)(unsafe.Pointer(coreglib.InternObject(baseStream).Native()))
	_arg2 = (*C.GConverter)(unsafe.Pointer(coreglib.InternObject(converter).Native()))

	_cret = C.g_converter_output_stream_new(_arg1, _arg2)
	runtime.KeepAlive(baseStream)
	runtime.KeepAlive(converter)

	var _converterOutputStream *ConverterOutputStream // out

	_converterOutputStream = wrapConverterOutputStream(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _converterOutputStream
}

// Converter gets the #GConverter that is used by converter_stream.
//
// The function returns the following values:
//
//   - converter of the converter output stream.
//
func (converterStream *ConverterOutputStream) Converter() *Converter {
	var _arg0 *C.GConverterOutputStream // out
	var _cret *C.GConverter             // in

	_arg0 = (*C.GConverterOutputStream)(unsafe.Pointer(coreglib.InternObject(converterStream).Native()))

	_cret = C.g_converter_output_stream_get_converter(_arg0)
	runtime.KeepAlive(converterStream)

	var _converter *Converter // out

	_converter = wrapConverter(coreglib.Take(unsafe.Pointer(_cret)))

	return _converter
}

// ConverterOutputStreamClass: instance of this type is always passed by
// reference.
type ConverterOutputStreamClass struct {
	*converterOutputStreamClass
}

// converterOutputStreamClass is the struct that's finalized.
type converterOutputStreamClass struct {
	native *C.GConverterOutputStreamClass
}

func (c *ConverterOutputStreamClass) ParentClass() *FilterOutputStreamClass {
	valptr := &c.native.parent_class
	var _v *FilterOutputStreamClass // out
	_v = (*FilterOutputStreamClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
