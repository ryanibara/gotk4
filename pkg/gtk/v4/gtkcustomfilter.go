// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gtk4_CustomFilterFunc(gpointer, gpointer);
// extern void callbackDelete(gpointer);
import "C"

// GTypeCustomFilter returns the GType for the type CustomFilter.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCustomFilter() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "CustomFilter").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalCustomFilter)
	return gtype
}

// CustomFilterFunc: user function that is called to determine if the item
// should be matched.
//
// If the filter matches the item, this function must return TRUE. If the item
// should be filtered out, FALSE must be returned.
type CustomFilterFunc func(item *coreglib.Object) (ok bool)

//export _gotk4_gtk4_CustomFilterFunc
func _gotk4_gtk4_CustomFilterFunc(arg1 C.gpointer, arg2 C.gpointer) (cret C.gboolean) {
	var fn CustomFilterFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(CustomFilterFunc)
	}

	var _item *coreglib.Object // out

	_item = coreglib.Take(unsafe.Pointer(arg1))

	ok := fn(_item)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// CustomFilterOverrider contains methods that are overridable.
type CustomFilterOverrider interface {
}

// CustomFilter: GtkCustomFilter determines whether to include items with a
// callback.
type CustomFilter struct {
	_ [0]func() // equal guard
	Filter
}

var (
	_ coreglib.Objector = (*CustomFilter)(nil)
)

func classInitCustomFilterer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapCustomFilter(obj *coreglib.Object) *CustomFilter {
	return &CustomFilter{
		Filter: Filter{
			Object: obj,
		},
	}
}

func marshalCustomFilter(p uintptr) (interface{}, error) {
	return wrapCustomFilter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCustomFilter creates a new filter using the given match_func to filter
// items.
//
// If match_func is NULL, the filter matches all items.
//
// If the filter func changes its filtering behavior, gtk_filter_changed() needs
// to be called.
//
// The function takes the following parameters:
//
//    - matchFunc (optional): function to filter items.
//
// The function returns the following values:
//
//    - customFilter: new GtkCustomFilter.
//
func NewCustomFilter(matchFunc CustomFilterFunc) *CustomFilter {
	var _args [3]girepository.Argument

	if matchFunc != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[0])) = (*[0]byte)(C._gotk4_gtk4_CustomFilterFunc)
		_args[1] = C.gpointer(gbox.Assign(matchFunc))
		_args[2] = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_info := girepository.MustFind("Gtk", "CustomFilter")
	_gret := _info.InvokeClassMethod("new_CustomFilter", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(matchFunc)

	var _customFilter *CustomFilter // out

	_customFilter = wrapCustomFilter(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _customFilter
}

// SetFilterFunc sets the function used for filtering items.
//
// If match_func is NULL, the filter matches all items.
//
// If the filter func changes its filtering behavior, gtk_filter_changed() needs
// to be called.
//
// If a previous function was set, its user_destroy will be called now.
//
// The function takes the following parameters:
//
//    - matchFunc (optional): function to filter items.
//
func (self *CustomFilter) SetFilterFunc(matchFunc CustomFilterFunc) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if matchFunc != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[1])) = (*[0]byte)(C._gotk4_gtk4_CustomFilterFunc)
		_args[2] = C.gpointer(gbox.Assign(matchFunc))
		_args[3] = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_info := girepository.MustFind("Gtk", "CustomFilter")
	_info.InvokeClassMethod("set_filter_func", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(matchFunc)
}
