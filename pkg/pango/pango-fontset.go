// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <pango/pango.h>
// gboolean _gotk4_pango1_FontsetForEachFunc(PangoFontset*, PangoFont*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_fontset_get_type()), F: marshalFontsetter},
		{T: externglib.Type(C.pango_fontset_simple_get_type()), F: marshalFontsetSimpler},
	})
}

// FontsetForEachFunc: callback used by pango_fontset_foreach() when enumerating
// fonts in a fontset.
type FontsetForEachFunc func(fontset Fontsetter, font Fonter) (ok bool)

//export _gotk4_pango1_FontsetForEachFunc
func _gotk4_pango1_FontsetForEachFunc(arg0 *C.PangoFontset, arg1 *C.PangoFont, arg2 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var fontset Fontsetter // out
	var font Fonter        // out

	{
		objptr := unsafe.Pointer(arg0)
		if objptr == nil {
			panic("object of type pango.Fontsetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(Fontsetter)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not pango.Fontsetter")
		}
		fontset = rv
	}
	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type pango.Fonter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(Fonter)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not pango.Fonter")
		}
		font = rv
	}

	fn := v.(FontsetForEachFunc)
	ok := fn(fontset, font)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// FontsetOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FontsetOverrider interface {
	// ForEach iterates through all the fonts in a fontset, calling func for
	// each one.
	//
	// If func returns TRUE, that stops the iteration.
	//
	// The function takes the following parameters:
	//
	//    - fn: callback function.
	//
	ForEach(fn FontsetForEachFunc)
	// Font returns the font in the fontset that contains the best glyph for a
	// Unicode character.
	//
	// The function takes the following parameters:
	//
	//    - wc: unicode character.
	//
	// The function returns the following values:
	//
	//    - font: PangoFont. The caller must call g_object_unref() when finished
	//      with the font.
	//
	Font(wc uint) Fonter
	// The function returns the following values:
	//
	Language() *Language
	// Metrics: get overall metric information for the fonts in the fontset.
	//
	// The function returns the following values:
	//
	//    - fontMetrics object. The caller must call pango_font_metrics_unref()
	//      when finished using the object.
	//
	Metrics() *FontMetrics
}

// Fontset: PangoFontset represents a set of PangoFont to use when rendering
// text.
//
// A PAngoFontset is the result of resolving a PangoFontDescription against a
// particular PangoContext. It has operations for finding the component font for
// a particular Unicode character, and for finding a composite set of metrics
// for the entire fontset.
type Fontset struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*Fontset)(nil)
)

// Fontsetter describes types inherited from class Fontset.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Fontsetter interface {
	externglib.Objector
	baseFontset() *Fontset
}

var _ Fontsetter = (*Fontset)(nil)

func wrapFontset(obj *externglib.Object) *Fontset {
	return &Fontset{
		Object: obj,
	}
}

func marshalFontsetter(p uintptr) (interface{}, error) {
	return wrapFontset(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (fontset *Fontset) baseFontset() *Fontset {
	return fontset
}

// BaseFontset returns the underlying base object.
func BaseFontset(obj Fontsetter) *Fontset {
	return obj.baseFontset()
}

// ForEach iterates through all the fonts in a fontset, calling func for each
// one.
//
// If func returns TRUE, that stops the iteration.
//
// The function takes the following parameters:
//
//    - fn: callback function.
//
func (fontset *Fontset) ForEach(fn FontsetForEachFunc) {
	var _arg0 *C.PangoFontset           // out
	var _arg1 C.PangoFontsetForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(fontset.Native()))
	_arg1 = (*[0]byte)(C._gotk4_pango1_FontsetForEachFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))
	defer gbox.Delete(uintptr(_arg2))

	C.pango_fontset_foreach(_arg0, _arg1, _arg2)
	runtime.KeepAlive(fontset)
	runtime.KeepAlive(fn)
}

// Font returns the font in the fontset that contains the best glyph for a
// Unicode character.
//
// The function takes the following parameters:
//
//    - wc: unicode character.
//
// The function returns the following values:
//
//    - font: PangoFont. The caller must call g_object_unref() when finished with
//      the font.
//
func (fontset *Fontset) Font(wc uint) Fonter {
	var _arg0 *C.PangoFontset // out
	var _arg1 C.guint         // out
	var _cret *C.PangoFont    // in

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(fontset.Native()))
	_arg1 = C.guint(wc)

	_cret = C.pango_fontset_get_font(_arg0, _arg1)
	runtime.KeepAlive(fontset)
	runtime.KeepAlive(wc)

	var _font Fonter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type pango.Fonter is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.Cast()
		rv, ok := casted.(Fonter)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not pango.Fonter")
		}
		_font = rv
	}

	return _font
}

// Metrics: get overall metric information for the fonts in the fontset.
//
// The function returns the following values:
//
//    - fontMetrics object. The caller must call pango_font_metrics_unref() when
//      finished using the object.
//
func (fontset *Fontset) Metrics() *FontMetrics {
	var _arg0 *C.PangoFontset     // out
	var _cret *C.PangoFontMetrics // in

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(fontset.Native()))

	_cret = C.pango_fontset_get_metrics(_arg0)
	runtime.KeepAlive(fontset)

	var _fontMetrics *FontMetrics // out

	_fontMetrics = (*FontMetrics)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_fontMetrics)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_font_metrics_unref((*C.PangoFontMetrics)(intern.C))
		},
	)

	return _fontMetrics
}

// FontsetSimple: PangoFontsetSimple is a implementation of the abstract
// PangoFontset base class as an array of fonts.
//
// When creating a PangoFontsetSimple, you have to provide the array of fonts
// that make up the fontset.
type FontsetSimple struct {
	Fontset
}

var (
	_ Fontsetter = (*FontsetSimple)(nil)
)

func wrapFontsetSimple(obj *externglib.Object) *FontsetSimple {
	return &FontsetSimple{
		Fontset: Fontset{
			Object: obj,
		},
	}
}

func marshalFontsetSimpler(p uintptr) (interface{}, error) {
	return wrapFontsetSimple(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFontsetSimple creates a new PangoFontsetSimple for the given language.
//
// The function takes the following parameters:
//
//    - language: PangoLanguage tag.
//
// The function returns the following values:
//
//    - fontsetSimple: newly allocated PangoFontsetSimple, which should be freed
//      with g_object_unref().
//
func NewFontsetSimple(language *Language) *FontsetSimple {
	var _arg1 *C.PangoLanguage      // out
	var _cret *C.PangoFontsetSimple // in

	_arg1 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))

	_cret = C.pango_fontset_simple_new(_arg1)
	runtime.KeepAlive(language)

	var _fontsetSimple *FontsetSimple // out

	_fontsetSimple = wrapFontsetSimple(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _fontsetSimple
}

// Append adds a font to the fontset.
//
// The function takes the following parameters:
//
//    - font: PangoFont.
//
func (fontset *FontsetSimple) Append(font Fonter) {
	var _arg0 *C.PangoFontsetSimple // out
	var _arg1 *C.PangoFont          // out

	_arg0 = (*C.PangoFontsetSimple)(unsafe.Pointer(fontset.Native()))
	_arg1 = (*C.PangoFont)(unsafe.Pointer(font.Native()))

	C.pango_fontset_simple_append(_arg0, _arg1)
	runtime.KeepAlive(fontset)
	runtime.KeepAlive(font)
}

// Size returns the number of fonts in the fontset.
//
// The function returns the following values:
//
//    - gint: size of fontset.
//
func (fontset *FontsetSimple) Size() int {
	var _arg0 *C.PangoFontsetSimple // out
	var _cret C.int                 // in

	_arg0 = (*C.PangoFontsetSimple)(unsafe.Pointer(fontset.Native()))

	_cret = C.pango_fontset_simple_size(_arg0)
	runtime.KeepAlive(fontset)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
