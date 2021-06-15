// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_separator_tool_item_get_type()), F: marshalSeparatorToolItem},
	})
}

// SeparatorToolItem: a SeparatorToolItem is a ToolItem that separates groups of
// other ToolItems. Depending on the theme, a SeparatorToolItem will often look
// like a vertical line on horizontally docked toolbars.
//
// If the Toolbar child property “expand” is true and the property
// SeparatorToolItem:draw is false, a SeparatorToolItem will act as a “spring”
// that forces other items to the ends of the toolbar.
//
// Use gtk_separator_tool_item_new() to create a new SeparatorToolItem.
//
//
// CSS nodes
//
// GtkSeparatorToolItem has a single CSS node with name separator.
type SeparatorToolItem interface {
	ToolItem
	Activatable
	Buildable

	// Draw returns whether @item is drawn as a line, or just blank. See
	// gtk_separator_tool_item_set_draw().
	Draw() bool
	// SetDraw: whether @item is drawn as a vertical line, or just blank.
	// Setting this to false along with gtk_tool_item_set_expand() is useful to
	// create an item that forces following items to the end of the toolbar.
	SetDraw(draw bool)
}

// separatorToolItem implements the SeparatorToolItem class.
type separatorToolItem struct {
	ToolItem
	Activatable
	Buildable
}

var _ SeparatorToolItem = (*separatorToolItem)(nil)

// WrapSeparatorToolItem wraps a GObject to the right type. It is
// primarily used internally.
func WrapSeparatorToolItem(obj *externglib.Object) SeparatorToolItem {
	return separatorToolItem{
		ToolItem:    WrapToolItem(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalSeparatorToolItem(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSeparatorToolItem(obj), nil
}

// NewSeparatorToolItem constructs a class SeparatorToolItem.
func NewSeparatorToolItem() SeparatorToolItem {
	var _cret C.GtkSeparatorToolItem // in

	_cret = C.gtk_separator_tool_item_new()

	var _separatorToolItem SeparatorToolItem // out

	_separatorToolItem = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(SeparatorToolItem)

	return _separatorToolItem
}

// Draw returns whether @item is drawn as a line, or just blank. See
// gtk_separator_tool_item_set_draw().
func (i separatorToolItem) Draw() bool {
	var _arg0 *C.GtkSeparatorToolItem // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkSeparatorToolItem)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_separator_tool_item_get_draw(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDraw: whether @item is drawn as a vertical line, or just blank.
// Setting this to false along with gtk_tool_item_set_expand() is useful to
// create an item that forces following items to the end of the toolbar.
func (i separatorToolItem) SetDraw(draw bool) {
	var _arg0 *C.GtkSeparatorToolItem // out
	var _arg1 C.gboolean              // out

	_arg0 = (*C.GtkSeparatorToolItem)(unsafe.Pointer(i.Native()))
	if draw {
		_arg1 = C.TRUE
	}

	C.gtk_separator_tool_item_set_draw(_arg0, _arg1)
}
