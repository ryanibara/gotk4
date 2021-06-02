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

type ImageAccessiblePrivate struct {
	native C.GtkImageAccessiblePrivate
}

// WrapImageAccessiblePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapImageAccessiblePrivate(ptr unsafe.Pointer) *ImageAccessiblePrivate {
	if ptr == nil {
		return nil
	}

	return (*ImageAccessiblePrivate)(ptr)
}

func marshalImageAccessiblePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapImageAccessiblePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (i *ImageAccessiblePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}
