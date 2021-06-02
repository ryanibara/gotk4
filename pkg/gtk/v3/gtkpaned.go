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

type PanedPrivate struct {
	native C.GtkPanedPrivate
}

// WrapPanedPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPanedPrivate(ptr unsafe.Pointer) *PanedPrivate {
	if ptr == nil {
		return nil
	}

	return (*PanedPrivate)(ptr)
}

func marshalPanedPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPanedPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *PanedPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}
