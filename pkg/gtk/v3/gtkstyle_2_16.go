// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// StyleProperty queries the value of a style property corresponding to a widget
// class is in the given style.
//
// The function takes the following parameters:
//
//    - widgetType of a descendant of Widget.
//    - propertyName: name of the style property to get.
//
// The function returns the following values:
//
//    - value where the value of the property being queried will be stored.
//
func (style *Style) StyleProperty(widgetType coreglib.Type, propertyName string) coreglib.Value {
	var _arg0 *C.GtkStyle // out
	var _arg1 C.GType     // out
	var _arg2 *C.gchar    // out
	var _arg3 C.GValue    // in

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(coreglib.InternObject(style).Native()))
	_arg1 = C.GType(widgetType)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(propertyName)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_style_get_style_property(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(style)
	runtime.KeepAlive(widgetType)
	runtime.KeepAlive(propertyName)

	var _value coreglib.Value // out

	_value = *coreglib.ValueFromNative(unsafe.Pointer((&_arg3)))

	return _value
}
