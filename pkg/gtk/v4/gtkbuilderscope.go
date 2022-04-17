// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"strings"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern GType _gotk4_gtk4_BuilderScopeInterface_get_type_from_function(GtkBuilderScope*, GtkBuilder*, char*);
// extern GType _gotk4_gtk4_BuilderScopeInterface_get_type_from_name(GtkBuilderScope*, GtkBuilder*, char*);
import "C"

// glib.Type values for gtkbuilderscope.go.
var (
	GTypeBuilderClosureFlags = externglib.Type(C.gtk_builder_closure_flags_get_type())
	GTypeBuilderScope        = externglib.Type(C.gtk_builder_scope_get_type())
	GTypeBuilderCScope       = externglib.Type(C.gtk_builder_cscope_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeBuilderClosureFlags, F: marshalBuilderClosureFlags},
		{T: GTypeBuilderScope, F: marshalBuilderScope},
		{T: GTypeBuilderCScope, F: marshalBuilderCScope},
	})
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
	return BuilderClosureFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
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
	externglib.Objector
	// The function takes the following parameters:
	//
	//    - builder
	//    - functionName
	//
	// The function returns the following values:
	//
	TypeFromFunction(builder *Builder, functionName string) externglib.Type
	// The function takes the following parameters:
	//
	//    - builder
	//    - typeName
	//
	// The function returns the following values:
	//
	TypeFromName(builder *Builder, typeName string) externglib.Type
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
type BuilderScope struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*BuilderScope)(nil)
)

// BuilderScoper describes BuilderScope's interface methods.
type BuilderScoper interface {
	externglib.Objector

	baseBuilderScope() *BuilderScope
}

var _ BuilderScoper = (*BuilderScope)(nil)

func ifaceInitBuilderScoper(gifacePtr, data C.gpointer) {
	iface := (*C.GtkBuilderScopeInterface)(unsafe.Pointer(gifacePtr))
	iface.get_type_from_function = (*[0]byte)(C._gotk4_gtk4_BuilderScopeInterface_get_type_from_function)
	iface.get_type_from_name = (*[0]byte)(C._gotk4_gtk4_BuilderScopeInterface_get_type_from_name)
}

//export _gotk4_gtk4_BuilderScopeInterface_get_type_from_function
func _gotk4_gtk4_BuilderScopeInterface_get_type_from_function(arg0 *C.GtkBuilderScope, arg1 *C.GtkBuilder, arg2 *C.char) (cret C.GType) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuilderScopeOverrider)

	var _builder *Builder    // out
	var _functionName string // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))
	_functionName = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	gType := iface.TypeFromFunction(_builder, _functionName)

	cret = C.GType(gType)

	return cret
}

//export _gotk4_gtk4_BuilderScopeInterface_get_type_from_name
func _gotk4_gtk4_BuilderScopeInterface_get_type_from_name(arg0 *C.GtkBuilderScope, arg1 *C.GtkBuilder, arg2 *C.char) (cret C.GType) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuilderScopeOverrider)

	var _builder *Builder // out
	var _typeName string  // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))
	_typeName = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	gType := iface.TypeFromName(_builder, _typeName)

	cret = C.GType(gType)

	return cret
}

func wrapBuilderScope(obj *externglib.Object) *BuilderScope {
	return &BuilderScope{
		Object: obj,
	}
}

func marshalBuilderScope(p uintptr) (interface{}, error) {
	return wrapBuilderScope(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *BuilderScope) baseBuilderScope() *BuilderScope {
	return self
}

// BaseBuilderScope returns the underlying base object.
func BaseBuilderScope(obj BuilderScoper) *BuilderScope {
	return obj.baseBuilderScope()
}

// BuilderCScopeOverrider contains methods that are overridable.
type BuilderCScopeOverrider interface {
	externglib.Objector
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
	*externglib.Object

	BuilderScope
}

var (
	_ externglib.Objector = (*BuilderCScope)(nil)
)

func classInitBuilderCScoper(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapBuilderCScope(obj *externglib.Object) *BuilderCScope {
	return &BuilderCScope{
		Object: obj,
		BuilderScope: BuilderScope{
			Object: obj,
		},
	}
}

func marshalBuilderCScope(p uintptr) (interface{}, error) {
	return wrapBuilderCScope(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
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
	var _cret *C.GtkBuilderScope // in

	_cret = C.gtk_builder_cscope_new()

	var _builderCScope *BuilderCScope // out

	_builderCScope = wrapBuilderCScope(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _builderCScope
}
