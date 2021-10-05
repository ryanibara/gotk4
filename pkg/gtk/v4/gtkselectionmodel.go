// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_selection_model_get_type()), F: marshalSelectionModeller},
	})
}

// SelectionModelOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SelectionModelOverrider interface {
	// SelectionInRange gets the set of selected items in a range.
	//
	// This function is an optimization for gtk.SelectionModel.GetSelection()
	// when you are only interested in part of the model's selected state. A
	// common use case is in response to the
	// gtk.SelectionModel::selection-changed signal.
	SelectionInRange(position, nItems uint) *Bitset
	// IsSelected checks if the given item is selected.
	IsSelected(position uint) bool
	// SelectAll requests to select all items in the model.
	SelectAll() bool
	// SelectItem requests to select an item in the model.
	SelectItem(position uint, unselectRest bool) bool
	// SelectRange requests to select a range of items in the model.
	SelectRange(position, nItems uint, unselectRest bool) bool
	// SetSelection: make selection changes.
	//
	// This is the most advanced selection updating method that allows the most
	// fine-grained control over selection changes. If you can, you should try
	// the simpler versions, as implementations are more likely to implement
	// support for those.
	//
	// Requests that the selection state of all positions set in mask be updated
	// to the respective value in the selected bitmask.
	//
	// In pseudocode, it would look something like this:
	//
	//    for (i = 0; i < n_items; i++)
	//      {
	//        // don't change values not in the mask
	//        if (!gtk_bitset_contains (mask, i))
	//          continue;
	//
	//        if (gtk_bitset_contains (selected, i))
	//          select_item (i);
	//        else
	//          unselect_item (i);
	//      }
	//
	//    gtk_selection_model_selection_changed (model,
	//                                           first_changed_item,
	//                                           n_changed_items);
	//
	//
	// mask and selected must not be modified. They may refer to the same
	// bitset, which would mean that every item in the set should be selected.
	SetSelection(selected, mask *Bitset) bool
	// UnselectAll requests to unselect all items in the model.
	UnselectAll() bool
	// UnselectItem requests to unselect an item in the model.
	UnselectItem(position uint) bool
	// UnselectRange requests to unselect a range of items in the model.
	UnselectRange(position, nItems uint) bool
}

// SelectionModel: GtkSelectionModel is an interface that add support for
// selection to list models.
//
// This support is then used by widgets using list models to add the ability to
// select and unselect various items.
//
// GTK provides default implementations of the most common selection modes such
// as gtk.SingleSelection, so you will only need to implement this interface if
// you want detailed control about how selections should be handled.
//
// A GtkSelectionModel supports a single boolean per item indicating if an item
// is selected or not. This can be queried via gtk.SelectionModel.IsSelected().
// When the selected state of one or more items changes, the model will emit the
// gtk.SelectionModel::selection-changed signal by calling the
// gtk.SelectionModel.SelectionChanged() function. The positions given in that
// signal may have their selection state changed, though that is not a
// requirement. If new items added to the model via the ::items-changed signal
// are selected or not is up to the implementation.
//
// Note that items added via ::items-changed may already be selected and no
// [Gtk.SelectionModel::selection-changed] will be emitted for them. So to track
// which items are selected, it is necessary to listen to both signals.
//
// Additionally, the interface can expose functionality to select and unselect
// items. If these functions are implemented, GTK's list widgets will allow
// users to select and unselect items. However, GtkSelectionModels are free to
// only implement them partially or not at all. In that case the widgets will
// not support the unimplemented operations.
//
// When selecting or unselecting is supported by a model, the return values of
// the selection functions do *not* indicate if selection or unselection
// happened. They are only meant to indicate complete failure, like when this
// mode of selecting is not supported by the model.
//
// Selections may happen asynchronously, so the only reliable way to find out
// when an item was selected is to listen to the signals that indicate
// selection.
type SelectionModel struct {
	gio.ListModel
}

// SelectionModeller describes SelectionModel's abstract methods.
type SelectionModeller interface {
	externglib.Objector

	// Selection gets the set containing all currently selected items in the
	// model.
	Selection() *Bitset
	// SelectionInRange gets the set of selected items in a range.
	SelectionInRange(position, nItems uint) *Bitset
	// IsSelected checks if the given item is selected.
	IsSelected(position uint) bool
	// SelectAll requests to select all items in the model.
	SelectAll() bool
	// SelectItem requests to select an item in the model.
	SelectItem(position uint, unselectRest bool) bool
	// SelectRange requests to select a range of items in the model.
	SelectRange(position, nItems uint, unselectRest bool) bool
	// SelectionChanged: helper function for implementations of
	// GtkSelectionModel.
	SelectionChanged(position, nItems uint)
	// SetSelection: make selection changes.
	SetSelection(selected, mask *Bitset) bool
	// UnselectAll requests to unselect all items in the model.
	UnselectAll() bool
	// UnselectItem requests to unselect an item in the model.
	UnselectItem(position uint) bool
	// UnselectRange requests to unselect a range of items in the model.
	UnselectRange(position, nItems uint) bool
}

var _ SelectionModeller = (*SelectionModel)(nil)

func wrapSelectionModel(obj *externglib.Object) *SelectionModel {
	return &SelectionModel{
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalSelectionModeller(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSelectionModel(obj), nil
}

// Selection gets the set containing all currently selected items in the model.
//
// This function may be slow, so if you are only interested in single item,
// consider using gtk.SelectionModel.IsSelected() or if you are only interested
// in a few, consider gtk.SelectionModel.GetSelectionInRange().
func (model *SelectionModel) Selection() *Bitset {
	var _arg0 *C.GtkSelectionModel // out
	var _cret *C.GtkBitset         // in

	_arg0 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_selection_model_get_selection(_arg0)
	runtime.KeepAlive(model)

	var _bitset *Bitset // out

	_bitset = (*Bitset)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bitset)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_bitset_unref((*C.GtkBitset)(intern.C))
		},
	)

	return _bitset
}

// SelectionInRange gets the set of selected items in a range.
//
// This function is an optimization for gtk.SelectionModel.GetSelection() when
// you are only interested in part of the model's selected state. A common use
// case is in response to the gtk.SelectionModel::selection-changed signal.
func (model *SelectionModel) SelectionInRange(position, nItems uint) *Bitset {
	var _arg0 *C.GtkSelectionModel // out
	var _arg1 C.guint              // out
	var _arg2 C.guint              // out
	var _cret *C.GtkBitset         // in

	_arg0 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
	_arg1 = C.guint(position)
	_arg2 = C.guint(nItems)

	_cret = C.gtk_selection_model_get_selection_in_range(_arg0, _arg1, _arg2)
	runtime.KeepAlive(model)
	runtime.KeepAlive(position)
	runtime.KeepAlive(nItems)

	var _bitset *Bitset // out

	_bitset = (*Bitset)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bitset)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_bitset_unref((*C.GtkBitset)(intern.C))
		},
	)

	return _bitset
}

// IsSelected checks if the given item is selected.
func (model *SelectionModel) IsSelected(position uint) bool {
	var _arg0 *C.GtkSelectionModel // out
	var _arg1 C.guint              // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
	_arg1 = C.guint(position)

	_cret = C.gtk_selection_model_is_selected(_arg0, _arg1)
	runtime.KeepAlive(model)
	runtime.KeepAlive(position)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectAll requests to select all items in the model.
func (model *SelectionModel) SelectAll() bool {
	var _arg0 *C.GtkSelectionModel // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_selection_model_select_all(_arg0)
	runtime.KeepAlive(model)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectItem requests to select an item in the model.
func (model *SelectionModel) SelectItem(position uint, unselectRest bool) bool {
	var _arg0 *C.GtkSelectionModel // out
	var _arg1 C.guint              // out
	var _arg2 C.gboolean           // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
	_arg1 = C.guint(position)
	if unselectRest {
		_arg2 = C.TRUE
	}

	_cret = C.gtk_selection_model_select_item(_arg0, _arg1, _arg2)
	runtime.KeepAlive(model)
	runtime.KeepAlive(position)
	runtime.KeepAlive(unselectRest)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectRange requests to select a range of items in the model.
func (model *SelectionModel) SelectRange(position, nItems uint, unselectRest bool) bool {
	var _arg0 *C.GtkSelectionModel // out
	var _arg1 C.guint              // out
	var _arg2 C.guint              // out
	var _arg3 C.gboolean           // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
	_arg1 = C.guint(position)
	_arg2 = C.guint(nItems)
	if unselectRest {
		_arg3 = C.TRUE
	}

	_cret = C.gtk_selection_model_select_range(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(model)
	runtime.KeepAlive(position)
	runtime.KeepAlive(nItems)
	runtime.KeepAlive(unselectRest)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectionChanged: helper function for implementations of GtkSelectionModel.
//
// Call this when a the selection changes to emit the
// gtk.SelectionModel::selection-changed signal.
func (model *SelectionModel) SelectionChanged(position, nItems uint) {
	var _arg0 *C.GtkSelectionModel // out
	var _arg1 C.guint              // out
	var _arg2 C.guint              // out

	_arg0 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
	_arg1 = C.guint(position)
	_arg2 = C.guint(nItems)

	C.gtk_selection_model_selection_changed(_arg0, _arg1, _arg2)
	runtime.KeepAlive(model)
	runtime.KeepAlive(position)
	runtime.KeepAlive(nItems)
}

// SetSelection: make selection changes.
//
// This is the most advanced selection updating method that allows the most
// fine-grained control over selection changes. If you can, you should try the
// simpler versions, as implementations are more likely to implement support for
// those.
//
// Requests that the selection state of all positions set in mask be updated to
// the respective value in the selected bitmask.
//
// In pseudocode, it would look something like this:
//
//    for (i = 0; i < n_items; i++)
//      {
//        // don't change values not in the mask
//        if (!gtk_bitset_contains (mask, i))
//          continue;
//
//        if (gtk_bitset_contains (selected, i))
//          select_item (i);
//        else
//          unselect_item (i);
//      }
//
//    gtk_selection_model_selection_changed (model,
//                                           first_changed_item,
//                                           n_changed_items);
//
//
// mask and selected must not be modified. They may refer to the same bitset,
// which would mean that every item in the set should be selected.
func (model *SelectionModel) SetSelection(selected, mask *Bitset) bool {
	var _arg0 *C.GtkSelectionModel // out
	var _arg1 *C.GtkBitset         // out
	var _arg2 *C.GtkBitset         // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
	_arg1 = (*C.GtkBitset)(gextras.StructNative(unsafe.Pointer(selected)))
	_arg2 = (*C.GtkBitset)(gextras.StructNative(unsafe.Pointer(mask)))

	_cret = C.gtk_selection_model_set_selection(_arg0, _arg1, _arg2)
	runtime.KeepAlive(model)
	runtime.KeepAlive(selected)
	runtime.KeepAlive(mask)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnselectAll requests to unselect all items in the model.
func (model *SelectionModel) UnselectAll() bool {
	var _arg0 *C.GtkSelectionModel // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_selection_model_unselect_all(_arg0)
	runtime.KeepAlive(model)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnselectItem requests to unselect an item in the model.
func (model *SelectionModel) UnselectItem(position uint) bool {
	var _arg0 *C.GtkSelectionModel // out
	var _arg1 C.guint              // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
	_arg1 = C.guint(position)

	_cret = C.gtk_selection_model_unselect_item(_arg0, _arg1)
	runtime.KeepAlive(model)
	runtime.KeepAlive(position)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnselectRange requests to unselect a range of items in the model.
func (model *SelectionModel) UnselectRange(position, nItems uint) bool {
	var _arg0 *C.GtkSelectionModel // out
	var _arg1 C.guint              // out
	var _arg2 C.guint              // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
	_arg1 = C.guint(position)
	_arg2 = C.guint(nItems)

	_cret = C.gtk_selection_model_unselect_range(_arg0, _arg1, _arg2)
	runtime.KeepAlive(model)
	runtime.KeepAlive(position)
	runtime.KeepAlive(nItems)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ConnectSelectionChanged: emitted when the selection state of some of the
// items in model changes.
//
// Note that this signal does not specify the new selection state of the items,
// they need to be queried manually. It is also not necessary for a model to
// change the selection state of any of the items in the selection model, though
// it would be rather useless to emit such a signal.
func (s *SelectionModel) ConnectSelectionChanged(f func(position, nItems uint)) glib.SignalHandle {
	return s.Connect("selection-changed", f)
}
