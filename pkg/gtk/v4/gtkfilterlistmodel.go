// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeFilterListModel = coreglib.Type(C.gtk_filter_list_model_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFilterListModel, F: marshalFilterListModel},
	})
}

// FilterListModelOverrides contains methods that are overridable.
type FilterListModelOverrides struct {
}

func defaultFilterListModelOverrides(v *FilterListModel) FilterListModelOverrides {
	return FilterListModelOverrides{}
}

// FilterListModel: GtkFilterListModel is a list model that filters the elements
// of the underlying model according to a GtkFilter.
//
// It hides some elements from the other model according to criteria given by a
// GtkFilter.
//
// The model can be set up to do incremental searching, so that filtering long
// lists doesn't block the UI. See gtk.FilterListModel.SetIncremental() for
// details.
type FilterListModel struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.ListModel
}

var (
	_ coreglib.Objector = (*FilterListModel)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*FilterListModel, *FilterListModelClass, FilterListModelOverrides](
		GTypeFilterListModel,
		initFilterListModelClass,
		wrapFilterListModel,
		defaultFilterListModelOverrides,
	)
}

func initFilterListModelClass(gclass unsafe.Pointer, overrides FilterListModelOverrides, classInitFunc func(*FilterListModelClass)) {
	if classInitFunc != nil {
		class := (*FilterListModelClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapFilterListModel(obj *coreglib.Object) *FilterListModel {
	return &FilterListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalFilterListModel(p uintptr) (interface{}, error) {
	return wrapFilterListModel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFilterListModel creates a new GtkFilterListModel that will filter model
// using the given filter.
//
// The function takes the following parameters:
//
//    - model (optional) to sort, or NULL.
//    - filter (optional) or NULL to not filter items.
//
// The function returns the following values:
//
//    - filterListModel: new GtkFilterListModel.
//
func NewFilterListModel(model gio.ListModeller, filter *Filter) *FilterListModel {
	var _arg1 *C.GListModel         // out
	var _arg2 *C.GtkFilter          // out
	var _cret *C.GtkFilterListModel // in

	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(model).Native()))
	}
	if filter != nil {
		_arg2 = (*C.GtkFilter)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(filter).Native()))
	}

	_cret = C.gtk_filter_list_model_new(_arg1, _arg2)
	runtime.KeepAlive(model)
	runtime.KeepAlive(filter)

	var _filterListModel *FilterListModel // out

	_filterListModel = wrapFilterListModel(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _filterListModel
}

// Filter gets the GtkFilter currently set on self.
//
// The function returns the following values:
//
//    - filter (optional) currently in use or NULL if the list isn't filtered.
//
func (self *FilterListModel) Filter() *Filter {
	var _arg0 *C.GtkFilterListModel // out
	var _cret *C.GtkFilter          // in

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_filter_list_model_get_filter(_arg0)
	runtime.KeepAlive(self)

	var _filter *Filter // out

	if _cret != nil {
		_filter = wrapFilter(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _filter
}

// Incremental returns whether incremental filtering is enabled.
//
// See gtk.FilterListModel.SetIncremental().
//
// The function returns the following values:
//
//    - ok: TRUE if incremental filtering is enabled.
//
func (self *FilterListModel) Incremental() bool {
	var _arg0 *C.GtkFilterListModel // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_filter_list_model_get_incremental(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Model gets the model currently filtered or NULL if none.
//
// The function returns the following values:
//
//    - listModel (optional): model that gets filtered.
//
func (self *FilterListModel) Model() *gio.ListModel {
	var _arg0 *C.GtkFilterListModel // out
	var _cret *C.GListModel         // in

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_filter_list_model_get_model(_arg0)
	runtime.KeepAlive(self)

	var _listModel *gio.ListModel // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_listModel = &gio.ListModel{
				Object: obj,
			}
		}
	}

	return _listModel
}

// Pending returns the number of items that have not been filtered yet.
//
// You can use this value to check if self is busy filtering by comparing the
// return value to 0 or you can compute the percentage of the filter remaining
// by dividing the return value by the total number of items in the underlying
// model:
//
//    pending = gtk_filter_list_model_get_pending (self);
//    model = gtk_filter_list_model_get_model (self);
//    percentage = pending / (double) g_list_model_get_n_items (model);
//
//
// If no filter operation is ongoing - in particular when
// gtk.FilterListModel:incremental is FALSE - this function returns 0.
//
// The function returns the following values:
//
//    - guint: number of items not yet filtered.
//
func (self *FilterListModel) Pending() uint {
	var _arg0 *C.GtkFilterListModel // out
	var _cret C.guint               // in

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_filter_list_model_get_pending(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SetFilter sets the filter used to filter items.
//
// The function takes the following parameters:
//
//    - filter (optional) to use or NULL to not filter items.
//
func (self *FilterListModel) SetFilter(filter *Filter) {
	var _arg0 *C.GtkFilterListModel // out
	var _arg1 *C.GtkFilter          // out

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if filter != nil {
		_arg1 = (*C.GtkFilter)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	}

	C.gtk_filter_list_model_set_filter(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(filter)
}

// SetIncremental sets the filter model to do an incremental sort.
//
// When incremental filtering is enabled, the GtkFilterListModel will not run
// filters immediately, but will instead queue an idle handler that
// incrementally filters the items and adds them to the list. This of course
// means that items are not instantly added to the list, but only appear
// incrementally.
//
// When your filter blocks the UI while filtering, you might consider turning
// this on. Depending on your model and filters, this may become interesting
// around 10,000 to 100,000 items.
//
// By default, incremental filtering is disabled.
//
// See gtk.FilterListModel.GetPending() for progress information about an
// ongoing incremental filtering operation.
//
// The function takes the following parameters:
//
//    - incremental: TRUE to enable incremental filtering.
//
func (self *FilterListModel) SetIncremental(incremental bool) {
	var _arg0 *C.GtkFilterListModel // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if incremental {
		_arg1 = C.TRUE
	}

	C.gtk_filter_list_model_set_incremental(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(incremental)
}

// SetModel sets the model to be filtered.
//
// Note that GTK makes no effort to ensure that model conforms to the item type
// of self. It assumes that the caller knows what they are doing and have set up
// an appropriate filter to ensure that item types match.
//
// The function takes the following parameters:
//
//    - model (optional) to be filtered.
//
func (self *FilterListModel) SetModel(model gio.ListModeller) {
	var _arg0 *C.GtkFilterListModel // out
	var _arg1 *C.GListModel         // out

	_arg0 = (*C.GtkFilterListModel)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	C.gtk_filter_list_model_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}

// FilterListModelClass: instance of this type is always passed by reference.
type FilterListModelClass struct {
	*filterListModelClass
}

// filterListModelClass is the struct that's finalized.
type filterListModelClass struct {
	native *C.GtkFilterListModelClass
}
