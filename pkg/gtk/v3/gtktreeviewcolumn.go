// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.gtk_tree_view_column_get_type()), F: marshalTreeViewColumn},
	})
}

// TreeCellDataFunc: a function to set the properties of a cell instead of just
// using the straight mapping between the cell and the model. This is useful for
// customizing the cell renderer. For example, a function might get an integer
// from the @tree_model, and render it to the “text” attribute of “cell” by
// converting it to its written equivalent. This is set by calling
// gtk_tree_view_column_set_cell_data_func()
type TreeCellDataFunc func(treeColumn TreeViewColumn, cell CellRenderer, treeModel TreeModel, iter *TreeIter)

//export gotk4_TreeCellDataFunc
func gotk4_TreeCellDataFunc(arg0 *C.GtkTreeViewColumn, arg1 *C.GtkCellRenderer, arg2 *C.GtkTreeModel, arg3 *C.GtkTreeIter, arg4 C.gpointer) {
	v := box.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(TreeCellDataFunc)
	fn(treeColumn, cell, treeModel, iter, data)
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
	Buildable
	CellLayout

	// AddAttribute adds an attribute mapping to the list in @tree_column. The
	// @column is the column of the model to get a value from, and the
	// @attribute is the parameter on @cell_renderer to be set from the value.
	// So for example if column 2 of the model contains strings, you could have
	// the “text” attribute of a CellRendererText get its values from column 2.
	AddAttribute(cellRenderer CellRenderer, attribute string, column int)
	// CellGetPosition obtains the horizontal position and size of a cell in a
	// column. If the cell is not found in the column, @start_pos and @width are
	// not changed and false is returned.
	CellGetPosition(cellRenderer CellRenderer) (xOffset int, width int, ok bool)
	// CellGetSize obtains the width and height needed to render the column.
	// This is used primarily by the TreeView.
	CellGetSize(cellArea *gdk.Rectangle) (xOffset int, yOffset int, width int, height int)
	// CellIsVisible returns true if any of the cells packed into the
	// @tree_column are visible. For this to be meaningful, you must first
	// initialize the cells with gtk_tree_view_column_cell_set_cell_data()
	CellIsVisible() bool
	// CellSetCellData sets the cell renderer based on the @tree_model and
	// @iter. That is, for every attribute mapping in @tree_column, it will get
	// a value from the set column on the @iter, and use that value to set the
	// attribute on the cell renderer. This is used primarily by the TreeView.
	CellSetCellData(treeModel TreeModel, iter *TreeIter, isExpander bool, isExpanded bool)
	// Clear unsets all the mappings on all renderers on the @tree_column.
	Clear()
	// ClearAttributes clears all existing attributes previously set with
	// gtk_tree_view_column_set_attributes().
	ClearAttributes(cellRenderer CellRenderer)
	// Clicked emits the “clicked” signal on the column. This function will only
	// work if @tree_column is clickable.
	Clicked()
	// FocusCell sets the current keyboard focus to be at @cell, if the column
	// contains 2 or more editable and activatable cells.
	FocusCell(cell CellRenderer)
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
	// PackEnd adds the @cell to end of the column. If @expand is false, then
	// the @cell is allocated no more space than it needs. Any unused space is
	// divided evenly between cells for which @expand is true.
	PackEnd(cell CellRenderer, expand bool)
	// PackStart packs the @cell into the beginning of the column. If @expand is
	// false, then the @cell is allocated no more space than it needs. Any
	// unused space is divided evenly between cells for which @expand is true.
	PackStart(cell CellRenderer, expand bool)
	// QueueResize flags the column, and the cell renderers added to this
	// column, to have their sizes renegotiated.
	QueueResize()
	// SetAlignment sets the alignment of the title or custom widget inside the
	// column header. The alignment determines its location inside the button --
	// 0.0 for left, 0.5 for center, 1.0 for right.
	SetAlignment(xalign float32)
	// SetCellDataFunc sets the TreeCellDataFunc to use for the column. This
	// function is used instead of the standard attributes mapping for setting
	// the column value, and should set the value of @tree_column's cell
	// renderer as appropriate. @func may be nil to remove an older one.
	SetCellDataFunc(cellRenderer CellRenderer, fn TreeCellDataFunc)
	// SetClickable sets the header to be active if @clickable is true. When the
	// header is active, then it can take keyboard focus, and can be clicked.
	SetClickable(clickable bool)
	// SetExpand sets the column to take available extra space. This space is
	// shared equally amongst all columns that have the expand set to true. If
	// no column has this option set, then the last column gets all extra space.
	// By default, every column is created with this false.
	//
	// Along with “fixed-width”, the “expand” property changes when the column
	// is resized by the user.
	SetExpand(expand bool)
	// SetFixedWidth: if @fixed_width is not -1, sets the fixed width of
	// @tree_column; otherwise unsets it. The effective value of @fixed_width is
	// clamped between the minimum and maximum width of the column; however, the
	// value stored in the “fixed-width” property is not clamped. If the column
	// sizing is K_TREE_VIEW_COLUMN_GROW_ONLY or K_TREE_VIEW_COLUMN_AUTOSIZE,
	// setting a fixed width overrides the automatically calculated width. Note
	// that @fixed_width is only a hint to GTK+; the width actually allocated to
	// the column may be greater or less than requested.
	//
	// Along with “expand”, the “fixed-width” property changes when the column
	// is resized by the user.
	SetFixedWidth(fixedWidth int)
	// SetMaxWidth sets the maximum width of the @tree_column. If @max_width is
	// -1, then the maximum width is unset. Note, the column can actually be
	// wider than max width if it’s the last column in a view. In this case, the
	// column expands to fill any extra space.
	SetMaxWidth(maxWidth int)
	// SetMinWidth sets the minimum width of the @tree_column. If @min_width is
	// -1, then the minimum width is unset.
	SetMinWidth(minWidth int)
	// SetReorderable: if @reorderable is true, then the column can be reordered
	// by the end user dragging the header.
	SetReorderable(reorderable bool)
	// SetResizable: if @resizable is true, then the user can explicitly resize
	// the column by grabbing the outer edge of the column button. If resizable
	// is true and sizing mode of the column is K_TREE_VIEW_COLUMN_AUTOSIZE,
	// then the sizing mode is changed to K_TREE_VIEW_COLUMN_GROW_ONLY.
	SetResizable(resizable bool)
	// SetSizing sets the growth behavior of @tree_column to @type.
	SetSizing(typ TreeViewColumnSizing)
	// SetSortColumnID sets the logical @sort_column_id that this column sorts
	// on when this column is selected for sorting. Doing so makes the column
	// header clickable.
	SetSortColumnID(sortColumnID int)
	// SetSortIndicator: call this function with a @setting of true to display
	// an arrow in the header button indicating the column is sorted. Call
	// gtk_tree_view_column_set_sort_order() to change the direction of the
	// arrow.
	SetSortIndicator(setting bool)
	// SetSortOrder changes the appearance of the sort indicator.
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
	SetSortOrder(order SortType)
	// SetSpacing sets the spacing field of @tree_column, which is the number of
	// pixels to place between cell renderers packed into it.
	SetSpacing(spacing int)
	// SetTitle sets the title of the @tree_column. If a custom widget has been
	// set, then this value is ignored.
	SetTitle(title string)
	// SetVisible sets the visibility of @tree_column.
	SetVisible(visible bool)
	// SetWidget sets the widget in the header to be @widget. If widget is nil,
	// then the header button is set with a Label set to the title of
	// @tree_column.
	SetWidget(widget Widget)
}

// treeViewColumn implements the TreeViewColumn interface.
type treeViewColumn struct {
	gextras.Objector
	Buildable
	CellLayout
}

var _ TreeViewColumn = (*treeViewColumn)(nil)

// WrapTreeViewColumn wraps a GObject to the right type. It is
// primarily used internally.
func WrapTreeViewColumn(obj *externglib.Object) TreeViewColumn {
	return TreeViewColumn{
		Objector:   obj,
		Buildable:  WrapBuildable(obj),
		CellLayout: WrapCellLayout(obj),
	}
}

func marshalTreeViewColumn(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeViewColumn(obj), nil
}

// NewTreeViewColumn constructs a class TreeViewColumn.
func NewTreeViewColumn() TreeViewColumn {
	var cret C.GtkTreeViewColumn
	var ret1 TreeViewColumn

	cret = C.gtk_tree_view_column_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(TreeViewColumn)

	return ret1
}

// NewTreeViewColumnWithArea constructs a class TreeViewColumn.
func NewTreeViewColumnWithArea(area CellArea) TreeViewColumn {
	var arg1 *C.GtkCellArea

	arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))

	var cret C.GtkTreeViewColumn
	var ret1 TreeViewColumn

	cret = C.gtk_tree_view_column_new_with_area(area)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(TreeViewColumn)

	return ret1
}

// AddAttribute adds an attribute mapping to the list in @tree_column. The
// @column is the column of the model to get a value from, and the
// @attribute is the parameter on @cell_renderer to be set from the value.
// So for example if column 2 of the model contains strings, you could have
// the “text” attribute of a CellRendererText get its values from column 2.
func (t treeViewColumn) AddAttribute(cellRenderer CellRenderer, attribute string, column int) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 *C.GtkCellRenderer
	var arg2 *C.gchar
	var arg3 C.gint

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer.Native()))
	arg2 = (*C.gchar)(C.CString(attribute))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.gint(column)

	C.gtk_tree_view_column_add_attribute(arg0, cellRenderer, attribute, column)
}

// CellGetPosition obtains the horizontal position and size of a cell in a
// column. If the cell is not found in the column, @start_pos and @width are
// not changed and false is returned.
func (t treeViewColumn) CellGetPosition(cellRenderer CellRenderer) (xOffset int, width int, ok bool) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 *C.GtkCellRenderer

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer.Native()))

	var arg2 C.gint
	var ret2 int
	var arg3 C.gint
	var ret3 int
	var cret C.gboolean
	var ret3 bool

	cret = C.gtk_tree_view_column_cell_get_position(arg0, cellRenderer, &arg2, &arg3)

	*ret2 = C.gint(arg2)
	*ret3 = C.gint(arg3)
	ret3 = C.bool(cret) != C.false

	return ret2, ret3, ret3
}

// CellGetSize obtains the width and height needed to render the column.
// This is used primarily by the TreeView.
func (t treeViewColumn) CellGetSize(cellArea *gdk.Rectangle) (xOffset int, yOffset int, width int, height int) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 *C.GdkRectangle

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GdkRectangle)(unsafe.Pointer(cellArea.Native()))

	var arg2 C.gint
	var ret2 int
	var arg3 C.gint
	var ret3 int
	var arg4 C.gint
	var ret4 int
	var arg5 C.gint
	var ret5 int

	C.gtk_tree_view_column_cell_get_size(arg0, cellArea, &arg2, &arg3, &arg4, &arg5)

	*ret2 = C.gint(arg2)
	*ret3 = C.gint(arg3)
	*ret4 = C.gint(arg4)
	*ret5 = C.gint(arg5)

	return ret2, ret3, ret4, ret5
}

// CellIsVisible returns true if any of the cells packed into the
// @tree_column are visible. For this to be meaningful, you must first
// initialize the cells with gtk_tree_view_column_cell_set_cell_data()
func (t treeViewColumn) CellIsVisible() bool {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_tree_view_column_cell_is_visible(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// CellSetCellData sets the cell renderer based on the @tree_model and
// @iter. That is, for every attribute mapping in @tree_column, it will get
// a value from the set column on the @iter, and use that value to set the
// attribute on the cell renderer. This is used primarily by the TreeView.
func (t treeViewColumn) CellSetCellData(treeModel TreeModel, iter *TreeIter, isExpander bool, isExpanded bool) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 *C.GtkTreeModel
	var arg2 *C.GtkTreeIter
	var arg3 C.gboolean
	var arg4 C.gboolean

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkTreeModel)(unsafe.Pointer(treeModel.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))
	if isExpander {
		arg3 = C.gboolean(1)
	}
	if isExpanded {
		arg4 = C.gboolean(1)
	}

	C.gtk_tree_view_column_cell_set_cell_data(arg0, treeModel, iter, isExpander, isExpanded)
}

// Clear unsets all the mappings on all renderers on the @tree_column.
func (t treeViewColumn) Clear() {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	C.gtk_tree_view_column_clear(arg0)
}

// ClearAttributes clears all existing attributes previously set with
// gtk_tree_view_column_set_attributes().
func (t treeViewColumn) ClearAttributes(cellRenderer CellRenderer) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 *C.GtkCellRenderer

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer.Native()))

	C.gtk_tree_view_column_clear_attributes(arg0, cellRenderer)
}

// Clicked emits the “clicked” signal on the column. This function will only
// work if @tree_column is clickable.
func (t treeViewColumn) Clicked() {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	C.gtk_tree_view_column_clicked(arg0)
}

// FocusCell sets the current keyboard focus to be at @cell, if the column
// contains 2 or more editable and activatable cells.
func (t treeViewColumn) FocusCell(cell CellRenderer) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 *C.GtkCellRenderer

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_tree_view_column_focus_cell(arg0, cell)
}

// Alignment returns the current x alignment of @tree_column. This value can
// range between 0.0 and 1.0.
func (t treeViewColumn) Alignment() float32 {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.gfloat
	var ret1 float32

	cret = C.gtk_tree_view_column_get_alignment(arg0)

	ret1 = C.gfloat(cret)

	return ret1
}

// Button returns the button used in the treeview column header
func (t treeViewColumn) Button() Widget {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_tree_view_column_get_button(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// Clickable returns true if the user can click on the header for the
// column.
func (t treeViewColumn) Clickable() bool {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_tree_view_column_get_clickable(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Expand returns true if the column expands to fill available space.
func (t treeViewColumn) Expand() bool {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_tree_view_column_get_expand(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// FixedWidth gets the fixed width of the column. This may not be the actual
// displayed width of the column; for that, use
// gtk_tree_view_column_get_width().
func (t treeViewColumn) FixedWidth() int {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.gint
	var ret1 int

	cret = C.gtk_tree_view_column_get_fixed_width(arg0)

	ret1 = C.gint(cret)

	return ret1
}

// MaxWidth returns the maximum width in pixels of the @tree_column, or -1
// if no maximum width is set.
func (t treeViewColumn) MaxWidth() int {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.gint
	var ret1 int

	cret = C.gtk_tree_view_column_get_max_width(arg0)

	ret1 = C.gint(cret)

	return ret1
}

// MinWidth returns the minimum width in pixels of the @tree_column, or -1
// if no minimum width is set.
func (t treeViewColumn) MinWidth() int {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.gint
	var ret1 int

	cret = C.gtk_tree_view_column_get_min_width(arg0)

	ret1 = C.gint(cret)

	return ret1
}

// Reorderable returns true if the @tree_column can be reordered by the
// user.
func (t treeViewColumn) Reorderable() bool {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_tree_view_column_get_reorderable(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Resizable returns true if the @tree_column can be resized by the end
// user.
func (t treeViewColumn) Resizable() bool {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_tree_view_column_get_resizable(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Sizing returns the current type of @tree_column.
func (t treeViewColumn) Sizing() TreeViewColumnSizing {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.GtkTreeViewColumnSizing
	var ret1 TreeViewColumnSizing

	cret = C.gtk_tree_view_column_get_sizing(arg0)

	ret1 = TreeViewColumnSizing(cret)

	return ret1
}

// SortColumnID gets the logical @sort_column_id that the model sorts on
// when this column is selected for sorting. See
// gtk_tree_view_column_set_sort_column_id().
func (t treeViewColumn) SortColumnID() int {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.gint
	var ret1 int

	cret = C.gtk_tree_view_column_get_sort_column_id(arg0)

	ret1 = C.gint(cret)

	return ret1
}

// SortIndicator gets the value set by
// gtk_tree_view_column_set_sort_indicator().
func (t treeViewColumn) SortIndicator() bool {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_tree_view_column_get_sort_indicator(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// SortOrder gets the value set by gtk_tree_view_column_set_sort_order().
func (t treeViewColumn) SortOrder() SortType {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.GtkSortType
	var ret1 SortType

	cret = C.gtk_tree_view_column_get_sort_order(arg0)

	ret1 = SortType(cret)

	return ret1
}

// Spacing returns the spacing of @tree_column.
func (t treeViewColumn) Spacing() int {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.gint
	var ret1 int

	cret = C.gtk_tree_view_column_get_spacing(arg0)

	ret1 = C.gint(cret)

	return ret1
}

// Title returns the title of the widget.
func (t treeViewColumn) Title() string {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_tree_view_column_get_title(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// TreeView returns the TreeView wherein @tree_column has been inserted. If
// @column is currently not inserted in any tree view, nil is returned.
func (t treeViewColumn) TreeView() Widget {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_tree_view_column_get_tree_view(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// Visible returns true if @tree_column is visible.
func (t treeViewColumn) Visible() bool {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_tree_view_column_get_visible(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Widget returns the Widget in the button on the column header. If a custom
// widget has not been set then nil is returned.
func (t treeViewColumn) Widget() Widget {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_tree_view_column_get_widget(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// Width returns the current size of @tree_column in pixels.
func (t treeViewColumn) Width() int {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.gint
	var ret1 int

	cret = C.gtk_tree_view_column_get_width(arg0)

	ret1 = C.gint(cret)

	return ret1
}

// XOffset returns the current X offset of @tree_column in pixels.
func (t treeViewColumn) XOffset() int {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var cret C.gint
	var ret1 int

	cret = C.gtk_tree_view_column_get_x_offset(arg0)

	ret1 = C.gint(cret)

	return ret1
}

// PackEnd adds the @cell to end of the column. If @expand is false, then
// the @cell is allocated no more space than it needs. Any unused space is
// divided evenly between cells for which @expand is true.
func (t treeViewColumn) PackEnd(cell CellRenderer, expand bool) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 *C.GtkCellRenderer
	var arg2 C.gboolean

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if expand {
		arg2 = C.gboolean(1)
	}

	C.gtk_tree_view_column_pack_end(arg0, cell, expand)
}

// PackStart packs the @cell into the beginning of the column. If @expand is
// false, then the @cell is allocated no more space than it needs. Any
// unused space is divided evenly between cells for which @expand is true.
func (t treeViewColumn) PackStart(cell CellRenderer, expand bool) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 *C.GtkCellRenderer
	var arg2 C.gboolean

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if expand {
		arg2 = C.gboolean(1)
	}

	C.gtk_tree_view_column_pack_start(arg0, cell, expand)
}

// QueueResize flags the column, and the cell renderers added to this
// column, to have their sizes renegotiated.
func (t treeViewColumn) QueueResize() {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	C.gtk_tree_view_column_queue_resize(arg0)
}

// SetAlignment sets the alignment of the title or custom widget inside the
// column header. The alignment determines its location inside the button --
// 0.0 for left, 0.5 for center, 1.0 for right.
func (t treeViewColumn) SetAlignment(xalign float32) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 C.gfloat

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = C.gfloat(xalign)

	C.gtk_tree_view_column_set_alignment(arg0, xalign)
}

// SetCellDataFunc sets the TreeCellDataFunc to use for the column. This
// function is used instead of the standard attributes mapping for setting
// the column value, and should set the value of @tree_column's cell
// renderer as appropriate. @func may be nil to remove an older one.
func (t treeViewColumn) SetCellDataFunc(cellRenderer CellRenderer, fn TreeCellDataFunc) {
	var arg0 *C.GtkTreeViewColumn

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	C.gtk_tree_view_column_set_cell_data_func(arg0, cellRenderer, fn, funcData, destroy)
}

// SetClickable sets the header to be active if @clickable is true. When the
// header is active, then it can take keyboard focus, and can be clicked.
func (t treeViewColumn) SetClickable(clickable bool) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 C.gboolean

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if clickable {
		arg1 = C.gboolean(1)
	}

	C.gtk_tree_view_column_set_clickable(arg0, clickable)
}

// SetExpand sets the column to take available extra space. This space is
// shared equally amongst all columns that have the expand set to true. If
// no column has this option set, then the last column gets all extra space.
// By default, every column is created with this false.
//
// Along with “fixed-width”, the “expand” property changes when the column
// is resized by the user.
func (t treeViewColumn) SetExpand(expand bool) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 C.gboolean

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if expand {
		arg1 = C.gboolean(1)
	}

	C.gtk_tree_view_column_set_expand(arg0, expand)
}

// SetFixedWidth: if @fixed_width is not -1, sets the fixed width of
// @tree_column; otherwise unsets it. The effective value of @fixed_width is
// clamped between the minimum and maximum width of the column; however, the
// value stored in the “fixed-width” property is not clamped. If the column
// sizing is K_TREE_VIEW_COLUMN_GROW_ONLY or K_TREE_VIEW_COLUMN_AUTOSIZE,
// setting a fixed width overrides the automatically calculated width. Note
// that @fixed_width is only a hint to GTK+; the width actually allocated to
// the column may be greater or less than requested.
//
// Along with “expand”, the “fixed-width” property changes when the column
// is resized by the user.
func (t treeViewColumn) SetFixedWidth(fixedWidth int) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 C.gint

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = C.gint(fixedWidth)

	C.gtk_tree_view_column_set_fixed_width(arg0, fixedWidth)
}

// SetMaxWidth sets the maximum width of the @tree_column. If @max_width is
// -1, then the maximum width is unset. Note, the column can actually be
// wider than max width if it’s the last column in a view. In this case, the
// column expands to fill any extra space.
func (t treeViewColumn) SetMaxWidth(maxWidth int) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 C.gint

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = C.gint(maxWidth)

	C.gtk_tree_view_column_set_max_width(arg0, maxWidth)
}

// SetMinWidth sets the minimum width of the @tree_column. If @min_width is
// -1, then the minimum width is unset.
func (t treeViewColumn) SetMinWidth(minWidth int) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 C.gint

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = C.gint(minWidth)

	C.gtk_tree_view_column_set_min_width(arg0, minWidth)
}

// SetReorderable: if @reorderable is true, then the column can be reordered
// by the end user dragging the header.
func (t treeViewColumn) SetReorderable(reorderable bool) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 C.gboolean

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if reorderable {
		arg1 = C.gboolean(1)
	}

	C.gtk_tree_view_column_set_reorderable(arg0, reorderable)
}

// SetResizable: if @resizable is true, then the user can explicitly resize
// the column by grabbing the outer edge of the column button. If resizable
// is true and sizing mode of the column is K_TREE_VIEW_COLUMN_AUTOSIZE,
// then the sizing mode is changed to K_TREE_VIEW_COLUMN_GROW_ONLY.
func (t treeViewColumn) SetResizable(resizable bool) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 C.gboolean

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if resizable {
		arg1 = C.gboolean(1)
	}

	C.gtk_tree_view_column_set_resizable(arg0, resizable)
}

// SetSizing sets the growth behavior of @tree_column to @type.
func (t treeViewColumn) SetSizing(typ TreeViewColumnSizing) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 C.GtkTreeViewColumnSizing

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = (C.GtkTreeViewColumnSizing)(typ)

	C.gtk_tree_view_column_set_sizing(arg0, typ)
}

// SetSortColumnID sets the logical @sort_column_id that this column sorts
// on when this column is selected for sorting. Doing so makes the column
// header clickable.
func (t treeViewColumn) SetSortColumnID(sortColumnID int) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 C.gint

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = C.gint(sortColumnID)

	C.gtk_tree_view_column_set_sort_column_id(arg0, sortColumnID)
}

// SetSortIndicator: call this function with a @setting of true to display
// an arrow in the header button indicating the column is sorted. Call
// gtk_tree_view_column_set_sort_order() to change the direction of the
// arrow.
func (t treeViewColumn) SetSortIndicator(setting bool) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 C.gboolean

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_tree_view_column_set_sort_indicator(arg0, setting)
}

// SetSortOrder changes the appearance of the sort indicator.
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
func (t treeViewColumn) SetSortOrder(order SortType) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 C.GtkSortType

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = (C.GtkSortType)(order)

	C.gtk_tree_view_column_set_sort_order(arg0, order)
}

// SetSpacing sets the spacing field of @tree_column, which is the number of
// pixels to place between cell renderers packed into it.
func (t treeViewColumn) SetSpacing(spacing int) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 C.gint

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = C.gint(spacing)

	C.gtk_tree_view_column_set_spacing(arg0, spacing)
}

// SetTitle sets the title of the @tree_column. If a custom widget has been
// set, then this value is ignored.
func (t treeViewColumn) SetTitle(title string) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 *C.gchar

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_tree_view_column_set_title(arg0, title)
}

// SetVisible sets the visibility of @tree_column.
func (t treeViewColumn) SetVisible(visible bool) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 C.gboolean

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if visible {
		arg1 = C.gboolean(1)
	}

	C.gtk_tree_view_column_set_visible(arg0, visible)
}

// SetWidget sets the widget in the header to be @widget. If widget is nil,
// then the header button is set with a Label set to the title of
// @tree_column.
func (t treeViewColumn) SetWidget(widget Widget) {
	var arg0 *C.GtkTreeViewColumn
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_tree_view_column_set_widget(arg0, widget)
}
