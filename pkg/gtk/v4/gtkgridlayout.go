// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkgridlayout.go.
var (
	GTypeGridLayout      = coreglib.Type(C.gtk_grid_layout_get_type())
	GTypeGridLayoutChild = coreglib.Type(C.gtk_grid_layout_child_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeGridLayout, F: marshalGridLayout},
		{T: GTypeGridLayoutChild, F: marshalGridLayoutChild},
	})
}

// GridLayoutOverrider contains methods that are overridable.
type GridLayoutOverrider interface {
}

// GridLayout: GtkGridLayout is a layout manager which arranges child widgets in
// rows and columns.
//
// Children have an "attach point" defined by the horizontal and vertical index
// of the cell they occupy; children can span multiple rows or columns. The
// layout properties for setting the attach points and spans are set using the
// gtk.GridLayoutChild associated to each child widget.
//
// The behaviour of GtkGridLayout when several children occupy the same grid
// cell is undefined.
//
// GtkGridLayout can be used like a GtkBoxLayout if all children are attached to
// the same row or column; however, if you only ever need a single row or
// column, you should consider using GtkBoxLayout.
type GridLayout struct {
	_ [0]func() // equal guard
	LayoutManager
}

var (
	_ LayoutManagerer = (*GridLayout)(nil)
)

func classInitGridLayouter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGridLayout(obj *coreglib.Object) *GridLayout {
	return &GridLayout{
		LayoutManager: LayoutManager{
			Object: obj,
		},
	}
}

func marshalGridLayout(p uintptr) (interface{}, error) {
	return wrapGridLayout(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewGridLayout creates a new GtkGridLayout.
//
// The function returns the following values:
//
//    - gridLayout: newly created GtkGridLayout.
//
func NewGridLayout() *GridLayout {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "GridLayout").InvokeMethod("new_GridLayout", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _gridLayout *GridLayout // out

	_gridLayout = wrapGridLayout(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gridLayout
}

// ColumnHomogeneous checks whether all columns of grid should have the same
// width.
//
// The function returns the following values:
//
//    - ok: TRUE if the columns are homogeneous, and FALSE otherwise.
//
func (grid *GridLayout) ColumnHomogeneous() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))
	*(**GridLayout)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "GridLayout").InvokeMethod("get_column_homogeneous", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(grid)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ColumnSpacing retrieves the spacing set with
// gtk_grid_layout_set_column_spacing().
//
// The function returns the following values:
//
//    - guint: spacing between consecutive columns.
//
func (grid *GridLayout) ColumnSpacing() uint {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))
	*(**GridLayout)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "GridLayout").InvokeMethod("get_column_spacing", args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(grid)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// RowHomogeneous checks whether all rows of grid should have the same height.
//
// The function returns the following values:
//
//    - ok: TRUE if the rows are homogeneous, and FALSE otherwise.
//
func (grid *GridLayout) RowHomogeneous() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))
	*(**GridLayout)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "GridLayout").InvokeMethod("get_row_homogeneous", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(grid)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowSpacing retrieves the spacing set with gtk_grid_layout_set_row_spacing().
//
// The function returns the following values:
//
//    - guint: spacing between consecutive rows.
//
func (grid *GridLayout) RowSpacing() uint {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))
	*(**GridLayout)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "GridLayout").InvokeMethod("get_row_spacing", args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(grid)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SetColumnHomogeneous sets whether all columns of grid should have the same
// width.
//
// The function takes the following parameters:
//
//    - homogeneous: TRUE to make columns homogeneous.
//
func (grid *GridLayout) SetColumnHomogeneous(homogeneous bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}
	*(**GridLayout)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "GridLayout").InvokeMethod("set_column_homogeneous", args[:], nil)

	runtime.KeepAlive(grid)
	runtime.KeepAlive(homogeneous)
}

// SetColumnSpacing sets the amount of space to insert between consecutive
// columns.
//
// The function takes the following parameters:
//
//    - spacing: amount of space between columns, in pixels.
//
func (grid *GridLayout) SetColumnSpacing(spacing uint) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))
	_arg1 = C.guint(spacing)
	*(**GridLayout)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "GridLayout").InvokeMethod("set_column_spacing", args[:], nil)

	runtime.KeepAlive(grid)
	runtime.KeepAlive(spacing)
}

// SetRowHomogeneous sets whether all rows of grid should have the same height.
//
// The function takes the following parameters:
//
//    - homogeneous: TRUE to make rows homogeneous.
//
func (grid *GridLayout) SetRowHomogeneous(homogeneous bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}
	*(**GridLayout)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "GridLayout").InvokeMethod("set_row_homogeneous", args[:], nil)

	runtime.KeepAlive(grid)
	runtime.KeepAlive(homogeneous)
}

// SetRowSpacing sets the amount of space to insert between consecutive rows.
//
// The function takes the following parameters:
//
//    - spacing: amount of space between rows, in pixels.
//
func (grid *GridLayout) SetRowSpacing(spacing uint) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))
	_arg1 = C.guint(spacing)
	*(**GridLayout)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "GridLayout").InvokeMethod("set_row_spacing", args[:], nil)

	runtime.KeepAlive(grid)
	runtime.KeepAlive(spacing)
}

// GridLayoutChildOverrider contains methods that are overridable.
type GridLayoutChildOverrider interface {
}

// GridLayoutChild: GtkLayoutChild subclass for children in a GtkGridLayout.
type GridLayoutChild struct {
	_ [0]func() // equal guard
	LayoutChild
}

var (
	_ LayoutChilder = (*GridLayoutChild)(nil)
)

func classInitGridLayoutChilder(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGridLayoutChild(obj *coreglib.Object) *GridLayoutChild {
	return &GridLayoutChild{
		LayoutChild: LayoutChild{
			Object: obj,
		},
	}
}

func marshalGridLayoutChild(p uintptr) (interface{}, error) {
	return wrapGridLayoutChild(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
