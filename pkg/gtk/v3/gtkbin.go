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

// GTypeBin returns the GType for the type Bin.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeBin() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Bin").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalBin)
	return gtype
}

// BinOverrider contains methods that are overridable.
type BinOverrider interface {
}

// Bin widget is a container with just one child. It is not very useful itself,
// but it is useful for deriving subclasses, since it provides common code
// needed for handling a single child widget.
//
// Many GTK+ widgets are subclasses of Bin, including Window, Button, Frame,
// HandleBox or ScrolledWindow.
type Bin struct {
	_ [0]func() // equal guard
	Container
}

var (
	_ Containerer = (*Bin)(nil)
)

// Binner describes types inherited from class Bin.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Binner interface {
	coreglib.Objector
	baseBin() *Bin
}

var _ Binner = (*Bin)(nil)

func classInitBinner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapBin(obj *coreglib.Object) *Bin {
	return &Bin{
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
	}
}

func marshalBin(p uintptr) (interface{}, error) {
	return wrapBin(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (bin *Bin) baseBin() *Bin {
	return bin
}

// BaseBin returns the underlying base object.
func BaseBin(obj Binner) *Bin {
	return obj.baseBin()
}

// Child gets the child of the Bin, or NULL if the bin contains no child widget.
// The returned widget does not have a reference added, so you do not need to
// unref it.
//
// The function returns the following values:
//
//    - widget (optional): child of bin, or NULL if it does not have a child.
//
func (bin *Bin) Child() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bin).Native()))

	_info := girepository.MustFind("Gtk", "Bin")
	_gret := _info.InvokeClassMethod("get_child", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bin)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}
