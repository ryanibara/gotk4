// Code generated by girgen. DO NOT EDIT.

package gobject

// #cgo pkg-config: gobject-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
import "C"

// VALUEINTERNEDSTRING: for string values, indicates that the string contained
// is canonical and will exist for the duration of the process. See
// g_value_set_interned_string().
const VALUE_INTERNED_STRING = 268435456

// VALUENOCOPYCONTENTS: if passed to G_VALUE_COLLECT(), allocated data won't be
// copied but used verbatim. This does not affect ref-counted types like
// objects. This does not affect usage of g_value_copy(), the data will be
// copied if it is not ref-counted.
const VALUE_NOCOPY_CONTENTS = 134217728