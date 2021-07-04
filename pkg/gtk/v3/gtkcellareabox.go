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
		{T: externglib.Type(C.gtk_cell_area_box_get_type()), F: marshalCellAreaBox},
	})
}

// CellAreaBox: the CellAreaBox renders cell renderers into a row or a column
// depending on its Orientation.
//
// GtkCellAreaBox uses a notion of packing. Packing refers to adding cell
// renderers with reference to a particular position in a CellAreaBox. There are
// two reference positions: the start and the end of the box. When the
// CellAreaBox is oriented in the GTK_ORIENTATION_VERTICAL orientation, the
// start is defined as the top of the box and the end is defined as the bottom.
// In the GTK_ORIENTATION_HORIZONTAL orientation start is defined as the left
// side and the end is defined as the right side.
//
// Alignments of CellRenderers rendered in adjacent rows can be configured by
// configuring the CellAreaBox align child cell property with
// gtk_cell_area_cell_set_property() or by specifying the "align" argument to
// gtk_cell_area_box_pack_start() and gtk_cell_area_box_pack_end().
type CellAreaBox interface {
	CellArea
	Orientable

	Spacing() int

	PackEndCellAreaBox(renderer CellRenderer, expand bool, align bool, fixed bool)

	PackStartCellAreaBox(renderer CellRenderer, expand bool, align bool, fixed bool)

	SetSpacingCellAreaBox(spacing int)
}

// cellAreaBox implements the CellAreaBox class.
type cellAreaBox struct {
	CellArea
}

// WrapCellAreaBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellAreaBox(obj *externglib.Object) CellAreaBox {
	return cellAreaBox{
		CellArea: WrapCellArea(obj),
	}
}

func marshalCellAreaBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellAreaBox(obj), nil
}

func NewCellAreaBox() CellAreaBox {
	var _cret *C.GtkCellArea // in

	_cret = C.gtk_cell_area_box_new()

	var _cellAreaBox CellAreaBox // out

	_cellAreaBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellAreaBox)

	return _cellAreaBox
}

func (b cellAreaBox) Spacing() int {
	var _arg0 *C.GtkCellAreaBox // out
	var _cret C.gint            // in

	_arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_cell_area_box_get_spacing(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (b cellAreaBox) PackEndCellAreaBox(renderer CellRenderer, expand bool, align bool, fixed bool) {
	var _arg0 *C.GtkCellAreaBox  // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.gboolean         // out
	var _arg3 C.gboolean         // out
	var _arg4 C.gboolean         // out

	_arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	if expand {
		_arg2 = C.TRUE
	}
	if align {
		_arg3 = C.TRUE
	}
	if fixed {
		_arg4 = C.TRUE
	}

	C.gtk_cell_area_box_pack_end(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (b cellAreaBox) PackStartCellAreaBox(renderer CellRenderer, expand bool, align bool, fixed bool) {
	var _arg0 *C.GtkCellAreaBox  // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.gboolean         // out
	var _arg3 C.gboolean         // out
	var _arg4 C.gboolean         // out

	_arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	if expand {
		_arg2 = C.TRUE
	}
	if align {
		_arg3 = C.TRUE
	}
	if fixed {
		_arg4 = C.TRUE
	}

	C.gtk_cell_area_box_pack_start(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (b cellAreaBox) SetSpacingCellAreaBox(spacing int) {
	var _arg0 *C.GtkCellAreaBox // out
	var _arg1 C.gint            // out

	_arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.gint(spacing)

	C.gtk_cell_area_box_set_spacing(_arg0, _arg1)
}

func (b cellAreaBox) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b cellAreaBox) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b cellAreaBox) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b cellAreaBox) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b cellAreaBox) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b cellAreaBox) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b cellAreaBox) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b cellAreaBox) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b cellAreaBox) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b cellAreaBox) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (c cellAreaBox) AddAttribute(cell CellRenderer, attribute string, column int) {
	WrapCellLayout(gextras.InternObject(c)).AddAttribute(cell, attribute, column)
}

func (c cellAreaBox) Clear() {
	WrapCellLayout(gextras.InternObject(c)).Clear()
}

func (c cellAreaBox) ClearAttributes(cell CellRenderer) {
	WrapCellLayout(gextras.InternObject(c)).ClearAttributes(cell)
}

func (c cellAreaBox) Area() CellArea {
	return WrapCellLayout(gextras.InternObject(c)).Area()
}

func (c cellAreaBox) PackEnd(cell CellRenderer, expand bool) {
	WrapCellLayout(gextras.InternObject(c)).PackEnd(cell, expand)
}

func (c cellAreaBox) PackStart(cell CellRenderer, expand bool) {
	WrapCellLayout(gextras.InternObject(c)).PackStart(cell, expand)
}

func (c cellAreaBox) Reorder(cell CellRenderer, position int) {
	WrapCellLayout(gextras.InternObject(c)).Reorder(cell, position)
}

func (o cellAreaBox) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o cellAreaBox) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
