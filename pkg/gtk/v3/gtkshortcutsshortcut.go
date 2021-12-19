// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcut_type_get_type()), F: marshalShortcutType},
		{T: externglib.Type(C.gtk_shortcuts_shortcut_get_type()), F: marshalShortcutsShortcutter},
	})
}

// ShortcutType specifies the kind of shortcut that is being described. More
// values may be added to this enumeration over time.
type ShortcutType C.gint

const (
	// ShortcutAccelerator: shortcut is a keyboard accelerator. The
	// ShortcutsShortcut:accelerator property will be used.
	ShortcutAccelerator ShortcutType = iota
	// ShortcutGesturePinch: shortcut is a pinch gesture. GTK+ provides an icon
	// and subtitle.
	ShortcutGesturePinch
	// ShortcutGestureStretch: shortcut is a stretch gesture. GTK+ provides an
	// icon and subtitle.
	ShortcutGestureStretch
	// ShortcutGestureRotateClockwise: shortcut is a clockwise rotation gesture.
	// GTK+ provides an icon and subtitle.
	ShortcutGestureRotateClockwise
	// ShortcutGestureRotateCounterclockwise: shortcut is a counterclockwise
	// rotation gesture. GTK+ provides an icon and subtitle.
	ShortcutGestureRotateCounterclockwise
	// ShortcutGestureTwoFingerSwipeLeft: shortcut is a two-finger swipe
	// gesture. GTK+ provides an icon and subtitle.
	ShortcutGestureTwoFingerSwipeLeft
	// ShortcutGestureTwoFingerSwipeRight: shortcut is a two-finger swipe
	// gesture. GTK+ provides an icon and subtitle.
	ShortcutGestureTwoFingerSwipeRight
	// ShortcutGesture: shortcut is a gesture. The ShortcutsShortcut:icon
	// property will be used.
	ShortcutGesture
)

func marshalShortcutType(p uintptr) (interface{}, error) {
	return ShortcutType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ShortcutType.
func (s ShortcutType) String() string {
	switch s {
	case ShortcutAccelerator:
		return "Accelerator"
	case ShortcutGesturePinch:
		return "GesturePinch"
	case ShortcutGestureStretch:
		return "GestureStretch"
	case ShortcutGestureRotateClockwise:
		return "GestureRotateClockwise"
	case ShortcutGestureRotateCounterclockwise:
		return "GestureRotateCounterclockwise"
	case ShortcutGestureTwoFingerSwipeLeft:
		return "GestureTwoFingerSwipeLeft"
	case ShortcutGestureTwoFingerSwipeRight:
		return "GestureTwoFingerSwipeRight"
	case ShortcutGesture:
		return "Gesture"
	default:
		return fmt.Sprintf("ShortcutType(%d)", s)
	}
}

// ShortcutsShortcutOverrider contains methods that are overridable.
type ShortcutsShortcutOverrider interface {
}

// ShortcutsShortcut represents a single keyboard shortcut or gesture with a
// short text. This widget is only meant to be used with ShortcutsWindow.
type ShortcutsShortcut struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer         = (*ShortcutsShortcut)(nil)
	_ externglib.Objector = (*ShortcutsShortcut)(nil)
)

func classInitShortcutsShortcutter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapShortcutsShortcut(obj *externglib.Object) *ShortcutsShortcut {
	return &ShortcutsShortcut{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalShortcutsShortcutter(p uintptr) (interface{}, error) {
	return wrapShortcutsShortcut(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
