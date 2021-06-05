// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_fontset_get_type()), F: marshalFontset},
		{T: externglib.Type(C.pango_fontset_simple_get_type()), F: marshalFontsetSimple},
	})
}

// FontsetForeachFunc: callback used by pango_fontset_foreach() when enumerating
// fonts in a fontset.
type FontsetForeachFunc func(fontset Fontset, font Font) bool

//export gotk4_FontsetForeachFunc
func gotk4_FontsetForeachFunc(arg0 *C.PangoFontset, arg1 *C.PangoFont, arg2 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(FontsetForeachFunc)
	ret := fn(fontset, font, userData)

	if ret {
		cret = C.gboolean(1)
	}

	return cret
}

// Fontset: a `PangoFontset` represents a set of `PangoFont` to use when
// rendering text.
//
// A `PAngoFontset` is the result of resolving a `PangoFontDescription` against
// a particular `PangoContext`. It has operations for finding the component font
// for a particular Unicode character, and for finding a composite set of
// metrics for the entire fontset.
type Fontset interface {
	gextras.Objector

	// Foreach iterates through all the fonts in a fontset, calling @func for
	// each one.
	//
	// If @func returns true, that stops the iteration.
	Foreach(fn FontsetForeachFunc)
	// Font returns the font in the fontset that contains the best glyph for a
	// Unicode character.
	Font(wc uint) Font
	// Metrics: get overall metric information for the fonts in the fontset.
	Metrics() *FontMetrics
}

// fontset implements the Fontset interface.
type fontset struct {
	gextras.Objector
}

var _ Fontset = (*fontset)(nil)

// WrapFontset wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontset(obj *externglib.Object) Fontset {
	return Fontset{
		Objector: obj,
	}
}

func marshalFontset(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontset(obj), nil
}

// Foreach iterates through all the fonts in a fontset, calling @func for
// each one.
//
// If @func returns true, that stops the iteration.
func (f fontset) Foreach(fn FontsetForeachFunc) {
	var arg0 *C.PangoFontset

	arg0 = (*C.PangoFontset)(unsafe.Pointer(f.Native()))

	C.pango_fontset_foreach(arg0, fn, data)
}

// Font returns the font in the fontset that contains the best glyph for a
// Unicode character.
func (f fontset) Font(wc uint) Font {
	var arg0 *C.PangoFontset
	var arg1 C.guint

	arg0 = (*C.PangoFontset)(unsafe.Pointer(f.Native()))
	arg1 = C.guint(wc)

	var cret *C.PangoFont
	var ret1 Font

	cret = C.pango_fontset_get_font(arg0, wc)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Font)

	return ret1
}

// Metrics: get overall metric information for the fonts in the fontset.
func (f fontset) Metrics() *FontMetrics {
	var arg0 *C.PangoFontset

	arg0 = (*C.PangoFontset)(unsafe.Pointer(f.Native()))

	var cret *C.PangoFontMetrics
	var ret1 *FontMetrics

	cret = C.pango_fontset_get_metrics(arg0)

	ret1 = WrapFontMetrics(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *FontMetrics) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// FontsetSimple: `PangoFontsetSimple` is a implementation of the abstract
// `PangoFontset` base class as an array of fonts.
//
// When creating a `PangoFontsetSimple`, you have to provide the array of fonts
// that make up the fontset.
type FontsetSimple interface {
	Fontset

	// Append adds a font to the fontset.
	Append(font Font)
	// Size returns the number of fonts in the fontset.
	Size() int
}

// fontsetSimple implements the FontsetSimple interface.
type fontsetSimple struct {
	Fontset
}

var _ FontsetSimple = (*fontsetSimple)(nil)

// WrapFontsetSimple wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontsetSimple(obj *externglib.Object) FontsetSimple {
	return FontsetSimple{
		Fontset: WrapFontset(obj),
	}
}

func marshalFontsetSimple(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontsetSimple(obj), nil
}

// NewFontsetSimple constructs a class FontsetSimple.
func NewFontsetSimple(language *Language) FontsetSimple {
	var arg1 *C.PangoLanguage

	arg1 = (*C.PangoLanguage)(unsafe.Pointer(language.Native()))

	var cret C.PangoFontsetSimple
	var ret1 FontsetSimple

	cret = C.pango_fontset_simple_new(language)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(FontsetSimple)

	return ret1
}

// Append adds a font to the fontset.
func (f fontsetSimple) Append(font Font) {
	var arg0 *C.PangoFontsetSimple
	var arg1 *C.PangoFont

	arg0 = (*C.PangoFontsetSimple)(unsafe.Pointer(f.Native()))
	arg1 = (*C.PangoFont)(unsafe.Pointer(font.Native()))

	C.pango_fontset_simple_append(arg0, font)
}

// Size returns the number of fonts in the fontset.
func (f fontsetSimple) Size() int {
	var arg0 *C.PangoFontsetSimple

	arg0 = (*C.PangoFontsetSimple)(unsafe.Pointer(f.Native()))

	var cret C.int
	var ret1 int

	cret = C.pango_fontset_simple_size(arg0)

	ret1 = C.int(cret)

	return ret1
}
