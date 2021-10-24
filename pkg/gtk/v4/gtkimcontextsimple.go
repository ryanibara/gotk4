// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_im_context_simple_get_type()), F: marshalIMContextSimpler},
	})
}

const MAX_COMPOSE_LEN = C.MAX_COMPOSE_LEN

// IMContextSimple: GtkIMContextSimple is an input method supporting table-based
// input methods.
//
// GtkIMContextSimple has a built-in table of compose sequences that is derived
// from the X11 Compose files.
//
// GtkIMContextSimple reads additional compose sequences from the first of the
// following files that is found: ~/.config/gtk-4.0/Compose, ~/.XCompose,
// /usr/share/X11/locale/$locale/Compose (for locales that have a nontrivial
// Compose file). The syntax of these files is described in the Compose(5)
// manual page.
//
//
// Unicode characters
//
// GtkIMContextSimple also supports numeric entry of Unicode characters by
// typing <kbd>Ctrl</kbd>-<kbd>Shift</kbd>-<kbd>u</kbd>, followed by a
// hexadecimal Unicode codepoint.
//
// For example,
//
//    Ctrl-Shift-u 1 2 3 Enter
//
// yields U+0123 LATIN SMALL LETTER G WITH CEDILLA, i.e. ģ.
type IMContextSimple struct {
	IMContext
}

func wrapIMContextSimple(obj *externglib.Object) *IMContextSimple {
	return &IMContextSimple{
		IMContext: IMContext{
			Object: obj,
		},
	}
}

func marshalIMContextSimpler(p uintptr) (interface{}, error) {
	return wrapIMContextSimple(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewIMContextSimple creates a new IMContextSimple.
func NewIMContextSimple() *IMContextSimple {
	var _cret *C.GtkIMContext // in

	_cret = C.gtk_im_context_simple_new()

	var _imContextSimple *IMContextSimple // out

	_imContextSimple = wrapIMContextSimple(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _imContextSimple
}

// AddComposeFile adds an additional table from the X11 compose file.
//
// The function takes the following parameters:
//
//    - composeFile: path of compose file.
//
func (contextSimple *IMContextSimple) AddComposeFile(composeFile string) {
	var _arg0 *C.GtkIMContextSimple // out
	var _arg1 *C.char               // out

	_arg0 = (*C.GtkIMContextSimple)(unsafe.Pointer(contextSimple.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(composeFile)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_im_context_simple_add_compose_file(_arg0, _arg1)
	runtime.KeepAlive(contextSimple)
	runtime.KeepAlive(composeFile)
}
