// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_settings_get_type()), F: marshalSettingser},
	})
}

// Settings: GtkSettings provides a mechanism to share global settings between
// applications.
//
// On the X window system, this sharing is realized by an XSettings
// (http://www.freedesktop.org/wiki/Specifications/xsettings-spec) manager that
// is usually part of the desktop environment, along with utilities that let the
// user change these settings.
//
// On Wayland, the settings are obtained either via a settings portal, or by
// reading desktop settings from DConf.
//
// In the absence of these sharing mechanisms, GTK reads default values for
// settings from settings.ini files in /etc/gtk-4.0, $XDG_CONFIG_DIRS/gtk-4.0
// and $XDG_CONFIG_HOME/gtk-4.0. These files must be valid key files (see
// GKeyFile), and have a section called Settings. Themes can also provide
// default values for settings by installing a settings.ini file next to their
// gtk.css file.
//
// Applications can override system-wide settings by setting the property of the
// GtkSettings object with g_object_set(). This should be restricted to special
// cases though; GtkSettings are not meant as an application configuration
// facility.
//
// There is one GtkSettings instance per display. It can be obtained with
// gtksettings.GetForDisplay, but in many cases, it is more convenient to use
// gtk.Widget.GetSettings().
type Settings struct {
	*externglib.Object

	StyleProvider
}

func wrapSettings(obj *externglib.Object) *Settings {
	return &Settings{
		Object: obj,
		StyleProvider: StyleProvider{
			Object: obj,
		},
	}
}

func marshalSettingser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSettings(obj), nil
}

// ResetProperty undoes the effect of calling g_object_set() to install an
// application-specific value for a setting.
//
// After this call, the setting will again follow the session-wide value for
// this setting.
func (settings *Settings) ResetProperty(name string) {
	var _arg0 *C.GtkSettings // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_settings_reset_property(_arg0, _arg1)
}

// SettingsGetDefault gets the GtkSettings object for the default display,
// creating it if necessary.
//
// See gtk.Settings.GetForDisplay.
func SettingsGetDefault() *Settings {
	var _cret *C.GtkSettings // in

	_cret = C.gtk_settings_get_default()

	var _settings *Settings // out

	if _cret != nil {
		_settings = wrapSettings(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _settings
}

// SettingsGetForDisplay gets the GtkSettings object for display, creating it if
// necessary.
func SettingsGetForDisplay(display *gdk.Display) *Settings {
	var _arg1 *C.GdkDisplay  // out
	var _cret *C.GtkSettings // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gtk_settings_get_for_display(_arg1)

	var _settings *Settings // out

	_settings = wrapSettings(externglib.Take(unsafe.Pointer(_cret)))

	return _settings
}
