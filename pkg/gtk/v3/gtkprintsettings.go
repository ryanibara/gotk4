// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

type PrintSettingsFunc func(key string, value string)

//export gotk4_PrintSettingsFunc
func gotk4_PrintSettingsFunc(arg0 *C.gchar, arg1 *C.gchar, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(PrintSettingsFunc)
	fn(key, value, userData)
}

// PageRange: see also gtk_print_settings_set_page_ranges().
type PageRange struct {
	native C.GtkPageRange
}

// WrapPageRange wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPageRange(ptr unsafe.Pointer) *PageRange {
	if ptr == nil {
		return nil
	}

	return (*PageRange)(ptr)
}

func marshalPageRange(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPageRange(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *PageRange) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// Start gets the field inside the struct.
func (p *PageRange) Start() int {
	v = C.gint(p.native.start)
}

// End gets the field inside the struct.
func (p *PageRange) End() int {
	v = C.gint(p.native.end)
}
