// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// LanguageGetDefault returns the PangoLanguage for the current locale of the
// process.
//
// On Unix systems, this is the return value is derived from setlocale
// (LC_CTYPE, NULL), and the user can affect this through the environment
// variables LC_ALL, LC_CTYPE or LANG (checked in that order). The locale string
// typically is in the form lang_COUNTRY, where lang is an ISO-639 language
// code, and COUNTRY is an ISO-3166 country code. For instance, sv_FI for
// Swedish as written in Finland or pt_BR for Portuguese as written in Brazil.
//
// On Windows, the C library does not use any such environment variables,
// and setting them won't affect the behavior of functions like ctime().
// The user sets the locale through the Regional Options in the Control Panel.
// The C library (in the setlocale() function) does not use country and language
// codes, but country and language names spelled out in English. However,
// this function does check the above environment variables, and does return a
// Unix-style locale string based on either said environment variables or the
// thread's current locale.
//
// Your application should call setlocale(LC_ALL, "") for the user settings to
// take effect. GTK does this in its initialization functions automatically (by
// calling gtk_set_locale()). See the setlocale() manpage for more details.
//
// Note that the default language can change over the life of an application.
//
// The function returns the following values:
//
//   - language: default language as a PangoLanguage, must not be freed.
//
func LanguageGetDefault() *Language {
	var _cret *C.PangoLanguage // in

	_cret = C.pango_language_get_default()

	var _language *Language // out

	_language = (*Language)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _language
}
