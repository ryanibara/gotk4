// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.g_icon_get_type()), F: marshalIcon},
	})
}

// Icon is a very minimal interface for icons. It provides functions for
// checking the equality of two icons, hashing of icons and serializing an icon
// to and from strings.
//
// #GIcon does not provide the actual pixmap for the icon as this is out of
// GIO's scope, however implementations of #GIcon may contain the name of an
// icon (see Icon), or the path to an icon (see Icon).
//
// To obtain a hash of a #GIcon, see g_icon_hash().
//
// To check if two #GIcons are equal, see g_icon_equal().
//
// For serializing a #GIcon, use g_icon_serialize() and g_icon_deserialize().
//
// If you want to consume #GIcon (for example, in a toolkit) you must be
// prepared to handle at least the three following cases: Icon, Icon and Icon.
// It may also make sense to have fast-paths for other cases (like handling
// Pixbuf directly, for example) but all compliant #GIcon implementations
// outside of GIO must implement Icon.
//
// If your application or library provides one or more #GIcon implementations
// you need to ensure that your new implementation also implements Icon.
// Additionally, you must provide an implementation of g_icon_serialize() that
// gives a result that is understood by g_icon_deserialize(), yielding one of
// the built-in icon types.
type Icon interface {
	gextras.Objector

	// Equal checks if two icons are equal.
	Equal(icon2 Icon) bool
	// Serialize serializes a #GIcon into a #GVariant. An equivalent #GIcon can
	// be retrieved back by calling g_icon_deserialize() on the returned value.
	// As serialization will avoid using raw icon data when possible, it only
	// makes sense to transfer the #GVariant between processes on the same
	// machine, (as opposed to over the network), and within the same file
	// system namespace.
	Serialize() glib.Variant
	// String generates a textual representation of @icon that can be used for
	// serialization such as when passing @icon to a different process or saving
	// it to persistent storage. Use g_icon_new_for_string() to get @icon back
	// from the returned string.
	//
	// The encoding of the returned string is proprietary to #GIcon except in
	// the following two cases
	//
	// - If @icon is a Icon, the returned string is a native path (such as
	// `/path/to/my icon.png`) without escaping if the #GFile for @icon is a
	// native file. If the file is not native, the returned string is the result
	// of g_file_get_uri() (such as `sftp://path/to/my20icon.png`).
	//
	// - If @icon is a Icon with exactly one name and no fallbacks, the encoding
	// is simply the name (such as `network-server`).
	String() string
}

// icon implements the Icon interface.
type icon struct {
	gextras.Objector
}

var _ Icon = (*icon)(nil)

// WrapIcon wraps a GObject to a type that implements
// interface Icon. It is primarily used internally.
func WrapIcon(obj *externglib.Object) Icon {
	return icon{
		Objector: obj,
	}
}

func marshalIcon(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapIcon(obj), nil
}

func (i icon) Equal(icon2 Icon) bool {
	var _arg0 *C.GIcon   // out
	var _arg1 *C.GIcon   // out
	var _cret C.gboolean // in

	_arg0 = (*C.GIcon)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(icon2.Native()))

	_cret = C.g_icon_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (i icon) Serialize() glib.Variant {
	var _arg0 *C.GIcon    // out
	var _cret *C.GVariant // in

	_arg0 = (*C.GIcon)(unsafe.Pointer(i.Native()))

	_cret = C.g_icon_serialize(_arg0)

	var _variant glib.Variant // out

	_variant = (glib.Variant)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_variant, func(v glib.Variant) {
		C.g_variant_unref((*C.GVariant)(unsafe.Pointer(v)))
	})

	return _variant
}

func (i icon) String() string {
	var _arg0 *C.GIcon // out
	var _cret *C.gchar // in

	_arg0 = (*C.GIcon)(unsafe.Pointer(i.Native()))

	_cret = C.g_icon_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
