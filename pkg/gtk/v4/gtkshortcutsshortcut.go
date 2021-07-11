// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcut_type_get_type()), F: marshalShortcutType},
		{T: externglib.Type(C.gtk_shortcuts_shortcut_get_type()), F: marshalShortcutsShortcutter},
	})
}

// ShortcutType specifies the kind of shortcut that is being described.
//
// More values may be added to this enumeration over time.
type ShortcutType int

const (
	// Accelerator: shortcut is a keyboard accelerator. The
	// GtkShortcutsShortcut:accelerator property will be used.
	ShortcutTypeAccelerator ShortcutType = iota
	// GesturePinch: shortcut is a pinch gesture. GTK provides an icon and
	// subtitle.
	ShortcutTypeGesturePinch
	// GestureStretch: shortcut is a stretch gesture. GTK provides an icon and
	// subtitle.
	ShortcutTypeGestureStretch
	// GestureRotateClockwise: shortcut is a clockwise rotation gesture. GTK
	// provides an icon and subtitle.
	ShortcutTypeGestureRotateClockwise
	// GestureRotateCounterclockwise: shortcut is a counterclockwise rotation
	// gesture. GTK provides an icon and subtitle.
	ShortcutTypeGestureRotateCounterclockwise
	// GestureTwoFingerSwipeLeft: shortcut is a two-finger swipe gesture. GTK
	// provides an icon and subtitle.
	ShortcutTypeGestureTwoFingerSwipeLeft
	// GestureTwoFingerSwipeRight: shortcut is a two-finger swipe gesture. GTK
	// provides an icon and subtitle.
	ShortcutTypeGestureTwoFingerSwipeRight
	// Gesture: shortcut is a gesture. The GtkShortcutsShortcut:icon property
	// will be used.
	ShortcutTypeGesture
	// GestureSwipeLeft: shortcut is a swipe gesture. GTK provides an icon and
	// subtitle.
	ShortcutTypeGestureSwipeLeft
	// GestureSwipeRight: shortcut is a swipe gesture. GTK provides an icon and
	// subtitle.
	ShortcutTypeGestureSwipeRight
)

func marshalShortcutType(p uintptr) (interface{}, error) {
	return ShortcutType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ShortcutsShortcutter describes ShortcutsShortcut's methods.
type ShortcutsShortcutter interface {
	privateShortcutsShortcut()
}

// ShortcutsShortcut: `GtkShortcutsShortcut` represents a single keyboard
// shortcut or gesture with a short text.
//
// This widget is only meant to be used with `GtkShortcutsWindow`.
type ShortcutsShortcut struct {
	Widget
}

var (
	_ ShortcutsShortcutter = (*ShortcutsShortcut)(nil)
	_ gextras.Nativer      = (*ShortcutsShortcut)(nil)
)

func wrapShortcutsShortcut(obj *externglib.Object) ShortcutsShortcutter {
	return &ShortcutsShortcut{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalShortcutsShortcutter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapShortcutsShortcut(obj), nil
}

func (*ShortcutsShortcut) privateShortcutsShortcut() {}
