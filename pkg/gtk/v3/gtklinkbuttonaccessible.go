// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtklinkbuttonaccessible.go.
var GTypeLinkButtonAccessible = externglib.Type(C.gtk_link_button_accessible_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeLinkButtonAccessible, F: marshalLinkButtonAccessible},
	})
}

// LinkButtonAccessibleOverrider contains methods that are overridable.
type LinkButtonAccessibleOverrider interface {
	externglib.Objector
}

// WrapLinkButtonAccessibleOverrider wraps the LinkButtonAccessibleOverrider
// interface implementation to access the instance methods.
func WrapLinkButtonAccessibleOverrider(obj LinkButtonAccessibleOverrider) *LinkButtonAccessible {
	return wrapLinkButtonAccessible(externglib.BaseObject(obj))
}

type LinkButtonAccessible struct {
	_ [0]func() // equal guard
	ButtonAccessible

	*externglib.Object
	atk.HyperlinkImpl
}

var (
	_ externglib.Objector = (*LinkButtonAccessible)(nil)
)

func classInitLinkButtonAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapLinkButtonAccessible(obj *externglib.Object) *LinkButtonAccessible {
	return &LinkButtonAccessible{
		ButtonAccessible: ButtonAccessible{
			ContainerAccessible: ContainerAccessible{
				WidgetAccessible: WidgetAccessible{
					Accessible: Accessible{
						ObjectClass: atk.ObjectClass{
							Object: obj,
						},
					},
					Component: atk.Component{
						Object: obj,
					},
				},
			},
			Object: obj,
			Action: atk.Action{
				Object: obj,
			},
			Image: atk.Image{
				Object: obj,
			},
		},
		Object: obj,
		HyperlinkImpl: atk.HyperlinkImpl{
			Object: obj,
		},
	}
}

func marshalLinkButtonAccessible(p uintptr) (interface{}, error) {
	return wrapLinkButtonAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
