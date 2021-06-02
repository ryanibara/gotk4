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

type RangeAccessiblePrivate struct {
	native C.GtkRangeAccessiblePrivate
}

// WrapRangeAccessiblePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRangeAccessiblePrivate(ptr unsafe.Pointer) *RangeAccessiblePrivate {
	if ptr == nil {
		return nil
	}

	return (*RangeAccessiblePrivate)(ptr)
}

func marshalRangeAccessiblePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRangeAccessiblePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RangeAccessiblePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}
