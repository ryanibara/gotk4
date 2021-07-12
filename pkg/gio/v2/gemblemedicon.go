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
		{T: externglib.Type(C.g_emblemed_icon_get_type()), F: marshalEmblemedIconer},
	})
}

// EmblemedIconer describes EmblemedIcon's methods.
type EmblemedIconer interface {
	// AddEmblem adds @emblem to the #GList of #GEmblems.
	AddEmblem(emblem Emblemer)
	// ClearEmblems removes all the emblems from @icon.
	ClearEmblems()
	// GetIcon gets the main icon for @emblemed.
	GetIcon() *Icon
}

// EmblemedIcon is an implementation of #GIcon that supports adding an emblem to
// an icon. Adding multiple emblems to an icon is ensured via
// g_emblemed_icon_add_emblem().
//
// Note that Icon allows no control over the position of the emblems. See also
// #GEmblem for more information.
type EmblemedIcon struct {
	*externglib.Object

	Icon
}

var (
	_ EmblemedIconer  = (*EmblemedIcon)(nil)
	_ gextras.Nativer = (*EmblemedIcon)(nil)
)

func wrapEmblemedIcon(obj *externglib.Object) *EmblemedIcon {
	return &EmblemedIcon{
		Object: obj,
		Icon: Icon{
			Object: obj,
		},
	}
}

func marshalEmblemedIconer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEmblemedIcon(obj), nil
}

// NewEmblemedIcon creates a new emblemed icon for @icon with the emblem
// @emblem.
func NewEmblemedIcon(icon Iconer, emblem Emblemer) *EmblemedIcon {
	var _arg1 *C.GIcon   // out
	var _arg2 *C.GEmblem // out
	var _cret *C.GIcon   // in

	_arg1 = (*C.GIcon)(unsafe.Pointer((icon).(gextras.Nativer).Native()))
	_arg2 = (*C.GEmblem)(unsafe.Pointer((emblem).(gextras.Nativer).Native()))

	_cret = C.g_emblemed_icon_new(_arg1, _arg2)

	var _emblemedIcon *EmblemedIcon // out

	_emblemedIcon = wrapEmblemedIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _emblemedIcon
}

// AddEmblem adds @emblem to the #GList of #GEmblems.
func (emblemed *EmblemedIcon) AddEmblem(emblem Emblemer) {
	var _arg0 *C.GEmblemedIcon // out
	var _arg1 *C.GEmblem       // out

	_arg0 = (*C.GEmblemedIcon)(unsafe.Pointer(emblemed.Native()))
	_arg1 = (*C.GEmblem)(unsafe.Pointer((emblem).(gextras.Nativer).Native()))

	C.g_emblemed_icon_add_emblem(_arg0, _arg1)
}

// ClearEmblems removes all the emblems from @icon.
func (emblemed *EmblemedIcon) ClearEmblems() {
	var _arg0 *C.GEmblemedIcon // out

	_arg0 = (*C.GEmblemedIcon)(unsafe.Pointer(emblemed.Native()))

	C.g_emblemed_icon_clear_emblems(_arg0)
}

// GetIcon gets the main icon for @emblemed.
func (emblemed *EmblemedIcon) GetIcon() *Icon {
	var _arg0 *C.GEmblemedIcon // out
	var _cret *C.GIcon         // in

	_arg0 = (*C.GEmblemedIcon)(unsafe.Pointer(emblemed.Native()))

	_cret = C.g_emblemed_icon_get_icon(_arg0)

	var _icon *Icon // out

	_icon = wrapIcon(externglib.Take(unsafe.Pointer(_cret)))

	return _icon
}
