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
		{T: externglib.Type(C.gtk_shortcuts_group_get_type()), F: marshalShortcutsGroup},
	})
}

// ShortcutsGroup: a GtkShortcutsGroup represents a group of related keyboard
// shortcuts or gestures. The group has a title. It may optionally be associated
// with a view of the application, which can be used to show only relevant
// shortcuts depending on the application context.
//
// This widget is only meant to be used with ShortcutsWindow.
type ShortcutsGroup interface {
	Box
}

// shortcutsGroup implements the ShortcutsGroup class.
type shortcutsGroup struct {
	Box
}

// WrapShortcutsGroup wraps a GObject to the right type. It is
// primarily used internally.
func WrapShortcutsGroup(obj *externglib.Object) ShortcutsGroup {
	return shortcutsGroup{
		Box: WrapBox(obj),
	}
}

func marshalShortcutsGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapShortcutsGroup(obj), nil
}

func (b shortcutsGroup) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b shortcutsGroup) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b shortcutsGroup) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b shortcutsGroup) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b shortcutsGroup) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b shortcutsGroup) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b shortcutsGroup) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b shortcutsGroup) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b shortcutsGroup) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b shortcutsGroup) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (o shortcutsGroup) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o shortcutsGroup) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
