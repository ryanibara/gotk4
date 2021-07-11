// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_builder_closure_flags_get_type()), F: marshalBuilderClosureFlags},
		{T: externglib.Type(C.gtk_builder_scope_get_type()), F: marshalBuilderScoper},
		{T: externglib.Type(C.gtk_builder_cscope_get_type()), F: marshalBuilderCScoper},
	})
}

// BuilderClosureFlags: list of flags that can be passed to
// gtk_builder_create_closure().
//
// New values may be added in the future for new features, so external
// implementations of [interface@Gtk.BuilderScope] should test the flags for
// unknown values and raise a GTK_BUILDER_ERROR_INVALID_ATTRIBUTE error when
// they encounter one.
type BuilderClosureFlags int

const (
	// BuilderClosureFlagsSwapped: closure should be created swapped. See
	// g_cclosure_new_swap() for details.
	BuilderClosureFlagsSwapped BuilderClosureFlags = 0b1
)

func marshalBuilderClosureFlags(p uintptr) (interface{}, error) {
	return BuilderClosureFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// BuilderScopeOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type BuilderScopeOverrider interface {
	TypeFromFunction(builder Builderer, functionName string) externglib.Type

	TypeFromName(builder Builderer, typeName string) externglib.Type
}

// BuilderScoper describes BuilderScope's methods.
type BuilderScoper interface {
	privateBuilderScope()
}

// BuilderScope: `GtkBuilderScope` is an interface to provide language binding
// support to `GtkBuilder`.
//
// The goal of `GtkBuilderScope` is to look up programming-language-specific
// values for strings that are given in a `GtkBuilder` UI file.
//
// The primary intended audience is bindings that want to provide deeper
// integration of `GtkBuilder` into the language.
//
// A `GtkBuilderScope` instance may be used with multiple `GtkBuilder` objects,
// even at once.
//
// By default, GTK will use its own implementation of `GtkBuilderScope` for the
// C language which can be created via [ctor@Gtk.BuilderCScope.new].
type BuilderScope struct {
	*externglib.Object
}

var (
	_ BuilderScoper   = (*BuilderScope)(nil)
	_ gextras.Nativer = (*BuilderScope)(nil)
)

func wrapBuilderScope(obj *externglib.Object) BuilderScoper {
	return &BuilderScope{
		Object: obj,
	}
}

func marshalBuilderScoper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapBuilderScope(obj), nil
}

func (*BuilderScope) privateBuilderScope() {}

// BuilderCScoper describes BuilderCScope's methods.
type BuilderCScoper interface {
	privateBuilderCScope()
}

// BuilderCScope: `GtkBuilderScope` implementation for the C language.
//
// `GtkBuilderCScope` instances use symbols explicitly added to @builder with
// prior calls to [method@Gtk.BuilderCScope.add_callback_symbol]. If developers
// want to do that, they are encouraged to create their own scopes for that
// purpose.
//
// In the case that symbols are not explicitly added; GTK will uses `GModule`’s
// introspective features (by opening the module nil) to look at the
// application’s symbol table. From here it tries to match the signal function
// names given in the interface description with symbols in the application.
//
// Note that unless [method@Gtk.BuilderCScope.add_callback_symbol] is called for
// all signal callbacks which are referenced by the loaded XML, this
// functionality will require that `GModule` be supported on the platform.
type BuilderCScope struct {
	*externglib.Object

	BuilderScope
}

var (
	_ BuilderCScoper  = (*BuilderCScope)(nil)
	_ gextras.Nativer = (*BuilderCScope)(nil)
)

func wrapBuilderCScope(obj *externglib.Object) BuilderCScoper {
	return &BuilderCScope{
		Object: obj,
		BuilderScope: BuilderScope{
			Object: obj,
		},
	}
}

func marshalBuilderCScoper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapBuilderCScope(obj), nil
}

// NewBuilderCScope creates a new `GtkBuilderCScope` object to use with future
// `GtkBuilder` instances.
//
// Calling this function is only necessary if you want to add custom callbacks
// via [method@Gtk.BuilderCScope.add_callback_symbol].
func NewBuilderCScope() *BuilderCScope {
	var _cret *C.GtkBuilderScope // in

	_cret = C.gtk_builder_cscope_new()

	var _builderCScope *BuilderCScope // out

	_builderCScope = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*BuilderCScope)

	return _builderCScope
}

func (*BuilderCScope) privateBuilderCScope() {}
