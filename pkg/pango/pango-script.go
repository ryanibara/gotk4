// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_script_iter_get_type()), F: marshalScriptIter},
	})
}

// ScriptForUnichar looks up the script for a particular character.
//
// The script of a character is defined by Unicode Standard Annex \#24. No check
// is made for @ch being a valid Unicode character; if you pass in invalid
// character, the result is undefined.
//
// Note that while the return type of this function is declared as
// `PangoScript`, as of Pango 1.18, this function simply returns the return
// value of g_unichar_get_script(). Callers must be prepared to handle unknown
// values.
func ScriptForUnichar(ch uint32) Script {
	var arg1 C.gunichar

	arg1 = C.gunichar(ch)

	var cret C.PangoScript
	var ret1 Script

	cret = C.pango_script_for_unichar(ch)

	ret1 = Script(cret)

	return ret1
}

// ScriptGetSampleLanguage finds a language tag that is reasonably
// representative of @script.
//
// The language will usually be the most widely spoken or used language written
// in that script: for instance, the sample language for PANGO_SCRIPT_CYRILLIC
// is ru (Russian), the sample language for PANGO_SCRIPT_ARABIC is ar.
//
// For some scripts, no sample language will be returned because there is no
// language that is sufficiently representative. The best example of this is
// PANGO_SCRIPT_HAN, where various different variants of written Chinese,
// Japanese, and Korean all use significantly different sets of Han characters
// and forms of shared characters. No sample language can be provided for many
// historical scripts as well.
//
// As of 1.18, this function checks the environment variables `PANGO_LANGUAGE`
// and `LANGUAGE` (checked in that order) first. If one of them is set, it is
// parsed as a list of language tags separated by colons or other separators.
// This function will return the first language in the parsed list that Pango
// believes may use @script for writing. This last predicate is tested using
// [method@Pango.Language.includes_script]. This can be used to control Pango's
// font selection for non-primary languages. For example, a `PANGO_LANGUAGE`
// enviroment variable set to "en:fa" makes Pango choose fonts suitable for
// Persian (fa) instead of Arabic (ar) when a segment of Arabic text is found in
// an otherwise non-Arabic text. The same trick can be used to choose a default
// language for PANGO_SCRIPT_HAN when setting context language is not feasible.
func ScriptGetSampleLanguage(script Script) *Language {
	var arg1 C.PangoScript

	arg1 = (C.PangoScript)(script)

	var cret *C.PangoLanguage
	var ret1 *Language

	cret = C.pango_script_get_sample_language(script)

	ret1 = WrapLanguage(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Language) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// ScriptIter: a `PangoScriptIter` is used to iterate through a string and
// identify ranges in different scripts.
type ScriptIter struct {
	native C.PangoScriptIter
}

// WrapScriptIter wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapScriptIter(ptr unsafe.Pointer) *ScriptIter {
	if ptr == nil {
		return nil
	}

	return (*ScriptIter)(ptr)
}

func marshalScriptIter(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapScriptIter(unsafe.Pointer(b)), nil
}

// NewScriptIter constructs a struct ScriptIter.
func NewScriptIter(text string, length int) *ScriptIter {
	var arg1 *C.char
	var arg2 C.int

	arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.int(length)

	var cret *C.PangoScriptIter
	var ret1 *ScriptIter

	cret = C.pango_script_iter_new(text, length)

	ret1 = WrapScriptIter(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *ScriptIter) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Native returns the underlying C source pointer.
func (s *ScriptIter) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// Free frees a ScriptIter created with pango_script_iter_new().
func (i *ScriptIter) Free() {
	var arg0 *C.PangoScriptIter

	arg0 = (*C.PangoScriptIter)(unsafe.Pointer(i.Native()))

	C.pango_script_iter_free(arg0)
}

// Range gets information about the range to which @iter currently points. The
// range is the set of locations p where *start <= p < *end. (That is, it
// doesn't include the character stored at *end)
//
// Note that while the type of the @script argument is declared as PangoScript,
// as of Pango 1.18, this function simply returns GUnicodeScript values. Callers
// must be prepared to handle unknown values.
func (i *ScriptIter) Range() (start string, end string, script Script) {
	var arg0 *C.PangoScriptIter

	arg0 = (*C.PangoScriptIter)(unsafe.Pointer(i.Native()))

	var arg1 *C.char
	var ret1 string
	var arg2 *C.char
	var ret2 string
	var arg3 C.PangoScript
	var ret3 *Script

	C.pango_script_iter_get_range(arg0, &arg1, &arg2, &arg3)

	ret1 = C.GoString(arg1)
	defer C.free(unsafe.Pointer(arg1))
	ret2 = C.GoString(arg2)
	defer C.free(unsafe.Pointer(arg2))
	ret3 = *Script(arg3)

	return ret1, ret2, ret3
}

// Next advances a ScriptIter to the next range. If @iter is already at the end,
// it is left unchanged and false is returned.
func (i *ScriptIter) Next() bool {
	var arg0 *C.PangoScriptIter

	arg0 = (*C.PangoScriptIter)(unsafe.Pointer(i.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.pango_script_iter_next(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}
