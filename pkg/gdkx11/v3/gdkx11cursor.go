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

// glib.Type values for gdkx11cursor.go.
var GTypeX11Cursor = externglib.Type(C.gdk_x11_cursor_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeX11Cursor, F: marshalX11Cursor},
	})
}

// X11CursorOverrider contains methods that are overridable.
type X11CursorOverrider interface {
	externglib.Objector
}

// WrapX11CursorOverrider wraps the X11CursorOverrider
// interface implementation to access the instance methods.
func WrapX11CursorOverrider(obj X11CursorOverrider) *X11Cursor {
	return wrapX11Cursor(externglib.BaseObject(obj))
}

type X11Cursor struct {
	_ [0]func() // equal guard
	gdk.Cursor
}

var (
	_ gdk.Cursorrer = (*X11Cursor)(nil)
)

func classInitX11Cursorrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapX11Cursor(obj *externglib.Object) *X11Cursor {
	return &X11Cursor{
		Cursor: gdk.Cursor{
			Object: obj,
		},
	}
}

func marshalX11Cursor(p uintptr) (interface{}, error) {
	return wrapX11Cursor(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
