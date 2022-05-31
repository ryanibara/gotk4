// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkfilterlistmodel.go.
var GTypeFilterListModel = coreglib.Type(C.gtk_filter_list_model_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeFilterListModel, F: marshalFilterListModel},
	})
}

// FilterListModelOverrider contains methods that are overridable.
type FilterListModelOverrider interface {
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

func classInitFilterListModeller(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

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
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	if model != nil {
		_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(model).Native()))
	}
	if filter != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(filter).Native()))
	}
	*(*gio.ListModeller)(unsafe.Pointer(&args[0])) = _arg0
	*(**Filter)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "FilterListModel").InvokeMethod("new_FilterListModel", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**FilterListModel)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "FilterListModel").InvokeMethod("get_filter", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**FilterListModel)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "FilterListModel").InvokeMethod("get_incremental", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**FilterListModel)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "FilterListModel").InvokeMethod("get_model", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
func (self *FilterListModel) Pending() uint32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**FilterListModel)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "FilterListModel").InvokeMethod("get_pending", args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// SetFilter sets the filter used to filter items.
//
// The function takes the following parameters:
//
//    - filter (optional) to use or NULL to not filter items.
//
func (self *FilterListModel) SetFilter(filter *Filter) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if filter != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	}
	*(**FilterListModel)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "FilterListModel").InvokeMethod("set_filter", args[:], nil)

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
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if incremental {
		_arg1 = C.TRUE
	}
	*(**FilterListModel)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "FilterListModel").InvokeMethod("set_incremental", args[:], nil)

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
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if model != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}
	*(**FilterListModel)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "FilterListModel").InvokeMethod("set_model", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}
