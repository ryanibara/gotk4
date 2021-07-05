// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_settings_get_type()), F: marshalSettings},
	})
}

// Settings provide a mechanism to share global settings between applications.
//
// On the X window system, this sharing is realized by an XSettings
// (http://www.freedesktop.org/wiki/Specifications/xsettings-spec) manager that
// is usually part of the desktop environment, along with utilities that let the
// user change these settings. In the absence of an Xsettings manager, GTK+
// reads default values for settings from `settings.ini` files in
// `/etc/gtk-3.0`, `$XDG_CONFIG_DIRS/gtk-3.0` and `$XDG_CONFIG_HOME/gtk-3.0`.
// These files must be valid key files (see File), and have a section called
// Settings. Themes can also provide default values for settings by installing a
// `settings.ini` file next to their `gtk.css` file.
//
// Applications can override system-wide settings by setting the property of the
// GtkSettings object with g_object_set(). This should be restricted to special
// cases though; GtkSettings are not meant as an application configuration
// facility. When doing so, you need to be aware that settings that are specific
// to individual widgets may not be available before the widget type has been
// realized at least once. The following example demonstrates a way to do this:
//
//      gtk_init (&argc, &argv);
//
//      // make sure the type is realized
//      g_type_class_unref (g_type_class_ref (GTK_TYPE_IMAGE_MENU_ITEM));
//
//      g_object_set (gtk_settings_get_default (), "gtk-enable-animations", FALSE, NULL);
//
// There is one GtkSettings instance per screen. It can be obtained with
// gtk_settings_get_for_screen(), but in many cases, it is more convenient to
// use gtk_widget_get_settings(). gtk_settings_get_default() returns the
// GtkSettings instance for the default screen.
type Settings interface {
	gextras.Objector

	// AsStyleProvider casts the class to the StyleProvider interface.
	AsStyleProvider() StyleProvider

	// ResetPropertySettings undoes the effect of calling g_object_set() to
	// install an application-specific value for a setting. After this call, the
	// setting will again follow the session-wide value for this setting.
	ResetPropertySettings(name string)
	// SetDoublePropertySettings: deprecated: since version 3.16.
	SetDoublePropertySettings(name string, vDouble float64, origin string)
	// SetLongPropertySettings: deprecated: since version 3.16.
	SetLongPropertySettings(name string, vLong int32, origin string)
	// SetStringPropertySettings: deprecated: since version 3.16.
	SetStringPropertySettings(name string, vString string, origin string)
}

// settings implements the Settings class.
type settings struct {
	gextras.Objector
}

// WrapSettings wraps a GObject to the right type. It is
// primarily used internally.
func WrapSettings(obj *externglib.Object) Settings {
	return settings{
		Objector: obj,
	}
}

func marshalSettings(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSettings(obj), nil
}

func (s settings) ResetPropertySettings(name string) {
	var _arg0 *C.GtkSettings // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_settings_reset_property(_arg0, _arg1)
}

func (s settings) SetDoublePropertySettings(name string, vDouble float64, origin string) {
	var _arg0 *C.GtkSettings // out
	var _arg1 *C.gchar       // out
	var _arg2 C.gdouble      // out
	var _arg3 *C.gchar       // out

	_arg0 = (*C.GtkSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gdouble(vDouble)
	_arg3 = (*C.gchar)(C.CString(origin))
	defer C.free(unsafe.Pointer(_arg3))

	C.gtk_settings_set_double_property(_arg0, _arg1, _arg2, _arg3)
}

func (s settings) SetLongPropertySettings(name string, vLong int32, origin string) {
	var _arg0 *C.GtkSettings // out
	var _arg1 *C.gchar       // out
	var _arg2 C.glong        // out
	var _arg3 *C.gchar       // out

	_arg0 = (*C.GtkSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.glong(vLong)
	_arg3 = (*C.gchar)(C.CString(origin))
	defer C.free(unsafe.Pointer(_arg3))

	C.gtk_settings_set_long_property(_arg0, _arg1, _arg2, _arg3)
}

func (s settings) SetStringPropertySettings(name string, vString string, origin string) {
	var _arg0 *C.GtkSettings // out
	var _arg1 *C.gchar       // out
	var _arg2 *C.gchar       // out
	var _arg3 *C.gchar       // out

	_arg0 = (*C.GtkSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(vString))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(origin))
	defer C.free(unsafe.Pointer(_arg3))

	C.gtk_settings_set_string_property(_arg0, _arg1, _arg2, _arg3)
}

func (s settings) AsStyleProvider() StyleProvider {
	return WrapStyleProvider(gextras.InternObject(s))
}
