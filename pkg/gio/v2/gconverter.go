// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern GConverterResult _gotk4_gio2_ConverterIface_convert(GConverter*, void*, gsize, void*, gsize, GConverterFlags, gsize*, gsize*, GError**);
// extern void _gotk4_gio2_ConverterIface_reset(GConverter*);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_converter_get_type()), F: marshalConverterer},
	})
}

// ConverterOverrider contains methods that are overridable.
type ConverterOverrider interface {
	// Convert: this is the main operation used when converting data. It is to
	// be called multiple times in a loop, and each time it will do some work,
	// i.e. producing some output (in outbuf) or consuming some input (from
	// inbuf) or both. If its not possible to do any work an error is returned.
	//
	// Note that a single call may not consume all input (or any input at all).
	// Also a call may produce output even if given no input, due to state
	// stored in the converter producing output.
	//
	// If any data was either produced or consumed, and then an error happens,
	// then only the successful conversion is reported and the error is returned
	// on the next call.
	//
	// A full conversion loop involves calling this method repeatedly, each time
	// giving it new input and space output space. When there is no more input
	// data after the data in inbuf, the flag G_CONVERTER_INPUT_AT_END must be
	// set. The loop will be (unless some error happens) returning
	// G_CONVERTER_CONVERTED each time until all data is consumed and all output
	// is produced, then G_CONVERTER_FINISHED is returned instead. Note, that
	// G_CONVERTER_FINISHED may be returned even if G_CONVERTER_INPUT_AT_END is
	// not set, for instance in a decompression converter where the end of data
	// is detectable from the data (and there might even be other data after the
	// end of the compressed data).
	//
	// When some data has successfully been converted bytes_read and is set to
	// the number of bytes read from inbuf, and bytes_written is set to indicate
	// how many bytes was written to outbuf. If there are more data to output or
	// consume (i.e. unless the G_CONVERTER_INPUT_AT_END is specified) then
	// G_CONVERTER_CONVERTED is returned, and if no more data is to be output
	// then G_CONVERTER_FINISHED is returned.
	//
	// On error G_CONVERTER_ERROR is returned and error is set accordingly. Some
	// errors need special handling:
	//
	// G_IO_ERROR_NO_SPACE is returned if there is not enough space to write the
	// resulting converted data, the application should call the function again
	// with a larger outbuf to continue.
	//
	// G_IO_ERROR_PARTIAL_INPUT is returned if there is not enough input to
	// fully determine what the conversion should produce, and the
	// G_CONVERTER_INPUT_AT_END flag is not set. This happens for example with
	// an incomplete multibyte sequence when converting text, or when a regexp
	// matches up to the end of the input (and may match further input). It may
	// also happen when inbuf_size is zero and there is no more data to produce.
	//
	// When this happens the application should read more input and then call
	// the function again. If further input shows that there is no more data
	// call the function again with the same data but with the
	// G_CONVERTER_INPUT_AT_END flag set. This may cause the conversion to
	// finish as e.g. in the regexp match case (or, to fail again with
	// G_IO_ERROR_PARTIAL_INPUT in e.g. a charset conversion where the input is
	// actually partial).
	//
	// After g_converter_convert() has returned G_CONVERTER_FINISHED the
	// converter object is in an invalid state where its not allowed to call
	// g_converter_convert() anymore. At this time you can only free the object
	// or call g_converter_reset() to reset it to the initial state.
	//
	// If the flag G_CONVERTER_FLUSH is set then conversion is modified to try
	// to write out all internal state to the output. The application has to
	// call the function multiple times with the flag set, and when the
	// available input has been consumed and all internal state has been
	// produced then G_CONVERTER_FLUSHED (or G_CONVERTER_FINISHED if really at
	// the end) is returned instead of G_CONVERTER_CONVERTED. This is somewhat
	// similar to what happens at the end of the input stream, but done in the
	// middle of the data.
	//
	// This has different meanings for different conversions. For instance in a
	// compression converter it would mean that we flush all the compression
	// state into output such that if you uncompress the compressed data you get
	// back all the input data. Doing this may make the final file larger due to
	// padding though. Another example is a regexp conversion, where if you at
	// the end of the flushed data have a match, but there is also a potential
	// longer match. In the non-flushed case we would ask for more input, but
	// when flushing we treat this as the end of input and do the match.
	//
	// Flushing is not always possible (like if a charset converter flushes at a
	// partial multibyte sequence). Converters are supposed to try to produce as
	// much output as possible and then return an error (typically
	// G_IO_ERROR_PARTIAL_INPUT).
	//
	// The function takes the following parameters:
	//
	//    - inbuf (optional): buffer containing the data to convert.
	//    - outbuf (optional): buffer to write converted data in.
	//    - flags controlling the conversion details.
	//
	// The function returns the following values:
	//
	//    - bytesRead will be set to the number of bytes read from inbuf on
	//      success.
	//    - bytesWritten will be set to the number of bytes written to outbuf on
	//      success.
	//    - converterResult G_CONVERTER_ERROR on error.
	//
	Convert(inbuf, outbuf []byte, flags ConverterFlags) (bytesRead uint, bytesWritten uint, converterResult ConverterResult, goerr error)
	// Reset resets all internal state in the converter, making it behave as if
	// it was just created. If the converter has any internal state that would
	// produce output then that output is lost.
	Reset()
}

// Converter is implemented by objects that convert binary data in various ways.
// The conversion can be stateful and may fail at any place.
//
// Some example conversions are: character set conversion, compression,
// decompression and regular expression replace.
type Converter struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Converter)(nil)
)

// Converterer describes Converter's interface methods.
type Converterer interface {
	externglib.Objector

	// Convert: this is the main operation used when converting data.
	Convert(inbuf, outbuf []byte, flags ConverterFlags) (bytesRead uint, bytesWritten uint, converterResult ConverterResult, goerr error)
	// Reset resets all internal state in the converter, making it behave as if
	// it was just created.
	Reset()
}

var _ Converterer = (*Converter)(nil)

func ifaceInitConverterer(gifacePtr, data C.gpointer) {
	iface := (*C.GConverterIface)(unsafe.Pointer(gifacePtr))
	iface.convert = (*[0]byte)(C._gotk4_gio2_ConverterIface_convert)
	iface.reset = (*[0]byte)(C._gotk4_gio2_ConverterIface_reset)
}

//export _gotk4_gio2_ConverterIface_convert
func _gotk4_gio2_ConverterIface_convert(arg0 *C.GConverter, arg1 *C.void, arg2 C.gsize, arg3 *C.void, arg4 C.gsize, arg5 C.GConverterFlags, arg6 *C.gsize, arg7 *C.gsize, _cerr **C.GError) (cret C.GConverterResult) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ConverterOverrider)

	var _inbuf []byte         // out
	var _outbuf []byte        // out
	var _flags ConverterFlags // out

	if arg1 != nil {
		_inbuf = make([]byte, arg2)
		copy(_inbuf, unsafe.Slice((*byte)(unsafe.Pointer(arg1)), arg2))
	}
	if arg3 != nil {
		_outbuf = make([]byte, arg4)
		copy(_outbuf, unsafe.Slice((*byte)(unsafe.Pointer(arg3)), arg4))
	}
	_flags = ConverterFlags(arg5)

	bytesRead, bytesWritten, converterResult, _goerr := iface.Convert(_inbuf, _outbuf, _flags)

	*arg6 = C.gsize(bytesRead)
	*arg7 = C.gsize(bytesWritten)
	cret = C.GConverterResult(converterResult)
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_ConverterIface_reset
func _gotk4_gio2_ConverterIface_reset(arg0 *C.GConverter) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ConverterOverrider)

	iface.Reset()
}

func wrapConverter(obj *externglib.Object) *Converter {
	return &Converter{
		Object: obj,
	}
}

func marshalConverterer(p uintptr) (interface{}, error) {
	return wrapConverter(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Convert: this is the main operation used when converting data. It is to be
// called multiple times in a loop, and each time it will do some work, i.e.
// producing some output (in outbuf) or consuming some input (from inbuf) or
// both. If its not possible to do any work an error is returned.
//
// Note that a single call may not consume all input (or any input at all). Also
// a call may produce output even if given no input, due to state stored in the
// converter producing output.
//
// If any data was either produced or consumed, and then an error happens, then
// only the successful conversion is reported and the error is returned on the
// next call.
//
// A full conversion loop involves calling this method repeatedly, each time
// giving it new input and space output space. When there is no more input data
// after the data in inbuf, the flag G_CONVERTER_INPUT_AT_END must be set. The
// loop will be (unless some error happens) returning G_CONVERTER_CONVERTED each
// time until all data is consumed and all output is produced, then
// G_CONVERTER_FINISHED is returned instead. Note, that G_CONVERTER_FINISHED may
// be returned even if G_CONVERTER_INPUT_AT_END is not set, for instance in a
// decompression converter where the end of data is detectable from the data
// (and there might even be other data after the end of the compressed data).
//
// When some data has successfully been converted bytes_read and is set to the
// number of bytes read from inbuf, and bytes_written is set to indicate how
// many bytes was written to outbuf. If there are more data to output or consume
// (i.e. unless the G_CONVERTER_INPUT_AT_END is specified) then
// G_CONVERTER_CONVERTED is returned, and if no more data is to be output then
// G_CONVERTER_FINISHED is returned.
//
// On error G_CONVERTER_ERROR is returned and error is set accordingly. Some
// errors need special handling:
//
// G_IO_ERROR_NO_SPACE is returned if there is not enough space to write the
// resulting converted data, the application should call the function again with
// a larger outbuf to continue.
//
// G_IO_ERROR_PARTIAL_INPUT is returned if there is not enough input to fully
// determine what the conversion should produce, and the
// G_CONVERTER_INPUT_AT_END flag is not set. This happens for example with an
// incomplete multibyte sequence when converting text, or when a regexp matches
// up to the end of the input (and may match further input). It may also happen
// when inbuf_size is zero and there is no more data to produce.
//
// When this happens the application should read more input and then call the
// function again. If further input shows that there is no more data call the
// function again with the same data but with the G_CONVERTER_INPUT_AT_END flag
// set. This may cause the conversion to finish as e.g. in the regexp match case
// (or, to fail again with G_IO_ERROR_PARTIAL_INPUT in e.g. a charset conversion
// where the input is actually partial).
//
// After g_converter_convert() has returned G_CONVERTER_FINISHED the converter
// object is in an invalid state where its not allowed to call
// g_converter_convert() anymore. At this time you can only free the object or
// call g_converter_reset() to reset it to the initial state.
//
// If the flag G_CONVERTER_FLUSH is set then conversion is modified to try to
// write out all internal state to the output. The application has to call the
// function multiple times with the flag set, and when the available input has
// been consumed and all internal state has been produced then
// G_CONVERTER_FLUSHED (or G_CONVERTER_FINISHED if really at the end) is
// returned instead of G_CONVERTER_CONVERTED. This is somewhat similar to what
// happens at the end of the input stream, but done in the middle of the data.
//
// This has different meanings for different conversions. For instance in a
// compression converter it would mean that we flush all the compression state
// into output such that if you uncompress the compressed data you get back all
// the input data. Doing this may make the final file larger due to padding
// though. Another example is a regexp conversion, where if you at the end of
// the flushed data have a match, but there is also a potential longer match. In
// the non-flushed case we would ask for more input, but when flushing we treat
// this as the end of input and do the match.
//
// Flushing is not always possible (like if a charset converter flushes at a
// partial multibyte sequence). Converters are supposed to try to produce as
// much output as possible and then return an error (typically
// G_IO_ERROR_PARTIAL_INPUT).
//
// The function takes the following parameters:
//
//    - inbuf: buffer containing the data to convert.
//    - outbuf: buffer to write converted data in.
//    - flags controlling the conversion details.
//
// The function returns the following values:
//
//    - bytesRead will be set to the number of bytes read from inbuf on success.
//    - bytesWritten will be set to the number of bytes written to outbuf on
//      success.
//    - converterResult G_CONVERTER_ERROR on error.
//
func (converter *Converter) Convert(inbuf, outbuf []byte, flags ConverterFlags) (bytesRead uint, bytesWritten uint, converterResult ConverterResult, goerr error) {
	var _arg0 *C.GConverter // out
	var _arg1 *C.void       // out
	var _arg2 C.gsize
	var _arg3 *C.void // out
	var _arg4 C.gsize
	var _arg5 C.GConverterFlags  // out
	var _arg6 C.gsize            // in
	var _arg7 C.gsize            // in
	var _cret C.GConverterResult // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GConverter)(unsafe.Pointer(converter.Native()))
	_arg2 = (C.gsize)(len(inbuf))
	if len(inbuf) > 0 {
		_arg1 = (*C.void)(unsafe.Pointer(&inbuf[0]))
	}
	_arg4 = (C.gsize)(len(outbuf))
	if len(outbuf) > 0 {
		_arg3 = (*C.void)(unsafe.Pointer(&outbuf[0]))
	}
	_arg5 = C.GConverterFlags(flags)

	_cret = C.g_converter_convert(_arg0, unsafe.Pointer(_arg1), _arg2, unsafe.Pointer(_arg3), _arg4, _arg5, &_arg6, &_arg7, &_cerr)
	runtime.KeepAlive(converter)
	runtime.KeepAlive(inbuf)
	runtime.KeepAlive(outbuf)
	runtime.KeepAlive(flags)

	var _bytesRead uint                  // out
	var _bytesWritten uint               // out
	var _converterResult ConverterResult // out
	var _goerr error                     // out

	_bytesRead = uint(_arg6)
	_bytesWritten = uint(_arg7)
	_converterResult = ConverterResult(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _bytesRead, _bytesWritten, _converterResult, _goerr
}

// Reset resets all internal state in the converter, making it behave as if it
// was just created. If the converter has any internal state that would produce
// output then that output is lost.
func (converter *Converter) Reset() {
	var _arg0 *C.GConverter // out

	_arg0 = (*C.GConverter)(unsafe.Pointer(converter.Native()))

	C.g_converter_reset(_arg0)
	runtime.KeepAlive(converter)
}
