// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
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

	RendererState(cell CellAccessible) CellRendererState

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
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_activate(arg0, cell)
}

func (p cellAccessibleParent) Edit(cell CellAccessible) {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_edit(arg0, cell)
}

func (p cellAccessibleParent) ExpandCollapse(cell CellAccessible) {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_expand_collapse(arg0, cell)
}

func (p cellAccessibleParent) CellArea(cell CellAccessible) gdk.Rectangle {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	var arg2 C.GdkRectangle
	var ret2 *gdk.Rectangle

	C.gtk_cell_accessible_parent_get_cell_area(arg0, cell, &arg2)

	*ret2 = gdk.WrapRectangle(unsafe.Pointer(arg2))

	return ret2
}

func (p cellAccessibleParent) CellPosition(cell CellAccessible) (row int, column int) {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	var arg2 C.gint
	var ret2 int
	var arg3 C.gint
	var ret3 int

	C.gtk_cell_accessible_parent_get_cell_position(arg0, cell, &arg2, &arg3)

	*ret2 = C.gint(arg2)
	*ret3 = C.gint(arg3)

	return ret2, ret3
}

func (p cellAccessibleParent) ChildIndex(cell CellAccessible) int {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	var cret C.int
	var ret1 int

	cret = C.gtk_cell_accessible_parent_get_child_index(arg0, cell)

	ret1 = C.int(cret)

	return ret1
}

func (p cellAccessibleParent) RendererState(cell CellAccessible) CellRendererState {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	var cret C.GtkCellRendererState
	var ret1 CellRendererState

	cret = C.gtk_cell_accessible_parent_get_renderer_state(arg0, cell)

	ret1 = CellRendererState(cret)

	return ret1
}

func (p cellAccessibleParent) GrabFocus(cell CellAccessible) bool {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_cell_accessible_parent_grab_focus(arg0, cell)

	ret1 = C.bool(cret) != C.false

	return ret1
}
