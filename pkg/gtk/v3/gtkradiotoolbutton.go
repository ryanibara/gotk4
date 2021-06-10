// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_radio_tool_button_get_type()), F: marshalRadioToolButton},
	})
}

// RadioToolButton: a RadioToolButton is a ToolItem that contains a radio
// button, that is, a button that is part of a group of toggle buttons where
// only one button can be active at a time.
//
// Use gtk_radio_tool_button_new() to create a new GtkRadioToolButton. Use
// gtk_radio_tool_button_new_from_widget() to create a new GtkRadioToolButton
// that is part of the same group as an existing GtkRadioToolButton.
//
//
// CSS nodes
//
// GtkRadioToolButton has a single CSS node with name toolbutton.
type RadioToolButton interface {
	ToggleToolButton
	Actionable
	Activatable
	Buildable

	// SetGroup adds @button to @group, removing it from the group it belonged
	// to before.
	SetGroup(group *glib.SList)
}

// radioToolButton implements the RadioToolButton interface.
type radioToolButton struct {
	ToggleToolButton
	Actionable
	Activatable
	Buildable
}

var _ RadioToolButton = (*radioToolButton)(nil)

// WrapRadioToolButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapRadioToolButton(obj *externglib.Object) RadioToolButton {
	return RadioToolButton{
		ToggleToolButton: WrapToggleToolButton(obj),
		Actionable:       WrapActionable(obj),
		Activatable:      WrapActivatable(obj),
		Buildable:        WrapBuildable(obj),
	}
}

func marshalRadioToolButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRadioToolButton(obj), nil
}

// SetGroup adds @button to @group, removing it from the group it belonged
// to before.
func (b radioToolButton) SetGroup(group *glib.SList) {
	var _arg0 *C.GtkRadioToolButton
	var _arg1 *C.GSList

	_arg0 = (*C.GtkRadioToolButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GSList)(unsafe.Pointer(group.Native()))

	C.gtk_radio_tool_button_set_group(_arg0, _arg1)
}
