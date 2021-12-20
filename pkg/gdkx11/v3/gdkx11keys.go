// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_keymap_get_type()), F: marshalX11Keymapper},
	})
}

type X11Keymap struct {
	gdk.Keymap
}

var (
	_ externglib.Objector = (*X11Keymap)(nil)
)

func wrapX11Keymap(obj *externglib.Object) *X11Keymap {
	return &X11Keymap{
		Keymap: gdk.Keymap{
			Object: obj,
		},
	}
}

func marshalX11Keymapper(p uintptr) (interface{}, error) {
	return wrapX11Keymap(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// GroupForState extracts the group from the state field sent in an X Key event.
// This is only needed for code processing raw X events, since EventKey directly
// includes an is_modifier field.
//
// The function takes the following parameters:
//
//    - state: raw state returned from X.
//
// The function returns the following values:
//
//    - gint: index of the active keyboard group for the event.
//
func (keymap *X11Keymap) GroupForState(state uint) int {
	var _arg0 *C.GdkKeymap // out
	var _arg1 C.guint      // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(keymap.Native()))
	_arg1 = C.guint(state)

	_cret = C.gdk_x11_keymap_get_group_for_state(_arg0, _arg1)
	runtime.KeepAlive(keymap)
	runtime.KeepAlive(state)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// KeyIsModifier determines whether a particular key code represents a key that
// is a modifier. That is, it’s a key that normally just affects the keyboard
// state and the behavior of other keys rather than producing a direct effect
// itself. This is only needed for code processing raw X events, since EventKey
// directly includes an is_modifier field.
//
// The function takes the following parameters:
//
//    - keycode: hardware keycode from a key event.
//
// The function returns the following values:
//
//    - ok: TRUE if the hardware keycode is a modifier key.
//
func (keymap *X11Keymap) KeyIsModifier(keycode uint) bool {
	var _arg0 *C.GdkKeymap // out
	var _arg1 C.guint      // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(keymap.Native()))
	_arg1 = C.guint(keycode)

	_cret = C.gdk_x11_keymap_key_is_modifier(_arg0, _arg1)
	runtime.KeepAlive(keymap)
	runtime.KeepAlive(keycode)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
