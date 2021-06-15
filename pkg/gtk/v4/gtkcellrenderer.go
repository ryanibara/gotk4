// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_mode_get_type()), F: marshalCellRendererMode},
		{T: externglib.Type(C.gtk_cell_renderer_state_get_type()), F: marshalCellRendererState},
		{T: externglib.Type(C.gtk_cell_renderer_get_type()), F: marshalCellRenderer},
	})
}

// CellRendererMode identifies how the user can interact with a particular cell.
type CellRendererMode int

const (
	// CellRendererModeInert: the cell is just for display and cannot be
	// interacted with. Note that this doesn’t mean that eg. the row being drawn
	// can’t be selected -- just that a particular element of it cannot be
	// individually modified.
	CellRendererModeInert CellRendererMode = 0
	// CellRendererModeActivatable: the cell can be clicked.
	CellRendererModeActivatable CellRendererMode = 1
	// CellRendererModeEditable: the cell can be edited or otherwise modified.
	CellRendererModeEditable CellRendererMode = 2
)

func marshalCellRendererMode(p uintptr) (interface{}, error) {
	return CellRendererMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// CellRendererState tells how a cell is to be rendered.
type CellRendererState int

const (
	// CellRendererStateSelected: the cell is currently selected, and probably
	// has a selection colored background to render to.
	CellRendererStateSelected CellRendererState = 1
	// CellRendererStatePrelit: the mouse is hovering over the cell.
	CellRendererStatePrelit CellRendererState = 2
	// CellRendererStateInsensitive: the cell is drawn in an insensitive manner
	CellRendererStateInsensitive CellRendererState = 4
	// CellRendererStateSorted: the cell is in a sorted row
	CellRendererStateSorted CellRendererState = 8
	// CellRendererStateFocused: the cell is in the focus row.
	CellRendererStateFocused CellRendererState = 16
	// CellRendererStateExpandable: the cell is in a row that can be expanded
	CellRendererStateExpandable CellRendererState = 32
	// CellRendererStateExpanded: the cell is in a row that is expanded
	CellRendererStateExpanded CellRendererState = 64
)

func marshalCellRendererState(p uintptr) (interface{}, error) {
	return CellRendererState(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// CellRenderer: an object for rendering a single cell
//
// The CellRenderer is a base class of a set of objects used for rendering a
// cell to a #cairo_t. These objects are used primarily by the TreeView widget,
// though they aren’t tied to them in any specific way. It is worth noting that
// CellRenderer is not a Widget and cannot be treated as such.
//
// The primary use of a CellRenderer is for drawing a certain graphical elements
// on a #cairo_t. Typically, one cell renderer is used to draw many cells on the
// screen. To this extent, it isn’t expected that a CellRenderer keep any
// permanent state around. Instead, any state is set just prior to use using
// #GObjects property system. Then, the cell is measured using
// gtk_cell_renderer_get_preferred_size(). Finally, the cell is rendered in the
// correct location using gtk_cell_renderer_snapshot().
//
// There are a number of rules that must be followed when writing a new
// CellRenderer. First and foremost, it’s important that a certain set of
// properties will always yield a cell renderer of the same size, barring a
// style change. The CellRenderer also has a number of generic properties that
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

	// Activate passes an activate event to the cell renderer for possible
	// processing. Some cell renderers may use events; for example,
	// CellRendererToggle toggles when it gets a mouse click.
	Activate(event gdk.Event, widget Widget, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) bool
	// AlignedArea gets the aligned area used by @cell inside @cell_area. Used
	// for finding the appropriate edit and focus rectangle.
	AlignedArea(widget Widget, flags CellRendererState, cellArea *gdk.Rectangle) gdk.Rectangle
	// Alignment fills in @xalign and @yalign with the appropriate values of
	// @cell.
	Alignment() (xalign float32, yalign float32)
	// FixedSize fills in @width and @height with the appropriate size of @cell.
	FixedSize() (width int, height int)
	// IsExpanded checks whether the given CellRenderer is expanded.
	IsExpanded() bool
	// IsExpander checks whether the given CellRenderer is an expander.
	IsExpander() bool
	// Padding fills in @xpad and @ypad with the appropriate values of @cell.
	Padding() (xpad int, ypad int)
	// PreferredHeight retrieves a renderer’s natural size when rendered to
	// @widget.
	PreferredHeight(widget Widget) (minimumSize int, naturalSize int)
	// PreferredHeightForWidth retrieves a cell renderers’s minimum and natural
	// height if it were rendered to @widget with the specified @width.
	PreferredHeightForWidth(widget Widget, width int) (minimumHeight int, naturalHeight int)
	// PreferredSize retrieves the minimum and natural size of a cell taking
	// into account the widget’s preference for height-for-width management.
	PreferredSize(widget Widget) (minimumSize Requisition, naturalSize Requisition)
	// PreferredWidth retrieves a renderer’s natural size when rendered to
	// @widget.
	PreferredWidth(widget Widget) (minimumSize int, naturalSize int)
	// PreferredWidthForHeight retrieves a cell renderers’s minimum and natural
	// width if it were rendered to @widget with the specified @height.
	PreferredWidthForHeight(widget Widget, height int) (minimumWidth int, naturalWidth int)
	// RequestMode gets whether the cell renderer prefers a height-for-width
	// layout or a width-for-height layout.
	RequestMode() SizeRequestMode
	// Sensitive returns the cell renderer’s sensitivity.
	Sensitive() bool
	// State translates the cell renderer state to StateFlags, based on the cell
	// renderer and widget sensitivity, and the given CellRendererState.
	State(widget Widget, cellState CellRendererState) StateFlags
	// Visible returns the cell renderer’s visibility.
	Visible() bool
	// IsActivatable checks whether the cell renderer can do something when
	// activated.
	IsActivatable() bool
	// SetAlignment sets the renderer’s alignment within its available space.
	SetAlignment(xalign float32, yalign float32)
	// SetFixedSize sets the renderer size to be explicit, independent of the
	// properties set.
	SetFixedSize(width int, height int)
	// SetIsExpanded sets whether the given CellRenderer is expanded.
	SetIsExpanded(isExpanded bool)
	// SetIsExpander sets whether the given CellRenderer is an expander.
	SetIsExpander(isExpander bool)
	// SetPadding sets the renderer’s padding.
	SetPadding(xpad int, ypad int)
	// SetSensitive sets the cell renderer’s sensitivity.
	SetSensitive(sensitive bool)
	// SetVisible sets the cell renderer’s visibility.
	SetVisible(visible bool)
	// Snapshot invokes the virtual render function of the CellRenderer. The
	// three passed-in rectangles are areas in @cr. Most renderers will draw
	// within @cell_area; the xalign, yalign, xpad, and ypad fields of the
	// CellRenderer should be honored with respect to @cell_area.
	// @background_area includes the blank space around the cell, and also the
	// area containing the tree expander; so the @background_area rectangles for
	// all cells tile to cover the entire @window.
	Snapshot(snapshot Snapshot, widget Widget, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState)
	// StartEditing starts editing the contents of this @cell, through a new
	// CellEditable widget created by the CellRendererClass.start_editing
	// virtual function.
	StartEditing(event gdk.Event, widget Widget, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) CellEditable
	// StopEditing informs the cell renderer that the editing is stopped. If
	// @canceled is true, the cell renderer will emit the
	// CellRenderer::editing-canceled signal.
	//
	// This function should be called by cell renderer implementations in
	// response to the CellEditable::editing-done signal of CellEditable.
	StopEditing(canceled bool)
}

// cellRenderer implements the CellRenderer class.
type cellRenderer struct {
	gextras.Objector
}

var _ CellRenderer = (*cellRenderer)(nil)

// WrapCellRenderer wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellRenderer(obj *externglib.Object) CellRenderer {
	return cellRenderer{
		Objector: obj,
	}
}

func marshalCellRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellRenderer(obj), nil
}

// Activate passes an activate event to the cell renderer for possible
// processing. Some cell renderers may use events; for example,
// CellRendererToggle toggles when it gets a mouse click.
func (c cellRenderer) Activate(event gdk.Event, widget Widget, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) bool {
	var _arg0 *C.GtkCellRenderer     // out
	var _arg1 *C.GdkEvent            // out
	var _arg2 *C.GtkWidget           // out
	var _arg3 *C.char                // out
	var _arg4 *C.GdkRectangle        // out
	var _arg5 *C.GdkRectangle        // out
	var _arg6 C.GtkCellRendererState // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.GdkRectangle)(unsafe.Pointer(backgroundArea.Native()))
	_arg5 = (*C.GdkRectangle)(unsafe.Pointer(cellArea.Native()))
	_arg6 = (C.GtkCellRendererState)(flags)

	_cret = C.gtk_cell_renderer_activate(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AlignedArea gets the aligned area used by @cell inside @cell_area. Used
// for finding the appropriate edit and focus rectangle.
func (c cellRenderer) AlignedArea(widget Widget, flags CellRendererState, cellArea *gdk.Rectangle) gdk.Rectangle {
	var _arg0 *C.GtkCellRenderer     // out
	var _arg1 *C.GtkWidget           // out
	var _arg2 C.GtkCellRendererState // out
	var _arg3 *C.GdkRectangle        // out
	var _alignedArea gdk.Rectangle

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = (C.GtkCellRendererState)(flags)
	_arg3 = (*C.GdkRectangle)(unsafe.Pointer(cellArea.Native()))

	C.gtk_cell_renderer_get_aligned_area(_arg0, _arg1, _arg2, _arg3, (*C.GdkRectangle)(unsafe.Pointer(&_alignedArea)))

	return _alignedArea
}

// Alignment fills in @xalign and @yalign with the appropriate values of
// @cell.
func (c cellRenderer) Alignment() (xalign float32, yalign float32) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.float            // in
	var _arg2 C.float            // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))

	C.gtk_cell_renderer_get_alignment(_arg0, &_arg1, &_arg2)

	var _xalign float32 // out
	var _yalign float32 // out

	_xalign = (float32)(_arg1)
	_yalign = (float32)(_arg2)

	return _xalign, _yalign
}

// FixedSize fills in @width and @height with the appropriate size of @cell.
func (c cellRenderer) FixedSize() (width int, height int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.int              // in
	var _arg2 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))

	C.gtk_cell_renderer_get_fixed_size(_arg0, &_arg1, &_arg2)

	var _width int  // out
	var _height int // out

	_width = (int)(_arg1)
	_height = (int)(_arg2)

	return _width, _height
}

// IsExpanded checks whether the given CellRenderer is expanded.
func (c cellRenderer) IsExpanded() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_cell_renderer_get_is_expanded(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsExpander checks whether the given CellRenderer is an expander.
func (c cellRenderer) IsExpander() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_cell_renderer_get_is_expander(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Padding fills in @xpad and @ypad with the appropriate values of @cell.
func (c cellRenderer) Padding() (xpad int, ypad int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.int              // in
	var _arg2 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))

	C.gtk_cell_renderer_get_padding(_arg0, &_arg1, &_arg2)

	var _xpad int // out
	var _ypad int // out

	_xpad = (int)(_arg1)
	_ypad = (int)(_arg2)

	return _xpad, _ypad
}

// PreferredHeight retrieves a renderer’s natural size when rendered to
// @widget.
func (c cellRenderer) PreferredHeight(widget Widget) (minimumSize int, naturalSize int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.int              // in
	var _arg3 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_cell_renderer_get_preferred_height(_arg0, _arg1, &_arg2, &_arg3)

	var _minimumSize int // out
	var _naturalSize int // out

	_minimumSize = (int)(_arg2)
	_naturalSize = (int)(_arg3)

	return _minimumSize, _naturalSize
}

// PreferredHeightForWidth retrieves a cell renderers’s minimum and natural
// height if it were rendered to @widget with the specified @width.
func (c cellRenderer) PreferredHeightForWidth(widget Widget, width int) (minimumHeight int, naturalHeight int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.int              // out
	var _arg3 C.int              // in
	var _arg4 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = (C.int)(width)

	C.gtk_cell_renderer_get_preferred_height_for_width(_arg0, _arg1, _arg2, &_arg3, &_arg4)

	var _minimumHeight int // out
	var _naturalHeight int // out

	_minimumHeight = (int)(_arg3)
	_naturalHeight = (int)(_arg4)

	return _minimumHeight, _naturalHeight
}

// PreferredSize retrieves the minimum and natural size of a cell taking
// into account the widget’s preference for height-for-width management.
func (c cellRenderer) PreferredSize(widget Widget) (minimumSize Requisition, naturalSize Requisition) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _minimumSize Requisition
	var _naturalSize Requisition

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_cell_renderer_get_preferred_size(_arg0, _arg1, (*C.GtkRequisition)(unsafe.Pointer(&_minimumSize)), (*C.GtkRequisition)(unsafe.Pointer(&_naturalSize)))

	return _minimumSize, _naturalSize
}

// PreferredWidth retrieves a renderer’s natural size when rendered to
// @widget.
func (c cellRenderer) PreferredWidth(widget Widget) (minimumSize int, naturalSize int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.int              // in
	var _arg3 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_cell_renderer_get_preferred_width(_arg0, _arg1, &_arg2, &_arg3)

	var _minimumSize int // out
	var _naturalSize int // out

	_minimumSize = (int)(_arg2)
	_naturalSize = (int)(_arg3)

	return _minimumSize, _naturalSize
}

// PreferredWidthForHeight retrieves a cell renderers’s minimum and natural
// width if it were rendered to @widget with the specified @height.
func (c cellRenderer) PreferredWidthForHeight(widget Widget, height int) (minimumWidth int, naturalWidth int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.int              // out
	var _arg3 C.int              // in
	var _arg4 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = (C.int)(height)

	C.gtk_cell_renderer_get_preferred_width_for_height(_arg0, _arg1, _arg2, &_arg3, &_arg4)

	var _minimumWidth int // out
	var _naturalWidth int // out

	_minimumWidth = (int)(_arg3)
	_naturalWidth = (int)(_arg4)

	return _minimumWidth, _naturalWidth
}

// RequestMode gets whether the cell renderer prefers a height-for-width
// layout or a width-for-height layout.
func (c cellRenderer) RequestMode() SizeRequestMode {
	var _arg0 *C.GtkCellRenderer   // out
	var _cret C.GtkSizeRequestMode // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_cell_renderer_get_request_mode(_arg0)

	var _sizeRequestMode SizeRequestMode // out

	_sizeRequestMode = SizeRequestMode(_cret)

	return _sizeRequestMode
}

// Sensitive returns the cell renderer’s sensitivity.
func (c cellRenderer) Sensitive() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_cell_renderer_get_sensitive(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// State translates the cell renderer state to StateFlags, based on the cell
// renderer and widget sensitivity, and the given CellRendererState.
func (c cellRenderer) State(widget Widget, cellState CellRendererState) StateFlags {
	var _arg0 *C.GtkCellRenderer     // out
	var _arg1 *C.GtkWidget           // out
	var _arg2 C.GtkCellRendererState // out
	var _cret C.GtkStateFlags        // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = (C.GtkCellRendererState)(cellState)

	_cret = C.gtk_cell_renderer_get_state(_arg0, _arg1, _arg2)

	var _stateFlags StateFlags // out

	_stateFlags = StateFlags(_cret)

	return _stateFlags
}

// Visible returns the cell renderer’s visibility.
func (c cellRenderer) Visible() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_cell_renderer_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsActivatable checks whether the cell renderer can do something when
// activated.
func (c cellRenderer) IsActivatable() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_cell_renderer_is_activatable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAlignment sets the renderer’s alignment within its available space.
func (c cellRenderer) SetAlignment(xalign float32, yalign float32) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	_arg1 = (C.float)(xalign)
	_arg2 = (C.float)(yalign)

	C.gtk_cell_renderer_set_alignment(_arg0, _arg1, _arg2)
}

// SetFixedSize sets the renderer size to be explicit, independent of the
// properties set.
func (c cellRenderer) SetFixedSize(width int, height int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.int              // out
	var _arg2 C.int              // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	_arg1 = (C.int)(width)
	_arg2 = (C.int)(height)

	C.gtk_cell_renderer_set_fixed_size(_arg0, _arg1, _arg2)
}

// SetIsExpanded sets whether the given CellRenderer is expanded.
func (c cellRenderer) SetIsExpanded(isExpanded bool) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	if isExpanded {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_is_expanded(_arg0, _arg1)
}

// SetIsExpander sets whether the given CellRenderer is an expander.
func (c cellRenderer) SetIsExpander(isExpander bool) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	if isExpander {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_is_expander(_arg0, _arg1)
}

// SetPadding sets the renderer’s padding.
func (c cellRenderer) SetPadding(xpad int, ypad int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.int              // out
	var _arg2 C.int              // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	_arg1 = (C.int)(xpad)
	_arg2 = (C.int)(ypad)

	C.gtk_cell_renderer_set_padding(_arg0, _arg1, _arg2)
}

// SetSensitive sets the cell renderer’s sensitivity.
func (c cellRenderer) SetSensitive(sensitive bool) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	if sensitive {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_sensitive(_arg0, _arg1)
}

// SetVisible sets the cell renderer’s visibility.
func (c cellRenderer) SetVisible(visible bool) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_visible(_arg0, _arg1)
}

// Snapshot invokes the virtual render function of the CellRenderer. The
// three passed-in rectangles are areas in @cr. Most renderers will draw
// within @cell_area; the xalign, yalign, xpad, and ypad fields of the
// CellRenderer should be honored with respect to @cell_area.
// @background_area includes the blank space around the cell, and also the
// area containing the tree expander; so the @background_area rectangles for
// all cells tile to cover the entire @window.
func (c cellRenderer) Snapshot(snapshot Snapshot, widget Widget, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) {
	var _arg0 *C.GtkCellRenderer     // out
	var _arg1 *C.GtkSnapshot         // out
	var _arg2 *C.GtkWidget           // out
	var _arg3 *C.GdkRectangle        // out
	var _arg4 *C.GdkRectangle        // out
	var _arg5 C.GtkCellRendererState // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.GdkRectangle)(unsafe.Pointer(backgroundArea.Native()))
	_arg4 = (*C.GdkRectangle)(unsafe.Pointer(cellArea.Native()))
	_arg5 = (C.GtkCellRendererState)(flags)

	C.gtk_cell_renderer_snapshot(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// StartEditing starts editing the contents of this @cell, through a new
// CellEditable widget created by the CellRendererClass.start_editing
// virtual function.
func (c cellRenderer) StartEditing(event gdk.Event, widget Widget, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) CellEditable {
	var _arg0 *C.GtkCellRenderer     // out
	var _arg1 *C.GdkEvent            // out
	var _arg2 *C.GtkWidget           // out
	var _arg3 *C.char                // out
	var _arg4 *C.GdkRectangle        // out
	var _arg5 *C.GdkRectangle        // out
	var _arg6 C.GtkCellRendererState // out
	var _cret *C.GtkCellEditable     // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.GdkRectangle)(unsafe.Pointer(backgroundArea.Native()))
	_arg5 = (*C.GdkRectangle)(unsafe.Pointer(cellArea.Native()))
	_arg6 = (C.GtkCellRendererState)(flags)

	_cret = C.gtk_cell_renderer_start_editing(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)

	var _cellEditable CellEditable // out

	_cellEditable = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(CellEditable)

	return _cellEditable
}

// StopEditing informs the cell renderer that the editing is stopped. If
// @canceled is true, the cell renderer will emit the
// CellRenderer::editing-canceled signal.
//
// This function should be called by cell renderer implementations in
// response to the CellEditable::editing-done signal of CellEditable.
func (c cellRenderer) StopEditing(canceled bool) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(c.Native()))
	if canceled {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_stop_editing(_arg0, _arg1)
}
