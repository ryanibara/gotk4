// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

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
