// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
// void gotk4_ListBoxForeachFunc(GtkListBox*, GtkListBoxRow*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_list_box_get_type()), F: marshalListBoxer},
		{T: externglib.Type(C.gtk_list_box_row_get_type()), F: marshalListBoxRower},
	})
}

// ListBoxCreateWidgetFunc: called for list boxes that are bound to a
// `GListModel` with gtk_list_box_bind_model() for each item that gets added to
// the model.
type ListBoxCreateWidgetFunc func(item *externglib.Object, userData cgo.Handle) (widget *Widget)

//export gotk4_ListBoxCreateWidgetFunc
func gotk4_ListBoxCreateWidgetFunc(arg0 C.gpointer, arg1 C.gpointer) (cret *C.GtkWidget) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var item *externglib.Object // out
	var userData cgo.Handle     // out

	item = externglib.Take(unsafe.Pointer(arg0))
	userData = (cgo.Handle)(unsafe.Pointer(arg1))

	fn := v.(ListBoxCreateWidgetFunc)
	widget := fn(item, userData)

	cret = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	return cret
}

// ListBoxFilterFunc: will be called whenever the row changes or is added and
// lets you control if the row should be visible or not.
type ListBoxFilterFunc func(row *ListBoxRow, userData cgo.Handle) (ok bool)

//export gotk4_ListBoxFilterFunc
func gotk4_ListBoxFilterFunc(arg0 *C.GtkListBoxRow, arg1 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var row *ListBoxRow     // out
	var userData cgo.Handle // out

	row = wrapListBoxRow(externglib.Take(unsafe.Pointer(arg0)))
	userData = (cgo.Handle)(unsafe.Pointer(arg1))

	fn := v.(ListBoxFilterFunc)
	ok := fn(row, userData)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ListBoxForeachFunc: function used by gtk_list_box_selected_foreach().
//
// It will be called on every selected child of the @box.
type ListBoxForeachFunc func(box *ListBox, row *ListBoxRow, userData cgo.Handle)

//export gotk4_ListBoxForeachFunc
func gotk4_ListBoxForeachFunc(arg0 *C.GtkListBox, arg1 *C.GtkListBoxRow, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var box *ListBox        // out
	var row *ListBoxRow     // out
	var userData cgo.Handle // out

	box = wrapListBox(externglib.Take(unsafe.Pointer(arg0)))
	row = wrapListBoxRow(externglib.Take(unsafe.Pointer(arg1)))
	userData = (cgo.Handle)(unsafe.Pointer(arg2))

	fn := v.(ListBoxForeachFunc)
	fn(box, row, userData)
}

// ListBoxSortFunc: compare two rows to determine which should be first.
type ListBoxSortFunc func(row1 *ListBoxRow, row2 *ListBoxRow, userData cgo.Handle) (gint int)

//export gotk4_ListBoxSortFunc
func gotk4_ListBoxSortFunc(arg0 *C.GtkListBoxRow, arg1 *C.GtkListBoxRow, arg2 C.gpointer) (cret C.int) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var row1 *ListBoxRow    // out
	var row2 *ListBoxRow    // out
	var userData cgo.Handle // out

	row1 = wrapListBoxRow(externglib.Take(unsafe.Pointer(arg0)))
	row2 = wrapListBoxRow(externglib.Take(unsafe.Pointer(arg1)))
	userData = (cgo.Handle)(unsafe.Pointer(arg2))

	fn := v.(ListBoxSortFunc)
	gint := fn(row1, row2, userData)

	cret = C.int(gint)

	return cret
}

// ListBoxUpdateHeaderFunc: whenever @row changes or which row is before @row
// changes this is called, which lets you update the header on @row.
//
// You may remove or set a new one via [method@Gtk.ListBoxRow.set_header] or
// just change the state of the current header widget.
type ListBoxUpdateHeaderFunc func(row *ListBoxRow, before *ListBoxRow, userData cgo.Handle)

//export gotk4_ListBoxUpdateHeaderFunc
func gotk4_ListBoxUpdateHeaderFunc(arg0 *C.GtkListBoxRow, arg1 *C.GtkListBoxRow, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var row *ListBoxRow     // out
	var before *ListBoxRow  // out
	var userData cgo.Handle // out

	row = wrapListBoxRow(externglib.Take(unsafe.Pointer(arg0)))
	before = wrapListBoxRow(externglib.Take(unsafe.Pointer(arg1)))
	userData = (cgo.Handle)(unsafe.Pointer(arg2))

	fn := v.(ListBoxUpdateHeaderFunc)
	fn(row, before, userData)
}

// ListBoxer describes ListBox's methods.
type ListBoxer interface {
	// Append a widget to the list.
	Append(child Widgeter)
	// DragHighlightRow: add a drag highlight to a row.
	DragHighlightRow(row ListBoxRower)
	// DragUnhighlightRow: if a row has previously been highlighted via
	// gtk_list_box_drag_highlight_row(), it will have the highlight removed.
	DragUnhighlightRow()
	// ActivateOnSingleClick returns whether rows activate on single clicks.
	ActivateOnSingleClick() bool
	// Adjustment gets the adjustment (if any) that the widget uses to for
	// vertical scrolling.
	Adjustment() *Adjustment
	// RowAtIndex gets the n-th child in the list (not counting headers).
	RowAtIndex(index_ int) *ListBoxRow
	// RowAtY gets the row at the @y position.
	RowAtY(y int) *ListBoxRow
	// SelectedRow gets the selected row, or nil if no rows are selected.
	SelectedRow() *ListBoxRow
	// SelectionMode gets the selection mode of the listbox.
	SelectionMode() SelectionMode
	// ShowSeparators returns whether the list box should show separators
	// between rows.
	ShowSeparators() bool
	// Insert the @child into the @box at @position.
	Insert(child Widgeter, position int)
	// InvalidateFilter: update the filtering for all rows.
	InvalidateFilter()
	// InvalidateHeaders: update the separators for all rows.
	InvalidateHeaders()
	// InvalidateSort: update the sorting for all rows.
	InvalidateSort()
	// Prepend a widget to the list.
	Prepend(child Widgeter)
	// Remove removes a child from @box.
	Remove(child Widgeter)
	// SelectAll: select all children of @box, if the selection mode allows it.
	SelectAll()
	// SelectRow: make @row the currently selected row.
	SelectRow(row ListBoxRower)
	// SelectedForeach calls a function for each selected child.
	SelectedForeach(fn ListBoxForeachFunc)
	// SetActivateOnSingleClick: if @single is true, rows will be activated when
	// you click on them, otherwise you need to double-click.
	SetActivateOnSingleClick(single bool)
	// SetAdjustment sets the adjustment (if any) that the widget uses to for
	// vertical scrolling.
	SetAdjustment(adjustment Adjustmenter)
	// SetPlaceholder sets the placeholder widget that is shown in the list when
	// it doesn't display any visible children.
	SetPlaceholder(placeholder Widgeter)
	// SetSelectionMode sets how selection works in the listbox.
	SetSelectionMode(mode SelectionMode)
	// SetShowSeparators sets whether the list box should show separators
	// between rows.
	SetShowSeparators(showSeparators bool)
	// UnselectAll: unselect all children of @box, if the selection mode allows
	// it.
	UnselectAll()
	// UnselectRow unselects a single row of @box, if the selection mode allows
	// it.
	UnselectRow(row ListBoxRower)
}

// ListBox: `GtkListBox` is a vertical list.
//
// A `GtkListBox` only contains `GtkListBoxRow` children. These rows can by
// dynamically sorted and filtered, and headers can be added dynamically
// depending on the row content. It also allows keyboard and mouse navigation
// and selection like a typical list.
//
// Using `GtkListBox` is often an alternative to `GtkTreeView`, especially when
// the list contents has a more complicated layout than what is allowed by a
// `GtkCellRenderer`, or when the contents is interactive (i.e. has a button in
// it).
//
// Although a `GtkListBox` must have only `GtkListBoxRow` children, you can add
// any kind of widget to it via [method@Gtk.ListBox.prepend],
// [method@Gtk.ListBox.append] and [method@Gtk.ListBox.insert] and a
// `GtkListBoxRow` widget will automatically be inserted between the list and
// the widget.
//
// `GtkListBoxRows` can be marked as activatable or selectable. If a row is
// activatable, [signal@Gtk.ListBox::row-activated] will be emitted for it when
// the user tries to activate it. If it is selectable, the row will be marked as
// selected when the user tries to select it.
//
//
// GtkListBox as GtkBuildable
//
// The `GtkListBox` implementation of the `GtkBuildable` interface supports
// setting a child as the placeholder by specifying “placeholder” as the “type”
// attribute of a <child> element. See [method@Gtk.ListBox.set_placeholder] for
// info.
//
// CSS nodes
//
//    list[.separators][.rich-list][.navigation-sidebar]
//    ╰── row[.activatable]
//
// `GtkListBox` uses a single CSS node named list. It may carry the .separators
// style class, when the [property@Gtk.ListBox:show-separators] property is set.
// Each `GtkListBoxRow` uses a single CSS node named row. The row nodes get the
// .activatable style class added when appropriate.
//
// The main list node may also carry style classes to select the style of list
// presentation (section-list-widget.html#list-styles): .rich-list,
// .navigation-sidebar or .data-table.
//
//
// Accessibility
//
// `GtkListBox` uses the GTK_ACCESSIBLE_ROLE_LIST role and `GtkListBoxRow` uses
// the GTK_ACCESSIBLE_ROLE_LIST_ITEM role.
type ListBox struct {
	Widget
}

var (
	_ ListBoxer       = (*ListBox)(nil)
	_ gextras.Nativer = (*ListBox)(nil)
)

func wrapListBox(obj *externglib.Object) *ListBox {
	return &ListBox{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalListBoxer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapListBox(obj), nil
}

// NewListBox creates a new `GtkListBox` container.
func NewListBox() *ListBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_list_box_new()

	var _listBox *ListBox // out

	_listBox = wrapListBox(externglib.Take(unsafe.Pointer(_cret)))

	return _listBox
}

// Append a widget to the list.
//
// If a sort function is set, the widget will actually be inserted at the
// calculated position.
func (box *ListBox) Append(child Widgeter) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_list_box_append(_arg0, _arg1)
}

// DragHighlightRow: add a drag highlight to a row.
//
// This is a helper function for implementing DnD onto a `GtkListBox`. The
// passed in @row will be highlighted by setting the GTK_STATE_FLAG_DROP_ACTIVE
// state and any previously highlighted row will be unhighlighted.
//
// The row will also be unhighlighted when the widget gets a drag leave event.
func (box *ListBox) DragHighlightRow(row ListBoxRower) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer((row).(gextras.Nativer).Native()))

	C.gtk_list_box_drag_highlight_row(_arg0, _arg1)
}

// DragUnhighlightRow: if a row has previously been highlighted via
// gtk_list_box_drag_highlight_row(), it will have the highlight removed.
func (box *ListBox) DragUnhighlightRow() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_drag_unhighlight_row(_arg0)
}

// ActivateOnSingleClick returns whether rows activate on single clicks.
func (box *ListBox) ActivateOnSingleClick() bool {
	var _arg0 *C.GtkListBox // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_activate_on_single_click(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Adjustment gets the adjustment (if any) that the widget uses to for vertical
// scrolling.
func (box *ListBox) Adjustment() *Adjustment {
	var _arg0 *C.GtkListBox    // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_adjustment(_arg0)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// RowAtIndex gets the n-th child in the list (not counting headers).
//
// If @index_ is negative or larger than the number of items in the list, nil is
// returned.
func (box *ListBox) RowAtIndex(index_ int) *ListBoxRow {
	var _arg0 *C.GtkListBox    // out
	var _arg1 C.int            // out
	var _cret *C.GtkListBoxRow // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.int(index_)

	_cret = C.gtk_list_box_get_row_at_index(_arg0, _arg1)

	var _listBoxRow *ListBoxRow // out

	_listBoxRow = wrapListBoxRow(externglib.Take(unsafe.Pointer(_cret)))

	return _listBoxRow
}

// RowAtY gets the row at the @y position.
func (box *ListBox) RowAtY(y int) *ListBoxRow {
	var _arg0 *C.GtkListBox    // out
	var _arg1 C.int            // out
	var _cret *C.GtkListBoxRow // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.int(y)

	_cret = C.gtk_list_box_get_row_at_y(_arg0, _arg1)

	var _listBoxRow *ListBoxRow // out

	_listBoxRow = wrapListBoxRow(externglib.Take(unsafe.Pointer(_cret)))

	return _listBoxRow
}

// SelectedRow gets the selected row, or nil if no rows are selected.
//
// Note that the box may allow multiple selection, in which case you should use
// [method@Gtk.ListBox.selected_foreach] to find all selected rows.
func (box *ListBox) SelectedRow() *ListBoxRow {
	var _arg0 *C.GtkListBox    // out
	var _cret *C.GtkListBoxRow // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_selected_row(_arg0)

	var _listBoxRow *ListBoxRow // out

	_listBoxRow = wrapListBoxRow(externglib.Take(unsafe.Pointer(_cret)))

	return _listBoxRow
}

// SelectionMode gets the selection mode of the listbox.
func (box *ListBox) SelectionMode() SelectionMode {
	var _arg0 *C.GtkListBox      // out
	var _cret C.GtkSelectionMode // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_selection_mode(_arg0)

	var _selectionMode SelectionMode // out

	_selectionMode = SelectionMode(_cret)

	return _selectionMode
}

// ShowSeparators returns whether the list box should show separators between
// rows.
func (box *ListBox) ShowSeparators() bool {
	var _arg0 *C.GtkListBox // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_show_separators(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Insert the @child into the @box at @position.
//
// If a sort function is set, the widget will actually be inserted at the
// calculated position.
//
// If @position is -1, or larger than the total number of items in the @box,
// then the @child will be appended to the end.
func (box *ListBox) Insert(child Widgeter, position int) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.int         // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = C.int(position)

	C.gtk_list_box_insert(_arg0, _arg1, _arg2)
}

// InvalidateFilter: update the filtering for all rows.
//
// Call this when result of the filter function on the @box is changed due to an
// external factor. For instance, this would be used if the filter function just
// looked for a specific search string and the entry with the search string has
// changed.
func (box *ListBox) InvalidateFilter() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_invalidate_filter(_arg0)
}

// InvalidateHeaders: update the separators for all rows.
//
// Call this when result of the header function on the @box is changed due to an
// external factor.
func (box *ListBox) InvalidateHeaders() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_invalidate_headers(_arg0)
}

// InvalidateSort: update the sorting for all rows.
//
// Call this when result of the sort function on the @box is changed due to an
// external factor.
func (box *ListBox) InvalidateSort() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_invalidate_sort(_arg0)
}

// Prepend a widget to the list.
//
// If a sort function is set, the widget will actually be inserted at the
// calculated position.
func (box *ListBox) Prepend(child Widgeter) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_list_box_prepend(_arg0, _arg1)
}

// Remove removes a child from @box.
func (box *ListBox) Remove(child Widgeter) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_list_box_remove(_arg0, _arg1)
}

// SelectAll: select all children of @box, if the selection mode allows it.
func (box *ListBox) SelectAll() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_select_all(_arg0)
}

// SelectRow: make @row the currently selected row.
func (box *ListBox) SelectRow(row ListBoxRower) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer((row).(gextras.Nativer).Native()))

	C.gtk_list_box_select_row(_arg0, _arg1)
}

// SelectedForeach calls a function for each selected child.
//
// Note that the selection cannot be modified from within this function.
func (box *ListBox) SelectedForeach(fn ListBoxForeachFunc) {
	var _arg0 *C.GtkListBox           // out
	var _arg1 C.GtkListBoxForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*[0]byte)(C.gotk4_ListBoxForeachFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))

	C.gtk_list_box_selected_foreach(_arg0, _arg1, _arg2)
}

// SetActivateOnSingleClick: if @single is true, rows will be activated when you
// click on them, otherwise you need to double-click.
func (box *ListBox) SetActivateOnSingleClick(single bool) {
	var _arg0 *C.GtkListBox // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	if single {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_set_activate_on_single_click(_arg0, _arg1)
}

// SetAdjustment sets the adjustment (if any) that the widget uses to for
// vertical scrolling.
//
// For instance, this is used to get the page size for PageUp/Down key handling.
//
// In the normal case when the @box is packed inside a `GtkScrolledWindow` the
// adjustment from that will be picked up automatically, so there is no need to
// manually do that.
func (box *ListBox) SetAdjustment(adjustment Adjustmenter) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer((adjustment).(gextras.Nativer).Native()))

	C.gtk_list_box_set_adjustment(_arg0, _arg1)
}

// SetPlaceholder sets the placeholder widget that is shown in the list when it
// doesn't display any visible children.
func (box *ListBox) SetPlaceholder(placeholder Widgeter) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((placeholder).(gextras.Nativer).Native()))

	C.gtk_list_box_set_placeholder(_arg0, _arg1)
}

// SetSelectionMode sets how selection works in the listbox.
func (box *ListBox) SetSelectionMode(mode SelectionMode) {
	var _arg0 *C.GtkListBox      // out
	var _arg1 C.GtkSelectionMode // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.GtkSelectionMode(mode)

	C.gtk_list_box_set_selection_mode(_arg0, _arg1)
}

// SetShowSeparators sets whether the list box should show separators between
// rows.
func (box *ListBox) SetShowSeparators(showSeparators bool) {
	var _arg0 *C.GtkListBox // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	if showSeparators {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_set_show_separators(_arg0, _arg1)
}

// UnselectAll: unselect all children of @box, if the selection mode allows it.
func (box *ListBox) UnselectAll() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_unselect_all(_arg0)
}

// UnselectRow unselects a single row of @box, if the selection mode allows it.
func (box *ListBox) UnselectRow(row ListBoxRower) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer((row).(gextras.Nativer).Native()))

	C.gtk_list_box_unselect_row(_arg0, _arg1)
}

// ListBoxRowOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ListBoxRowOverrider interface {
	Activate()
}

// ListBoxRower describes ListBoxRow's methods.
type ListBoxRower interface {
	// Changed marks @row as changed, causing any state that depends on this to
	// be updated.
	Changed()
	// Activatable gets whether the row is activatable.
	Activatable() bool
	// Child gets the child widget of @row.
	Child() *Widget
	// Header returns the current header of the @row.
	Header() *Widget
	// Index gets the current index of the @row in its `GtkListBox` container.
	Index() int
	// Selectable gets whether the row can be selected.
	Selectable() bool
	// IsSelected returns whether the child is currently selected in its
	// `GtkListBox` container.
	IsSelected() bool
	// SetActivatable: set whether the row is activatable.
	SetActivatable(activatable bool)
	// SetChild sets the child widget of @self.
	SetChild(child Widgeter)
	// SetHeader sets the current header of the @row.
	SetHeader(header Widgeter)
	// SetSelectable: set whether the row can be selected.
	SetSelectable(selectable bool)
}

// ListBoxRow: `GtkListBoxRow` is the kind of widget that can be added to a
// `GtkListBox`.
type ListBoxRow struct {
	Widget

	Actionable
}

var (
	_ ListBoxRower    = (*ListBoxRow)(nil)
	_ gextras.Nativer = (*ListBoxRow)(nil)
)

func wrapListBoxRow(obj *externglib.Object) *ListBoxRow {
	return &ListBoxRow{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Actionable: Actionable{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
		},
	}
}

func marshalListBoxRower(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapListBoxRow(obj), nil
}

// NewListBoxRow creates a new `GtkListBoxRow`.
func NewListBoxRow() *ListBoxRow {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_list_box_row_new()

	var _listBoxRow *ListBoxRow // out

	_listBoxRow = wrapListBoxRow(externglib.Take(unsafe.Pointer(_cret)))

	return _listBoxRow
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *ListBoxRow) Native() uintptr {
	return v.Widget.InitiallyUnowned.Object.Native()
}

// Changed marks @row as changed, causing any state that depends on this to be
// updated.
//
// This affects sorting, filtering and headers.
//
// Note that calls to this method must be in sync with the data used for the row
// functions. For instance, if the list is mirroring some external data set, and
// *two* rows changed in the external data set then when you call
// gtk_list_box_row_changed() on the first row the sort function must only read
// the new data for the first of the two changed rows, otherwise the resorting
// of the rows will be wrong.
//
// This generally means that if you don’t fully control the data model you have
// to duplicate the data that affects the listbox row functions into the row
// widgets themselves. Another alternative is to call
// [method@Gtk.ListBox.invalidate_sort] on any model change, but that is more
// expensive.
func (row *ListBoxRow) Changed() {
	var _arg0 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_row_changed(_arg0)
}

// Activatable gets whether the row is activatable.
func (row *ListBoxRow) Activatable() bool {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_activatable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Child gets the child widget of @row.
func (row *ListBoxRow) Child() *Widget {
	var _arg0 *C.GtkListBoxRow // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_child(_arg0)

	var _widget *Widget // out

	_widget = wrapWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _widget
}

// Header returns the current header of the @row.
//
// This can be used in a [callback@Gtk.ListBoxUpdateHeaderFunc] to see if there
// is a header set already, and if so to update the state of it.
func (row *ListBoxRow) Header() *Widget {
	var _arg0 *C.GtkListBoxRow // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_header(_arg0)

	var _widget *Widget // out

	_widget = wrapWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _widget
}

// Index gets the current index of the @row in its `GtkListBox` container.
func (row *ListBoxRow) Index() int {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.int            // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_index(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Selectable gets whether the row can be selected.
func (row *ListBoxRow) Selectable() bool {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_selectable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSelected returns whether the child is currently selected in its
// `GtkListBox` container.
func (row *ListBoxRow) IsSelected() bool {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_is_selected(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActivatable: set whether the row is activatable.
func (row *ListBoxRow) SetActivatable(activatable bool) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))
	if activatable {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_row_set_activatable(_arg0, _arg1)
}

// SetChild sets the child widget of @self.
func (row *ListBoxRow) SetChild(child Widgeter) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_list_box_row_set_child(_arg0, _arg1)
}

// SetHeader sets the current header of the @row.
//
// This is only allowed to be called from a
// [callback@Gtk.ListBoxUpdateHeaderFunc]. It will replace any existing header
// in the row, and be shown in front of the row in the listbox.
func (row *ListBoxRow) SetHeader(header Widgeter) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((header).(gextras.Nativer).Native()))

	C.gtk_list_box_row_set_header(_arg0, _arg1)
}

// SetSelectable: set whether the row can be selected.
func (row *ListBoxRow) SetSelectable(selectable bool) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))
	if selectable {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_row_set_selectable(_arg0, _arg1)
}
