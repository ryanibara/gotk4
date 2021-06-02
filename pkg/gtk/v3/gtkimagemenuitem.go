// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

type ImageMenuItemPrivate struct {
	native C.GtkImageMenuItemPrivate
}

// WrapImageMenuItemPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapImageMenuItemPrivate(ptr unsafe.Pointer) *ImageMenuItemPrivate {
	if ptr == nil {
		return nil
	}

	return (*ImageMenuItemPrivate)(ptr)
}

func marshalImageMenuItemPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapImageMenuItemPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (i *ImageMenuItemPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}
