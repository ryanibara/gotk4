// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_accessible_get_type()), F: marshalCellAccessibler},
	})
}

// CellAccessibleOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CellAccessibleOverrider interface {
	// The function takes the following parameters:
	//
	UpdateCache(emitSignal bool)
}

type CellAccessible struct {
	_ [0]func() // equal guard
	Accessible

	*externglib.Object
	atk.Action
	atk.Component
	atk.ObjectClass
	atk.TableCell
}

var (
	_ externglib.Objector = (*CellAccessible)(nil)
)

func wrapCellAccessible(obj *externglib.Object) *CellAccessible {
	return &CellAccessible{
		Accessible: Accessible{
			ObjectClass: atk.ObjectClass{
				Object: obj,
			},
		},
		Object: obj,
		Action: atk.Action{
			Object: obj,
		},
		Component: atk.Component{
			Object: obj,
		},
		ObjectClass: atk.ObjectClass{
			Object: obj,
		},
		TableCell: atk.TableCell{
			ObjectClass: atk.ObjectClass{
				Object: obj,
			},
		},
	}
}

func marshalCellAccessibler(p uintptr) (interface{}, error) {
	return wrapCellAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
