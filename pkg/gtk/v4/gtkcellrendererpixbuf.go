// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GTypeCellRendererPixbuf returns the GType for the type CellRendererPixbuf.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCellRendererPixbuf() coreglib.Type {
	gtype := coreglib.Type(C.gtk_cell_renderer_pixbuf_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalCellRendererPixbuf)
	return gtype
}

// CellRendererPixbuf renders a pixbuf in a cell
//
// A CellRendererPixbuf can be used to render an image in a cell. It allows to
// render either a given Pixbuf (set via the CellRendererPixbuf:pixbuf property)
// or a named icon (set via the CellRendererPixbuf:icon-name property).
//
// To support the tree view, CellRendererPixbuf also supports rendering two
// alternative pixbufs, when the CellRenderer:is-expander property is TRUE. If
// the CellRenderer:is-expanded property is TRUE and the
// CellRendererPixbuf:pixbuf-expander-open property is set to a pixbuf, it
// renders that pixbuf, if the CellRenderer:is-expanded property is FALSE and
// the CellRendererPixbuf:pixbuf-expander-closed property is set to a pixbuf, it
// renders that one.
type CellRendererPixbuf struct {
	_ [0]func() // equal guard
	CellRenderer
}

var (
	_ CellRendererer = (*CellRendererPixbuf)(nil)
)

func wrapCellRendererPixbuf(obj *coreglib.Object) *CellRendererPixbuf {
	return &CellRendererPixbuf{
		CellRenderer: CellRenderer{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
		},
	}
}

func marshalCellRendererPixbuf(p uintptr) (interface{}, error) {
	return wrapCellRendererPixbuf(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCellRendererPixbuf creates a new CellRendererPixbuf. Adjust rendering
// parameters using object properties. Object properties can be set globally
// (with g_object_set()). Also, with TreeViewColumn, you can bind a property to
// a value in a TreeModel. For example, you can bind the “pixbuf” property on
// the cell renderer to a pixbuf value in the model, thus rendering a different
// image in each row of the TreeView.
//
// The function returns the following values:
//
//    - cellRendererPixbuf: new cell renderer.
//
func NewCellRendererPixbuf() *CellRendererPixbuf {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_pixbuf_new()

	var _cellRendererPixbuf *CellRendererPixbuf // out

	_cellRendererPixbuf = wrapCellRendererPixbuf(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererPixbuf
}
