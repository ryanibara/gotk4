// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// GtkWidget* _gotk4_gtk3_ListBoxCreateWidgetFunc(gpointer, gpointer);
// extern void callbackDelete(gpointer);
// gboolean _gotk4_gtk3_ListBoxFilterFunc(GtkListBoxRow*, gpointer);
// gint _gotk4_gtk3_ListBoxSortFunc(GtkListBoxRow*, GtkListBoxRow*, gpointer);
// void _gotk4_gtk3_ListBoxForeachFunc(GtkListBox*, GtkListBoxRow*, gpointer);
// void _gotk4_gtk3_ListBoxUpdateHeaderFunc(GtkListBoxRow*, GtkListBoxRow*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_list_box_get_type()), F: marshalListBoxer},
		{T: externglib.Type(C.gtk_list_box_row_get_type()), F: marshalListBoxRower},
	})
}

// ListBoxCreateWidgetFunc: called for list boxes that are bound to a Model with
// gtk_list_box_bind_model() for each item that gets added to the model.
//
// Versions of GTK+ prior to 3.18 called gtk_widget_show_all() on the rows
// created by the GtkListBoxCreateWidgetFunc, but this forced all widgets inside
// the row to be shown, and is no longer the case. Applications should be
// updated to show the desired row widgets.
type ListBoxCreateWidgetFunc func(item *externglib.Object) (widget Widgetter)

//export _gotk4_gtk3_ListBoxCreateWidgetFunc
func _gotk4_gtk3_ListBoxCreateWidgetFunc(arg0 C.gpointer, arg1 C.gpointer) (cret *C.GtkWidget) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var item *externglib.Object // out

	item = externglib.Take(unsafe.Pointer(arg0))

	fn := v.(ListBoxCreateWidgetFunc)
	widget := fn(item)

	cret = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	return cret
}

// ListBoxFilterFunc will be called whenever the row changes or is added and
// lets you control if the row should be visible or not.
type ListBoxFilterFunc func(row *ListBoxRow) (ok bool)

//export _gotk4_gtk3_ListBoxFilterFunc
func _gotk4_gtk3_ListBoxFilterFunc(arg0 *C.GtkListBoxRow, arg1 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var row *ListBoxRow // out

	row = wrapListBoxRow(externglib.Take(unsafe.Pointer(arg0)))

	fn := v.(ListBoxFilterFunc)
	ok := fn(row)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ListBoxForeachFunc: function used by gtk_list_box_selected_foreach(). It will
// be called on every selected child of the box.
type ListBoxForeachFunc func(box *ListBox, row *ListBoxRow)

//export _gotk4_gtk3_ListBoxForeachFunc
func _gotk4_gtk3_ListBoxForeachFunc(arg0 *C.GtkListBox, arg1 *C.GtkListBoxRow, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var box *ListBox    // out
	var row *ListBoxRow // out

	box = wrapListBox(externglib.Take(unsafe.Pointer(arg0)))
	row = wrapListBoxRow(externglib.Take(unsafe.Pointer(arg1)))

	fn := v.(ListBoxForeachFunc)
	fn(box, row)
}

// ListBoxSortFunc: compare two rows to determine which should be first.
type ListBoxSortFunc func(row1 *ListBoxRow, row2 *ListBoxRow) (gint int)

//export _gotk4_gtk3_ListBoxSortFunc
func _gotk4_gtk3_ListBoxSortFunc(arg0 *C.GtkListBoxRow, arg1 *C.GtkListBoxRow, arg2 C.gpointer) (cret C.gint) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var row1 *ListBoxRow // out
	var row2 *ListBoxRow // out

	row1 = wrapListBoxRow(externglib.Take(unsafe.Pointer(arg0)))
	row2 = wrapListBoxRow(externglib.Take(unsafe.Pointer(arg1)))

	fn := v.(ListBoxSortFunc)
	gint := fn(row1, row2)

	cret = C.gint(gint)

	return cret
}

// ListBoxUpdateHeaderFunc: whenever row changes or which row is before row
// changes this is called, which lets you update the header on row. You may
// remove or set a new one via gtk_list_box_row_set_header() or just change the
// state of the current header widget.
type ListBoxUpdateHeaderFunc func(row *ListBoxRow, before *ListBoxRow)

//export _gotk4_gtk3_ListBoxUpdateHeaderFunc
func _gotk4_gtk3_ListBoxUpdateHeaderFunc(arg0 *C.GtkListBoxRow, arg1 *C.GtkListBoxRow, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var row *ListBoxRow    // out
	var before *ListBoxRow // out

	row = wrapListBoxRow(externglib.Take(unsafe.Pointer(arg0)))
	if arg1 != nil {
		before = wrapListBoxRow(externglib.Take(unsafe.Pointer(arg1)))
	}

	fn := v.(ListBoxUpdateHeaderFunc)
	fn(row, before)
}

// ListBoxOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ListBoxOverrider interface {
	ActivateCursorRow()
	MoveCursor(step MovementStep, count int)
	RowActivated(row *ListBoxRow)
	RowSelected(row *ListBoxRow)
	// SelectAll: select all children of box, if the selection mode allows it.
	SelectAll()
	SelectedRowsChanged()
	ToggleCursorRow()
	// UnselectAll: unselect all children of box, if the selection mode allows
	// it.
	UnselectAll()
}

// ListBox is a vertical container that contains GtkListBoxRow children. These
// rows can by dynamically sorted and filtered, and headers can be added
// dynamically depending on the row content. It also allows keyboard and mouse
// navigation and selection like a typical list.
//
// Using GtkListBox is often an alternative to TreeView, especially when the
// list contents has a more complicated layout than what is allowed by a
// CellRenderer, or when the contents is interactive (i.e. has a button in it).
//
// Although a ListBox must have only ListBoxRow children you can add any kind of
// widget to it via gtk_container_add(), and a ListBoxRow widget will
// automatically be inserted between the list and the widget.
//
// ListBoxRows can be marked as activatable or selectable. If a row is
// activatable, ListBox::row-activated will be emitted for it when the user
// tries to activate it. If it is selectable, the row will be marked as selected
// when the user tries to select it.
//
// The GtkListBox widget was added in GTK+ 3.10.
//
//
// GtkListBox as GtkBuildable
//
// The GtkListBox implementation of the Buildable interface supports setting a
// child as the placeholder by specifying “placeholder” as the “type” attribute
// of a <child> element. See gtk_list_box_set_placeholder() for info.
//
// CSS nodes
//
//    list
//    ╰── row[.activatable]
//
// GtkListBox uses a single CSS node named list. Each GtkListBoxRow uses a
// single CSS node named row. The row nodes get the .activatable style class
// added when appropriate.
type ListBox struct {
	Container
}

func wrapListBox(obj *externglib.Object) *ListBox {
	return &ListBox{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
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

// NewListBox creates a new ListBox container.
func NewListBox() *ListBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_list_box_new()

	var _listBox *ListBox // out

	_listBox = wrapListBox(externglib.Take(unsafe.Pointer(_cret)))

	return _listBox
}

// BindModel binds model to box.
//
// If box was already bound to a model, that previous binding is destroyed.
//
// The contents of box are cleared and then filled with widgets that represent
// items from model. box is updated whenever model changes. If model is NULL,
// box is left empty.
//
// It is undefined to add or remove widgets directly (for example, with
// gtk_list_box_insert() or gtk_container_add()) while box is bound to a model.
//
// Note that using a model is incompatible with the filtering and sorting
// functionality in GtkListBox. When using a model, filtering and sorting should
// be implemented by the model.
func (box *ListBox) BindModel(model gio.ListModeller, createWidgetFunc ListBoxCreateWidgetFunc) {
	var _arg0 *C.GtkListBox                // out
	var _arg1 *C.GListModel                // out
	var _arg2 C.GtkListBoxCreateWidgetFunc // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
	}
	if createWidgetFunc != nil {
		_arg2 = (*[0]byte)(C._gotk4_gtk3_ListBoxCreateWidgetFunc)
		_arg3 = C.gpointer(gbox.Assign(createWidgetFunc))
		_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_list_box_bind_model(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(box)
	runtime.KeepAlive(model)
	runtime.KeepAlive(createWidgetFunc)
}

// DragHighlightRow: this is a helper function for implementing DnD onto a
// ListBox. The passed in row will be highlighted via gtk_drag_highlight(), and
// any previously highlighted row will be unhighlighted.
//
// The row will also be unhighlighted when the widget gets a drag leave event.
func (box *ListBox) DragHighlightRow(row *ListBoxRow) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_drag_highlight_row(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(row)
}

// DragUnhighlightRow: if a row has previously been highlighted via
// gtk_list_box_drag_highlight_row() it will have the highlight removed.
func (box *ListBox) DragUnhighlightRow() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_drag_unhighlight_row(_arg0)
	runtime.KeepAlive(box)
}

// ActivateOnSingleClick returns whether rows activate on single clicks.
func (box *ListBox) ActivateOnSingleClick() bool {
	var _arg0 *C.GtkListBox // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_activate_on_single_click(_arg0)
	runtime.KeepAlive(box)

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
	runtime.KeepAlive(box)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// RowAtIndex gets the n-th child in the list (not counting headers). If _index
// is negative or larger than the number of items in the list, NULL is returned.
func (box *ListBox) RowAtIndex(index_ int) *ListBoxRow {
	var _arg0 *C.GtkListBox    // out
	var _arg1 C.gint           // out
	var _cret *C.GtkListBoxRow // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.gint(index_)

	_cret = C.gtk_list_box_get_row_at_index(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(index_)

	var _listBoxRow *ListBoxRow // out

	if _cret != nil {
		_listBoxRow = wrapListBoxRow(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _listBoxRow
}

// RowAtY gets the row at the y position.
func (box *ListBox) RowAtY(y int) *ListBoxRow {
	var _arg0 *C.GtkListBox    // out
	var _arg1 C.gint           // out
	var _cret *C.GtkListBoxRow // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.gint(y)

	_cret = C.gtk_list_box_get_row_at_y(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(y)

	var _listBoxRow *ListBoxRow // out

	if _cret != nil {
		_listBoxRow = wrapListBoxRow(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _listBoxRow
}

// SelectedRow gets the selected row.
//
// Note that the box may allow multiple selection, in which case you should use
// gtk_list_box_selected_foreach() to find all selected rows.
func (box *ListBox) SelectedRow() *ListBoxRow {
	var _arg0 *C.GtkListBox    // out
	var _cret *C.GtkListBoxRow // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_selected_row(_arg0)
	runtime.KeepAlive(box)

	var _listBoxRow *ListBoxRow // out

	_listBoxRow = wrapListBoxRow(externglib.Take(unsafe.Pointer(_cret)))

	return _listBoxRow
}

// SelectedRows creates a list of all selected children.
func (box *ListBox) SelectedRows() []ListBoxRow {
	var _arg0 *C.GtkListBox // out
	var _cret *C.GList      // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_selected_rows(_arg0)
	runtime.KeepAlive(box)

	var _list []ListBoxRow // out

	_list = make([]ListBoxRow, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkListBoxRow)(v)
		var dst ListBoxRow // out
		dst = *wrapListBoxRow(externglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// SelectionMode gets the selection mode of the listbox.
func (box *ListBox) SelectionMode() SelectionMode {
	var _arg0 *C.GtkListBox      // out
	var _cret C.GtkSelectionMode // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_selection_mode(_arg0)
	runtime.KeepAlive(box)

	var _selectionMode SelectionMode // out

	_selectionMode = SelectionMode(_cret)

	return _selectionMode
}

// Insert the child into the box at position. If a sort function is set, the
// widget will actually be inserted at the calculated position and this function
// has the same effect of gtk_container_add().
//
// If position is -1, or larger than the total number of items in the box, then
// the child will be appended to the end.
func (box *ListBox) Insert(child Widgetter, position int) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gint        // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.gint(position)

	C.gtk_list_box_insert(_arg0, _arg1, _arg2)
	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
	runtime.KeepAlive(position)
}

// InvalidateFilter: update the filtering for all rows. Call this when result of
// the filter function on the box is changed due to an external factor. For
// instance, this would be used if the filter function just looked for a
// specific search string and the entry with the search string has changed.
func (box *ListBox) InvalidateFilter() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_invalidate_filter(_arg0)
	runtime.KeepAlive(box)
}

// InvalidateHeaders: update the separators for all rows. Call this when result
// of the header function on the box is changed due to an external factor.
func (box *ListBox) InvalidateHeaders() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_invalidate_headers(_arg0)
	runtime.KeepAlive(box)
}

// InvalidateSort: update the sorting for all rows. Call this when result of the
// sort function on the box is changed due to an external factor.
func (box *ListBox) InvalidateSort() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_invalidate_sort(_arg0)
	runtime.KeepAlive(box)
}

// Prepend a widget to the list. If a sort function is set, the widget will
// actually be inserted at the calculated position and this function has the
// same effect of gtk_container_add().
func (box *ListBox) Prepend(child Widgetter) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_list_box_prepend(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
}

// SelectAll: select all children of box, if the selection mode allows it.
func (box *ListBox) SelectAll() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_select_all(_arg0)
	runtime.KeepAlive(box)
}

// SelectRow: make row the currently selected row.
func (box *ListBox) SelectRow(row *ListBoxRow) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	if row != nil {
		_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))
	}

	C.gtk_list_box_select_row(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(row)
}

// SelectedForeach calls a function for each selected child.
//
// Note that the selection cannot be modified from within this function.
func (box *ListBox) SelectedForeach(fn ListBoxForeachFunc) {
	var _arg0 *C.GtkListBox           // out
	var _arg1 C.GtkListBoxForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_ListBoxForeachFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))
	defer gbox.Delete(uintptr(_arg2))

	C.gtk_list_box_selected_foreach(_arg0, _arg1, _arg2)
	runtime.KeepAlive(box)
	runtime.KeepAlive(fn)
}

// SetActivateOnSingleClick: if single is TRUE, rows will be activated when you
// click on them, otherwise you need to double-click.
func (box *ListBox) SetActivateOnSingleClick(single bool) {
	var _arg0 *C.GtkListBox // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	if single {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_set_activate_on_single_click(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(single)
}

// SetAdjustment sets the adjustment (if any) that the widget uses to for
// vertical scrolling. For instance, this is used to get the page size for
// PageUp/Down key handling.
//
// In the normal case when the box is packed inside a ScrolledWindow the
// adjustment from that will be picked up automatically, so there is no need to
// manually do that.
func (box *ListBox) SetAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	if adjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	}

	C.gtk_list_box_set_adjustment(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(adjustment)
}

// SetFilterFunc: by setting a filter function on the box one can decide
// dynamically which of the rows to show. For instance, to implement a search
// function on a list that filters the original list to only show the matching
// rows.
//
// The filter_func will be called for each row after the call, and it will
// continue to be called each time a row changes (via
// gtk_list_box_row_changed()) or when gtk_list_box_invalidate_filter() is
// called.
//
// Note that using a filter function is incompatible with using a model (see
// gtk_list_box_bind_model()).
func (box *ListBox) SetFilterFunc(filterFunc ListBoxFilterFunc) {
	var _arg0 *C.GtkListBox          // out
	var _arg1 C.GtkListBoxFilterFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	if filterFunc != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk3_ListBoxFilterFunc)
		_arg2 = C.gpointer(gbox.Assign(filterFunc))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_list_box_set_filter_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(box)
	runtime.KeepAlive(filterFunc)
}

// SetHeaderFunc: by setting a header function on the box one can dynamically
// add headers in front of rows, depending on the contents of the row and its
// position in the list. For instance, one could use it to add headers in front
// of the first item of a new kind, in a list sorted by the kind.
//
// The update_header can look at the current header widget using
// gtk_list_box_row_get_header() and either update the state of the widget as
// needed, or set a new one using gtk_list_box_row_set_header(). If no header is
// needed, set the header to NULL.
//
// Note that you may get many calls update_header to this for a particular row
// when e.g. changing things that don’t affect the header. In this case it is
// important for performance to not blindly replace an existing header with an
// identical one.
//
// The update_header function will be called for each row after the call, and it
// will continue to be called each time a row changes (via
// gtk_list_box_row_changed()) and when the row before changes (either by
// gtk_list_box_row_changed() on the previous row, or when the previous row
// becomes a different row). It is also called for all rows when
// gtk_list_box_invalidate_headers() is called.
func (box *ListBox) SetHeaderFunc(updateHeader ListBoxUpdateHeaderFunc) {
	var _arg0 *C.GtkListBox                // out
	var _arg1 C.GtkListBoxUpdateHeaderFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	if updateHeader != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk3_ListBoxUpdateHeaderFunc)
		_arg2 = C.gpointer(gbox.Assign(updateHeader))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_list_box_set_header_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(box)
	runtime.KeepAlive(updateHeader)
}

// SetPlaceholder sets the placeholder widget that is shown in the list when it
// doesn't display any visible children.
func (box *ListBox) SetPlaceholder(placeholder Widgetter) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	if placeholder != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(placeholder.Native()))
	}

	C.gtk_list_box_set_placeholder(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(placeholder)
}

// SetSelectionMode sets how selection works in the listbox. See SelectionMode
// for details.
func (box *ListBox) SetSelectionMode(mode SelectionMode) {
	var _arg0 *C.GtkListBox      // out
	var _arg1 C.GtkSelectionMode // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.GtkSelectionMode(mode)

	C.gtk_list_box_set_selection_mode(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(mode)
}

// SetSortFunc: by setting a sort function on the box one can dynamically
// reorder the rows of the list, based on the contents of the rows.
//
// The sort_func will be called for each row after the call, and will continue
// to be called each time a row changes (via gtk_list_box_row_changed()) and
// when gtk_list_box_invalidate_sort() is called.
//
// Note that using a sort function is incompatible with using a model (see
// gtk_list_box_bind_model()).
func (box *ListBox) SetSortFunc(sortFunc ListBoxSortFunc) {
	var _arg0 *C.GtkListBox        // out
	var _arg1 C.GtkListBoxSortFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	if sortFunc != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk3_ListBoxSortFunc)
		_arg2 = C.gpointer(gbox.Assign(sortFunc))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_list_box_set_sort_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(box)
	runtime.KeepAlive(sortFunc)
}

// UnselectAll: unselect all children of box, if the selection mode allows it.
func (box *ListBox) UnselectAll() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_unselect_all(_arg0)
	runtime.KeepAlive(box)
}

// UnselectRow unselects a single row of box, if the selection mode allows it.
func (box *ListBox) UnselectRow(row *ListBoxRow) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_unselect_row(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(row)
}

// ListBoxRowOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ListBoxRowOverrider interface {
	Activate()
}

type ListBoxRow struct {
	Bin

	Actionable
	*externglib.Object
}

func wrapListBoxRow(obj *externglib.Object) *ListBoxRow {
	return &ListBoxRow{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					Object: obj,
				},
			},
		},
		Actionable: Actionable{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				Object: obj,
			},
		},
		Object: obj,
	}
}

func marshalListBoxRower(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapListBoxRow(obj), nil
}

// NewListBoxRow creates a new ListBoxRow, to be used as a child of a ListBox.
func NewListBoxRow() *ListBoxRow {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_list_box_row_new()

	var _listBoxRow *ListBoxRow // out

	_listBoxRow = wrapListBoxRow(externglib.Take(unsafe.Pointer(_cret)))

	return _listBoxRow
}

// Changed marks row as changed, causing any state that depends on this to be
// updated. This affects sorting, filtering and headers.
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
// gtk_list_box_invalidate_sort() on any model change, but that is more
// expensive.
func (row *ListBoxRow) Changed() {
	var _arg0 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_row_changed(_arg0)
	runtime.KeepAlive(row)
}

// Activatable gets the value of the ListBoxRow:activatable property for this
// row.
func (row *ListBoxRow) Activatable() bool {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_activatable(_arg0)
	runtime.KeepAlive(row)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Header returns the current header of the row. This can be used in a
// ListBoxUpdateHeaderFunc to see if there is a header set already, and if so to
// update the state of it.
func (row *ListBoxRow) Header() Widgetter {
	var _arg0 *C.GtkListBoxRow // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_header(_arg0)
	runtime.KeepAlive(row)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Index gets the current index of the row in its ListBox container.
func (row *ListBoxRow) Index() int {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gint           // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_index(_arg0)
	runtime.KeepAlive(row)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Selectable gets the value of the ListBoxRow:selectable property for this row.
func (row *ListBoxRow) Selectable() bool {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_selectable(_arg0)
	runtime.KeepAlive(row)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSelected returns whether the child is currently selected in its ListBox
// container.
func (row *ListBoxRow) IsSelected() bool {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_is_selected(_arg0)
	runtime.KeepAlive(row)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActivatable: set the ListBoxRow:activatable property for this row.
func (row *ListBoxRow) SetActivatable(activatable bool) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))
	if activatable {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_row_set_activatable(_arg0, _arg1)
	runtime.KeepAlive(row)
	runtime.KeepAlive(activatable)
}

// SetHeader sets the current header of the row. This is only allowed to be
// called from a ListBoxUpdateHeaderFunc. It will replace any existing header in
// the row, and be shown in front of the row in the listbox.
func (row *ListBoxRow) SetHeader(header Widgetter) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))
	if header != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(header.Native()))
	}

	C.gtk_list_box_row_set_header(_arg0, _arg1)
	runtime.KeepAlive(row)
	runtime.KeepAlive(header)
}

// SetSelectable: set the ListBoxRow:selectable property for this row.
func (row *ListBoxRow) SetSelectable(selectable bool) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))
	if selectable {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_row_set_selectable(_arg0, _arg1)
	runtime.KeepAlive(row)
	runtime.KeepAlive(selectable)
}
