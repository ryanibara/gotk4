// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_charset_converter_get_type()), F: marshalCharsetConverter},
	})
}

// CharsetConverter is an implementation of #GConverter based on GIConv.
type CharsetConverter interface {
	Converter
	Initable

	NumFallbacks() uint

	UseFallback() bool

	SetUseFallbackCharsetConverter(useFallback bool)
}

// charsetConverter implements the CharsetConverter class.
type charsetConverter struct {
	gextras.Objector
}

// WrapCharsetConverter wraps a GObject to the right type. It is
// primarily used internally.
func WrapCharsetConverter(obj *externglib.Object) CharsetConverter {
	return charsetConverter{
		Objector: obj,
	}
}

func marshalCharsetConverter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCharsetConverter(obj), nil
}

func NewCharsetConverter(toCharset string, fromCharset string) (CharsetConverter, error) {
	var _arg1 *C.gchar             // out
	var _arg2 *C.gchar             // out
	var _cret *C.GCharsetConverter // in
	var _cerr *C.GError            // in

	_arg1 = (*C.gchar)(C.CString(toCharset))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(fromCharset))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_charset_converter_new(_arg1, _arg2, &_cerr)

	var _charsetConverter CharsetConverter // out
	var _goerr error                       // out

	_charsetConverter = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(CharsetConverter)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _charsetConverter, _goerr
}

func (c charsetConverter) NumFallbacks() uint {
	var _arg0 *C.GCharsetConverter // out
	var _cret C.guint              // in

	_arg0 = (*C.GCharsetConverter)(unsafe.Pointer(c.Native()))

	_cret = C.g_charset_converter_get_num_fallbacks(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (c charsetConverter) UseFallback() bool {
	var _arg0 *C.GCharsetConverter // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GCharsetConverter)(unsafe.Pointer(c.Native()))

	_cret = C.g_charset_converter_get_use_fallback(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c charsetConverter) SetUseFallbackCharsetConverter(useFallback bool) {
	var _arg0 *C.GCharsetConverter // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GCharsetConverter)(unsafe.Pointer(c.Native()))
	if useFallback {
		_arg1 = C.TRUE
	}

	C.g_charset_converter_set_use_fallback(_arg0, _arg1)
}

func (c charsetConverter) Convert(inbuf []byte, outbuf []byte, flags ConverterFlags) (bytesRead uint, bytesWritten uint, converterResult ConverterResult, goerr error) {
	return WrapConverter(gextras.InternObject(c)).Convert(inbuf, outbuf, flags)
}

func (c charsetConverter) Reset() {
	WrapConverter(gextras.InternObject(c)).Reset()
}

func (i charsetConverter) Init(cancellable Cancellable) error {
	return WrapInitable(gextras.InternObject(i)).Init(cancellable)
}
