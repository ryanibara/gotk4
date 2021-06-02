// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdbool.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_get_type()), F: marshalCellRenderer},
	})
}

type CellRendererClassPrivate struct {
	native C.GtkCellRendererClassPrivate
}

// WrapCellRendererClassPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapCellRendererClassPrivate(ptr unsafe.Pointer) *CellRendererClassPrivate {
	if ptr == nil {
		return nil
	}

	return (*CellRendererClassPrivate)(ptr)
}

func marshalCellRendererClassPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapCellRendererClassPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *CellRendererClassPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

type CellRendererPrivate struct {
	native C.GtkCellRendererPrivate
}

// WrapCellRendererPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapCellRendererPrivate(ptr unsafe.Pointer) *CellRendererPrivate {
	if ptr == nil {
		return nil
	}

	return (*CellRendererPrivate)(ptr)
}

func marshalCellRendererPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapCellRendererPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *CellRendererPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// CellRenderer: the CellRenderer is a base class of a set of objects used for
// rendering a cell to a #cairo_t. These objects are used primarily by the
// TreeView widget, though they aren’t tied to them in any specific way. It is
// worth noting that CellRenderer is not a Widget and cannot be treated as such.
//
// The primary use of a CellRenderer is for drawing a certain graphical elements
// on a #cairo_t. Typically, one cell renderer is used to draw many cells on the
// screen. To this extent, it isn’t expected that a CellRenderer keep any
// permanent state around. Instead, any state is set just prior to use using
// #GObjects property system. Then, the cell is measured using
// gtk_cell_renderer_get_size(). Finally, the cell is rendered in the correct
// location using gtk_cell_renderer_render().
//
// There are a number of rules that must be followed when writing a new
// CellRenderer. First and foremost, it’s important that a certain set of
// properties will always yield a cell renderer of the same size, barring a
// Style change. The CellRenderer also has a number of generic properties that
// are expected to be honored by all children.
//
// Beyond merely rendering a cell, cell renderers can optionally provide active
// user interface elements. A cell renderer can be “activatable” like
// CellRendererToggle, which toggles when it gets activated by a mouse click, or
// it can be “editable” like CellRendererText, which allows the user to edit the
// text using a widget implementing the CellEditable interface, e.g. Entry. To
// make a cell renderer activatable or editable, you have to implement the
// CellRendererClass.activate or CellRendererClass.start_editing virtual
// functions, respectively.
//
// Many properties of CellRenderer and its subclasses have a corresponding “set”
// property, e.g. “cell-background-set” corresponds to “cell-background”. These
// “set” properties reflect whether a property has been set or not. You should
// not set them independently.
type CellRenderer interface {
	gextras.Objector

	// AlignedArea gets the aligned area used by @cell inside @cell_area. Used
	// for finding the appropriate edit and focus rectangle.
	AlignedArea(widget Widget, flags CellRendererState, cellArea *gdk.Rectangle) gdk.Rectangle
	// Alignment fills in @xalign and @yalign with the appropriate values of
	// @cell.
	Alignment() (xalign float32, yalign float32)
	// FixedSize fills in @width and @height with the appropriate size of @cell.
	FixedSize() (width int, height int)
	// Padding fills in @xpad and @ypad with the appropriate values of @cell.
	Padding() (xpad int, ypad int)
	// PreferredHeight retreives a renderer’s natural size when rendered to
	// @widget.
	PreferredHeight(widget Widget) (minimumSize int, naturalSize int)
	// PreferredHeightForWidth retreives a cell renderers’s minimum and natural
	// height if it were rendered to @widget with the specified @width.
	PreferredHeightForWidth(widget Widget, width int) (minimumHeight int, naturalHeight int)
	// PreferredSize retrieves the minimum and natural size of a cell taking
	// into account the widget’s preference for height-for-width management.
	PreferredSize(widget Widget) (minimumSize Requisition, naturalSize Requisition)
	// PreferredWidth retreives a renderer’s natural size when rendered to
	// @widget.
	PreferredWidth(widget Widget) (minimumSize int, naturalSize int)
	// PreferredWidthForHeight retreives a cell renderers’s minimum and natural
	// width if it were rendered to @widget with the specified @height.
	PreferredWidthForHeight(widget Widget, height int) (minimumWidth int, naturalWidth int)
	// RequestMode gets whether the cell renderer prefers a height-for-width
	// layout or a width-for-height layout.
	RequestMode() SizeRequestMode
	// Sensitive returns the cell renderer’s sensitivity.
	Sensitive() bool
	// Size obtains the width and height needed to render the cell. Used by view
	// widgets to determine the appropriate size for the cell_area passed to
	// gtk_cell_renderer_render(). If @cell_area is not nil, fills in the x and
	// y offsets (if set) of the cell relative to this location.
	//
	// Please note that the values set in @width and @height, as well as those
	// in @x_offset and @y_offset are inclusive of the xpad and ypad properties.
	Size(widget Widget, cellArea *gdk.Rectangle) (xOffset int, yOffset int, width int, height int)
	// State translates the cell renderer state to StateFlags, based on the cell
	// renderer and widget sensitivity, and the given CellRendererState.
	State(widget Widget, cellState CellRendererState) StateFlags
	// Visible returns the cell renderer’s visibility.
	Visible() bool
	// IsActivatable checks whether the cell renderer can do something when
	// activated.
	IsActivatable() bool
	// Render invokes the virtual render function of the CellRenderer. The three
	// passed-in rectangles are areas in @cr. Most renderers will draw within
	// @cell_area; the xalign, yalign, xpad, and ypad fields of the CellRenderer
	// should be honored with respect to @cell_area. @background_area includes
	// the blank space around the cell, and also the area containing the tree
	// expander; so the @background_area rectangles for all cells tile to cover
	// the entire @window.
	Render(cr *cairo.Context, widget Widget, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState)
	// SetAlignment sets the renderer’s alignment within its available space.
	SetAlignment(xalign float32, yalign float32)
	// SetFixedSize sets the renderer size to be explicit, independent of the
	// properties set.
	SetFixedSize(width int, height int)
	// SetPadding sets the renderer’s padding.
	SetPadding(xpad int, ypad int)
	// SetSensitive sets the cell renderer’s sensitivity.
	SetSensitive(sensitive bool)
	// SetVisible sets the cell renderer’s visibility.
	SetVisible(visible bool)
	// StartEditing starts editing the contents of this @cell, through a new
	// CellEditable widget created by the CellRendererClass.start_editing
	// virtual function.
	StartEditing(event *gdk.Event, widget Widget, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) CellEditable
	// StopEditing informs the cell renderer that the editing is stopped. If
	// @canceled is true, the cell renderer will emit the
	// CellRenderer::editing-canceled signal.
	//
	// This function should be called by cell renderer implementations in
	// response to the CellEditable::editing-done signal of CellEditable.
	StopEditing(canceled bool)
}

// cellRenderer implements the CellRenderer interface.
type cellRenderer struct {
	gextras.Objector
}

var _ CellRenderer = (*cellRenderer)(nil)

// WrapCellRenderer wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellRenderer(obj *externglib.Object) CellRenderer {
	return CellRenderer{
		Objector: obj,
	}
}

func marshalCellRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellRenderer(obj), nil
}

// AlignedArea gets the aligned area used by @cell inside @cell_area. Used
// for finding the appropriate edit and focus rectangle.
func (c cellRenderer) AlignedArea(widget Widget, flags CellRendererState, cellArea *gdk.Rectangle) gdk.Rectangle {
	var arg0 *C.GtkCellRenderer
	var arg1 *C.GtkWidget
	var arg2 C.GtkCellRendererState
	var arg3 *C.GdkRectangle
	var arg4 *C.GdkRectangle // out

	arg0 = (*C.GtkCellRenderer)(c.Native())
	arg1 = (*C.GtkWidget)(widget.Native())
	arg2 = (C.GtkCellRendererState)(flags)
	arg3 = (*C.GdkRectangle)(cellArea.Native())

	C.gtk_cell_renderer_get_aligned_area(arg0, arg1, arg2, arg3, &arg4)

	var ret0 *gdk.Rectangle

	{
		ret0 = gdk.WrapRectangle(unsafe.Pointer(arg4))
	}

	return ret0
}

// Alignment fills in @xalign and @yalign with the appropriate values of
// @cell.
func (c cellRenderer) Alignment() (xalign float32, yalign float32) {
	var arg0 *C.GtkCellRenderer
	var arg1 *C.gfloat // out
	var arg2 *C.gfloat // out

	arg0 = (*C.GtkCellRenderer)(c.Native())

	C.gtk_cell_renderer_get_alignment(arg0, &arg1, &arg2)

	var ret0 float32
	var ret1 float32

	ret0 = float32(arg1)

	ret1 = float32(arg2)

	return ret0, ret1
}

// FixedSize fills in @width and @height with the appropriate size of @cell.
func (c cellRenderer) FixedSize() (width int, height int) {
	var arg0 *C.GtkCellRenderer
	var arg1 *C.gint // out
	var arg2 *C.gint // out

	arg0 = (*C.GtkCellRenderer)(c.Native())

	C.gtk_cell_renderer_get_fixed_size(arg0, &arg1, &arg2)

	var ret0 int
	var ret1 int

	ret0 = int(arg1)

	ret1 = int(arg2)

	return ret0, ret1
}

// Padding fills in @xpad and @ypad with the appropriate values of @cell.
func (c cellRenderer) Padding() (xpad int, ypad int) {
	var arg0 *C.GtkCellRenderer
	var arg1 *C.gint // out
	var arg2 *C.gint // out

	arg0 = (*C.GtkCellRenderer)(c.Native())

	C.gtk_cell_renderer_get_padding(arg0, &arg1, &arg2)

	var ret0 int
	var ret1 int

	ret0 = int(arg1)

	ret1 = int(arg2)

	return ret0, ret1
}

// PreferredHeight retreives a renderer’s natural size when rendered to
// @widget.
func (c cellRenderer) PreferredHeight(widget Widget) (minimumSize int, naturalSize int) {
	var arg0 *C.GtkCellRenderer
	var arg1 *C.GtkWidget
	var arg2 *C.gint // out
	var arg3 *C.gint // out

	arg0 = (*C.GtkCellRenderer)(c.Native())
	arg1 = (*C.GtkWidget)(widget.Native())

	C.gtk_cell_renderer_get_preferred_height(arg0, arg1, &arg2, &arg3)

	var ret0 int
	var ret1 int

	ret0 = int(arg2)

	ret1 = int(arg3)

	return ret0, ret1
}

// PreferredHeightForWidth retreives a cell renderers’s minimum and natural
// height if it were rendered to @widget with the specified @width.
func (c cellRenderer) PreferredHeightForWidth(widget Widget, width int) (minimumHeight int, naturalHeight int) {
	var arg0 *C.GtkCellRenderer
	var arg1 *C.GtkWidget
	var arg2 C.gint
	var arg3 *C.gint // out
	var arg4 *C.gint // out

	arg0 = (*C.GtkCellRenderer)(c.Native())
	arg1 = (*C.GtkWidget)(widget.Native())
	arg2 = C.gint(width)

	C.gtk_cell_renderer_get_preferred_height_for_width(arg0, arg1, arg2, &arg3, &arg4)

	var ret0 int
	var ret1 int

	ret0 = int(arg3)

	ret1 = int(arg4)

	return ret0, ret1
}

// PreferredSize retrieves the minimum and natural size of a cell taking
// into account the widget’s preference for height-for-width management.
func (c cellRenderer) PreferredSize(widget Widget) (minimumSize Requisition, naturalSize Requisition) {
	var arg0 *C.GtkCellRenderer
	var arg1 *C.GtkWidget
	var arg2 *C.GtkRequisition // out
	var arg3 *C.GtkRequisition // out

	arg0 = (*C.GtkCellRenderer)(c.Native())
	arg1 = (*C.GtkWidget)(widget.Native())

	C.gtk_cell_renderer_get_preferred_size(arg0, arg1, &arg2, &arg3)

	var ret0 *Requisition
	var ret1 *Requisition

	{
		ret0 = WrapRequisition(unsafe.Pointer(arg2))
	}

	{
		ret1 = WrapRequisition(unsafe.Pointer(arg3))
	}

	return ret0, ret1
}

// PreferredWidth retreives a renderer’s natural size when rendered to
// @widget.
func (c cellRenderer) PreferredWidth(widget Widget) (minimumSize int, naturalSize int) {
	var arg0 *C.GtkCellRenderer
	var arg1 *C.GtkWidget
	var arg2 *C.gint // out
	var arg3 *C.gint // out

	arg0 = (*C.GtkCellRenderer)(c.Native())
	arg1 = (*C.GtkWidget)(widget.Native())

	C.gtk_cell_renderer_get_preferred_width(arg0, arg1, &arg2, &arg3)

	var ret0 int
	var ret1 int

	ret0 = int(arg2)

	ret1 = int(arg3)

	return ret0, ret1
}

// PreferredWidthForHeight retreives a cell renderers’s minimum and natural
// width if it were rendered to @widget with the specified @height.
func (c cellRenderer) PreferredWidthForHeight(widget Widget, height int) (minimumWidth int, naturalWidth int) {
	var arg0 *C.GtkCellRenderer
	var arg1 *C.GtkWidget
	var arg2 C.gint
	var arg3 *C.gint // out
	var arg4 *C.gint // out

	arg0 = (*C.GtkCellRenderer)(c.Native())
	arg1 = (*C.GtkWidget)(widget.Native())
	arg2 = C.gint(height)

	C.gtk_cell_renderer_get_preferred_width_for_height(arg0, arg1, arg2, &arg3, &arg4)

	var ret0 int
	var ret1 int

	ret0 = int(arg3)

	ret1 = int(arg4)

	return ret0, ret1
}

// RequestMode gets whether the cell renderer prefers a height-for-width
// layout or a width-for-height layout.
func (c cellRenderer) RequestMode() SizeRequestMode {
	var arg0 *C.GtkCellRenderer

	arg0 = (*C.GtkCellRenderer)(c.Native())

	ret := C.gtk_cell_renderer_get_request_mode(arg0)

	var ret0 SizeRequestMode

	ret0 = SizeRequestMode(ret)

	return ret0
}

// Sensitive returns the cell renderer’s sensitivity.
func (c cellRenderer) Sensitive() bool {
	var arg0 *C.GtkCellRenderer

	arg0 = (*C.GtkCellRenderer)(c.Native())

	ret := C.gtk_cell_renderer_get_sensitive(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Size obtains the width and height needed to render the cell. Used by view
// widgets to determine the appropriate size for the cell_area passed to
// gtk_cell_renderer_render(). If @cell_area is not nil, fills in the x and
// y offsets (if set) of the cell relative to this location.
//
// Please note that the values set in @width and @height, as well as those
// in @x_offset and @y_offset are inclusive of the xpad and ypad properties.
func (c cellRenderer) Size(widget Widget, cellArea *gdk.Rectangle) (xOffset int, yOffset int, width int, height int) {
	var arg0 *C.GtkCellRenderer
	var arg1 *C.GtkWidget
	var arg2 *C.GdkRectangle
	var arg3 *C.gint // out
	var arg4 *C.gint // out
	var arg5 *C.gint // out
	var arg6 *C.gint // out

	arg0 = (*C.GtkCellRenderer)(c.Native())
	arg1 = (*C.GtkWidget)(widget.Native())
	arg2 = (*C.GdkRectangle)(cellArea.Native())

	C.gtk_cell_renderer_get_size(arg0, arg1, arg2, &arg3, &arg4, &arg5, &arg6)

	var ret0 int
	var ret1 int
	var ret2 int
	var ret3 int

	ret0 = int(arg3)

	ret1 = int(arg4)

	ret2 = int(arg5)

	ret3 = int(arg6)

	return ret0, ret1, ret2, ret3
}

// State translates the cell renderer state to StateFlags, based on the cell
// renderer and widget sensitivity, and the given CellRendererState.
func (c cellRenderer) State(widget Widget, cellState CellRendererState) StateFlags {
	var arg0 *C.GtkCellRenderer
	var arg1 *C.GtkWidget
	var arg2 C.GtkCellRendererState

	arg0 = (*C.GtkCellRenderer)(c.Native())
	arg1 = (*C.GtkWidget)(widget.Native())
	arg2 = (C.GtkCellRendererState)(cellState)

	ret := C.gtk_cell_renderer_get_state(arg0, arg1, arg2)

	var ret0 StateFlags

	ret0 = StateFlags(ret)

	return ret0
}

// Visible returns the cell renderer’s visibility.
func (c cellRenderer) Visible() bool {
	var arg0 *C.GtkCellRenderer

	arg0 = (*C.GtkCellRenderer)(c.Native())

	ret := C.gtk_cell_renderer_get_visible(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// IsActivatable checks whether the cell renderer can do something when
// activated.
func (c cellRenderer) IsActivatable() bool {
	var arg0 *C.GtkCellRenderer

	arg0 = (*C.GtkCellRenderer)(c.Native())

	ret := C.gtk_cell_renderer_is_activatable(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Render invokes the virtual render function of the CellRenderer. The three
// passed-in rectangles are areas in @cr. Most renderers will draw within
// @cell_area; the xalign, yalign, xpad, and ypad fields of the CellRenderer
// should be honored with respect to @cell_area. @background_area includes
// the blank space around the cell, and also the area containing the tree
// expander; so the @background_area rectangles for all cells tile to cover
// the entire @window.
func (c cellRenderer) Render(cr *cairo.Context, widget Widget, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) {
	var arg0 *C.GtkCellRenderer
	var arg1 *C.cairo_t
	var arg2 *C.GtkWidget
	var arg3 *C.GdkRectangle
	var arg4 *C.GdkRectangle
	var arg5 C.GtkCellRendererState

	arg0 = (*C.GtkCellRenderer)(c.Native())
	arg1 = (*C.cairo_t)(cr.Native())
	arg2 = (*C.GtkWidget)(widget.Native())
	arg3 = (*C.GdkRectangle)(backgroundArea.Native())
	arg4 = (*C.GdkRectangle)(cellArea.Native())
	arg5 = (C.GtkCellRendererState)(flags)

	C.gtk_cell_renderer_render(arg0, arg1, arg2, arg3, arg4, arg5)
}

// SetAlignment sets the renderer’s alignment within its available space.
func (c cellRenderer) SetAlignment(xalign float32, yalign float32) {
	var arg0 *C.GtkCellRenderer
	var arg1 C.gfloat
	var arg2 C.gfloat

	arg0 = (*C.GtkCellRenderer)(c.Native())
	arg1 = C.gfloat(xalign)
	arg2 = C.gfloat(yalign)

	C.gtk_cell_renderer_set_alignment(arg0, arg1, arg2)
}

// SetFixedSize sets the renderer size to be explicit, independent of the
// properties set.
func (c cellRenderer) SetFixedSize(width int, height int) {
	var arg0 *C.GtkCellRenderer
	var arg1 C.gint
	var arg2 C.gint

	arg0 = (*C.GtkCellRenderer)(c.Native())
	arg1 = C.gint(width)
	arg2 = C.gint(height)

	C.gtk_cell_renderer_set_fixed_size(arg0, arg1, arg2)
}

// SetPadding sets the renderer’s padding.
func (c cellRenderer) SetPadding(xpad int, ypad int) {
	var arg0 *C.GtkCellRenderer
	var arg1 C.gint
	var arg2 C.gint

	arg0 = (*C.GtkCellRenderer)(c.Native())
	arg1 = C.gint(xpad)
	arg2 = C.gint(ypad)

	C.gtk_cell_renderer_set_padding(arg0, arg1, arg2)
}

// SetSensitive sets the cell renderer’s sensitivity.
func (c cellRenderer) SetSensitive(sensitive bool) {
	var arg0 *C.GtkCellRenderer
	var arg1 C.gboolean

	arg0 = (*C.GtkCellRenderer)(c.Native())
	if sensitive {
		arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_sensitive(arg0, arg1)
}

// SetVisible sets the cell renderer’s visibility.
func (c cellRenderer) SetVisible(visible bool) {
	var arg0 *C.GtkCellRenderer
	var arg1 C.gboolean

	arg0 = (*C.GtkCellRenderer)(c.Native())
	if visible {
		arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_visible(arg0, arg1)
}

// StartEditing starts editing the contents of this @cell, through a new
// CellEditable widget created by the CellRendererClass.start_editing
// virtual function.
func (c cellRenderer) StartEditing(event *gdk.Event, widget Widget, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) CellEditable {
	var arg0 *C.GtkCellRenderer
	var arg1 *C.GdkEvent
	var arg2 *C.GtkWidget
	var arg3 *C.gchar
	var arg4 *C.GdkRectangle
	var arg5 *C.GdkRectangle
	var arg6 C.GtkCellRendererState

	arg0 = (*C.GtkCellRenderer)(c.Native())
	arg2 = (*C.GtkWidget)(widget.Native())
	arg3 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = (*C.GdkRectangle)(backgroundArea.Native())
	arg5 = (*C.GdkRectangle)(cellArea.Native())
	arg6 = (C.GtkCellRendererState)(flags)

	ret := C.gtk_cell_renderer_start_editing(arg0, arg1, arg2, arg3, arg4, arg5, arg6)

	var ret0 CellEditable

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(CellEditable)

	return ret0
}

// StopEditing informs the cell renderer that the editing is stopped. If
// @canceled is true, the cell renderer will emit the
// CellRenderer::editing-canceled signal.
//
// This function should be called by cell renderer implementations in
// response to the CellEditable::editing-done signal of CellEditable.
func (c cellRenderer) StopEditing(canceled bool) {
	var arg0 *C.GtkCellRenderer
	var arg1 C.gboolean

	arg0 = (*C.GtkCellRenderer)(c.Native())
	if canceled {
		arg1 = C.TRUE
	}

	C.gtk_cell_renderer_stop_editing(arg0, arg1)
}
