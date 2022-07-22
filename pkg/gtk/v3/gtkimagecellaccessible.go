// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeImageCellAccessible = coreglib.Type(C.gtk_image_cell_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeImageCellAccessible, F: marshalImageCellAccessible},
	})
}

// ImageCellAccessibleOverrider contains methods that are overridable.
type ImageCellAccessibleOverrider interface {
}

type ImageCellAccessible struct {
	_ [0]func() // equal guard
	RendererCellAccessible

	*coreglib.Object
	atk.AtkObject
	atk.Image
}

var (
	_ coreglib.Objector = (*ImageCellAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:     GTypeImageCellAccessible,
		GoType:    reflect.TypeOf((*ImageCellAccessible)(nil)),
		InitClass: initClassImageCellAccessible,
	})
}

func initClassImageCellAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		InitImageCellAccessible(*ImageCellAccessibleClass)
	}); ok {
		klass := (*ImageCellAccessibleClass)(gextras.NewStructNative(gclass))
		goval.InitImageCellAccessible(klass)
	}
}

func wrapImageCellAccessible(obj *coreglib.Object) *ImageCellAccessible {
	return &ImageCellAccessible{
		RendererCellAccessible: RendererCellAccessible{
			CellAccessible: CellAccessible{
				Accessible: Accessible{
					AtkObject: atk.AtkObject{
						Object: obj,
					},
				},
				Object: obj,
				Action: atk.Action{
					Object: obj,
				},
				AtkObject: atk.AtkObject{
					Object: obj,
				},
				Component: atk.Component{
					Object: obj,
				},
				TableCell: atk.TableCell{
					AtkObject: atk.AtkObject{
						Object: obj,
					},
				},
			},
		},
		Object: obj,
		AtkObject: atk.AtkObject{
			Object: obj,
		},
		Image: atk.Image{
			Object: obj,
		},
	}
}

func marshalImageCellAccessible(p uintptr) (interface{}, error) {
	return wrapImageCellAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ImageCellAccessibleClass: instance of this type is always passed by
// reference.
type ImageCellAccessibleClass struct {
	*imageCellAccessibleClass
}

// imageCellAccessibleClass is the struct that's finalized.
type imageCellAccessibleClass struct {
	native *C.GtkImageCellAccessibleClass
}

func (i *ImageCellAccessibleClass) ParentClass() *RendererCellAccessibleClass {
	valptr := &i.native.parent_class
	var v *RendererCellAccessibleClass // out
	v = (*RendererCellAccessibleClass)(gextras.NewStructNative(unsafe.Pointer((&*valptr))))
	return v
}
