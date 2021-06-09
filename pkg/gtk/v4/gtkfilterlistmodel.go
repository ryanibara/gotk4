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
		{T: externglib.Type(C.gtk_filter_list_model_get_type()), F: marshalFilterListModel},
	})
}

// FilterListModel: `GtkFilterListModel` is a list model that filters the
// elements of the underlying model according to a `GtkFilter`.
//
// It hides some elements from the other model according to criteria given by a
// `GtkFilter`.
//
// The model can be set up to do incremental searching, so that filtering long
// lists doesn't block the UI. See [method@Gtk.FilterListModel.set_incremental]
// for details.
type FilterListModel interface {
	gextras.Objector
	gio.ListModel

	// Filter gets the `GtkFilter` currently set on @self.
	Filter() Filter
	// Incremental returns whether incremental filtering is enabled.
	//
	// See [method@Gtk.FilterListModel.set_incremental].
	Incremental() bool
	// Model gets the model currently filtered or nil if none.
	Model() gio.ListModel
	// Pending returns the number of items that have not been filtered yet.
	//
	// You can use this value to check if @self is busy filtering by comparing
	// the return value to 0 or you can compute the percentage of the filter
	// remaining by dividing the return value by the total number of items in
	// the underlying model:
	//
	// “`c pending = gtk_filter_list_model_get_pending (self); model =
	// gtk_filter_list_model_get_model (self); percentage = pending / (double)
	// g_list_model_get_n_items (model); “`
	//
	// If no filter operation is ongoing - in particular when
	// [property@Gtk.FilterListModel:incremental] is false - this function
	// returns 0.
	Pending() uint
	// SetFilter sets the filter used to filter items.
	SetFilter(filter Filter)
	// SetIncremental sets the filter model to do an incremental sort.
	//
	// When incremental filtering is enabled, the `GtkFilterListModel` will not
	// run filters immediately, but will instead queue an idle handler that
	// incrementally filters the items and adds them to the list. This of course
	// means that items are not instantly added to the list, but only appear
	// incrementally.
	//
	// When your filter blocks the UI while filtering, you might consider
	// turning this on. Depending on your model and filters, this may become
	// interesting around 10,000 to 100,000 items.
	//
	// By default, incremental filtering is disabled.
	//
	// See [method@Gtk.FilterListModel.get_pending] for progress information
	// about an ongoing incremental filtering operation.
	SetIncremental(incremental bool)
	// SetModel sets the model to be filtered.
	//
	// Note that GTK makes no effort to ensure that @model conforms to the item
	// type of @self. It assumes that the caller knows what they are doing and
	// have set up an appropriate filter to ensure that item types match.
	SetModel(model gio.ListModel)
}

// filterListModel implements the FilterListModel interface.
type filterListModel struct {
	gextras.Objector
	gio.ListModel
}

var _ FilterListModel = (*filterListModel)(nil)

// WrapFilterListModel wraps a GObject to the right type. It is
// primarily used internally.
func WrapFilterListModel(obj *externglib.Object) FilterListModel {
	return FilterListModel{
		Objector:      obj,
		gio.ListModel: gio.WrapListModel(obj),
	}
}

func marshalFilterListModel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFilterListModel(obj), nil
}

// NewFilterListModel constructs a class FilterListModel.
func NewFilterListModel(model gio.ListModel, filter Filter) FilterListModel {
	var _arg1 *C.GListModel
	var _arg2 *C.GtkFilter

	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
	_arg2 = (*C.GtkFilter)(unsafe.Pointer(filter.Native()))

	var _cret C.GtkFilterListModel

	cret = C.gtk_filter_list_model_new(_arg1, _arg2)

	var _filterListModel FilterListModel

	_filterListModel = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(FilterListModel)

	return _filterListModel
}

// Filter gets the `GtkFilter` currently set on @self.
func (s filterListModel) Filter() Filter {
	var _arg0 *C.GtkFilterListModel

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkFilter

	cret = C.gtk_filter_list_model_get_filter(_arg0)

	var _filter Filter

	_filter = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Filter)

	return _filter
}

// Incremental returns whether incremental filtering is enabled.
//
// See [method@Gtk.FilterListModel.set_incremental].
func (s filterListModel) Incremental() bool {
	var _arg0 *C.GtkFilterListModel

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	cret = C.gtk_filter_list_model_get_incremental(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Model gets the model currently filtered or nil if none.
func (s filterListModel) Model() gio.ListModel {
	var _arg0 *C.GtkFilterListModel

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(s.Native()))

	var _cret *C.GListModel

	cret = C.gtk_filter_list_model_get_model(_arg0)

	var _listModel gio.ListModel

	_listModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.ListModel)

	return _listModel
}

// Pending returns the number of items that have not been filtered yet.
//
// You can use this value to check if @self is busy filtering by comparing
// the return value to 0 or you can compute the percentage of the filter
// remaining by dividing the return value by the total number of items in
// the underlying model:
//
// “`c pending = gtk_filter_list_model_get_pending (self); model =
// gtk_filter_list_model_get_model (self); percentage = pending / (double)
// g_list_model_get_n_items (model); “`
//
// If no filter operation is ongoing - in particular when
// [property@Gtk.FilterListModel:incremental] is false - this function
// returns 0.
func (s filterListModel) Pending() uint {
	var _arg0 *C.GtkFilterListModel

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(s.Native()))

	var _cret C.guint

	cret = C.gtk_filter_list_model_get_pending(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// SetFilter sets the filter used to filter items.
func (s filterListModel) SetFilter(filter Filter) {
	var _arg0 *C.GtkFilterListModel
	var _arg1 *C.GtkFilter

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_filter_list_model_set_filter(_arg0, _arg1)
}

// SetIncremental sets the filter model to do an incremental sort.
//
// When incremental filtering is enabled, the `GtkFilterListModel` will not
// run filters immediately, but will instead queue an idle handler that
// incrementally filters the items and adds them to the list. This of course
// means that items are not instantly added to the list, but only appear
// incrementally.
//
// When your filter blocks the UI while filtering, you might consider
// turning this on. Depending on your model and filters, this may become
// interesting around 10,000 to 100,000 items.
//
// By default, incremental filtering is disabled.
//
// See [method@Gtk.FilterListModel.get_pending] for progress information
// about an ongoing incremental filtering operation.
func (s filterListModel) SetIncremental(incremental bool) {
	var _arg0 *C.GtkFilterListModel
	var _arg1 C.gboolean

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(s.Native()))
	if incremental {
		_arg1 = C.gboolean(1)
	}

	C.gtk_filter_list_model_set_incremental(_arg0, _arg1)
}

// SetModel sets the model to be filtered.
//
// Note that GTK makes no effort to ensure that @model conforms to the item
// type of @self. It assumes that the caller knows what they are doing and
// have set up an appropriate filter to ensure that item types match.
func (s filterListModel) SetModel(model gio.ListModel) {
	var _arg0 *C.GtkFilterListModel
	var _arg1 *C.GListModel

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.gtk_filter_list_model_set_model(_arg0, _arg1)
}