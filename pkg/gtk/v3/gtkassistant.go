// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
//
// gint gotk4_AssistantPageFunc(gint, gpointer);
import "C"

// AssistantPageFunc: a function used by gtk_assistant_set_forward_page_func()
// to know which is the next page given a current one. It’s called both for
// computing the next page when the user presses the “forward” button and for
// handling the behavior of the “last” button.
type AssistantPageFunc func(currentPage int) int

//export gotk4_AssistantPageFunc
func gotk4_AssistantPageFunc(arg0 C.gint, arg1 C.gpointer) C.gint {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var currentPage int

	currentPage = int(arg0)

	gint := v.(AssistantPageFunc)(currentPage)
}

type AssistantPrivate struct {
	native C.GtkAssistantPrivate
}

// WrapAssistantPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAssistantPrivate(ptr unsafe.Pointer) *AssistantPrivate {
	if ptr == nil {
		return nil
	}

	return (*AssistantPrivate)(ptr)
}

func marshalAssistantPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAssistantPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AssistantPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}
