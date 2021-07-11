// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_script_get_type()), F: marshalScript},
		{T: externglib.Type(C.pango_script_iter_get_type()), F: marshalScriptIter},
	})
}

// Script: `PangoScript` enumeration identifies different writing systems.
//
// The values correspond to the names as defined in the Unicode standard. See
// Unicode Standard Annex 24: Script names
// (http://www.unicode.org/reports/tr24/)
//
// Note that this enumeration is deprecated and will not be updated to include
// values in newer versions of the Unicode standard. Applications should use the
// `GUnicodeScript` enumeration instead, whose values are interchangeable with
// `PangoScript`.
type Script int

const (
	// InvalidCode: value never returned from pango_script_for_unichar()
	ScriptInvalidCode Script = -1
	// Common: character used by multiple different scripts
	ScriptCommon Script = 0
	// Inherited: mark glyph that takes its script from the base glyph to which
	// it is attached
	ScriptInherited Script = 1
	// Arabic: arabic
	ScriptArabic Script = 2
	// Armenian: armenian
	ScriptArmenian Script = 3
	// Bengali: bengali
	ScriptBengali Script = 4
	// Bopomofo: bopomofo
	ScriptBopomofo Script = 5
	// Cherokee: cherokee
	ScriptCherokee Script = 6
	// Coptic: coptic
	ScriptCoptic Script = 7
	// Cyrillic: cyrillic
	ScriptCyrillic Script = 8
	// Deseret: deseret
	ScriptDeseret Script = 9
	// Devanagari: devanagari
	ScriptDevanagari Script = 10
	// Ethiopic: ethiopic
	ScriptEthiopic Script = 11
	// Georgian: georgian
	ScriptGeorgian Script = 12
	// Gothic: gothic
	ScriptGothic Script = 13
	// Greek: greek
	ScriptGreek Script = 14
	// Gujarati: gujarati
	ScriptGujarati Script = 15
	// Gurmukhi: gurmukhi
	ScriptGurmukhi Script = 16
	// Han: han
	ScriptHan Script = 17
	// Hangul: hangul
	ScriptHangul Script = 18
	// Hebrew: hebrew
	ScriptHebrew Script = 19
	// Hiragana: hiragana
	ScriptHiragana Script = 20
	// Kannada: kannada
	ScriptKannada Script = 21
	// Katakana: katakana
	ScriptKatakana Script = 22
	// Khmer: khmer
	ScriptKhmer Script = 23
	// Lao: lao
	ScriptLao Script = 24
	// Latin: latin
	ScriptLatin Script = 25
	// Malayalam: malayalam
	ScriptMalayalam Script = 26
	// Mongolian: mongolian
	ScriptMongolian Script = 27
	// Myanmar: myanmar
	ScriptMyanmar Script = 28
	// Ogham: ogham
	ScriptOgham Script = 29
	// OldItalic: old Italic
	ScriptOldItalic Script = 30
	// Oriya: oriya
	ScriptOriya Script = 31
	// Runic: runic
	ScriptRunic Script = 32
	// Sinhala: sinhala
	ScriptSinhala Script = 33
	// Syriac: syriac
	ScriptSyriac Script = 34
	// Tamil: tamil
	ScriptTamil Script = 35
	// Telugu: telugu
	ScriptTelugu Script = 36
	// Thaana: thaana
	ScriptThaana Script = 37
	// Thai: thai
	ScriptThai Script = 38
	// Tibetan: tibetan
	ScriptTibetan Script = 39
	// CanadianAboriginal: canadian Aboriginal
	ScriptCanadianAboriginal Script = 40
	// Yi: yi
	ScriptYi Script = 41
	// Tagalog: tagalog
	ScriptTagalog Script = 42
	// Hanunoo: hanunoo
	ScriptHanunoo Script = 43
	// Buhid: buhid
	ScriptBuhid Script = 44
	// Tagbanwa: tagbanwa
	ScriptTagbanwa Script = 45
	// Braille: braille
	ScriptBraille Script = 46
	// Cypriot: cypriot
	ScriptCypriot Script = 47
	// Limbu: limbu
	ScriptLimbu Script = 48
	// Osmanya: osmanya
	ScriptOsmanya Script = 49
	// Shavian: shavian
	ScriptShavian Script = 50
	// LinearB: linear B
	ScriptLinearB Script = 51
	// TaiLe: tai Le
	ScriptTaiLe Script = 52
	// Ugaritic: ugaritic
	ScriptUgaritic Script = 53
	// NewTaiLue: new Tai Lue. Since 1.10
	ScriptNewTaiLue Script = 54
	// Buginese: buginese. Since 1.10
	ScriptBuginese Script = 55
	// Glagolitic: glagolitic. Since 1.10
	ScriptGlagolitic Script = 56
	// Tifinagh: tifinagh. Since 1.10
	ScriptTifinagh Script = 57
	// SylotiNagri: syloti Nagri. Since 1.10
	ScriptSylotiNagri Script = 58
	// OldPersian: old Persian. Since 1.10
	ScriptOldPersian Script = 59
	// Kharoshthi: kharoshthi. Since 1.10
	ScriptKharoshthi Script = 60
	// Unknown: unassigned code point. Since 1.14
	ScriptUnknown Script = 61
	// Balinese: balinese. Since 1.14
	ScriptBalinese Script = 62
	// Cuneiform: cuneiform. Since 1.14
	ScriptCuneiform Script = 63
	// Phoenician: phoenician. Since 1.14
	ScriptPhoenician Script = 64
	// PhagsPa: phags-pa. Since 1.14
	ScriptPhagsPa Script = 65
	// Nko: n'Ko. Since 1.14
	ScriptNko Script = 66
	// KayahLi: kayah Li. Since 1.20.1
	ScriptKayahLi Script = 67
	// Lepcha: lepcha. Since 1.20.1
	ScriptLepcha Script = 68
	// Rejang: rejang. Since 1.20.1
	ScriptRejang Script = 69
	// Sundanese: sundanese. Since 1.20.1
	ScriptSundanese Script = 70
	// Saurashtra: saurashtra. Since 1.20.1
	ScriptSaurashtra Script = 71
	// Cham: cham. Since 1.20.1
	ScriptCham Script = 72
	// OlChiki: ol Chiki. Since 1.20.1
	ScriptOlChiki Script = 73
	// Vai: vai. Since 1.20.1
	ScriptVai Script = 74
	// Carian: carian. Since 1.20.1
	ScriptCarian Script = 75
	// Lycian: lycian. Since 1.20.1
	ScriptLycian Script = 76
	// Lydian: lydian. Since 1.20.1
	ScriptLydian Script = 77
	// Batak: batak. Since 1.32
	ScriptBatak Script = 78
	// Brahmi: brahmi. Since 1.32
	ScriptBrahmi Script = 79
	// Mandaic: mandaic. Since 1.32
	ScriptMandaic Script = 80
	// Chakma: chakma. Since: 1.32
	ScriptChakma Script = 81
	// MeroiticCursive: meroitic Cursive. Since: 1.32
	ScriptMeroiticCursive Script = 82
	// MeroiticHieroglyphs: meroitic Hieroglyphs. Since: 1.32
	ScriptMeroiticHieroglyphs Script = 83
	// Miao: miao. Since: 1.32
	ScriptMiao Script = 84
	// Sharada: sharada. Since: 1.32
	ScriptSharada Script = 85
	// SoraSompeng: sora Sompeng. Since: 1.32
	ScriptSoraSompeng Script = 86
	// Takri: takri. Since: 1.32
	ScriptTakri Script = 87
	// BassaVah: bassa. Since: 1.40
	ScriptBassaVah Script = 88
	// CaucasianAlbanian: caucasian Albanian. Since: 1.40
	ScriptCaucasianAlbanian Script = 89
	// Duployan: duployan. Since: 1.40
	ScriptDuployan Script = 90
	// Elbasan: elbasan. Since: 1.40
	ScriptElbasan Script = 91
	// Grantha: grantha. Since: 1.40
	ScriptGrantha Script = 92
	// Khojki: kjohki. Since: 1.40
	ScriptKhojki Script = 93
	// Khudawadi: khudawadi, Sindhi. Since: 1.40
	ScriptKhudawadi Script = 94
	// LinearA: linear A. Since: 1.40
	ScriptLinearA Script = 95
	// Mahajani: mahajani. Since: 1.40
	ScriptMahajani Script = 96
	// Manichaean: manichaean. Since: 1.40
	ScriptManichaean Script = 97
	// MendeKikakui: mende Kikakui. Since: 1.40
	ScriptMendeKikakui Script = 98
	// Modi: modi. Since: 1.40
	ScriptModi Script = 99
	// Mro: mro. Since: 1.40
	ScriptMro Script = 100
	// Nabataean: nabataean. Since: 1.40
	ScriptNabataean Script = 101
	// OldNorthArabian: old North Arabian. Since: 1.40
	ScriptOldNorthArabian Script = 102
	// OldPermic: old Permic. Since: 1.40
	ScriptOldPermic Script = 103
	// PahawhHmong: pahawh Hmong. Since: 1.40
	ScriptPahawhHmong Script = 104
	// Palmyrene: palmyrene. Since: 1.40
	ScriptPalmyrene Script = 105
	// PauCinHau: pau Cin Hau. Since: 1.40
	ScriptPauCinHau Script = 106
	// PsalterPahlavi: psalter Pahlavi. Since: 1.40
	ScriptPsalterPahlavi Script = 107
	// Siddham: siddham. Since: 1.40
	ScriptSiddham Script = 108
	// Tirhuta: tirhuta. Since: 1.40
	ScriptTirhuta Script = 109
	// WarangCiti: warang Citi. Since: 1.40
	ScriptWarangCiti Script = 110
	// Ahom: ahom. Since: 1.40
	ScriptAhom Script = 111
	// AnatolianHieroglyphs: anatolian Hieroglyphs. Since: 1.40
	ScriptAnatolianHieroglyphs Script = 112
	// Hatran: hatran. Since: 1.40
	ScriptHatran Script = 113
	// Multani: multani. Since: 1.40
	ScriptMultani Script = 114
	// OldHungarian: old Hungarian. Since: 1.40
	ScriptOldHungarian Script = 115
	// Signwriting: signwriting. Since: 1.40
	ScriptSignwriting Script = 116
)

func marshalScript(p uintptr) (interface{}, error) {
	return Script(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ScriptIter: `PangoScriptIter` is used to iterate through a string and
// identify ranges in different scripts.
type ScriptIter struct {
	native C.PangoScriptIter
}

func marshalScriptIter(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*ScriptIter)(unsafe.Pointer(b)), nil
}

// NewScriptIter constructs a struct ScriptIter.
func NewScriptIter(text string, length int) *ScriptIter {
	var _arg1 *C.char            // out
	var _arg2 C.int              // out
	var _cret *C.PangoScriptIter // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(length)

	_cret = C.pango_script_iter_new(_arg1, _arg2)

	var _scriptIter *ScriptIter // out

	_scriptIter = (*ScriptIter)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_scriptIter, func(v *ScriptIter) {
		C.pango_script_iter_free((*C.PangoScriptIter)(unsafe.Pointer(v)))
	})

	return _scriptIter
}

// Native returns the underlying C source pointer.
func (s *ScriptIter) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// Free frees a ScriptIter created with pango_script_iter_new().
func (iter *ScriptIter) free() {
	var _arg0 *C.PangoScriptIter // out

	_arg0 = (*C.PangoScriptIter)(unsafe.Pointer(iter))

	C.pango_script_iter_free(_arg0)
}

// Range gets information about the range to which @iter currently points. The
// range is the set of locations p where *start <= p < *end. (That is, it
// doesn't include the character stored at *end)
//
// Note that while the type of the @script argument is declared as PangoScript,
// as of Pango 1.18, this function simply returns GUnicodeScript values. Callers
// must be prepared to handle unknown values.
func (iter *ScriptIter) Range() (start string, end string, script Script) {
	var _arg0 *C.PangoScriptIter // out
	var _arg1 *C.char            // in
	var _arg2 *C.char            // in
	var _arg3 C.PangoScript      // in

	_arg0 = (*C.PangoScriptIter)(unsafe.Pointer(iter))

	C.pango_script_iter_get_range(_arg0, &_arg1, &_arg2, &_arg3)

	var _start string  // out
	var _end string    // out
	var _script Script // out

	_start = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	defer C.free(unsafe.Pointer(_arg1))
	_end = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	defer C.free(unsafe.Pointer(_arg2))
	_script = Script(_arg3)

	return _start, _end, _script
}

// Next advances a ScriptIter to the next range. If @iter is already at the end,
// it is left unchanged and false is returned.
func (iter *ScriptIter) Next() bool {
	var _arg0 *C.PangoScriptIter // out
	var _cret C.gboolean         // in

	_arg0 = (*C.PangoScriptIter)(unsafe.Pointer(iter))

	_cret = C.pango_script_iter_next(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
