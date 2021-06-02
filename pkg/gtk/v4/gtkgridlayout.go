// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
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
	// RowBaselinePosition returns the baseline position of @row.
	//
	// If no value has been set with
	// [method@Gtk.GridLayout.set_row_baseline_position], the default value of
	// GTK_BASELINE_POSITION_CENTER is returned.
	RowBaselinePosition(row int) BaselinePosition
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

// NewGridLayout constructs a class GridLayout.
func NewGridLayout() GridLayout {
	ret := C.gtk_grid_layout_new()

	var ret0 GridLayout

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(GridLayout)

	return ret0
}

// BaselineRow retrieves the row set with
// gtk_grid_layout_set_baseline_row().
func (g gridLayout) BaselineRow() int {
	var arg0 *C.GtkGridLayout

	arg0 = (*C.GtkGridLayout)(g.Native())

	ret := C.gtk_grid_layout_get_baseline_row(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// ColumnHomogeneous checks whether all columns of @grid should have the
// same width.
func (g gridLayout) ColumnHomogeneous() bool {
	var arg0 *C.GtkGridLayout

	arg0 = (*C.GtkGridLayout)(g.Native())

	ret := C.gtk_grid_layout_get_column_homogeneous(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// ColumnSpacing retrieves the spacing set with
// gtk_grid_layout_set_column_spacing().
func (g gridLayout) ColumnSpacing() uint {
	var arg0 *C.GtkGridLayout

	arg0 = (*C.GtkGridLayout)(g.Native())

	ret := C.gtk_grid_layout_get_column_spacing(arg0)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// RowBaselinePosition returns the baseline position of @row.
//
// If no value has been set with
// [method@Gtk.GridLayout.set_row_baseline_position], the default value of
// GTK_BASELINE_POSITION_CENTER is returned.
func (g gridLayout) RowBaselinePosition(row int) BaselinePosition {
	var arg0 *C.GtkGridLayout
	var arg1 C.int

	arg0 = (*C.GtkGridLayout)(g.Native())
	arg1 = C.int(row)

	ret := C.gtk_grid_layout_get_row_baseline_position(arg0, arg1)

	var ret0 BaselinePosition

	ret0 = BaselinePosition(ret)

	return ret0
}

// RowHomogeneous checks whether all rows of @grid should have the same
// height.
func (g gridLayout) RowHomogeneous() bool {
	var arg0 *C.GtkGridLayout

	arg0 = (*C.GtkGridLayout)(g.Native())

	ret := C.gtk_grid_layout_get_row_homogeneous(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// RowSpacing retrieves the spacing set with
// gtk_grid_layout_set_row_spacing().
func (g gridLayout) RowSpacing() uint {
	var arg0 *C.GtkGridLayout

	arg0 = (*C.GtkGridLayout)(g.Native())

	ret := C.gtk_grid_layout_get_row_spacing(arg0)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// SetBaselineRow sets which row defines the global baseline for the entire
// grid.
//
// Each row in the grid can have its own local baseline, but only one of
// those is global, meaning it will be the baseline in the parent of the
// @grid.
func (g gridLayout) SetBaselineRow(row int) {
	var arg0 *C.GtkGridLayout
	var arg1 C.int

	arg0 = (*C.GtkGridLayout)(g.Native())
	arg1 = C.int(row)

	C.gtk_grid_layout_set_baseline_row(arg0, arg1)
}

// SetColumnHomogeneous sets whether all columns of @grid should have the
// same width.
func (g gridLayout) SetColumnHomogeneous(homogeneous bool) {
	var arg0 *C.GtkGridLayout
	var arg1 C.gboolean

	arg0 = (*C.GtkGridLayout)(g.Native())
	if homogeneous {
		arg1 = C.TRUE
	}

	C.gtk_grid_layout_set_column_homogeneous(arg0, arg1)
}

// SetColumnSpacing sets the amount of space to insert between consecutive
// columns.
func (g gridLayout) SetColumnSpacing(spacing uint) {
	var arg0 *C.GtkGridLayout
	var arg1 C.guint

	arg0 = (*C.GtkGridLayout)(g.Native())
	arg1 = C.guint(spacing)

	C.gtk_grid_layout_set_column_spacing(arg0, arg1)
}

// SetRowBaselinePosition sets how the baseline should be positioned on @row
// of the grid, in case that row is assigned more space than is requested.
func (g gridLayout) SetRowBaselinePosition(row int, pos BaselinePosition) {
	var arg0 *C.GtkGridLayout
	var arg1 C.int
	var arg2 C.GtkBaselinePosition

	arg0 = (*C.GtkGridLayout)(g.Native())
	arg1 = C.int(row)
	arg2 = (C.GtkBaselinePosition)(pos)

	C.gtk_grid_layout_set_row_baseline_position(arg0, arg1, arg2)
}

// SetRowHomogeneous sets whether all rows of @grid should have the same
// height.
func (g gridLayout) SetRowHomogeneous(homogeneous bool) {
	var arg0 *C.GtkGridLayout
	var arg1 C.gboolean

	arg0 = (*C.GtkGridLayout)(g.Native())
	if homogeneous {
		arg1 = C.TRUE
	}

	C.gtk_grid_layout_set_row_homogeneous(arg0, arg1)
}

// SetRowSpacing sets the amount of space to insert between consecutive
// rows.
func (g gridLayout) SetRowSpacing(spacing uint) {
	var arg0 *C.GtkGridLayout
	var arg1 C.guint

	arg0 = (*C.GtkGridLayout)(g.Native())
	arg1 = C.guint(spacing)

	C.gtk_grid_layout_set_row_spacing(arg0, arg1)
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
	var arg0 *C.GtkGridLayoutChild

	arg0 = (*C.GtkGridLayoutChild)(c.Native())

	ret := C.gtk_grid_layout_child_get_column(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// ColumnSpan retrieves the number of columns that @child spans to.
func (c gridLayoutChild) ColumnSpan() int {
	var arg0 *C.GtkGridLayoutChild

	arg0 = (*C.GtkGridLayoutChild)(c.Native())

	ret := C.gtk_grid_layout_child_get_column_span(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// Row retrieves the row number to which @child attaches its top side.
func (c gridLayoutChild) Row() int {
	var arg0 *C.GtkGridLayoutChild

	arg0 = (*C.GtkGridLayoutChild)(c.Native())

	ret := C.gtk_grid_layout_child_get_row(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// RowSpan retrieves the number of rows that @child spans to.
func (c gridLayoutChild) RowSpan() int {
	var arg0 *C.GtkGridLayoutChild

	arg0 = (*C.GtkGridLayoutChild)(c.Native())

	ret := C.gtk_grid_layout_child_get_row_span(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// SetColumn sets the column number to attach the left side of @child.
func (c gridLayoutChild) SetColumn(column int) {
	var arg0 *C.GtkGridLayoutChild
	var arg1 C.int

	arg0 = (*C.GtkGridLayoutChild)(c.Native())
	arg1 = C.int(column)

	C.gtk_grid_layout_child_set_column(arg0, arg1)
}

// SetColumnSpan sets the number of columns @child spans to.
func (c gridLayoutChild) SetColumnSpan(span int) {
	var arg0 *C.GtkGridLayoutChild
	var arg1 C.int

	arg0 = (*C.GtkGridLayoutChild)(c.Native())
	arg1 = C.int(span)

	C.gtk_grid_layout_child_set_column_span(arg0, arg1)
}

// SetRow sets the row to place @child in.
func (c gridLayoutChild) SetRow(row int) {
	var arg0 *C.GtkGridLayoutChild
	var arg1 C.int

	arg0 = (*C.GtkGridLayoutChild)(c.Native())
	arg1 = C.int(row)

	C.gtk_grid_layout_child_set_row(arg0, arg1)
}

// SetRowSpan sets the number of rows @child spans to.
func (c gridLayoutChild) SetRowSpan(span int) {
	var arg0 *C.GtkGridLayoutChild
	var arg1 C.int

	arg0 = (*C.GtkGridLayoutChild)(c.Native())
	arg1 = C.int(span)

	C.gtk_grid_layout_child_set_row_span(arg0, arg1)
}
