// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gemblem.go.
var GTypeEmblem = coreglib.Type(C.g_emblem_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeEmblem, F: marshalEmblem},
	})
}

// EmblemOverrider contains methods that are overridable.
type EmblemOverrider interface {
}

// Emblem is an implementation of #GIcon that supports having an emblem, which
// is an icon with additional properties. It can than be added to a Icon.
//
// Currently, only metainformation about the emblem's origin is supported. More
// may be added in the future.
type Emblem struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Icon
}

var (
	_ coreglib.Objector = (*Emblem)(nil)
)

func classInitEmblemmer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapEmblem(obj *coreglib.Object) *Emblem {
	return &Emblem{
		Object: obj,
		Icon: Icon{
			Object: obj,
		},
	}
}

func marshalEmblem(p uintptr) (interface{}, error) {
	return wrapEmblem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewEmblem creates a new emblem for icon.
//
// The function takes the following parameters:
//
//    - icon: GIcon containing the icon.
//
// The function returns the following values:
//
//    - emblem: new #GEmblem.
//
func NewEmblem(icon Iconner) *Emblem {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(icon).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "Emblem").InvokeMethod("new_Emblem", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(icon)

	var _emblem *Emblem // out

	_emblem = wrapEmblem(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _emblem
}

// GetIcon gives back the icon from emblem.
//
// The function returns the following values:
//
//    - icon The returned object belongs to the emblem and should not be modified
//      or freed.
//
func (emblem *Emblem) GetIcon() *Icon {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(emblem).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "Emblem").InvokeMethod("get_icon", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(emblem)

	var _icon *Icon // out

	_icon = wrapIcon(coreglib.Take(unsafe.Pointer(_cret)))

	return _icon
}
