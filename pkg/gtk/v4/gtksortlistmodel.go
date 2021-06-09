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
		{T: externglib.Type(C.gtk_sort_list_model_get_type()), F: marshalSortListModel},
	})
}

// SortListModel: `GtkSortListModel` is a list model that sorts the elements of
// the underlying model according to a `GtkSorter`.
//
// The model can be set up to do incremental sorting, so that sorting long lists
// doesn't block the UI. See [method@Gtk.SortListModel.set_incremental] for
// details.
//
// `GtkSortListModel` is a generic model and because of that it cannot take
// advantage of any external knowledge when sorting. If you run into performance
// issues with `GtkSortListModel`, it is strongly recommended that you write
// your own sorting list model.
type SortListModel interface {
	gextras.Objector
	gio.ListModel

	// Incremental returns whether incremental sorting is enabled.
	//
	// See [method@Gtk.SortListModel.set_incremental].
	Incremental() bool
	// Model gets the model currently sorted or nil if none.
	Model() gio.ListModel
	// Pending estimates progress of an ongoing sorting operation.
	//
	// The estimate is the number of items that would still need to be sorted to
	// finish the sorting operation if this was a linear algorithm. So this
	// number is not related to how many items are already correctly sorted.
	//
	// If you want to estimate the progress, you can use code like this: “`c
	// pending = gtk_sort_list_model_get_pending (self); model =
	// gtk_sort_list_model_get_model (self); progress = 1.0 - pending / (double)
	// MAX (1, g_list_model_get_n_items (model)); “`
	//
	// If no sort operation is ongoing - in particular when
	// [property@Gtk.SortListModel:incremental] is false - this function returns
	// 0.
	Pending() uint
	// Sorter gets the sorter that is used to sort @self.
	Sorter() Sorter
	// SetIncremental sets the sort model to do an incremental sort.
	//
	// When incremental sorting is enabled, the `GtkSortListModel` will not do a
	// complete sort immediately, but will instead queue an idle handler that
	// incrementally sorts the items towards their correct position. This of
	// course means that items do not instantly appear in the right place. It
	// also means that the total sorting time is a lot slower.
	//
	// When your filter blocks the UI while sorting, you might consider turning
	// this on. Depending on your model and sorters, this may become interesting
	// around 10,000 to 100,000 items.
	//
	// By default, incremental sorting is disabled.
	//
	// See [method@Gtk.SortListModel.get_pending] for progress information about
	// an ongoing incremental sorting operation.
	SetIncremental(incremental bool)
	// SetModel sets the model to be sorted.
	//
	// The @model's item type must conform to the item type of @self.
	SetModel(model gio.ListModel)
	// SetSorter sets a new sorter on @self.
	SetSorter(sorter Sorter)
}

// sortListModel implements the SortListModel interface.
type sortListModel struct {
	gextras.Objector
	gio.ListModel
}

var _ SortListModel = (*sortListModel)(nil)

// WrapSortListModel wraps a GObject to the right type. It is
// primarily used internally.
func WrapSortListModel(obj *externglib.Object) SortListModel {
	return SortListModel{
		Objector:      obj,
		gio.ListModel: gio.WrapListModel(obj),
	}
}

func marshalSortListModel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSortListModel(obj), nil
}

// NewSortListModel constructs a class SortListModel.
func NewSortListModel(model gio.ListModel, sorter Sorter) SortListModel {
	var _arg1 *C.GListModel
	var _arg2 *C.GtkSorter

	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
	_arg2 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))

	var _cret C.GtkSortListModel

	cret = C.gtk_sort_list_model_new(_arg1, _arg2)

	var _sortListModel SortListModel

	_sortListModel = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(SortListModel)

	return _sortListModel
}

// Incremental returns whether incremental sorting is enabled.
//
// See [method@Gtk.SortListModel.set_incremental].
func (s sortListModel) Incremental() bool {
	var _arg0 *C.GtkSortListModel

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	cret = C.gtk_sort_list_model_get_incremental(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Model gets the model currently sorted or nil if none.
func (s sortListModel) Model() gio.ListModel {
	var _arg0 *C.GtkSortListModel

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer(s.Native()))

	var _cret *C.GListModel

	cret = C.gtk_sort_list_model_get_model(_arg0)

	var _listModel gio.ListModel

	_listModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.ListModel)

	return _listModel
}

// Pending estimates progress of an ongoing sorting operation.
//
// The estimate is the number of items that would still need to be sorted to
// finish the sorting operation if this was a linear algorithm. So this
// number is not related to how many items are already correctly sorted.
//
// If you want to estimate the progress, you can use code like this: “`c
// pending = gtk_sort_list_model_get_pending (self); model =
// gtk_sort_list_model_get_model (self); progress = 1.0 - pending / (double)
// MAX (1, g_list_model_get_n_items (model)); “`
//
// If no sort operation is ongoing - in particular when
// [property@Gtk.SortListModel:incremental] is false - this function returns
// 0.
func (s sortListModel) Pending() uint {
	var _arg0 *C.GtkSortListModel

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer(s.Native()))

	var _cret C.guint

	cret = C.gtk_sort_list_model_get_pending(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// Sorter gets the sorter that is used to sort @self.
func (s sortListModel) Sorter() Sorter {
	var _arg0 *C.GtkSortListModel

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkSorter

	cret = C.gtk_sort_list_model_get_sorter(_arg0)

	var _sorter Sorter

	_sorter = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Sorter)

	return _sorter
}

// SetIncremental sets the sort model to do an incremental sort.
//
// When incremental sorting is enabled, the `GtkSortListModel` will not do a
// complete sort immediately, but will instead queue an idle handler that
// incrementally sorts the items towards their correct position. This of
// course means that items do not instantly appear in the right place. It
// also means that the total sorting time is a lot slower.
//
// When your filter blocks the UI while sorting, you might consider turning
// this on. Depending on your model and sorters, this may become interesting
// around 10,000 to 100,000 items.
//
// By default, incremental sorting is disabled.
//
// See [method@Gtk.SortListModel.get_pending] for progress information about
// an ongoing incremental sorting operation.
func (s sortListModel) SetIncremental(incremental bool) {
	var _arg0 *C.GtkSortListModel
	var _arg1 C.gboolean

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer(s.Native()))
	if incremental {
		_arg1 = C.gboolean(1)
	}

	C.gtk_sort_list_model_set_incremental(_arg0, _arg1)
}

// SetModel sets the model to be sorted.
//
// The @model's item type must conform to the item type of @self.
func (s sortListModel) SetModel(model gio.ListModel) {
	var _arg0 *C.GtkSortListModel
	var _arg1 *C.GListModel

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.gtk_sort_list_model_set_model(_arg0, _arg1)
}

// SetSorter sets a new sorter on @self.
func (s sortListModel) SetSorter(sorter Sorter) {
	var _arg0 *C.GtkSortListModel
	var _arg1 *C.GtkSorter

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))

	C.gtk_sort_list_model_set_sorter(_arg0, _arg1)
}