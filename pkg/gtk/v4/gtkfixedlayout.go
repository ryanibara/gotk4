// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gsk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeFixedLayout      = coreglib.Type(C.gtk_fixed_layout_get_type())
	GTypeFixedLayoutChild = coreglib.Type(C.gtk_fixed_layout_child_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFixedLayout, F: marshalFixedLayout},
		coreglib.TypeMarshaler{T: GTypeFixedLayoutChild, F: marshalFixedLayoutChild},
	})
}

// FixedLayoutOverrider contains methods that are overridable.
type FixedLayoutOverrider interface {
}

// FixedLayout: GtkFixedLayout is a layout manager which can place child widgets
// at fixed positions.
//
// Most applications should never use this layout manager; fixed positioning and
// sizing requires constant recalculations on where children need to be
// positioned and sized. Other layout managers perform this kind of work
// internally so that application developers don't need to do it. Specifically,
// widgets positioned in a fixed layout manager will need to take into account:
//
// - Themes, which may change widget sizes.
//
// - Fonts other than the one you used to write the app will of course change
// the size of widgets containing text; keep in mind that users may use a larger
// font because of difficulty reading the default, or they may be using a
// different OS that provides different fonts.
//
// - Translation of text into other languages changes its size. Also, display of
// non-English text will use a different font in many cases.
//
// In addition, GtkFixedLayout does not pay attention to text direction and thus
// may produce unwanted results if your app is run under right-to-left languages
// such as Hebrew or Arabic. That is: normally GTK will order containers
// appropriately depending on the text direction, e.g. to put labels to the
// right of the thing they label when using an RTL language; GtkFixedLayout
// won't be able to do that for you.
//
// Finally, fixed positioning makes it kind of annoying to add/remove UI
// elements, since you have to reposition all the other elements. This is a
// long-term maintenance problem for your application.
type FixedLayout struct {
	_ [0]func() // equal guard
	LayoutManager
}

var (
	_ LayoutManagerer = (*FixedLayout)(nil)
)

func classInitFixedLayouter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFixedLayout(obj *coreglib.Object) *FixedLayout {
	return &FixedLayout{
		LayoutManager: LayoutManager{
			Object: obj,
		},
	}
}

func marshalFixedLayout(p uintptr) (interface{}, error) {
	return wrapFixedLayout(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFixedLayout creates a new GtkFixedLayout.
//
// The function returns the following values:
//
//    - fixedLayout: newly created GtkFixedLayout.
//
func NewFixedLayout() *FixedLayout {
	var _cret *C.GtkLayoutManager // in

	_cret = C.gtk_fixed_layout_new()

	var _fixedLayout *FixedLayout // out

	_fixedLayout = wrapFixedLayout(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _fixedLayout
}

// FixedLayoutChildOverrider contains methods that are overridable.
type FixedLayoutChildOverrider interface {
}

// FixedLayoutChild: GtkLayoutChild subclass for children in a GtkFixedLayout.
type FixedLayoutChild struct {
	_ [0]func() // equal guard
	LayoutChild
}

var (
	_ LayoutChilder = (*FixedLayoutChild)(nil)
)

func classInitFixedLayoutChilder(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFixedLayoutChild(obj *coreglib.Object) *FixedLayoutChild {
	return &FixedLayoutChild{
		LayoutChild: LayoutChild{
			Object: obj,
		},
	}
}

func marshalFixedLayoutChild(p uintptr) (interface{}, error) {
	return wrapFixedLayoutChild(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Transform retrieves the transformation of the child.
//
// The function returns the following values:
//
//    - transform (optional): GskTransform.
//
func (child *FixedLayoutChild) Transform() *gsk.Transform {
	var _arg0 *C.GtkFixedLayoutChild // out
	var _cret *C.GskTransform        // in

	_arg0 = (*C.GtkFixedLayoutChild)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_cret = C.gtk_fixed_layout_child_get_transform(_arg0)
	runtime.KeepAlive(child)

	var _transform *gsk.Transform // out

	if _cret != nil {
		_transform = (*gsk.Transform)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.gsk_transform_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_transform)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gsk_transform_unref((*C.GskTransform)(intern.C))
			},
		)
	}

	return _transform
}

// SetTransform sets the transformation of the child of a GtkFixedLayout.
//
// The function takes the following parameters:
//
//    - transform: GskTransform.
//
func (child *FixedLayoutChild) SetTransform(transform *gsk.Transform) {
	var _arg0 *C.GtkFixedLayoutChild // out
	var _arg1 *C.GskTransform        // out

	_arg0 = (*C.GtkFixedLayoutChild)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	_arg1 = (*C.GskTransform)(gextras.StructNative(unsafe.Pointer(transform)))

	C.gtk_fixed_layout_child_set_transform(_arg0, _arg1)
	runtime.KeepAlive(child)
	runtime.KeepAlive(transform)
}
