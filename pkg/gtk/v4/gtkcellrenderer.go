// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_mode_get_type()), F: marshalCellRendererMode},
		{T: externglib.Type(C.gtk_cell_renderer_state_get_type()), F: marshalCellRendererState},
		{T: externglib.Type(C.gtk_cell_renderer_get_type()), F: marshalCellRendererer},
	})
}

// CellRendererMode identifies how the user can interact with a particular cell.
type CellRendererMode int

const (
	// Inert: cell is just for display and cannot be interacted with. Note that
	// this doesn’t mean that eg. the row being drawn can’t be selected -- just
	// that a particular element of it cannot be individually modified.
	CellRendererModeInert CellRendererMode = iota
	// Activatable: cell can be clicked.
	CellRendererModeActivatable
	// Editable: cell can be edited or otherwise modified.
	CellRendererModeEditable
)

func marshalCellRendererMode(p uintptr) (interface{}, error) {
	return CellRendererMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// CellRendererState tells how a cell is to be rendered.
type CellRendererState int

const (
	// CellRendererStateSelected: cell is currently selected, and probably has a
	// selection colored background to render to.
	CellRendererStateSelected CellRendererState = 0b1
	// CellRendererStatePrelit: mouse is hovering over the cell.
	CellRendererStatePrelit CellRendererState = 0b10
	// CellRendererStateInsensitive: cell is drawn in an insensitive manner
	CellRendererStateInsensitive CellRendererState = 0b100
	// CellRendererStateSorted: cell is in a sorted row
	CellRendererStateSorted CellRendererState = 0b1000
	// CellRendererStateFocused: cell is in the focus row.
	CellRendererStateFocused CellRendererState = 0b10000
	// CellRendererStateExpandable: cell is in a row that can be expanded
	CellRendererStateExpandable CellRendererState = 0b100000
	// CellRendererStateExpanded: cell is in a row that is expanded
	CellRendererStateExpanded CellRendererState = 0b1000000
)

func marshalCellRendererState(p uintptr) (interface{}, error) {
	return CellRendererState(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// CellRendererOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CellRendererOverrider interface {
	EditingCanceled()

	EditingStarted(editable CellEditabler, path string)
	// PreferredHeight retrieves a renderer’s natural size when rendered to
	// @widget.
	PreferredHeight(widget Widgetter) (minimumSize int, naturalSize int)
	// PreferredHeightForWidth retrieves a cell renderers’s minimum and natural
	// height if it were rendered to @widget with the specified @width.
	PreferredHeightForWidth(widget Widgetter, width int) (minimumHeight int, naturalHeight int)
	// PreferredWidth retrieves a renderer’s natural size when rendered to
	// @widget.
	PreferredWidth(widget Widgetter) (minimumSize int, naturalSize int)
	// PreferredWidthForHeight retrieves a cell renderers’s minimum and natural
	// width if it were rendered to @widget with the specified @height.
	PreferredWidthForHeight(widget Widgetter, height int) (minimumWidth int, naturalWidth int)
	// RequestMode gets whether the cell renderer prefers a height-for-width
	// layout or a width-for-height layout.
	RequestMode() SizeRequestMode
}

// CellRendererer describes CellRenderer's methods.
type CellRendererer interface {
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
	PreferredHeight(widget Widgetter) (minimumSize int, naturalSize int)
	// PreferredHeightForWidth retrieves a cell renderers’s minimum and natural
	// height if it were rendered to @widget with the specified @width.
	PreferredHeightForWidth(widget Widgetter, width int) (minimumHeight int, naturalHeight int)
	// PreferredSize retrieves the minimum and natural size of a cell taking
	// into account the widget’s preference for height-for-width management.
	PreferredSize(widget Widgetter) (minimumSize Requisition, naturalSize Requisition)
	// PreferredWidth retrieves a renderer’s natural size when rendered to
	// @widget.
	PreferredWidth(widget Widgetter) (minimumSize int, naturalSize int)
	// PreferredWidthForHeight retrieves a cell renderers’s minimum and natural
	// width if it were rendered to @widget with the specified @height.
	PreferredWidthForHeight(widget Widgetter, height int) (minimumWidth int, naturalWidth int)
	// RequestMode gets whether the cell renderer prefers a height-for-width
	// layout or a width-for-height layout.
	RequestMode() SizeRequestMode
	// Sensitive returns the cell renderer’s sensitivity.
	Sensitive() bool
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
	// StopEditing informs the cell renderer that the editing is stopped.
	StopEditing(canceled bool)
}

// CellRenderer: object for rendering a single cell
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
type CellRenderer struct {
	externglib.InitiallyUnowned
}

var (
	_ CellRendererer  = (*CellRenderer)(nil)
	_ gextras.Nativer = (*CellRenderer)(nil)
)

func wrapCellRenderer(obj *externglib.Object) CellRendererer {
	return &CellRenderer{
		InitiallyUnowned: externglib.InitiallyUnowned{
			Object: obj,
		},
	}
}

func marshalCellRendererer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCellRenderer(obj), nil
}

// Alignment fills in @xalign and @yalign with the appropriate values of @cell.
func (cell *CellRenderer) Alignment() (xalign float32, yalign float32) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.float            // in
	var _arg2 C.float            // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_renderer_get_alignment(_arg0, &_arg1, &_arg2)

	var _xalign float32 // out
	var _yalign float32 // out

	_xalign = float32(_arg1)
	_yalign = float32(_arg2)

	return _xalign, _yalign
}

// FixedSize fills in @width and @height with the appropriate size of @cell.
func (cell *CellRenderer) FixedSize() (width int, height int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.int              // in
	var _arg2 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_renderer_get_fixed_size(_arg0, &_arg1, &_arg2)

	var _width int  // out
	var _height int // out

	_width = int(_arg1)
	_height = int(_arg2)

	return _width, _height
}

// IsExpanded checks whether the given CellRenderer is expanded.
func (cell *CellRenderer) IsExpanded() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_get_is_expanded(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsExpander checks whether the given CellRenderer is an expander.
func (cell *CellRenderer) IsExpander() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_get_is_expander(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Padding fills in @xpad and @ypad with the appropriate values of @cell.
func (cell *CellRenderer) Padding() (xpad int, ypad int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.int              // in
	var _arg2 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_renderer_get_padding(_arg0, &_arg1, &_arg2)

	var _xpad int // out
	var _ypad int // out

	_xpad = int(_arg1)
	_ypad = int(_arg2)

	return _xpad, _ypad
}

// PreferredHeight retrieves a renderer’s natural size when rendered to @widget.
func (cell *CellRenderer) PreferredHeight(widget Widgetter) (minimumSize int, naturalSize int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.int              // in
	var _arg3 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))

	C.gtk_cell_renderer_get_preferred_height(_arg0, _arg1, &_arg2, &_arg3)

	var _minimumSize int // out
	var _naturalSize int // out

	_minimumSize = int(_arg2)
	_naturalSize = int(_arg3)

	return _minimumSize, _naturalSize
}

// PreferredHeightForWidth retrieves a cell renderers’s minimum and natural
// height if it were rendered to @widget with the specified @width.
func (cell *CellRenderer) PreferredHeightForWidth(widget Widgetter, width int) (minimumHeight int, naturalHeight int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.int              // out
	var _arg3 C.int              // in
	var _arg4 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))
	_arg2 = C.int(width)

	C.gtk_cell_renderer_get_preferred_height_for_width(_arg0, _arg1, _arg2, &_arg3, &_arg4)

	var _minimumHeight int // out
	var _naturalHeight int // out

	_minimumHeight = int(_arg3)
	_naturalHeight = int(_arg4)

	return _minimumHeight, _naturalHeight
}

// PreferredSize retrieves the minimum and natural size of a cell taking into
// account the widget’s preference for height-for-width management.
func (cell *CellRenderer) PreferredSize(widget Widgetter) (minimumSize Requisition, naturalSize Requisition) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.GtkRequisition   // in
	var _arg3 C.GtkRequisition   // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))

	C.gtk_cell_renderer_get_preferred_size(_arg0, _arg1, &_arg2, &_arg3)

	var _minimumSize Requisition // out
	var _naturalSize Requisition // out

	_minimumSize = *(*Requisition)(unsafe.Pointer((&_arg2)))
	_naturalSize = *(*Requisition)(unsafe.Pointer((&_arg3)))

	return _minimumSize, _naturalSize
}

// PreferredWidth retrieves a renderer’s natural size when rendered to @widget.
func (cell *CellRenderer) PreferredWidth(widget Widgetter) (minimumSize int, naturalSize int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.int              // in
	var _arg3 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))

	C.gtk_cell_renderer_get_preferred_width(_arg0, _arg1, &_arg2, &_arg3)

	var _minimumSize int // out
	var _naturalSize int // out

	_minimumSize = int(_arg2)
	_naturalSize = int(_arg3)

	return _minimumSize, _naturalSize
}

// PreferredWidthForHeight retrieves a cell renderers’s minimum and natural
// width if it were rendered to @widget with the specified @height.
func (cell *CellRenderer) PreferredWidthForHeight(widget Widgetter, height int) (minimumWidth int, naturalWidth int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.int              // out
	var _arg3 C.int              // in
	var _arg4 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))
	_arg2 = C.int(height)

	C.gtk_cell_renderer_get_preferred_width_for_height(_arg0, _arg1, _arg2, &_arg3, &_arg4)

	var _minimumWidth int // out
	var _naturalWidth int // out

	_minimumWidth = int(_arg3)
	_naturalWidth = int(_arg4)

	return _minimumWidth, _naturalWidth
}

// RequestMode gets whether the cell renderer prefers a height-for-width layout
// or a width-for-height layout.
func (cell *CellRenderer) RequestMode() SizeRequestMode {
	var _arg0 *C.GtkCellRenderer   // out
	var _cret C.GtkSizeRequestMode // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_get_request_mode(_arg0)

	var _sizeRequestMode SizeRequestMode // out

	_sizeRequestMode = (SizeRequestMode)(_cret)

	return _sizeRequestMode
}

// Sensitive returns the cell renderer’s sensitivity.
func (cell *CellRenderer) Sensitive() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_get_sensitive(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Visible returns the cell renderer’s visibility.
func (cell *CellRenderer) Visible() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsActivatable checks whether the cell renderer can do something when
// activated.
func (cell *CellRenderer) IsActivatable() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_is_activatable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAlignment sets the renderer’s alignment within its available space.
func (cell *CellRenderer) SetAlignment(xalign float32, yalign float32) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = C.float(xalign)
	_arg2 = C.float(yalign)

	C.gtk_cell_renderer_set_alignment(_arg0, _arg1, _arg2)
}

// SetFixedSize sets the renderer size to be explicit, independent of the
// properties set.
func (cell *CellRenderer) SetFixedSize(width int, height int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.int              // out
	var _arg2 C.int              // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	C.gtk_cell_renderer_set_fixed_size(_arg0, _arg1, _arg2)
}

// SetIsExpanded sets whether the given CellRenderer is expanded.
func (cell *CellRenderer) SetIsExpanded(isExpanded bool) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if isExpanded {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_is_expanded(_arg0, _arg1)
}

// SetIsExpander sets whether the given CellRenderer is an expander.
func (cell *CellRenderer) SetIsExpander(isExpander bool) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if isExpander {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_is_expander(_arg0, _arg1)
}

// SetPadding sets the renderer’s padding.
func (cell *CellRenderer) SetPadding(xpad int, ypad int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.int              // out
	var _arg2 C.int              // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = C.int(xpad)
	_arg2 = C.int(ypad)

	C.gtk_cell_renderer_set_padding(_arg0, _arg1, _arg2)
}

// SetSensitive sets the cell renderer’s sensitivity.
func (cell *CellRenderer) SetSensitive(sensitive bool) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if sensitive {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_sensitive(_arg0, _arg1)
}

// SetVisible sets the cell renderer’s visibility.
func (cell *CellRenderer) SetVisible(visible bool) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_visible(_arg0, _arg1)
}

// StopEditing informs the cell renderer that the editing is stopped. If
// @canceled is true, the cell renderer will emit the
// CellRenderer::editing-canceled signal.
//
// This function should be called by cell renderer implementations in response
// to the CellEditable::editing-done signal of CellEditable.
func (cell *CellRenderer) StopEditing(canceled bool) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if canceled {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_stop_editing(_arg0, _arg1)
}
