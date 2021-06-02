// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
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
		{T: externglib.Type(C.g_icon_get_type()), F: marshalIcon},
	})
}

// IconDeserialize deserializes a #GIcon previously serialized using
// g_icon_serialize().
func IconDeserialize(value *glib.Variant) Icon {
	var arg1 *C.GVariant

	arg1 = (*C.GVariant)(value.Native())

	ret := C.g_icon_deserialize(arg1)

	var ret0 Icon

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(Icon)

	return ret0
}

// IconHash gets a hash for an icon.
func IconHash(icon interface{}) uint {
	var arg1 C.gpointer

	arg1 = C.gpointer(box.Assign(icon))

	ret := C.g_icon_hash(arg1)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// IconNewForString: generate a #GIcon instance from @str. This function can
// fail if @str is not valid - see g_icon_to_string() for discussion.
//
// If your application or library provides one or more #GIcon implementations
// you need to ensure that each #GType is registered with the type system prior
// to calling g_icon_new_for_string().
func IconNewForString(str string) (icon Icon, err error) {
	var arg1 *C.gchar
	var gError *C.GError

	arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.g_icon_new_for_string(arg1, &gError)

	var ret0 Icon
	var goError error

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(Icon)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// IconOverrider contains methods that are overridable. This
// interface is a subset of the interface Icon.
type IconOverrider interface {
	// Equal checks if two icons are equal.
	Equal(icon2 Icon) bool
	// Hash gets a hash for an icon.
	Hash() uint
	// Serialize serializes a #GIcon into a #GVariant. An equivalent #GIcon can
	// be retrieved back by calling g_icon_deserialize() on the returned value.
	// As serialization will avoid using raw icon data when possible, it only
	// makes sense to transfer the #GVariant between processes on the same
	// machine, (as opposed to over the network), and within the same file
	// system namespace.
	Serialize() *glib.Variant
	// ToTokens generates a textual representation of @icon that can be used for
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
	ToTokens(tokens []interface{}, outVersion int) bool
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
	IconOverrider

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

// WrapIcon wraps a GObject to a type that implements interface
// Icon. It is primarily used internally.
func WrapIcon(obj *externglib.Object) Icon {
	return Icon{
		Objector: obj,
	}
}

func marshalIcon(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapIcon(obj), nil
}

// Equal checks if two icons are equal.
func (i icon) Equal(icon2 Icon) bool {
	var arg0 *C.GIcon
	var arg1 *C.GIcon

	arg0 = (*C.GIcon)(i.Native())
	arg1 = (*C.GIcon)(icon2.Native())

	ret := C.g_icon_equal(arg0, arg1)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Serialize serializes a #GIcon into a #GVariant. An equivalent #GIcon can
// be retrieved back by calling g_icon_deserialize() on the returned value.
// As serialization will avoid using raw icon data when possible, it only
// makes sense to transfer the #GVariant between processes on the same
// machine, (as opposed to over the network), and within the same file
// system namespace.
func (i icon) Serialize() *glib.Variant {
	var arg0 *C.GIcon

	arg0 = (*C.GIcon)(i.Native())

	ret := C.g_icon_serialize(arg0)

	var ret0 *glib.Variant

	{
		ret0 = glib.WrapVariant(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *glib.Variant) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

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
func (i icon) String() string {
	var arg0 *C.GIcon

	arg0 = (*C.GIcon)(i.Native())

	ret := C.g_icon_to_string(arg0)

	var ret0 string

	ret0 = C.GoString(ret)
	C.free(unsafe.Pointer(ret))

	return ret0
}
