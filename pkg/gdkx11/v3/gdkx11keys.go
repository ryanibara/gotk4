// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gdk/gdkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_keymap_get_type()), F: marshalX11Keymap},
	})
}

type X11Keymap interface {
	gdk.Keymap

	// GroupForState extracts the group from the state field sent in an X Key
	// event. This is only needed for code processing raw X events, since
	// EventKey directly includes an is_modifier field.
	GroupForState(state uint) int
	// KeyIsModifier determines whether a particular key code represents a key
	// that is a modifier. That is, it’s a key that normally just affects the
	// keyboard state and the behavior of other keys rather than producing a
	// direct effect itself. This is only needed for code processing raw X
	// events, since EventKey directly includes an is_modifier field.
	KeyIsModifier(keycode uint) bool
}

// x11Keymap implements the X11Keymap interface.
type x11Keymap struct {
	gdk.Keymap
}

var _ X11Keymap = (*x11Keymap)(nil)

// WrapX11Keymap wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Keymap(obj *externglib.Object) X11Keymap {
	return X11Keymap{
		gdk.Keymap: gdk.WrapKeymap(obj),
	}
}

func marshalX11Keymap(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Keymap(obj), nil
}

// GroupForState extracts the group from the state field sent in an X Key
// event. This is only needed for code processing raw X events, since
// EventKey directly includes an is_modifier field.
func (k x11Keymap) GroupForState(state uint) int {
	var arg0 *C.GdkKeymap
	var arg1 C.guint

	arg0 = (*C.GdkKeymap)(unsafe.Pointer(k.Native()))
	arg1 = C.guint(state)

	var cret C.gint
	var ret1 int

	cret = C.gdk_x11_keymap_get_group_for_state(arg0, state)

	ret1 = C.gint(cret)

	return ret1
}

// KeyIsModifier determines whether a particular key code represents a key
// that is a modifier. That is, it’s a key that normally just affects the
// keyboard state and the behavior of other keys rather than producing a
// direct effect itself. This is only needed for code processing raw X
// events, since EventKey directly includes an is_modifier field.
func (k x11Keymap) KeyIsModifier(keycode uint) bool {
	var arg0 *C.GdkKeymap
	var arg1 C.guint

	arg0 = (*C.GdkKeymap)(unsafe.Pointer(k.Native()))
	arg1 = C.guint(keycode)

	var cret C.gboolean
	var ret1 bool

	cret = C.gdk_x11_keymap_key_is_modifier(arg0, keycode)

	ret1 = C.bool(cret) != C.false

	return ret1
}
