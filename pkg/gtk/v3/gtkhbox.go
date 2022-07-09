// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeHBox returns the GType for the type HBox.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeHBox() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "HBox").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalHBox)
	return gtype
}

// HBoxOverrider contains methods that are overridable.
type HBoxOverrider interface {
}

// HBox is a container that organizes child widgets into a single row.
//
// Use the Box packing interface to determine the arrangement, spacing, width,
// and alignment of HBox children.
//
// All children are allocated the same height.
//
// GtkHBox has been deprecated. You can use Box instead, which is a very quick
// and easy change. If you have derived your own classes from GtkHBox, you can
// simply change the inheritance to derive directly from Box. No further changes
// are needed, since the default value of the Orientable:orientation property is
// GTK_ORIENTATION_HORIZONTAL.
//
// If you have a grid-like layout composed of nested boxes, and you don’t need
// first-child or last-child styling, the recommendation is to switch to Grid.
// For more information about migrating to Grid, see [Migrating from other
// containers to GtkGrid][gtk-migrating-GtkGrid].
type HBox struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer       = (*HBox)(nil)
	_ coreglib.Objector = (*HBox)(nil)
)

func classInitHBoxer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapHBox(obj *coreglib.Object) *HBox {
	return &HBox{
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

func marshalHBox(p uintptr) (interface{}, error) {
	return wrapHBox(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewHBox creates a new HBox.
//
// Deprecated: You can use gtk_box_new() with GTK_ORIENTATION_HORIZONTAL
// instead, which is a quick and easy change. But the recommendation is to
// switch to Grid, since Box is going to go away eventually. See [Migrating from
// other containers to GtkGrid][gtk-migrating-GtkGrid].
//
// The function takes the following parameters:
//
//    - homogeneous: TRUE if all children are to be given equal space allotments.
//    - spacing: number of pixels to place by default between children.
//
// The function returns the following values:
//
//    - hBox: new HBox.
//
func NewHBox(homogeneous bool, spacing int32) *HBox {
	var _args [2]girepository.Argument

	if homogeneous {
		*(*C.gboolean)(unsafe.Pointer(&_args[0])) = C.TRUE
	}
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(spacing)

	_info := girepository.MustFind("Gtk", "HBox")
	_gret := _info.InvokeClassMethod("new_HBox", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(homogeneous)
	runtime.KeepAlive(spacing)

	var _hBox *HBox // out

	_hBox = wrapHBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _hBox
}
