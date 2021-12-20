// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_cell_accessible_parent_get_type()), F: marshalCellAccessibleParenter},
	})
}

// CellAccessibleParentOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CellAccessibleParentOverrider interface {
	// The function takes the following parameters:
	//
	Activate(cell *CellAccessible)
	// The function takes the following parameters:
	//
	Edit(cell *CellAccessible)
	// The function takes the following parameters:
	//
	ExpandCollapse(cell *CellAccessible)
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	CellArea(cell *CellAccessible) *gdk.Rectangle
	// The function takes the following parameters:
	//
	//    - cell
	//    - coordType
	//
	// The function returns the following values:
	//
	//    - x
	//    - y
	//    - width
	//    - height
	//
	CellExtents(cell *CellAccessible, coordType atk.CoordType) (x int, y int, width int, height int)
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	//    - row
	//    - column
	//
	CellPosition(cell *CellAccessible) (row int, column int)
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	ChildIndex(cell *CellAccessible) int
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	RendererState(cell *CellAccessible) CellRendererState
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	GrabFocus(cell *CellAccessible) bool
	// The function takes the following parameters:
	//
	//    - cell
	//    - relationset
	//
	UpdateRelationset(cell *CellAccessible, relationset *atk.RelationSet)
}

type CellAccessibleParent struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*CellAccessibleParent)(nil)
)

// CellAccessibleParenter describes CellAccessibleParent's interface methods.
type CellAccessibleParenter interface {
	externglib.Objector

	Activate(cell *CellAccessible)
	Edit(cell *CellAccessible)
	ExpandCollapse(cell *CellAccessible)
	CellArea(cell *CellAccessible) *gdk.Rectangle
	CellExtents(cell *CellAccessible, coordType atk.CoordType) (x int, y int, width int, height int)
	CellPosition(cell *CellAccessible) (row int, column int)
	ChildIndex(cell *CellAccessible) int
	RendererState(cell *CellAccessible) CellRendererState
	GrabFocus(cell *CellAccessible) bool
	UpdateRelationset(cell *CellAccessible, relationset *atk.RelationSet)
}

var _ CellAccessibleParenter = (*CellAccessibleParent)(nil)

func wrapCellAccessibleParent(obj *externglib.Object) *CellAccessibleParent {
	return &CellAccessibleParent{
		Object: obj,
	}
}

func marshalCellAccessibleParenter(p uintptr) (interface{}, error) {
	return wrapCellAccessibleParent(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function takes the following parameters:
//
func (parent *CellAccessibleParent) Activate(cell *CellAccessible) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_activate(_arg0, _arg1)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)
}

// The function takes the following parameters:
//
func (parent *CellAccessibleParent) Edit(cell *CellAccessible) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_edit(_arg0, _arg1)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)
}

// The function takes the following parameters:
//
func (parent *CellAccessibleParent) ExpandCollapse(cell *CellAccessible) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_expand_collapse(_arg0, _arg1)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (parent *CellAccessibleParent) CellArea(cell *CellAccessible) *gdk.Rectangle {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _arg2 C.GdkRectangle             // in

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_get_cell_area(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)

	var _cellRect *gdk.Rectangle // out

	_cellRect = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _cellRect
}

// The function takes the following parameters:
//
//    - cell
//    - coordType
//
// The function returns the following values:
//
//    - x
//    - y
//    - width
//    - height
//
func (parent *CellAccessibleParent) CellExtents(cell *CellAccessible, coordType atk.CoordType) (x int, y int, width int, height int) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _arg2 C.gint                     // in
	var _arg3 C.gint                     // in
	var _arg4 C.gint                     // in
	var _arg5 C.gint                     // in
	var _arg6 C.AtkCoordType             // out

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))
	_arg6 = C.AtkCoordType(coordType)

	C.gtk_cell_accessible_parent_get_cell_extents(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5, _arg6)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(coordType)

	var _x int      // out
	var _y int      // out
	var _width int  // out
	var _height int // out

	_x = int(_arg2)
	_y = int(_arg3)
	_width = int(_arg4)
	_height = int(_arg5)

	return _x, _y, _width, _height
}

// The function takes the following parameters:
//
// The function returns the following values:
//
//    - row
//    - column
//
func (parent *CellAccessibleParent) CellPosition(cell *CellAccessible) (row int, column int) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _arg2 C.gint                     // in
	var _arg3 C.gint                     // in

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_get_cell_position(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)

	var _row int    // out
	var _column int // out

	_row = int(_arg2)
	_column = int(_arg3)

	return _row, _column
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (parent *CellAccessibleParent) ChildIndex(cell *CellAccessible) int {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _cret C.int                      // in

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_accessible_parent_get_child_index(_arg0, _arg1)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (parent *CellAccessibleParent) RendererState(cell *CellAccessible) CellRendererState {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _cret C.GtkCellRendererState     // in

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_accessible_parent_get_renderer_state(_arg0, _arg1)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)

	var _cellRendererState CellRendererState // out

	_cellRendererState = CellRendererState(_cret)

	return _cellRendererState
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (parent *CellAccessibleParent) GrabFocus(cell *CellAccessible) bool {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_accessible_parent_grab_focus(_arg0, _arg1)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
//    - cell
//    - relationset
//
func (parent *CellAccessibleParent) UpdateRelationset(cell *CellAccessible, relationset *atk.RelationSet) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _arg2 *C.AtkRelationSet          // out

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))
	_arg2 = (*C.AtkRelationSet)(unsafe.Pointer(relationset.Native()))

	C.gtk_cell_accessible_parent_update_relationset(_arg0, _arg1, _arg2)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(relationset)
}
