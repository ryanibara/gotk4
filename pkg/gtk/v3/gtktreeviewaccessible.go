// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_tree_view_accessible_get_type()), F: marshalTreeViewAccessible},
	})
}

type TreeViewAccessible interface {
	ContainerAccessible
	atk.Table
	CellAccessibleParent
}

// treeViewAccessible implements the TreeViewAccessible class.
type treeViewAccessible struct {
	ContainerAccessible
}

// WrapTreeViewAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapTreeViewAccessible(obj *externglib.Object) TreeViewAccessible {
	return treeViewAccessible{
		ContainerAccessible: WrapContainerAccessible(obj),
	}
}

func marshalTreeViewAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeViewAccessible(obj), nil
}

func (t treeViewAccessible) AddColumnSelection(column int) bool {
	return atk.WrapTable(gextras.InternObject(t)).AddColumnSelection(column)
}

func (t treeViewAccessible) AddRowSelection(row int) bool {
	return atk.WrapTable(gextras.InternObject(t)).AddRowSelection(row)
}

func (t treeViewAccessible) Caption() atk.Object {
	return atk.WrapTable(gextras.InternObject(t)).Caption()
}

func (t treeViewAccessible) ColumnAtIndex(index_ int) int {
	return atk.WrapTable(gextras.InternObject(t)).ColumnAtIndex(index_)
}

func (t treeViewAccessible) ColumnDescription(column int) string {
	return atk.WrapTable(gextras.InternObject(t)).ColumnDescription(column)
}

func (t treeViewAccessible) ColumnExtentAt(row int, column int) int {
	return atk.WrapTable(gextras.InternObject(t)).ColumnExtentAt(row, column)
}

func (t treeViewAccessible) ColumnHeader(column int) atk.Object {
	return atk.WrapTable(gextras.InternObject(t)).ColumnHeader(column)
}

func (t treeViewAccessible) IndexAt(row int, column int) int {
	return atk.WrapTable(gextras.InternObject(t)).IndexAt(row, column)
}

func (t treeViewAccessible) NColumns() int {
	return atk.WrapTable(gextras.InternObject(t)).NColumns()
}

func (t treeViewAccessible) NRows() int {
	return atk.WrapTable(gextras.InternObject(t)).NRows()
}

func (t treeViewAccessible) RowAtIndex(index_ int) int {
	return atk.WrapTable(gextras.InternObject(t)).RowAtIndex(index_)
}

func (t treeViewAccessible) RowDescription(row int) string {
	return atk.WrapTable(gextras.InternObject(t)).RowDescription(row)
}

func (t treeViewAccessible) RowExtentAt(row int, column int) int {
	return atk.WrapTable(gextras.InternObject(t)).RowExtentAt(row, column)
}

func (t treeViewAccessible) RowHeader(row int) atk.Object {
	return atk.WrapTable(gextras.InternObject(t)).RowHeader(row)
}

func (t treeViewAccessible) SelectedColumns(selected **int) int {
	return atk.WrapTable(gextras.InternObject(t)).SelectedColumns(selected)
}

func (t treeViewAccessible) SelectedRows(selected **int) int {
	return atk.WrapTable(gextras.InternObject(t)).SelectedRows(selected)
}

func (t treeViewAccessible) Summary() atk.Object {
	return atk.WrapTable(gextras.InternObject(t)).Summary()
}

func (t treeViewAccessible) IsColumnSelected(column int) bool {
	return atk.WrapTable(gextras.InternObject(t)).IsColumnSelected(column)
}

func (t treeViewAccessible) IsRowSelected(row int) bool {
	return atk.WrapTable(gextras.InternObject(t)).IsRowSelected(row)
}

func (t treeViewAccessible) IsSelected(row int, column int) bool {
	return atk.WrapTable(gextras.InternObject(t)).IsSelected(row, column)
}

func (t treeViewAccessible) RefAt(row int, column int) atk.Object {
	return atk.WrapTable(gextras.InternObject(t)).RefAt(row, column)
}

func (t treeViewAccessible) RemoveColumnSelection(column int) bool {
	return atk.WrapTable(gextras.InternObject(t)).RemoveColumnSelection(column)
}

func (t treeViewAccessible) RemoveRowSelection(row int) bool {
	return atk.WrapTable(gextras.InternObject(t)).RemoveRowSelection(row)
}

func (t treeViewAccessible) SetCaption(caption atk.Object) {
	atk.WrapTable(gextras.InternObject(t)).SetCaption(caption)
}

func (t treeViewAccessible) SetColumnDescription(column int, description string) {
	atk.WrapTable(gextras.InternObject(t)).SetColumnDescription(column, description)
}

func (t treeViewAccessible) SetColumnHeader(column int, header atk.Object) {
	atk.WrapTable(gextras.InternObject(t)).SetColumnHeader(column, header)
}

func (t treeViewAccessible) SetRowDescription(row int, description string) {
	atk.WrapTable(gextras.InternObject(t)).SetRowDescription(row, description)
}

func (t treeViewAccessible) SetRowHeader(row int, header atk.Object) {
	atk.WrapTable(gextras.InternObject(t)).SetRowHeader(row, header)
}

func (t treeViewAccessible) SetSummary(accessible atk.Object) {
	atk.WrapTable(gextras.InternObject(t)).SetSummary(accessible)
}

func (p treeViewAccessible) Activate(cell CellAccessible) {
	WrapCellAccessibleParent(gextras.InternObject(p)).Activate(cell)
}

func (p treeViewAccessible) Edit(cell CellAccessible) {
	WrapCellAccessibleParent(gextras.InternObject(p)).Edit(cell)
}

func (p treeViewAccessible) ExpandCollapse(cell CellAccessible) {
	WrapCellAccessibleParent(gextras.InternObject(p)).ExpandCollapse(cell)
}

func (p treeViewAccessible) CellArea(cell CellAccessible) gdk.Rectangle {
	return WrapCellAccessibleParent(gextras.InternObject(p)).CellArea(cell)
}

func (p treeViewAccessible) CellExtents(cell CellAccessible, coordType atk.CoordType) (x int, y int, width int, height int) {
	return WrapCellAccessibleParent(gextras.InternObject(p)).CellExtents(cell, coordType)
}

func (p treeViewAccessible) CellPosition(cell CellAccessible) (row int, column int) {
	return WrapCellAccessibleParent(gextras.InternObject(p)).CellPosition(cell)
}

func (p treeViewAccessible) ChildIndex(cell CellAccessible) int {
	return WrapCellAccessibleParent(gextras.InternObject(p)).ChildIndex(cell)
}

func (p treeViewAccessible) RendererState(cell CellAccessible) CellRendererState {
	return WrapCellAccessibleParent(gextras.InternObject(p)).RendererState(cell)
}

func (p treeViewAccessible) GrabFocus(cell CellAccessible) bool {
	return WrapCellAccessibleParent(gextras.InternObject(p)).GrabFocus(cell)
}

func (p treeViewAccessible) UpdateRelationset(cell CellAccessible, relationset atk.RelationSet) {
	WrapCellAccessibleParent(gextras.InternObject(p)).UpdateRelationset(cell, relationset)
}
