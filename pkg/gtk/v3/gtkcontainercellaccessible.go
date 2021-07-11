// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_container_cell_accessible_get_type()), F: marshalContainerCellAccessibler},
	})
}

// ContainerCellAccessibler describes ContainerCellAccessible's methods.
type ContainerCellAccessibler interface {
	AddChild(child CellAccessibler)

	RemoveChild(child CellAccessibler)
}

type ContainerCellAccessible struct {
	CellAccessible
}

var (
	_ ContainerCellAccessibler = (*ContainerCellAccessible)(nil)
	_ gextras.Nativer          = (*ContainerCellAccessible)(nil)
)

func wrapContainerCellAccessible(obj *externglib.Object) ContainerCellAccessibler {
	return &ContainerCellAccessible{
		CellAccessible: CellAccessible{
			Accessible: Accessible{
				ObjectClass: atk.ObjectClass{
					Object: obj,
				},
			},
			Action: atk.Action{
				Object: obj,
			},
			Component: atk.Component{
				Object: obj,
			},
		},
	}
}

func marshalContainerCellAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapContainerCellAccessible(obj), nil
}

func NewContainerCellAccessible() *ContainerCellAccessible {
	var _cret *C.GtkContainerCellAccessible // in

	_cret = C.gtk_container_cell_accessible_new()

	var _containerCellAccessible *ContainerCellAccessible // out

	_containerCellAccessible = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*ContainerCellAccessible)

	return _containerCellAccessible
}

func (container *ContainerCellAccessible) AddChild(child CellAccessibler) {
	var _arg0 *C.GtkContainerCellAccessible // out
	var _arg1 *C.GtkCellAccessible          // out

	_arg0 = (*C.GtkContainerCellAccessible)(unsafe.Pointer(container.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_container_cell_accessible_add_child(_arg0, _arg1)
}

func (container *ContainerCellAccessible) RemoveChild(child CellAccessibler) {
	var _arg0 *C.GtkContainerCellAccessible // out
	var _arg1 *C.GtkCellAccessible          // out

	_arg0 = (*C.GtkContainerCellAccessible)(unsafe.Pointer(container.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_container_cell_accessible_remove_child(_arg0, _arg1)
}
