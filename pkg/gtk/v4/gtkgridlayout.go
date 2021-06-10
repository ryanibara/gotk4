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
		{T: externglib.Type(C.gtk_grid_layout_get_type()), F: marshalGridLayout},
		{T: externglib.Type(C.gtk_grid_layout_child_get_type()), F: marshalGridLayoutChild},
	})
}

// GridLayout: `GtkGridLayout` is a layout manager which arranges child widgets
// in rows and columns.
//
// Children have an "attach point" defined by the horizontal and vertical index
// of the cell they occupy; children can span multiple rows or columns. The
// layout properties for setting the attach points and spans are set using the
// [class@Gtk.GridLayoutChild] associated to each child widget.
//
// The behaviour of `GtkGridLayout` when several children occupy the same grid
// cell is undefined.
//
// `GtkGridLayout` can be used like a `GtkBoxLayout` if all children are
// attached to the same row or column; however, if you only ever need a single
// row or column, you should consider using `GtkBoxLayout`.
type GridLayout interface {
	LayoutManager

	// BaselineRow retrieves the row set with
	// gtk_grid_layout_set_baseline_row().
	BaselineRow() int
	// ColumnHomogeneous checks whether all columns of @grid should have the
	// same width.
	ColumnHomogeneous() bool
	// ColumnSpacing retrieves the spacing set with
	// gtk_grid_layout_set_column_spacing().
	ColumnSpacing() uint
	// RowHomogeneous checks whether all rows of @grid should have the same
	// height.
	RowHomogeneous() bool
	// RowSpacing retrieves the spacing set with
	// gtk_grid_layout_set_row_spacing().
	RowSpacing() uint
	// SetBaselineRow sets which row defines the global baseline for the entire
	// grid.
	//
	// Each row in the grid can have its own local baseline, but only one of
	// those is global, meaning it will be the baseline in the parent of the
	// @grid.
	SetBaselineRow(row int)
	// SetColumnHomogeneous sets whether all columns of @grid should have the
	// same width.
	SetColumnHomogeneous(homogeneous bool)
	// SetColumnSpacing sets the amount of space to insert between consecutive
	// columns.
	SetColumnSpacing(spacing uint)
	// SetRowBaselinePosition sets how the baseline should be positioned on @row
	// of the grid, in case that row is assigned more space than is requested.
	SetRowBaselinePosition(row int, pos BaselinePosition)
	// SetRowHomogeneous sets whether all rows of @grid should have the same
	// height.
	SetRowHomogeneous(homogeneous bool)
	// SetRowSpacing sets the amount of space to insert between consecutive
	// rows.
	SetRowSpacing(spacing uint)
}

// gridLayout implements the GridLayout interface.
type gridLayout struct {
	LayoutManager
}

var _ GridLayout = (*gridLayout)(nil)

// WrapGridLayout wraps a GObject to the right type. It is
// primarily used internally.
func WrapGridLayout(obj *externglib.Object) GridLayout {
	return GridLayout{
		LayoutManager: WrapLayoutManager(obj),
	}
}

func marshalGridLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGridLayout(obj), nil
}

// BaselineRow retrieves the row set with
// gtk_grid_layout_set_baseline_row().
func (g gridLayout) BaselineRow() int {
	var _arg0 *C.GtkGridLayout

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(g.Native()))

	var _cret C.int

	_cret = C.gtk_grid_layout_get_baseline_row(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// ColumnHomogeneous checks whether all columns of @grid should have the
// same width.
func (g gridLayout) ColumnHomogeneous() bool {
	var _arg0 *C.GtkGridLayout

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(g.Native()))

	var _cret C.gboolean

	_cret = C.gtk_grid_layout_get_column_homogeneous(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ColumnSpacing retrieves the spacing set with
// gtk_grid_layout_set_column_spacing().
func (g gridLayout) ColumnSpacing() uint {
	var _arg0 *C.GtkGridLayout

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(g.Native()))

	var _cret C.guint

	_cret = C.gtk_grid_layout_get_column_spacing(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// RowHomogeneous checks whether all rows of @grid should have the same
// height.
func (g gridLayout) RowHomogeneous() bool {
	var _arg0 *C.GtkGridLayout

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(g.Native()))

	var _cret C.gboolean

	_cret = C.gtk_grid_layout_get_row_homogeneous(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// RowSpacing retrieves the spacing set with
// gtk_grid_layout_set_row_spacing().
func (g gridLayout) RowSpacing() uint {
	var _arg0 *C.GtkGridLayout

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(g.Native()))

	var _cret C.guint

	_cret = C.gtk_grid_layout_get_row_spacing(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// SetBaselineRow sets which row defines the global baseline for the entire
// grid.
//
// Each row in the grid can have its own local baseline, but only one of
// those is global, meaning it will be the baseline in the parent of the
// @grid.
func (g gridLayout) SetBaselineRow(row int) {
	var _arg0 *C.GtkGridLayout
	var _arg1 C.int

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(g.Native()))
	_arg1 = C.int(row)

	C.gtk_grid_layout_set_baseline_row(_arg0, _arg1)
}

// SetColumnHomogeneous sets whether all columns of @grid should have the
// same width.
func (g gridLayout) SetColumnHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkGridLayout
	var _arg1 C.gboolean

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(g.Native()))
	if homogeneous {
		_arg1 = C.gboolean(1)
	}

	C.gtk_grid_layout_set_column_homogeneous(_arg0, _arg1)
}

// SetColumnSpacing sets the amount of space to insert between consecutive
// columns.
func (g gridLayout) SetColumnSpacing(spacing uint) {
	var _arg0 *C.GtkGridLayout
	var _arg1 C.guint

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(g.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_grid_layout_set_column_spacing(_arg0, _arg1)
}

// SetRowBaselinePosition sets how the baseline should be positioned on @row
// of the grid, in case that row is assigned more space than is requested.
func (g gridLayout) SetRowBaselinePosition(row int, pos BaselinePosition) {
	var _arg0 *C.GtkGridLayout
	var _arg1 C.int
	var _arg2 C.GtkBaselinePosition

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(g.Native()))
	_arg1 = C.int(row)
	_arg2 = (C.GtkBaselinePosition)(pos)

	C.gtk_grid_layout_set_row_baseline_position(_arg0, _arg1, _arg2)
}

// SetRowHomogeneous sets whether all rows of @grid should have the same
// height.
func (g gridLayout) SetRowHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkGridLayout
	var _arg1 C.gboolean

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(g.Native()))
	if homogeneous {
		_arg1 = C.gboolean(1)
	}

	C.gtk_grid_layout_set_row_homogeneous(_arg0, _arg1)
}

// SetRowSpacing sets the amount of space to insert between consecutive
// rows.
func (g gridLayout) SetRowSpacing(spacing uint) {
	var _arg0 *C.GtkGridLayout
	var _arg1 C.guint

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(g.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_grid_layout_set_row_spacing(_arg0, _arg1)
}

// GridLayoutChild: `GtkLayoutChild` subclass for children in a `GtkGridLayout`.
type GridLayoutChild interface {
	LayoutChild

	// Column retrieves the column number to which @child attaches its left
	// side.
	Column() int
	// ColumnSpan retrieves the number of columns that @child spans to.
	ColumnSpan() int
	// Row retrieves the row number to which @child attaches its top side.
	Row() int
	// RowSpan retrieves the number of rows that @child spans to.
	RowSpan() int
	// SetColumn sets the column number to attach the left side of @child.
	SetColumn(column int)
	// SetColumnSpan sets the number of columns @child spans to.
	SetColumnSpan(span int)
	// SetRow sets the row to place @child in.
	SetRow(row int)
	// SetRowSpan sets the number of rows @child spans to.
	SetRowSpan(span int)
}

// gridLayoutChild implements the GridLayoutChild interface.
type gridLayoutChild struct {
	LayoutChild
}

var _ GridLayoutChild = (*gridLayoutChild)(nil)

// WrapGridLayoutChild wraps a GObject to the right type. It is
// primarily used internally.
func WrapGridLayoutChild(obj *externglib.Object) GridLayoutChild {
	return GridLayoutChild{
		LayoutChild: WrapLayoutChild(obj),
	}
}

func marshalGridLayoutChild(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGridLayoutChild(obj), nil
}

// Column retrieves the column number to which @child attaches its left
// side.
func (c gridLayoutChild) Column() int {
	var _arg0 *C.GtkGridLayoutChild

	_arg0 = (*C.GtkGridLayoutChild)(unsafe.Pointer(c.Native()))

	var _cret C.int

	_cret = C.gtk_grid_layout_child_get_column(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// ColumnSpan retrieves the number of columns that @child spans to.
func (c gridLayoutChild) ColumnSpan() int {
	var _arg0 *C.GtkGridLayoutChild

	_arg0 = (*C.GtkGridLayoutChild)(unsafe.Pointer(c.Native()))

	var _cret C.int

	_cret = C.gtk_grid_layout_child_get_column_span(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Row retrieves the row number to which @child attaches its top side.
func (c gridLayoutChild) Row() int {
	var _arg0 *C.GtkGridLayoutChild

	_arg0 = (*C.GtkGridLayoutChild)(unsafe.Pointer(c.Native()))

	var _cret C.int

	_cret = C.gtk_grid_layout_child_get_row(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// RowSpan retrieves the number of rows that @child spans to.
func (c gridLayoutChild) RowSpan() int {
	var _arg0 *C.GtkGridLayoutChild

	_arg0 = (*C.GtkGridLayoutChild)(unsafe.Pointer(c.Native()))

	var _cret C.int

	_cret = C.gtk_grid_layout_child_get_row_span(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// SetColumn sets the column number to attach the left side of @child.
func (c gridLayoutChild) SetColumn(column int) {
	var _arg0 *C.GtkGridLayoutChild
	var _arg1 C.int

	_arg0 = (*C.GtkGridLayoutChild)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(column)

	C.gtk_grid_layout_child_set_column(_arg0, _arg1)
}

// SetColumnSpan sets the number of columns @child spans to.
func (c gridLayoutChild) SetColumnSpan(span int) {
	var _arg0 *C.GtkGridLayoutChild
	var _arg1 C.int

	_arg0 = (*C.GtkGridLayoutChild)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(span)

	C.gtk_grid_layout_child_set_column_span(_arg0, _arg1)
}

// SetRow sets the row to place @child in.
func (c gridLayoutChild) SetRow(row int) {
	var _arg0 *C.GtkGridLayoutChild
	var _arg1 C.int

	_arg0 = (*C.GtkGridLayoutChild)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(row)

	C.gtk_grid_layout_child_set_row(_arg0, _arg1)
}

// SetRowSpan sets the number of rows @child spans to.
func (c gridLayoutChild) SetRowSpan(span int) {
	var _arg0 *C.GtkGridLayoutChild
	var _arg1 C.int

	_arg0 = (*C.GtkGridLayoutChild)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(span)

	C.gtk_grid_layout_child_set_row_span(_arg0, _arg1)
}
