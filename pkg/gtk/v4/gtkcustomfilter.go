// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_custom_filter_get_type()), F: marshalCustomFilterer},
	})
}

// CustomFilterFunc: user function that is called to determine if the @item
// should be matched.
//
// If the filter matches the item, this function must return true. If the item
// should be filtered out, false must be returned.
type CustomFilterFunc func(item *externglib.Object, userData interface{}) (ok bool)

//export gotk4_CustomFilterFunc
func gotk4_CustomFilterFunc(arg0 C.gpointer, arg1 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var item *externglib.Object // out
	var userData interface{}    // out

	item = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*externglib.Object)
	userData = box.Get(uintptr(arg1))

	fn := v.(CustomFilterFunc)
	ok := fn(item, userData)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// CustomFilterer describes CustomFilter's methods.
type CustomFilterer interface {
	privateCustomFilter()
}

// CustomFilter: `GtkCustomFilter` determines whether to include items with a
// callback.
type CustomFilter struct {
	Filter
}

var (
	_ CustomFilterer  = (*CustomFilter)(nil)
	_ gextras.Nativer = (*CustomFilter)(nil)
)

func wrapCustomFilter(obj *externglib.Object) CustomFilterer {
	return &CustomFilter{
		Filter: Filter{
			Object: obj,
		},
	}
}

func marshalCustomFilterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCustomFilter(obj), nil
}

func (*CustomFilter) privateCustomFilter() {}
