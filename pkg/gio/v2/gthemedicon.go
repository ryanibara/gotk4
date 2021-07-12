// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_themed_icon_get_type()), F: marshalThemedIconer},
	})
}

// ThemedIconer describes ThemedIcon's methods.
type ThemedIconer interface {
	// AppendName: append a name to the list of icons from within @icon.
	AppendName(iconname string)
	// Names gets the names of icons from within @icon.
	Names() []string
	// PrependName: prepend a name to the list of icons from within @icon.
	PrependName(iconname string)
}

// ThemedIcon is an implementation of #GIcon that supports icon themes. Icon
// contains a list of all of the icons present in an icon theme, so that icons
// can be looked up quickly. Icon does not provide actual pixmaps for icons,
// just the icon names. Ideally something like gtk_icon_theme_choose_icon()
// should be used to resolve the list of names so that fallback icons work
// nicely with themes that inherit other themes.
type ThemedIcon struct {
	*externglib.Object

	Icon
}

var (
	_ ThemedIconer    = (*ThemedIcon)(nil)
	_ gextras.Nativer = (*ThemedIcon)(nil)
)

func wrapThemedIcon(obj *externglib.Object) *ThemedIcon {
	return &ThemedIcon{
		Object: obj,
		Icon: Icon{
			Object: obj,
		},
	}
}

func marshalThemedIconer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapThemedIcon(obj), nil
}

// NewThemedIcon creates a new themed icon for @iconname.
func NewThemedIcon(iconname string) *ThemedIcon {
	var _arg1 *C.char  // out
	var _cret *C.GIcon // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconname)))

	_cret = C.g_themed_icon_new(_arg1)

	var _themedIcon *ThemedIcon // out

	_themedIcon = wrapThemedIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _themedIcon
}

// NewThemedIconFromNames creates a new themed icon for @iconnames.
func NewThemedIconFromNames(iconnames []string) *ThemedIcon {
	var _arg1 **C.char
	var _arg2 C.int
	var _cret *C.GIcon // in

	_arg2 = C.int(len(iconnames))
	_arg1 = (**C.char)(C.malloc(C.ulong(len(iconnames)) * C.ulong(unsafe.Sizeof(uint(0)))))
	{
		out := unsafe.Slice((**C.char)(_arg1), len(iconnames))
		for i := range iconnames {
			out[i] = (*C.char)(unsafe.Pointer(C.CString(iconnames[i])))
		}
	}

	_cret = C.g_themed_icon_new_from_names(_arg1, _arg2)

	var _themedIcon *ThemedIcon // out

	_themedIcon = wrapThemedIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _themedIcon
}

// NewThemedIconWithDefaultFallbacks creates a new themed icon for @iconname,
// and all the names that can be created by shortening @iconname at '-'
// characters.
//
// In the following example, @icon1 and @icon2 are equivalent:
//
//    const char *names[] = {
//      "gnome-dev-cdrom-audio",
//      "gnome-dev-cdrom",
//      "gnome-dev",
//      "gnome"
//    };
//
//    icon1 = g_themed_icon_new_from_names (names, 4);
//    icon2 = g_themed_icon_new_with_default_fallbacks ("gnome-dev-cdrom-audio");
func NewThemedIconWithDefaultFallbacks(iconname string) *ThemedIcon {
	var _arg1 *C.char  // out
	var _cret *C.GIcon // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconname)))

	_cret = C.g_themed_icon_new_with_default_fallbacks(_arg1)

	var _themedIcon *ThemedIcon // out

	_themedIcon = wrapThemedIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _themedIcon
}

// AppendName: append a name to the list of icons from within @icon.
//
// Note that doing so invalidates the hash computed by prior calls to
// g_icon_hash().
func (icon *ThemedIcon) AppendName(iconname string) {
	var _arg0 *C.GThemedIcon // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GThemedIcon)(unsafe.Pointer(icon.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconname)))

	C.g_themed_icon_append_name(_arg0, _arg1)
}

// Names gets the names of icons from within @icon.
func (icon *ThemedIcon) Names() []string {
	var _arg0 *C.GThemedIcon // out
	var _cret **C.gchar

	_arg0 = (*C.GThemedIcon)(unsafe.Pointer(icon.Native()))

	_cret = C.g_themed_icon_get_names(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// PrependName: prepend a name to the list of icons from within @icon.
//
// Note that doing so invalidates the hash computed by prior calls to
// g_icon_hash().
func (icon *ThemedIcon) PrependName(iconname string) {
	var _arg0 *C.GThemedIcon // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GThemedIcon)(unsafe.Pointer(icon.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconname)))

	C.g_themed_icon_prepend_name(_arg0, _arg1)
}
