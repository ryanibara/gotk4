// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcut_type_get_type()), F: marshalShortcutType},
		{T: externglib.Type(C.gtk_shortcuts_shortcut_get_type()), F: marshalShortcutsShortcuter},
	})
}

// ShortcutType specifies the kind of shortcut that is being described. More
// values may be added to this enumeration over time.
type ShortcutType int

const (
	// Accelerator: shortcut is a keyboard accelerator. The
	// ShortcutsShortcut:accelerator property will be used.
	ShortcutTypeAccelerator ShortcutType = iota
	// GesturePinch: shortcut is a pinch gesture. GTK+ provides an icon and
	// subtitle.
	ShortcutTypeGesturePinch
	// GestureStretch: shortcut is a stretch gesture. GTK+ provides an icon and
	// subtitle.
	ShortcutTypeGestureStretch
	// GestureRotateClockwise: shortcut is a clockwise rotation gesture. GTK+
	// provides an icon and subtitle.
	ShortcutTypeGestureRotateClockwise
	// GestureRotateCounterclockwise: shortcut is a counterclockwise rotation
	// gesture. GTK+ provides an icon and subtitle.
	ShortcutTypeGestureRotateCounterclockwise
	// GestureTwoFingerSwipeLeft: shortcut is a two-finger swipe gesture. GTK+
	// provides an icon and subtitle.
	ShortcutTypeGestureTwoFingerSwipeLeft
	// GestureTwoFingerSwipeRight: shortcut is a two-finger swipe gesture. GTK+
	// provides an icon and subtitle.
	ShortcutTypeGestureTwoFingerSwipeRight
	// Gesture: shortcut is a gesture. The ShortcutsShortcut:icon property will
	// be used.
	ShortcutTypeGesture
)

func marshalShortcutType(p uintptr) (interface{}, error) {
	return ShortcutType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ShortcutsShortcuter describes ShortcutsShortcut's methods.
type ShortcutsShortcuter interface {
	privateShortcutsShortcut()
}

// ShortcutsShortcut represents a single keyboard shortcut or gesture with a
// short text. This widget is only meant to be used with ShortcutsWindow.
type ShortcutsShortcut struct {
	Box
}

var (
	_ ShortcutsShortcuter = (*ShortcutsShortcut)(nil)
	_ gextras.Nativer     = (*ShortcutsShortcut)(nil)
)

func wrapShortcutsShortcut(obj *externglib.Object) *ShortcutsShortcut {
	return &ShortcutsShortcut{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalShortcutsShortcuter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapShortcutsShortcut(obj), nil
}

func (*ShortcutsShortcut) privateShortcutsShortcut() {}
