// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gdkx11keys.go.
var GTypeX11Keymap = coreglib.Type(C.gdk_x11_keymap_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeX11Keymap, F: marshalX11Keymap},
	})
}

// X11KeymapOverrider contains methods that are overridable.
type X11KeymapOverrider interface {
}

type X11Keymap struct {
	_ [0]func() // equal guard
	gdk.Keymap
}

var (
	_ coreglib.Objector = (*X11Keymap)(nil)
)

func classInitX11Keymapper(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapX11Keymap(obj *coreglib.Object) *X11Keymap {
	return &X11Keymap{
		Keymap: gdk.Keymap{
			Object: obj,
		},
	}
}

func marshalX11Keymap(p uintptr) (interface{}, error) {
	return wrapX11Keymap(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
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
func (keymap *X11Keymap) GroupForState(state uint32) int32 {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(keymap).Native()))
	_arg1 = C.guint(state)
	*(**X11Keymap)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("GdkX11", "X11Keymap").InvokeMethod("get_group_for_state", args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(keymap)
	runtime.KeepAlive(state)

	var _gint int32 // out

	_gint = int32(_cret)

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
func (keymap *X11Keymap) KeyIsModifier(keycode uint32) bool {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.guint    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(keymap).Native()))
	_arg1 = C.guint(keycode)
	*(**X11Keymap)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("GdkX11", "X11Keymap").InvokeMethod("key_is_modifier", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(keymap)
	runtime.KeepAlive(keycode)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
