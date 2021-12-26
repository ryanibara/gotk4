// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_themed_icon_get_type()), F: marshalThemedIconner},
	})
}

// ThemedIcon is an implementation of #GIcon that supports icon themes. Icon
// contains a list of all of the icons present in an icon theme, so that icons
// can be looked up quickly. Icon does not provide actual pixmaps for icons,
// just the icon names. Ideally something like gtk_icon_theme_choose_icon()
// should be used to resolve the list of names so that fallback icons work
// nicely with themes that inherit other themes.
type ThemedIcon struct {
	_ [0]func() // equal guard
	*externglib.Object

	Icon
}

var (
	_ externglib.Objector = (*ThemedIcon)(nil)
)

func wrapThemedIcon(obj *externglib.Object) *ThemedIcon {
	return &ThemedIcon{
		Object: obj,
		Icon: Icon{
			Object: obj,
		},
	}
}

func marshalThemedIconner(p uintptr) (interface{}, error) {
	return wrapThemedIcon(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewThemedIcon creates a new themed icon for iconname.
//
// The function takes the following parameters:
//
//    - iconname: string containing an icon name.
//
// The function returns the following values:
//
//    - themedIcon: new Icon.
//
func NewThemedIcon(iconname string) *ThemedIcon {
	var _arg1 *C.char  // out
	var _cret *C.GIcon // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconname)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_themed_icon_new(_arg1)
	runtime.KeepAlive(iconname)

	var _themedIcon *ThemedIcon // out

	_themedIcon = wrapThemedIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _themedIcon
}

// NewThemedIconFromNames creates a new themed icon for iconnames.
//
// The function takes the following parameters:
//
//    - iconnames: array of strings containing icon names.
//
// The function returns the following values:
//
//    - themedIcon: new Icon.
//
func NewThemedIconFromNames(iconnames []string) *ThemedIcon {
	var _arg1 **C.char // out
	var _arg2 C.int
	var _cret *C.GIcon // in

	_arg2 = (C.int)(len(iconnames))
	_arg1 = (**C.char)(C.calloc(C.size_t(len(iconnames)), C.size_t(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((**C.char)(_arg1), len(iconnames))
		for i := range iconnames {
			out[i] = (*C.char)(unsafe.Pointer(C.CString(iconnames[i])))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	_cret = C.g_themed_icon_new_from_names(_arg1, _arg2)
	runtime.KeepAlive(iconnames)

	var _themedIcon *ThemedIcon // out

	_themedIcon = wrapThemedIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _themedIcon
}

// NewThemedIconWithDefaultFallbacks creates a new themed icon for iconname, and
// all the names that can be created by shortening iconname at '-' characters.
//
// In the following example, icon1 and icon2 are equivalent:
//
//    const char *names[] = {
//      "gnome-dev-cdrom-audio",
//      "gnome-dev-cdrom",
//      "gnome-dev",
//      "gnome"
//    };
//
//    icon1 = g_themed_icon_new_from_names (names, 4);
//    icon2 = g_themed_icon_new_with_default_fallbacks ("gnome-dev-cdrom-audio");.
//
// The function takes the following parameters:
//
//    - iconname: string containing an icon name.
//
// The function returns the following values:
//
//    - themedIcon: new Icon.
//
func NewThemedIconWithDefaultFallbacks(iconname string) *ThemedIcon {
	var _arg1 *C.char  // out
	var _cret *C.GIcon // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconname)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_themed_icon_new_with_default_fallbacks(_arg1)
	runtime.KeepAlive(iconname)

	var _themedIcon *ThemedIcon // out

	_themedIcon = wrapThemedIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _themedIcon
}

// AppendName: append a name to the list of icons from within icon.
//
// Note that doing so invalidates the hash computed by prior calls to
// g_icon_hash().
//
// The function takes the following parameters:
//
//    - iconname: name of icon to append to list of icons from within icon.
//
func (icon *ThemedIcon) AppendName(iconname string) {
	var _arg0 *C.GThemedIcon // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GThemedIcon)(unsafe.Pointer(icon.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconname)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_themed_icon_append_name(_arg0, _arg1)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(iconname)
}

// Names gets the names of icons from within icon.
//
// The function returns the following values:
//
//    - utf8s: list of icon names.
//
func (icon *ThemedIcon) Names() []string {
	var _arg0 *C.GThemedIcon // out
	var _cret **C.gchar      // in

	_arg0 = (*C.GThemedIcon)(unsafe.Pointer(icon.Native()))

	_cret = C.g_themed_icon_get_names(_arg0)
	runtime.KeepAlive(icon)

	var _utf8s []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// PrependName: prepend a name to the list of icons from within icon.
//
// Note that doing so invalidates the hash computed by prior calls to
// g_icon_hash().
//
// The function takes the following parameters:
//
//    - iconname: name of icon to prepend to list of icons from within icon.
//
func (icon *ThemedIcon) PrependName(iconname string) {
	var _arg0 *C.GThemedIcon // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GThemedIcon)(unsafe.Pointer(icon.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconname)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_themed_icon_prepend_name(_arg0, _arg1)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(iconname)
}
