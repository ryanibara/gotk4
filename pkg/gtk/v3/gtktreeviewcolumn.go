// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_tree_view_column_sizing_get_type()), F: marshalTreeViewColumnSizing},
		{T: externglib.Type(C.gtk_tree_view_column_get_type()), F: marshalTreeViewColumn},
	})
}

// TreeViewColumnSizing: the sizing method the column uses to determine its
// width. Please note that @GTK_TREE_VIEW_COLUMN_AUTOSIZE are inefficient for
// large views, and can make columns appear choppy.
type TreeViewColumnSizing int

const (
	// GrowOnly columns only get bigger in reaction to changes in the model
	TreeViewColumnSizingGrowOnly TreeViewColumnSizing = 0
	// autosize columns resize to be the optimal size everytime the model
	// changes.
	TreeViewColumnSizingAutosize TreeViewColumnSizing = 1
	// fixed columns are a fixed numbers of pixels wide.
	TreeViewColumnSizingFixed TreeViewColumnSizing = 2
)

func marshalTreeViewColumnSizing(p uintptr) (interface{}, error) {
	return TreeViewColumnSizing(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// TreeCellDataFunc: function to set the properties of a cell instead of just
// using the straight mapping between the cell and the model. This is useful for
// customizing the cell renderer. For example, a function might get an integer
// from the @tree_model, and render it to the “text” attribute of “cell” by
// converting it to its written equivalent. This is set by calling
// gtk_tree_view_column_set_cell_data_func()
type TreeCellDataFunc func(treeColumn TreeViewColumn, cell CellRenderer, treeModel TreeModel, iter TreeIter)

//export gotk4_TreeCellDataFunc
func gotk4_TreeCellDataFunc(arg0 *C.GtkTreeViewColumn, arg1 *C.GtkCellRenderer, arg2 *C.GtkTreeModel, arg3 *C.GtkTreeIter, arg4 C.gpointer) {
	v := box.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	var treeColumn TreeViewColumn // out
	var cell CellRenderer         // out
	var treeModel TreeModel       // out
	var iter TreeIter             // out

	treeColumn = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(TreeViewColumn)
	cell = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1))).(CellRenderer)
	treeModel = gextras.CastObject(externglib.Take(unsafe.Pointer(arg2))).(TreeModel)
	iter = (TreeIter)(unsafe.Pointer(arg3))

	fn := v.(TreeCellDataFunc)
	fn(treeColumn, cell, treeModel, iter)
}

// TreeViewColumn: the GtkTreeViewColumn object represents a visible column in a
// TreeView widget. It allows to set properties of the column header, and
// functions as a holding pen for the cell renderers which determine how the
// data in the column is displayed.
//
// Please refer to the [tree widget conceptual overview][TreeWidget] for an
// overview of all the objects and data types related to the tree widget and how
// they work together.
type TreeViewColumn interface {
	gextras.Objector

	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsCellLayout casts the class to the CellLayout interface.
	AsCellLayout() CellLayout

	// AddAttributeTreeViewColumn adds an attribute mapping to the list in
	// @tree_column. The @column is the column of the model to get a value from,
	// and the @attribute is the parameter on @cell_renderer to be set from the
	// value. So for example if column 2 of the model contains strings, you
	// could have the “text” attribute of a CellRendererText get its values from
	// column 2.
	AddAttributeTreeViewColumn(cellRenderer CellRenderer, attribute string, column int)
	// CellGetPositionTreeViewColumn obtains the horizontal position and size of
	// a cell in a column. If the cell is not found in the column, @start_pos
	// and @width are not changed and false is returned.
	CellGetPositionTreeViewColumn(cellRenderer CellRenderer) (xOffset int, width int, ok bool)
	// CellGetSizeTreeViewColumn obtains the width and height needed to render
	// the column. This is used primarily by the TreeView.
	CellGetSizeTreeViewColumn(cellArea gdk.Rectangle) (xOffset int, yOffset int, width int, height int)
	// CellIsVisibleTreeViewColumn returns true if any of the cells packed into
	// the @tree_column are visible. For this to be meaningful, you must first
	// initialize the cells with gtk_tree_view_column_cell_set_cell_data()
	CellIsVisibleTreeViewColumn() bool
	// CellSetCellDataTreeViewColumn sets the cell renderer based on the
	// @tree_model and @iter. That is, for every attribute mapping in
	// @tree_column, it will get a value from the set column on the @iter, and
	// use that value to set the attribute on the cell renderer. This is used
	// primarily by the TreeView.
	CellSetCellDataTreeViewColumn(treeModel TreeModel, iter TreeIter, isExpander bool, isExpanded bool)
	// ClearTreeViewColumn unsets all the mappings on all renderers on the
	// @tree_column.
	ClearTreeViewColumn()
	// ClearAttributesTreeViewColumn clears all existing attributes previously
	// set with gtk_tree_view_column_set_attributes().
	ClearAttributesTreeViewColumn(cellRenderer CellRenderer)
	// ClickedTreeViewColumn emits the “clicked” signal on the column. This
	// function will only work if @tree_column is clickable.
	ClickedTreeViewColumn()
	// FocusCellTreeViewColumn sets the current keyboard focus to be at @cell,
	// if the column contains 2 or more editable and activatable cells.
	FocusCellTreeViewColumn(cell CellRenderer)
	// Alignment returns the current x alignment of @tree_column. This value can
	// range between 0.0 and 1.0.
	Alignment() float32
	// Button returns the button used in the treeview column header
	Button() Widget
	// Clickable returns true if the user can click on the header for the
	// column.
	Clickable() bool
	// Expand returns true if the column expands to fill available space.
	Expand() bool
	// FixedWidth gets the fixed width of the column. This may not be the actual
	// displayed width of the column; for that, use
	// gtk_tree_view_column_get_width().
	FixedWidth() int
	// MaxWidth returns the maximum width in pixels of the @tree_column, or -1
	// if no maximum width is set.
	MaxWidth() int
	// MinWidth returns the minimum width in pixels of the @tree_column, or -1
	// if no minimum width is set.
	MinWidth() int
	// Reorderable returns true if the @tree_column can be reordered by the
	// user.
	Reorderable() bool
	// Resizable returns true if the @tree_column can be resized by the end
	// user.
	Resizable() bool
	// Sizing returns the current type of @tree_column.
	Sizing() TreeViewColumnSizing
	// SortColumnID gets the logical @sort_column_id that the model sorts on
	// when this column is selected for sorting. See
	// gtk_tree_view_column_set_sort_column_id().
	SortColumnID() int
	// SortIndicator gets the value set by
	// gtk_tree_view_column_set_sort_indicator().
	SortIndicator() bool
	// SortOrder gets the value set by gtk_tree_view_column_set_sort_order().
	SortOrder() SortType
	// Spacing returns the spacing of @tree_column.
	Spacing() int
	// Title returns the title of the widget.
	Title() string
	// TreeView returns the TreeView wherein @tree_column has been inserted. If
	// @column is currently not inserted in any tree view, nil is returned.
	TreeView() Widget
	// Visible returns true if @tree_column is visible.
	Visible() bool
	// Widget returns the Widget in the button on the column header. If a custom
	// widget has not been set then nil is returned.
	Widget() Widget
	// Width returns the current size of @tree_column in pixels.
	Width() int
	// XOffset returns the current X offset of @tree_column in pixels.
	XOffset() int
	// PackEndTreeViewColumn adds the @cell to end of the column. If @expand is
	// false, then the @cell is allocated no more space than it needs. Any
	// unused space is divided evenly between cells for which @expand is true.
	PackEndTreeViewColumn(cell CellRenderer, expand bool)
	// PackStartTreeViewColumn packs the @cell into the beginning of the column.
	// If @expand is false, then the @cell is allocated no more space than it
	// needs. Any unused space is divided evenly between cells for which @expand
	// is true.
	PackStartTreeViewColumn(cell CellRenderer, expand bool)
	// QueueResizeTreeViewColumn flags the column, and the cell renderers added
	// to this column, to have their sizes renegotiated.
	QueueResizeTreeViewColumn()
	// SetAlignmentTreeViewColumn sets the alignment of the title or custom
	// widget inside the column header. The alignment determines its location
	// inside the button -- 0.0 for left, 0.5 for center, 1.0 for right.
	SetAlignmentTreeViewColumn(xalign float32)
	// SetClickableTreeViewColumn sets the header to be active if @clickable is
	// true. When the header is active, then it can take keyboard focus, and can
	// be clicked.
	SetClickableTreeViewColumn(clickable bool)
	// SetExpandTreeViewColumn sets the column to take available extra space.
	// This space is shared equally amongst all columns that have the expand set
	// to true. If no column has this option set, then the last column gets all
	// extra space. By default, every column is created with this false.
	//
	// Along with “fixed-width”, the “expand” property changes when the column
	// is resized by the user.
	SetExpandTreeViewColumn(expand bool)
	// SetFixedWidthTreeViewColumn: if @fixed_width is not -1, sets the fixed
	// width of @tree_column; otherwise unsets it. The effective value of
	// @fixed_width is clamped between the minimum and maximum width of the
	// column; however, the value stored in the “fixed-width” property is not
	// clamped. If the column sizing is K_TREE_VIEW_COLUMN_GROW_ONLY or
	// K_TREE_VIEW_COLUMN_AUTOSIZE, setting a fixed width overrides the
	// automatically calculated width. Note that @fixed_width is only a hint to
	// GTK+; the width actually allocated to the column may be greater or less
	// than requested.
	//
	// Along with “expand”, the “fixed-width” property changes when the column
	// is resized by the user.
	SetFixedWidthTreeViewColumn(fixedWidth int)
	// SetMaxWidthTreeViewColumn sets the maximum width of the @tree_column. If
	// @max_width is -1, then the maximum width is unset. Note, the column can
	// actually be wider than max width if it’s the last column in a view. In
	// this case, the column expands to fill any extra space.
	SetMaxWidthTreeViewColumn(maxWidth int)
	// SetMinWidthTreeViewColumn sets the minimum width of the @tree_column. If
	// @min_width is -1, then the minimum width is unset.
	SetMinWidthTreeViewColumn(minWidth int)
	// SetReorderableTreeViewColumn: if @reorderable is true, then the column
	// can be reordered by the end user dragging the header.
	SetReorderableTreeViewColumn(reorderable bool)
	// SetResizableTreeViewColumn: if @resizable is true, then the user can
	// explicitly resize the column by grabbing the outer edge of the column
	// button. If resizable is true and sizing mode of the column is
	// K_TREE_VIEW_COLUMN_AUTOSIZE, then the sizing mode is changed to
	// K_TREE_VIEW_COLUMN_GROW_ONLY.
	SetResizableTreeViewColumn(resizable bool)
	// SetSizingTreeViewColumn sets the growth behavior of @tree_column to
	// @type.
	SetSizingTreeViewColumn(typ TreeViewColumnSizing)
	// SetSortColumnIDTreeViewColumn sets the logical @sort_column_id that this
	// column sorts on when this column is selected for sorting. Doing so makes
	// the column header clickable.
	SetSortColumnIDTreeViewColumn(sortColumnId int)
	// SetSortIndicatorTreeViewColumn: call this function with a @setting of
	// true to display an arrow in the header button indicating the column is
	// sorted. Call gtk_tree_view_column_set_sort_order() to change the
	// direction of the arrow.
	SetSortIndicatorTreeViewColumn(setting bool)
	// SetSortOrderTreeViewColumn changes the appearance of the sort indicator.
	//
	// This does not actually sort the model. Use
	// gtk_tree_view_column_set_sort_column_id() if you want automatic sorting
	// support. This function is primarily for custom sorting behavior, and
	// should be used in conjunction with gtk_tree_sortable_set_sort_column_id()
	// to do that. For custom models, the mechanism will vary.
	//
	// The sort indicator changes direction to indicate normal sort or reverse
	// sort. Note that you must have the sort indicator enabled to see anything
	// when calling this function; see
	// gtk_tree_view_column_set_sort_indicator().
	SetSortOrderTreeViewColumn(order SortType)
	// SetSpacingTreeViewColumn sets the spacing field of @tree_column, which is
	// the number of pixels to place between cell renderers packed into it.
	SetSpacingTreeViewColumn(spacing int)
	// SetTitleTreeViewColumn sets the title of the @tree_column. If a custom
	// widget has been set, then this value is ignored.
	SetTitleTreeViewColumn(title string)
	// SetVisibleTreeViewColumn sets the visibility of @tree_column.
	SetVisibleTreeViewColumn(visible bool)
	// SetWidgetTreeViewColumn sets the widget in the header to be @widget. If
	// widget is nil, then the header button is set with a Label set to the
	// title of @tree_column.
	SetWidgetTreeViewColumn(widget Widget)
}

// treeViewColumn implements the TreeViewColumn class.
type treeViewColumn struct {
	gextras.Objector
}

// WrapTreeViewColumn wraps a GObject to the right type. It is
// primarily used internally.
func WrapTreeViewColumn(obj *externglib.Object) TreeViewColumn {
	return treeViewColumn{
		Objector: obj,
	}
}

func marshalTreeViewColumn(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeViewColumn(obj), nil
}

// NewTreeViewColumn creates a new TreeViewColumn.
func NewTreeViewColumn() TreeViewColumn {
	var _cret *C.GtkTreeViewColumn // in

	_cret = C.gtk_tree_view_column_new()

	var _treeViewColumn TreeViewColumn // out

	_treeViewColumn = WrapTreeViewColumn(externglib.Take(unsafe.Pointer(_cret)))

	return _treeViewColumn
}

// NewTreeViewColumnWithArea creates a new TreeViewColumn using @area to render
// its cells.
func NewTreeViewColumnWithArea(area CellArea) TreeViewColumn {
	var _arg1 *C.GtkCellArea       // out
	var _cret *C.GtkTreeViewColumn // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_tree_view_column_new_with_area(_arg1)

	var _treeViewColumn TreeViewColumn // out

	_treeViewColumn = WrapTreeViewColumn(externglib.Take(unsafe.Pointer(_cret)))

	return _treeViewColumn
}

func (t treeViewColumn) AddAttributeTreeViewColumn(cellRenderer CellRenderer, attribute string, column int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out
	var _arg2 *C.gchar             // out
	var _arg3 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer.Native()))
	_arg2 = (*C.gchar)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gint(column)

	C.gtk_tree_view_column_add_attribute(_arg0, _arg1, _arg2, _arg3)
}

func (t treeViewColumn) CellGetPositionTreeViewColumn(cellRenderer CellRenderer) (xOffset int, width int, ok bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out
	var _arg2 *C.gint              // in
	var _arg3 *C.gint              // in
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer.Native()))

	_cret = C.gtk_tree_view_column_cell_get_position(_arg0, _arg1, &_arg2, &_arg3)

	var _xOffset int // out
	var _width int   // out
	var _ok bool     // out

	_xOffset = int(_arg2)
	_width = int(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _xOffset, _width, _ok
}

func (t treeViewColumn) CellGetSizeTreeViewColumn(cellArea gdk.Rectangle) (xOffset int, yOffset int, width int, height int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GdkRectangle      // out
	var _arg2 *C.gint              // in
	var _arg3 *C.gint              // in
	var _arg4 *C.gint              // in
	var _arg5 *C.gint              // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(cellArea))

	C.gtk_tree_view_column_cell_get_size(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)

	var _xOffset int // out
	var _yOffset int // out
	var _width int   // out
	var _height int  // out

	_xOffset = int(_arg2)
	_yOffset = int(_arg3)
	_width = int(_arg4)
	_height = int(_arg5)

	return _xOffset, _yOffset, _width, _height
}

func (t treeViewColumn) CellIsVisibleTreeViewColumn() bool {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_cell_is_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t treeViewColumn) CellSetCellDataTreeViewColumn(treeModel TreeModel, iter TreeIter, isExpander bool, isExpanded bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkTreeModel      // out
	var _arg2 *C.GtkTreeIter       // out
	var _arg3 C.gboolean           // out
	var _arg4 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(treeModel.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(iter))
	if isExpander {
		_arg3 = C.TRUE
	}
	if isExpanded {
		_arg4 = C.TRUE
	}

	C.gtk_tree_view_column_cell_set_cell_data(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (t treeViewColumn) ClearTreeViewColumn() {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	C.gtk_tree_view_column_clear(_arg0)
}

func (t treeViewColumn) ClearAttributesTreeViewColumn(cellRenderer CellRenderer) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer.Native()))

	C.gtk_tree_view_column_clear_attributes(_arg0, _arg1)
}

func (t treeViewColumn) ClickedTreeViewColumn() {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	C.gtk_tree_view_column_clicked(_arg0)
}

func (t treeViewColumn) FocusCellTreeViewColumn(cell CellRenderer) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_tree_view_column_focus_cell(_arg0, _arg1)
}

func (t treeViewColumn) Alignment() float32 {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gfloat             // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_alignment(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

func (t treeViewColumn) Button() Widget {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_button(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (t treeViewColumn) Clickable() bool {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_clickable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t treeViewColumn) Expand() bool {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_expand(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t treeViewColumn) FixedWidth() int {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_fixed_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (t treeViewColumn) MaxWidth() int {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_max_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (t treeViewColumn) MinWidth() int {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_min_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (t treeViewColumn) Reorderable() bool {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_reorderable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t treeViewColumn) Resizable() bool {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_resizable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t treeViewColumn) Sizing() TreeViewColumnSizing {
	var _arg0 *C.GtkTreeViewColumn      // out
	var _cret C.GtkTreeViewColumnSizing // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_sizing(_arg0)

	var _treeViewColumnSizing TreeViewColumnSizing // out

	_treeViewColumnSizing = TreeViewColumnSizing(_cret)

	return _treeViewColumnSizing
}

func (t treeViewColumn) SortColumnID() int {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_sort_column_id(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (t treeViewColumn) SortIndicator() bool {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_sort_indicator(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t treeViewColumn) SortOrder() SortType {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.GtkSortType        // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_sort_order(_arg0)

	var _sortType SortType // out

	_sortType = SortType(_cret)

	return _sortType
}

func (t treeViewColumn) Spacing() int {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_spacing(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (t treeViewColumn) Title() string {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (t treeViewColumn) TreeView() Widget {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_tree_view(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (t treeViewColumn) Visible() bool {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t treeViewColumn) Widget() Widget {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (t treeViewColumn) Width() int {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (t treeViewColumn) XOffset() int {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_tree_view_column_get_x_offset(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (t treeViewColumn) PackEndTreeViewColumn(cell CellRenderer, expand bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out
	var _arg2 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if expand {
		_arg2 = C.TRUE
	}

	C.gtk_tree_view_column_pack_end(_arg0, _arg1, _arg2)
}

func (t treeViewColumn) PackStartTreeViewColumn(cell CellRenderer, expand bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out
	var _arg2 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if expand {
		_arg2 = C.TRUE
	}

	C.gtk_tree_view_column_pack_start(_arg0, _arg1, _arg2)
}

func (t treeViewColumn) QueueResizeTreeViewColumn() {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	C.gtk_tree_view_column_queue_resize(_arg0)
}

func (t treeViewColumn) SetAlignmentTreeViewColumn(xalign float32) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gfloat             // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = C.gfloat(xalign)

	C.gtk_tree_view_column_set_alignment(_arg0, _arg1)
}

func (t treeViewColumn) SetClickableTreeViewColumn(clickable bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if clickable {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_clickable(_arg0, _arg1)
}

func (t treeViewColumn) SetExpandTreeViewColumn(expand bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if expand {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_expand(_arg0, _arg1)
}

func (t treeViewColumn) SetFixedWidthTreeViewColumn(fixedWidth int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(fixedWidth)

	C.gtk_tree_view_column_set_fixed_width(_arg0, _arg1)
}

func (t treeViewColumn) SetMaxWidthTreeViewColumn(maxWidth int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(maxWidth)

	C.gtk_tree_view_column_set_max_width(_arg0, _arg1)
}

func (t treeViewColumn) SetMinWidthTreeViewColumn(minWidth int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(minWidth)

	C.gtk_tree_view_column_set_min_width(_arg0, _arg1)
}

func (t treeViewColumn) SetReorderableTreeViewColumn(reorderable bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if reorderable {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_reorderable(_arg0, _arg1)
}

func (t treeViewColumn) SetResizableTreeViewColumn(resizable bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if resizable {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_resizable(_arg0, _arg1)
}

func (t treeViewColumn) SetSizingTreeViewColumn(typ TreeViewColumnSizing) {
	var _arg0 *C.GtkTreeViewColumn      // out
	var _arg1 C.GtkTreeViewColumnSizing // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = C.GtkTreeViewColumnSizing(typ)

	C.gtk_tree_view_column_set_sizing(_arg0, _arg1)
}

func (t treeViewColumn) SetSortColumnIDTreeViewColumn(sortColumnId int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(sortColumnId)

	C.gtk_tree_view_column_set_sort_column_id(_arg0, _arg1)
}

func (t treeViewColumn) SetSortIndicatorTreeViewColumn(setting bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_sort_indicator(_arg0, _arg1)
}

func (t treeViewColumn) SetSortOrderTreeViewColumn(order SortType) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.GtkSortType        // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = C.GtkSortType(order)

	C.gtk_tree_view_column_set_sort_order(_arg0, _arg1)
}

func (t treeViewColumn) SetSpacingTreeViewColumn(spacing int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(spacing)

	C.gtk_tree_view_column_set_spacing(_arg0, _arg1)
}

func (t treeViewColumn) SetTitleTreeViewColumn(title string) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.gchar             // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_tree_view_column_set_title(_arg0, _arg1)
}

func (t treeViewColumn) SetVisibleTreeViewColumn(visible bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_visible(_arg0, _arg1)
}

func (t treeViewColumn) SetWidgetTreeViewColumn(widget Widget) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkWidget         // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_tree_view_column_set_widget(_arg0, _arg1)
}

func (t treeViewColumn) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(t))
}

func (t treeViewColumn) AsCellLayout() CellLayout {
	return WrapCellLayout(gextras.InternObject(t))
}
