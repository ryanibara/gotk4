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

type ArrowPrivate struct {
	native C.GtkArrowPrivate
}

// WrapArrowPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapArrowPrivate(ptr unsafe.Pointer) *ArrowPrivate {
	if ptr == nil {
		return nil
	}

	return (*ArrowPrivate)(ptr)
}

func marshalArrowPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapArrowPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *ArrowPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}
