// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_area_context_get_type()), F: marshalCellAreaContext},
	})
}

// CellAreaContext: the CellAreaContext object is created by a given CellArea
// implementation via its CellAreaClass.create_context() virtual method and is
// used to store cell sizes and alignments for a series of TreeModel rows that
// are requested and rendered in the same context.
//
// CellLayout widgets can create any number of contexts in which to request and
// render groups of data rows. However, it’s important that the same context
// which was used to request sizes for a given TreeModel row also be used for
// the same row when calling other CellArea APIs such as gtk_cell_area_render()
// and gtk_cell_area_event().
type CellAreaContext interface {
	gextras.Objector

	// Allocate allocates a width and/or a height for all rows which are to be
	// rendered with @context.
	//
	// Usually allocation is performed only horizontally or sometimes vertically
	// since a group of rows are usually rendered side by side vertically or
	// horizontally and share either the same width or the same height.
	// Sometimes they are allocated in both horizontal and vertical orientations
	// producing a homogeneous effect of the rows. This is generally the case
	// for TreeView when TreeView:fixed-height-mode is enabled.
	Allocate(width int, height int)
	// Allocation fetches the current allocation size for @context.
	//
	// If the context was not allocated in width or height, or if the context
	// was recently reset with gtk_cell_area_context_reset(), the returned value
	// will be -1.
	Allocation() (width int, height int)
	// Area fetches the CellArea this @context was created by.
	//
	// This is generally unneeded by layouting widgets; however, it is important
	// for the context implementation itself to fetch information about the area
	// it is being used for.
	//
	// For instance at CellAreaContextClass.allocate() time it’s important to
	// know details about any cell spacing that the CellArea is configured with
	// in order to compute a proper allocation.
	Area() CellArea
	// PreferredHeight gets the accumulative preferred height for all rows which
	// have been requested with this context.
	//
	// After gtk_cell_area_context_reset() is called and/or before ever
	// requesting the size of a CellArea, the returned values are 0.
	PreferredHeight() (minimumHeight int, naturalHeight int)
	// PreferredHeightForWidth gets the accumulative preferred height for @width
	// for all rows which have been requested for the same said @width with this
	// context.
	//
	// After gtk_cell_area_context_reset() is called and/or before ever
	// requesting the size of a CellArea, the returned values are -1.
	PreferredHeightForWidth(width int) (minimumHeight int, naturalHeight int)
	// PreferredWidth gets the accumulative preferred width for all rows which
	// have been requested with this context.
	//
	// After gtk_cell_area_context_reset() is called and/or before ever
	// requesting the size of a CellArea, the returned values are 0.
	PreferredWidth() (minimumWidth int, naturalWidth int)
	// PreferredWidthForHeight gets the accumulative preferred width for @height
	// for all rows which have been requested for the same said @height with
	// this context.
	//
	// After gtk_cell_area_context_reset() is called and/or before ever
	// requesting the size of a CellArea, the returned values are -1.
	PreferredWidthForHeight(height int) (minimumWidth int, naturalWidth int)
	// PushPreferredHeight causes the minimum and/or natural height to grow if
	// the new proposed sizes exceed the current minimum and natural height.
	//
	// This is used by CellAreaContext implementations during the request
	// process over a series of TreeModel rows to progressively push the
	// requested height over a series of gtk_cell_area_get_preferred_height()
	// requests.
	PushPreferredHeight(minimumHeight int, naturalHeight int)
	// PushPreferredWidth causes the minimum and/or natural width to grow if the
	// new proposed sizes exceed the current minimum and natural width.
	//
	// This is used by CellAreaContext implementations during the request
	// process over a series of TreeModel rows to progressively push the
	// requested width over a series of gtk_cell_area_get_preferred_width()
	// requests.
	PushPreferredWidth(minimumWidth int, naturalWidth int)
	// Reset resets any previously cached request and allocation data.
	//
	// When underlying TreeModel data changes its important to reset the context
	// if the content size is allowed to shrink. If the content size is only
	// allowed to grow (this is usually an option for views rendering large data
	// stores as a measure of optimization), then only the row that changed or
	// was inserted needs to be (re)requested with
	// gtk_cell_area_get_preferred_width().
	//
	// When the new overall size of the context requires that the allocated size
	// changes (or whenever this allocation changes at all), the variable row
	// sizes need to be re-requested for every row.
	//
	// For instance, if the rows are displayed all with the same width from top
	// to bottom then a change in the allocated width necessitates a
	// recalculation of all the displayed row heights using
	// gtk_cell_area_get_preferred_height_for_width().
	Reset()
}

// cellAreaContext implements the CellAreaContext interface.
type cellAreaContext struct {
	gextras.Objector
}

var _ CellAreaContext = (*cellAreaContext)(nil)

// WrapCellAreaContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellAreaContext(obj *externglib.Object) CellAreaContext {
	return CellAreaContext{
		Objector: obj,
	}
}

func marshalCellAreaContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellAreaContext(obj), nil
}

// Allocate allocates a width and/or a height for all rows which are to be
// rendered with @context.
//
// Usually allocation is performed only horizontally or sometimes vertically
// since a group of rows are usually rendered side by side vertically or
// horizontally and share either the same width or the same height.
// Sometimes they are allocated in both horizontal and vertical orientations
// producing a homogeneous effect of the rows. This is generally the case
// for TreeView when TreeView:fixed-height-mode is enabled.
func (c cellAreaContext) Allocate(width int, height int) {
	var arg0 *C.GtkCellAreaContext
	var arg1 C.int
	var arg2 C.int

	arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(c.Native()))
	arg1 = C.int(width)
	arg2 = C.int(height)

	C.gtk_cell_area_context_allocate(arg0, width, height)
}

// Allocation fetches the current allocation size for @context.
//
// If the context was not allocated in width or height, or if the context
// was recently reset with gtk_cell_area_context_reset(), the returned value
// will be -1.
func (c cellAreaContext) Allocation() (width int, height int) {
	var arg0 *C.GtkCellAreaContext

	arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(c.Native()))

	var arg1 C.int
	var ret1 int
	var arg2 C.int
	var ret2 int

	C.gtk_cell_area_context_get_allocation(arg0, &arg1, &arg2)

	*ret1 = C.int(arg1)
	*ret2 = C.int(arg2)

	return ret1, ret2
}

// Area fetches the CellArea this @context was created by.
//
// This is generally unneeded by layouting widgets; however, it is important
// for the context implementation itself to fetch information about the area
// it is being used for.
//
// For instance at CellAreaContextClass.allocate() time it’s important to
// know details about any cell spacing that the CellArea is configured with
// in order to compute a proper allocation.
func (c cellAreaContext) Area() CellArea {
	var arg0 *C.GtkCellAreaContext

	arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(c.Native()))

	var cret *C.GtkCellArea
	var ret1 CellArea

	cret = C.gtk_cell_area_context_get_area(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(CellArea)

	return ret1
}

// PreferredHeight gets the accumulative preferred height for all rows which
// have been requested with this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever
// requesting the size of a CellArea, the returned values are 0.
func (c cellAreaContext) PreferredHeight() (minimumHeight int, naturalHeight int) {
	var arg0 *C.GtkCellAreaContext

	arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(c.Native()))

	var arg1 C.int
	var ret1 int
	var arg2 C.int
	var ret2 int

	C.gtk_cell_area_context_get_preferred_height(arg0, &arg1, &arg2)

	*ret1 = C.int(arg1)
	*ret2 = C.int(arg2)

	return ret1, ret2
}

// PreferredHeightForWidth gets the accumulative preferred height for @width
// for all rows which have been requested for the same said @width with this
// context.
//
// After gtk_cell_area_context_reset() is called and/or before ever
// requesting the size of a CellArea, the returned values are -1.
func (c cellAreaContext) PreferredHeightForWidth(width int) (minimumHeight int, naturalHeight int) {
	var arg0 *C.GtkCellAreaContext
	var arg1 C.int

	arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(c.Native()))
	arg1 = C.int(width)

	var arg2 C.int
	var ret2 int
	var arg3 C.int
	var ret3 int

	C.gtk_cell_area_context_get_preferred_height_for_width(arg0, width, &arg2, &arg3)

	*ret2 = C.int(arg2)
	*ret3 = C.int(arg3)

	return ret2, ret3
}

// PreferredWidth gets the accumulative preferred width for all rows which
// have been requested with this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever
// requesting the size of a CellArea, the returned values are 0.
func (c cellAreaContext) PreferredWidth() (minimumWidth int, naturalWidth int) {
	var arg0 *C.GtkCellAreaContext

	arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(c.Native()))

	var arg1 C.int
	var ret1 int
	var arg2 C.int
	var ret2 int

	C.gtk_cell_area_context_get_preferred_width(arg0, &arg1, &arg2)

	*ret1 = C.int(arg1)
	*ret2 = C.int(arg2)

	return ret1, ret2
}

// PreferredWidthForHeight gets the accumulative preferred width for @height
// for all rows which have been requested for the same said @height with
// this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever
// requesting the size of a CellArea, the returned values are -1.
func (c cellAreaContext) PreferredWidthForHeight(height int) (minimumWidth int, naturalWidth int) {
	var arg0 *C.GtkCellAreaContext
	var arg1 C.int

	arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(c.Native()))
	arg1 = C.int(height)

	var arg2 C.int
	var ret2 int
	var arg3 C.int
	var ret3 int

	C.gtk_cell_area_context_get_preferred_width_for_height(arg0, height, &arg2, &arg3)

	*ret2 = C.int(arg2)
	*ret3 = C.int(arg3)

	return ret2, ret3
}

// PushPreferredHeight causes the minimum and/or natural height to grow if
// the new proposed sizes exceed the current minimum and natural height.
//
// This is used by CellAreaContext implementations during the request
// process over a series of TreeModel rows to progressively push the
// requested height over a series of gtk_cell_area_get_preferred_height()
// requests.
func (c cellAreaContext) PushPreferredHeight(minimumHeight int, naturalHeight int) {
	var arg0 *C.GtkCellAreaContext
	var arg1 C.int
	var arg2 C.int

	arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(c.Native()))
	arg1 = C.int(minimumHeight)
	arg2 = C.int(naturalHeight)

	C.gtk_cell_area_context_push_preferred_height(arg0, minimumHeight, naturalHeight)
}

// PushPreferredWidth causes the minimum and/or natural width to grow if the
// new proposed sizes exceed the current minimum and natural width.
//
// This is used by CellAreaContext implementations during the request
// process over a series of TreeModel rows to progressively push the
// requested width over a series of gtk_cell_area_get_preferred_width()
// requests.
func (c cellAreaContext) PushPreferredWidth(minimumWidth int, naturalWidth int) {
	var arg0 *C.GtkCellAreaContext
	var arg1 C.int
	var arg2 C.int

	arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(c.Native()))
	arg1 = C.int(minimumWidth)
	arg2 = C.int(naturalWidth)

	C.gtk_cell_area_context_push_preferred_width(arg0, minimumWidth, naturalWidth)
}

// Reset resets any previously cached request and allocation data.
//
// When underlying TreeModel data changes its important to reset the context
// if the content size is allowed to shrink. If the content size is only
// allowed to grow (this is usually an option for views rendering large data
// stores as a measure of optimization), then only the row that changed or
// was inserted needs to be (re)requested with
// gtk_cell_area_get_preferred_width().
//
// When the new overall size of the context requires that the allocated size
// changes (or whenever this allocation changes at all), the variable row
// sizes need to be re-requested for every row.
//
// For instance, if the rows are displayed all with the same width from top
// to bottom then a change in the allocated width necessitates a
// recalculation of all the displayed row heights using
// gtk_cell_area_get_preferred_height_for_width().
func (c cellAreaContext) Reset() {
	var arg0 *C.GtkCellAreaContext

	arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(c.Native()))

	C.gtk_cell_area_context_reset(arg0)
}

type CellAreaContextPrivate struct {
	native C.GtkCellAreaContextPrivate
}

// WrapCellAreaContextPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapCellAreaContextPrivate(ptr unsafe.Pointer) *CellAreaContextPrivate {
	if ptr == nil {
		return nil
	}

	return (*CellAreaContextPrivate)(ptr)
}

func marshalCellAreaContextPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapCellAreaContextPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *CellAreaContextPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}
