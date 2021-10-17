// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_column_view_get_type()), F: marshalColumnViewer},
	})
}

// ColumnView: GtkColumnView presents a large dynamic list of items using
// multiple columns with headers.
//
// GtkColumnView uses the factories of its columns to generate a cell widget for
// each column, for each visible item and displays them together as the row for
// this item.
//
// The gtk.ColumnView:show-row-separators and
// [propertyGtk.ColumnView:show-column-separators] properties offer a simple way
// to display separators between the rows or columns.
//
// GtkColumnView allows the user to select items according to the selection
// characteristics of the model. For models that allow multiple selected items,
// it is possible to turn on *rubberband selection*, using
// gtk.ColumnView:enable-rubberband.
//
// The column view supports sorting that can be customized by the user by
// clicking on column headers. To set this up, the GtkSorter returned by
// gtk.ColumnView.GetSorter() must be attached to a sort model for the data that
// the view is showing, and the columns must have sorters attached to them by
// calling gtk.ColumnViewColumn.SetSorter(). The initial sort order can be set
// with gtk.ColumnView.SortByColumn().
//
// The column view also supports interactive resizing and reordering of columns,
// via Drag-and-Drop of the column headers. This can be enabled or disabled with
// the gtk.ColumnView:reorderable and gtk.ColumnViewColumn:resizable properties.
//
// To learn more about the list widget framework, see the overview
// (section-list-widget.html).
//
// CSS nodes
//
//    columnview[.column-separators][.rich-list][.navigation-sidebar][.data-table]
//    ├── header
//    │   ├── <column header>
//    ┊   ┊
//    │   ╰── <column header>
//    │
//    ├── listview
//    │
//    ┊
//    ╰── [rubberband]
//
//
// GtkColumnView uses a single CSS node named columnview. It may carry the
// .column-separators style class, when gtk.ColumnView:show-column-separators
// property is set. Header widgets appear below a node with name header. The
// rows are contained in a GtkListView widget, so there is a listview node with
// the same structure as for a standalone GtkListView widget. If
// gtk.ColumnView:show-row-separators is set, it will be passed on to the list
// view, causing its CSS node to carry the .separators style class. For
// rubberband selection, a node with name rubberband is used.
//
// The main columnview node may also carry style classes to select the style of
// list presentation (section-list-widget.html#list-styles): .rich-list,
// .navigation-sidebar or .data-table.
//
//
// Accessibility
//
// GtkColumnView uses the GTK_ACCESSIBLE_ROLE_TREE_GRID role, header title
// widgets are using the GTK_ACCESSIBLE_ROLE_COLUMN_HEADER role. The row widgets
// are using the GTK_ACCESSIBLE_ROLE_ROW role, and individual cells are using
// the GTK_ACCESSIBLE_ROLE_GRID_CELL role.
type ColumnView struct {
	Widget

	Scrollable
	*externglib.Object
}

func wrapColumnView(obj *externglib.Object) *ColumnView {
	return &ColumnView{
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
			Object: obj,
		},
		Scrollable: Scrollable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalColumnViewer(p uintptr) (interface{}, error) {
	return wrapColumnView(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewColumnView creates a new GtkColumnView.
//
// You most likely want to call gtk.ColumnView.AppendColumn() to add columns
// next.
//
// The function takes the following parameters:
//
//    - model: list model to use, or NULL.
//
func NewColumnView(model SelectionModeller) *ColumnView {
	var _arg1 *C.GtkSelectionModel // out
	var _cret *C.GtkWidget         // in

	if model != nil {
		_arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
		C.g_object_ref(C.gpointer(model.Native()))
	}

	_cret = C.gtk_column_view_new(_arg1)
	runtime.KeepAlive(model)

	var _columnView *ColumnView // out

	_columnView = wrapColumnView(externglib.Take(unsafe.Pointer(_cret)))

	return _columnView
}

// AppendColumn appends the column to the end of the columns in self.
//
// The function takes the following parameters:
//
//    - column: GtkColumnViewColumn that hasn't been added to a GtkColumnView
//    yet.
//
func (self *ColumnView) AppendColumn(column *ColumnViewColumn) {
	var _arg0 *C.GtkColumnView       // out
	var _arg1 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))

	C.gtk_column_view_append_column(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(column)
}

// Columns gets the list of columns in this column view.
//
// This list is constant over the lifetime of self and can be used to monitor
// changes to the columns of self by connecting to the ::items-changed signal.
func (self *ColumnView) Columns() gio.ListModeller {
	var _arg0 *C.GtkColumnView // out
	var _cret *C.GListModel    // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_column_view_get_columns(_arg0)
	runtime.KeepAlive(self)

	var _listModel gio.ListModeller // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.ListModeller is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(gio.ListModeller)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.ListModeller")
		}
		_listModel = rv
	}

	return _listModel
}

// EnableRubberband returns whether rows can be selected by dragging with the
// mouse.
func (self *ColumnView) EnableRubberband() bool {
	var _arg0 *C.GtkColumnView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_column_view_get_enable_rubberband(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Model gets the model that's currently used to read the items displayed.
func (self *ColumnView) Model() SelectionModeller {
	var _arg0 *C.GtkColumnView     // out
	var _cret *C.GtkSelectionModel // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_column_view_get_model(_arg0)
	runtime.KeepAlive(self)

	var _selectionModel SelectionModeller // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(SelectionModeller)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.SelectionModeller")
			}
			_selectionModel = rv
		}
	}

	return _selectionModel
}

// Reorderable returns whether columns are reorderable.
func (self *ColumnView) Reorderable() bool {
	var _arg0 *C.GtkColumnView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_column_view_get_reorderable(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowColumnSeparators returns whether the list should show separators between
// columns.
func (self *ColumnView) ShowColumnSeparators() bool {
	var _arg0 *C.GtkColumnView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_column_view_get_show_column_separators(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowRowSeparators returns whether the list should show separators between
// rows.
func (self *ColumnView) ShowRowSeparators() bool {
	var _arg0 *C.GtkColumnView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_column_view_get_show_row_separators(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SingleClickActivate returns whether rows will be activated on single click
// and selected on hover.
func (self *ColumnView) SingleClickActivate() bool {
	var _arg0 *C.GtkColumnView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_column_view_get_single_click_activate(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Sorter returns a special sorter that reflects the users sorting choices in
// the column view.
//
// To allow users to customizable sorting by clicking on column headers, this
// sorter needs to be set on the sort model underneath the model that is
// displayed by the view.
//
// See gtk.ColumnViewColumn.SetSorter() for setting up per-column sorting.
//
// Here is an example:
//
//    gtk_column_view_column_set_sorter (column, sorter);
//    gtk_column_view_append_column (view, column);
//    sorter = g_object_ref (gtk_column_view_get_sorter (view)));
//    model = gtk_sort_list_model_new (store, sorter);
//    selection = gtk_no_selection_new (model);
//    gtk_column_view_set_model (view, selection);.
func (self *ColumnView) Sorter() *Sorter {
	var _arg0 *C.GtkColumnView // out
	var _cret *C.GtkSorter     // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_column_view_get_sorter(_arg0)
	runtime.KeepAlive(self)

	var _sorter *Sorter // out

	if _cret != nil {
		_sorter = wrapSorter(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _sorter
}

// InsertColumn inserts a column at the given position in the columns of self.
//
// If column is already a column of self, it will be repositioned.
//
// The function takes the following parameters:
//
//    - position to insert column at.
//    - column: GtkColumnViewColumn to insert.
//
func (self *ColumnView) InsertColumn(position uint, column *ColumnViewColumn) {
	var _arg0 *C.GtkColumnView       // out
	var _arg1 C.guint                // out
	var _arg2 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(position)
	_arg2 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))

	C.gtk_column_view_insert_column(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)
	runtime.KeepAlive(column)
}

// RemoveColumn removes the column from the list of columns of self.
//
// The function takes the following parameters:
//
//    - column: GtkColumnViewColumn that's part of self.
//
func (self *ColumnView) RemoveColumn(column *ColumnViewColumn) {
	var _arg0 *C.GtkColumnView       // out
	var _arg1 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))

	C.gtk_column_view_remove_column(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(column)
}

// SetEnableRubberband sets whether selections can be changed by dragging with
// the mouse.
//
// The function takes the following parameters:
//
//    - enableRubberband: TRUE to enable rubberband selection.
//
func (self *ColumnView) SetEnableRubberband(enableRubberband bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	if enableRubberband {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_enable_rubberband(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(enableRubberband)
}

// SetModel sets the model to use.
//
// This must be a gtk.SelectionModel.
//
// The function takes the following parameters:
//
//    - model to use or NULL for none.
//
func (self *ColumnView) SetModel(model SelectionModeller) {
	var _arg0 *C.GtkColumnView     // out
	var _arg1 *C.GtkSelectionModel // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	if model != nil {
		_arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_column_view_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}

// SetReorderable sets whether columns should be reorderable by dragging.
//
// The function takes the following parameters:
//
//    - reorderable: whether columns should be reorderable.
//
func (self *ColumnView) SetReorderable(reorderable bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	if reorderable {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_reorderable(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(reorderable)
}

// SetShowColumnSeparators sets whether the list should show separators between
// columns.
//
// The function takes the following parameters:
//
//    - showColumnSeparators: TRUE to show column separators.
//
func (self *ColumnView) SetShowColumnSeparators(showColumnSeparators bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	if showColumnSeparators {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_show_column_separators(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(showColumnSeparators)
}

// SetShowRowSeparators sets whether the list should show separators between
// rows.
//
// The function takes the following parameters:
//
//    - showRowSeparators: TRUE to show row separators.
//
func (self *ColumnView) SetShowRowSeparators(showRowSeparators bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	if showRowSeparators {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_show_row_separators(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(showRowSeparators)
}

// SetSingleClickActivate sets whether rows should be activated on single click
// and selected on hover.
//
// The function takes the following parameters:
//
//    - singleClickActivate: TRUE to activate items on single click.
//
func (self *ColumnView) SetSingleClickActivate(singleClickActivate bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	if singleClickActivate {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_single_click_activate(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(singleClickActivate)
}

// SortByColumn sets the sorting of the view.
//
// This function should be used to set up the initial sorting. At runtime, users
// can change the sorting of a column view by clicking on the list headers.
//
// This call only has an effect if the sorter returned by
// gtk.ColumnView.GetSorter() is set on a sort model, and
// gtk.ColumnViewColumn.SetSorter() has been called on column to associate a
// sorter with the column.
//
// If column is NULL, the view will be unsorted.
//
// The function takes the following parameters:
//
//    - column: GtkColumnViewColumn to sort by, or NULL.
//    - direction to sort in.
//
func (self *ColumnView) SortByColumn(column *ColumnViewColumn, direction SortType) {
	var _arg0 *C.GtkColumnView       // out
	var _arg1 *C.GtkColumnViewColumn // out
	var _arg2 C.GtkSortType          // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	if column != nil {
		_arg1 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))
	}
	_arg2 = C.GtkSortType(direction)

	C.gtk_column_view_sort_by_column(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(column)
	runtime.KeepAlive(direction)
}

// ConnectActivate: emitted when a row has been activated by the user, usually
// via activating the GtkListBase|list.activate-item action.
//
// This allows for a convenient way to handle activation in a columnview. See
// gtk.ListItem.SetActivatable() for details on how to use this signal.
func (self *ColumnView) ConnectActivate(f func(position uint)) externglib.SignalHandle {
	return self.Connect("activate", f)
}
