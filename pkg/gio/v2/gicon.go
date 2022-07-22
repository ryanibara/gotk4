// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern GVariant* _gotk4_gio2_IconIface_serialize(GIcon*);
// extern gboolean _gotk4_gio2_IconIface_equal(GIcon*, GIcon*);
// extern guint _gotk4_gio2_IconIface_hash(GIcon*);
import "C"

// GType values.
var (
	GTypeIcon = coreglib.Type(C.g_icon_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeIcon, F: marshalIcon},
	})
}

// IconOverrider contains methods that are overridable.
type IconOverrider interface {
	// Equal checks if two icons are equal.
	//
	// The function takes the following parameters:
	//
	//    - icon2 (optional): pointer to the second #GIcon.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if icon1 is equal to icon2. FALSE otherwise.
	//
	Equal(icon2 Iconner) bool
	// Hash gets a hash for an icon.
	//
	// The function returns the following values:
	//
	//    - guint containing a hash for the icon, suitable for use in a Table or
	//      similar data structure.
	//
	Hash() uint
	// Serialize serializes a #GIcon into a #GVariant. An equivalent #GIcon can
	// be retrieved back by calling g_icon_deserialize() on the returned value.
	// As serialization will avoid using raw icon data when possible, it only
	// makes sense to transfer the #GVariant between processes on the same
	// machine, (as opposed to over the network), and within the same file
	// system namespace.
	//
	// The function returns the following values:
	//
	//    - variant (optional) or NULL when serialization fails. The #GVariant
	//      will not be floating.
	//
	Serialize() *glib.Variant
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
//
// Icon wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Icon struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Icon)(nil)
)

// Iconner describes Icon's interface methods.
type Iconner interface {
	coreglib.Objector

	// Equal checks if two icons are equal.
	Equal(icon2 Iconner) bool
	// Serialize serializes a #GIcon into a #GVariant.
	Serialize() *glib.Variant
	// String generates a textual representation of icon that can be used for
	// serialization such as when passing icon to a different process or saving
	// it to persistent storage.
	String() string
}

var _ Iconner = (*Icon)(nil)

func ifaceInitIconner(gifacePtr, data C.gpointer) {
	iface := (*C.GIconIface)(unsafe.Pointer(gifacePtr))
	iface.equal = (*[0]byte)(C._gotk4_gio2_IconIface_equal)
	iface.hash = (*[0]byte)(C._gotk4_gio2_IconIface_hash)
	iface.serialize = (*[0]byte)(C._gotk4_gio2_IconIface_serialize)
}

//export _gotk4_gio2_IconIface_equal
func _gotk4_gio2_IconIface_equal(arg0 *C.GIcon, arg1 *C.GIcon) (cret C.gboolean) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(IconOverrider)

	var _icon2 Iconner // out

	if arg1 != nil {
		{
			objptr := unsafe.Pointer(arg1)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Iconner)
				return ok
			})
			rv, ok := casted.(Iconner)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Iconner")
			}
			_icon2 = rv
		}
	}

	ok := iface.Equal(_icon2)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_IconIface_hash
func _gotk4_gio2_IconIface_hash(arg0 *C.GIcon) (cret C.guint) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(IconOverrider)

	guint := iface.Hash()

	cret = C.guint(guint)

	return cret
}

//export _gotk4_gio2_IconIface_serialize
func _gotk4_gio2_IconIface_serialize(arg0 *C.GIcon) (cret *C.GVariant) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(IconOverrider)

	variant := iface.Serialize()

	if variant != nil {
		cret = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(variant)))
	}

	return cret
}

func wrapIcon(obj *coreglib.Object) *Icon {
	return &Icon{
		Object: obj,
	}
}

func marshalIcon(p uintptr) (interface{}, error) {
	return wrapIcon(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Equal checks if two icons are equal.
//
// The function takes the following parameters:
//
//    - icon2 (optional): pointer to the second #GIcon.
//
// The function returns the following values:
//
//    - ok: TRUE if icon1 is equal to icon2. FALSE otherwise.
//
func (icon1 *Icon) Equal(icon2 Iconner) bool {
	var _arg0 *C.GIcon   // out
	var _arg1 *C.GIcon   // out
	var _cret C.gboolean // in

	if icon1 != nil {
		_arg0 = (*C.GIcon)(unsafe.Pointer(coreglib.InternObject(icon1).Native()))
	}
	if icon2 != nil {
		_arg1 = (*C.GIcon)(unsafe.Pointer(coreglib.InternObject(icon2).Native()))
	}

	_cret = C.g_icon_equal(_arg0, _arg1)
	runtime.KeepAlive(icon1)
	runtime.KeepAlive(icon2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Serialize serializes a #GIcon into a #GVariant. An equivalent #GIcon can be
// retrieved back by calling g_icon_deserialize() on the returned value. As
// serialization will avoid using raw icon data when possible, it only makes
// sense to transfer the #GVariant between processes on the same machine, (as
// opposed to over the network), and within the same file system namespace.
//
// The function returns the following values:
//
//    - variant (optional) or NULL when serialization fails. The #GVariant will
//      not be floating.
//
func (icon *Icon) Serialize() *glib.Variant {
	var _arg0 *C.GIcon    // out
	var _cret *C.GVariant // in

	_arg0 = (*C.GIcon)(unsafe.Pointer(coreglib.InternObject(icon).Native()))

	_cret = C.g_icon_serialize(_arg0)
	runtime.KeepAlive(icon)

	var _variant *glib.Variant // out

	if _cret != nil {
		_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_variant)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}

	return _variant
}

// String generates a textual representation of icon that can be used for
// serialization such as when passing icon to a different process or saving it
// to persistent storage. Use g_icon_new_for_string() to get icon back from the
// returned string.
//
// The encoding of the returned string is proprietary to #GIcon except in the
// following two cases
//
// - If icon is a Icon, the returned string is a native path (such as
// /path/to/my icon.png) without escaping if the #GFile for icon is a native
// file. If the file is not native, the returned string is the result of
// g_file_get_uri() (such as sftp://path/to/my20icon.png).
//
// - If icon is a Icon with exactly one name and no fallbacks, the encoding is
// simply the name (such as network-server).
//
// The function returns the following values:
//
//    - utf8 (optional): allocated NUL-terminated UTF8 string or NULL if icon
//      can't be serialized. Use g_free() to free.
//
func (icon *Icon) String() string {
	var _arg0 *C.GIcon // out
	var _cret *C.gchar // in

	_arg0 = (*C.GIcon)(unsafe.Pointer(coreglib.InternObject(icon).Native()))

	_cret = C.g_icon_to_string(_arg0)
	runtime.KeepAlive(icon)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// IconDeserialize deserializes a #GIcon previously serialized using
// g_icon_serialize().
//
// The function takes the following parameters:
//
//    - value created with g_icon_serialize().
//
// The function returns the following values:
//
//    - icon (optional) or NULL when deserialization fails.
//
func IconDeserialize(value *glib.Variant) *Icon {
	var _arg1 *C.GVariant // out
	var _cret *C.GIcon    // in

	_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))

	_cret = C.g_icon_deserialize(_arg1)
	runtime.KeepAlive(value)

	var _icon *Icon // out

	if _cret != nil {
		_icon = wrapIcon(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _icon
}

// IconHash gets a hash for an icon.
//
// The function takes the following parameters:
//
//    - icon to an icon object.
//
// The function returns the following values:
//
//    - guint containing a hash for the icon, suitable for use in a Table or
//      similar data structure.
//
func IconHash(icon unsafe.Pointer) uint {
	var _arg1 C.gconstpointer // out
	var _cret C.guint         // in

	_arg1 = (C.gconstpointer)(unsafe.Pointer(icon))

	_cret = C.g_icon_hash(_arg1)
	runtime.KeepAlive(icon)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// NewIconForString: generate a #GIcon instance from str. This function can fail
// if str is not valid - see g_icon_to_string() for discussion.
//
// If your application or library provides one or more #GIcon implementations
// you need to ensure that each #GType is registered with the type system prior
// to calling g_icon_new_for_string().
//
// The function takes the following parameters:
//
//    - str: string obtained via g_icon_to_string().
//
// The function returns the following values:
//
//    - icon: object implementing the #GIcon interface or NULL if error is set.
//
func NewIconForString(str string) (*Icon, error) {
	var _arg1 *C.gchar  // out
	var _cret *C.GIcon  // in
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_icon_new_for_string(_arg1, &_cerr)
	runtime.KeepAlive(str)

	var _icon *Icon  // out
	var _goerr error // out

	_icon = wrapIcon(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _icon, _goerr
}
