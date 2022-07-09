// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeSettingsSchema returns the GType for the type SettingsSchema.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeSettingsSchema() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "SettingsSchema").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalSettingsSchema)
	return gtype
}

// GTypeSettingsSchemaKey returns the GType for the type SettingsSchemaKey.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeSettingsSchemaKey() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "SettingsSchemaKey").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalSettingsSchemaKey)
	return gtype
}

// GTypeSettingsSchemaSource returns the GType for the type SettingsSchemaSource.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeSettingsSchemaSource() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "SettingsSchemaSource").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalSettingsSchemaSource)
	return gtype
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
//
// An instance of this type is always passed by reference.
type SettingsSchema struct {
	*settingsSchema
}

// settingsSchema is the struct that's finalized.
type settingsSchema struct {
	native unsafe.Pointer
}

func marshalSettingsSchema(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &SettingsSchema{&settingsSchema{(unsafe.Pointer)(b)}}, nil
}

// ID: get the ID of schema.
//
// The function returns the following values:
//
//    - utf8: ID.
//
func (schema *SettingsSchema) ID() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(schema)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(schema)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Key gets the key named name from schema.
//
// It is a programmer error to request a key that does not exist. See
// g_settings_schema_list_keys().
//
// The function takes the following parameters:
//
//    - name of a key.
//
// The function returns the following values:
//
//    - settingsSchemaKey for name.
//
func (schema *SettingsSchema) Key(name string) *SettingsSchemaKey {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(schema)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(schema)
	runtime.KeepAlive(name)

	var _settingsSchemaKey *SettingsSchemaKey // out

	_settingsSchemaKey = (*SettingsSchemaKey)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_settingsSchemaKey)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

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
//
// The function returns the following values:
//
//    - utf8 (optional): path of the schema, or NULL.
//
func (schema *SettingsSchema) Path() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(schema)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(schema)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// HasKey checks if schema has a key named name.
//
// The function takes the following parameters:
//
//    - name of a key.
//
// The function returns the following values:
//
//    - ok: TRUE if such a key exists.
//
func (schema *SettingsSchema) HasKey(name string) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(schema)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(schema)
	runtime.KeepAlive(name)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ListChildren gets the list of children in schema.
//
// You should free the return value with g_strfreev() when you are done with it.
//
// The function returns the following values:
//
//    - utf8s: list of the children on settings, in no defined order.
//
func (schema *SettingsSchema) ListChildren() []string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(schema)))

	_cret = *(***C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(schema)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.void
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

// ListKeys introspects the list of keys on schema.
//
// You should probably not be calling this function from "normal" code (since
// you should already know what keys are in your schema). This function is
// intended for introspection reasons.
//
// The function returns the following values:
//
//    - utf8s: list of the keys on schema, in no defined order.
//
func (schema *SettingsSchema) ListKeys() []string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(schema)))

	_cret = *(***C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(schema)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.void
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

// SettingsSchemaKey is an opaque data structure and can only be accessed using
// the following functions.
//
// An instance of this type is always passed by reference.
type SettingsSchemaKey struct {
	*settingsSchemaKey
}

// settingsSchemaKey is the struct that's finalized.
type settingsSchemaKey struct {
	native unsafe.Pointer
}

func marshalSettingsSchemaKey(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &SettingsSchemaKey{&settingsSchemaKey{(unsafe.Pointer)(b)}}, nil
}

// DefaultValue gets the default value for key.
//
// Note that this is the default value according to the schema. System
// administrator defaults and lockdown are not visible via this API.
//
// The function returns the following values:
//
//    - variant: default value for the key.
//
func (key *SettingsSchemaKey) DefaultValue() *glib.Variant {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(key)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(key)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

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
//
// The function returns the following values:
//
//    - utf8 (optional): description for key, or NULL.
//
func (key *SettingsSchemaKey) Description() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(key)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(key)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Name gets the name of key.
//
// The function returns the following values:
//
//    - utf8: name of key.
//
func (key *SettingsSchemaKey) Name() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(key)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(key)

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
//
// The function returns the following values:
//
//    - variant describing the range.
//
func (key *SettingsSchemaKey) Range() *glib.Variant {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(key)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(key)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

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
//
// The function returns the following values:
//
//    - utf8 (optional): summary for key, or NULL.
//
func (key *SettingsSchemaKey) Summary() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(key)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(key)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ValueType gets the Type of key.
//
// The function returns the following values:
//
//    - variantType: type of key.
//
func (key *SettingsSchemaKey) ValueType() *glib.VariantType {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(key)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(key)

	var _variantType *glib.VariantType // out

	_variantType = (*glib.VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _variantType
}

// RangeCheck checks if the given value is of the correct type and within the
// permitted range for key.
//
// It is a programmer error if value is not of the correct type -- you must
// check for this first.
//
// The function takes the following parameters:
//
//    - value to check.
//
// The function returns the following values:
//
//    - ok: TRUE if value is valid for key.
//
func (key *SettingsSchemaKey) RangeCheck(value *glib.Variant) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(key)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(value)))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(key)
	runtime.KeepAlive(value)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SettingsSchemaSource: this is an opaque structure type. You may not access it
// directly.
//
// An instance of this type is always passed by reference.
type SettingsSchemaSource struct {
	*settingsSchemaSource
}

// settingsSchemaSource is the struct that's finalized.
type settingsSchemaSource struct {
	native unsafe.Pointer
}

func marshalSettingsSchemaSource(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &SettingsSchemaSource{&settingsSchemaSource{(unsafe.Pointer)(b)}}, nil
}

// NewSettingsSchemaSourceFromDirectory constructs a struct SettingsSchemaSource.
func NewSettingsSchemaSourceFromDirectory(directory string, parent *SettingsSchemaSource, trusted bool) (*SettingsSchemaSource, error) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(directory)))
	defer C.free(unsafe.Pointer(_args[0]))
	if parent != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(parent)))
	}
	if trusted {
		*(*C.gboolean)(unsafe.Pointer(&_args[2])) = C.TRUE
	}

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(directory)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(trusted)

	var _settingsSchemaSource *SettingsSchemaSource // out
	var _goerr error                                // out

	_settingsSchemaSource = (*SettingsSchemaSource)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_settingsSchemaSource)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _settingsSchemaSource, _goerr
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
//
// The function takes the following parameters:
//
//    - recursive: if we should recurse.
//
// The function returns the following values:
//
//    - nonRelocatable: the list of non-relocatable schemas, in no defined order.
//    - relocatable: list of relocatable schemas, in no defined order.
//
func (source *SettingsSchemaSource) ListSchemas(recursive bool) (nonRelocatable []string, relocatable []string) {
	var _args [2]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(source)))
	if recursive {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	runtime.KeepAlive(source)
	runtime.KeepAlive(recursive)

	var _nonRelocatable []string // out
	var _relocatable []string    // out

	defer C.free(unsafe.Pointer(_outs[0]))
	{
		var i int
		var z *C.void
		for p := _outs[0]; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_outs[0], i)
		_nonRelocatable = make([]string, i)
		for i := range src {
			_nonRelocatable[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	defer C.free(unsafe.Pointer(_outs[1]))
	{
		var i int
		var z *C.void
		for p := _outs[1]; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_outs[1], i)
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
//
// The function takes the following parameters:
//
//    - schemaId: schema ID.
//    - recursive: TRUE if the lookup should be recursive.
//
// The function returns the following values:
//
//    - settingsSchema (optional): new Schema.
//
func (source *SettingsSchemaSource) Lookup(schemaId string, recursive bool) *SettingsSchema {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(source)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(schemaId)))
	defer C.free(unsafe.Pointer(_args[1]))
	if recursive {
		*(*C.gboolean)(unsafe.Pointer(&_args[2])) = C.TRUE
	}

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(source)
	runtime.KeepAlive(schemaId)
	runtime.KeepAlive(recursive)

	var _settingsSchema *SettingsSchema // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_settingsSchema = (*SettingsSchema)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_settingsSchema)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _settingsSchema
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
//
// The function returns the following values:
//
//    - settingsSchemaSource (optional): default schema source.
//
func SettingsSchemaSourceGetDefault() *SettingsSchemaSource {
	_gret := girepository.MustFind("Gio", "get_default").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _settingsSchemaSource *SettingsSchemaSource // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_settingsSchemaSource = (*SettingsSchemaSource)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.g_settings_schema_source_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_settingsSchemaSource)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _settingsSchemaSource
}
