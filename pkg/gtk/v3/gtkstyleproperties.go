// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
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

// StyleProperties gtkStyleProperties provides the storage for style information
// that is used by StyleContext and other StyleProvider implementations.
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
	StyleProvider

	// Clear clears all style information from @props.
	Clear()
	// Property gets a style property from @props for the given state. When done
	// with @value, g_value_unset() needs to be called to free any allocated
	// memory.
	Property(property string, state StateFlags) (*externglib.Value, bool)
	// LookupColor returns the symbolic color that is mapped to @name.
	LookupColor(name string) *SymbolicColor
	// MapColor maps @color so it can be referenced by @name. See
	// gtk_style_properties_lookup_color()
	MapColor(name string, color *SymbolicColor)
	// Merge merges into @props all the style information contained in
	// @props_to_merge. If @replace is true, the values will be overwritten, if
	// it is false, the older values will prevail.
	Merge(propsToMerge StyleProperties, replace bool)
	// SetProperty sets a styling property in @props.
	SetProperty(property string, state StateFlags, value **externglib.Value)
	// UnsetProperty unsets a style property in @props.
	UnsetProperty(property string, state StateFlags)
}

// styleProperties implements the StyleProperties interface.
type styleProperties struct {
	gextras.Objector
	StyleProvider
}

var _ StyleProperties = (*styleProperties)(nil)

// WrapStyleProperties wraps a GObject to the right type. It is
// primarily used internally.
func WrapStyleProperties(obj *externglib.Object) StyleProperties {
	return StyleProperties{
		Objector:      obj,
		StyleProvider: WrapStyleProvider(obj),
	}
}

func marshalStyleProperties(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStyleProperties(obj), nil
}

// NewStyleProperties constructs a class StyleProperties.
func NewStyleProperties() StyleProperties {
	var _cret C.GtkStyleProperties

	cret = C.gtk_style_properties_new()

	var _styleProperties StyleProperties

	_styleProperties = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(StyleProperties)

	return _styleProperties
}

// Clear clears all style information from @props.
func (p styleProperties) Clear() {
	var _arg0 *C.GtkStyleProperties

	_arg0 = (*C.GtkStyleProperties)(unsafe.Pointer(p.Native()))

	C.gtk_style_properties_clear(_arg0)
}

// Property gets a style property from @props for the given state. When done
// with @value, g_value_unset() needs to be called to free any allocated
// memory.
func (p styleProperties) Property(property string, state StateFlags) (*externglib.Value, bool) {
	var _arg0 *C.GtkStyleProperties
	var _arg1 *C.gchar
	var _arg2 C.GtkStateFlags

	_arg0 = (*C.GtkStyleProperties)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.GtkStateFlags)(state)

	var _arg3 C.GValue
	var _cret C.gboolean

	cret = C.gtk_style_properties_get_property(_arg0, _arg1, _arg2, &_arg3)

	var _value *externglib.Value
	var _ok bool

	_value = externglib.ValueFromNative(unsafe.Pointer(_arg3))
	runtime.SetFinalizer(_value, func(v *externglib.Value) {
		C.g_value_unset((*C.GValue)(v.GValue))
	})
	if _cret {
		_ok = true
	}

	return _value, _ok
}

// LookupColor returns the symbolic color that is mapped to @name.
func (p styleProperties) LookupColor(name string) *SymbolicColor {
	var _arg0 *C.GtkStyleProperties
	var _arg1 *C.gchar

	_arg0 = (*C.GtkStyleProperties)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.GtkSymbolicColor

	cret = C.gtk_style_properties_lookup_color(_arg0, _arg1)

	var _symbolicColor *SymbolicColor

	_symbolicColor = WrapSymbolicColor(unsafe.Pointer(_cret))

	return _symbolicColor
}

// MapColor maps @color so it can be referenced by @name. See
// gtk_style_properties_lookup_color()
func (p styleProperties) MapColor(name string, color *SymbolicColor) {
	var _arg0 *C.GtkStyleProperties
	var _arg1 *C.gchar
	var _arg2 *C.GtkSymbolicColor

	_arg0 = (*C.GtkStyleProperties)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkSymbolicColor)(unsafe.Pointer(color.Native()))

	C.gtk_style_properties_map_color(_arg0, _arg1, _arg2)
}

// Merge merges into @props all the style information contained in
// @props_to_merge. If @replace is true, the values will be overwritten, if
// it is false, the older values will prevail.
func (p styleProperties) Merge(propsToMerge StyleProperties, replace bool) {
	var _arg0 *C.GtkStyleProperties
	var _arg1 *C.GtkStyleProperties
	var _arg2 C.gboolean

	_arg0 = (*C.GtkStyleProperties)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkStyleProperties)(unsafe.Pointer(propsToMerge.Native()))
	if replace {
		_arg2 = C.gboolean(1)
	}

	C.gtk_style_properties_merge(_arg0, _arg1, _arg2)
}

// SetProperty sets a styling property in @props.
func (p styleProperties) SetProperty(property string, state StateFlags, value **externglib.Value) {
	var _arg0 *C.GtkStyleProperties
	var _arg1 *C.gchar
	var _arg2 C.GtkStateFlags
	var _arg3 *C.GValue

	_arg0 = (*C.GtkStyleProperties)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.GtkStateFlags)(state)
	_arg3 = (*C.GValue)(value.GValue)

	C.gtk_style_properties_set_property(_arg0, _arg1, _arg2, _arg3)
}

// UnsetProperty unsets a style property in @props.
func (p styleProperties) UnsetProperty(property string, state StateFlags) {
	var _arg0 *C.GtkStyleProperties
	var _arg1 *C.gchar
	var _arg2 C.GtkStateFlags

	_arg0 = (*C.GtkStyleProperties)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.GtkStateFlags)(state)

	C.gtk_style_properties_unset_property(_arg0, _arg1, _arg2)
}

// Gradient: gtkGradient is a boxed type that represents a gradient. It is the
// result of parsing a [gradient expression][gtkcssprovider-gradients]. To
// obtain the gradient represented by a GtkGradient, it has to be resolved with
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
	if ptr == nil {
		return nil
	}

	return (*Gradient)(ptr)
}

func marshalGradient(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapGradient(unsafe.Pointer(b)), nil
}

// NewGradientLinear constructs a struct Gradient.
func NewGradientLinear(x0 float64, y0 float64, x1 float64, y1 float64) *Gradient {
	var _arg1 C.gdouble
	var _arg2 C.gdouble
	var _arg3 C.gdouble
	var _arg4 C.gdouble

	_arg1 = C.gdouble(x0)
	_arg2 = C.gdouble(y0)
	_arg3 = C.gdouble(x1)
	_arg4 = C.gdouble(y1)

	var _cret *C.GtkGradient

	cret = C.gtk_gradient_new_linear(_arg1, _arg2, _arg3, _arg4)

	var _gradient *Gradient

	_gradient = WrapGradient(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_gradient, func(v *Gradient) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _gradient
}

// NewGradientRadial constructs a struct Gradient.
func NewGradientRadial(x0 float64, y0 float64, radius0 float64, x1 float64, y1 float64, radius1 float64) *Gradient {
	var _arg1 C.gdouble
	var _arg2 C.gdouble
	var _arg3 C.gdouble
	var _arg4 C.gdouble
	var _arg5 C.gdouble
	var _arg6 C.gdouble

	_arg1 = C.gdouble(x0)
	_arg2 = C.gdouble(y0)
	_arg3 = C.gdouble(radius0)
	_arg4 = C.gdouble(x1)
	_arg5 = C.gdouble(y1)
	_arg6 = C.gdouble(radius1)

	var _cret *C.GtkGradient

	cret = C.gtk_gradient_new_radial(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)

	var _gradient *Gradient

	_gradient = WrapGradient(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_gradient, func(v *Gradient) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _gradient
}

// Native returns the underlying C source pointer.
func (g *Gradient) Native() unsafe.Pointer {
	return unsafe.Pointer(&g.native)
}

// AddColorStop adds a stop color to @gradient.
func (g *Gradient) AddColorStop(offset float64, color *SymbolicColor) {
	var _arg0 *C.GtkGradient
	var _arg1 C.gdouble
	var _arg2 *C.GtkSymbolicColor

	_arg0 = (*C.GtkGradient)(unsafe.Pointer(g.Native()))
	_arg1 = C.gdouble(offset)
	_arg2 = (*C.GtkSymbolicColor)(unsafe.Pointer(color.Native()))

	C.gtk_gradient_add_color_stop(_arg0, _arg1, _arg2)
}

// Ref increases the reference count of @gradient.
func (g *Gradient) Ref() *Gradient {
	var _arg0 *C.GtkGradient

	_arg0 = (*C.GtkGradient)(unsafe.Pointer(g.Native()))

	var _cret *C.GtkGradient

	cret = C.gtk_gradient_ref(_arg0)

	var _ret *Gradient

	_ret = WrapGradient(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_ret, func(v *Gradient) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _ret
}

// Resolve: if @gradient is resolvable, @resolved_gradient will be filled in
// with the resolved gradient as a cairo_pattern_t, and true will be returned.
// Generally, if @gradient can’t be resolved, it is due to it being defined on
// top of a named color that doesn't exist in @props.
func (g *Gradient) Resolve(props StyleProperties) (*cairo.Pattern, bool) {
	var _arg0 *C.GtkGradient
	var _arg1 *C.GtkStyleProperties

	_arg0 = (*C.GtkGradient)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GtkStyleProperties)(unsafe.Pointer(props.Native()))

	var _resolvedGradient *cairo.Pattern
	var _cret C.gboolean

	cret = C.gtk_gradient_resolve(_arg0, _arg1, (**C.cairo_pattern_t)(unsafe.Pointer(&_resolvedGradient)))

	var _ok bool

	if _cret {
		_ok = true
	}

	return _resolvedGradient, _ok
}

func (g *Gradient) ResolveForContext(context StyleContext) *cairo.Pattern {
	var _arg0 *C.GtkGradient
	var _arg1 *C.GtkStyleContext

	_arg0 = (*C.GtkGradient)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))

	var _cret *C.cairo_pattern_t

	cret = C.gtk_gradient_resolve_for_context(_arg0, _arg1)

	var _pattern *cairo.Pattern

	_pattern = cairo.WrapPattern(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_pattern, func(v *cairo.Pattern) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _pattern
}

// String creates a string representation for @gradient that is suitable for
// using in GTK CSS files.
func (g *Gradient) String() string {
	var _arg0 *C.GtkGradient

	_arg0 = (*C.GtkGradient)(unsafe.Pointer(g.Native()))

	var _cret *C.char

	cret = C.gtk_gradient_to_string(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Unref decreases the reference count of @gradient, freeing its memory if the
// reference count reaches 0.
func (g *Gradient) Unref() {
	var _arg0 *C.GtkGradient

	_arg0 = (*C.GtkGradient)(unsafe.Pointer(g.Native()))

	C.gtk_gradient_unref(_arg0)
}

// SymbolicColor: gtkSymbolicColor is a boxed type that represents a symbolic
// color. It is the result of parsing a [color
// expression][gtkcssprovider-symbolic-colors]. To obtain the color represented
// by a GtkSymbolicColor, it has to be resolved with
// gtk_symbolic_color_resolve(), which replaces all symbolic color references by
// the colors they refer to (in a given context) and evaluates mix, shade and
// other expressions, resulting in a RGBA value.
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
	if ptr == nil {
		return nil
	}

	return (*SymbolicColor)(ptr)
}

func marshalSymbolicColor(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapSymbolicColor(unsafe.Pointer(b)), nil
}

// NewSymbolicColorAlpha constructs a struct SymbolicColor.
func NewSymbolicColorAlpha(color *SymbolicColor, factor float64) *SymbolicColor {
	var _arg1 *C.GtkSymbolicColor
	var _arg2 C.gdouble

	_arg1 = (*C.GtkSymbolicColor)(unsafe.Pointer(color.Native()))
	_arg2 = C.gdouble(factor)

	var _cret *C.GtkSymbolicColor

	cret = C.gtk_symbolic_color_new_alpha(_arg1, _arg2)

	var _symbolicColor *SymbolicColor

	_symbolicColor = WrapSymbolicColor(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_symbolicColor, func(v *SymbolicColor) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _symbolicColor
}

// NewSymbolicColorLiteral constructs a struct SymbolicColor.
func NewSymbolicColorLiteral(color *gdk.RGBA) *SymbolicColor {
	var _arg1 *C.GdkRGBA

	_arg1 = (*C.GdkRGBA)(unsafe.Pointer(color.Native()))

	var _cret *C.GtkSymbolicColor

	cret = C.gtk_symbolic_color_new_literal(_arg1)

	var _symbolicColor *SymbolicColor

	_symbolicColor = WrapSymbolicColor(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_symbolicColor, func(v *SymbolicColor) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _symbolicColor
}

// NewSymbolicColorMix constructs a struct SymbolicColor.
func NewSymbolicColorMix(color1 *SymbolicColor, color2 *SymbolicColor, factor float64) *SymbolicColor {
	var _arg1 *C.GtkSymbolicColor
	var _arg2 *C.GtkSymbolicColor
	var _arg3 C.gdouble

	_arg1 = (*C.GtkSymbolicColor)(unsafe.Pointer(color1.Native()))
	_arg2 = (*C.GtkSymbolicColor)(unsafe.Pointer(color2.Native()))
	_arg3 = C.gdouble(factor)

	var _cret *C.GtkSymbolicColor

	cret = C.gtk_symbolic_color_new_mix(_arg1, _arg2, _arg3)

	var _symbolicColor *SymbolicColor

	_symbolicColor = WrapSymbolicColor(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_symbolicColor, func(v *SymbolicColor) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _symbolicColor
}

// NewSymbolicColorName constructs a struct SymbolicColor.
func NewSymbolicColorName(name string) *SymbolicColor {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.GtkSymbolicColor

	cret = C.gtk_symbolic_color_new_name(_arg1)

	var _symbolicColor *SymbolicColor

	_symbolicColor = WrapSymbolicColor(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_symbolicColor, func(v *SymbolicColor) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _symbolicColor
}

// NewSymbolicColorShade constructs a struct SymbolicColor.
func NewSymbolicColorShade(color *SymbolicColor, factor float64) *SymbolicColor {
	var _arg1 *C.GtkSymbolicColor
	var _arg2 C.gdouble

	_arg1 = (*C.GtkSymbolicColor)(unsafe.Pointer(color.Native()))
	_arg2 = C.gdouble(factor)

	var _cret *C.GtkSymbolicColor

	cret = C.gtk_symbolic_color_new_shade(_arg1, _arg2)

	var _symbolicColor *SymbolicColor

	_symbolicColor = WrapSymbolicColor(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_symbolicColor, func(v *SymbolicColor) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _symbolicColor
}

// NewSymbolicColorWin32 constructs a struct SymbolicColor.
func NewSymbolicColorWin32(themeClass string, id int) *SymbolicColor {
	var _arg1 *C.gchar
	var _arg2 C.gint

	_arg1 = (*C.gchar)(C.CString(themeClass))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(id)

	var _cret *C.GtkSymbolicColor

	cret = C.gtk_symbolic_color_new_win32(_arg1, _arg2)

	var _symbolicColor *SymbolicColor

	_symbolicColor = WrapSymbolicColor(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_symbolicColor, func(v *SymbolicColor) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _symbolicColor
}

// Native returns the underlying C source pointer.
func (s *SymbolicColor) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// Ref increases the reference count of @color
func (c *SymbolicColor) Ref() *SymbolicColor {
	var _arg0 *C.GtkSymbolicColor

	_arg0 = (*C.GtkSymbolicColor)(unsafe.Pointer(c.Native()))

	var _cret *C.GtkSymbolicColor

	cret = C.gtk_symbolic_color_ref(_arg0)

	var _symbolicColor *SymbolicColor

	_symbolicColor = WrapSymbolicColor(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_symbolicColor, func(v *SymbolicColor) {
		C.free(unsafe.Pointer(v.Native()))
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
func (c *SymbolicColor) Resolve(props StyleProperties) (gdk.RGBA, bool) {
	var _arg0 *C.GtkSymbolicColor
	var _arg1 *C.GtkStyleProperties

	_arg0 = (*C.GtkSymbolicColor)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkStyleProperties)(unsafe.Pointer(props.Native()))

	var _resolvedColor gdk.RGBA
	var _cret C.gboolean

	cret = C.gtk_symbolic_color_resolve(_arg0, _arg1, (*C.GdkRGBA)(unsafe.Pointer(&_resolvedColor)))

	var _ok bool

	if _cret {
		_ok = true
	}

	return _resolvedColor, _ok
}

// String converts the given @color to a string representation. This is useful
// both for debugging and for serialization of strings. The format of the string
// may change between different versions of GTK, but it is guaranteed that the
// GTK css parser is able to read the string and create the same symbolic color
// from it.
func (c *SymbolicColor) String() string {
	var _arg0 *C.GtkSymbolicColor

	_arg0 = (*C.GtkSymbolicColor)(unsafe.Pointer(c.Native()))

	var _cret *C.char

	cret = C.gtk_symbolic_color_to_string(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Unref decreases the reference count of @color, freeing its memory if the
// reference count reaches 0.
func (c *SymbolicColor) Unref() {
	var _arg0 *C.GtkSymbolicColor

	_arg0 = (*C.GtkSymbolicColor)(unsafe.Pointer(c.Native()))

	C.gtk_symbolic_color_unref(_arg0)
}