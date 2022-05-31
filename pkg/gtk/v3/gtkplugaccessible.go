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
import "C"

// glib.Type values for gtkplugaccessible.go.
var GTypePlugAccessible = coreglib.Type(C.gtk_plug_accessible_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypePlugAccessible, F: marshalPlugAccessible},
	})
}

// PlugAccessibleOverrider contains methods that are overridable.
type PlugAccessibleOverrider interface {
}

type PlugAccessible struct {
	_ [0]func() // equal guard
	WindowAccessible
}

var (
	_ coreglib.Objector = (*PlugAccessible)(nil)
)

func classInitPlugAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapPlugAccessible(obj *coreglib.Object) *PlugAccessible {
	return &PlugAccessible{
		WindowAccessible: WindowAccessible{
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
			Window: atk.Window{
				ObjectClass: atk.ObjectClass{
					Object: obj,
				},
			},
		},
	}
}

func marshalPlugAccessible(p uintptr) (interface{}, error) {
	return wrapPlugAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function returns the following values:
//
func (plug *PlugAccessible) ID() string {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(plug).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "PlugAccessible").InvokeMethod("get_id", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(plug)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
