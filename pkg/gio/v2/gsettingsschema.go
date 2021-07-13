// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_settings_schema_get_type()), F: marshalSettingsSchema},
		{T: externglib.Type(C.g_settings_schema_key_get_type()), F: marshalSettingsSchemaKey},
		{T: externglib.Type(C.g_settings_schema_source_get_type()), F: marshalSettingsSchemaSource},
	})
}

// SettingsSchema and Schema APIs provide a mechanism for advanced control over
// the loading of schemas and a mechanism for introspecting their content.
//
// Plugin loading systems that wish to provide plugins a way to access settings
// face the problem of how to make the schemas for these settings visible to
// GSettings. Typically, a plugin will want to ship the schema along with itself
// and it won't be installed into the standard system directories for schemas.
//
// SchemaSource provides a mechanism for dealing with this by allowing the
// creation of a new 'schema source' from which schemas can be acquired. This
// schema source can then become part of the metadata associated with the plugin
// and queried whenever the plugin requires access to some settings.
//
// Consider the following example:
//
//    {
//      GSettings *settings;
//      gint some_value;
//
//      settings = plugin_get_settings (self, NULL);
//      some_value = g_settings_get_int (settings, "some-value");
//      ...
//    }
//
// It's also possible that the plugin system expects the schema source files
// (ie: .gschema.xml files) instead of a gschemas.compiled file. In that case,
// the plugin loading system must compile the schemas for itself before
// attempting to create the settings source.
type SettingsSchema struct {
	native C.GSettingsSchema
}

func marshalSettingsSchema(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*SettingsSchema)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *SettingsSchema) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// ID: get the ID of schema.
func (schema *SettingsSchema) ID() string {
	var _arg0 *C.GSettingsSchema // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GSettingsSchema)(unsafe.Pointer(schema))

	_cret = C.g_settings_schema_get_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Key gets the key named name from schema.
//
// It is a programmer error to request a key that does not exist. See
// g_settings_schema_list_keys().
func (schema *SettingsSchema) Key(name string) *SettingsSchemaKey {
	var _arg0 *C.GSettingsSchema    // out
	var _arg1 *C.gchar              // out
	var _cret *C.GSettingsSchemaKey // in

	_arg0 = (*C.GSettingsSchema)(unsafe.Pointer(schema))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))

	_cret = C.g_settings_schema_get_key(_arg0, _arg1)

	var _settingsSchemaKey *SettingsSchemaKey // out

	_settingsSchemaKey = (*SettingsSchemaKey)(unsafe.Pointer(_cret))
	C.g_settings_schema_key_ref(_cret)
	runtime.SetFinalizer(_settingsSchemaKey, func(v *SettingsSchemaKey) {
		C.g_settings_schema_key_unref((*C.GSettingsSchemaKey)(unsafe.Pointer(v)))
	})

	return _settingsSchemaKey
}

// Path gets the path associated with schema, or NULL.
//
// Schemas may be single-instance or relocatable. Single-instance schemas
// correspond to exactly one set of keys in the backend database: those located
// at the path returned by this function.
//
// Relocatable schemas can be referenced by other schemas and can therefore
// describe multiple sets of keys at different locations. For relocatable
// schemas, this function will return NULL.
func (schema *SettingsSchema) Path() string {
	var _arg0 *C.GSettingsSchema // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GSettingsSchema)(unsafe.Pointer(schema))

	_cret = C.g_settings_schema_get_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// HasKey checks if schema has a key named name.
func (schema *SettingsSchema) HasKey(name string) bool {
	var _arg0 *C.GSettingsSchema // out
	var _arg1 *C.gchar           // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GSettingsSchema)(unsafe.Pointer(schema))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))

	_cret = C.g_settings_schema_has_key(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ListChildren gets the list of children in schema.
//
// You should free the return value with g_strfreev() when you are done with it.
func (schema *SettingsSchema) ListChildren() []string {
	var _arg0 *C.GSettingsSchema // out
	var _cret **C.gchar

	_arg0 = (*C.GSettingsSchema)(unsafe.Pointer(schema))

	_cret = C.g_settings_schema_list_children(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
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

// ListKeys introspects the list of keys on schema.
//
// You should probably not be calling this function from "normal" code (since
// you should already know what keys are in your schema). This function is
// intended for introspection reasons.
func (schema *SettingsSchema) ListKeys() []string {
	var _arg0 *C.GSettingsSchema // out
	var _cret **C.gchar

	_arg0 = (*C.GSettingsSchema)(unsafe.Pointer(schema))

	_cret = C.g_settings_schema_list_keys(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
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

// Ref: increase the reference count of schema, returning a new reference.
func (schema *SettingsSchema) ref() *SettingsSchema {
	var _arg0 *C.GSettingsSchema // out
	var _cret *C.GSettingsSchema // in

	_arg0 = (*C.GSettingsSchema)(unsafe.Pointer(schema))

	_cret = C.g_settings_schema_ref(_arg0)

	var _settingsSchema *SettingsSchema // out

	_settingsSchema = (*SettingsSchema)(unsafe.Pointer(_cret))
	C.g_settings_schema_ref(_cret)
	runtime.SetFinalizer(_settingsSchema, func(v *SettingsSchema) {
		C.g_settings_schema_unref((*C.GSettingsSchema)(unsafe.Pointer(v)))
	})

	return _settingsSchema
}

// Unref: decrease the reference count of schema, possibly freeing it.
func (schema *SettingsSchema) unref() {
	var _arg0 *C.GSettingsSchema // out

	_arg0 = (*C.GSettingsSchema)(unsafe.Pointer(schema))

	C.g_settings_schema_unref(_arg0)
}

// SettingsSchemaKey is an opaque data structure and can only be accessed using
// the following functions.
type SettingsSchemaKey struct {
	native C.GSettingsSchemaKey
}

func marshalSettingsSchemaKey(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*SettingsSchemaKey)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *SettingsSchemaKey) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// DefaultValue gets the default value for key.
//
// Note that this is the default value according to the schema. System
// administrator defaults and lockdown are not visible via this API.
func (key *SettingsSchemaKey) DefaultValue() *glib.Variant {
	var _arg0 *C.GSettingsSchemaKey // out
	var _cret *C.GVariant           // in

	_arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(key))

	_cret = C.g_settings_schema_key_get_default_value(_arg0)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(unsafe.Pointer(v)))
	})

	return _variant
}

// Description gets the description for key.
//
// If no description has been provided in the schema for key, returns NULL.
//
// The description can be one sentence to several paragraphs in length.
// Paragraphs are delimited with a double newline. Descriptions can be
// translated and the value returned from this function is is the current
// locale.
//
// This function is slow. The summary and description information for the
// schemas is not stored in the compiled schema database so this function has to
// parse all of the source XML files in the schema directory.
func (key *SettingsSchemaKey) Description() string {
	var _arg0 *C.GSettingsSchemaKey // out
	var _cret *C.gchar              // in

	_arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(key))

	_cret = C.g_settings_schema_key_get_description(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Name gets the name of key.
func (key *SettingsSchemaKey) Name() string {
	var _arg0 *C.GSettingsSchemaKey // out
	var _cret *C.gchar              // in

	_arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(key))

	_cret = C.g_settings_schema_key_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Range queries the range of a key.
//
// This function will return a #GVariant that fully describes the range of
// values that are valid for key.
//
// The type of #GVariant returned is (sv). The string describes the type of
// range restriction in effect. The type and meaning of the value contained in
// the variant depends on the string.
//
// If the string is 'type' then the variant contains an empty array. The element
// type of that empty array is the expected type of value and all values of that
// type are valid.
//
// If the string is 'enum' then the variant contains an array enumerating the
// possible values. Each item in the array is a possible valid value and no
// other values are valid.
//
// If the string is 'flags' then the variant contains an array. Each item in the
// array is a value that may appear zero or one times in an array to be used as
// the value for this key. For example, if the variant contained the array ['x',
// 'y'] then the valid values for the key would be [], ['x'], ['y'], ['x', 'y']
// and ['y', 'x'].
//
// Finally, if the string is 'range' then the variant contains a pair of
// like-typed values -- the minimum and maximum permissible values for this key.
//
// This information should not be used by normal programs. It is considered to
// be a hint for introspection purposes. Normal programs should already know
// what is permitted by their own schema. The format may change in any way in
// the future -- but particularly, new forms may be added to the possibilities
// described above.
//
// You should free the returned value with g_variant_unref() when it is no
// longer needed.
func (key *SettingsSchemaKey) Range() *glib.Variant {
	var _arg0 *C.GSettingsSchemaKey // out
	var _cret *C.GVariant           // in

	_arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(key))

	_cret = C.g_settings_schema_key_get_range(_arg0)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(unsafe.Pointer(v)))
	})

	return _variant
}

// Summary gets the summary for key.
//
// If no summary has been provided in the schema for key, returns NULL.
//
// The summary is a short description of the purpose of the key; usually one
// short sentence. Summaries can be translated and the value returned from this
// function is is the current locale.
//
// This function is slow. The summary and description information for the
// schemas is not stored in the compiled schema database so this function has to
// parse all of the source XML files in the schema directory.
func (key *SettingsSchemaKey) Summary() string {
	var _arg0 *C.GSettingsSchemaKey // out
	var _cret *C.gchar              // in

	_arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(key))

	_cret = C.g_settings_schema_key_get_summary(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ValueType gets the Type of key.
func (key *SettingsSchemaKey) ValueType() *glib.VariantType {
	var _arg0 *C.GSettingsSchemaKey // out
	var _cret *C.GVariantType       // in

	_arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(key))

	_cret = C.g_settings_schema_key_get_value_type(_arg0)

	var _variantType *glib.VariantType // out

	_variantType = (*glib.VariantType)(unsafe.Pointer(_cret))

	return _variantType
}

// RangeCheck checks if the given value is of the correct type and within the
// permitted range for key.
//
// It is a programmer error if value is not of the correct type -- you must
// check for this first.
func (key *SettingsSchemaKey) RangeCheck(value *glib.Variant) bool {
	var _arg0 *C.GSettingsSchemaKey // out
	var _arg1 *C.GVariant           // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(key))
	_arg1 = (*C.GVariant)(unsafe.Pointer(value))

	_cret = C.g_settings_schema_key_range_check(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Ref: increase the reference count of key, returning a new reference.
func (key *SettingsSchemaKey) ref() *SettingsSchemaKey {
	var _arg0 *C.GSettingsSchemaKey // out
	var _cret *C.GSettingsSchemaKey // in

	_arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(key))

	_cret = C.g_settings_schema_key_ref(_arg0)

	var _settingsSchemaKey *SettingsSchemaKey // out

	_settingsSchemaKey = (*SettingsSchemaKey)(unsafe.Pointer(_cret))
	C.g_settings_schema_key_ref(_cret)
	runtime.SetFinalizer(_settingsSchemaKey, func(v *SettingsSchemaKey) {
		C.g_settings_schema_key_unref((*C.GSettingsSchemaKey)(unsafe.Pointer(v)))
	})

	return _settingsSchemaKey
}

// Unref: decrease the reference count of key, possibly freeing it.
func (key *SettingsSchemaKey) unref() {
	var _arg0 *C.GSettingsSchemaKey // out

	_arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(key))

	C.g_settings_schema_key_unref(_arg0)
}

// SettingsSchemaSource: this is an opaque structure type. You may not access it
// directly.
type SettingsSchemaSource struct {
	native C.GSettingsSchemaSource
}

func marshalSettingsSchemaSource(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*SettingsSchemaSource)(unsafe.Pointer(b)), nil
}

// NewSettingsSchemaSourceFromDirectory constructs a struct SettingsSchemaSource.
func NewSettingsSchemaSourceFromDirectory(directory string, parent *SettingsSchemaSource, trusted bool) (*SettingsSchemaSource, error) {
	var _arg1 *C.gchar                 // out
	var _arg2 *C.GSettingsSchemaSource // out
	var _arg3 C.gboolean               // out
	var _cret *C.GSettingsSchemaSource // in
	var _cerr *C.GError                // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(directory)))
	_arg2 = (*C.GSettingsSchemaSource)(unsafe.Pointer(parent))
	if trusted {
		_arg3 = C.TRUE
	}

	_cret = C.g_settings_schema_source_new_from_directory(_arg1, _arg2, _arg3, &_cerr)

	var _settingsSchemaSource *SettingsSchemaSource // out
	var _goerr error                                // out

	_settingsSchemaSource = (*SettingsSchemaSource)(unsafe.Pointer(_cret))
	C.g_settings_schema_source_ref(_cret)
	runtime.SetFinalizer(_settingsSchemaSource, func(v *SettingsSchemaSource) {
		C.g_settings_schema_source_unref((*C.GSettingsSchemaSource)(unsafe.Pointer(v)))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _settingsSchemaSource, _goerr
}

// Native returns the underlying C source pointer.
func (s *SettingsSchemaSource) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// ListSchemas lists the schemas in a given source.
//
// If recursive is TRUE then include parent sources. If FALSE then only include
// the schemas from one source (ie: one directory). You probably want TRUE.
//
// Non-relocatable schemas are those for which you can call g_settings_new().
// Relocatable schemas are those for which you must use
// g_settings_new_with_path().
//
// Do not call this function from normal programs. This is designed for use by
// database editors, commandline tools, etc.
func (source *SettingsSchemaSource) ListSchemas(recursive bool) (nonRelocatable []string, relocatable []string) {
	var _arg0 *C.GSettingsSchemaSource // out
	var _arg1 C.gboolean               // out
	var _arg2 **C.gchar
	var _arg3 **C.gchar

	_arg0 = (*C.GSettingsSchemaSource)(unsafe.Pointer(source))
	if recursive {
		_arg1 = C.TRUE
	}

	C.g_settings_schema_source_list_schemas(_arg0, _arg1, &_arg2, &_arg3)

	var _nonRelocatable []string
	var _relocatable []string

	{
		var i int
		var z *C.gchar
		for p := _arg2; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_arg2, i)
		_nonRelocatable = make([]string, i)
		for i := range src {
			_nonRelocatable[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	{
		var i int
		var z *C.gchar
		for p := _arg3; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_arg3, i)
		_relocatable = make([]string, i)
		for i := range src {
			_relocatable[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _nonRelocatable, _relocatable
}

// Lookup looks up a schema with the identifier schema_id in source.
//
// This function is not required for normal uses of #GSettings but it may be
// useful to authors of plugin management systems or to those who want to
// introspect the content of schemas.
//
// If the schema isn't found directly in source and recursive is TRUE then the
// parent sources will also be checked.
//
// If the schema isn't found, NULL is returned.
func (source *SettingsSchemaSource) Lookup(schemaId string, recursive bool) *SettingsSchema {
	var _arg0 *C.GSettingsSchemaSource // out
	var _arg1 *C.gchar                 // out
	var _arg2 C.gboolean               // out
	var _cret *C.GSettingsSchema       // in

	_arg0 = (*C.GSettingsSchemaSource)(unsafe.Pointer(source))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(schemaId)))
	if recursive {
		_arg2 = C.TRUE
	}

	_cret = C.g_settings_schema_source_lookup(_arg0, _arg1, _arg2)

	var _settingsSchema *SettingsSchema // out

	_settingsSchema = (*SettingsSchema)(unsafe.Pointer(_cret))
	C.g_settings_schema_ref(_cret)
	runtime.SetFinalizer(_settingsSchema, func(v *SettingsSchema) {
		C.g_settings_schema_unref((*C.GSettingsSchema)(unsafe.Pointer(v)))
	})

	return _settingsSchema
}

// Ref: increase the reference count of source, returning a new reference.
func (source *SettingsSchemaSource) ref() *SettingsSchemaSource {
	var _arg0 *C.GSettingsSchemaSource // out
	var _cret *C.GSettingsSchemaSource // in

	_arg0 = (*C.GSettingsSchemaSource)(unsafe.Pointer(source))

	_cret = C.g_settings_schema_source_ref(_arg0)

	var _settingsSchemaSource *SettingsSchemaSource // out

	_settingsSchemaSource = (*SettingsSchemaSource)(unsafe.Pointer(_cret))
	C.g_settings_schema_source_ref(_cret)
	runtime.SetFinalizer(_settingsSchemaSource, func(v *SettingsSchemaSource) {
		C.g_settings_schema_source_unref((*C.GSettingsSchemaSource)(unsafe.Pointer(v)))
	})

	return _settingsSchemaSource
}

// Unref: decrease the reference count of source, possibly freeing it.
func (source *SettingsSchemaSource) unref() {
	var _arg0 *C.GSettingsSchemaSource // out

	_arg0 = (*C.GSettingsSchemaSource)(unsafe.Pointer(source))

	C.g_settings_schema_source_unref(_arg0)
}

// SettingsSchemaSourceGetDefault gets the default system schema source.
//
// This function is not required for normal uses of #GSettings but it may be
// useful to authors of plugin management systems or to those who want to
// introspect the content of schemas.
//
// If no schemas are installed, NULL will be returned.
//
// The returned source may actually consist of multiple schema sources from
// different directories, depending on which directories were given in
// XDG_DATA_DIRS and GSETTINGS_SCHEMA_DIR. For this reason, all lookups
// performed against the default source should probably be done recursively.
func SettingsSchemaSourceGetDefault() *SettingsSchemaSource {
	var _cret *C.GSettingsSchemaSource // in

	_cret = C.g_settings_schema_source_get_default()

	var _settingsSchemaSource *SettingsSchemaSource // out

	_settingsSchemaSource = (*SettingsSchemaSource)(unsafe.Pointer(_cret))
	C.g_settings_schema_source_ref(_cret)
	runtime.SetFinalizer(_settingsSchemaSource, func(v *SettingsSchemaSource) {
		C.g_settings_schema_source_unref((*C.GSettingsSchemaSource)(unsafe.Pointer(v)))
	})

	return _settingsSchemaSource
}
