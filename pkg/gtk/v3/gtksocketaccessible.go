// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
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
	GTypeSocketAccessible = coreglib.Type(C.gtk_socket_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSocketAccessible, F: marshalSocketAccessible},
	})
}

// SocketAccessibleOverrides contains methods that are overridable.
type SocketAccessibleOverrides struct {
}

func defaultSocketAccessibleOverrides(v *SocketAccessible) SocketAccessibleOverrides {
	return SocketAccessibleOverrides{}
}

type SocketAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible
}

var (
	_ coreglib.Objector = (*SocketAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*SocketAccessible, *SocketAccessibleClass, SocketAccessibleOverrides](
		GTypeSocketAccessible,
		initSocketAccessibleClass,
		wrapSocketAccessible,
		defaultSocketAccessibleOverrides,
	)
}

func initSocketAccessibleClass(gclass unsafe.Pointer, overrides SocketAccessibleOverrides, classInitFunc func(*SocketAccessibleClass)) {
	if classInitFunc != nil {
		class := (*SocketAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSocketAccessible(obj *coreglib.Object) *SocketAccessible {
	return &SocketAccessible{
		ContainerAccessible: ContainerAccessible{
			WidgetAccessible: WidgetAccessible{
				Accessible: Accessible{
					AtkObject: atk.AtkObject{
						Object: obj,
					},
				},
				Component: atk.Component{
					Object: obj,
				},
			},
		},
	}
}

func marshalSocketAccessible(p uintptr) (interface{}, error) {
	return wrapSocketAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function takes the following parameters:
//
func (socket *SocketAccessible) Embed(path string) {
	var _arg0 *C.GtkSocketAccessible // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GtkSocketAccessible)(unsafe.Pointer(coreglib.InternObject(socket).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_socket_accessible_embed(_arg0, _arg1)
	runtime.KeepAlive(socket)
	runtime.KeepAlive(path)
}

// SocketAccessibleClass: instance of this type is always passed by reference.
type SocketAccessibleClass struct {
	*socketAccessibleClass
}

// socketAccessibleClass is the struct that's finalized.
type socketAccessibleClass struct {
	native *C.GtkSocketAccessibleClass
}

func (s *SocketAccessibleClass) ParentClass() *ContainerAccessibleClass {
	valptr := &s.native.parent_class
	var _v *ContainerAccessibleClass // out
	_v = (*ContainerAccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
