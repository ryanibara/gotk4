// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void callbackDelete(gpointer);
// gboolean _gotk4_gtk4_CustomFilterFunc(gpointer, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_custom_filter_get_type()), F: marshalCustomFilterer},
	})
}

// CustomFilterFunc: user function that is called to determine if the item
// should be matched.
//
// If the filter matches the item, this function must return TRUE. If the item
// should be filtered out, FALSE must be returned.
type CustomFilterFunc func(item *externglib.Object) (ok bool)

//export _gotk4_gtk4_CustomFilterFunc
func _gotk4_gtk4_CustomFilterFunc(arg0 C.gpointer, arg1 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var item *externglib.Object // out

	item = externglib.Take(unsafe.Pointer(arg0))

	fn := v.(CustomFilterFunc)
	ok := fn(item)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// CustomFilter: GtkCustomFilter determines whether to include items with a
// callback.
type CustomFilter struct {
	Filter
}

func wrapCustomFilter(obj *externglib.Object) *CustomFilter {
	return &CustomFilter{
		Filter: Filter{
			Object: obj,
		},
	}
}

func marshalCustomFilterer(p uintptr) (interface{}, error) {
	return wrapCustomFilter(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
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
//    - matchFunc: function to filter items.
//
func NewCustomFilter(matchFunc CustomFilterFunc) *CustomFilter {
	var _arg1 C.GtkCustomFilterFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify
	var _cret *C.GtkCustomFilter // in

	if matchFunc != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk4_CustomFilterFunc)
		_arg2 = C.gpointer(gbox.Assign(matchFunc))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_cret = C.gtk_custom_filter_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(matchFunc)

	var _customFilter *CustomFilter // out

	_customFilter = wrapCustomFilter(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

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
//    - matchFunc: function to filter items.
//
func (self *CustomFilter) SetFilterFunc(matchFunc CustomFilterFunc) {
	var _arg0 *C.GtkCustomFilter    // out
	var _arg1 C.GtkCustomFilterFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkCustomFilter)(unsafe.Pointer(self.Native()))
	if matchFunc != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk4_CustomFilterFunc)
		_arg2 = C.gpointer(gbox.Assign(matchFunc))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_custom_filter_set_filter_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(matchFunc)
}
