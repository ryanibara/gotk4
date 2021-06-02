// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0
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
		{T: externglib.Type(C.g_emblemed_icon_get_type()), F: marshalEmblemedIcon},
	})
}

type EmblemedIconPrivate struct {
	native C.GEmblemedIconPrivate
}

// WrapEmblemedIconPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEmblemedIconPrivate(ptr unsafe.Pointer) *EmblemedIconPrivate {
	if ptr == nil {
		return nil
	}

	return (*EmblemedIconPrivate)(ptr)
}

func marshalEmblemedIconPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapEmblemedIconPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (e *EmblemedIconPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EmblemedIcon is an implementation of #GIcon that supports adding an emblem to
// an icon. Adding multiple emblems to an icon is ensured via
// g_emblemed_icon_add_emblem().
//
// Note that Icon allows no control over the position of the emblems. See also
// #GEmblem for more information.
type EmblemedIcon interface {
	gextras.Objector
	Icon

	// AddEmblem adds @emblem to the #GList of #GEmblems.
	AddEmblem(emblem Emblem)
	// ClearEmblems removes all the emblems from @icon.
	ClearEmblems()
	// Emblems gets the list of emblems for the @icon.
	Emblems() *glib.List
	// Icon gets the main icon for @emblemed.
	Icon() Icon
}

// emblemedIcon implements the EmblemedIcon interface.
type emblemedIcon struct {
	gextras.Objector
	Icon
}

var _ EmblemedIcon = (*emblemedIcon)(nil)

// WrapEmblemedIcon wraps a GObject to the right type. It is
// primarily used internally.
func WrapEmblemedIcon(obj *externglib.Object) EmblemedIcon {
	return EmblemedIcon{
		Objector: obj,
		Icon:     WrapIcon(obj),
	}
}

func marshalEmblemedIcon(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEmblemedIcon(obj), nil
}

// NewEmblemedIcon constructs a class EmblemedIcon.
func NewEmblemedIcon(icon Icon, emblem Emblem) EmblemedIcon {
	var arg1 *C.GIcon
	var arg2 *C.GEmblem

	arg1 = (*C.GIcon)(icon.Native())
	arg2 = (*C.GEmblem)(emblem.Native())

	ret := C.g_emblemed_icon_new(arg1, arg2)

	var ret0 EmblemedIcon

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(EmblemedIcon)

	return ret0
}

// AddEmblem adds @emblem to the #GList of #GEmblems.
func (e emblemedIcon) AddEmblem(emblem Emblem) {
	var arg0 *C.GEmblemedIcon
	var arg1 *C.GEmblem

	arg0 = (*C.GEmblemedIcon)(e.Native())
	arg1 = (*C.GEmblem)(emblem.Native())

	C.g_emblemed_icon_add_emblem(arg0, arg1)
}

// ClearEmblems removes all the emblems from @icon.
func (e emblemedIcon) ClearEmblems() {
	var arg0 *C.GEmblemedIcon

	arg0 = (*C.GEmblemedIcon)(e.Native())

	C.g_emblemed_icon_clear_emblems(arg0)
}

// Emblems gets the list of emblems for the @icon.
func (e emblemedIcon) Emblems() *glib.List {
	var arg0 *C.GEmblemedIcon

	arg0 = (*C.GEmblemedIcon)(e.Native())

	ret := C.g_emblemed_icon_get_emblems(arg0)

	var ret0 *glib.List

	{
		ret0 = glib.WrapList(unsafe.Pointer(ret))
	}

	return ret0
}

// Icon gets the main icon for @emblemed.
func (e emblemedIcon) Icon() Icon {
	var arg0 *C.GEmblemedIcon

	arg0 = (*C.GEmblemedIcon)(e.Native())

	ret := C.g_emblemed_icon_get_icon(arg0)

	var ret0 Icon

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Icon)

	return ret0
}
