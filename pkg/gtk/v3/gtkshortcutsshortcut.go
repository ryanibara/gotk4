// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcut_type_get_type()), F: marshalShortcutType},
		{T: externglib.Type(C.gtk_shortcuts_shortcut_get_type()), F: marshalShortcutsShortcut},
	})
}

// ShortcutType: gtkShortcutType specifies the kind of shortcut that is being
// described. More values may be added to this enumeration over time.
type ShortcutType int

const (
	// accelerator: the shortcut is a keyboard accelerator. The
	// ShortcutsShortcut:accelerator property will be used.
	ShortcutTypeAccelerator ShortcutType = 0
	// GesturePinch: the shortcut is a pinch gesture. GTK+ provides an icon and
	// subtitle.
	ShortcutTypeGesturePinch ShortcutType = 1
	// GestureStretch: the shortcut is a stretch gesture. GTK+ provides an icon
	// and subtitle.
	ShortcutTypeGestureStretch ShortcutType = 2
	// GestureRotateClockwise: the shortcut is a clockwise rotation gesture.
	// GTK+ provides an icon and subtitle.
	ShortcutTypeGestureRotateClockwise ShortcutType = 3
	// GestureRotateCounterclockwise: the shortcut is a counterclockwise
	// rotation gesture. GTK+ provides an icon and subtitle.
	ShortcutTypeGestureRotateCounterclockwise ShortcutType = 4
	// GestureTwoFingerSwipeLeft: the shortcut is a two-finger swipe gesture.
	// GTK+ provides an icon and subtitle.
	ShortcutTypeGestureTwoFingerSwipeLeft ShortcutType = 5
	// GestureTwoFingerSwipeRight: the shortcut is a two-finger swipe gesture.
	// GTK+ provides an icon and subtitle.
	ShortcutTypeGestureTwoFingerSwipeRight ShortcutType = 6
	// gesture: the shortcut is a gesture. The ShortcutsShortcut:icon property
	// will be used.
	ShortcutTypeGesture ShortcutType = 7
)

func marshalShortcutType(p uintptr) (interface{}, error) {
	return ShortcutType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ShortcutsShortcut: a GtkShortcutsShortcut represents a single keyboard
// shortcut or gesture with a short text. This widget is only meant to be used
// with ShortcutsWindow.
type ShortcutsShortcut interface {
	Box
}

// shortcutsShortcut implements the ShortcutsShortcut class.
type shortcutsShortcut struct {
	Box
}

// WrapShortcutsShortcut wraps a GObject to the right type. It is
// primarily used internally.
func WrapShortcutsShortcut(obj *externglib.Object) ShortcutsShortcut {
	return shortcutsShortcut{
		Box: WrapBox(obj),
	}
}

func marshalShortcutsShortcut(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapShortcutsShortcut(obj), nil
}

func (b shortcutsShortcut) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b shortcutsShortcut) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b shortcutsShortcut) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b shortcutsShortcut) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b shortcutsShortcut) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b shortcutsShortcut) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b shortcutsShortcut) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b shortcutsShortcut) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b shortcutsShortcut) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b shortcutsShortcut) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (o shortcutsShortcut) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o shortcutsShortcut) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
