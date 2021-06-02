// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
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
	var arg1 *C.char
	var arg2 *C.GtkListItemFactory

	arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GtkListItemFactory)(factory.Native())

	ret := C.gtk_column_view_column_new(arg1, arg2)

	var ret0 ColumnViewColumn

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(ColumnViewColumn)

	return ret0
}

// ColumnView gets the column view that's currently displaying this column.
//
// If @self has not been added to a column view yet, nil is returned.
func (s columnViewColumn) ColumnView() ColumnView {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(s.Native())

	ret := C.gtk_column_view_column_get_column_view(arg0)

	var ret0 ColumnView

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(ColumnView)

	return ret0
}

// Expand returns whether this column should expand.
func (s columnViewColumn) Expand() bool {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(s.Native())

	ret := C.gtk_column_view_column_get_expand(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != C.false

	return ret0
}

// Factory gets the factory that's currently used to populate list items for
// this column.
func (s columnViewColumn) Factory() ListItemFactory {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(s.Native())

	ret := C.gtk_column_view_column_get_factory(arg0)

	var ret0 ListItemFactory

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(ListItemFactory)

	return ret0
}

// FixedWidth gets the fixed width of the column.
func (s columnViewColumn) FixedWidth() int {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(s.Native())

	ret := C.gtk_column_view_column_get_fixed_width(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// HeaderMenu gets the menu model that is used to create the context menu
// for the column header.
func (s columnViewColumn) HeaderMenu() gio.MenuModel {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(s.Native())

	ret := C.gtk_column_view_column_get_header_menu(arg0)

	var ret0 gio.MenuModel

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(gio.MenuModel)

	return ret0
}

// Resizable returns whether this column is resizable.
func (s columnViewColumn) Resizable() bool {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(s.Native())

	ret := C.gtk_column_view_column_get_resizable(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != C.false

	return ret0
}

// Sorter returns the sorter that is associated with the column.
func (s columnViewColumn) Sorter() Sorter {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(s.Native())

	ret := C.gtk_column_view_column_get_sorter(arg0)

	var ret0 Sorter

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Sorter)

	return ret0
}

// Title returns the title set with gtk_column_view_column_set_title().
func (s columnViewColumn) Title() string {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(s.Native())

	ret := C.gtk_column_view_column_get_title(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// Visible returns whether this column is visible.
func (s columnViewColumn) Visible() bool {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(s.Native())

	ret := C.gtk_column_view_column_get_visible(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != C.false

	return ret0
}

// SetExpand sets the column to take available extra space.
//
// The extra space is shared equally amongst all columns that have the
// expand set to true.
func (s columnViewColumn) SetExpand(expand bool) {
	var arg0 *C.GtkColumnViewColumn
	var arg1 C.gboolean

	arg0 = (*C.GtkColumnViewColumn)(s.Native())
	if expand {
		arg1 = C.TRUE
	}

	C.gtk_column_view_column_set_expand(arg0, arg1)
}

// SetFactory sets the `GtkListItemFactory` to use for populating list items
// for this column.
func (s columnViewColumn) SetFactory(factory ListItemFactory) {
	var arg0 *C.GtkColumnViewColumn
	var arg1 *C.GtkListItemFactory

	arg0 = (*C.GtkColumnViewColumn)(s.Native())
	arg1 = (*C.GtkListItemFactory)(factory.Native())

	C.gtk_column_view_column_set_factory(arg0, arg1)
}

// SetFixedWidth: if @fixed_width is not -1, sets the fixed width of
// @column; otherwise unsets it.
//
// Setting a fixed width overrides the automatically calculated width.
// Interactive resizing also sets the “fixed-width” property.
func (s columnViewColumn) SetFixedWidth(fixedWidth int) {
	var arg0 *C.GtkColumnViewColumn
	var arg1 C.int

	arg0 = (*C.GtkColumnViewColumn)(s.Native())
	arg1 = C.int(fixedWidth)

	C.gtk_column_view_column_set_fixed_width(arg0, arg1)
}

// SetHeaderMenu sets the menu model that is used to create the context menu
// for the column header.
func (s columnViewColumn) SetHeaderMenu(menu gio.MenuModel) {
	var arg0 *C.GtkColumnViewColumn
	var arg1 *C.GMenuModel

	arg0 = (*C.GtkColumnViewColumn)(s.Native())
	arg1 = (*C.GMenuModel)(menu.Native())

	C.gtk_column_view_column_set_header_menu(arg0, arg1)
}

// SetResizable sets whether this column should be resizable by dragging.
func (s columnViewColumn) SetResizable(resizable bool) {
	var arg0 *C.GtkColumnViewColumn
	var arg1 C.gboolean

	arg0 = (*C.GtkColumnViewColumn)(s.Native())
	if resizable {
		arg1 = C.TRUE
	}

	C.gtk_column_view_column_set_resizable(arg0, arg1)
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
	var arg0 *C.GtkColumnViewColumn
	var arg1 *C.GtkSorter

	arg0 = (*C.GtkColumnViewColumn)(s.Native())
	arg1 = (*C.GtkSorter)(sorter.Native())

	C.gtk_column_view_column_set_sorter(arg0, arg1)
}

// SetTitle sets the title of this column.
//
// The title is displayed in the header of a `GtkColumnView` for this column
// and is therefore user-facing text that should be translated.
func (s columnViewColumn) SetTitle(title string) {
	var arg0 *C.GtkColumnViewColumn
	var arg1 *C.char

	arg0 = (*C.GtkColumnViewColumn)(s.Native())
	arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_column_view_column_set_title(arg0, arg1)
}

// SetVisible sets whether this column should be visible in views.
func (s columnViewColumn) SetVisible(visible bool) {
	var arg0 *C.GtkColumnViewColumn
	var arg1 C.gboolean

	arg0 = (*C.GtkColumnViewColumn)(s.Native())
	if visible {
		arg1 = C.TRUE
	}

	C.gtk_column_view_column_set_visible(arg0, arg1)
}
