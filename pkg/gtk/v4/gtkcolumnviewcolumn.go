// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_column_view_column_get_type()), F: marshalColumnViewColumn},
	})
}

// ColumnViewColumn: `GtkColumnViewColumn` represents the columns being added to
// `GtkColumnView`.
//
// The main ingredient for a `GtkColumnViewColumn` is the `GtkListItemFactory`
// that tells the columnview how to create cells for this column from items in
// the model.
//
// Columns have a title, and can optionally have a header menu set with
// [method@Gtk.ColumnViewColumn.set_header_menu].
//
// A sorter can be associated with a column using
// [method@Gtk.ColumnViewColumn.set_sorter], to let users influence sorting by
// clicking on the column header.
type ColumnViewColumn interface {
	gextras.Objector

	// ColumnView gets the column view that's currently displaying this column.
	//
	// If @self has not been added to a column view yet, nil is returned.
	ColumnView() ColumnView
	// Expand returns whether this column should expand.
	Expand() bool
	// Factory gets the factory that's currently used to populate list items for
	// this column.
	Factory() ListItemFactory
	// FixedWidth gets the fixed width of the column.
	FixedWidth() int
	// HeaderMenu gets the menu model that is used to create the context menu
	// for the column header.
	HeaderMenu() gio.MenuModel
	// Resizable returns whether this column is resizable.
	Resizable() bool
	// Sorter returns the sorter that is associated with the column.
	Sorter() Sorter
	// Title returns the title set with gtk_column_view_column_set_title().
	Title() string
	// Visible returns whether this column is visible.
	Visible() bool
	// SetExpand sets the column to take available extra space.
	//
	// The extra space is shared equally amongst all columns that have the
	// expand set to true.
	SetExpand(expand bool)
	// SetFactory sets the `GtkListItemFactory` to use for populating list items
	// for this column.
	SetFactory(factory ListItemFactory)
	// SetFixedWidth: if @fixed_width is not -1, sets the fixed width of
	// @column; otherwise unsets it.
	//
	// Setting a fixed width overrides the automatically calculated width.
	// Interactive resizing also sets the “fixed-width” property.
	SetFixedWidth(fixedWidth int)
	// SetHeaderMenu sets the menu model that is used to create the context menu
	// for the column header.
	SetHeaderMenu(menu gio.MenuModel)
	// SetResizable sets whether this column should be resizable by dragging.
	SetResizable(resizable bool)
	// SetSorter associates a sorter with the column.
	//
	// If @sorter is nil, the column will not let users change the sorting by
	// clicking on its header.
	//
	// This sorter can be made active by clicking on the column header, or by
	// calling [method@Gtk.ColumnView.sort_by_column].
	//
	// See [method@Gtk.ColumnView.get_sorter] for the necessary steps for
	// setting up customizable sorting for [class@Gtk.ColumnView].
	SetSorter(sorter Sorter)
	// SetTitle sets the title of this column.
	//
	// The title is displayed in the header of a `GtkColumnView` for this column
	// and is therefore user-facing text that should be translated.
	SetTitle(title string)
	// SetVisible sets whether this column should be visible in views.
	SetVisible(visible bool)
}

// columnViewColumn implements the ColumnViewColumn interface.
type columnViewColumn struct {
	gextras.Objector
}

var _ ColumnViewColumn = (*columnViewColumn)(nil)

// WrapColumnViewColumn wraps a GObject to the right type. It is
// primarily used internally.
func WrapColumnViewColumn(obj *externglib.Object) ColumnViewColumn {
	return ColumnViewColumn{
		Objector: obj,
	}
}

func marshalColumnViewColumn(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColumnViewColumn(obj), nil
}

// NewColumnViewColumn constructs a class ColumnViewColumn.
func NewColumnViewColumn(title string, factory ListItemFactory) ColumnViewColumn {
	var _arg1 *C.char
	var _arg2 *C.GtkListItemFactory

	_arg1 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))

	var _cret C.GtkColumnViewColumn

	cret = C.gtk_column_view_column_new(_arg1, _arg2)

	var _columnViewColumn ColumnViewColumn

	_columnViewColumn = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ColumnViewColumn)

	return _columnViewColumn
}

// ColumnView gets the column view that's currently displaying this column.
//
// If @self has not been added to a column view yet, nil is returned.
func (s columnViewColumn) ColumnView() ColumnView {
	var _arg0 *C.GtkColumnViewColumn

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkColumnView

	cret = C.gtk_column_view_column_get_column_view(_arg0)

	var _columnView ColumnView

	_columnView = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ColumnView)

	return _columnView
}

// Expand returns whether this column should expand.
func (s columnViewColumn) Expand() bool {
	var _arg0 *C.GtkColumnViewColumn

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	cret = C.gtk_column_view_column_get_expand(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Factory gets the factory that's currently used to populate list items for
// this column.
func (s columnViewColumn) Factory() ListItemFactory {
	var _arg0 *C.GtkColumnViewColumn

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkListItemFactory

	cret = C.gtk_column_view_column_get_factory(_arg0)

	var _listItemFactory ListItemFactory

	_listItemFactory = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ListItemFactory)

	return _listItemFactory
}

// FixedWidth gets the fixed width of the column.
func (s columnViewColumn) FixedWidth() int {
	var _arg0 *C.GtkColumnViewColumn

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var _cret C.int

	cret = C.gtk_column_view_column_get_fixed_width(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// HeaderMenu gets the menu model that is used to create the context menu
// for the column header.
func (s columnViewColumn) HeaderMenu() gio.MenuModel {
	var _arg0 *C.GtkColumnViewColumn

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var _cret *C.GMenuModel

	cret = C.gtk_column_view_column_get_header_menu(_arg0)

	var _menuModel gio.MenuModel

	_menuModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.MenuModel)

	return _menuModel
}

// Resizable returns whether this column is resizable.
func (s columnViewColumn) Resizable() bool {
	var _arg0 *C.GtkColumnViewColumn

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	cret = C.gtk_column_view_column_get_resizable(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Sorter returns the sorter that is associated with the column.
func (s columnViewColumn) Sorter() Sorter {
	var _arg0 *C.GtkColumnViewColumn

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkSorter

	cret = C.gtk_column_view_column_get_sorter(_arg0)

	var _sorter Sorter

	_sorter = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Sorter)

	return _sorter
}

// Title returns the title set with gtk_column_view_column_set_title().
func (s columnViewColumn) Title() string {
	var _arg0 *C.GtkColumnViewColumn

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var _cret *C.char

	cret = C.gtk_column_view_column_get_title(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Visible returns whether this column is visible.
func (s columnViewColumn) Visible() bool {
	var _arg0 *C.GtkColumnViewColumn

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	cret = C.gtk_column_view_column_get_visible(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetExpand sets the column to take available extra space.
//
// The extra space is shared equally amongst all columns that have the
// expand set to true.
func (s columnViewColumn) SetExpand(expand bool) {
	var _arg0 *C.GtkColumnViewColumn
	var _arg1 C.gboolean

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	if expand {
		_arg1 = C.gboolean(1)
	}

	C.gtk_column_view_column_set_expand(_arg0, _arg1)
}

// SetFactory sets the `GtkListItemFactory` to use for populating list items
// for this column.
func (s columnViewColumn) SetFactory(factory ListItemFactory) {
	var _arg0 *C.GtkColumnViewColumn
	var _arg1 *C.GtkListItemFactory

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))

	C.gtk_column_view_column_set_factory(_arg0, _arg1)
}

// SetFixedWidth: if @fixed_width is not -1, sets the fixed width of
// @column; otherwise unsets it.
//
// Setting a fixed width overrides the automatically calculated width.
// Interactive resizing also sets the “fixed-width” property.
func (s columnViewColumn) SetFixedWidth(fixedWidth int) {
	var _arg0 *C.GtkColumnViewColumn
	var _arg1 C.int

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(fixedWidth)

	C.gtk_column_view_column_set_fixed_width(_arg0, _arg1)
}

// SetHeaderMenu sets the menu model that is used to create the context menu
// for the column header.
func (s columnViewColumn) SetHeaderMenu(menu gio.MenuModel) {
	var _arg0 *C.GtkColumnViewColumn
	var _arg1 *C.GMenuModel

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(menu.Native()))

	C.gtk_column_view_column_set_header_menu(_arg0, _arg1)
}

// SetResizable sets whether this column should be resizable by dragging.
func (s columnViewColumn) SetResizable(resizable bool) {
	var _arg0 *C.GtkColumnViewColumn
	var _arg1 C.gboolean

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	if resizable {
		_arg1 = C.gboolean(1)
	}

	C.gtk_column_view_column_set_resizable(_arg0, _arg1)
}

// SetSorter associates a sorter with the column.
//
// If @sorter is nil, the column will not let users change the sorting by
// clicking on its header.
//
// This sorter can be made active by clicking on the column header, or by
// calling [method@Gtk.ColumnView.sort_by_column].
//
// See [method@Gtk.ColumnView.get_sorter] for the necessary steps for
// setting up customizable sorting for [class@Gtk.ColumnView].
func (s columnViewColumn) SetSorter(sorter Sorter) {
	var _arg0 *C.GtkColumnViewColumn
	var _arg1 *C.GtkSorter

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))

	C.gtk_column_view_column_set_sorter(_arg0, _arg1)
}

// SetTitle sets the title of this column.
//
// The title is displayed in the header of a `GtkColumnView` for this column
// and is therefore user-facing text that should be translated.
func (s columnViewColumn) SetTitle(title string) {
	var _arg0 *C.GtkColumnViewColumn
	var _arg1 *C.char

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_column_view_column_set_title(_arg0, _arg1)
}

// SetVisible sets whether this column should be visible in views.
func (s columnViewColumn) SetVisible(visible bool) {
	var _arg0 *C.GtkColumnViewColumn
	var _arg1 C.gboolean

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	if visible {
		_arg1 = C.gboolean(1)
	}

	C.gtk_column_view_column_set_visible(_arg0, _arg1)
}