// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gchar* _gotk4_atk1_PlugClass_get_object_id(void*);
import "C"

// glib.Type values for atkplug.go.
var GTypePlug = coreglib.Type(C.atk_plug_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypePlug, F: marshalPlug},
	})
}

// PlugOverrider contains methods that are overridable.
type PlugOverrider interface {
	// The function returns the following values:
	//
	ObjectID() string
}

// Plug: see Socket.
type Plug struct {
	_ [0]func() // equal guard
	ObjectClass

	*coreglib.Object
	Component
}

var (
	_ coreglib.Objector = (*Plug)(nil)
)

func classInitPlugger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.AtkPlugClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.AtkPlugClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ ObjectID() string }); ok {
		pclass.get_object_id = (*[0]byte)(C._gotk4_atk1_PlugClass_get_object_id)
	}
}

//export _gotk4_atk1_PlugClass_get_object_id
func _gotk4_atk1_PlugClass_get_object_id(arg0 *C.void) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ObjectID() string })

	utf8 := iface.ObjectID()

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

func wrapPlug(obj *coreglib.Object) *Plug {
	return &Plug{
		ObjectClass: ObjectClass{
			Object: obj,
		},
		Object: obj,
		Component: Component{
			Object: obj,
		},
	}
}

func marshalPlug(p uintptr) (interface{}, error) {
	return wrapPlug(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPlug creates a new Plug instance.
//
// The function returns the following values:
//
//    - plug: newly created Plug.
//
func NewPlug() *Plug {
	_gret := girepository.MustFind("Atk", "Plug").InvokeMethod("new_Plug", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _plug *Plug // out

	_plug = wrapPlug(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _plug
}

// ID gets the unique ID of an Plug object, which can be used to embed inside of
// an Socket using atk_socket_embed().
//
// Internally, this calls a class function that should be registered by the IPC
// layer (usually at-spi2-atk). The implementor of an Plug object should call
// this function (after atk-bridge is loaded) and pass the value to the process
// implementing the Socket, so it could embed the plug.
//
// The function returns the following values:
//
//    - utf8: unique ID for the plug.
//
func (plug *Plug) ID() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(plug).Native()))

	_gret := girepository.MustFind("Atk", "Plug").InvokeMethod("get_id", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(plug)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SetChild sets child as accessible child of plug and plug as accessible parent
// of child. child can be NULL.
//
// In some cases, one can not use the AtkPlug type directly as accessible object
// for the toplevel widget of the application. For instance in the gtk case,
// GtkPlugAccessible can not inherit both from GtkWindowAccessible and from
// AtkPlug. In such a case, one can create, in addition to the standard
// accessible object for the toplevel widget, an AtkPlug object, and make the
// former the child of the latter by calling atk_plug_set_child().
//
// The function takes the following parameters:
//
//    - child to be set as accessible child of plug.
//
func (plug *Plug) SetChild(child *ObjectClass) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(plug).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	girepository.MustFind("Atk", "Plug").InvokeMethod("set_child", _args[:], nil)

	runtime.KeepAlive(plug)
	runtime.KeepAlive(child)
}
