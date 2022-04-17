// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <pango/pango.h>
// extern PangoFont* _gotk4_pango1_FontsetClass_get_font(PangoFontset*, guint);
// extern PangoFontMetrics* _gotk4_pango1_FontsetClass_get_metrics(PangoFontset*);
// extern PangoLanguage* _gotk4_pango1_FontsetClass_get_language(PangoFontset*);
// extern gboolean _gotk4_pango1_FontsetForEachFunc(PangoFontset*, PangoFont*, gpointer);
import "C"

// glib.Type values for pango-fontset.go.
var (
	GTypeFontset       = externglib.Type(C.pango_fontset_get_type())
	GTypeFontsetSimple = externglib.Type(C.pango_fontset_simple_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeFontset, F: marshalFontset},
		{T: GTypeFontsetSimple, F: marshalFontsetSimple},
	})
}

// FontsetForEachFunc: callback used by pango_fontset_foreach() when enumerating
// fonts in a fontset.
type FontsetForEachFunc func(fontset Fontsetter, font Fonter) (ok bool)

//export _gotk4_pango1_FontsetForEachFunc
func _gotk4_pango1_FontsetForEachFunc(arg1 *C.PangoFontset, arg2 *C.PangoFont, arg3 C.gpointer) (cret C.gboolean) {
	var fn FontsetForEachFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(FontsetForEachFunc)
	}

	var _fontset Fontsetter // out
	var _font Fonter        // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type pango.Fontsetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Fontsetter)
			return ok
		})
		rv, ok := casted.(Fontsetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.Fontsetter")
		}
		_fontset = rv
	}
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type pango.Fonter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Fonter)
			return ok
		})
		rv, ok := casted.(Fonter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.Fonter")
		}
		_font = rv
	}

	ok := fn(_fontset, _font)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// FontsetOverrider contains methods that are overridable.
type FontsetOverrider interface {
	externglib.Objector
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

// WrapFontsetOverrider wraps the FontsetOverrider
// interface implementation to access the instance methods.
func WrapFontsetOverrider(obj FontsetOverrider) *Fontset {
	return wrapFontset(externglib.BaseObject(obj))
}

// Fontset: PangoFontset represents a set of PangoFont to use when rendering
// text.
//
// A PAngoFontset is the result of resolving a PangoFontDescription against a
// particular PangoContext. It has operations for finding the component font for
// a particular Unicode character, and for finding a composite set of metrics
// for the entire fontset.
type Fontset struct {
	_ [0]func() // equal guard
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

func classInitFontsetter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.PangoFontsetClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.PangoFontsetClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Font(wc uint) Fonter }); ok {
		pclass.get_font = (*[0]byte)(C._gotk4_pango1_FontsetClass_get_font)
	}

	if _, ok := goval.(interface{ Language() *Language }); ok {
		pclass.get_language = (*[0]byte)(C._gotk4_pango1_FontsetClass_get_language)
	}

	if _, ok := goval.(interface{ Metrics() *FontMetrics }); ok {
		pclass.get_metrics = (*[0]byte)(C._gotk4_pango1_FontsetClass_get_metrics)
	}
}

//export _gotk4_pango1_FontsetClass_get_font
func _gotk4_pango1_FontsetClass_get_font(arg0 *C.PangoFontset, arg1 C.guint) (cret *C.PangoFont) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Font(wc uint) Fonter })

	var _wc uint // out

	_wc = uint(arg1)

	font := iface.Font(_wc)

	cret = (*C.PangoFont)(unsafe.Pointer(externglib.InternObject(font).Native()))
	C.g_object_ref(C.gpointer(externglib.InternObject(font).Native()))

	return cret
}

//export _gotk4_pango1_FontsetClass_get_language
func _gotk4_pango1_FontsetClass_get_language(arg0 *C.PangoFontset) (cret *C.PangoLanguage) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Language() *Language })

	language := iface.Language()

	cret = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(language)), nil)

	return cret
}

//export _gotk4_pango1_FontsetClass_get_metrics
func _gotk4_pango1_FontsetClass_get_metrics(arg0 *C.PangoFontset) (cret *C.PangoFontMetrics) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Metrics() *FontMetrics })

	fontMetrics := iface.Metrics()

	cret = (*C.PangoFontMetrics)(gextras.StructNative(unsafe.Pointer(fontMetrics)))

	return cret
}

func wrapFontset(obj *externglib.Object) *Fontset {
	return &Fontset{
		Object: obj,
	}
}

func marshalFontset(p uintptr) (interface{}, error) {
	return wrapFontset(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (fontset *Fontset) baseFontset() *Fontset {
	return fontset
}

// BaseFontset returns the underlying base object from the
// interface.
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

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(externglib.InternObject(fontset).Native()))
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

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(externglib.InternObject(fontset).Native()))
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
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Fonter)
			return ok
		})
		rv, ok := casted.(Fonter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.Fonter")
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

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(externglib.InternObject(fontset).Native()))

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

// FontsetSimpleOverrider contains methods that are overridable.
type FontsetSimpleOverrider interface {
	externglib.Objector
}

// WrapFontsetSimpleOverrider wraps the FontsetSimpleOverrider
// interface implementation to access the instance methods.
func WrapFontsetSimpleOverrider(obj FontsetSimpleOverrider) *FontsetSimple {
	return wrapFontsetSimple(externglib.BaseObject(obj))
}

// FontsetSimple: PangoFontsetSimple is a implementation of the abstract
// PangoFontset base class as an array of fonts.
//
// When creating a PangoFontsetSimple, you have to provide the array of fonts
// that make up the fontset.
type FontsetSimple struct {
	_ [0]func() // equal guard
	Fontset
}

var (
	_ Fontsetter = (*FontsetSimple)(nil)
)

func classInitFontsetSimpler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFontsetSimple(obj *externglib.Object) *FontsetSimple {
	return &FontsetSimple{
		Fontset: Fontset{
			Object: obj,
		},
	}
}

func marshalFontsetSimple(p uintptr) (interface{}, error) {
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

	_arg0 = (*C.PangoFontsetSimple)(unsafe.Pointer(externglib.InternObject(fontset).Native()))
	_arg1 = (*C.PangoFont)(unsafe.Pointer(externglib.InternObject(font).Native()))

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

	_arg0 = (*C.PangoFontsetSimple)(unsafe.Pointer(externglib.InternObject(fontset).Native()))

	_cret = C.pango_fontset_simple_size(_arg0)
	runtime.KeepAlive(fontset)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
