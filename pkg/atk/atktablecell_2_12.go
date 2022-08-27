// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// AtkObject* _gotk4_atk1_TableCell_virtual_get_table(void* fnptr, AtkTableCell* arg0) {
//   return ((AtkObject* (*)(AtkTableCell*))(fnptr))(arg0);
// };
// gboolean _gotk4_atk1_TableCell_virtual_get_position(void* fnptr, AtkTableCell* arg0, gint* arg1, gint* arg2) {
//   return ((gboolean (*)(AtkTableCell*, gint*, gint*))(fnptr))(arg0, arg1, arg2);
// };
// gboolean _gotk4_atk1_TableCell_virtual_get_row_column_span(void* fnptr, AtkTableCell* arg0, gint* arg1, gint* arg2, gint* arg3, gint* arg4) {
//   return ((gboolean (*)(AtkTableCell*, gint*, gint*, gint*, gint*))(fnptr))(arg0, arg1, arg2, arg3, arg4);
// };
// gint _gotk4_atk1_TableCell_virtual_get_column_span(void* fnptr, AtkTableCell* arg0) {
//   return ((gint (*)(AtkTableCell*))(fnptr))(arg0);
// };
// gint _gotk4_atk1_TableCell_virtual_get_row_span(void* fnptr, AtkTableCell* arg0) {
//   return ((gint (*)(AtkTableCell*))(fnptr))(arg0);
// };
import "C"

// ColumnSpan returns the number of columns occupied by this cell accessible.
//
// The function returns the following values:
//
//    - gint representing the number of columns occupied by this cell, or 0 if
//      the cell does not implement this method.
//
func (cell *TableCell) ColumnSpan() int {
	var _arg0 *C.AtkTableCell // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkTableCell)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	_cret = C.atk_table_cell_get_column_span(_arg0)
	runtime.KeepAlive(cell)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Position retrieves the tabular position of this cell.
//
// The function returns the following values:
//
//    - row of the given cell.
//    - column of the given cell.
//    - ok: TRUE if successful; FALSE otherwise.
//
func (cell *TableCell) Position() (row, column int, ok bool) {
	var _arg0 *C.AtkTableCell // out
	var _arg1 C.gint          // in
	var _arg2 C.gint          // in
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkTableCell)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	_cret = C.atk_table_cell_get_position(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(cell)

	var _row int    // out
	var _column int // out
	var _ok bool    // out

	_row = int(_arg1)
	_column = int(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _row, _column, _ok
}

// RowColumnSpan gets the row and column indexes and span of this cell
// accessible.
//
// Note: If the object does not implement this function, then, by default, atk
// will implement this function by calling get_row_span and get_column_span on
// the object.
//
// The function returns the following values:
//
//    - row index of the given cell.
//    - column index of the given cell.
//    - rowSpan: number of rows occupied by this cell.
//    - columnSpan: number of columns occupied by this cell.
//    - ok: TRUE if successful; FALSE otherwise.
//
func (cell *TableCell) RowColumnSpan() (row, column, rowSpan, columnSpan int, ok bool) {
	var _arg0 *C.AtkTableCell // out
	var _arg1 C.gint          // in
	var _arg2 C.gint          // in
	var _arg3 C.gint          // in
	var _arg4 C.gint          // in
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkTableCell)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	_cret = C.atk_table_cell_get_row_column_span(_arg0, &_arg1, &_arg2, &_arg3, &_arg4)
	runtime.KeepAlive(cell)

	var _row int        // out
	var _column int     // out
	var _rowSpan int    // out
	var _columnSpan int // out
	var _ok bool        // out

	_row = int(_arg1)
	_column = int(_arg2)
	_rowSpan = int(_arg3)
	_columnSpan = int(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _row, _column, _rowSpan, _columnSpan, _ok
}

// RowSpan returns the number of rows occupied by this cell accessible.
//
// The function returns the following values:
//
//    - gint representing the number of rows occupied by this cell, or 0 if the
//      cell does not implement this method.
//
func (cell *TableCell) RowSpan() int {
	var _arg0 *C.AtkTableCell // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkTableCell)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	_cret = C.atk_table_cell_get_row_span(_arg0)
	runtime.KeepAlive(cell)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Table returns a reference to the accessible of the containing table.
//
// The function returns the following values:
//
//    - object: atk object for the containing table.
//
func (cell *TableCell) Table() *AtkObject {
	var _arg0 *C.AtkTableCell // out
	var _cret *C.AtkObject    // in

	_arg0 = (*C.AtkTableCell)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	_cret = C.atk_table_cell_get_table(_arg0)
	runtime.KeepAlive(cell)

	var _object *AtkObject // out

	_object = wrapObject(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _object
}

// columnSpan returns the number of columns occupied by this cell accessible.
//
// The function returns the following values:
//
//    - gint representing the number of columns occupied by this cell, or 0 if
//      the cell does not implement this method.
//
func (cell *TableCell) columnSpan() int {
	gclass := (*C.AtkTableCellIface)(coreglib.PeekParentClass(cell))
	fnarg := gclass.get_column_span

	var _arg0 *C.AtkTableCell // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkTableCell)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	_cret = C._gotk4_atk1_TableCell_virtual_get_column_span(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(cell)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Position retrieves the tabular position of this cell.
//
// The function returns the following values:
//
//    - row of the given cell.
//    - column of the given cell.
//    - ok: TRUE if successful; FALSE otherwise.
//
func (cell *TableCell) position() (row, column int, ok bool) {
	gclass := (*C.AtkTableCellIface)(coreglib.PeekParentClass(cell))
	fnarg := gclass.get_position

	var _arg0 *C.AtkTableCell // out
	var _arg1 C.gint          // in
	var _arg2 C.gint          // in
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkTableCell)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	_cret = C._gotk4_atk1_TableCell_virtual_get_position(unsafe.Pointer(fnarg), _arg0, &_arg1, &_arg2)
	runtime.KeepAlive(cell)

	var _row int    // out
	var _column int // out
	var _ok bool    // out

	_row = int(_arg1)
	_column = int(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _row, _column, _ok
}

// rowColumnSpan gets the row and column indexes and span of this cell
// accessible.
//
// Note: If the object does not implement this function, then, by default, atk
// will implement this function by calling get_row_span and get_column_span on
// the object.
//
// The function returns the following values:
//
//    - row index of the given cell.
//    - column index of the given cell.
//    - rowSpan: number of rows occupied by this cell.
//    - columnSpan: number of columns occupied by this cell.
//    - ok: TRUE if successful; FALSE otherwise.
//
func (cell *TableCell) rowColumnSpan() (row, column, rowSpan, columnSpan int, ok bool) {
	gclass := (*C.AtkTableCellIface)(coreglib.PeekParentClass(cell))
	fnarg := gclass.get_row_column_span

	var _arg0 *C.AtkTableCell // out
	var _arg1 C.gint          // in
	var _arg2 C.gint          // in
	var _arg3 C.gint          // in
	var _arg4 C.gint          // in
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkTableCell)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	_cret = C._gotk4_atk1_TableCell_virtual_get_row_column_span(unsafe.Pointer(fnarg), _arg0, &_arg1, &_arg2, &_arg3, &_arg4)
	runtime.KeepAlive(cell)

	var _row int        // out
	var _column int     // out
	var _rowSpan int    // out
	var _columnSpan int // out
	var _ok bool        // out

	_row = int(_arg1)
	_column = int(_arg2)
	_rowSpan = int(_arg3)
	_columnSpan = int(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _row, _column, _rowSpan, _columnSpan, _ok
}

// rowSpan returns the number of rows occupied by this cell accessible.
//
// The function returns the following values:
//
//    - gint representing the number of rows occupied by this cell, or 0 if the
//      cell does not implement this method.
//
func (cell *TableCell) rowSpan() int {
	gclass := (*C.AtkTableCellIface)(coreglib.PeekParentClass(cell))
	fnarg := gclass.get_row_span

	var _arg0 *C.AtkTableCell // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkTableCell)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	_cret = C._gotk4_atk1_TableCell_virtual_get_row_span(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(cell)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Table returns a reference to the accessible of the containing table.
//
// The function returns the following values:
//
//    - object: atk object for the containing table.
//
func (cell *TableCell) table() *AtkObject {
	gclass := (*C.AtkTableCellIface)(coreglib.PeekParentClass(cell))
	fnarg := gclass.get_table

	var _arg0 *C.AtkTableCell // out
	var _cret *C.AtkObject    // in

	_arg0 = (*C.AtkTableCell)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	_cret = C._gotk4_atk1_TableCell_virtual_get_table(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(cell)

	var _object *AtkObject // out

	_object = wrapObject(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _object
}

// TableCellIface: atkTableCell is an interface for cells inside an Table.
//
// An instance of this type is always passed by reference.
type TableCellIface struct {
	*tableCellIface
}

// tableCellIface is the struct that's finalized.
type tableCellIface struct {
	native *C.AtkTableCellIface
}
