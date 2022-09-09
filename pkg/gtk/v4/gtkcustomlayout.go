// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeCustomLayout = coreglib.Type(C.gtk_custom_layout_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCustomLayout, F: marshalCustomLayout},
	})
}

// CustomLayoutOverrides contains methods that are overridable.
type CustomLayoutOverrides struct {
}

func defaultCustomLayoutOverrides(v *CustomLayout) CustomLayoutOverrides {
	return CustomLayoutOverrides{}
}

// CustomLayout: GtkCustomLayout uses closures for size negotiation.
//
// A GtkCustomLayout uses closures matching to the old GtkWidget virtual
// functions for size negotiation, as a convenience API to ease the porting
// towards the corresponding `GtkLayoutManager virtual functions.
type CustomLayout struct {
	_ [0]func() // equal guard
	LayoutManager
}

var (
	_ LayoutManagerer = (*CustomLayout)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*CustomLayout, *CustomLayoutClass, CustomLayoutOverrides](
		GTypeCustomLayout,
		initCustomLayoutClass,
		wrapCustomLayout,
		defaultCustomLayoutOverrides,
	)
}

func initCustomLayoutClass(gclass unsafe.Pointer, overrides CustomLayoutOverrides, classInitFunc func(*CustomLayoutClass)) {
	if classInitFunc != nil {
		class := (*CustomLayoutClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapCustomLayout(obj *coreglib.Object) *CustomLayout {
	return &CustomLayout{
		LayoutManager: LayoutManager{
			Object: obj,
		},
	}
}

func marshalCustomLayout(p uintptr) (interface{}, error) {
	return wrapCustomLayout(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CustomLayoutClass: instance of this type is always passed by reference.
type CustomLayoutClass struct {
	*customLayoutClass
}

// customLayoutClass is the struct that's finalized.
type customLayoutClass struct {
	native *C.GtkCustomLayoutClass
}

func (c *CustomLayoutClass) ParentClass() *LayoutManagerClass {
	valptr := &c.native.parent_class
	var _v *LayoutManagerClass // out
	_v = (*LayoutManagerClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
