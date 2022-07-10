// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeBuilderClosureFlags returns the GType for the type BuilderClosureFlags.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeBuilderClosureFlags() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "BuilderClosureFlags").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalBuilderClosureFlags)
	return gtype
}

// GTypeBuilderScope returns the GType for the type BuilderScope.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeBuilderScope() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "BuilderScope").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalBuilderScope)
	return gtype
}

// GTypeBuilderCScope returns the GType for the type BuilderCScope.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeBuilderCScope() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "BuilderCScope").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalBuilderCScope)
	return gtype
}

// BuilderClosureFlags: list of flags that can be passed to
// gtk_builder_create_closure().
//
// New values may be added in the future for new features, so external
// implementations of gtk.BuilderScope should test the flags for unknown values
// and raise a GTK_BUILDER_ERROR_INVALID_ATTRIBUTE error when they encounter
// one.
type BuilderClosureFlags C.guint

const (
	// BuilderClosureSwapped: closure should be created swapped. See
	// g_cclosure_new_swap() for details.
	BuilderClosureSwapped BuilderClosureFlags = 0b1
)

func marshalBuilderClosureFlags(p uintptr) (interface{}, error) {
	return BuilderClosureFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for BuilderClosureFlags.
func (b BuilderClosureFlags) String() string {
	if b == 0 {
		return "BuilderClosureFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(21)

	for b != 0 {
		next := b & (b - 1)
		bit := b - next

		switch bit {
		case BuilderClosureSwapped:
			builder.WriteString("Swapped|")
		default:
			builder.WriteString(fmt.Sprintf("BuilderClosureFlags(0b%b)|", bit))
		}

		b = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if b contains other.
func (b BuilderClosureFlags) Has(other BuilderClosureFlags) bool {
	return (b & other) == other
}

// BuilderScopeOverrider contains methods that are overridable.
type BuilderScopeOverrider interface {
}

// BuilderScope: GtkBuilderScope is an interface to provide language binding
// support to GtkBuilder.
//
// The goal of GtkBuilderScope is to look up programming-language-specific
// values for strings that are given in a GtkBuilder UI file.
//
// The primary intended audience is bindings that want to provide deeper
// integration of GtkBuilder into the language.
//
// A GtkBuilderScope instance may be used with multiple GtkBuilder objects, even
// at once.
//
// By default, GTK will use its own implementation of GtkBuilderScope for the C
// language which can be created via gtk.BuilderCScope.New.
//
// BuilderScope wraps an interface. This means the user can get the
// underlying type by calling Cast().
type BuilderScope struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*BuilderScope)(nil)
)

// BuilderScoper describes BuilderScope's interface methods.
type BuilderScoper interface {
	coreglib.Objector

	baseBuilderScope() *BuilderScope
}

var _ BuilderScoper = (*BuilderScope)(nil)

func ifaceInitBuilderScoper(gifacePtr, data C.gpointer) {
}

func wrapBuilderScope(obj *coreglib.Object) *BuilderScope {
	return &BuilderScope{
		Object: obj,
	}
}

func marshalBuilderScope(p uintptr) (interface{}, error) {
	return wrapBuilderScope(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *BuilderScope) baseBuilderScope() *BuilderScope {
	return v
}

// BaseBuilderScope returns the underlying base object.
func BaseBuilderScope(obj BuilderScoper) *BuilderScope {
	return obj.baseBuilderScope()
}

// BuilderCScopeOverrider contains methods that are overridable.
type BuilderCScopeOverrider interface {
}

// BuilderCScope: GtkBuilderScope implementation for the C language.
//
// GtkBuilderCScope instances use symbols explicitly added to builder with prior
// calls to gtk.BuilderCScope.AddCallbackSymbol(). If developers want to do
// that, they are encouraged to create their own scopes for that purpose.
//
// In the case that symbols are not explicitly added; GTK will uses GModule’s
// introspective features (by opening the module NULL) to look at the
// application’s symbol table. From here it tries to match the signal function
// names given in the interface description with symbols in the application.
//
// Note that unless gtk.BuilderCScope.AddCallbackSymbol() is called for all
// signal callbacks which are referenced by the loaded XML, this functionality
// will require that GModule be supported on the platform.
type BuilderCScope struct {
	_ [0]func() // equal guard
	*coreglib.Object

	BuilderScope
}

var (
	_ coreglib.Objector = (*BuilderCScope)(nil)
)

func classInitBuilderCScoper(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapBuilderCScope(obj *coreglib.Object) *BuilderCScope {
	return &BuilderCScope{
		Object: obj,
		BuilderScope: BuilderScope{
			Object: obj,
		},
	}
}

func marshalBuilderCScope(p uintptr) (interface{}, error) {
	return wrapBuilderCScope(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBuilderCScope creates a new GtkBuilderCScope object to use with future
// GtkBuilder instances.
//
// Calling this function is only necessary if you want to add custom callbacks
// via gtk.BuilderCScope.AddCallbackSymbol().
//
// The function returns the following values:
//
//    - builderCScope: new GtkBuilderCScope.
//
func NewBuilderCScope() *BuilderCScope {
	_info := girepository.MustFind("Gtk", "BuilderCScope")
	_gret := _info.InvokeClassMethod("new_BuilderCScope", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _builderCScope *BuilderCScope // out

	_builderCScope = wrapBuilderCScope(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _builderCScope
}
