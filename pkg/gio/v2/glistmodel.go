// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gpointer _gotk4_gio2_ListModelInterface_get_item(void*, guint);
// extern guint _gotk4_gio2_ListModelInterface_get_n_items(void*);
// extern void _gotk4_gio2_ListModel_ConnectItemsChanged(gpointer, guint, guint, guint, guintptr);
import "C"

// GTypeListModel returns the GType for the type ListModel.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeListModel() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "ListModel").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalListModel)
	return gtype
}

// ListModelOverrider contains methods that are overridable.
type ListModelOverrider interface {
	// Item: get the item at position. If position is greater than the number of
	// items in list, NULL is returned.
	//
	// NULL is never returned for an index that is smaller than the length of
	// the list. See g_list_model_get_n_items().
	//
	// The function takes the following parameters:
	//
	//    - position of the item to fetch.
	//
	// The function returns the following values:
	//
	//    - object (optional) at position.
	//
	Item(position uint32) *coreglib.Object
	// NItems gets the number of items in list.
	//
	// Depending on the model implementation, calling this function may be less
	// efficient than iterating the list with increasing values for position
	// until g_list_model_get_item() returns NULL.
	//
	// The function returns the following values:
	//
	//    - guint: number of items in list.
	//
	NItems() uint32
}

// ListModel is an interface that represents a mutable list of #GObjects. Its
// main intention is as a model for various widgets in user interfaces, such as
// list views, but it can also be used as a convenient method of returning lists
// of data, with support for updates.
//
// Each object in the list may also report changes in itself via some mechanism
// (normally the #GObject::notify signal). Taken together with the
// Model::items-changed signal, this provides for a list that can change its
// membership, and in which the members can change their individual properties.
//
// A good example would be the list of visible wireless network access points,
// where each access point can report dynamic properties such as signal
// strength.
//
// It is important to note that the Model itself does not report changes to the
// individual items. It only reports changes to the list membership. If you want
// to observe changes to the objects themselves then you need to connect signals
// to the objects that you are interested in.
//
// All items in a Model are of (or derived from) the same type.
// g_list_model_get_item_type() returns that type. The type may be an interface,
// in which case all objects in the list must implement it.
//
// The semantics are close to that of an array: g_list_model_get_n_items()
// returns the number of items in the list and g_list_model_get_item() returns
// an item at a (0-based) position. In order to allow implementations to
// calculate the list length lazily, you can also iterate over items: starting
// from 0, repeatedly call g_list_model_get_item() until it returns NULL.
//
// An implementation may create objects lazily, but must take care to return the
// same object for a given position until all references to it are gone.
//
// On the other side, a consumer is expected only to hold references on objects
// that are currently "user visible", in order to facilitate the maximum level
// of laziness in the implementation of the list and to reduce the required
// number of signal connections at a given time.
//
// This interface is intended only to be used from a single thread. The thread
// in which it is appropriate to use it depends on the particular
// implementation, but typically it will be from the thread that owns the
// [thread-default main context][g-main-context-push-thread-default] in effect
// at the time that the model was created.
//
// ListModel wraps an interface. This means the user can get the
// underlying type by calling Cast().
type ListModel struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ListModel)(nil)
)

// ListModeller describes ListModel's interface methods.
type ListModeller interface {
	coreglib.Objector

	// NItems gets the number of items in list.
	NItems() uint32
	// Item: get the item at position.
	Item(position uint32) *coreglib.Object
	// ItemsChanged emits the Model::items-changed signal on list.
	ItemsChanged(position, removed, added uint32)

	// Items-changed: this signal is emitted whenever items were added to or
	// removed from list.
	ConnectItemsChanged(func(position, removed, added uint32)) coreglib.SignalHandle
}

var _ ListModeller = (*ListModel)(nil)

func ifaceInitListModeller(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Gio", "ListModelInterface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_item"))) = unsafe.Pointer(C._gotk4_gio2_ListModelInterface_get_item)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_n_items"))) = unsafe.Pointer(C._gotk4_gio2_ListModelInterface_get_n_items)
}

//export _gotk4_gio2_ListModelInterface_get_item
func _gotk4_gio2_ListModelInterface_get_item(arg0 *C.void, arg1 C.guint) (cret C.gpointer) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ListModelOverrider)

	var _position uint32 // out

	_position = uint32(arg1)

	object := iface.Item(_position)

	cret = C.gpointer(unsafe.Pointer(object.Native()))
	C.g_object_ref(C.gpointer(object.Native()))

	return cret
}

//export _gotk4_gio2_ListModelInterface_get_n_items
func _gotk4_gio2_ListModelInterface_get_n_items(arg0 *C.void) (cret C.guint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ListModelOverrider)

	guint := iface.NItems()

	cret = C.guint(guint)

	return cret
}

func wrapListModel(obj *coreglib.Object) *ListModel {
	return &ListModel{
		Object: obj,
	}
}

func marshalListModel(p uintptr) (interface{}, error) {
	return wrapListModel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_ListModel_ConnectItemsChanged
func _gotk4_gio2_ListModel_ConnectItemsChanged(arg0 C.gpointer, arg1 C.guint, arg2 C.guint, arg3 C.guint, arg4 C.guintptr) {
	var f func(position, removed, added uint32)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(position, removed, added uint32))
	}

	var _position uint32 // out
	var _removed uint32  // out
	var _added uint32    // out

	_position = uint32(arg1)
	_removed = uint32(arg2)
	_added = uint32(arg3)

	f(_position, _removed, _added)
}

// ConnectItemsChanged: this signal is emitted whenever items were added to or
// removed from list. At position, removed items were removed and added items
// were added in their place.
//
// Note: If removed != added, the positions of all later items in the model
// change.
func (list *ListModel) ConnectItemsChanged(f func(position, removed, added uint32)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(list, "items-changed", false, unsafe.Pointer(C._gotk4_gio2_ListModel_ConnectItemsChanged), f)
}

// NItems gets the number of items in list.
//
// Depending on the model implementation, calling this function may be less
// efficient than iterating the list with increasing values for position until
// g_list_model_get_item() returns NULL.
//
// The function returns the following values:
//
//    - guint: number of items in list.
//
func (list *ListModel) NItems() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(list).Native()))

	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(list)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// Item: get the item at position. If position is greater than the number of
// items in list, NULL is returned.
//
// NULL is never returned for an index that is smaller than the length of the
// list. See g_list_model_get_n_items().
//
// The function takes the following parameters:
//
//    - position of the item to fetch.
//
// The function returns the following values:
//
//    - object (optional) at position.
//
func (list *ListModel) Item(position uint32) *coreglib.Object {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(list).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(position)

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(list)
	runtime.KeepAlive(position)

	var _object *coreglib.Object // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_object = coreglib.AssumeOwnership(unsafe.Pointer(_cret))
	}

	return _object
}

// ItemsChanged emits the Model::items-changed signal on list.
//
// This function should only be called by classes implementing Model. It has to
// be called after the internal representation of list has been updated, because
// handlers connected to this signal might query the new state of the list.
//
// Implementations must only make changes to the model (as visible to its
// consumer) in places that will not cause problems for that consumer. For
// models that are driven directly by a write API (such as Store), changes can
// be reported in response to uses of that API. For models that represent remote
// data, changes should only be made from a fresh mainloop dispatch. It is
// particularly not permitted to make changes in response to a call to the Model
// consumer API.
//
// Stated another way: in general, it is assumed that code making a series of
// accesses to the model via the API, without returning to the mainloop, and
// without calling other code, will continue to view the same contents of the
// model.
//
// The function takes the following parameters:
//
//    - position at which list changed.
//    - removed: number of items removed.
//    - added: number of items added.
//
func (list *ListModel) ItemsChanged(position, removed, added uint32) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(list).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(position)
	*(*C.guint)(unsafe.Pointer(&_args[2])) = C.guint(removed)
	*(*C.guint)(unsafe.Pointer(&_args[3])) = C.guint(added)

	runtime.KeepAlive(list)
	runtime.KeepAlive(position)
	runtime.KeepAlive(removed)
	runtime.KeepAlive(added)
}
