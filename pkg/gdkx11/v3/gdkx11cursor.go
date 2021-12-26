// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_cursor_get_type()), F: marshalX11Cursorrer},
	})
}

type X11Cursor struct {
	_ [0]func() // equal guard
	gdk.Cursor
}

var (
	_ gdk.Cursorrer = (*X11Cursor)(nil)
)

func wrapX11Cursor(obj *externglib.Object) *X11Cursor {
	return &X11Cursor{
		Cursor: gdk.Cursor{
			Object: obj,
		},
	}
}

func marshalX11Cursorrer(p uintptr) (interface{}, error) {
	return wrapX11Cursor(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
