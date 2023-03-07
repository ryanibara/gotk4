// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeBytesIcon      = coreglib.Type(C.g_bytes_icon_get_type())
	GTypePropertyAction = coreglib.Type(C.g_property_action_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeBytesIcon, F: marshalBytesIcon},
		coreglib.TypeMarshaler{T: GTypePropertyAction, F: marshalPropertyAction},
	})
}

// BytesIcon specifies an image held in memory in a common format (usually png)
// to be used as icon.
type BytesIcon struct {
	_ [0]func() // equal guard
	*coreglib.Object

	LoadableIcon
}

var (
	_ coreglib.Objector = (*BytesIcon)(nil)
)

func wrapBytesIcon(obj *coreglib.Object) *BytesIcon {
	return &BytesIcon{
		Object: obj,
		LoadableIcon: LoadableIcon{
			Icon: Icon{
				Object: obj,
			},
		},
	}
}

func marshalBytesIcon(p uintptr) (interface{}, error) {
	return wrapBytesIcon(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBytesIcon creates a new icon for a bytes.
//
// This cannot fail, but loading and interpreting the bytes may fail later on
// (for example, if g_loadable_icon_load() is called) if the image is invalid.
//
// The function takes the following parameters:
//
//    - bytes: #GBytes.
//
// The function returns the following values:
//
//    - bytesIcon for the given bytes.
//
func NewBytesIcon(bytes *glib.Bytes) *BytesIcon {
	var _arg1 *C.GBytes // out
	var _cret *C.GIcon  // in

	_arg1 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(bytes)))

	_cret = C.g_bytes_icon_new(_arg1)
	runtime.KeepAlive(bytes)

	var _bytesIcon *BytesIcon // out

	_bytesIcon = wrapBytesIcon(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _bytesIcon
}

// Bytes gets the #GBytes associated with the given icon.
//
// The function returns the following values:
//
//    - bytes: #GBytes.
//
func (icon *BytesIcon) Bytes() *glib.Bytes {
	var _arg0 *C.GBytesIcon // out
	var _cret *C.GBytes     // in

	_arg0 = (*C.GBytesIcon)(unsafe.Pointer(coreglib.InternObject(icon).Native()))

	_cret = C.g_bytes_icon_get_bytes(_arg0)
	runtime.KeepAlive(icon)

	var _bytes *glib.Bytes // out

	_bytes = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_bytes_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	return _bytes
}

// PropertyAction is a way to get a #GAction with a state value reflecting and
// controlling the value of a #GObject property.
//
// The state of the action will correspond to the value of the property.
// Changing it will change the property (assuming the requested value matches
// the requirements as specified in the Spec).
//
// Only the most common types are presently supported. Booleans are mapped to
// booleans, strings to strings, signed/unsigned integers to int32/uint32 and
// floats and doubles to doubles.
//
// If the property is an enum then the state will be string-typed and conversion
// will automatically be performed between the enum value and "nick" string as
// per the Value table.
//
// Flags types are not currently supported.
//
// Properties of object types, boxed types and pointer types are not supported
// and probably never will be.
//
// Properties of #GVariant types are not currently supported.
//
// If the property is boolean-valued then the action will have a NULL parameter
// type, and activating the action (with no parameter) will toggle the value of
// the property.
//
// In all other cases, the parameter type will correspond to the type of the
// property.
//
// The general idea here is to reduce the number of locations where a particular
// piece of state is kept (and therefore has to be synchronised between). Action
// does not have a separate state that is kept in sync with the property value
// -- its state is the property value.
//
// For example, it might be useful to create a #GAction corresponding to the
// "visible-child-name" property of a Stack so that the current page can be
// switched from a menu. The active radio indication in the menu is then
// directly determined from the active page of the Stack.
//
// An anti-example would be binding the "active-id" property on a ComboBox. This
// is because the state of the combobox itself is probably uninteresting and is
// actually being used to control something else.
//
// Another anti-example would be to bind to the "visible-child-name" property of
// a Stack if this value is actually stored in #GSettings. In that case, the
// real source of the value is #GSettings. If you want a #GAction to control a
// setting stored in #GSettings, see g_settings_create_action() instead, and
// possibly combine its use with g_settings_bind().
type PropertyAction struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Action
}

var (
	_ coreglib.Objector = (*PropertyAction)(nil)
)

func wrapPropertyAction(obj *coreglib.Object) *PropertyAction {
	return &PropertyAction{
		Object: obj,
		Action: Action{
			Object: obj,
		},
	}
}

func marshalPropertyAction(p uintptr) (interface{}, error) {
	return wrapPropertyAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPropertyAction creates a #GAction corresponding to the value of property
// property_name on object.
//
// The property must be existent and readable and writable (and not
// construct-only).
//
// This function takes a reference on object and doesn't release it until the
// action is destroyed.
//
// The function takes the following parameters:
//
//    - name of the action to create.
//    - object that has the property to wrap.
//    - propertyName: name of the property.
//
// The function returns the following values:
//
//    - propertyAction: new Action.
//
func NewPropertyAction(name string, object *coreglib.Object, propertyName string) *PropertyAction {
	var _arg1 *C.gchar           // out
	var _arg2 C.gpointer         // out
	var _arg3 *C.gchar           // out
	var _cret *C.GPropertyAction // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gpointer(unsafe.Pointer(object.Native()))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(propertyName)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.g_property_action_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(name)
	runtime.KeepAlive(object)
	runtime.KeepAlive(propertyName)

	var _propertyAction *PropertyAction // out

	_propertyAction = wrapPropertyAction(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _propertyAction
}