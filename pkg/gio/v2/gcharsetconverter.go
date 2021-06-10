// Code generated by girgen. DO NOT EDIT.

package gio

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
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
		{T: externglib.Type(C.g_charset_converter_get_type()), F: marshalCharsetConverter},
	})
}

// CharsetConverter is an implementation of #GConverter based on GIConv.
type CharsetConverter interface {
	gextras.Objector
	Converter
	Initable

	// NumFallbacks gets the number of fallbacks that @converter has applied so
	// far.
	NumFallbacks() uint
	// UseFallback gets the Converter:use-fallback property.
	UseFallback() bool
	// SetUseFallback sets the Converter:use-fallback property.
	SetUseFallback(useFallback bool)
}

// charsetConverter implements the CharsetConverter interface.
type charsetConverter struct {
	gextras.Objector
	Converter
	Initable
}

var _ CharsetConverter = (*charsetConverter)(nil)

// WrapCharsetConverter wraps a GObject to the right type. It is
// primarily used internally.
func WrapCharsetConverter(obj *externglib.Object) CharsetConverter {
	return CharsetConverter{
		Objector:  obj,
		Converter: WrapConverter(obj),
		Initable:  WrapInitable(obj),
	}
}

func marshalCharsetConverter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCharsetConverter(obj), nil
}

// NumFallbacks gets the number of fallbacks that @converter has applied so
// far.
func (c charsetConverter) NumFallbacks() uint {
	var _arg0 *C.GCharsetConverter

	_arg0 = (*C.GCharsetConverter)(unsafe.Pointer(c.Native()))

	var _cret C.guint

	_cret = C.g_charset_converter_get_num_fallbacks(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// UseFallback gets the Converter:use-fallback property.
func (c charsetConverter) UseFallback() bool {
	var _arg0 *C.GCharsetConverter

	_arg0 = (*C.GCharsetConverter)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean

	_cret = C.g_charset_converter_get_use_fallback(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetUseFallback sets the Converter:use-fallback property.
func (c charsetConverter) SetUseFallback(useFallback bool) {
	var _arg0 *C.GCharsetConverter
	var _arg1 C.gboolean

	_arg0 = (*C.GCharsetConverter)(unsafe.Pointer(c.Native()))
	if useFallback {
		_arg1 = C.gboolean(1)
	}

	C.g_charset_converter_set_use_fallback(_arg0, _arg1)
}
