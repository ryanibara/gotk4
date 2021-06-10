// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_cell_accessible_parent_get_type()), F: marshalCellAccessibleParent},
	})
}

// CellAccessibleParentOverrider contains methods that are overridable. This
// interface is a subset of the interface CellAccessibleParent.
type CellAccessibleParentOverrider interface {
	Activate(cell CellAccessible)

	Edit(cell CellAccessible)

	ExpandCollapse(cell CellAccessible)

	CellArea(cell CellAccessible) gdk.Rectangle

	CellPosition(cell CellAccessible) (row int, column int)

	ChildIndex(cell CellAccessible) int

	GrabFocus(cell CellAccessible) bool
}

type CellAccessibleParent interface {
	gextras.Objector
	CellAccessibleParentOverrider
}

// cellAccessibleParent implements the CellAccessibleParent interface.
type cellAccessibleParent struct {
	gextras.Objector
}

var _ CellAccessibleParent = (*cellAccessibleParent)(nil)

// WrapCellAccessibleParent wraps a GObject to a type that implements interface
// CellAccessibleParent. It is primarily used internally.
func WrapCellAccessibleParent(obj *externglib.Object) CellAccessibleParent {
	return CellAccessibleParent{
		Objector: obj,
	}
}

func marshalCellAccessibleParent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellAccessibleParent(obj), nil
}

func (p cellAccessibleParent) Activate(cell CellAccessible) {
	var _arg0 *C.GtkCellAccessibleParent
	var _arg1 *C.GtkCellAccessible

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_activate(_arg0, _arg1)
}

func (p cellAccessibleParent) Edit(cell CellAccessible) {
	var _arg0 *C.GtkCellAccessibleParent
	var _arg1 *C.GtkCellAccessible

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_edit(_arg0, _arg1)
}

func (p cellAccessibleParent) ExpandCollapse(cell CellAccessible) {
	var _arg0 *C.GtkCellAccessibleParent
	var _arg1 *C.GtkCellAccessible

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_expand_collapse(_arg0, _arg1)
}

func (p cellAccessibleParent) CellArea(cell CellAccessible) gdk.Rectangle {
	var _arg0 *C.GtkCellAccessibleParent
	var _arg1 *C.GtkCellAccessible

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	var _cellRect gdk.Rectangle

	C.gtk_cell_accessible_parent_get_cell_area(_arg0, _arg1, (*C.GdkRectangle)(unsafe.Pointer(&_cellRect)))

	return _cellRect
}

func (p cellAccessibleParent) CellPosition(cell CellAccessible) (row int, column int) {
	var _arg0 *C.GtkCellAccessibleParent
	var _arg1 *C.GtkCellAccessible

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	var _arg2 C.gint
	var _arg3 C.gint

	C.gtk_cell_accessible_parent_get_cell_position(_arg0, _arg1, &_arg2, &_arg3)

	var _row int
	var _column int

	_row = (int)(_arg2)
	_column = (int)(_arg3)

	return _row, _column
}

func (p cellAccessibleParent) ChildIndex(cell CellAccessible) int {
	var _arg0 *C.GtkCellAccessibleParent
	var _arg1 *C.GtkCellAccessible

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	var _cret C.int

	_cret = C.gtk_cell_accessible_parent_get_child_index(_arg0, _arg1)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

func (p cellAccessibleParent) GrabFocus(cell CellAccessible) bool {
	var _arg0 *C.GtkCellAccessibleParent
	var _arg1 *C.GtkCellAccessible

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	var _cret C.gboolean

	_cret = C.gtk_cell_accessible_parent_grab_focus(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}
