// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_emblemed_icon_get_type()), F: marshalEmblemedIconner},
	})
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

func wrapEmblemedIcon(obj *externglib.Object) *EmblemedIcon {
	return &EmblemedIcon{
		Object: obj,
		Icon: Icon{
			Object: obj,
		},
	}
}

func marshalEmblemedIconner(p uintptr) (interface{}, error) {
	return wrapEmblemedIcon(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewEmblemedIcon creates a new emblemed icon for icon with the emblem emblem.
//
// The function takes the following parameters:
//
//    - icon: #GIcon.
//    - emblem or NULL.
//
func NewEmblemedIcon(icon Iconner, emblem *Emblem) *EmblemedIcon {
	var _arg1 *C.GIcon   // out
	var _arg2 *C.GEmblem // out
	var _cret *C.GIcon   // in

	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))
	if emblem != nil {
		_arg2 = (*C.GEmblem)(unsafe.Pointer(emblem.Native()))
	}

	_cret = C.g_emblemed_icon_new(_arg1, _arg2)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(emblem)

	var _emblemedIcon *EmblemedIcon // out

	_emblemedIcon = wrapEmblemedIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _emblemedIcon
}

// AddEmblem adds emblem to the #GList of #GEmblems.
//
// The function takes the following parameters:
//
//    - emblem: #GEmblem.
//
func (emblemed *EmblemedIcon) AddEmblem(emblem *Emblem) {
	var _arg0 *C.GEmblemedIcon // out
	var _arg1 *C.GEmblem       // out

	_arg0 = (*C.GEmblemedIcon)(unsafe.Pointer(emblemed.Native()))
	_arg1 = (*C.GEmblem)(unsafe.Pointer(emblem.Native()))

	C.g_emblemed_icon_add_emblem(_arg0, _arg1)
	runtime.KeepAlive(emblemed)
	runtime.KeepAlive(emblem)
}

// ClearEmblems removes all the emblems from icon.
func (emblemed *EmblemedIcon) ClearEmblems() {
	var _arg0 *C.GEmblemedIcon // out

	_arg0 = (*C.GEmblemedIcon)(unsafe.Pointer(emblemed.Native()))

	C.g_emblemed_icon_clear_emblems(_arg0)
	runtime.KeepAlive(emblemed)
}

// Emblems gets the list of emblems for the icon.
func (emblemed *EmblemedIcon) Emblems() []Emblem {
	var _arg0 *C.GEmblemedIcon // out
	var _cret *C.GList         // in

	_arg0 = (*C.GEmblemedIcon)(unsafe.Pointer(emblemed.Native()))

	_cret = C.g_emblemed_icon_get_emblems(_arg0)
	runtime.KeepAlive(emblemed)

	var _list []Emblem // out

	_list = make([]Emblem, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.GEmblem)(v)
		var dst Emblem // out
		dst = *wrapEmblem(externglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// GetIcon gets the main icon for emblemed.
func (emblemed *EmblemedIcon) GetIcon() Iconner {
	var _arg0 *C.GEmblemedIcon // out
	var _cret *C.GIcon         // in

	_arg0 = (*C.GEmblemedIcon)(unsafe.Pointer(emblemed.Native()))

	_cret = C.g_emblemed_icon_get_icon(_arg0)
	runtime.KeepAlive(emblemed)

	var _icon Iconner // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Iconner is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Iconner)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.Iconner")
		}
		_icon = rv
	}

	return _icon
}
