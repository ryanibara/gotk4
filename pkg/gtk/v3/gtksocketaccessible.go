// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_socket_accessible_get_type()), F: marshalSocketAccessible},
	})
}

type SocketAccessible interface {
	ContainerAccessible

	Embed(path string)
}

// socketAccessible implements the SocketAccessible interface.
type socketAccessible struct {
	ContainerAccessible
}

var _ SocketAccessible = (*socketAccessible)(nil)

// WrapSocketAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapSocketAccessible(obj *externglib.Object) SocketAccessible {
	return SocketAccessible{
		ContainerAccessible: WrapContainerAccessible(obj),
	}
}

func marshalSocketAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSocketAccessible(obj), nil
}

func (s socketAccessible) Embed(path string) {
	var arg0 *C.GtkSocketAccessible
	var arg1 *C.gchar

	arg0 = (*C.GtkSocketAccessible)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_socket_accessible_embed(arg0, path)
}
