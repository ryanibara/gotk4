// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_text_get_type()), F: marshalCellRendererTexter},
	})
}

// CellRendererTextOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CellRendererTextOverrider interface {
	// The function takes the following parameters:
	//
	//    - path
	//    - newText
	//
	Edited(path, newText string)
}

// CellRendererText renders a given text in its cell, using the font, color and
// style information provided by its properties. The text will be ellipsized if
// it is too long and the CellRendererText:ellipsize property allows it.
//
// If the CellRenderer:mode is GTK_CELL_RENDERER_MODE_EDITABLE, the
// CellRendererText allows to edit its text using an entry.
type CellRendererText struct {
	CellRenderer

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ CellRendererer = (*CellRendererText)(nil)
)

func wrapCellRendererText(obj *externglib.Object) *CellRendererText {
	return &CellRendererText{
		CellRenderer: CellRenderer{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
		},
	}
}

func marshalCellRendererTexter(p uintptr) (interface{}, error) {
	return wrapCellRendererText(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectEdited: this signal is emitted after renderer has been edited.
//
// It is the responsibility of the application to update the model and store
// new_text at the position indicated by path.
func (renderer *CellRendererText) ConnectEdited(f func(path, newText string)) externglib.SignalHandle {
	return renderer.Connect("edited", f)
}

// NewCellRendererText creates a new CellRendererText. Adjust how text is drawn
// using object properties. Object properties can be set globally (with
// g_object_set()). Also, with TreeViewColumn, you can bind a property to a
// value in a TreeModel. For example, you can bind the “text” property on the
// cell renderer to a string value in the model, thus rendering a different
// string in each row of the TreeView.
//
// The function returns the following values:
//
//    - cellRendererText: new cell renderer.
//
func NewCellRendererText() *CellRendererText {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_text_new()

	var _cellRendererText *CellRendererText // out

	_cellRendererText = wrapCellRendererText(externglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererText
}

// SetFixedHeightFromFont sets the height of a renderer to explicitly be
// determined by the “font” and “y_pad” property set on it. Further changes in
// these properties do not affect the height, so they must be accompanied by a
// subsequent call to this function. Using this function is unflexible, and
// should really only be used if calculating the size of a cell is too slow (ie,
// a massive number of cells displayed). If number_of_rows is -1, then the fixed
// height is unset, and the height is determined by the properties again.
//
// The function takes the following parameters:
//
//    - numberOfRows: number of rows of text each cell renderer is allocated, or
//      -1.
//
func (renderer *CellRendererText) SetFixedHeightFromFont(numberOfRows int) {
	var _arg0 *C.GtkCellRendererText // out
	var _arg1 C.gint                 // out

	_arg0 = (*C.GtkCellRendererText)(unsafe.Pointer(renderer.Native()))
	_arg1 = C.gint(numberOfRows)

	C.gtk_cell_renderer_text_set_fixed_height_from_font(_arg0, _arg1)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(numberOfRows)
}
