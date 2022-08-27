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
import "C"

// NewSettings creates a new #GSettings object with the schema specified by
// schema_id.
//
// It is an error for the schema to not exist: schemas are an essential part of
// a program, as they provide type information. If schemas need to be
// dynamically loaded (for example, from an optional runtime dependency),
// g_settings_schema_source_lookup() can be used to test for their existence
// before loading them.
//
// Signals on the newly created #GSettings object will be dispatched via the
// thread-default Context in effect at the time of the call to g_settings_new().
// The new #GSettings will hold a reference on the context. See
// g_main_context_push_thread_default().
//
// The function takes the following parameters:
//
//    - schemaId: id of the schema.
//
// The function returns the following values:
//
//    - settings: new #GSettings object.
//
func NewSettings(schemaId string) *Settings {
	var _arg1 *C.gchar     // out
	var _cret *C.GSettings // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(schemaId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_new(_arg1)
	runtime.KeepAlive(schemaId)

	var _settings *Settings // out

	_settings = wrapSettings(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _settings
}

// NewSettingsWithPath creates a new #GSettings object with the relocatable
// schema specified by schema_id and a given path.
//
// You only need to do this if you want to directly create a settings object
// with a schema that doesn't have a specified path of its own. That's quite
// rare.
//
// It is a programmer error to call this function for a schema that has an
// explicitly specified path.
//
// It is a programmer error if path is not a valid path. A valid path begins and
// ends with '/' and does not contain two consecutive '/' characters.
//
// The function takes the following parameters:
//
//    - schemaId: id of the schema.
//    - path to use.
//
// The function returns the following values:
//
//    - settings: new #GSettings object.
//
func NewSettingsWithPath(schemaId, path string) *Settings {
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _cret *C.GSettings // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(schemaId)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_settings_new_with_path(_arg1, _arg2)
	runtime.KeepAlive(schemaId)
	runtime.KeepAlive(path)

	var _settings *Settings // out

	_settings = wrapSettings(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _settings
}

// Bind: create a binding between the key in the settings object and the
// property property of object.
//
// The binding uses the default GIO mapping functions to map between the
// settings and property values. These functions handle booleans, numeric types
// and string types in a straightforward way. Use g_settings_bind_with_mapping()
// if you need a custom mapping, or map between types that are not supported by
// the default mapping functions.
//
// Unless the flags include G_SETTINGS_BIND_NO_SENSITIVITY, this function also
// establishes a binding between the writability of key and the "sensitive"
// property of object (if object has a boolean property by that name). See
// g_settings_bind_writable() for more details about writable bindings.
//
// Note that the lifecycle of the binding is tied to object, and that you can
// have only one binding per object property. If you bind the same property
// twice on the same object, the second binding overrides the first one.
//
// The function takes the following parameters:
//
//    - key to bind.
//    - object: #GObject.
//    - property: name of the property to bind.
//    - flags for the binding.
//
func (settings *Settings) Bind(key string, object *coreglib.Object, property string, flags SettingsBindFlags) {
	var _arg0 *C.GSettings         // out
	var _arg1 *C.gchar             // out
	var _arg2 C.gpointer           // out
	var _arg3 *C.gchar             // out
	var _arg4 C.GSettingsBindFlags // out

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gpointer(unsafe.Pointer(object.Native()))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(property)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.GSettingsBindFlags(flags)

	C.g_settings_bind(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)
	runtime.KeepAlive(object)
	runtime.KeepAlive(property)
	runtime.KeepAlive(flags)
}

// BindWritable: create a binding between the writability of key in the settings
// object and the property property of object. The property must be boolean;
// "sensitive" or "visible" properties of widgets are the most likely
// candidates.
//
// Writable bindings are always uni-directional; changes of the writability of
// the setting will be propagated to the object property, not the other way.
//
// When the inverted argument is TRUE, the binding inverts the value as it
// passes from the setting to the object, i.e. property will be set to TRUE if
// the key is not writable.
//
// Note that the lifecycle of the binding is tied to object, and that you can
// have only one binding per object property. If you bind the same property
// twice on the same object, the second binding overrides the first one.
//
// The function takes the following parameters:
//
//    - key to bind.
//    - object: #GObject.
//    - property: name of a boolean property to bind.
//    - inverted: whether to 'invert' the value.
//
func (settings *Settings) BindWritable(key string, object *coreglib.Object, property string, inverted bool) {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 C.gpointer   // out
	var _arg3 *C.gchar     // out
	var _arg4 C.gboolean   // out

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gpointer(unsafe.Pointer(object.Native()))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(property)))
	defer C.free(unsafe.Pointer(_arg3))
	if inverted {
		_arg4 = C.TRUE
	}

	C.g_settings_bind_writable(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)
	runtime.KeepAlive(object)
	runtime.KeepAlive(property)
	runtime.KeepAlive(inverted)
}

// Delay changes the #GSettings object into 'delay-apply' mode. In this mode,
// changes to settings are not immediately propagated to the backend, but kept
// locally until g_settings_apply() is called.
func (settings *Settings) Delay() {
	var _arg0 *C.GSettings // out

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))

	C.g_settings_delay(_arg0)
	runtime.KeepAlive(settings)
}

// Boolean gets the value that is stored at key in settings.
//
// A convenience variant of g_settings_get() for booleans.
//
// It is a programmer error to give a key that isn't specified as having a
// boolean type in the schema for settings.
//
// The function takes the following parameters:
//
//    - key to get the value for.
//
// The function returns the following values:
//
//    - ok: boolean.
//
func (settings *Settings) Boolean(key string) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_boolean(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Child creates a child settings object which has a base path of
// base-path/name, where base-path is the base path of settings.
//
// The schema for the child settings object must have been declared in the
// schema of settings using a <child> element.
//
// The function takes the following parameters:
//
//    - name of the child schema.
//
// The function returns the following values:
//
//    - ret: 'child' settings object.
//
func (settings *Settings) Child(name string) *Settings {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret *C.GSettings // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_child(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(name)

	var _ret *Settings // out

	_ret = wrapSettings(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _ret
}

// Double gets the value that is stored at key in settings.
//
// A convenience variant of g_settings_get() for doubles.
//
// It is a programmer error to give a key that isn't specified as having a
// 'double' type in the schema for settings.
//
// The function takes the following parameters:
//
//    - key to get the value for.
//
// The function returns the following values:
//
//    - gdouble: double.
//
func (settings *Settings) Double(key string) float64 {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret C.gdouble    // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_double(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Enum gets the value that is stored in settings for key and converts it to the
// enum value that it represents.
//
// In order to use this function the type of the value must be a string and it
// must be marked in the schema file as an enumerated type.
//
// It is a programmer error to give a key that isn't contained in the schema for
// settings or is not marked as an enumerated type.
//
// If the value stored in the configuration database is not a valid value for
// the enumerated type then this function will return the default value.
//
// The function takes the following parameters:
//
//    - key to get the value for.
//
// The function returns the following values:
//
//    - gint: enum value.
//
func (settings *Settings) Enum(key string) int {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret C.gint       // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_enum(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Flags gets the value that is stored in settings for key and converts it to
// the flags value that it represents.
//
// In order to use this function the type of the value must be an array of
// strings and it must be marked in the schema file as a flags type.
//
// It is a programmer error to give a key that isn't contained in the schema for
// settings or is not marked as a flags type.
//
// If the value stored in the configuration database is not a valid value for
// the flags type then this function will return the default value.
//
// The function takes the following parameters:
//
//    - key to get the value for.
//
// The function returns the following values:
//
//    - guint flags value.
//
func (settings *Settings) Flags(key string) uint {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret C.guint      // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_flags(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// HasUnapplied returns whether the #GSettings object has any unapplied changes.
// This can only be the case if it is in 'delayed-apply' mode.
//
// The function returns the following values:
//
//    - ok: TRUE if settings has unapplied changes.
//
func (settings *Settings) HasUnapplied() bool {
	var _arg0 *C.GSettings // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))

	_cret = C.g_settings_get_has_unapplied(_arg0)
	runtime.KeepAlive(settings)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Int gets the value that is stored at key in settings.
//
// A convenience variant of g_settings_get() for 32-bit integers.
//
// It is a programmer error to give a key that isn't specified as having a int32
// type in the schema for settings.
//
// The function takes the following parameters:
//
//    - key to get the value for.
//
// The function returns the following values:
//
//    - gint: integer.
//
func (settings *Settings) Int(key string) int {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret C.gint       // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_int(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// String gets the value that is stored at key in settings.
//
// A convenience variant of g_settings_get() for strings.
//
// It is a programmer error to give a key that isn't specified as having a
// string type in the schema for settings.
//
// The function takes the following parameters:
//
//    - key to get the value for.
//
// The function returns the following values:
//
//    - utf8: newly-allocated string.
//
func (settings *Settings) String(key string) string {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_string(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Strv: convenience variant of g_settings_get() for string arrays.
//
// It is a programmer error to give a key that isn't specified as having an
// array of strings type in the schema for settings.
//
// The function takes the following parameters:
//
//    - key to get the value for.
//
// The function returns the following values:
//
//    - utf8s: a newly-allocated, NULL-terminated array of strings, the value
//      that is stored at key in settings.
//
func (settings *Settings) Strv(key string) []string {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret **C.gchar    // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_strv(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// Value gets the value that is stored in settings for key.
//
// It is a programmer error to give a key that isn't contained in the schema for
// settings.
//
// The function takes the following parameters:
//
//    - key to get the value for.
//
// The function returns the following values:
//
//    - variant: new #GVariant.
//
func (settings *Settings) Value(key string) *glib.Variant {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret *C.GVariant  // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_value(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	return _variant
}

// IsWritable finds out if a key can be written or not.
//
// The function takes the following parameters:
//
//    - name of a key.
//
// The function returns the following values:
//
//    - ok: TRUE if the key name is writable.
//
func (settings *Settings) IsWritable(name string) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_is_writable(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(name)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetBoolean sets key in settings to value.
//
// A convenience variant of g_settings_set() for booleans.
//
// It is a programmer error to give a key that isn't specified as having a
// boolean type in the schema for settings.
//
// The function takes the following parameters:
//
//    - key: name of the key to set.
//    - value to set it to.
//
// The function returns the following values:
//
//    - ok: TRUE if setting the key succeeded, FALSE if the key was not writable.
//
func (settings *Settings) SetBoolean(key string, value bool) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 C.gboolean   // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	if value {
		_arg2 = C.TRUE
	}

	_cret = C.g_settings_set_boolean(_arg0, _arg1, _arg2)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDouble sets key in settings to value.
//
// A convenience variant of g_settings_set() for doubles.
//
// It is a programmer error to give a key that isn't specified as having a
// 'double' type in the schema for settings.
//
// The function takes the following parameters:
//
//    - key: name of the key to set.
//    - value to set it to.
//
// The function returns the following values:
//
//    - ok: TRUE if setting the key succeeded, FALSE if the key was not writable.
//
func (settings *Settings) SetDouble(key string, value float64) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 C.gdouble    // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gdouble(value)

	_cret = C.g_settings_set_double(_arg0, _arg1, _arg2)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetInt sets key in settings to value.
//
// A convenience variant of g_settings_set() for 32-bit integers.
//
// It is a programmer error to give a key that isn't specified as having a int32
// type in the schema for settings.
//
// The function takes the following parameters:
//
//    - key: name of the key to set.
//    - value to set it to.
//
// The function returns the following values:
//
//    - ok: TRUE if setting the key succeeded, FALSE if the key was not writable.
//
func (settings *Settings) SetInt(key string, value int) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 C.gint       // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(value)

	_cret = C.g_settings_set_int(_arg0, _arg1, _arg2)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetString sets key in settings to value.
//
// A convenience variant of g_settings_set() for strings.
//
// It is a programmer error to give a key that isn't specified as having a
// string type in the schema for settings.
//
// The function takes the following parameters:
//
//    - key: name of the key to set.
//    - value to set it to.
//
// The function returns the following values:
//
//    - ok: TRUE if setting the key succeeded, FALSE if the key was not writable.
//
func (settings *Settings) SetString(key, value string) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_settings_set_string(_arg0, _arg1, _arg2)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetStrv sets key in settings to value.
//
// A convenience variant of g_settings_set() for string arrays. If value is
// NULL, then key is set to be the empty array.
//
// It is a programmer error to give a key that isn't specified as having an
// array of strings type in the schema for settings.
//
// The function takes the following parameters:
//
//    - key: name of the key to set.
//    - value (optional) to set it to, or NULL.
//
// The function returns the following values:
//
//    - ok: TRUE if setting the key succeeded, FALSE if the key was not writable.
//
func (settings *Settings) SetStrv(key string, value []string) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 **C.gchar    // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		_arg2 = (**C.gchar)(C.calloc(C.size_t((len(value) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(value)+1)
			var zero *C.gchar
			out[len(value)] = zero
			for i := range value {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(value[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.g_settings_set_strv(_arg0, _arg1, _arg2)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetValue sets key in settings to value.
//
// It is a programmer error to give a key that isn't contained in the schema for
// settings or for value to have the incorrect type, per the schema.
//
// If value is floating then this function consumes the reference.
//
// The function takes the following parameters:
//
//    - key: name of the key to set.
//    - value of the correct type.
//
// The function returns the following values:
//
//    - ok: TRUE if setting the key succeeded, FALSE if the key was not writable.
//
func (settings *Settings) SetValue(key string, value *glib.Variant) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 *C.GVariant  // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))

	_cret = C.g_settings_set_value(_arg0, _arg1, _arg2)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SettingsListSchemas: deprecated.
//
// Deprecated: Use g_settings_schema_source_list_schemas() instead. If you used
// g_settings_list_schemas() to check for the presence of a particular schema,
// use g_settings_schema_source_lookup() instead of your whole loop.
//
// The function returns the following values:
//
//    - utf8s: list of #GSettings schemas that are available, in no defined
//      order. The list must not be modified or freed.
//
func SettingsListSchemas() []string {
	var _cret **C.gchar // in

	_cret = C.g_settings_list_schemas()

	var _utf8s []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// SettingsUnbind removes an existing binding for property on object.
//
// Note that bindings are automatically removed when the object is finalized, so
// it is rarely necessary to call this function.
//
// The function takes the following parameters:
//
//    - object: object.
//    - property whose binding is removed.
//
func SettingsUnbind(object *coreglib.Object, property string) {
	var _arg1 C.gpointer // out
	var _arg2 *C.gchar   // out

	_arg1 = C.gpointer(unsafe.Pointer(object.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(property)))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_settings_unbind(_arg1, _arg2)
	runtime.KeepAlive(object)
	runtime.KeepAlive(property)
}
