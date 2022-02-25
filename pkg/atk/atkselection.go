// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
// extern AtkObject* _gotk4_atk1_SelectionIface_ref_selection(AtkSelection*, gint);
// extern gboolean _gotk4_atk1_SelectionIface_add_selection(AtkSelection*, gint);
// extern gboolean _gotk4_atk1_SelectionIface_clear_selection(AtkSelection*);
// extern gboolean _gotk4_atk1_SelectionIface_is_child_selected(AtkSelection*, gint);
// extern gboolean _gotk4_atk1_SelectionIface_remove_selection(AtkSelection*, gint);
// extern gboolean _gotk4_atk1_SelectionIface_select_all_selection(AtkSelection*);
// extern gint _gotk4_atk1_SelectionIface_get_selection_count(AtkSelection*);
// extern void _gotk4_atk1_SelectionIface_selection_changed(AtkSelection*);
// extern void _gotk4_atk1_Selection_ConnectSelectionChanged(gpointer, guintptr);
import "C"

// glib.Type values for atkselection.go.
var GTypeSelection = externglib.Type(C.atk_selection_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeSelection, F: marshalSelection},
	})
}

// SelectionOverrider contains methods that are overridable.
type SelectionOverrider interface {
	// AddSelection adds the specified accessible child of the object to the
	// object's selection.
	//
	// The function takes the following parameters:
	//
	//    - i specifying the child index.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if success, FALSE otherwise.
	//
	AddSelection(i int) bool
	// ClearSelection clears the selection in the object so that no children in
	// the object are selected.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if success, FALSE otherwise.
	//
	ClearSelection() bool
	// SelectionCount gets the number of accessible children currently selected.
	// Note: callers should not rely on NULL or on a zero value for indication
	// of whether AtkSelectionIface is implemented, they should use type
	// checking/interface checking macros or the atk_get_accessible_value()
	// convenience method.
	//
	// The function returns the following values:
	//
	//    - gint representing the number of items selected, or 0 if selection
	//      does not implement this interface.
	//
	SelectionCount() int
	// IsChildSelected determines if the current child of this object is
	// selected Note: callers should not rely on NULL or on a zero value for
	// indication of whether AtkSelectionIface is implemented, they should use
	// type checking/interface checking macros or the atk_get_accessible_value()
	// convenience method.
	//
	// The function takes the following parameters:
	//
	//    - i specifying the child index.
	//
	// The function returns the following values:
	//
	//    - ok: gboolean representing the specified child is selected, or 0 if
	//      selection does not implement this interface.
	//
	IsChildSelected(i int) bool
	// RefSelection gets a reference to the accessible object representing the
	// specified selected child of the object. Note: callers should not rely on
	// NULL or on a zero value for indication of whether AtkSelectionIface is
	// implemented, they should use type checking/interface checking macros or
	// the atk_get_accessible_value() convenience method.
	//
	// The function takes the following parameters:
	//
	//    - i specifying the index in the selection set. (e.g. the ith selection
	//      as opposed to the ith child).
	//
	// The function returns the following values:
	//
	//    - object (optional) representing the selected accessible, or NULL if
	//      selection does not implement this interface.
	//
	RefSelection(i int) *ObjectClass
	// RemoveSelection removes the specified child of the object from the
	// object's selection.
	//
	// The function takes the following parameters:
	//
	//    - i specifying the index in the selection set. (e.g. the ith selection
	//      as opposed to the ith child).
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if success, FALSE otherwise.
	//
	RemoveSelection(i int) bool
	// SelectAllSelection causes every child of the object to be selected if the
	// object supports multiple selections.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if success, FALSE otherwise.
	//
	SelectAllSelection() bool
	SelectionChanged()
}

// Selection should be implemented by UI components with children which are
// exposed by #atk_object_ref_child and #atk_object_get_n_children, if the use
// of the parent UI component ordinarily involves selection of one or more of
// the objects corresponding to those Object children - for example, selectable
// lists.
//
// Note that other types of "selection" (for instance text selection) are
// accomplished a other ATK interfaces - Selection is limited to the
// selection/deselection of children.
type Selection struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Selection)(nil)
)

// Selectioner describes Selection's interface methods.
type Selectioner interface {
	externglib.Objector

	// AddSelection adds the specified accessible child of the object to the
	// object's selection.
	AddSelection(i int) bool
	// ClearSelection clears the selection in the object so that no children in
	// the object are selected.
	ClearSelection() bool
	// SelectionCount gets the number of accessible children currently selected.
	SelectionCount() int
	// IsChildSelected determines if the current child of this object is
	// selected Note: callers should not rely on NULL or on a zero value for
	// indication of whether AtkSelectionIface is implemented, they should use
	// type checking/interface checking macros or the atk_get_accessible_value()
	// convenience method.
	IsChildSelected(i int) bool
	// RefSelection gets a reference to the accessible object representing the
	// specified selected child of the object.
	RefSelection(i int) *ObjectClass
	// RemoveSelection removes the specified child of the object from the
	// object's selection.
	RemoveSelection(i int) bool
	// SelectAllSelection causes every child of the object to be selected if the
	// object supports multiple selections.
	SelectAllSelection() bool
}

var _ Selectioner = (*Selection)(nil)

func ifaceInitSelectioner(gifacePtr, data C.gpointer) {
	iface := (*C.AtkSelectionIface)(unsafe.Pointer(gifacePtr))
	iface.add_selection = (*[0]byte)(C._gotk4_atk1_SelectionIface_add_selection)
	iface.clear_selection = (*[0]byte)(C._gotk4_atk1_SelectionIface_clear_selection)
	iface.get_selection_count = (*[0]byte)(C._gotk4_atk1_SelectionIface_get_selection_count)
	iface.is_child_selected = (*[0]byte)(C._gotk4_atk1_SelectionIface_is_child_selected)
	iface.ref_selection = (*[0]byte)(C._gotk4_atk1_SelectionIface_ref_selection)
	iface.remove_selection = (*[0]byte)(C._gotk4_atk1_SelectionIface_remove_selection)
	iface.select_all_selection = (*[0]byte)(C._gotk4_atk1_SelectionIface_select_all_selection)
	iface.selection_changed = (*[0]byte)(C._gotk4_atk1_SelectionIface_selection_changed)
}

//export _gotk4_atk1_SelectionIface_add_selection
func _gotk4_atk1_SelectionIface_add_selection(arg0 *C.AtkSelection, arg1 C.gint) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SelectionOverrider)

	var _i int // out

	_i = int(arg1)

	ok := iface.AddSelection(_i)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_atk1_SelectionIface_clear_selection
func _gotk4_atk1_SelectionIface_clear_selection(arg0 *C.AtkSelection) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SelectionOverrider)

	ok := iface.ClearSelection()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_atk1_SelectionIface_get_selection_count
func _gotk4_atk1_SelectionIface_get_selection_count(arg0 *C.AtkSelection) (cret C.gint) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SelectionOverrider)

	gint := iface.SelectionCount()

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_SelectionIface_is_child_selected
func _gotk4_atk1_SelectionIface_is_child_selected(arg0 *C.AtkSelection, arg1 C.gint) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SelectionOverrider)

	var _i int // out

	_i = int(arg1)

	ok := iface.IsChildSelected(_i)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_atk1_SelectionIface_ref_selection
func _gotk4_atk1_SelectionIface_ref_selection(arg0 *C.AtkSelection, arg1 C.gint) (cret *C.AtkObject) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SelectionOverrider)

	var _i int // out

	_i = int(arg1)

	object := iface.RefSelection(_i)

	if object != nil {
		cret = (*C.AtkObject)(unsafe.Pointer(externglib.InternObject(object).Native()))
		C.g_object_ref(C.gpointer(externglib.InternObject(object).Native()))
	}

	return cret
}

//export _gotk4_atk1_SelectionIface_remove_selection
func _gotk4_atk1_SelectionIface_remove_selection(arg0 *C.AtkSelection, arg1 C.gint) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SelectionOverrider)

	var _i int // out

	_i = int(arg1)

	ok := iface.RemoveSelection(_i)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_atk1_SelectionIface_select_all_selection
func _gotk4_atk1_SelectionIface_select_all_selection(arg0 *C.AtkSelection) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SelectionOverrider)

	ok := iface.SelectAllSelection()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_atk1_SelectionIface_selection_changed
func _gotk4_atk1_SelectionIface_selection_changed(arg0 *C.AtkSelection) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SelectionOverrider)

	iface.SelectionChanged()
}

func wrapSelection(obj *externglib.Object) *Selection {
	return &Selection{
		Object: obj,
	}
}

func marshalSelection(p uintptr) (interface{}, error) {
	return wrapSelection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_atk1_Selection_ConnectSelectionChanged
func _gotk4_atk1_Selection_ConnectSelectionChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectSelectionChanged: "selection-changed" signal is emitted by an object
// which implements AtkSelection interface when the selection changes.
func (selection *Selection) ConnectSelectionChanged(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(selection, "selection-changed", false, unsafe.Pointer(C._gotk4_atk1_Selection_ConnectSelectionChanged), f)
}

// AddSelection adds the specified accessible child of the object to the
// object's selection.
//
// The function takes the following parameters:
//
//    - i specifying the child index.
//
// The function returns the following values:
//
//    - ok: TRUE if success, FALSE otherwise.
//
func (selection *Selection) AddSelection(i int) bool {
	var _arg0 *C.AtkSelection // out
	var _arg1 C.gint          // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(externglib.InternObject(selection).Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_selection_add_selection(_arg0, _arg1)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(i)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ClearSelection clears the selection in the object so that no children in the
// object are selected.
//
// The function returns the following values:
//
//    - ok: TRUE if success, FALSE otherwise.
//
func (selection *Selection) ClearSelection() bool {
	var _arg0 *C.AtkSelection // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(externglib.InternObject(selection).Native()))

	_cret = C.atk_selection_clear_selection(_arg0)
	runtime.KeepAlive(selection)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectionCount gets the number of accessible children currently selected.
// Note: callers should not rely on NULL or on a zero value for indication of
// whether AtkSelectionIface is implemented, they should use type
// checking/interface checking macros or the atk_get_accessible_value()
// convenience method.
//
// The function returns the following values:
//
//    - gint representing the number of items selected, or 0 if selection does
//      not implement this interface.
//
func (selection *Selection) SelectionCount() int {
	var _arg0 *C.AtkSelection // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(externglib.InternObject(selection).Native()))

	_cret = C.atk_selection_get_selection_count(_arg0)
	runtime.KeepAlive(selection)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IsChildSelected determines if the current child of this object is selected
// Note: callers should not rely on NULL or on a zero value for indication of
// whether AtkSelectionIface is implemented, they should use type
// checking/interface checking macros or the atk_get_accessible_value()
// convenience method.
//
// The function takes the following parameters:
//
//    - i specifying the child index.
//
// The function returns the following values:
//
//    - ok: gboolean representing the specified child is selected, or 0 if
//      selection does not implement this interface.
//
func (selection *Selection) IsChildSelected(i int) bool {
	var _arg0 *C.AtkSelection // out
	var _arg1 C.gint          // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(externglib.InternObject(selection).Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_selection_is_child_selected(_arg0, _arg1)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(i)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RefSelection gets a reference to the accessible object representing the
// specified selected child of the object. Note: callers should not rely on NULL
// or on a zero value for indication of whether AtkSelectionIface is
// implemented, they should use type checking/interface checking macros or the
// atk_get_accessible_value() convenience method.
//
// The function takes the following parameters:
//
//    - i specifying the index in the selection set. (e.g. the ith selection as
//      opposed to the ith child).
//
// The function returns the following values:
//
//    - object (optional) representing the selected accessible, or NULL if
//      selection does not implement this interface.
//
func (selection *Selection) RefSelection(i int) *ObjectClass {
	var _arg0 *C.AtkSelection // out
	var _arg1 C.gint          // out
	var _cret *C.AtkObject    // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(externglib.InternObject(selection).Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_selection_ref_selection(_arg0, _arg1)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(i)

	var _object *ObjectClass // out

	if _cret != nil {
		_object = wrapObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _object
}

// RemoveSelection removes the specified child of the object from the object's
// selection.
//
// The function takes the following parameters:
//
//    - i specifying the index in the selection set. (e.g. the ith selection as
//      opposed to the ith child).
//
// The function returns the following values:
//
//    - ok: TRUE if success, FALSE otherwise.
//
func (selection *Selection) RemoveSelection(i int) bool {
	var _arg0 *C.AtkSelection // out
	var _arg1 C.gint          // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(externglib.InternObject(selection).Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_selection_remove_selection(_arg0, _arg1)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(i)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectAllSelection causes every child of the object to be selected if the
// object supports multiple selections.
//
// The function returns the following values:
//
//    - ok: TRUE if success, FALSE otherwise.
//
func (selection *Selection) SelectAllSelection() bool {
	var _arg0 *C.AtkSelection // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(externglib.InternObject(selection).Native()))

	_cret = C.atk_selection_select_all_selection(_arg0)
	runtime.KeepAlive(selection)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
