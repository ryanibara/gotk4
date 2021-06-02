// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <gdk/gdk.h>
import "C"

// DragActionIsUnique checks if @action represents a single action or includes
// multiple actions.
//
// When @action is 0 - ie no action was given, true is returned.
func DragActionIsUnique(action DragAction) bool {
	var arg1 C.GdkDragAction

	arg1 = (C.GdkDragAction)(action)

	ret := C.gdk_drag_action_is_unique(arg1)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}
