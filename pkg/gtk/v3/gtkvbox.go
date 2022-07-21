// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GTypeVBox returns the GType for the type VBox.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeVBox() coreglib.Type {
	gtype := coreglib.Type(C.gtk_vbox_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalVBox)
	return gtype
}

// VBoxOverrider contains methods that are overridable.
type VBoxOverrider interface {
}

// VBox is a container that organizes child widgets into a single column.
//
// Use the Box packing interface to determine the arrangement, spacing, height,
// and alignment of VBox children.
//
// All children are allocated the same width.
//
// GtkVBox has been deprecated. You can use Box with a Orientable:orientation
// set to GTK_ORIENTATION_VERTICAL instead when calling gtk_box_new(), which is
// a very quick and easy change.
//
// If you have derived your own classes from GtkVBox, you can change the
// inheritance to derive directly from Box, and set the Orientable:orientation
// property to GTK_ORIENTATION_VERTICAL in your instance init function, with a
// call like:
//
//    gtk_orientable_set_orientation (GTK_ORIENTABLE (object),
//                                    GTK_ORIENTATION_VERTICAL);
//
// If you have a grid-like layout composed of nested boxes, and you don’t need
// first-child or last-child styling, the recommendation is to switch to Grid.
// For more information about migrating to Grid, see [Migrating from other
// containers to GtkGrid][gtk-migrating-GtkGrid].
type VBox struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer       = (*VBox)(nil)
	_ coreglib.Objector = (*VBox)(nil)
)

func classInitVBoxer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapVBox(obj *coreglib.Object) *VBox {
	return &VBox{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalVBox(p uintptr) (interface{}, error) {
	return wrapVBox(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewVBox creates a new VBox.
//
// Deprecated: You can use gtk_box_new() with GTK_ORIENTATION_VERTICAL instead,
// which is a quick and easy change. But the recommendation is to switch to
// Grid, since Box is going to go away eventually. See [Migrating from other
// containers to GtkGrid][gtk-migrating-GtkGrid].
//
// The function takes the following parameters:
//
//    - homogeneous: TRUE if all children are to be given equal space allotments.
//    - spacing: number of pixels to place by default between children.
//
// The function returns the following values:
//
//    - vBox: new VBox.
//
func NewVBox(homogeneous bool, spacing int32) *VBox {
	var _arg1 C.gboolean   // out
	var _arg2 C.gint       // out
	var _cret *C.GtkWidget // in

	if homogeneous {
		_arg1 = C.TRUE
	}
	_arg2 = C.gint(spacing)

	_cret = C.gtk_vbox_new(_arg1, _arg2)
	runtime.KeepAlive(homogeneous)
	runtime.KeepAlive(spacing)

	var _vBox *VBox // out

	_vBox = wrapVBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _vBox
}
