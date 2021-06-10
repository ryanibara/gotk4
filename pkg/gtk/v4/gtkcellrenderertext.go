// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_text_get_type()), F: marshalCellRendererText},
	})
}

// CellRendererText renders text in a cell
//
// A CellRendererText renders a given text in its cell, using the font, color
// and style information provided by its properties. The text will be ellipsized
// if it is too long and the CellRendererText:ellipsize property allows it.
//
// If the CellRenderer:mode is GTK_CELL_RENDERER_MODE_EDITABLE, the
// CellRendererText allows to edit its text using an entry.
type CellRendererText interface {
	CellRenderer

	// SetFixedHeightFromFont sets the height of a renderer to explicitly be
	// determined by the “font” and “y_pad” property set on it. Further changes
	// in these properties do not affect the height, so they must be accompanied
	// by a subsequent call to this function. Using this function is inflexible,
	// and should really only be used if calculating the size of a cell is too
	// slow (ie, a massive number of cells displayed). If @number_of_rows is -1,
	// then the fixed height is unset, and the height is determined by the
	// properties again.
	SetFixedHeightFromFont(numberOfRows int)
}

// cellRendererText implements the CellRendererText interface.
type cellRendererText struct {
	CellRenderer
}

var _ CellRendererText = (*cellRendererText)(nil)

// WrapCellRendererText wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellRendererText(obj *externglib.Object) CellRendererText {
	return CellRendererText{
		CellRenderer: WrapCellRenderer(obj),
	}
}

func marshalCellRendererText(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellRendererText(obj), nil
}

// SetFixedHeightFromFont sets the height of a renderer to explicitly be
// determined by the “font” and “y_pad” property set on it. Further changes
// in these properties do not affect the height, so they must be accompanied
// by a subsequent call to this function. Using this function is inflexible,
// and should really only be used if calculating the size of a cell is too
// slow (ie, a massive number of cells displayed). If @number_of_rows is -1,
// then the fixed height is unset, and the height is determined by the
// properties again.
func (r cellRendererText) SetFixedHeightFromFont(numberOfRows int) {
	var _arg0 *C.GtkCellRendererText
	var _arg1 C.int

	_arg0 = (*C.GtkCellRendererText)(unsafe.Pointer(r.Native()))
	_arg1 = C.int(numberOfRows)

	C.gtk_cell_renderer_text_set_fixed_height_from_font(_arg0, _arg1)
}
