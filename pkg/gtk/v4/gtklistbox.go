// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_list_box_row_get_type()), F: marshalListBoxRow},
	})
}

// ListBoxCreateWidgetFunc: called for list boxes that are bound to a
// `GListModel` with gtk_list_box_bind_model() for each item that gets added to
// the model.
type ListBoxCreateWidgetFunc func(item gextras.Objector) (widget Widget)

//export gotk4_ListBoxCreateWidgetFunc
func gotk4_ListBoxCreateWidgetFunc(arg0 C.gpointer, arg1 C.gpointer) *C.GtkWidget {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var item gextras.Objector

	item = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0.Native()))).(gextras.Objector)

	fn := v.(ListBoxCreateWidgetFunc)
	widget := fn(item)

	cret = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	return widget
}

// ListBoxRow: `GtkListBoxRow` is the kind of widget that can be added to a
// `GtkListBox`.
type ListBoxRow interface {
	Widget
	Accessible
	Actionable
	Buildable
	ConstraintTarget

	// Changed marks @row as changed, causing any state that depends on this to
	// be updated.
	//
	// This affects sorting, filtering and headers.
	//
	// Note that calls to this method must be in sync with the data used for the
	// row functions. For instance, if the list is mirroring some external data
	// set, and *two* rows changed in the external data set then when you call
	// gtk_list_box_row_changed() on the first row the sort function must only
	// read the new data for the first of the two changed rows, otherwise the
	// resorting of the rows will be wrong.
	//
	// This generally means that if you don’t fully control the data model you
	// have to duplicate the data that affects the listbox row functions into
	// the row widgets themselves. Another alternative is to call
	// [method@Gtk.ListBox.invalidate_sort] on any model change, but that is
	// more expensive.
	Changed()
	// Activatable gets whether the row is activatable.
	Activatable() bool
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
	SetChild(child Widget)
	// SetHeader sets the current header of the @row.
	//
	// This is only allowed to be called from a
	// [callback@Gtk.ListBoxUpdateHeaderFunc]. It will replace any existing
	// header in the row, and be shown in front of the row in the listbox.
	SetHeader(header Widget)
	// SetSelectable: set whether the row can be selected.
	SetSelectable(selectable bool)
}

// listBoxRow implements the ListBoxRow interface.
type listBoxRow struct {
	Widget
	Accessible
	Actionable
	Buildable
	ConstraintTarget
}

var _ ListBoxRow = (*listBoxRow)(nil)

// WrapListBoxRow wraps a GObject to the right type. It is
// primarily used internally.
func WrapListBoxRow(obj *externglib.Object) ListBoxRow {
	return ListBoxRow{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Actionable:       WrapActionable(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalListBoxRow(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapListBoxRow(obj), nil
}

// Changed marks @row as changed, causing any state that depends on this to
// be updated.
//
// This affects sorting, filtering and headers.
//
// Note that calls to this method must be in sync with the data used for the
// row functions. For instance, if the list is mirroring some external data
// set, and *two* rows changed in the external data set then when you call
// gtk_list_box_row_changed() on the first row the sort function must only
// read the new data for the first of the two changed rows, otherwise the
// resorting of the rows will be wrong.
//
// This generally means that if you don’t fully control the data model you
// have to duplicate the data that affects the listbox row functions into
// the row widgets themselves. Another alternative is to call
// [method@Gtk.ListBox.invalidate_sort] on any model change, but that is
// more expensive.
func (r listBoxRow) Changed() {
	var _arg0 *C.GtkListBoxRow

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	C.gtk_list_box_row_changed(_arg0)
}

// Activatable gets whether the row is activatable.
func (r listBoxRow) Activatable() bool {
	var _arg0 *C.GtkListBoxRow

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	var _cret C.gboolean

	_cret = C.gtk_list_box_row_get_activatable(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Index gets the current index of the @row in its `GtkListBox` container.
func (r listBoxRow) Index() int {
	var _arg0 *C.GtkListBoxRow

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	var _cret C.int

	_cret = C.gtk_list_box_row_get_index(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Selectable gets whether the row can be selected.
func (r listBoxRow) Selectable() bool {
	var _arg0 *C.GtkListBoxRow

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	var _cret C.gboolean

	_cret = C.gtk_list_box_row_get_selectable(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsSelected returns whether the child is currently selected in its
// `GtkListBox` container.
func (r listBoxRow) IsSelected() bool {
	var _arg0 *C.GtkListBoxRow

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	var _cret C.gboolean

	_cret = C.gtk_list_box_row_is_selected(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetActivatable: set whether the row is activatable.
func (r listBoxRow) SetActivatable(activatable bool) {
	var _arg0 *C.GtkListBoxRow
	var _arg1 C.gboolean

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))
	if activatable {
		_arg1 = C.gboolean(1)
	}

	C.gtk_list_box_row_set_activatable(_arg0, _arg1)
}

// SetChild sets the child widget of @self.
func (r listBoxRow) SetChild(child Widget) {
	var _arg0 *C.GtkListBoxRow
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_list_box_row_set_child(_arg0, _arg1)
}

// SetHeader sets the current header of the @row.
//
// This is only allowed to be called from a
// [callback@Gtk.ListBoxUpdateHeaderFunc]. It will replace any existing
// header in the row, and be shown in front of the row in the listbox.
func (r listBoxRow) SetHeader(header Widget) {
	var _arg0 *C.GtkListBoxRow
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(header.Native()))

	C.gtk_list_box_row_set_header(_arg0, _arg1)
}

// SetSelectable: set whether the row can be selected.
func (r listBoxRow) SetSelectable(selectable bool) {
	var _arg0 *C.GtkListBoxRow
	var _arg1 C.gboolean

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))
	if selectable {
		_arg1 = C.gboolean(1)
	}

	C.gtk_list_box_row_set_selectable(_arg0, _arg1)
}
