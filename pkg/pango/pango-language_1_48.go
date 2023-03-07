// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// LanguageGetPreferred returns the list of languages that the user prefers.
//
// The list is specified by the PANGO_LANGUAGE or LANGUAGE environment
// variables, in order of preference. Note that this list does not necessarily
// include the language returned by pango.Language.GetDefault.
//
// When choosing language-specific resources, such as the sample text returned
// by pango.Language.GetSampleString(), you should first try the default
// language, followed by the languages returned by this function.
//
// The function returns the following values:
//
//    - language (optional): NULL-terminated array of PangoLanguage*.
//
func LanguageGetPreferred() *Language {
	var _cret **C.PangoLanguage // in

	_cret = C.pango_language_get_preferred()

	var _language *Language // out

	if _cret != nil {
		_language = (*Language)(gextras.NewStructNative(unsafe.Pointer((*_cret))))
	}

	return _language
}