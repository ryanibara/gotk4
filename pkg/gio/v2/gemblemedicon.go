// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.g_emblemed_icon_get_type()), F: marshalEmblemedIcon},
	})
}

// EmblemedIcon is an implementation of #GIcon that supports adding an emblem to
// an icon. Adding multiple emblems to an icon is ensured via
// g_emblemed_icon_add_emblem().
//
// Note that Icon allows no control over the position of the emblems. See also
// #GEmblem for more information.
type EmblemedIcon interface {
	gextras.Objector

	// AsIcon casts the class to the Icon interface.
	AsIcon() Icon

	// AddEmblemEmblemedIcon adds @emblem to the #GList of #GEmblems.
	AddEmblemEmblemedIcon(emblem Emblem)
	// ClearEmblemsEmblemedIcon removes all the emblems from @icon.
	ClearEmblemsEmblemedIcon()
	// Icon gets the main icon for @emblemed.
	Icon() Icon
}

// emblemedIcon implements the EmblemedIcon class.
type emblemedIcon struct {
	gextras.Objector
}

// WrapEmblemedIcon wraps a GObject to the right type. It is
// primarily used internally.
func WrapEmblemedIcon(obj *externglib.Object) EmblemedIcon {
	return emblemedIcon{
		Objector: obj,
	}
}

func marshalEmblemedIcon(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEmblemedIcon(obj), nil
}

// NewEmblemedIcon creates a new emblemed icon for @icon with the emblem
// @emblem.
func NewEmblemedIcon(icon Icon, emblem Emblem) EmblemedIcon {
	var _arg1 *C.GIcon   // out
	var _arg2 *C.GEmblem // out
	var _cret *C.GIcon   // in

	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))
	_arg2 = (*C.GEmblem)(unsafe.Pointer(emblem.Native()))

	_cret = C.g_emblemed_icon_new(_arg1, _arg2)

	var _emblemedIcon EmblemedIcon // out

	_emblemedIcon = WrapEmblemedIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _emblemedIcon
}

func (e emblemedIcon) AddEmblemEmblemedIcon(emblem Emblem) {
	var _arg0 *C.GEmblemedIcon // out
	var _arg1 *C.GEmblem       // out

	_arg0 = (*C.GEmblemedIcon)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GEmblem)(unsafe.Pointer(emblem.Native()))

	C.g_emblemed_icon_add_emblem(_arg0, _arg1)
}

func (e emblemedIcon) ClearEmblemsEmblemedIcon() {
	var _arg0 *C.GEmblemedIcon // out

	_arg0 = (*C.GEmblemedIcon)(unsafe.Pointer(e.Native()))

	C.g_emblemed_icon_clear_emblems(_arg0)
}

func (e emblemedIcon) Icon() Icon {
	var _arg0 *C.GEmblemedIcon // out
	var _cret *C.GIcon         // in

	_arg0 = (*C.GEmblemedIcon)(unsafe.Pointer(e.Native()))

	_cret = C.g_emblemed_icon_get_icon(_arg0)

	var _icon Icon // out

	_icon = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Icon)

	return _icon
}

func (e emblemedIcon) AsIcon() Icon {
	return WrapIcon(gextras.InternObject(e))
}
