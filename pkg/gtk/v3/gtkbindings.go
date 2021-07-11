// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// BindingsActivateEvent looks up key bindings for @object to find one matching
// @event, and if one was found, activate it.
func BindingsActivateEvent(object gextras.Objector, event *gdk.EventKey) bool {
	var _arg1 *C.GObject     // out
	var _arg2 *C.GdkEventKey // out
	var _cret C.gboolean     // in

	_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))
	_arg2 = (*C.GdkEventKey)(unsafe.Pointer(event))

	_cret = C.gtk_bindings_activate_event(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// BindingArg holds the data associated with an argument for a key binding
// signal emission as stored in BindingSignal.
type BindingArg struct {
	native C.GtkBindingArg
}

// Native returns the underlying C source pointer.
func (b *BindingArg) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// BindingEntry: each key binding element of a binding sets binding list is
// represented by a GtkBindingEntry.
type BindingEntry struct {
	native C.GtkBindingEntry
}

// Native returns the underlying C source pointer.
func (b *BindingEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// BindingSet: binding set maintains a list of activatable key bindings. A
// single binding set can match multiple types of widgets. Similar to style
// contexts, can be matched by any information contained in a widgets
// WidgetPath. When a binding within a set is matched upon activation, an action
// signal is emitted on the target widget to carry out the actual activation.
type BindingSet struct {
	native C.GtkBindingSet
}

// Native returns the underlying C source pointer.
func (b *BindingSet) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// BindingSignal stores the necessary information to activate a widget in
// response to a key press via a signal emission.
type BindingSignal struct {
	native C.GtkBindingSignal
}

// Native returns the underlying C source pointer.
func (b *BindingSignal) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}
