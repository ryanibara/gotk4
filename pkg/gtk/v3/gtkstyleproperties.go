// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_style_properties_get_type()), F: marshalStyleProperties},
		{T: externglib.Type(C.gtk_gradient_get_type()), F: marshalGradient},
		{T: externglib.Type(C.gtk_symbolic_color_get_type()), F: marshalSymbolicColor},
	})
}

// StyleProperties provides the storage for style information that is used by
// StyleContext and other StyleProvider implementations.
//
// Before style properties can be stored in GtkStyleProperties, they must be
// registered with gtk_style_properties_register_property().
//
// Unless you are writing a StyleProvider implementation, you are unlikely to
// use this API directly, as gtk_style_context_get() and its variants are the
// preferred way to access styling information from widget implementations and
// theming engine implementations should use the APIs provided by ThemingEngine
// instead.
//
// StyleProperties has been deprecated in GTK 3.16. The CSS machinery does not
// use it anymore and all users of this object have been deprecated.
type StyleProperties interface {
	gextras.Objector

	// AsStyleProvider casts the class to the StyleProvider interface.
	AsStyleProvider() StyleProvider

	// ClearStyleProperties clears all style information from @props.
	//
	// Deprecated: since version 3.16.
	ClearStyleProperties()
	// Property gets a style property from @props for the given state. When done
	// with @value, g_value_unset() needs to be called to free any allocated
	// memory.
	//
	// Deprecated: since version 3.16.
	Property(property string, state StateFlags) (externglib.Value, bool)
	// LookupColorStyleProperties returns the symbolic color that is mapped to
	// @name.
	//
	// Deprecated: since version 3.8.
	LookupColorStyleProperties(name string) SymbolicColor
	// MapColorStyleProperties maps @color so it can be referenced by @name. See
	// gtk_style_properties_lookup_color()
	//
	// Deprecated: since version 3.8.
	MapColorStyleProperties(name string, color SymbolicColor)
	// MergeStyleProperties merges into @props all the style information
	// contained in @props_to_merge. If @replace is true, the values will be
	// overwritten, if it is false, the older values will prevail.
	//
	// Deprecated: since version 3.16.
	MergeStyleProperties(propsToMerge StyleProperties, replace bool)
	// SetPropertyStyleProperties sets a styling property in @props.
	//
	// Deprecated: since version 3.16.
	SetPropertyStyleProperties(property string, state StateFlags, value externglib.Value)
	// UnsetPropertyStyleProperties unsets a style property in @props.
	//
	// Deprecated: since version 3.16.
	UnsetPropertyStyleProperties(property string, state StateFlags)
}

// styleProperties implements the StyleProperties class.
type styleProperties struct {
	gextras.Objector
}

// WrapStyleProperties wraps a GObject to the right type. It is
// primarily used internally.
func WrapStyleProperties(obj *externglib.Object) StyleProperties {
	return styleProperties{
		Objector: obj,
	}
}

func marshalStyleProperties(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStyleProperties(obj), nil
}

// NewStyleProperties returns a newly created StyleProperties
//
// Deprecated: since version 3.16.
func NewStyleProperties() StyleProperties {
	var _cret *C.GtkStyleProperties // in

	_cret = C.gtk_style_properties_new()

	var _styleProperties StyleProperties // out

	_styleProperties = WrapStyleProperties(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _styleProperties
}

func (p styleProperties) ClearStyleProperties() {
	var _arg0 *C.GtkStyleProperties // out

	_arg0 = (*C.GtkStyleProperties)(unsafe.Pointer(p.Native()))

	C.gtk_style_properties_clear(_arg0)
}

func (p styleProperties) Property(property string, state StateFlags) (externglib.Value, bool) {
	var _arg0 *C.GtkStyleProperties // out
	var _arg1 *C.gchar              // out
	var _arg2 C.GtkStateFlags       // out
	var _arg3 *C.GValue             // in
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkStyleProperties)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkStateFlags(state)

	_cret = C.gtk_style_properties_get_property(_arg0, _arg1, _arg2, _arg3)

	var _value externglib.Value // out
	var _ok bool                // out

	_value = externglib.ValueFromNative(unsafe.Pointer(_arg3))
	runtime.SetFinalizer(_value, func(v *externglib.Value) {
		C.g_value_unset((*C.GValue)(v.GValue))
	})
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

func (p styleProperties) LookupColorStyleProperties(name string) SymbolicColor {
	var _arg0 *C.GtkStyleProperties // out
	var _arg1 *C.gchar              // out
	var _cret *C.GtkSymbolicColor   // in

	_arg0 = (*C.GtkStyleProperties)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_style_properties_lookup_color(_arg0, _arg1)

	var _symbolicColor SymbolicColor // out

	_symbolicColor = (SymbolicColor)(unsafe.Pointer(_cret))
	C.gtk_symbolic_color_ref(_cret)

	return _symbolicColor
}

func (p styleProperties) MapColorStyleProperties(name string, color SymbolicColor) {
	var _arg0 *C.GtkStyleProperties // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.GtkSymbolicColor   // out

	_arg0 = (*C.GtkStyleProperties)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkSymbolicColor)(unsafe.Pointer(color))

	C.gtk_style_properties_map_color(_arg0, _arg1, _arg2)
}

func (p styleProperties) MergeStyleProperties(propsToMerge StyleProperties, replace bool) {
	var _arg0 *C.GtkStyleProperties // out
	var _arg1 *C.GtkStyleProperties // out
	var _arg2 C.gboolean            // out

	_arg0 = (*C.GtkStyleProperties)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkStyleProperties)(unsafe.Pointer(propsToMerge.Native()))
	if replace {
		_arg2 = C.TRUE
	}

	C.gtk_style_properties_merge(_arg0, _arg1, _arg2)
}

func (p styleProperties) SetPropertyStyleProperties(property string, state StateFlags, value externglib.Value) {
	var _arg0 *C.GtkStyleProperties // out
	var _arg1 *C.gchar              // out
	var _arg2 C.GtkStateFlags       // out
	var _arg3 *C.GValue             // out

	_arg0 = (*C.GtkStyleProperties)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkStateFlags(state)
	_arg3 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	C.gtk_style_properties_set_property(_arg0, _arg1, _arg2, _arg3)
}

func (p styleProperties) UnsetPropertyStyleProperties(property string, state StateFlags) {
	var _arg0 *C.GtkStyleProperties // out
	var _arg1 *C.gchar              // out
	var _arg2 C.GtkStateFlags       // out

	_arg0 = (*C.GtkStyleProperties)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkStateFlags(state)

	C.gtk_style_properties_unset_property(_arg0, _arg1, _arg2)
}

func (s styleProperties) AsStyleProvider() StyleProvider {
	return WrapStyleProvider(gextras.InternObject(s))
}

// Gradient is a boxed type that represents a gradient. It is the result of
// parsing a [gradient expression][gtkcssprovider-gradients]. To obtain the
// gradient represented by a GtkGradient, it has to be resolved with
// gtk_gradient_resolve(), which replaces all symbolic color references by the
// colors they refer to (in a given context) and constructs a #cairo_pattern_t
// value.
//
// It is not normally necessary to deal directly with Gradients, since they are
// mostly used behind the scenes by StyleContext and CssProvider.
//
// Gradient is deprecated. It was used internally by GTK’s CSS engine to
// represent gradients. As its handling is not conforming to modern web
// standards, it is not used anymore. If you want to use gradients in your own
// code, please use Cairo directly.
type Gradient struct {
	native C.GtkGradient
}

// WrapGradient wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapGradient(ptr unsafe.Pointer) *Gradient {
	return (*Gradient)(ptr)
}

func marshalGradient(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Gradient)(unsafe.Pointer(b)), nil
}

// NewGradientLinear constructs a struct Gradient.
func NewGradientLinear(x0 float64, y0 float64, x1 float64, y1 float64) Gradient {
	var _arg1 C.gdouble      // out
	var _arg2 C.gdouble      // out
	var _arg3 C.gdouble      // out
	var _arg4 C.gdouble      // out
	var _cret *C.GtkGradient // in

	_arg1 = C.gdouble(x0)
	_arg2 = C.gdouble(y0)
	_arg3 = C.gdouble(x1)
	_arg4 = C.gdouble(y1)

	_cret = C.gtk_gradient_new_linear(_arg1, _arg2, _arg3, _arg4)

	var _gradient Gradient // out

	_gradient = (Gradient)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_gradient, func(v Gradient) {
		C.gtk_gradient_unref((*C.GtkGradient)(unsafe.Pointer(v)))
	})

	return _gradient
}

// NewGradientRadial constructs a struct Gradient.
func NewGradientRadial(x0 float64, y0 float64, radius0 float64, x1 float64, y1 float64, radius1 float64) Gradient {
	var _arg1 C.gdouble      // out
	var _arg2 C.gdouble      // out
	var _arg3 C.gdouble      // out
	var _arg4 C.gdouble      // out
	var _arg5 C.gdouble      // out
	var _arg6 C.gdouble      // out
	var _cret *C.GtkGradient // in

	_arg1 = C.gdouble(x0)
	_arg2 = C.gdouble(y0)
	_arg3 = C.gdouble(radius0)
	_arg4 = C.gdouble(x1)
	_arg5 = C.gdouble(y1)
	_arg6 = C.gdouble(radius1)

	_cret = C.gtk_gradient_new_radial(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)

	var _gradient Gradient // out

	_gradient = (Gradient)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_gradient, func(v Gradient) {
		C.gtk_gradient_unref((*C.GtkGradient)(unsafe.Pointer(v)))
	})

	return _gradient
}

// Native returns the underlying C source pointer.
func (g *Gradient) Native() unsafe.Pointer {
	return unsafe.Pointer(&g.native)
}

// AddColorStop adds a stop color to @gradient.
//
// Deprecated: since version 3.8.
func (g *Gradient) AddColorStop(offset float64, color SymbolicColor) {
	var _arg0 *C.GtkGradient      // out
	var _arg1 C.gdouble           // out
	var _arg2 *C.GtkSymbolicColor // out

	_arg0 = (*C.GtkGradient)(unsafe.Pointer(g))
	_arg1 = C.gdouble(offset)
	_arg2 = (*C.GtkSymbolicColor)(unsafe.Pointer(color))

	C.gtk_gradient_add_color_stop(_arg0, _arg1, _arg2)
}

// Ref increases the reference count of @gradient.
//
// Deprecated: since version 3.8.
func (g *Gradient) Ref() Gradient {
	var _arg0 *C.GtkGradient // out
	var _cret *C.GtkGradient // in

	_arg0 = (*C.GtkGradient)(unsafe.Pointer(g))

	_cret = C.gtk_gradient_ref(_arg0)

	var _ret Gradient // out

	_ret = (Gradient)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_ret, func(v Gradient) {
		C.gtk_gradient_unref((*C.GtkGradient)(unsafe.Pointer(v)))
	})

	return _ret
}

// Resolve: if @gradient is resolvable, @resolved_gradient will be filled in
// with the resolved gradient as a cairo_pattern_t, and true will be returned.
// Generally, if @gradient can’t be resolved, it is due to it being defined on
// top of a named color that doesn't exist in @props.
//
// Deprecated: since version 3.8.
func (g *Gradient) Resolve(props StyleProperties) (*cairo.Pattern, bool) {
	var _arg0 *C.GtkGradient        // out
	var _arg1 *C.GtkStyleProperties // out
	var _arg2 **C.cairo_pattern_t   // in
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkGradient)(unsafe.Pointer(g))
	_arg1 = (*C.GtkStyleProperties)(unsafe.Pointer(props.Native()))

	_cret = C.gtk_gradient_resolve(_arg0, _arg1, &_arg2)

	var _resolvedGradient *cairo.Pattern // out
	var _ok bool                         // out

	{
		var refTmpIn *C.cairo_pattern_t
		var refTmpOut *cairo.Pattern

		refTmpIn = *_arg2

		refTmpOut = (*cairo.Pattern)(unsafe.Pointer(refTmpIn))
		runtime.SetFinalizer(refTmpOut, func(v *cairo.Pattern) {
			C.free(unsafe.Pointer(v))
		})

		_resolvedGradient = refTmpOut
	}
	if _cret != 0 {
		_ok = true
	}

	return _resolvedGradient, _ok
}

func (g *Gradient) ResolveForContext(context StyleContext) cairo.Pattern {
	var _arg0 *C.GtkGradient     // out
	var _arg1 *C.GtkStyleContext // out
	var _cret *C.cairo_pattern_t // in

	_arg0 = (*C.GtkGradient)(unsafe.Pointer(g))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_gradient_resolve_for_context(_arg0, _arg1)

	var _pattern cairo.Pattern // out

	_pattern = (cairo.Pattern)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_pattern, func(v cairo.Pattern) {
		C.free(unsafe.Pointer(v))
	})

	return _pattern
}

// String creates a string representation for @gradient that is suitable for
// using in GTK CSS files.
//
// Deprecated: since version 3.8.
func (g *Gradient) String() string {
	var _arg0 *C.GtkGradient // out
	var _cret *C.char        // in

	_arg0 = (*C.GtkGradient)(unsafe.Pointer(g))

	_cret = C.gtk_gradient_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Unref decreases the reference count of @gradient, freeing its memory if the
// reference count reaches 0.
//
// Deprecated: since version 3.8.
func (g *Gradient) Unref() {
	var _arg0 *C.GtkGradient // out

	_arg0 = (*C.GtkGradient)(unsafe.Pointer(g))

	C.gtk_gradient_unref(_arg0)
}

// SymbolicColor is a boxed type that represents a symbolic color. It is the
// result of parsing a [color expression][gtkcssprovider-symbolic-colors]. To
// obtain the color represented by a GtkSymbolicColor, it has to be resolved
// with gtk_symbolic_color_resolve(), which replaces all symbolic color
// references by the colors they refer to (in a given context) and evaluates
// mix, shade and other expressions, resulting in a RGBA value.
//
// It is not normally necessary to deal directly with SymbolicColors, since they
// are mostly used behind the scenes by StyleContext and CssProvider.
//
// SymbolicColor is deprecated. Symbolic colors are considered an implementation
// detail of GTK+.
type SymbolicColor struct {
	native C.GtkSymbolicColor
}

// WrapSymbolicColor wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapSymbolicColor(ptr unsafe.Pointer) *SymbolicColor {
	return (*SymbolicColor)(ptr)
}

func marshalSymbolicColor(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*SymbolicColor)(unsafe.Pointer(b)), nil
}

// NewSymbolicColorAlpha constructs a struct SymbolicColor.
func NewSymbolicColorAlpha(color SymbolicColor, factor float64) SymbolicColor {
	var _arg1 *C.GtkSymbolicColor // out
	var _arg2 C.gdouble           // out
	var _cret *C.GtkSymbolicColor // in

	_arg1 = (*C.GtkSymbolicColor)(unsafe.Pointer(color))
	_arg2 = C.gdouble(factor)

	_cret = C.gtk_symbolic_color_new_alpha(_arg1, _arg2)

	var _symbolicColor SymbolicColor // out

	_symbolicColor = (SymbolicColor)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_symbolicColor, func(v SymbolicColor) {
		C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(unsafe.Pointer(v)))
	})

	return _symbolicColor
}

// NewSymbolicColorLiteral constructs a struct SymbolicColor.
func NewSymbolicColorLiteral(color gdk.RGBA) SymbolicColor {
	var _arg1 *C.GdkRGBA          // out
	var _cret *C.GtkSymbolicColor // in

	_arg1 = (*C.GdkRGBA)(unsafe.Pointer(color))

	_cret = C.gtk_symbolic_color_new_literal(_arg1)

	var _symbolicColor SymbolicColor // out

	_symbolicColor = (SymbolicColor)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_symbolicColor, func(v SymbolicColor) {
		C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(unsafe.Pointer(v)))
	})

	return _symbolicColor
}

// NewSymbolicColorMix constructs a struct SymbolicColor.
func NewSymbolicColorMix(color1 SymbolicColor, color2 SymbolicColor, factor float64) SymbolicColor {
	var _arg1 *C.GtkSymbolicColor // out
	var _arg2 *C.GtkSymbolicColor // out
	var _arg3 C.gdouble           // out
	var _cret *C.GtkSymbolicColor // in

	_arg1 = (*C.GtkSymbolicColor)(unsafe.Pointer(color1))
	_arg2 = (*C.GtkSymbolicColor)(unsafe.Pointer(color2))
	_arg3 = C.gdouble(factor)

	_cret = C.gtk_symbolic_color_new_mix(_arg1, _arg2, _arg3)

	var _symbolicColor SymbolicColor // out

	_symbolicColor = (SymbolicColor)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_symbolicColor, func(v SymbolicColor) {
		C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(unsafe.Pointer(v)))
	})

	return _symbolicColor
}

// NewSymbolicColorName constructs a struct SymbolicColor.
func NewSymbolicColorName(name string) SymbolicColor {
	var _arg1 *C.gchar            // out
	var _cret *C.GtkSymbolicColor // in

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_symbolic_color_new_name(_arg1)

	var _symbolicColor SymbolicColor // out

	_symbolicColor = (SymbolicColor)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_symbolicColor, func(v SymbolicColor) {
		C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(unsafe.Pointer(v)))
	})

	return _symbolicColor
}

// NewSymbolicColorShade constructs a struct SymbolicColor.
func NewSymbolicColorShade(color SymbolicColor, factor float64) SymbolicColor {
	var _arg1 *C.GtkSymbolicColor // out
	var _arg2 C.gdouble           // out
	var _cret *C.GtkSymbolicColor // in

	_arg1 = (*C.GtkSymbolicColor)(unsafe.Pointer(color))
	_arg2 = C.gdouble(factor)

	_cret = C.gtk_symbolic_color_new_shade(_arg1, _arg2)

	var _symbolicColor SymbolicColor // out

	_symbolicColor = (SymbolicColor)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_symbolicColor, func(v SymbolicColor) {
		C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(unsafe.Pointer(v)))
	})

	return _symbolicColor
}

// NewSymbolicColorWin32 constructs a struct SymbolicColor.
func NewSymbolicColorWin32(themeClass string, id int) SymbolicColor {
	var _arg1 *C.gchar            // out
	var _arg2 C.gint              // out
	var _cret *C.GtkSymbolicColor // in

	_arg1 = (*C.gchar)(C.CString(themeClass))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(id)

	_cret = C.gtk_symbolic_color_new_win32(_arg1, _arg2)

	var _symbolicColor SymbolicColor // out

	_symbolicColor = (SymbolicColor)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_symbolicColor, func(v SymbolicColor) {
		C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(unsafe.Pointer(v)))
	})

	return _symbolicColor
}

// Native returns the underlying C source pointer.
func (s *SymbolicColor) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// Ref increases the reference count of @color
//
// Deprecated: since version 3.8.
func (c *SymbolicColor) Ref() SymbolicColor {
	var _arg0 *C.GtkSymbolicColor // out
	var _cret *C.GtkSymbolicColor // in

	_arg0 = (*C.GtkSymbolicColor)(unsafe.Pointer(c))

	_cret = C.gtk_symbolic_color_ref(_arg0)

	var _symbolicColor SymbolicColor // out

	_symbolicColor = (SymbolicColor)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_symbolicColor, func(v SymbolicColor) {
		C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(unsafe.Pointer(v)))
	})

	return _symbolicColor
}

// Resolve: if @color is resolvable, @resolved_color will be filled in with the
// resolved color, and true will be returned. Generally, if @color can’t be
// resolved, it is due to it being defined on top of a named color that doesn’t
// exist in @props.
//
// When @props is nil, resolving of named colors will fail, so if your @color is
// or references such a color, this function will return false.
//
// Deprecated: since version 3.8.
func (c *SymbolicColor) Resolve(props StyleProperties) (gdk.RGBA, bool) {
	var _arg0 *C.GtkSymbolicColor   // out
	var _arg1 *C.GtkStyleProperties // out
	var _arg2 *C.GdkRGBA            // in
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkSymbolicColor)(unsafe.Pointer(c))
	_arg1 = (*C.GtkStyleProperties)(unsafe.Pointer(props.Native()))

	_cret = C.gtk_symbolic_color_resolve(_arg0, _arg1, &_arg2)

	var _resolvedColor gdk.RGBA // out
	var _ok bool                // out

	_resolvedColor = (gdk.RGBA)(unsafe.Pointer(_arg2))
	if _cret != 0 {
		_ok = true
	}

	return _resolvedColor, _ok
}

// String converts the given @color to a string representation. This is useful
// both for debugging and for serialization of strings. The format of the string
// may change between different versions of GTK, but it is guaranteed that the
// GTK css parser is able to read the string and create the same symbolic color
// from it.
//
// Deprecated: since version 3.8.
func (c *SymbolicColor) String() string {
	var _arg0 *C.GtkSymbolicColor // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkSymbolicColor)(unsafe.Pointer(c))

	_cret = C.gtk_symbolic_color_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Unref decreases the reference count of @color, freeing its memory if the
// reference count reaches 0.
//
// Deprecated: since version 3.8.
func (c *SymbolicColor) Unref() {
	var _arg0 *C.GtkSymbolicColor // out

	_arg0 = (*C.GtkSymbolicColor)(unsafe.Pointer(c))

	C.gtk_symbolic_color_unref(_arg0)
}
